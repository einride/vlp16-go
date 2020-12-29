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
		var rawPacket RawPacket
		copy(rawPacket[:], sc.Bytes())
		var packet Packet
		packet.UnmarshalRawPacket(&rawPacket)
		numPackets++
		assert.Equal(t, ReturnModeStrongest, packet.ReturnMode)
		assert.Equal(t, ProductIDVLP16, packet.ProductID)
		assert.Assert(
			t,
			time.Duration(packet.Timestamp)*time.Microsecond <= time.Hour,
			"timestamp should be <= number of microseconds in an hour",
		)
	}
	assert.Equal(t, 1000, numPackets)
}

func BenchmarkPacket_Unmarshal_Example(b *testing.B) {
	packet := &Packet{}
	exampleData := exampleRawPacket(b)
	for i := 0; i < b.N; i++ {
		packet.UnmarshalRawPacket(exampleData)
	}
}

func TestPacket_Unmarshal_Example(t *testing.T) {
	actual := &Packet{}
	actual.UnmarshalRawPacket(exampleRawPacket(t))
	assert.DeepEqual(t, examplePacket(), actual)
}
