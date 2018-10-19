package vlp16

import (
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"net"

	"github.com/pkg/errors"
)

const (
	protocol       = "udp"
	packetSize     = 1206
	dataBlockCount = 12
	dataBlockSize  = 100
	channelsCount  = 32
	distanceFactor = 0.002 // 2/1000. A reported value of 51154 represents 102.308 meter
	azimuthFactor  = 0.01  // Azimuth is uint16 representing an angle in one hundredth of a degree
	maxAzimuth     = 35999 // Azimuth max value as binary
)

var verticalAngles = [16]float64{Deg2Rad(-15),
	Deg2Rad(1),
	Deg2Rad(-13),
	Deg2Rad(3),
	Deg2Rad(-11),
	Deg2Rad(5),
	Deg2Rad(-9),
	Deg2Rad(7),
	Deg2Rad(-7),
	Deg2Rad(9),
	Deg2Rad(-5),
	Deg2Rad(11),
	Deg2Rad(-3),
	Deg2Rad(13),
	Deg2Rad(-1),
	Deg2Rad(15)}

type Statistics struct {
	packetsCount uint64
	pointsCount  uint64
}

func InitVLP16(port int) error {
	stats := Statistics{}
	udpAddr, err := net.ResolveUDPAddr(protocol, fmt.Sprintf(":%v", port))
	if err != nil {
		return errors.Wrap(err, "Couldn't resolve address")
	}

	conn, err := net.ListenUDP(protocol, udpAddr)
	if err != nil {
		return errors.Wrap(err, "Couldn't connect")
	}

	packet := make([]byte, packetSize)

	for {
		err := binary.Read(conn, binary.LittleEndian, &packet)
		if err != nil {
			log.Printf("Error reading from connection. %v", err) // TODO: Don't want to return. Better solution?
		}
		err = ParsePacket(packet, &stats)
		if err != nil {
			return errors.Wrap(err, "Error parsing packet")
		}
	}
}

func ParsePacket(packet []byte, statistics *Statistics) error {
	statistics.packetsCount++
	for i := 0; i < dataBlockCount; i++ {
		err := parseBlock(packet[dataBlockSize*i:dataBlockSize*(i+1)], i, packet, statistics)
		if err != nil {
			return errors.Wrap(err, "Error parsing block")
		}
	}
	return nil
}

func parseBlock(block []byte, blockIndex int, packet []byte, statistics *Statistics) error {
	if len(block) != dataBlockSize {
		return errors.New(fmt.Sprintf("Block size incorrenct, len %v != %v", len(block), dataBlockSize))
	}

	azimuth := binary.LittleEndian.Uint16(block[2:4])
	dataPoints := block[4:dataBlockSize]
	for j := 0; j < channelsCount; j++ {
		if j == 16 {
			azimuth = InterpolateAzimuth(azimuth, blockIndex, packet)
		}

		distance := binary.LittleEndian.Uint16(dataPoints[j*3 : j*3+2])
		if distance == 0 { // skip when distance = 0, invalid return from LiDAR.
			continue
		}

		reflectivity := dataPoints[j*3+2]
		X, Y, Z := Spherical2XYZ(j, azimuth, distance)
		statistics.pointsCount++

		// TODO: Fix output
		if azimuth > 0 && azimuth < 500 && j == 11 {
			fmt.Printf("X: %3.3f, Y: %3.3f, Z: %3.3f, reflectivity: %v\n", X, Y, Z, reflectivity)
		}
	}
	return nil
}

func Spherical2XYZ(laserID int, azimuth uint16, distance uint16) (float64, float64, float64) {
	omega := VerticalAngle(laserID)
	r := float64(distance) * distanceFactor
	alpha := Deg2Rad(float64(azimuth) * azimuthFactor)

	X := r * math.Cos(omega) * math.Sin(alpha)
	Y := r * math.Cos(omega) * math.Cos(alpha)
	Z := r * math.Sin(omega)

	return X, Y, Z
}

func VerticalAngle(laserID int) float64 {
	if laserID > 15 { // Account for second firing
		laserID -= 16
	}
	omega := verticalAngles[laserID]
	return omega
}

func InterpolateAzimuth(azimuth uint16, blockIndex int, packet []byte) uint16 {
	/*
		Interpolates azimuth angle by using the either the next blocks azimuth or if it's the last block
		it uses the next to last.
	*/
	var azimuthMin uint16
	var azimuthMax uint16
	if blockIndex == dataBlockCount-1 {
		azimuthMin = binary.LittleEndian.Uint16(packet[(blockIndex-1)*dataBlockSize+2 : (blockIndex-1)*dataBlockSize+4])
		azimuthMax = azimuth
	} else {
		azimuthMin = azimuth
		azimuthMax = binary.LittleEndian.Uint16(packet[(blockIndex+1)*dataBlockSize+2 : (blockIndex+1)*dataBlockSize+4])
	}

	if azimuthMax < azimuthMin { // If the bigger angle has gone over 360 degrees (35999 -> 0)
		azimuthMax += 360 / azimuthFactor
	}

	azimuth += (azimuthMax - azimuthMin) / 2

	if azimuth > maxAzimuth {
		azimuth -= 360 / azimuthFactor
	}

	return azimuth
}

func Deg2Rad(degree float64) float64 {
	return degree * math.Pi / 180
}
