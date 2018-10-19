package vlp16

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	filename = "../../test/testdata"
)

func TestSimulateRead(t *testing.T) {
	test := assert.New(t)
	fileContents, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatalf("Error reading file, %v", err)
	}
	stats := Statistics{}
	for i := 0; i < len(fileContents)/packetSize; i++ {
		packet := fileContents[i*packetSize : (i+1)*packetSize]
		err := ParsePacket(packet, &stats)
		if err != nil {
			test.Error(err)
		}
	}
	fmt.Printf("packets: %v:, points: %v", stats.packetsCount, stats.pointsCount)
}

func TestInterpolateAzimuth(t *testing.T) {
	test := assert.New(t)
	fileContents, err := ioutil.ReadFile(filename)

	if err != nil {
		test.Errorf(err, "Error reading file")
	}

	for packetIndex := 0; packetIndex < len(fileContents)/packetSize; packetIndex++ {
		packet := fileContents[packetIndex*packetSize : (packetIndex+1)*packetSize]
		for blockIndex := 0; blockIndex < dataBlockCount; blockIndex++ {
			block := packet[dataBlockSize*blockIndex : dataBlockSize*(blockIndex+1)]
			azimuth := binary.LittleEndian.Uint16(block[2:4])
			for j := 0; j < channelsCount; j++ {
				if j == 16 {
					newAzimuth := InterpolateAzimuth(azimuth, blockIndex, packet)
					var testBool bool
					if azimuth > 35980 { // Approximate that holds as of now...
						testBool = azimuth > newAzimuth
					} else {
						testBool = newAzimuth > azimuth
					}
					test.True(testBool)
				}
			}
		}
	}
}

func TestVerticalAngle(t *testing.T) {
	test := assert.New(t)
	verticalAngles = [16]float64{
		Deg2Rad(-15),
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

	test.Equal(Deg2Rad(-15), VerticalAngle(0))
	test.Equal(Deg2Rad(15), VerticalAngle(15))
	test.Equal(Deg2Rad(7), VerticalAngle(7))
}

func TestSpherical2XYZ(t *testing.T) {
	test := assert.New(t)
	laserID := 15                          // 15 degrees
	azimuth := uint16(9000)                // 90 degrees
	distance := uint16(1 / distanceFactor) // 1 meter
	X, Y, Z := Spherical2XYZ(laserID, azimuth, distance)
	test.InEpsilon(X, 0.965925826289068, 0.000001)
	test.InEpsilon(Y, 5.914589856893349e-17, 0.000001)
	test.InEpsilon(Z, 0.258819045102521, 0.000001)
}

func TestDeg2Rad(t *testing.T) {
	test := assert.New(t)
	test.Equal(Deg2Rad(0.0), 0.0)
	test.Equal(Deg2Rad(90), math.Pi/2)
	test.Equal(Deg2Rad(180), math.Pi)
}
