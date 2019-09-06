package vlp16

import (
	"time"

	"github.com/einride/unit"
)

// PointsPerCloud is the number of points in a point cloud.
const PointsPerCloud = BlocksPerPacket * ChannelsPerBlock

type PointCloud struct {
	TimeSinceTopOfHour time.Duration
	Points             [PointsPerCloud]Point
}

func (pc *PointCloud) Azimuth

type Point struct {
	Row      uint8
	Column   uint8
	Distance unit.Distance
	// Azimuth          unit.Angle
	// Elevation        unit.Angle
	Reflectivity     uint8
	IsLastReflection bool
	TimeOffset       time.Duration
}

const (
	distanceFactor = 0.002
	azimuthFactor  = 0.01
)

func (s *PointCloud) UnmarshalPacket(packet *Packet) {
	s.TimeSinceTopOfHour = time.Duration(packet.Timestamp) * time.Microsecond
	timingOffsets := calculateTimingOffset(packet.ReturnMode)
	for i := range packet.Blocks {
		block := &packet.Blocks[i]
		azimuth := block.Azimuth
		for j := range block.Channels {
			channel := &block.Channels[j]
			if j == 16 {
				azimuth = interpolateAzimuth(i, packet)
			}
			point := &s.Points[i*ChannelsPerBlock+j]
			point.Distance = unit.Distance(channel.Distance) * distanceFactor * unit.Metre
			point.Azimuth = unit.Angle(azimuth) * azimuthFactor * unit.Radian
			point.Elevation = verticalAngle(j)
			point.Reflectivity = packet.Blocks[i].Channels[j].Reflectivity
			point.IsLastReflection = packet.ReturnMode == ReturnModeLastReturn ||
				// dual return mode: even number blocks (0,2,4,...) contain last return
				packet.ReturnMode == ReturnModeDualReturn && i%2 == 0
			point.TimeOffset = timingOffsets[j][i]
		}
	}
}
