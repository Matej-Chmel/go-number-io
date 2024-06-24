package internal

const (
	DecimalDot uint = 10
	MinusSign  uint = 11
	Newline    uint = 12
	WhiteSpace uint = 13
	ErrDigit   uint = 14

	Break       uint = 0x01
	HasDecimals uint = 0x02
	HasNewline  uint = 0x04
	HasValue    uint = 0x08
	IsNegative  uint = 0x10
)
