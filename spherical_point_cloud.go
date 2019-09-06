package vlp16

import (
	"time"

	"github.com/einride/unit"
)

const (
	PointsPerCloud   = BlocksPerPacket * ChannelsPerBlock // Number of points in a point cloud.
	ColumnsPerPacket = BlocksPerPacket * 2                // Number of columns in point cloud matrix.
	RowsPerColumn    = 16                                 // Number of rows in point cloud matrix.
)

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
	for i := range packet.Blocks {
		block := &packet.Blocks[i]
		for j := range block.Channels {
			channel := &block.Channels[j]
			point := &s.Points[i*ChannelsPerBlock+j]
			point.Column, point.Row, point.Return = indices(uint8(i), uint8(j), packet.ReturnMode)
			point.Distance = unit.Distance(channel.Distance) * distanceFactor * unit.Metre
			point.Reflectivity = packet.Blocks[i].Channels[j].Reflectivity
			point.IsLastReflection = packet.ReturnMode == ReturnModeLastReturn ||
				// dual return mode: even number blocks (0,2,4,...) contain last return
				packet.ReturnMode == ReturnModeDualReturn && i%2 == 0
			switch j {
			case 0:
				s.Azimuths[point.Column] =
					unit.Angle(float64(block.Azimuth)*azimuthFactor) * unit.Degree
			case 16:
				s.Azimuths[point.Column] =
					unit.Angle(float64(interpolateAzimuth(i, packet))*azimuthFactor) * unit.Degree
			}
		}
	}
}

// indices calculates the column, row and return index of a channel in a VLP-16 block.
func indices(blockIndex uint8, channelIndex uint8, returnMode ReturnMode) (uint8, uint8, uint8) {
	if returnMode == ReturnModeDualReturn {
		return blockIndex - (blockIndex % 2) + (channelIndex / numberOfRows),
			rowIndex(channelIndex),
			channelIndex % numberOfRows
	}
	return blockIndex*2 + channelIndex/numberOfRows, rowIndex(channelIndex), 0
}

func rowIndex(channelIndex uint8) uint8 {
	if channelIndex > numberOfRows-1 { // Account for second firing
		channelIndex -= numberOfRows
	}
	return channelIndex/2 + channelIndex%2*8
}
