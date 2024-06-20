package gonumberio

import (
	"io"

	ite "github.com/Matej-Chmel/go-number-io/internal"
)

const (
	DefaultChunkSize int = 32768
)

type ByteReader = ite.ByteReader

func Read1D[T any](
	r io.Reader, chunkSize int, conv func(*ByteReader) (T, error)) ([]T, error) {

	reader, err := ite.RunSliceReader(r, chunkSize, conv, 1)
	return reader.Buf1, err
}

func Read2D[T any](
	r io.Reader, chunkSize int, conv func(*ByteReader) (T, error)) ([][]T, error) {

	reader, err := ite.RunSliceReader(r, chunkSize, conv, 1)
	return reader.Buf2, err
}

func Read3D[T any](
	r io.Reader, chunkSize int, conv func(*ByteReader) (T, error)) ([][][]T, error) {

	reader, err := ite.RunSliceReader(r, chunkSize, conv, 1)
	return reader.Buf3, err
}
