package vlp16

import (
	"context"
	"fmt"
	"net"

	"golang.org/x/net/ipv4"
)

// ListenUDP listens for VLP-16 UDP packets on the specified address.
func ListenUDP(ctx context.Context, addr string, receiverOpts ...ReceiverOption) (*Receiver, error) {
	opts := defaultReceiverOptions()
	for _, receiverOpt := range receiverOpts {
		receiverOpt(opts)
	}
	var listenConfig net.ListenConfig
	packetConn, err := listenConfig.ListenPacket(ctx, "udp4", addr)
	if err != nil {
		return nil, fmt.Errorf("VLP-16: listen UDP: %w", err)
	}
	udpConn := packetConn.(*net.UDPConn)
	if err := udpConn.SetReadBuffer(opts.bufferSizeBytes); err != nil {
		return nil, fmt.Errorf("VLP-16: listen UDP: %w", err)
	}
	conn := ipv4.NewPacketConn(udpConn)
	c := &Receiver{conn: conn}
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

// Receiver receives VLP-16 packets.
type Receiver struct {
	conn             *ipv4.PacketConn
	messages         []ipv4.Message
	rawPacketBuf     []*RawPacket
	messageBufSize   int
	currMessageIndex int
	currPacket       Packet
}

// Receive a VLP-16 packet.
func (c *Receiver) Receive(ctx context.Context) error {
	c.currMessageIndex++
	if c.currMessageIndex >= c.messageBufSize {
		c.currMessageIndex = 0
		deadline, _ := ctx.Deadline()
		if err := c.conn.SetReadDeadline(deadline); err != nil {
			return fmt.Errorf("VLP-16 receiver: %w", err)
		}
		n, err := c.conn.ReadBatch(c.messages, 0)
		if err != nil {
			return fmt.Errorf("VLP-16 receiver: %w", err)
		}
		c.messageBufSize = n
	}
	c.currPacket.UnmarshalRawPacket(c.rawPacketBuf[c.currMessageIndex])
	return nil
}

// SourceIP returns the source IP of the last received VLP-16 packet.
func (c *Receiver) SourceIP() net.IP {
	return c.messages[c.currMessageIndex].Addr.(*net.UDPAddr).IP
}

// RawPacket returns the raw, unprocessed last received VLP-16 packet.
func (c *Receiver) RawPacket() *RawPacket {
	return c.rawPacketBuf[c.currMessageIndex]
}

// Packet returns the last received VLP-16 packet.
func (c *Receiver) Packet() *Packet {
	return &c.currPacket
}

// Close the client's underlying UDP connection.
func (c *Receiver) Close() error {
	if err := c.conn.Close(); err != nil {
		return fmt.Errorf("VLP-16 receiver: close: %w", err)
	}
	return nil
}
