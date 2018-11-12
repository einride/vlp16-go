package vlp16

import (
	"time"

	"github.com/pkg/errors"
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
}

func (cloud *SphericalPointCloud) UnmarshalPacket(packet *Packet) error {
	for i := 0; i < len(packet.Blocks); i++ {
		if err := cloud.parseBlock(i, packet); err != nil {
			return errors.Wrap(err, "error parsing Block")
		}
		// Duration is in nanoseconds and Velodyne timestamp in microseconds
		cloud.Timestamp = time.Duration(packet.Timestamp) * time.Microsecond
	}
	return nil
}

func (cloud *SphericalPointCloud) parseBlock(blockIndex int, packet *Packet) error {
	azimuth := packet.Blocks[blockIndex].Azimuth
	for j := 0; j < len(packet.Blocks[0].Channels); j++ {
		if j == 16 {
			azimuth = interpolateAzimuth(blockIndex, packet)
		}
		distance := packet.Blocks[blockIndex].Channels[j].Distance
		if distance == 0 { // skip when distance = 0, invalid return from LiDAR.
			continue
		}
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

		cloud.SphericalPoints = append(cloud.SphericalPoints, point)
	}
	return nil
}
