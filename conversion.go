package gonumberio

import (
	"fmt"
	r "reflect"

	ite "github.com/Matej-Chmel/go-number-io/internal"

	"golang.org/x/exp/constraints"
)

// Conversion function for type bool
func ConvertBool(r *ByteReader) (bool, uint, error) {
	for {
		b, err := r.NextByteConvertNewline()

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

// Conversion function for type byte
func ConvertByte(r *ByteReader) (byte, uint, error) {
	b, err := r.NextByte()
	return b, ite.HasValue, err
}

// Conversion function for type float
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

	return ite.ConvertSignedTemplate(r, ite.ProcessFloatNonDigit, processDigit)
}

// Conversion function for floats and integers
func ConvertSigned[T constraints.Signed](r *ByteReader) (T, uint, error) {
	return ite.ConvertSignedTemplate[T](r, ite.ProcessIntNonDigit, ite.ProcessDigit)
}

// Conversion function for unsigned integers
func ConvertUnsigned[T constraints.Unsigned](r *ByteReader) (T, uint, error) {
	return ite.ConvertTemplate[T](r, ite.ProcessUintNonDigit, ite.ProcessDigit)
}

// Returns conversion function for generic type T
func GetConversion[T any]() func(r *ByteReader) (T, uint, error) {
	a := getConversionImpl(ite.GetType[T]())

	if fn, ok := a.(func(r *ByteReader) (T, uint, error)); ok {
		return fn
	}

	return nil
}

// Internal implementation of GetConversion[T]
func getConversionImpl(t r.Type) any {
	switch kind := t.Kind(); kind {
	case r.Bool:
		return ConvertBool
	case r.Float32:
		return ConvertFloat[float32]
	case r.Float64:
		return ConvertFloat[float64]
	case r.Int:
		return ConvertSigned[int]
	case r.Int8:
		return ConvertSigned[int8]
	case r.Int16:
		return ConvertSigned[int16]
	case r.Int32:
		return ConvertSigned[int32]
	case r.Int64:
		return ConvertSigned[int64]
	case r.Uint:
		return ConvertUnsigned[uint]
	case r.Uint8:
		return ConvertUnsigned[uint8]
	case r.Uint16:
		return ConvertUnsigned[uint16]
	case r.Uint32:
		return ConvertUnsigned[uint32]
	case r.Uint64:
		return ConvertUnsigned[uint64]
	}

	return nil
}
