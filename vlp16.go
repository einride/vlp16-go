package vlp16

import (
	"math"
	"time"

	"github.com/einride/unit"
)

const (
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

	numberOfRows = 16
	maxAzimuth   = 35999 // Azimuth max value as binary
)

// compile-time assertion on full firing time.
var _ [FullFiringTime]struct{} = [numberOfRows*SingleFiringTime + RechargeTime]struct{}{}

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
			timingOffsets[y][x] = FullFiringTime*float64(dataBlockIndex) +
				SingleFiringTime*float64(dataPointIndex)
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

func verticalAngle(channelIndex int) unit.Angle {
	laserIndex := channelIndex
	if channelIndex > 15 { // account for second firing
		laserIndex = -16
	}
	return [16]unit.Angle{
		-15 * unit.Degree,
		1 * unit.Degree,
		-13 * unit.Degree,
		3 * unit.Degree,
		-11 * unit.Degree,
		5 * unit.Degree,
		-9 * unit.Degree,
		7 * unit.Degree,
		-7 * unit.Degree,
		9 * unit.Degree,
		-5 * unit.Degree,
		11 * unit.Degree,
		-3 * unit.Degree,
		13 * unit.Degree,
		-1 * unit.Degree,
		15 * unit.Degree,
	}[laserIndex]
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
