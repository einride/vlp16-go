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
	Reflectivity Reflectivity
}
type Reflectivity uint8

type ReturnMode uint8

const (
	ReturnModeStrongest  ReturnMode = 0x37
	ReturnModeLastReturn ReturnMode = 0x38
	ReturnModeDualReturn ReturnMode = 0x39
)

type ProductID uint8

const (
	ProductIDHDL32E    ProductID = 0x21
	ProductIDVLP16     ProductID = 0x22
	ProductIDPuckLITE  ProductID = 0x22
	ProductIDPuckHiRes ProductID = 0x24
	ProductIDVLP32C    ProductID = 0x28
	ProductIDVelarray  ProductID = 0x31
	ProductIDVLS128    ProductID = 0x63
)

func (p *Packet) Read(r io.Reader) error {
	// Read 12 blocks
	for i := 0; i < len(p.Blocks); i++ {
		if err := p.Blocks[i].Read(r); err != nil {
			return err
		}
	}
	// Read timestamp
	if err := binary.Read(r, binary.LittleEndian, &p.Timestamp); err != nil {
		return err
	}
	// Read factory flags
	if err := binary.Read(r, binary.LittleEndian, &p.ReturnMode); err != nil {
		return err
	}
	if err := binary.Read(r, binary.LittleEndian, &p.ProductID); err != nil {
		return err
	}
	return nil
}

func (b *Block) Read(r io.Reader) error {
	var flag uint16
	if err := binary.Read(r, binary.LittleEndian, &flag); err != nil {
		return err
	}
	if flag != 0xeeff {
		return errors.Errorf("unexpected flag: %v", flag)
	}
	if err := binary.Read(r, binary.LittleEndian, &b.Azimuth); err != nil {
		return err
	}
	for i := 0; i < len(b.Channels); i++ {
		if err := b.Channels[i].Read(r); err != nil {
			return err
		}
	}
	return nil
}

func (c *Channel) Read(r io.Reader) error {
	if err := binary.Read(r, binary.LittleEndian, &c.Distance); err != nil {
		return err
	}
	if err := binary.Read(r, binary.LittleEndian, &c.Reflectivity); err != nil {
		return err
	}
	return nil
}
