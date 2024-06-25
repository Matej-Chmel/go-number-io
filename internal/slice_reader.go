package internal

import (
	"errors"
	"io"
)

type SliceReader[T any] struct {
	Buf1        []T
	Buf2        [][]T
	Buf3        [][][]T
	byteReader  *ByteReader
	conv        func(*ByteReader) (T, uint, error)
	dim         uint
	prevNewline bool
}

func NewSliceReader[T any](
	byteReader *ByteReader, conv func(*ByteReader) (T, uint, error), dim uint,
) *SliceReader[T] {
	res := &SliceReader[T]{
		Buf1:        make([]T, 0),
		Buf2:        nil,
		Buf3:        nil,
		byteReader:  byteReader,
		conv:        conv,
		dim:         dim,
		prevNewline: false,
	}

	if dim >= 2 {
		res.Buf2 = make([][]T, 0)
	}

	if dim >= 3 {
		res.Buf3 = make([][][]T, 0)
	}

	return res
}

func (s *SliceReader[T]) add1Dto2D() {
	if len(s.Buf1) > 0 {
		s.Buf2 = append(s.Buf2, s.Buf1)
		s.Buf1 = make([]T, 0)
	}
}

func (s *SliceReader[T]) add2Dto3D() {
	if len(s.Buf2) > 0 {
		s.Buf3 = append(s.Buf3, s.Buf2)
		s.Buf2 = make([][]T, 0)
	}
}

func (s *SliceReader[T]) processNewline() {
	if s.dim == 2 {
		s.add1Dto2D()
		return
	}

	if s.dim != 3 {
		return
	}

	if s.prevNewline {
		s.add2Dto3D()
	} else {
		s.add1Dto2D()
	}

	s.prevNewline = !s.prevNewline
}

func (s *SliceReader[T]) Run() error {
	if s.conv == nil {
		return errors.New("Conversion function is nil")
	}

	for {
		val, flags, err := s.conv(s.byteReader)

		if (flags & HasValue) == HasValue {
			s.Buf1 = append(s.Buf1, val)
			s.prevNewline = false
		}

		if err != nil {
			if err == io.EOF {
				break
			}

			s.Buf1, s.Buf2, s.Buf3 = nil, nil, nil
			return err
		}

		if (flags & HasNewline) == HasNewline {
			s.processNewline()
		}
	}

	if s.dim >= 2 {
		s.add1Dto2D()
	}

	if s.dim == 3 {
		s.add2Dto3D()
	}

	return nil
}

func RunSliceReader[T any](
	r io.Reader, chunkSize int,
	conv func(*ByteReader) (T, uint, error), dim uint) (*SliceReader[T], error) {

	byteReader := NewByteReader(r, chunkSize)
	sliceReader := NewSliceReader(byteReader, conv, dim)
	err := sliceReader.Run()
	return sliceReader, err
}
