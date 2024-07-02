package gonumberio

import (
	"errors"
	"io"

	ite "github.com/Matej-Chmel/go-number-io/internal"
)

const (
	// Default buffer size for ByteReader
	DefaultChunkSize int = 32768
)

// Exported ByteReader
type ByteReader = ite.ByteReader

// Read one element of type T from a Reader
func Read0D[T any](r io.Reader) (T, error) {
	return Read0DCustom(r, DefaultChunkSize, GetConversion[T]())
}

// Read one element of type T from a Reader with options
func Read0DCustom[T any](
	r io.Reader, chunkSize int, conv func(*ByteReader) (T, uint, error)) (T, error) {

	arr, err := Read1DCustom[T](r, chunkSize, conv)

	if err != nil {
		var res T
		return res, err
	}

	if len(arr) < 1 {
		var res T
		return res, errors.New("Empty file")
	}

	return arr[0], nil
}

// Read a 1D slice of type T from a Reader
func Read1D[T any](r io.Reader) ([]T, error) {
	return Read1DCustom(r, DefaultChunkSize, GetConversion[T]())
}

// Read a 1D slice of type T from a Reader with options
func Read1DCustom[T any](
	r io.Reader, chunkSize int, conv func(*ByteReader) (T, uint, error)) ([]T, error) {

	reader, err := ite.RunSliceReader(r, chunkSize, conv, 1)
	return reader.Buf1, err
}

// Read a 2D slice of type T from a Reader
func Read2D[T any](r io.Reader) ([][]T, error) {
	return Read2DCustom(r, DefaultChunkSize, GetConversion[T]())
}

// Read a 2D slice of type T from a Reader with options
func Read2DCustom[T any](
	r io.Reader, chunkSize int, conv func(*ByteReader) (T, uint, error)) ([][]T, error) {

	reader, err := ite.RunSliceReader(r, chunkSize, conv, 2)
	return reader.Buf2, err
}

// Read a 3D slice of type T from a Reader
func Read3D[T any](r io.Reader) ([][][]T, error) {
	return Read3DCustom(r, DefaultChunkSize, GetConversion[T]())
}

// Read a 3D slice of type T from a Reader with options
func Read3DCustom[T any](
	r io.Reader, chunkSize int, conv func(*ByteReader) (T, uint, error)) ([][][]T, error) {

	reader, err := ite.RunSliceReader(r, chunkSize, conv, 3)
	return reader.Buf3, err
}
