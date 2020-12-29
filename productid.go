package vlp16

// ProductID represents a Velodyne product ID.
type ProductID uint8

//go:generate stringer -type ProductID -trimprefix ProductID

const (
	ProductIDHDL32E    ProductID = 0x21
	ProductIDVLP16     ProductID = 0x22
	ProductIDPuckLITE  ProductID = 0x22
	ProductIDPuckHiRes ProductID = 0x24
	ProductIDVLP32C    ProductID = 0x28
	ProductIDVelarray  ProductID = 0x31
	ProductIDVLS128    ProductID = 0x63
)
