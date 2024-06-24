package internal

import (
	"errors"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Float | constraints.Integer
}

type SignedNumber interface {
	constraints.Float | constraints.Signed
}

type SliceItem interface {
	bool | Number
}

func ProcessDigit[T constraints.Integer](digit uint, flags uint, res T) (T, uint) {
	return res*T(10) + T(digit), flags | HasValue
}

func ProcessFloatNonDigit[T constraints.Float](digit uint, flags uint, res T) (uint, error) {
	if digit == DecimalDot {
		if (flags & HasDecimals) == HasDecimals {
			return 0, errors.New("Two decimal dots")
		}

		return flags | HasDecimals | HasValue, nil
	}

	return ProcessSignedNonDigit(digit, flags, res)
}

func ProcessIntNonDigit[T constraints.Signed](digit uint, flags uint, res T) (uint, error) {
	if digit == DecimalDot {
		return 0, errors.New("Decimal dot in signed integer")
	}

	return ProcessSignedNonDigit(digit, flags, res)
}

func ProcessNonDigit[T Number](digit uint, flags uint, res T) (uint, error) {
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

func ProcessUintNonDigit[T constraints.Unsigned](digit uint, flags uint, res T) (uint, error) {
	if digit == DecimalDot {
		return 0, errors.New("Decimal dot in unsigned integer")
	}

	if digit == MinusSign {
		return 0, errors.New("Negative sign in unsigned integer")
	}

	return ProcessNonDigit(digit, flags, res)
}
