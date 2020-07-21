package vlp16

import (
	"math"
	"testing"

	"gotest.tools/v3/assert"
)

func TestSphericalPointCloud_UnmarshalExamplePacket(t *testing.T) {
	actual := &PointCloud{}
	actual.UnmarshalPacket(examplePacket())
	requirePointCloudEqual(t, examplePointCloud(), actual)
}

func requirePointCloudEqual(t *testing.T, p *PointCloud, pc *PointCloud) {
	const delta = 1e-5
	assert.Equal(t, p.TimeSinceTopOfHour, pc.TimeSinceTopOfHour)
	for i := range p.Azimuths {
		assert.Assert(t, math.Abs(p.Azimuths[i].Radians()-pc.Azimuths[i].Radians()) < delta)
	}
	for i := range p.Points {
		assert.Assert(t, math.Abs(p.Points[i].Distance.Metres()-pc.Points[i].Distance.Metres()) < delta)
		assert.Equal(t, p.Points[i].Column, pc.Points[i].Column)
		assert.Equal(t, p.Points[i].Row, pc.Points[i].Row)
		assert.Equal(t, p.Points[i].Reflectivity, pc.Points[i].Reflectivity)
		assert.Equal(t, p.Points[i].IsLastReflection, pc.Points[i].IsLastReflection)
	}
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
					newAzimuth := interpolateAzimuth(uint8(blockIndex), &packet)
					var testBool bool
					if azimuth > 35980 { // approximate that holds as of now...
						testBool = azimuth > newAzimuth
					} else {
						testBool = newAzimuth > azimuth
					}
					assert.Assert(t, testBool)
				}
			}
		}
	}
}
