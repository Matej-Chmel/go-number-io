package gonumberio

import (
	"errors"
	"fmt"
	"io"
	r "reflect"

	ite "github.com/Matej-Chmel/go-number-io/internal"
)

// Internal conversion of any to T
func convertAny[T any](a any, info *ite.DescendInfo) (T, error) {
	if info.Dimensions == 0 {
		return convertAny0D[T](a)
	} else if res, ok := a.(T); ok {
		return res, nil
	}

	var res T
	return res, fmt.Errorf("Unable to convert %T to %T", a, res)
}

// Internal conversion of any to T where T is not a slice
func convertAny0D[T any](a any) (T, error) {
	if slice, ok := a.([]T); ok {
		if len(slice) < 1 {
			var res T
			return res, errors.New("Empty file")
		} else {
			return slice[0], nil
		}
	}

	var res T
	return res, fmt.Errorf("Unable to convert %T to %T", a, res)
}

// Error for a situation in which a default conversion was not found
func dynamicError(aType r.Type) (any, error) {
	return nil, fmt.Errorf("Type %v doesn't have default conversion", aType)
}

// Reads 1D slice of type aType
func dynamicRead1D(reader io.Reader, chunkSize int, aType r.Type) (any, error) {
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

// Reads 2D slice of type aType
func dynamicRead2D(reader io.Reader, chunkSize int, aType r.Type) (any, error) {
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

// Reads 3D slice of type aType
func dynamicRead3D(reader io.Reader, chunkSize int, aType r.Type) (any, error) {
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

// Reads T from reader.
// T can be a single element or 1D, 2D or 3D slice.
func Read[T any](reader io.Reader) (T, error) {
	return ReadCustom[T](reader, DefaultChunkSize)
}

// Reads T from reader with options.
// T can be a single element or 1D, 2D or 3D slice.
func ReadCustom[T any](
	reader io.Reader, chunkSize int) (T, error) {

	info := ite.Descend[T]()
	a, err := readAny[T](chunkSize, &info, reader)

	if err != nil {
		var res T
		return res, err
	}

	return convertAny[T](a, &info)
}

// Internal implementation of ReadCustom[T]
func readAny[T any](
	chunkSize int, info *ite.DescendInfo, reader io.Reader) (any, error) {

	if !info.Supported {
		var res T
		return nil, fmt.Errorf("Type %T is not supported", res)
	}

	switch info.Dimensions {
	case 0, 1:
		return dynamicRead1D(reader, chunkSize, info.ElementType)
	case 2:
		return dynamicRead2D(reader, chunkSize, info.ElementType)
	case 3:
		return dynamicRead3D(reader, chunkSize, info.ElementType)
	}

	var res T
	return nil, fmt.Errorf(
		"Type %T has %d dimensions, only 0-3 dimensions are supported",
		res, info.Dimensions)
}
