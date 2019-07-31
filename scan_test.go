package vlp16

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestScanPackets(t *testing.T) {
	f, err := os.Open("testdata/recording.bin")
	require.NoError(t, err)
	sc := bufio.NewScanner(f)
	sc.Split(ScanPackets)
	var numPackets int
	for sc.Scan() {
		numPackets++
		require.Len(t, sc.Bytes(), lengthOfPacket)
		var rawPacket [lengthOfPacket]byte
		copy(rawPacket[:], sc.Bytes())
		var packet Packet
		packet.unmarshal(&rawPacket)
		require.NoError(t, packet.Validate())
	}
	require.Equal(t, 1000, numPackets)
}
