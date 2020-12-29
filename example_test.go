package vlp16

func exampleData() *[lengthOfPacket]byte {
	return &[lengthOfPacket]byte{
		0xff, 0xee, 0x66, 0x28, 0x0, 0x0, 0x2, 0xff, 0x2, 0x3, 0x0, 0x0, 0x3, 0x10, 0x3, 0x4b, 0x0, 0x0, 0x2,
		0xac, 0x2, 0x39, 0x0, 0x0, 0x3, 0xbc, 0x2, 0x38, 0x0, 0x0, 0x1, 0xbd, 0x2, 0x45, 0xbe, 0x2, 0x0, 0xc8,
		0x2, 0x57, 0xa9, 0x2, 0x23, 0xc5, 0x2, 0x32, 0xc3, 0x2, 0x1d, 0xd9, 0x2, 0x4a, 0x0, 0x0, 0x2, 0xf, 0x3,
		0x7, 0x0, 0x0, 0x2, 0xe, 0x3, 0x46, 0x0, 0x0, 0x4, 0xac, 0x2, 0x3d, 0x0, 0x0, 0x3, 0xc4, 0x2, 0x41,
		0x0, 0x0, 0x2, 0xba, 0x2, 0x38, 0x0, 0x0, 0x3, 0xc6, 0x2, 0x3e, 0xb5, 0x2, 0x26, 0xd6, 0x2, 0x39, 0xc6,
		0x2, 0x18, 0xd3, 0x2, 0x32, 0xff, 0xee, 0x8e, 0x28, 0x0, 0x0, 0x2, 0x0, 0x0, 0x64, 0x0, 0x0, 0x3, 0x18,
		0x3, 0x4c, 0x0, 0x0, 0x2, 0xad, 0x2, 0x3e, 0x0, 0x0, 0x3, 0xc4, 0x2, 0x4f, 0x0, 0x0, 0x2, 0xc1, 0x2,
		0x41, 0xdc, 0x2, 0x0, 0xcc, 0x2, 0x57, 0xb8, 0x2, 0x25, 0xd7, 0x2, 0x46, 0xf9, 0x2, 0x21, 0xd6, 0x2, 0x44,
		0x0, 0x0, 0x2, 0x1a, 0x3, 0x3c, 0x0, 0x0, 0x2, 0x1c, 0x3, 0x4c, 0x0, 0x0, 0x2, 0xac, 0x2, 0x3d, 0x0,
		0x0, 0x3, 0xca, 0x2, 0x5b, 0x0, 0x0, 0x3, 0xc7, 0x2, 0x54, 0xdc, 0x2, 0x0, 0xc9, 0x2, 0x63, 0xb7, 0x2,
		0x23, 0xe1, 0x2, 0x55, 0x3, 0x3, 0x25, 0xd6, 0x2, 0x51, 0xff, 0xee, 0xb6, 0x28, 0x0, 0x0, 0x2, 0x1c, 0x3,
		0x45, 0x0, 0x0, 0x2, 0x1a, 0x3, 0x54, 0x0, 0x0, 0x2, 0xae, 0x2, 0x39, 0x0, 0x0, 0x3, 0xc2, 0x2, 0x4f,
		0x0, 0x0, 0x3, 0xc7, 0x2, 0x53, 0xc8, 0x2, 0x0, 0xc9, 0x2, 0x57, 0xb5, 0x2, 0x24, 0xd3, 0x2, 0x48, 0x8,
		0x3, 0x33, 0xdd, 0x2, 0x4c, 0x0, 0x0, 0x2, 0x1f, 0x3, 0x41, 0x0, 0x0, 0x3, 0x1a, 0x3, 0x4b, 0x0, 0x0,
		0x2, 0xb4, 0x2, 0x34, 0x0, 0x0, 0x3, 0xc4, 0x2, 0x4d, 0x0, 0x0, 0x2, 0xc6, 0x2, 0x49, 0xd0, 0x2, 0x0,
		0xcb, 0x2, 0x59, 0xbc, 0x2, 0x27, 0xd8, 0x2, 0x43, 0x12, 0x3, 0x3b, 0xdf, 0x2, 0x44, 0xff, 0xee, 0xde, 0x28,
		0x0, 0x0, 0x2, 0x21, 0x3, 0x41, 0x0, 0x0, 0x2, 0x1c, 0x3, 0x51, 0x0, 0x0, 0x2, 0xb6, 0x2, 0x3d, 0x0,
		0x0, 0x3, 0xc8, 0x2, 0x3f, 0x0, 0x0, 0x2, 0xc3, 0x2, 0x41, 0xce, 0x2, 0x0, 0xcf, 0x2, 0x44, 0xc1, 0x2, 0x2d,
		0xd0, 0x2, 0x3a, 0x19, 0x3, 0x2e, 0xdd, 0x2, 0x38, 0x0, 0x0, 0x2, 0x27, 0x3, 0x3d, 0x0, 0x0, 0x2, 0x20,
		0x3, 0x51, 0x0, 0x0, 0x4, 0xb4, 0x2, 0x39, 0x0, 0x0, 0x3, 0xc9, 0x2, 0x3c, 0x0, 0x0, 0x1, 0xc1, 0x2,
		0x3f, 0xce, 0x2, 0x3, 0xcd, 0x2, 0x47, 0xc6, 0x2, 0x33, 0xd2, 0x2, 0x3f, 0x21, 0x3, 0x30, 0xd5, 0x2, 0x38,
		0xff, 0xee, 0x6, 0x29, 0x0, 0x0, 0x2, 0x22, 0x3, 0x43, 0x0, 0x0, 0x2, 0x28, 0x3, 0x48, 0x0, 0x0, 0x2,
		0xb5, 0x2, 0x3e, 0x0, 0x0, 0x3, 0xc2, 0x2, 0x4b, 0x0, 0x0, 0x1, 0xc2, 0x2, 0x49, 0x0, 0x0, 0x64, 0xc5,
		0x2, 0x49, 0xc4, 0x2, 0x40, 0xdb, 0x2, 0x46, 0x1a, 0x3, 0x3e, 0xd3, 0x2, 0x3f, 0x0, 0x0, 0x2, 0x25, 0x3,
		0x41, 0x0, 0x0, 0x2, 0x28, 0x3, 0x48, 0x0, 0x0, 0x2, 0xba, 0x2, 0x3d, 0x0, 0x0, 0x3, 0xbe, 0x2, 0x4f,
		0x0, 0x0, 0x1, 0xc4, 0x2, 0x4e, 0xd1, 0x2, 0x29, 0xc7, 0x2, 0x4b, 0xc4, 0x2, 0x40, 0xd2, 0x2, 0x3e, 0x20,
		0x3, 0x3e, 0xd7, 0x2, 0x3b, 0xff, 0xee, 0x2e, 0x29, 0x0, 0x0, 0x2, 0x2b, 0x3, 0x41, 0x0, 0x0, 0x2, 0x24,
		0x3, 0x4b, 0x0, 0x0, 0x2, 0xbf, 0x2, 0x40, 0x0, 0x0, 0x3, 0xc0, 0x2, 0x3f, 0x0, 0x0, 0x1, 0xc3, 0x2,
		0x41, 0xd1, 0x2, 0x2a, 0xca, 0x2, 0x3c, 0xc4, 0x2, 0x43, 0xcd, 0x2, 0x35, 0x18, 0x3, 0x3b, 0xdd, 0x2, 0x35,
		0x0, 0x0, 0x2, 0x2f, 0x3, 0x3d, 0x0, 0x0, 0x2, 0x26, 0x3, 0x4c, 0x0, 0x0, 0x4, 0xbe, 0x2, 0x3b, 0x0,
		0x0, 0x3, 0xc6, 0x2, 0x35, 0x0, 0x0, 0x1, 0xc5, 0x2, 0x3c, 0xd1, 0x2, 0x2d, 0xcc, 0x2, 0x3c, 0xca, 0x2,
		0x43, 0xcf, 0x2, 0x31, 0x10, 0x3, 0x12, 0xda, 0x2, 0x30, 0xff, 0xee, 0x56, 0x29, 0x0, 0x0, 0x2, 0x31, 0x3,
		0x3d, 0x0, 0x0, 0x2, 0x30, 0x3, 0x4c, 0x0, 0x0, 0x2, 0xbd, 0x2, 0x40, 0x0, 0x0, 0x3, 0xca, 0x2, 0x33,
		0x0, 0x0, 0x1, 0xc3, 0x2, 0x3a, 0xd1, 0x2, 0x3f, 0xce, 0x2, 0x3c, 0xd1, 0x2, 0x4c, 0xde, 0x2, 0x35, 0xe5,
		0x2, 0x6, 0xdb, 0x2, 0x31, 0x0, 0x0, 0x2, 0x2d, 0x3, 0x3f, 0x0, 0x0, 0x2, 0x34, 0x3, 0x4f, 0x0, 0x0,
		0x2, 0xc1, 0x2, 0x40, 0x0, 0x0, 0x3, 0xce, 0x2, 0x3b, 0x0, 0x0, 0x1, 0xc9, 0x2, 0x3c, 0xd5, 0x2, 0x3c,
		0xce, 0x2, 0x39, 0xd0, 0x2, 0x4e, 0xe0, 0x2, 0x38, 0xe9, 0x2, 0x8, 0xdd, 0x2, 0x35, 0xff, 0xee, 0x7d, 0x29,
		0x0, 0x0, 0x2, 0x33, 0x3, 0x3d, 0x0, 0x0, 0x2, 0x34, 0x3, 0x51, 0x0, 0x0, 0x2, 0xc5, 0x2, 0x40, 0x0,
		0x0, 0x3, 0xcc, 0x2, 0x45, 0x0, 0x0, 0x1, 0xd3, 0x2, 0x3d, 0xdb, 0x2, 0x28, 0xd1, 0x2, 0x45, 0xcf, 0x2,
		0x53, 0xe0, 0x2, 0x43, 0xe5, 0x2, 0x4, 0xe3, 0x2, 0x3d, 0x0, 0x0, 0x2, 0x39, 0x3, 0x3c, 0x0, 0x0, 0x2,
		0x2e, 0x3, 0x4c, 0x0, 0x0, 0x2, 0xc9, 0x2, 0x45, 0x0, 0x0, 0x3, 0xcc, 0x2, 0x57, 0x0, 0x0, 0x1, 0xd2,
		0x2, 0x4e, 0xde, 0x2, 0x13, 0xd8, 0x2, 0x54, 0xd3, 0x2, 0x53, 0xe4, 0x2, 0x55, 0xfc, 0x2, 0x5, 0xe8, 0x2,
		0x4b, 0xff, 0xee, 0xa5, 0x29, 0x0, 0x0, 0x2, 0x3a, 0x3, 0x38, 0x0, 0x0, 0x2, 0x36, 0x3, 0x51, 0x0, 0x0,
		0x2, 0xc9, 0x2, 0x45, 0x0, 0x0, 0x3, 0xce, 0x2, 0x51, 0x0, 0x0, 0x1, 0xcd, 0x2, 0x58, 0xdd, 0x2, 0x23,
		0xd9, 0x2, 0x55, 0xd8, 0x2, 0x4e, 0xdd, 0x2, 0x52, 0x0, 0x0, 0x64, 0xea, 0x2, 0x52, 0x0, 0x0, 0x2, 0x3a,
		0x3, 0x3a, 0x0, 0x0, 0x2, 0x3c, 0x3, 0x4f, 0x0, 0x0, 0x2, 0xc7, 0x2, 0x4b, 0x0, 0x0, 0x3, 0xd4, 0x2,
		0x4b, 0x0, 0x0, 0x1, 0xcf, 0x2, 0x51, 0xaf, 0x2, 0x19, 0xda, 0x2, 0x57, 0xda, 0x2, 0x4b, 0xe2, 0x2, 0x46,
		0x9, 0x3, 0x9, 0xe2, 0x2, 0x4b, 0xff, 0xee, 0xcc, 0x29, 0x0, 0x0, 0x2, 0x39, 0x3, 0x3c, 0x0, 0x0, 0x2,
		0x44, 0x3, 0x4b, 0x0, 0x0, 0x2, 0xc9, 0x2, 0x47, 0x0, 0x0, 0x3, 0xd2, 0x2, 0x3c, 0x0, 0x0, 0x1, 0xd3,
		0x2, 0x3d, 0xd8, 0x2, 0x50, 0xd2, 0x2, 0x3b, 0xda, 0x2, 0x4b, 0xe5, 0x2, 0x3f, 0x11, 0x3, 0x8, 0xde, 0x2,
		0x3b, 0x0, 0x0, 0x2, 0x3e, 0x3, 0x3a, 0x0, 0x0, 0x2, 0x44, 0x3, 0x4c, 0x0, 0x0, 0x2, 0xca, 0x2, 0x38,
		0x0, 0x0, 0x3, 0xcd, 0x2, 0x3c, 0x0, 0x0, 0x1, 0xd5, 0x2, 0x3d, 0xe1, 0x2, 0x3f, 0xd7, 0x2, 0x48, 0xd6,
		0x2, 0x49, 0xe3, 0x2, 0x36, 0xf, 0x3, 0x7, 0xe5, 0x2, 0x3e, 0xff, 0xee, 0xf5, 0x29, 0x0, 0x0, 0x2, 0x44,
		0x3, 0x3a, 0x0, 0x0, 0x3, 0x40, 0x3, 0x4c, 0x0, 0x0, 0x2, 0xce, 0x2, 0x30, 0x0, 0x0, 0x3, 0xce, 0x2,
		0x32, 0x0, 0x0, 0x1, 0xd8, 0x2, 0x35, 0xe9, 0x2, 0x3c, 0xd4, 0x2, 0x39, 0xd9, 0x2, 0x29, 0xe0, 0x2, 0x35,
		0x7, 0x3, 0x4, 0xe7, 0x2, 0x36, 0x0, 0x0, 0x2, 0x45, 0x3, 0x3d, 0x0, 0x0, 0x2, 0x46, 0x3, 0x48, 0x0,
		0x0, 0x2, 0xcc, 0x2, 0x2d, 0x0, 0x0, 0x3, 0xd6, 0x2, 0x35, 0x0, 0x0, 0x4d, 0xd4, 0x2, 0x36, 0xea, 0x2,
		0x3b, 0xdf, 0x2, 0x3f, 0xd5, 0x2, 0x15, 0xe2, 0x2, 0x32, 0xff, 0x2, 0x5, 0xeb, 0x2, 0x38, 0xff, 0xee, 0x1d,
		0x2a, 0x0, 0x0, 0x2, 0x46, 0x3, 0x38, 0x0, 0x0, 0x2, 0x4a, 0x3, 0x4b, 0x0, 0x0, 0x2, 0xcd, 0x2, 0x33,
		0x0, 0x0, 0x3, 0xda, 0x2, 0x36, 0xf4, 0x2, 0x28, 0xd5, 0x2, 0x3c, 0xe4, 0x2, 0x3b, 0xde, 0x2, 0x3c, 0xdd,
		0x2, 0x41, 0xea, 0x2, 0x35, 0x1, 0x3, 0x5, 0xe9, 0x2, 0x35, 0x0, 0x0, 0x2, 0x44, 0x3, 0x3a, 0x0, 0x0,
		0x2, 0x4e, 0x3, 0x48, 0x0, 0x0, 0x2, 0xce, 0x2, 0x3d, 0x0, 0x0, 0x3, 0xda, 0x2, 0x39, 0xc4, 0x2, 0x29,
		0xdc, 0x2, 0x36, 0xd5, 0x2, 0x28, 0xdd, 0x2, 0x47, 0xbf, 0x2, 0x2e, 0xef, 0x2, 0x38, 0x4, 0x3, 0x5, 0xeb,
		0x2, 0x3a, 0x9b, 0x6b, 0xea, 0x9e, 0x37, 0x22,
	}
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
