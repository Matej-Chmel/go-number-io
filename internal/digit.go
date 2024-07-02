package internal

import (
	"errors"

	"golang.org/x/exp/constraints"
)

// Generic type constraint for all numbers
type Number interface {
	constraints.Float | constraints.Integer
}

// Generic type constraint for floats and integers
type SignedNumber interface {
	constraints.Float | constraints.Signed
}

// Combines digit with the current result
func ProcessDigit[T constraints.Integer](digit uint, flags uint, res T) (T, uint) {
	return res*T(10) + T(digit), flags | HasValue
}

// Processes non-digit symbols for floats
func ProcessFloatNonDigit[T constraints.Float](digit uint, flags uint, res T) (uint, error) {
	if digit == DecimalDot {
		if (flags & HasDecimals) == HasDecimals {
			return 0, errors.New("Two decimal dots")
		}

		return flags | HasDecimals | HasValue, nil
	}

	return ProcessSignedNonDigit(digit, flags, res)
}

// Processes non-digit symbols for integers
func ProcessIntNonDigit[T constraints.Signed](digit uint, flags uint, res T) (uint, error) {
	if digit == DecimalDot {
		return 0, errors.New("Decimal dot in signed integer")
	}

	return ProcessSignedNonDigit(digit, flags, res)
}

// Processes non-digit symbols for all number types
func ProcessNonDigit[T Number](digit uint, flags uint, res T) (uint, error) {
	if digit == Letter {
		if (flags & HasValue) == 0 {
			return 0, errors.New("Letter in number")
		}

		return flags | Break, nil
	}

	if digit == Newline {
		return flags | HasNewline | Break, nil
	}

	if digit == WhiteSpace {
		if (flags & HasValue) == 0 {
			return flags, nil
		}

		return flags | Break, nil
	}

	if digit == 0 &&
		res == T(0) &&
		(flags&HasValue) == HasValue &&
		(flags&HasDecimals) == 0 {
		return 0, errors.New("Bad leading sequence")
	}

	return flags, nil
}

// Processes non-digit symbols for floats and integers
func ProcessSignedNonDigit[T SignedNumber](digit uint, flags uint, res T) (uint, error) {
	if digit == MinusSign {
		if (flags & IsNegative) == IsNegative {
			return 0, errors.New("Double negative integer")
		}

		if (flags & HasValue) == HasValue {
			return 0, errors.New("Minus sign after digit or dot")
		}

		return flags | HasValue | IsNegative, nil
	}

	return ProcessNonDigit(digit, flags, res)
}

// Processes non-digit symbols for unsigned integers
func ProcessUintNonDigit[T constraints.Unsigned](digit uint, flags uint, res T) (uint, error) {
	if digit == DecimalDot {
		return 0, errors.New("Decimal dot in unsigned integer")
	}

	if digit == MinusSign {
		return 0, errors.New("Negative sign in unsigned integer")
	}

	return ProcessNonDigit(digit, flags, res)
}
