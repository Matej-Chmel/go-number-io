package internal

import "io"

type SliceReader[T any] struct {
	Buf1        []T
	Buf2        [][]T
	Buf3        [][][]T
	byteReader  *ByteReader
	conv        func(*ByteReader) (T, error)
	dim         uint
	prevNewLine bool
}

func NewSliceReader[T any](
	byteReader *ByteReader, conv func(*ByteReader) (T, error), dim uint,
) *SliceReader[T] {
	res := &SliceReader[T]{
		Buf1:        make([]T, 0),
		Buf2:        nil,
		Buf3:        nil,
		byteReader:  byteReader,
		conv:        conv,
		dim:         dim,
		prevNewLine: false,
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

func (s *SliceReader[T]) processNewLine() {
	if s.dim <= 1 {
		return
	} else if s.dim == 2 {
		s.add1Dto2D()
	} else if s.prevNewLine {
		s.add2Dto3D()
		s.prevNewLine = false
	} else {
		s.prevNewLine = true
	}
}

func (s *SliceReader[T]) Run() error {
	for {
		val, err := s.conv(s.byteReader)

		if err != nil {
			if err == ErrNewLine {
				s.processNewLine()
				continue
			}

			if err == ErrNewLineValue {
				s.Buf1 = append(s.Buf1, val)
				s.processNewLine()
				continue
			}

			if err == io.EOF {
				break
			}

			if err == ErrEOFValue {
				s.Buf1 = append(s.Buf1, val)
				break
			}

			s.Buf1 = nil
			s.Buf2 = nil
			s.Buf3 = nil
			return err
		}

		s.Buf1 = append(s.Buf1, val)
	}

	return nil
}

func RunSliceReader[T any](
	r io.Reader, chunkSize int,
	conv func(*ByteReader) (T, error), dim uint) (*SliceReader[T], error) {

	byteReader := NewByteReader(r, chunkSize)
	sliceReader := NewSliceReader(byteReader, conv, dim)
	err := sliceReader.Run()
	return sliceReader, err
}
