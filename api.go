package gonumberio

import (
	"io"

	ite "github.com/Matej-Chmel/go-number-io/internal"
)

const (
	DefaultChunkSize int = 32768
)

type ByteReader = ite.ByteReader

func Read1D[T any](r io.Reader) ([]T, error) {
	return Read1DCustom(r, DefaultChunkSize, GetConversion[T]())
}

func Read1DCustom[T any](
	r io.Reader, chunkSize int, conv func(*ByteReader) (T, uint, error)) ([]T, error) {

	reader, err := ite.RunSliceReader(r, chunkSize, conv, 1)
	return reader.Buf1, err
}

func Read2D[T any](r io.Reader) ([][]T, error) {
	return Read2DCustom(r, DefaultChunkSize, GetConversion[T]())
}

func Read2DCustom[T any](
	r io.Reader, chunkSize int, conv func(*ByteReader) (T, uint, error)) ([][]T, error) {

	reader, err := ite.RunSliceReader(r, chunkSize, conv, 2)
	return reader.Buf2, err
}

func Read3D[T any](r io.Reader) ([][][]T, error) {
	return Read3DCustom(r, DefaultChunkSize, GetConversion[T]())
}

func Read3DCustom[T any](
	r io.Reader, chunkSize int, conv func(*ByteReader) (T, uint, error)) ([][][]T, error) {

	reader, err := ite.RunSliceReader(r, chunkSize, conv, 3)
	return reader.Buf3, err
}
