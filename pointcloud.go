package vlp16

import (
	"time"

	"github.com/einride/unit"
)

const (
	// PointsPerCloud is the number of points in a full firing point cloud.
	PointsPerCloud = BlocksPerPacket * ChannelsPerBlock
	// CloumnsPerPacket is the number of columns in point cloud matrix.
	ColumnsPerPacket = BlocksPerPacket * 2
	// RowsPerColumnSingleReturn is the number of rows in point cloud in single return mode.
	RowsPerColumnSingleReturn = 16
	// FullFiringTime is the total time for laser firings plus recharge (55.296 µs).
	FullFiringTime = 55296 * time.Nanosecond
	// SingleFiringTime is the time for one laser firing (2.304 µs).
	SingleFiringTime = 2304 * time.Nanosecond
	// RechargeTime is the recharge time between laser firings (18.432 µs).
	RechargeTime = 18432 * time.Nanosecond
	// LowestElevation is the elevation angle of first row of measurements.
	LowestElevation = -15 * unit.Degree
	// DeltaElevation is the angle difference between two rows.
	DeltaElevation = 2 * unit.Degree
)

// compile-time assertion on full firing time.
var _ [FullFiringTime]struct{} = [RowsPerColumnSingleReturn*SingleFiringTime + RechargeTime]struct{}{}

type PointCloud struct {
	TimeSinceTopOfHour time.Duration
	Points             [PointsPerCloud]Point
	Azimuths           [ColumnsPerPacket]unit.Angle
}

type Point struct {
	Row              uint8
	Column           uint8
	Return           uint8
	Reflectivity     uint8
	IsLastReflection bool
	Distance         unit.Distance
}

const (
	distanceFactor = 0.002
	azimuthFactor  = 0.01
)

func (s *PointCloud) UnmarshalPacket(packet *Packet) {
	s.TimeSinceTopOfHour = time.Duration(packet.Timestamp) * time.Microsecond
	for blockIndex := range packet.Blocks {
		block := &packet.Blocks[blockIndex]
		for channelIndex := range block.Channels {
			channel := &block.Channels[channelIndex]
			point := &s.Points[blockIndex*ChannelsPerBlock+channelIndex]
			if packet.ReturnMode == ReturnModeDualReturn {
				point.Column = uint8(blockIndex - (blockIndex % 2) + (channelIndex / RowsPerColumnSingleReturn))
				point.Return = uint8(channelIndex % RowsPerColumnSingleReturn)
			} else {
				point.Column = uint8(blockIndex*2 + channelIndex/RowsPerColumnSingleReturn)
				point.Return = 0
			}
			if channelIndex >= RowsPerColumnSingleReturn { // account for second firing
				point.Row = uint8((channelIndex-RowsPerColumnSingleReturn)/2 +
					(channelIndex-RowsPerColumnSingleReturn)%2*8)
			} else {
				point.Row = uint8(channelIndex/2 + channelIndex%2*8)
			}
			point.Distance = unit.Distance(channel.Distance) * distanceFactor * unit.Metre
			point.Reflectivity = packet.Blocks[blockIndex].Channels[channelIndex].Reflectivity
			point.IsLastReflection = packet.ReturnMode == ReturnModeLastReturn ||
				// dual return mode: even number blocks (0,2,4,...) contain last return
				packet.ReturnMode == ReturnModeDualReturn && blockIndex%2 == 0
			switch channelIndex {
			case 0:
				s.Azimuths[point.Column] =
					unit.Angle(float64(block.Azimuth)*azimuthFactor) * unit.Degree
			case 16:
				s.Azimuths[point.Column] =
					unit.Angle(float64(interpolateAzimuth(uint8(blockIndex), packet))*azimuthFactor) * unit.Degree
			}
		}
	}
}

func interpolateAzimuth(blockIndex uint8, packet *Packet) uint16 {
	// TODO: Interpolate azimuth with high precision algorithm
	// Interpolates azimuth angle by using the either the next blocks azimuth or if it's the last block
	// it uses the next to last.
	var azimuthMin uint16
	var azimuthMax uint16
	if blockIndex == uint8(len(packet.Blocks)-1) {
		azimuthMin = packet.Blocks[blockIndex-1].Azimuth
		azimuthMax = packet.Blocks[blockIndex].Azimuth
	} else {
		azimuthMin = packet.Blocks[blockIndex].Azimuth
		azimuthMax = packet.Blocks[blockIndex+1].Azimuth
	}
	if azimuthMax < azimuthMin { // If the bigger angle has gone over 360 degrees (35999 -> 0)
		azimuthMax += 360 / azimuthFactor
	}
	azimuth := packet.Blocks[blockIndex].Azimuth + (azimuthMax-azimuthMin)/2
	const maxAzimuth = 35999 // Azimuth max value as binary
	if azimuth > maxAzimuth {
		azimuth -= 360 / azimuthFactor
	}
	return azimuth
}
