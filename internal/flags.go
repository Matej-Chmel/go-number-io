package internal

const (
	// Floating point symbol found when reading the next digit
	DecimalDot uint = 10
	// Minus sign symbol found when reading the next digit
	MinusSign uint = 11
	// Newline found when reading the next digit
	Newline uint = 12
	// Whitespace found when reading the next digit
	WhiteSpace uint = 13
	// Error found when reading the next digit
	ErrDigit uint = 14
	// Letter found when reading the next digit
	Letter uint = 15

	// The slice reader should break the loop for current element
	Break uint = 0x01
	// Current element has decimal places
	HasDecimals uint = 0x02
	// Newline found when reading current element
	HasNewline uint = 0x04
	// Current element has at least one digit
	HasValue uint = 0x08
	// Current element has minus sign
	IsNegative uint = 0x10
)
