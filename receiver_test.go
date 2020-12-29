package vlp16

import (
	"context"
	"net"
	"sync"
	"testing"
	"time"

	"golang.org/x/net/nettest"
	"gotest.tools/v3/assert"
)

func TestClient_Receive(t *testing.T) {
	const testTimeout = time.Second
	ctx, cancel := context.WithTimeout(context.Background(), testTimeout)
	defer cancel()
	addr := getFreeAddress(t)
	rx, err := ListenUDP(ctx, addr)
	assert.NilError(t, err)
	var g sync.WaitGroup
	g.Add(1)
	go func() {
		defer g.Done()
		if err := rx.Receive(ctx); err != nil {
			t.Error(err)
		}
	}()
	conn, err := net.Dial("udp4", addr)
	assert.NilError(t, err)
	_, err = conn.Write(exampleRawPacket(t)[:])
	assert.NilError(t, err)
	assert.NilError(t, conn.Close())
	g.Wait()
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
