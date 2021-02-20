package vlp16

import (
	"bufio"
	"os"
	"testing"

	"gotest.tools/v3/assert"
)

func newPacketRecordingScanner(t *testing.T, filename string) (*bufio.Scanner, func()) {
	t.Helper()
	f, err := os.Open(filename)
	assert.NilError(t, err)
	sc := bufio.NewScanner(f)
	sc.Split(ScanPackets)
	return sc, func() {
		assert.NilError(t, sc.Err())
		assert.NilError(t, f.Close())
	}
}
