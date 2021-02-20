package vlp16

import (
	"bufio"
	"os"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestScanPackets(t *testing.T) {
	t.Parallel()
	f, err := os.Open("testdata/recording.bin")
	assert.NilError(t, err)
	sc := bufio.NewScanner(f)
	sc.Split(ScanPackets)
	var numPackets int
	for sc.Scan() {
		numPackets++
		assert.Assert(t, is.Len(sc.Bytes(), SizeOfPacket))
		var rawPacket RawPacket
		copy(rawPacket[:], sc.Bytes())
		var packet Packet
		packet.UnmarshalRawPacket(&rawPacket)
		assert.NilError(t, packet.Validate())
	}
	assert.Equal(t, 1000, numPackets)
}
