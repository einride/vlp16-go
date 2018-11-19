package vlp16

import (
	"io"
	"math"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testDataFile = "test/testdata"
)

func TestCalculateTimeOffset(t *testing.T) {
	test := assert.New(t)

	eps := 0.005

	// Expected values taken from Figure 9-9 Single Return Mode Timing Offsets (µs), VLP-16 Manual
	timingOffsetsLastReturn := calculateTimingOffset(ReturnModeLastReturn)
	test.InDelta(0.0, timingOffsetsLastReturn[0][0], eps)
	test.InDelta(34.560, timingOffsetsLastReturn[15][0], eps)
	test.InDelta(55.296, timingOffsetsLastReturn[16][0], eps)
	test.InDelta(89.856, timingOffsetsLastReturn[31][0], eps)
	test.InDelta(110.592, timingOffsetsLastReturn[0][1], eps)
	test.InDelta(145.152, timingOffsetsLastReturn[15][1], eps)
	test.InDelta(165.888, timingOffsetsLastReturn[16][1], eps)
	test.InDelta(200.448, timingOffsetsLastReturn[31][1], eps)
	test.InDelta(1306.37, timingOffsetsLastReturn[31][11], eps)

	// Expected values taken from Figure 9-10 Dual Return Mode Timing Offsets (µs), VLP-16 Manual
	timingOffsetsDualReturn := calculateTimingOffset(ReturnModeDualReturn)
	test.InDelta(0.0, timingOffsetsDualReturn[0][0], eps)
	test.InDelta(34.560, timingOffsetsDualReturn[15][0], eps)
	test.InDelta(55.296, timingOffsetsDualReturn[16][0], eps)
	test.InDelta(89.856, timingOffsetsDualReturn[31][0], eps)
	test.InDelta(0.0, timingOffsetsDualReturn[0][1], eps)
	test.InDelta(34.560, timingOffsetsDualReturn[15][1], eps)
	test.InDelta(55.296, timingOffsetsDualReturn[16][1], eps)
	test.InDelta(89.856, timingOffsetsDualReturn[31][1], eps)
	test.InDelta(642.816, timingOffsetsDualReturn[31][11], eps)
}

func TestInterpolateAzimuth(t *testing.T) {
	test := assert.New(t)
	testData, err := os.Open(testDataFile)
	require.NoError(t, err)

	var packet Packet
	for {
		err = packet.Read(testData)
		if err == io.EOF {
			break
		}
		for blockIndex := 0; blockIndex < len(packet.Blocks); blockIndex++ {
			block := packet.Blocks[blockIndex]
			azimuth := block.Azimuth
			for j := 0; j < len(packet.Blocks[0].Channels); j++ {
				if j == 16 {
					newAzimuth := interpolateAzimuth(blockIndex, &packet)
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

	test.Equal(deg2Rad(-15), verticalAngle(0))
	test.Equal(deg2Rad(15), verticalAngle(15))
	test.Equal(deg2Rad(7), verticalAngle(7))
}

func TestSpherical2XYZ(t *testing.T) {
	test := assert.New(t)
	laserID := 15                          // 15 degrees
	azimuth := uint16(9000)                // 90 degrees
	distance := uint16(1 / distanceFactor) // 1 meter
	X, Y, Z := spherical2XYZ(laserID, azimuth, distance)
	test.InEpsilon(X, 0.965925826289068, 0.000001)
	test.InEpsilon(Y, 5.914589856893349e-17, 0.000001)
	test.InEpsilon(Z, 0.258819045102521, 0.000001)
}

func TestDeg2Rad(t *testing.T) {
	test := assert.New(t)
	test.Equal(deg2Rad(0.0), 0.0)
	test.Equal(deg2Rad(90), math.Pi/2)
	test.Equal(deg2Rad(180), math.Pi)
}
