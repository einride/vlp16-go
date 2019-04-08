package vlp16

import (
	"time"
)

type SphericalPointCloud struct {
	SphericalPoints []SphericalPoint
	Timestamp       time.Duration
}

type SphericalPoint struct {
	Distance       float64
	Azimuth        float64
	Elevation      float64
	Reflectivity   Reflectivity
	LastReflection bool
	TimingOffset   float64
}

func (cloud *SphericalPointCloud) UnmarshalPacket(packet *Packet) {
	for i := 0; i < len(packet.Blocks); i++ {
		cloud.parseBlock(i, packet)
		// Duration is in nanoseconds and Velodyne timestamp in microseconds
		cloud.Timestamp = time.Duration(packet.Timestamp) * time.Microsecond
	}
}

func (cloud *SphericalPointCloud) parseBlock(blockIndex int, packet *Packet) {
	azimuth := packet.Blocks[blockIndex].Azimuth
	timingOffsets := calculateTimingOffset(packet.ReturnMode)
	for j := 0; j < len(packet.Blocks[0].Channels); j++ {
		if j == 16 {
			azimuth = interpolateAzimuth(blockIndex, packet)
		}
		distance := packet.Blocks[blockIndex].Channels[j].Distance
		lastReturn := false
		switch packet.ReturnMode {
		case ReturnModeDualReturn:
			// Even number blocks (0,2,4,...) contain last return
			if blockIndex%2 == 0 {
				lastReturn = true
			}
		case ReturnModeLastReturn:
			lastReturn = true
		}

		point := SphericalPoint{}
		point.Distance = float64(distance) * distanceFactor
		point.Azimuth = deg2Rad(float64(azimuth) * azimuthFactor)
		point.Elevation = verticalAngle(j)
		point.Reflectivity = packet.Blocks[blockIndex].Channels[j].Reflectivity
		point.LastReflection = lastReturn
		point.TimingOffset = timingOffsets[j][blockIndex]

		cloud.SphericalPoints = append(cloud.SphericalPoints, point)
	}
}