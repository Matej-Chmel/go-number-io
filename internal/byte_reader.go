package internal

import (
	"io"
)

type ByteReader struct {
	buf    []byte
	bufLen int
	impl   io.Reader
	index  int
}

func NewByteReader(r io.Reader, chunkSize int) *ByteReader {
	return &ByteReader{
		buf:    make([]byte, chunkSize),
		bufLen: 0,
		impl:   r,
		index:  0,
	}
}

func (r *ByteReader) NextByte() (byte, error) {
	if r.index >= r.bufLen {
		var err error
		r.bufLen, err = r.impl.Read(r.buf)

		if err != nil {
			return 0, err
		}

		r.index = 0
	}

	res := r.buf[r.index]
	r.index++
	return res, nil
}

func (r *ByteReader) NextDataByte() (byte, error) {
	b, err := r.NextByte()

	if err != nil {
		return 0, err
	}

	if b == '\r' {
		if nb, err := r.NextByte(); err != nil {
			return 0, err
		} else if nb == '\n' {
			return 0, ErrNewLine
		}

		r.index--
	} else if b == '\n' {
		return 0, ErrNewLine
	}

	return b, nil
}

func (r *ByteReader) NextDigit() (uint, error) {
	b, err := r.NextDataByte()

	if err != nil {
		return 0, err
	}

	if b >= '0' && b <= '9' {
		return uint(b - '0'), nil
	} else if b == '-' {
		return MinusSign, nil
	} else if b == '.' {
		return DecimalDot, nil
	} else if b == ' ' || b == '\t' || b == '\r' {
		return WhiteSpace, nil
	}

	return 0, NewCustomError(CodeLetter, "Letter %c in number", b)
}
