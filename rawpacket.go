package vlp16

import "encoding/hex"

// RawPacket is a raw, binary VLP-16 packet.
type RawPacket [SizeOfPacket]byte

// String implements Stringer for raw packets.
func (r *RawPacket) String() string {
	return hex.EncodeToString(r[:])
}
