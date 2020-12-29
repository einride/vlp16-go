package vlp16

// ReturnMode represents a VLP-16 return mode.
type ReturnMode uint8

//go:generate stringer -type ReturnMode -trimprefix ReturnMode

const (
	ReturnModeStrongest  ReturnMode = 0x37
	ReturnModeLastReturn ReturnMode = 0x38
	ReturnModeDualReturn ReturnMode = 0x39
)
