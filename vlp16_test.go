package vlp16

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculateTimeOffset(t *testing.T) {
	eps := 0.005
	// Expected values taken from Figure 9-9 Single Return Mode Timing Offsets (µs), VLP-16 Manual
	timingOffsetsLastReturn := calculateTimingOffset(ReturnModeLastReturn)
	require.InDelta(t, 0.0, timingOffsetsLastReturn[0][0], eps)
	require.InDelta(t, 34.560, timingOffsetsLastReturn[15][0], eps)
	require.InDelta(t, 55.296, timingOffsetsLastReturn[16][0], eps)
	require.InDelta(t, 89.856, timingOffsetsLastReturn[31][0], eps)
	require.InDelta(t, 110.592, timingOffsetsLastReturn[0][1], eps)
	require.InDelta(t, 145.152, timingOffsetsLastReturn[15][1], eps)
	require.InDelta(t, 165.888, timingOffsetsLastReturn[16][1], eps)
	require.InDelta(t, 200.448, timingOffsetsLastReturn[31][1], eps)
	require.InDelta(t, 1306.37, timingOffsetsLastReturn[31][11], eps)
	// Expected values taken from Figure 9-10 Dual Return Mode Timing Offsets (µs), VLP-16 Manual
	timingOffsetsDualReturn := calculateTimingOffset(ReturnModeDualReturn)
	require.InDelta(t, 0.0, timingOffsetsDualReturn[0][0], eps)
	require.InDelta(t, 34.560, timingOffsetsDualReturn[15][0], eps)
	require.InDelta(t, 55.296, timingOffsetsDualReturn[16][0], eps)
	require.InDelta(t, 89.856, timingOffsetsDualReturn[31][0], eps)
	require.InDelta(t, 0.0, timingOffsetsDualReturn[0][1], eps)
	require.InDelta(t, 34.560, timingOffsetsDualReturn[15][1], eps)
	require.InDelta(t, 55.296, timingOffsetsDualReturn[16][1], eps)
	require.InDelta(t, 89.856, timingOffsetsDualReturn[31][1], eps)
	require.InDelta(t, 642.816, timingOffsetsDualReturn[31][11], eps)
}

func TestInterpolateAzimuth(t *testing.T) {
	sc, done := newPacketRecordingScanner(t, "testdata/recording.bin")
	defer done()
	for sc.Scan() {
		var data [lengthOfPacket]byte
		copy(data[:], sc.Bytes())
		var packet Packet
		packet.unmarshal(&data)
		for blockIndex := 0; blockIndex < len(packet.Blocks); blockIndex++ {
			block := packet.Blocks[blockIndex]
			azimuth := block.Azimuth
			for j := 0; j < len(packet.Blocks[0].Channels); j++ {
				if j == 16 {
					newAzimuth := interpolateAzimuth(blockIndex, &packet)
					var testBool bool
					if azimuth > 35980 { // approximate that holds as of now...
						testBool = azimuth > newAzimuth
					} else {
						testBool = newAzimuth > azimuth
					}
					require.True(t, testBool)
				}
			}
		}
	}
}

func TestVerticalAngle(t *testing.T) {
	require.Equal(t, deg2Rad(-15), verticalAngle(0))
	require.Equal(t, deg2Rad(15), verticalAngle(15))
	require.Equal(t, deg2Rad(7), verticalAngle(7))
}

func TestSpherical2XYZ(t *testing.T) {
	laserID := 15                          // 15 degrees
	azimuth := uint16(9000)                // 90 degrees
	distance := uint16(1 / distanceFactor) // 1 meter
	X, Y, Z := spherical2XYZ(laserID, azimuth, distance)
	require.InEpsilon(t, X, 0.965925826289068, 0.000001)
	require.InEpsilon(t, Y, 5.914589856893349e-17, 0.000001)
	require.InEpsilon(t, Z, 0.258819045102521, 0.000001)
}

func TestDeg2Rad(t *testing.T) {
	require.Equal(t, deg2Rad(0.0), 0.0)
	require.Equal(t, deg2Rad(90), math.Pi/2)
	require.Equal(t, deg2Rad(180), math.Pi)
}
