package vlp16

import (
	"encoding/binary"

	"golang.org/x/xerrors"
)

// ScanPackets is a bufio.Scanner SplitFunc that splits a stream into VLP-16 packets.
func ScanPackets(data []byte, atEOF bool) (int, []byte, error) {
	if len(data) < lengthOfPacket {
		if len(data) > 0 && atEOF {
			return 0, nil, xerrors.Errorf("incomplete packet at EOF")
		}
		return 0, nil, nil
	}
	packetFlag := binary.LittleEndian.Uint16(data)
	if packetFlag != flag {
		return 0, nil, xerrors.Errorf("invalid flag at packet start: %#x", packetFlag)
	}
	return lengthOfPacket, data[:lengthOfPacket], nil
}
