package internal

import (
	"io"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Float | constraints.Integer
}

func ProcessDigit[T Number](digit uint, res T, hasDigits bool) (T, bool, error) {
	if digit == WhiteSpace {
		if !hasDigits {
			return res, hasDigits, nil
		}

		return res, hasDigits, ErrBreak
	}

	return res*10 + T(digit), true, nil
}

func ProcessError[T Number](err error, res T, hasDigits bool) (T, error) {
	if hasDigits {
		if err == ErrNewLine {
			return res, ErrNewLineValue
		}

		if err == io.EOF {
			return res, ErrEOFValue
		}
	}

	return 0, err
}
