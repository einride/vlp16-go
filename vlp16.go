package vlp16

import (
	"math"
)

const (
	distanceFactor   = 0.002  // 2/1000. A reported value of 51154 represents 102.308 meter
	azimuthFactor    = 0.01   // Azimuth is uint16 representing an angle in one hundredth of a degree
	maxAzimuth       = 35999  // Azimuth max value as binary
	fullFiringTime   = 55.296 // Total time for laser firings plus recharge (µs)
	singleFiringTime = 2.304  // Time for one laser firing (µs)
)

func calculateTimingOffset(returnMode ReturnMode) [32][12]float64 {
	var timingOffsets [32][12]float64
	var dataBlockIndex int
	for y, inner := range timingOffsets {
		for x := range inner {
			if returnMode == ReturnModeDualReturn {
				dataBlockIndex = (x - (x % 2)) + (y / 16)
			} else {
				dataBlockIndex = (x * 2) + (y / 16)
			}
			dataPointIndex := y % 16
			timingOffsets[y][x] = fullFiringTime*float64(dataBlockIndex) +
				singleFiringTime*float64(dataPointIndex)
		}
	}
	return timingOffsets
}

func spherical2XYZ(laserID int, azimuth uint16, distance uint16) (float64, float64, float64) {
	omega := verticalAngle(laserID)
	r := float64(distance) * distanceFactor
	alpha := deg2Rad(float64(azimuth) * azimuthFactor)

	X := r * math.Cos(omega) * math.Sin(alpha)
	Y := r * math.Cos(omega) * math.Cos(alpha)
	Z := r * math.Sin(omega)

	return X, Y, Z
}

func verticalAngle(laserID int) float64 {
	verticalAngles := [16]float64{
		deg2Rad(-15),
		deg2Rad(1),
		deg2Rad(-13),
		deg2Rad(3),
		deg2Rad(-11),
		deg2Rad(5),
		deg2Rad(-9),
		deg2Rad(7),
		deg2Rad(-7),
		deg2Rad(9),
		deg2Rad(-5),
		deg2Rad(11),
		deg2Rad(-3),
		deg2Rad(13),
		deg2Rad(-1),
		deg2Rad(15),
	}
	if laserID > 15 { // Account for second firing
		laserID -= 16
	}
	return verticalAngles[laserID]
}

func interpolateAzimuth(blockIndex int, packet *Packet) uint16 {
	// TODO: Interpolate azimuth with high precision algorithm
	// Interpolates azimuth angle by using the either the next blocks azimuth or if it's the last block
	// it uses the next to last.
	var azimuthMin uint16
	var azimuthMax uint16

	if blockIndex == len(packet.Blocks)-1 {
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

	if azimuth > maxAzimuth {
		azimuth -= 360 / azimuthFactor
	}

	return azimuth
}

func deg2Rad(degree float64) float64 {
	return degree * math.Pi / 180
}
