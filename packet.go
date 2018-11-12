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
	Flag     uint16
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
	if err := binary.Read(r, binary.LittleEndian, p); err != nil {
		return err
	}
	// validate flags
	for i := 0; i < len(p.Blocks); i++ {
		if p.Blocks[i].Flag != 0xeeff {
			return errors.Errorf("invalid flag value %v in block %v", p.Blocks[i].Flag, i)
		}
	}
	return nil
}
