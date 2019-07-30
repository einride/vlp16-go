package vlp16

import (
	"bufio"
	"encoding/hex"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func newPacketRecordingScanner(t *testing.T, filename string) (*bufio.Scanner, func()) {
	f, err := os.Open(filename)
	require.NoError(t, err)
	sc := bufio.NewScanner(f)
	sc.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if len(data) < lengthOfPacket {
			if len(data) > 0 && atEOF {
				require.Fail(t, "remaining data: ", hex.EncodeToString(data))
			}
			return 0, nil, nil
		}
		return lengthOfPacket, data[:lengthOfPacket], nil
	})
	return sc, func() {
		require.NoError(t, sc.Err())
		require.NoError(t, f.Close())
	}
}
