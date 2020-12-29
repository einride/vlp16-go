package vlp16

import (
	"io/ioutil"

	"gotest.tools/v3/assert"
)

func exampleRawPacket(t assert.TestingT) *RawPacket {
	data, err := ioutil.ReadFile("testdata/packet.bin")
	assert.NilError(t, err)
	var rawPacket RawPacket
	copy(rawPacket[:], data)
	return &rawPacket
}

func examplePacket() *Packet {
	return &Packet{
		Blocks: [12]Block{
			{
				Azimuth: 0x2866,
				Flag:    0xeeff,
				Channels: [32]Channel{
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x2ff, Reflectivity: 0x3},
					{Distance: 0x0, Reflectivity: 0x3},
					{Distance: 0x310, Reflectivity: 0x4b},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x2ac, Reflectivity: 0x39},
					{Distance: 0x0, Reflectivity: 0x3},
					{Distance: 0x2bc, Reflectivity: 0x38},
					{Distance: 0x0, Reflectivity: 0x1},
					{Distance: 0x2bd, Reflectivity: 0x45},
					{Distance: 0x2be, Reflectivity: 0x0},
					{Distance: 0x2c8, Reflectivity: 0x57},
					{Distance: 0x2a9, Reflectivity: 0x23},
					{Distance: 0x2c5, Reflectivity: 0x32},
					{Distance: 0x2c3, Reflectivity: 0x1d},
					{Distance: 0x2d9, Reflectivity: 0x4a},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x30f, Reflectivity: 0x7},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x30e, Reflectivity: 0x46},
					{Distance: 0x0, Reflectivity: 0x4},
					{Distance: 0x2ac, Reflectivity: 0x3d},
					{Distance: 0x0, Reflectivity: 0x3},
					{Distance: 0x2c4, Reflectivity: 0x41},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x2ba, Reflectivity: 0x38},
					{Distance: 0x0, Reflectivity: 0x3},
					{Distance: 0x2c6, Reflectivity: 0x3e},
					{Distance: 0x2b5, Reflectivity: 0x26},
					{Distance: 0x2d6, Reflectivity: 0x39},
					{Distance: 0x2c6, Reflectivity: 0x18},
					{Distance: 0x2d3, Reflectivity: 0x32},
				},
			},
			{
				Azimuth: 0x288e,
				Flag:    0xeeff,
				Channels: [32]Channel{
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x0, Reflectivity: 0x64},
					{Distance: 0x0, Reflectivity: 0x3},
					{Distance: 0x318, Reflectivity: 0x4c},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x2ad, Reflectivity: 0x3e},
					{Distance: 0x0, Reflectivity: 0x3},
					{Distance: 0x2c4, Reflectivity: 0x4f},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x2c1, Reflectivity: 0x41},
					{Distance: 0x2dc, Reflectivity: 0x0},
					{Distance: 0x2cc, Reflectivity: 0x57},
					{Distance: 0x2b8, Reflectivity: 0x25},
					{Distance: 0x2d7, Reflectivity: 0x46},
					{Distance: 0x2f9, Reflectivity: 0x21},
					{Distance: 0x2d6, Reflectivity: 0x44},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x31a, Reflectivity: 0x3c},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x31c, Reflectivity: 0x4c},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x2ac, Reflectivity: 0x3d},
					{Distance: 0x0, Reflectivity: 0x3},
					{Distance: 0x2ca, Reflectivity: 0x5b},
					{Distance: 0x0, Reflectivity: 0x3},
					{Distance: 0x2c7, Reflectivity: 0x54},
					{Distance: 0x2dc, Reflectivity: 0x0},
					{Distance: 0x2c9, Reflectivity: 0x63},
					{Distance: 0x2b7, Reflectivity: 0x23},
					{Distance: 0x2e1, Reflectivity: 0x55},
					{Distance: 0x303, Reflectivity: 0x25},
					{Distance: 0x2d6, Reflectivity: 0x51},
				},
			},
			{
				Azimuth: 0x28b6,
				Flag:    0xeeff,
				Channels: [32]Channel{
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x31c, Reflectivity: 0x45},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x31a, Reflectivity: 0x54},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x2ae, Reflectivity: 0x39},
					{Distance: 0x0, Reflectivity: 0x3},
					{Distance: 0x2c2, Reflectivity: 0x4f},
					{Distance: 0x0, Reflectivity: 0x3},
					{Distance: 0x2c7, Reflectivity: 0x53},
					{Distance: 0x2c8, Reflectivity: 0x0},
					{Distance: 0x2c9, Reflectivity: 0x57},
					{Distance: 0x2b5, Reflectivity: 0x24},
					{Distance: 0x2d3, Reflectivity: 0x48},
					{Distance: 0x308, Reflectivity: 0x33},
					{Distance: 0x2dd, Reflectivity: 0x4c},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x31f, Reflectivity: 0x41},
					{Distance: 0x0, Reflectivity: 0x3},
					{Distance: 0x31a, Reflectivity: 0x4b},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x2b4, Reflectivity: 0x34},
					{Distance: 0x0, Reflectivity: 0x3},
					{Distance: 0x2c4, Reflectivity: 0x4d},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x2c6, Reflectivity: 0x49},
					{Distance: 0x2d0, Reflectivity: 0x0},
					{Distance: 0x2cb, Reflectivity: 0x59},
					{Distance: 0x2bc, Reflectivity: 0x27},
					{Distance: 0x2d8, Reflectivity: 0x43},
					{Distance: 0x312, Reflectivity: 0x3b},
					{Distance: 0x2df, Reflectivity: 0x44},
				},
			},
			{
				Azimuth: 0x28de,
				Flag:    0xeeff,
				Channels: [32]Channel{
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x321, Reflectivity: 0x41},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x31c, Reflectivity: 0x51},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x2b6, Reflectivity: 0x3d},
					{Distance: 0x0, Reflectivity: 0x3},
					{Distance: 0x2c8, Reflectivity: 0x3f},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x2c3, Reflectivity: 0x41},
					{Distance: 0x2ce, Reflectivity: 0x0},
					{Distance: 0x2cf, Reflectivity: 0x44},
					{Distance: 0x2c1, Reflectivity: 0x2d},
					{Distance: 0x2d0, Reflectivity: 0x3a},
					{Distance: 0x319, Reflectivity: 0x2e},
					{Distance: 0x2dd, Reflectivity: 0x38},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x327, Reflectivity: 0x3d},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x320, Reflectivity: 0x51},
					{Distance: 0x0, Reflectivity: 0x4},
					{Distance: 0x2b4, Reflectivity: 0x39},
					{Distance: 0x0, Reflectivity: 0x3},
					{Distance: 0x2c9, Reflectivity: 0x3c},
					{Distance: 0x0, Reflectivity: 0x1},
					{Distance: 0x2c1, Reflectivity: 0x3f},
					{Distance: 0x2ce, Reflectivity: 0x3},
					{Distance: 0x2cd, Reflectivity: 0x47},
					{Distance: 0x2c6, Reflectivity: 0x33},
					{Distance: 0x2d2, Reflectivity: 0x3f},
					{Distance: 0x321, Reflectivity: 0x30},
					{Distance: 0x2d5, Reflectivity: 0x38},
				},
			},
			{
				Azimuth: 0x2906,
				Flag:    0xeeff,
				Channels: [32]Channel{
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x322, Reflectivity: 0x43},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x328, Reflectivity: 0x48},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x2b5, Reflectivity: 0x3e},
					{Distance: 0x0, Reflectivity: 0x3},
					{Distance: 0x2c2, Reflectivity: 0x4b},
					{Distance: 0x0, Reflectivity: 0x1},
					{Distance: 0x2c2, Reflectivity: 0x49},
					{Distance: 0x0, Reflectivity: 0x64},
					{Distance: 0x2c5, Reflectivity: 0x49},
					{Distance: 0x2c4, Reflectivity: 0x40},
					{Distance: 0x2db, Reflectivity: 0x46},
					{Distance: 0x31a, Reflectivity: 0x3e},
					{Distance: 0x2d3, Reflectivity: 0x3f},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x325, Reflectivity: 0x41},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x328, Reflectivity: 0x48},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x2ba, Reflectivity: 0x3d},
					{Distance: 0x0, Reflectivity: 0x3},
					{Distance: 0x2be, Reflectivity: 0x4f},
					{Distance: 0x0, Reflectivity: 0x1},
					{Distance: 0x2c4, Reflectivity: 0x4e},
					{Distance: 0x2d1, Reflectivity: 0x29},
					{Distance: 0x2c7, Reflectivity: 0x4b},
					{Distance: 0x2c4, Reflectivity: 0x40},
					{Distance: 0x2d2, Reflectivity: 0x3e},
					{Distance: 0x320, Reflectivity: 0x3e},
					{Distance: 0x2d7, Reflectivity: 0x3b},
				},
			},
			{
				Azimuth: 0x292e,
				Flag:    0xeeff,
				Channels: [32]Channel{
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x32b, Reflectivity: 0x41},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x324, Reflectivity: 0x4b},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x2bf, Reflectivity: 0x40},
					{Distance: 0x0, Reflectivity: 0x3},
					{Distance: 0x2c0, Reflectivity: 0x3f},
					{Distance: 0x0, Reflectivity: 0x1},
					{Distance: 0x2c3, Reflectivity: 0x41},
					{Distance: 0x2d1, Reflectivity: 0x2a},
					{Distance: 0x2ca, Reflectivity: 0x3c},
					{Distance: 0x2c4, Reflectivity: 0x43},
					{Distance: 0x2cd, Reflectivity: 0x35},
					{Distance: 0x318, Reflectivity: 0x3b},
					{Distance: 0x2dd, Reflectivity: 0x35},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x32f, Reflectivity: 0x3d},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x326, Reflectivity: 0x4c},
					{Distance: 0x0, Reflectivity: 0x4},
					{Distance: 0x2be, Reflectivity: 0x3b},
					{Distance: 0x0, Reflectivity: 0x3},
					{Distance: 0x2c6, Reflectivity: 0x35},
					{Distance: 0x0, Reflectivity: 0x1},
					{Distance: 0x2c5, Reflectivity: 0x3c},
					{Distance: 0x2d1, Reflectivity: 0x2d},
					{Distance: 0x2cc, Reflectivity: 0x3c},
					{Distance: 0x2ca, Reflectivity: 0x43},
					{Distance: 0x2cf, Reflectivity: 0x31},
					{Distance: 0x310, Reflectivity: 0x12},
					{Distance: 0x2da, Reflectivity: 0x30},
				},
			},
			{
				Azimuth: 0x2956,
				Flag:    0xeeff,
				Channels: [32]Channel{
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x331, Reflectivity: 0x3d},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x330, Reflectivity: 0x4c},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x2bd, Reflectivity: 0x40},
					{Distance: 0x0, Reflectivity: 0x3},
					{Distance: 0x2ca, Reflectivity: 0x33},
					{Distance: 0x0, Reflectivity: 0x1},
					{Distance: 0x2c3, Reflectivity: 0x3a},
					{Distance: 0x2d1, Reflectivity: 0x3f},
					{Distance: 0x2ce, Reflectivity: 0x3c},
					{Distance: 0x2d1, Reflectivity: 0x4c},
					{Distance: 0x2de, Reflectivity: 0x35},
					{Distance: 0x2e5, Reflectivity: 0x6},
					{Distance: 0x2db, Reflectivity: 0x31},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x32d, Reflectivity: 0x3f},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x334, Reflectivity: 0x4f},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x2c1, Reflectivity: 0x40},
					{Distance: 0x0, Reflectivity: 0x3},
					{Distance: 0x2ce, Reflectivity: 0x3b},
					{Distance: 0x0, Reflectivity: 0x1},
					{Distance: 0x2c9, Reflectivity: 0x3c},
					{Distance: 0x2d5, Reflectivity: 0x3c},
					{Distance: 0x2ce, Reflectivity: 0x39},
					{Distance: 0x2d0, Reflectivity: 0x4e},
					{Distance: 0x2e0, Reflectivity: 0x38},
					{Distance: 0x2e9, Reflectivity: 0x8},
					{Distance: 0x2dd, Reflectivity: 0x35},
				},
			},
			{
				Azimuth: 0x297d,
				Flag:    0xeeff,
				Channels: [32]Channel{
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x333, Reflectivity: 0x3d},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x334, Reflectivity: 0x51},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x2c5, Reflectivity: 0x40},
					{Distance: 0x0, Reflectivity: 0x3},
					{Distance: 0x2cc, Reflectivity: 0x45},
					{Distance: 0x0, Reflectivity: 0x1},
					{Distance: 0x2d3, Reflectivity: 0x3d},
					{Distance: 0x2db, Reflectivity: 0x28},
					{Distance: 0x2d1, Reflectivity: 0x45},
					{Distance: 0x2cf, Reflectivity: 0x53},
					{Distance: 0x2e0, Reflectivity: 0x43},
					{Distance: 0x2e5, Reflectivity: 0x4},
					{Distance: 0x2e3, Reflectivity: 0x3d},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x339, Reflectivity: 0x3c},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x32e, Reflectivity: 0x4c},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x2c9, Reflectivity: 0x45},
					{Distance: 0x0, Reflectivity: 0x3},
					{Distance: 0x2cc, Reflectivity: 0x57},
					{Distance: 0x0, Reflectivity: 0x1},
					{Distance: 0x2d2, Reflectivity: 0x4e},
					{Distance: 0x2de, Reflectivity: 0x13},
					{Distance: 0x2d8, Reflectivity: 0x54},
					{Distance: 0x2d3, Reflectivity: 0x53},
					{Distance: 0x2e4, Reflectivity: 0x55},
					{Distance: 0x2fc, Reflectivity: 0x5},
					{Distance: 0x2e8, Reflectivity: 0x4b},
				},
			},
			{
				Azimuth: 0x29a5,
				Flag:    0xeeff,
				Channels: [32]Channel{
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x33a, Reflectivity: 0x38},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x336, Reflectivity: 0x51},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x2c9, Reflectivity: 0x45},
					{Distance: 0x0, Reflectivity: 0x3},
					{Distance: 0x2ce, Reflectivity: 0x51},
					{Distance: 0x0, Reflectivity: 0x1},
					{Distance: 0x2cd, Reflectivity: 0x58},
					{Distance: 0x2dd, Reflectivity: 0x23},
					{Distance: 0x2d9, Reflectivity: 0x55},
					{Distance: 0x2d8, Reflectivity: 0x4e},
					{Distance: 0x2dd, Reflectivity: 0x52},
					{Distance: 0x0, Reflectivity: 0x64},
					{Distance: 0x2ea, Reflectivity: 0x52},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x33a, Reflectivity: 0x3a},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x33c, Reflectivity: 0x4f},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x2c7, Reflectivity: 0x4b},
					{Distance: 0x0, Reflectivity: 0x3},
					{Distance: 0x2d4, Reflectivity: 0x4b},
					{Distance: 0x0, Reflectivity: 0x1},
					{Distance: 0x2cf, Reflectivity: 0x51},
					{Distance: 0x2af, Reflectivity: 0x19},
					{Distance: 0x2da, Reflectivity: 0x57},
					{Distance: 0x2da, Reflectivity: 0x4b},
					{Distance: 0x2e2, Reflectivity: 0x46},
					{Distance: 0x309, Reflectivity: 0x9},
					{Distance: 0x2e2, Reflectivity: 0x4b},
				},
			},
			{
				Azimuth: 0x29cc,
				Flag:    0xeeff,
				Channels: [32]Channel{
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x339, Reflectivity: 0x3c},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x344, Reflectivity: 0x4b},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x2c9, Reflectivity: 0x47},
					{Distance: 0x0, Reflectivity: 0x3},
					{Distance: 0x2d2, Reflectivity: 0x3c},
					{Distance: 0x0, Reflectivity: 0x1},
					{Distance: 0x2d3, Reflectivity: 0x3d},
					{Distance: 0x2d8, Reflectivity: 0x50},
					{Distance: 0x2d2, Reflectivity: 0x3b},
					{Distance: 0x2da, Reflectivity: 0x4b},
					{Distance: 0x2e5, Reflectivity: 0x3f},
					{Distance: 0x311, Reflectivity: 0x8},
					{Distance: 0x2de, Reflectivity: 0x3b},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x33e, Reflectivity: 0x3a},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x344, Reflectivity: 0x4c},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x2ca, Reflectivity: 0x38},
					{Distance: 0x0, Reflectivity: 0x3},
					{Distance: 0x2cd, Reflectivity: 0x3c},
					{Distance: 0x0, Reflectivity: 0x1},
					{Distance: 0x2d5, Reflectivity: 0x3d},
					{Distance: 0x2e1, Reflectivity: 0x3f},
					{Distance: 0x2d7, Reflectivity: 0x48},
					{Distance: 0x2d6, Reflectivity: 0x49},
					{Distance: 0x2e3, Reflectivity: 0x36},
					{Distance: 0x30f, Reflectivity: 0x7},
					{Distance: 0x2e5, Reflectivity: 0x3e},
				},
			},
			{
				Azimuth: 0x29f5,
				Flag:    0xeeff,
				Channels: [32]Channel{
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x344, Reflectivity: 0x3a},
					{Distance: 0x0, Reflectivity: 0x3},
					{Distance: 0x340, Reflectivity: 0x4c},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x2ce, Reflectivity: 0x30},
					{Distance: 0x0, Reflectivity: 0x3},
					{Distance: 0x2ce, Reflectivity: 0x32},
					{Distance: 0x0, Reflectivity: 0x1},
					{Distance: 0x2d8, Reflectivity: 0x35},
					{Distance: 0x2e9, Reflectivity: 0x3c},
					{Distance: 0x2d4, Reflectivity: 0x39},
					{Distance: 0x2d9, Reflectivity: 0x29},
					{Distance: 0x2e0, Reflectivity: 0x35},
					{Distance: 0x307, Reflectivity: 0x4},
					{Distance: 0x2e7, Reflectivity: 0x36},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x345, Reflectivity: 0x3d},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x346, Reflectivity: 0x48},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x2cc, Reflectivity: 0x2d},
					{Distance: 0x0, Reflectivity: 0x3},
					{Distance: 0x2d6, Reflectivity: 0x35},
					{Distance: 0x0, Reflectivity: 0x4d},
					{Distance: 0x2d4, Reflectivity: 0x36},
					{Distance: 0x2ea, Reflectivity: 0x3b},
					{Distance: 0x2df, Reflectivity: 0x3f},
					{Distance: 0x2d5, Reflectivity: 0x15},
					{Distance: 0x2e2, Reflectivity: 0x32},
					{Distance: 0x2ff, Reflectivity: 0x5},
					{Distance: 0x2eb, Reflectivity: 0x38},
				},
			},
			{
				Azimuth: 0x2a1d,
				Flag:    0xeeff,
				Channels: [32]Channel{
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x346, Reflectivity: 0x38},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x34a, Reflectivity: 0x4b},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x2cd, Reflectivity: 0x33},
					{Distance: 0x0, Reflectivity: 0x3},
					{Distance: 0x2da, Reflectivity: 0x36},
					{Distance: 0x2f4, Reflectivity: 0x28},
					{Distance: 0x2d5, Reflectivity: 0x3c},
					{Distance: 0x2e4, Reflectivity: 0x3b},
					{Distance: 0x2de, Reflectivity: 0x3c},
					{Distance: 0x2dd, Reflectivity: 0x41},
					{Distance: 0x2ea, Reflectivity: 0x35},
					{Distance: 0x301, Reflectivity: 0x5},
					{Distance: 0x2e9, Reflectivity: 0x35},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x344, Reflectivity: 0x3a},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x34e, Reflectivity: 0x48},
					{Distance: 0x0, Reflectivity: 0x2},
					{Distance: 0x2ce, Reflectivity: 0x3d},
					{Distance: 0x0, Reflectivity: 0x3},
					{Distance: 0x2da, Reflectivity: 0x39},
					{Distance: 0x2c4, Reflectivity: 0x29},
					{Distance: 0x2dc, Reflectivity: 0x36},
					{Distance: 0x2d5, Reflectivity: 0x28},
					{Distance: 0x2dd, Reflectivity: 0x47},
					{Distance: 0x2bf, Reflectivity: 0x2e},
					{Distance: 0x2ef, Reflectivity: 0x38},
					{Distance: 0x304, Reflectivity: 0x5},
					{Distance: 0x2eb, Reflectivity: 0x3a},
				},
			},
		},
		Timestamp:  0x9eea6b9b,
		ReturnMode: 0x37,
		ProductID:  0x22,
	}
}
