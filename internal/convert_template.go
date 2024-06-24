package internal

func ConvertSignedTemplate[T SignedNumber](r *ByteReader,
	processNonDigit func(uint, uint, T) (uint, error),
	processDigit func(uint, uint, T) (T, uint),
) (T, uint, error) {
	res, flags, err := ConvertTemplate(r, processNonDigit, processDigit)

	if (flags & IsNegative) == IsNegative {
		res *= T(-1)
	}

	return res, flags, err
}

func ConvertTemplate[T Number](
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

		if err != nil || (flags&Break) == Break {
			break
		}

		if digit >= DecimalDot {
			continue
		}

		res, flags = processDigit(digit, flags, res)
	}

	return res, flags, err
}
