package vlp16

import (
	"time"
)

// PointsPerCloud is the number of points in a point cloud.
const PointsPerCloud = BlocksPerPacket * ChannelsPerBlock

type SphericalPointCloud struct {
	Timestamp       time.Duration
	SphericalPoints [PointsPerCloud]SphericalPoint
}

type SphericalPoint struct {
	Distance       float64
	Azimuth        float64
	Elevation      float64
	Reflectivity   uint8
	LastReflection bool
	TimingOffset   float64
}

func (s *SphericalPointCloud) UnmarshalPacket(packet *Packet) {
	// duration is in nanoseconds and Velodyne timestamp in microseconds
	s.Timestamp = time.Duration(packet.Timestamp) * time.Microsecond
	timingOffsets := calculateTimingOffset(packet.ReturnMode)
	for i := range packet.Blocks {
		block := &packet.Blocks[i]
		azimuth := block.Azimuth
		for j := range block.Channels {
			channel := &block.Channels[j]
			if j == 16 {
				azimuth = interpolateAzimuth(i, packet)
			}
			point := &s.SphericalPoints[i*ChannelsPerBlock+j]
			point.Distance = float64(channel.Distance) * distanceFactor
			point.Azimuth = deg2Rad(float64(azimuth) * azimuthFactor)
			point.Elevation = verticalAngle(j)
			point.Reflectivity = packet.Blocks[i].Channels[j].Reflectivity
			point.LastReflection = packet.ReturnMode == ReturnModeLastReturn ||
				// dual return mode: even number blocks (0,2,4,...) contain last return
				packet.ReturnMode == ReturnModeDualReturn && i%2 == 0
			point.TimingOffset = timingOffsets[j][i]
		}
	}
}
