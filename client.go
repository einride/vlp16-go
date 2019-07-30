package vlp16

import (
	"context"
	"net"
	"time"

	"golang.org/x/xerrors"
)

// UDPConn is the interface used by the Client to read UDP packets.
type UDPConn interface {
	ReadFromUDP(b []byte) (int, *net.UDPAddr, error)
	SetReadDeadline(time.Time) error
	Close() error
}

// Client is a VLP-16 client.
type Client struct {
	conn                UDPConn
	senderAddr          *net.UDPAddr
	buf                 [lengthOfPacket]byte
	packet              Packet
	sphericalPointCloud SphericalPointCloud
}

// NewClient returns a new VLP-16 client with the provided UDP connection.
func NewClient(conn UDPConn) *Client {
	return &Client{conn: conn}
}

// Receive a VLP-16 packet.
func (c *Client) Receive(ctx context.Context) error {
	deadline, _ := ctx.Deadline()
	if err := c.conn.SetReadDeadline(deadline); err != nil {
		return xerrors.Errorf("VLP-16 client: receive: %w", err)
	}
	n, senderAddr, err := c.conn.ReadFromUDP(c.buf[:])
	if err != nil {
		return xerrors.Errorf("VLP-16 client: receive: %w", err)
	}
	c.senderAddr = senderAddr
	if n != lengthOfPacket {
		return xerrors.Errorf("VLP-16 client: receive: unexpected packet length: %d (expected %d)", n, lengthOfPacket)
	}
	c.packet.unmarshal(&c.buf)
	c.sphericalPointCloud.UnmarshalPacket(&c.packet)
	return nil
}

// SenderAddr returns the sender address of the last received VLP-16 packet.
func (c *Client) SenderAddr() *net.UDPAddr {
	return c.senderAddr
}

// RawPacket returns the raw bytes of the last received VLP-16 packet.
func (c *Client) RawPacket() []byte {
	return c.buf[:]
}

// Packet returns the last received VLP-16 packet.
func (c *Client) Packet() *Packet {
	return &c.packet
}

// Packet returns the last received VLP-16 packet decoded as a spherical point cloud.
func (c *Client) SphericalPointCloud() *SphericalPointCloud {
	return &c.sphericalPointCloud
}

// Close the client's underlying UDP connection.
func (c *Client) Close() error {
	if err := c.conn.Close(); err != nil {
		return xerrors.Errorf("VLP-16 client: close: %w", err)
	}
	return nil
}
