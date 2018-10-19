package vlp16

import (
	"log"
	"math"
	"net"
	"time"

	"github.com/pkg/errors"
)

type SphericalPointCloud struct {
	SphericalPoints []SphericalPoint
	Timestamp       time.Time
}

type SphericalPoint struct {
	Distance     float64
	Azimuth      float64
	Omega        float64
	Reflectivity Reflectivity
}

type CartesianPointCloud struct {
	CartesianPoints []CartesianPoint
	TimeStamp       time.Time
}

type CartesianPoint struct {
	X            float64
	Y            float64
	Z            float64
	Reflectivity Reflectivity
}

const (
	distanceFactor = 0.002 // 2/1000. A reported value of 51154 represents 102.308 meter
	azimuthFactor  = 0.01  // Azimuth is uint16 representing an angle in one hundredth of a degree
	maxAzimuth     = 35999 // Azimuth max value as binary
)

var verticalAngles = [16]float64{
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
	deg2Rad(15)}

func InitVLP16(conn *net.UDPConn) error {

	packet := Packet{}
	for {
		if err := packet.Read(conn); err != nil {
			log.Printf("Error reading from connection. %v", err)
		}

		if err := parsePacket(&packet); err != nil {
			return errors.Wrap(err, "Error parsing packet")
		}
	}
}

func parsePacket(packet *Packet) error {
	for i := 0; i < len(packet.Blocks); i++ {
		if err := parseBlock(i, packet); err != nil {
			return errors.Wrap(err, "Error parsing Block")
		}
	}
	return nil
}

func parseBlock(blockIndex int, packet *Packet) error {
	azimuth := packet.Blocks[blockIndex].Azimuth
	for j := 0; j < len(packet.Blocks[0].Channels); j++ {
		if j == 16 {
			azimuth = interpolateAzimuth(blockIndex, packet)
		}

		distance := packet.Blocks[blockIndex].Channels[j].Distance
		if distance == 0 { // skip when distance = 0, invalid return from LiDAR.
			continue
		}
		X, Y, Z := spherical2XYZ(j, azimuth, distance)
		log.Printf("X: %v, Y: %v, Z: %v, Reflectivity: %v", X, Y, Z, packet.Blocks[blockIndex].Channels[j].Reflectivity)
	}
	return nil
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
	if laserID > 15 { // Account for second firing
		laserID -= 16
	}
	omega := verticalAngles[laserID]
	return omega
}

func interpolateAzimuth(blockIndex int, packet *Packet) uint16 {
	/*
		Interpolates azimuth angle by using the either the next blocks azimuth or if it's the last block
		it uses the next to last.
	*/
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
