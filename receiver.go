package vlp16

import (
	"context"
	"net"

	"golang.org/x/net/ipv4"
	"golang.org/x/xerrors"
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
		return nil, xerrors.Errorf("VLP-16: listen UDP: %w", err)
	}
	udpConn := packetConn.(*net.UDPConn)
	if err := udpConn.SetReadBuffer(opts.bufferSizeBytes); err != nil {
		return nil, xerrors.Errorf("VLP-16: listen UDP: %w", err)
	}
	conn := ipv4.NewPacketConn(udpConn)
	c := &Receiver{conn: conn}
	// allocate memory for batch reads
	c.messages = make([]ipv4.Message, 0, opts.batchSize)
	for i := 0; i < opts.batchSize; i++ {
		c.packetBuf = append(c.packetBuf, &[lengthOfPacket]byte{})
		c.messages = append(c.messages, ipv4.Message{
			Buffers: [][]byte{c.packetBuf[i][:]},
		})
	}
	return c, nil
}

// Receiver receives VLP-16 packets.
type Receiver struct {
	conn             *ipv4.PacketConn
	messages         []ipv4.Message
	packetBuf        []*[lengthOfPacket]byte
	messageBufSize   int
	currMessageIndex int
	currPacket       Packet
	currPointCloud   PointCloud
}

// Receive a VLP-16 packet.
func (c *Receiver) Receive(ctx context.Context) error {
	c.currMessageIndex++
	if c.currMessageIndex >= c.messageBufSize {
		c.currMessageIndex = 0
		deadline, _ := ctx.Deadline()
		if err := c.conn.SetReadDeadline(deadline); err != nil {
			return xerrors.Errorf("VLP-16 receiver: %w", err)
		}
		n, err := c.conn.ReadBatch(c.messages, 0)
		if err != nil {
			return xerrors.Errorf("VLP-16 receiver: %w", err)
		}
		c.messageBufSize = n
	}
	c.currPacket.unmarshal(c.packetBuf[c.currMessageIndex])
	c.currPointCloud.UnmarshalPacket(&c.currPacket)
	return nil
}

// SourceIP returns the source IP of the last received VLP-16 packet.
func (c *Receiver) SourceIP() net.IP {
	return c.messages[c.currMessageIndex].Addr.(*net.UDPAddr).IP
}

// RawPacket returns the raw bytes of the last received VLP-16 packet.
func (c *Receiver) RawPacket() []byte {
	return c.packetBuf[c.currMessageIndex][:c.messages[c.currMessageIndex].N]
}

// Packet returns the last received VLP-16 packet.
func (c *Receiver) Packet() *Packet {
	return &c.currPacket
}

// PointCloud returns the point cloud representation of the last received packet.
func (c *Receiver) PointCloud() *PointCloud {
	return &c.currPointCloud
}

// Close the client's underlying UDP connection.
func (c *Receiver) Close() error {
	if err := c.conn.Close(); err != nil {
		return xerrors.Errorf("VLP-16 receiver: close: %w", err)
	}
	return nil
}
