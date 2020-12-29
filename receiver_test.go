package vlp16

import (
	"context"
	"net"
	"testing"
	"time"

	"golang.org/x/net/nettest"
	"golang.org/x/sync/errgroup"
	"gotest.tools/v3/assert"
)

func TestClient_Receive(t *testing.T) {
	const testTimeout = time.Second
	ctx, cancel := context.WithTimeout(context.Background(), testTimeout)
	defer cancel()
	addr := getFreeAddress(t)
	rx, err := ListenUDP(ctx, addr)
	assert.NilError(t, err)
	var g errgroup.Group
	g.Go(func() error {
		return rx.Receive(ctx)
	})
	conn, err := net.Dial("udp4", addr)
	assert.NilError(t, err)
	_, err = conn.Write(exampleRawPacket(t)[:])
	assert.NilError(t, err)
	assert.NilError(t, conn.Close())
	assert.NilError(t, g.Wait())
	assert.DeepEqual(t, exampleRawPacket(t), rx.RawPacket())
	assert.DeepEqual(t, examplePacket(), rx.Packet())
}

func getFreeAddress(t *testing.T) string {
	t.Helper()
	l, err := nettest.NewLocalPacketListener("udp4")
	assert.NilError(t, err)
	defer func() {
		assert.NilError(t, l.Close())
	}()
	return l.LocalAddr().String()
}
