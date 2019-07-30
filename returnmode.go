package vlp16

// ReturnMode represents a VLP-16 return mode.
type ReturnMode uint8

const (
	ReturnModeStrongest  ReturnMode = 0x37
	ReturnModeLastReturn ReturnMode = 0x38
	ReturnModeDualReturn ReturnMode = 0x39
)
