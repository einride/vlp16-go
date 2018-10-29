package vlp16

import (
	"encoding/binary"
	"io"

	"github.com/pkg/errors"
)

type Packet struct {
	Blocks     [12]Block
	Timestamp  uint32
	ReturnMode ReturnMode
	ProductID  ProductID
}

type Block struct {
	Azimuth  uint16
	Channels [32]Channel
}

type Channel struct {
	Distance     uint16
	Reflectivity uint8
}

type ReturnMode uint8

const (
	ReturnModeStrongest  ReturnMode = 0x37
	ReturnModeLastReturn ReturnMode = 0x38
	ReturnModeDualReturn ReturnMode = 0x39
)

type ProductID uint8

const (
	ProductIDHDL32E ProductID = 0x21
	ProductIDVLP16  ProductID = 0x22
)

func (p *Packet) ReadFrom(r io.Reader) (n int64, err error) {
	// Read 12 blocks
	for i := 0; i < len(p.Blocks); i++ {
		if _, err := p.Blocks[i].ReadFrom(r); err != nil {
			return 0, err
		}
	}
	// Read timestamp
	if err := binary.Read(r, binary.LittleEndian, &p.Timestamp); err != nil {
		return 0, err
	}
	// Read factory flags
	if err := binary.Read(r, binary.LittleEndian, &p.ReturnMode); err != nil {
		return 0, err
	}
	if err := binary.Read(r, binary.LittleEndian, &p.ProductID); err != nil {
		return 0, err
	}
	return 1206, nil
}

func (b *Block) ReadFrom(r io.Reader) (n int64, err error) {
	var flag uint16
	if err := binary.Read(r, binary.LittleEndian, &flag); err != nil {
		return 0, err
	}
	if flag != 0xeeff {
		return 0, errors.Errorf("unexpected flag: %v", flag)
	}
	if err := binary.Read(r, binary.LittleEndian, &b.Azimuth); err != nil {
		return 0, err
	}
	for i := 0; i < len(b.Channels); i++ {
		if _, err := b.Channels[i].ReadFrom(r); err != nil {
			return 0, err
		}
	}
	return 0, nil
}

func (c *Channel) ReadFrom(r io.Reader) (n int64, err error) {
	if err := binary.Read(r, binary.LittleEndian, &c.Distance); err != nil {
		return 0, err
	}
	if err := binary.Read(r, binary.LittleEndian, &c.Reflectivity); err != nil {
		return 0, err
	}
	return 0, nil
}
