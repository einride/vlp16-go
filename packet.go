package vlp16

import (
	"encoding/binary"

	"golang.org/x/xerrors"
)

const (
	// BlocksPerPacket is the number of blocks in a packet.
	BlocksPerPacket = 12
	// ChannelsPerBlock is the number of channels in a block.
	ChannelsPerBlock = 32
)

// Packet is a VLP-16 packet.
type Packet struct {
	Blocks     [BlocksPerPacket]Block
	Timestamp  uint32
	ReturnMode ReturnMode
	ProductID  ProductID
}

// Block is a VLP-16 block.
type Block struct {
	Flag     uint16
	Azimuth  uint16
	Channels [ChannelsPerBlock]Channel
}

// Channel is a VLP-16 channel.
type Channel struct {
	Distance     uint16
	Reflectivity uint8
}

// structure of packet.
const (
	// lengths
	lengthOfDistance     = 2
	lengthOfReflectivity = 1
	lengthOfAzimuth      = 2
	lengthOfFlag         = 2
	lengthOfProductID    = 1
	lengthOfReturnMode   = 1
	lengthOfTimestamp    = 4
	lengthOfChannel      = lengthOfDistance + lengthOfReflectivity
	lengthOfBlock        = lengthOfFlag + lengthOfAzimuth + lengthOfChannel*ChannelsPerBlock
	lengthOfPacket       = lengthOfBlock*BlocksPerPacket + lengthOfTimestamp + lengthOfReturnMode + lengthOfProductID
	// channel indices
	indexOfDistanceInChannel     = 0
	indexOfReflectivityInChannel = lengthOfDistance
	// block indices
	indexOfFlagInBlock     = 0
	indexOfAzimuthInBlock  = lengthOfFlag
	indexOfChannelsInBlock = indexOfAzimuthInBlock + lengthOfAzimuth
	// packet indices
	indexOfBlocksInPacket     = 0
	indexOfTimestampInPacket  = lengthOfBlock * BlocksPerPacket
	indexOfReturnModeInPacket = indexOfTimestampInPacket + lengthOfTimestamp
	indexOfProductIDInPacket  = indexOfReturnModeInPacket + lengthOfReturnMode
)

// compile-time assertion on structure of packet
var _ [1206]struct{} = [lengthOfPacket]struct{}{}

// flag is the magic value of the flag field.
const flag = 0xeeff

// Validate the packet.
func (p *Packet) Validate() error {
	for i := range p.Blocks {
		if p.Blocks[i].Flag != flag {
			return xerrors.Errorf("validate packet: invalid flag %#x in block %v", p.Blocks[i].Flag, i)
		}
	}
	return nil
}

func (p *Packet) unmarshal(b *[lengthOfPacket]byte) {
	for iBlock := range p.Blocks {
		block := &p.Blocks[iBlock]
		baseBlock := indexOfBlocksInPacket + iBlock*lengthOfBlock
		block.Flag = binary.LittleEndian.Uint16(b[baseBlock+indexOfFlagInBlock:])
		block.Azimuth = binary.LittleEndian.Uint16(b[baseBlock+indexOfAzimuthInBlock:])
		for iChannel := range p.Blocks[iBlock].Channels {
			channel := &block.Channels[iChannel]
			baseChannel := baseBlock + indexOfChannelsInBlock + iChannel*lengthOfChannel
			channel.Distance = binary.LittleEndian.Uint16(b[baseChannel+indexOfDistanceInChannel:])
			channel.Reflectivity = b[baseChannel+indexOfReflectivityInChannel]
		}
	}
	p.Timestamp = binary.LittleEndian.Uint32(b[indexOfTimestampInPacket:])
	p.ReturnMode = ReturnMode(b[indexOfReturnModeInPacket])
	p.ProductID = ProductID(b[indexOfProductIDInPacket])
}
