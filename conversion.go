package gonumberio

import (
	"fmt"
	r "reflect"

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

func ConvertSigned[T constraints.Signed](r *ByteReader) (T, uint, error) {
	return ite.ConvertSignedTemplate[T](r, ite.ProcessIntNonDigit, ite.ProcessDigit)
}

func ConvertUnsigned[T constraints.Unsigned](r *ByteReader) (T, uint, error) {
	return ite.ConvertTemplate[T](r, ite.ProcessUintNonDigit, ite.ProcessDigit)
}

func GetConversion[T any]() func(r *ByteReader) (T, uint, error) {
	var ifc interface{}

	switch kind := ite.GetTypeKind[T](); kind {
	case r.Bool:
		ifc = ConvertBool
	case r.Float32:
		ifc = ConvertFloat[float32]
	case r.Float64:
		ifc = ConvertFloat[float64]
	case r.Int:
		ifc = ConvertSigned[int]
	case r.Int8:
		ifc = ConvertSigned[int8]
	case r.Int16:
		ifc = ConvertSigned[int16]
	case r.Int32:
		ifc = ConvertSigned[int32]
	case r.Int64:
		ifc = ConvertSigned[int64]
	case r.Uint:
		ifc = ConvertUnsigned[uint]
	case r.Uint8:
		ifc = ConvertUnsigned[uint8]
	case r.Uint16:
		ifc = ConvertUnsigned[uint16]
	case r.Uint32:
		ifc = ConvertUnsigned[uint32]
	case r.Uint64:
		ifc = ConvertUnsigned[uint64]
	}

	if fn, ok := ifc.(func(r *ByteReader) (T, uint, error)); ok {
		return fn
	}

	return nil
}
