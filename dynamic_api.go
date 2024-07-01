package gonumberio

import (
	"fmt"
	"io"
	r "reflect"

	ite "github.com/Matej-Chmel/go-number-io/internal"
)

func dynamicError(aType r.Type) (interface{}, error) {
	return nil, fmt.Errorf("Type %v doesn't have default conversion", aType)
}

func dynamicRead1D(reader io.Reader, chunkSize int, aType r.Type) (interface{}, error) {
	switch kind := aType.Kind(); kind {
	case r.Bool:
		return Read1DCustom(reader, chunkSize, ConvertBool)
	case r.Float32:
		return Read1DCustom(reader, chunkSize, ConvertFloat[float32])
	case r.Float64:
		return Read1DCustom(reader, chunkSize, ConvertFloat[float64])
	case r.Int:
		return Read1DCustom(reader, chunkSize, ConvertSigned[int])
	case r.Int8:
		return Read1DCustom(reader, chunkSize, ConvertSigned[int8])
	case r.Int16:
		return Read1DCustom(reader, chunkSize, ConvertSigned[int16])
	case r.Int32:
		return Read1DCustom(reader, chunkSize, ConvertSigned[int32])
	case r.Int64:
		return Read1DCustom(reader, chunkSize, ConvertSigned[int64])
	case r.Uint:
		return Read1DCustom(reader, chunkSize, ConvertUnsigned[uint])
	case r.Uint8:
		return Read1DCustom(reader, chunkSize, ConvertUnsigned[uint8])
	case r.Uint16:
		return Read1DCustom(reader, chunkSize, ConvertUnsigned[uint16])
	case r.Uint32:
		return Read1DCustom(reader, chunkSize, ConvertUnsigned[uint32])
	case r.Uint64:
		return Read1DCustom(reader, chunkSize, ConvertUnsigned[uint64])
	}

	return dynamicError(aType)
}

func dynamicRead2D(reader io.Reader, chunkSize int, aType r.Type) (interface{}, error) {
	switch kind := aType.Kind(); kind {
	case r.Bool:
		return Read2DCustom(reader, chunkSize, ConvertBool)
	case r.Float32:
		return Read2DCustom(reader, chunkSize, ConvertFloat[float32])
	case r.Float64:
		return Read2DCustom(reader, chunkSize, ConvertFloat[float64])
	case r.Int:
		return Read2DCustom(reader, chunkSize, ConvertSigned[int])
	case r.Int8:
		return Read2DCustom(reader, chunkSize, ConvertSigned[int8])
	case r.Int16:
		return Read2DCustom(reader, chunkSize, ConvertSigned[int16])
	case r.Int32:
		return Read2DCustom(reader, chunkSize, ConvertSigned[int32])
	case r.Int64:
		return Read2DCustom(reader, chunkSize, ConvertSigned[int64])
	case r.Uint:
		return Read2DCustom(reader, chunkSize, ConvertUnsigned[uint])
	case r.Uint8:
		return Read2DCustom(reader, chunkSize, ConvertUnsigned[uint8])
	case r.Uint16:
		return Read2DCustom(reader, chunkSize, ConvertUnsigned[uint16])
	case r.Uint32:
		return Read2DCustom(reader, chunkSize, ConvertUnsigned[uint32])
	case r.Uint64:
		return Read2DCustom(reader, chunkSize, ConvertUnsigned[uint64])
	}

	return dynamicError(aType)
}

func dynamicRead3D(reader io.Reader, chunkSize int, aType r.Type) (interface{}, error) {
	switch kind := aType.Kind(); kind {
	case r.Bool:
		return Read3DCustom(reader, chunkSize, ConvertBool)
	case r.Float32:
		return Read3DCustom(reader, chunkSize, ConvertFloat[float32])
	case r.Float64:
		return Read3DCustom(reader, chunkSize, ConvertFloat[float64])
	case r.Int:
		return Read3DCustom(reader, chunkSize, ConvertSigned[int])
	case r.Int8:
		return Read3DCustom(reader, chunkSize, ConvertSigned[int8])
	case r.Int16:
		return Read3DCustom(reader, chunkSize, ConvertSigned[int16])
	case r.Int32:
		return Read3DCustom(reader, chunkSize, ConvertSigned[int32])
	case r.Int64:
		return Read3DCustom(reader, chunkSize, ConvertSigned[int64])
	case r.Uint:
		return Read3DCustom(reader, chunkSize, ConvertUnsigned[uint])
	case r.Uint8:
		return Read3DCustom(reader, chunkSize, ConvertUnsigned[uint8])
	case r.Uint16:
		return Read3DCustom(reader, chunkSize, ConvertUnsigned[uint16])
	case r.Uint32:
		return Read3DCustom(reader, chunkSize, ConvertUnsigned[uint32])
	case r.Uint64:
		return Read3DCustom(reader, chunkSize, ConvertUnsigned[uint64])
	}

	return dynamicError(aType)
}

func Read[T any](reader io.Reader) (T, error) {
	return ReadCustom[T](reader, DefaultChunkSize)
}

func ReadCustom[T any](
	reader io.Reader, chunkSize int) (T, error) {

	var err error = nil
	var ifc interface{}

	switch info := ite.Descend[T](); info.Dimensions {
	case 0:
		var res T
		return res, fmt.Errorf("Type %T isn't 1D, 2D or 3D slice", res)

	case 1:
		ifc, err = dynamicRead1D(reader, chunkSize, info.ElementType)
	case 2:
		ifc, err = dynamicRead2D(reader, chunkSize, info.ElementType)
	case 3:
		ifc, err = dynamicRead3D(reader, chunkSize, info.ElementType)
	}

	if err != nil {
		var res T
		return res, err
	}

	if res, ok := ifc.(T); ok {
		return res, nil
	}

	var res T
	return res, fmt.Errorf("Unable to convert %T to %T", ifc, res)
}
