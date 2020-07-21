package vlp16

import (
	"testing"
	"time"

	"gotest.tools/v3/assert"
)

func TestPacket_Unmarshal_TestData(t *testing.T) {
	sc, done := newPacketRecordingScanner(t, "testdata/recording.bin")
	defer done()
	var numPackets int
	for sc.Scan() {
		var data [lengthOfPacket]byte
		copy(data[:], sc.Bytes())
		var packet Packet
		packet.unmarshal(&data)
		numPackets++
		assert.Equal(t, ReturnModeStrongest, packet.ReturnMode)
		assert.Equal(t, ProductIDVLP16, packet.ProductID)
		assert.Assert(t,
			time.Duration(packet.Timestamp)*time.Microsecond <= time.Hour,
			"timestamp should be <= number of microseconds in an hour",
		)
	}
	assert.Equal(t, 1000, numPackets)
}

func BenchmarkPacket_Unmarshal_Example(b *testing.B) {
	packet := &Packet{}
	exampleData := exampleData()
	for i := 0; i < b.N; i++ {
		packet.unmarshal(exampleData)
	}
}

func TestPacket_Unmarshal_Example(t *testing.T) {
	actual := &Packet{}
	actual.unmarshal(exampleData())
	assert.DeepEqual(t, examplePacket(), actual)
}
