package gonumberio

import (
	"fmt"

	ite "github.com/Matej-Chmel/go-number-io/internal"

	"golang.org/x/exp/constraints"
)

func ConvertBool(r *ByteReader) (bool, uint, error) {
	for {
		b, err := r.NextDataByte()

		if err != nil {
			return false, 0, err
		}

		if b == '\n' {
			return false, ite.HasNewline, nil
		}

		if b == '0' {
			return false, ite.HasValue, nil
		} else if b == '1' {
			return true, ite.HasValue, nil
		} else if b == '\t' || b == ' ' {
			continue
		} else {
			return false, 0, fmt.Errorf("Unknown bool symbol %c", b)
		}
	}
}

func ConvertByte(r *ByteReader) (byte, uint, error) {
	b, err := r.NextByte()
	return b, ite.HasValue, err
}

func convertSignedTemplate[T ite.SignedNumber](r *ByteReader,
	processNonDigit func(uint, uint, T) (uint, error),
	processDigit func(uint, uint, T) (T, uint),
) (T, uint, error) {
	res, flags, err := convertTemplate(r, processNonDigit, processDigit)

	if (flags & ite.IsNegative) == ite.IsNegative {
		res *= T(-1)
	}

	return res, flags, err
}

func convertTemplate[T ite.Number](
	r *ByteReader,
	processNonDigit func(uint, uint, T) (uint, error),
	processDigit func(uint, uint, T) (T, uint),
) (T, uint, error) {
	var digit uint = 0
	var err error = nil
	var flags uint = 0
	res := T(0)

	for {
		digit, err = r.NextDigit()

		if err != nil {
			break
		}

		flags, err = processNonDigit(digit, flags, res)

		if err != nil || (flags&ite.Break) == ite.Break {
			break
		}

		if digit >= ite.DecimalDot {
			continue
		}

		res, flags = processDigit(digit, flags, res)
	}

	return res, flags, err
}

func ConvertFloat[T constraints.Float](r *ByteReader) (T, uint, error) {
	decMult := T(.1)
	processDigit := func(digit uint, flags uint, res T) (T, uint) {
		if (flags & ite.HasDecimals) == ite.HasDecimals {
			res += decMult * T(digit)
			decMult *= T(.1)
		} else {
			res = res*T(10.) + T(digit)
		}

		return res, flags | ite.HasValue
	}

	return convertSignedTemplate(r, ite.ProcessFloatNonDigit, processDigit)
}

func ConvertSigned[T constraints.Signed](r *ByteReader) (T, uint, error) {
	return convertSignedTemplate[T](r, ite.ProcessIntNonDigit, ite.ProcessDigit)
}

func ConvertUnsigned[T constraints.Unsigned](r *ByteReader) (T, uint, error) {
	return convertTemplate[T](r, ite.ProcessUintNonDigit, ite.ProcessDigit)
}
