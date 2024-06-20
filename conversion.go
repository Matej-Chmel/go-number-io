package gonumberio

import (
	ite "github.com/Matej-Chmel/go-number-io/internal"

	"golang.org/x/exp/constraints"
)

func ConvertBool(r *ByteReader) (bool, error) {
	for {
		b, err := r.NextDataByte()

		if err != nil {
			return false, err
		}

		if b == '0' {
			return false, nil
		} else if b == '1' {
			return true, nil
		}
	}
}

func ConvertByte(r *ByteReader) (byte, error) {
	return r.NextByte()
}

func ConvertFloat[T constraints.Float](r *ByteReader) (T, error) {
	add, mult, res := T(1.0), T(10.0), T(0.0)
	hasDigits, isNegative := false, false

	for {
		digit, err := r.NextDigit()

		if err != nil {
			return ite.ProcessError(err, res, hasDigits)
		}

		if digit == ite.MinusSign {
			if isNegative {
				return 0, ite.NewCustomError(
					ite.CodeBadFormat, "Double negative float")
			}

			isNegative = true
			continue
		}

		if digit == ite.DecimalDot {
			if mult < 0 {
				return 0.0, ite.NewCustomError(
					ite.CodeBadFormat, "Double decimal dot")
			}

			add, hasDigits, mult = 0.1, true, 1/mult
			continue
		}

		if digit == ite.WhiteSpace {
			if !hasDigits {
				continue
			}

			break
		}

		if digit == 0 && res == 0.0 && hasDigits {
			return 0.0, ite.NewCustomError(
				ite.CodeBadFormat, "Double leading zero")
		}

		res += add * T(digit)
		add *= mult
		hasDigits = true
	}

	if isNegative {
		res *= -1.0
	}

	return res, nil
}

func ConvertSigned[T constraints.Signed](r *ByteReader) (T, error) {
	hasDigits, isNegative := false, false
	var res T = 0

	for {
		digit, err := r.NextDigit()

		if err != nil {
			return ite.ProcessError(err, res, hasDigits)
		}

		if digit == ite.MinusSign {
			if isNegative {
				return 0, ite.NewCustomError(
					ite.CodeBadFormat, "Double negative signed")
			}

			isNegative = true
			continue
		}

		if digit == ite.DecimalDot {
			return 0, ite.NewCustomError(
				ite.CodeBadFormat, "Decimal dot in signed integer")
		}

		res, hasDigits, err = ite.ProcessDigit(digit, res, hasDigits)

		if err == ite.ErrBreak {
			break
		}
	}

	if isNegative {
		res *= -1
	}

	return res, nil
}

func ConvertUnsigned[T constraints.Unsigned](r *ByteReader) (T, error) {
	hasDigits := false
	var res T = 0

	for {
		digit, err := r.NextDigit()

		if err != nil {
			return ite.ProcessError(err, res, hasDigits)
		}

		if digit == ite.MinusSign {
			return 0, ite.NewCustomError(
				ite.CodeBadFormat, "Negative sign in unsigned integer")
		}

		if digit == ite.DecimalDot {
			return 0, ite.NewCustomError(
				ite.CodeBadFormat, "Decimal dot in unsigned integer")
		}

		res, hasDigits, err = ite.ProcessDigit(digit, res, hasDigits)

		if err == ite.ErrBreak {
			break
		}
	}

	return res, nil
}
