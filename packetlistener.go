package vlp16

import (
	"context"
	"fmt"
	"net"
	"time"

	"golang.org/x/net/ipv4"
)

// ListenUDP listens for VLP-16 UDP packets on the specified address.
func ListenUDP(ctx context.Context, addr string, listenOpts ...ListenOption) (_ *PacketListener, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("listen VLP-16 UDP packets: %w", err)
		}
	}()
	opts := defaultListenOptions()
	for _, receiverOpt := range listenOpts {
		receiverOpt(opts)
	}
	var listenConfig net.ListenConfig
	packetConn, err := listenConfig.ListenPacket(ctx, "udp4", addr)
	if err != nil {
		return nil, err
	}
	udpConn := packetConn.(*net.UDPConn)
	if err := udpConn.SetReadBuffer(opts.bufferSizeBytes); err != nil {
		return nil, err
	}
	conn := ipv4.NewPacketConn(udpConn)
	c := &PacketListener{conn: conn}
	// allocate memory for batch reads
	c.messages = make([]ipv4.Message, 0, opts.batchSize)
	for i := 0; i < opts.batchSize; i++ {
		c.rawPacketBuf = append(c.rawPacketBuf, &RawPacket{})
		c.messages = append(c.messages, ipv4.Message{
			Buffers: [][]byte{c.rawPacketBuf[i][:]},
		})
	}
	return c, nil
}

// PacketListener listens for VLP-16 packets.
type PacketListener struct {
	conn             *ipv4.PacketConn
	messages         []ipv4.Message
	rawPacketBuf     []*RawPacket
	messageBufSize   int
	currMessageIndex int
	currPacket       Packet
}

// ReadPacket reads a VLP-16 packet.
func (c *PacketListener) ReadPacket() error {
	c.currMessageIndex++
	if c.currMessageIndex >= c.messageBufSize {
		c.currMessageIndex = 0
		n, err := c.conn.ReadBatch(c.messages, 0)
		if err != nil {
			return fmt.Errorf("read VLP-16 packet: %w", err)
		}
		c.messageBufSize = n
	}
	c.currPacket.UnmarshalRawPacket(c.rawPacketBuf[c.currMessageIndex])
	return nil
}

// SetReadDeadline sets the read deadline associated with the listener's underlying UDP connection.
func (c *PacketListener) SetReadDeadline(deadline time.Time) error {
	if err := c.conn.SetReadDeadline(deadline); err != nil {
		return fmt.Errorf("set VLP-16 listener read deadline: %w", err)
	}
	return nil
}

// SourceIP returns the source IP of the last received VLP-16 packet.
func (c *PacketListener) SourceIP() net.IP {
	return c.messages[c.currMessageIndex].Addr.(*net.UDPAddr).IP
}

// RawPacket returns the raw, unprocessed last received VLP-16 packet.
func (c *PacketListener) RawPacket() *RawPacket {
	return c.rawPacketBuf[c.currMessageIndex]
}

// Packet returns the last received VLP-16 packet.
func (c *PacketListener) Packet() *Packet {
	return &c.currPacket
}

// Close the client's underlying UDP connection.
func (c *PacketListener) Close() error {
	if err := c.conn.Close(); err != nil {
		return fmt.Errorf("close VLP-16 listener: %w", err)
	}
	return nil
}
