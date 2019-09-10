package vlp16

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"golang.org/x/net/nettest"
	"golang.org/x/sync/errgroup"
)

func TestClient_Receive(t *testing.T) {
	const testTimeout = time.Second
	ctx, cancel := context.WithTimeout(context.Background(), testTimeout)
	defer cancel()
	addr := getFreeAddress(t)
	rx, err := ListenUDP(ctx, addr)
	require.NoError(t, err)
	var g errgroup.Group
	g.Go(func() error {
		return rx.Receive(ctx)
	})
	conn, err := net.Dial("udp4", addr)
	require.NoError(t, err)
	_, err = conn.Write(exampleData()[:])
	require.NoError(t, err)
	require.NoError(t, conn.Close())
	require.NoError(t, g.Wait())
	require.Equal(t, exampleData()[:], rx.RawPacket())
	require.Equal(t, examplePacket(), rx.Packet())
	requirePointCloudEqual(t, examplePointCloud(), rx.PointCloud())
}

func getFreeAddress(t *testing.T) string {
	t.Helper()
	l, err := nettest.NewLocalPacketListener("udp4")
	require.NoError(t, err)
	defer func() {
		require.NoError(t, l.Close())
	}()
	return l.LocalAddr().String()
}
