package vlp16

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func newPacketRecordingScanner(t *testing.T, filename string) (*bufio.Scanner, func()) {
	f, err := os.Open(filename)
	require.NoError(t, err)
	sc := bufio.NewScanner(f)
	sc.Split(ScanPackets)
	return sc, func() {
		require.NoError(t, sc.Err())
		require.NoError(t, f.Close())
	}
}
