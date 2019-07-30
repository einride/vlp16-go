package vlp16

import (
	"context"
	"net"
	"testing"
	"time"

	mockvlp16 "github.com/einride/vlp-16-go/test/mocks/vlp16"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"golang.org/x/xerrors"
)

func TestClient_Close(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	conn := mockvlp16.NewMockUDPConn(ctrl)
	client := NewClient(conn)
	err := xerrors.New("boom")
	conn.EXPECT().Close().Return(err)
	require.True(t, xerrors.Is(client.Close(), err))
}

func TestClient_Receive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	conn := mockvlp16.NewMockUDPConn(ctrl)
	client := NewClient(conn)
	addr := &net.UDPAddr{Port: 1234}
	exampleData := exampleData()[:]
	conn.EXPECT().ReadFromUDP(gomock.Any()).DoAndReturn(func(b []byte) (int, *net.UDPAddr, error) {
		require.Len(t, b, lengthOfPacket)
		copy(b, exampleData)
		return lengthOfPacket, addr, nil
	})
	deadline := time.Unix(0, 1)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()
	conn.EXPECT().SetReadDeadline(deadline)
	require.NoError(t, client.Receive(ctx))
	require.Equal(t, exampleData, client.RawPacket())
	require.Equal(t, examplePacket(), client.Packet())
	require.Equal(t, exampleSphericalPointCloud(), client.SphericalPointCloud())
}
