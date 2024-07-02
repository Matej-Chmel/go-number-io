package internal

import "io"

// Buffered reader of bytes
type ByteReader struct {
	buf    []byte
	bufLen int
	impl   io.Reader
	index  int
}

// Constructs new ByteReader
func NewByteReader(r io.Reader, chunkSize int) *ByteReader {
	return &ByteReader{
		buf:    make([]byte, chunkSize),
		bufLen: 0,
		impl:   r,
		index:  0,
	}
}

// Searches for byte b, if found moves back one character behind b
func (r *ByteReader) LookAheadFor(b byte) (bool, error) {
	for {
		next, err := r.NextByteConvertNewline()

		if err != nil {
			return false, err
		}

		if next == '\n' {
			return true, nil
		}

		if b == next {
			r.MoveBack()
			return false, nil
		}
	}
}

// Moves the head one place backwards
func (r *ByteReader) MoveBack() {
	r.index--
}

// Returns the next byte without any conversions.
// If current buffer is exhausted, a new buffer is read.
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

// Returns next byte and convert "\r\n" to "\n"
func (r *ByteReader) NextByteConvertNewline() (byte, error) {
	b, err := r.NextByte()

	if err != nil {
		return 0, err
	}

	if b == '\r' {
		if nb, err := r.NextByte(); err != nil {
			return 0, err
		} else if nb == '\n' {
			return '\n', nil
		}

		r.MoveBack()
	} else if b == '\n' {
		return '\n', nil
	}

	return b, nil
}

// Returns next number symbol as uint
func (r *ByteReader) NextDigit() (uint, error) {
	b, err := r.NextByteConvertNewline()

	if err != nil {
		return ErrDigit, err
	}

	if b >= '0' && b <= '9' {
		return uint(b - '0'), nil
	}

	if b == '-' {
		return MinusSign, nil
	}

	if b == '.' {
		return DecimalDot, nil
	}

	if b == '\n' {
		return Newline, nil
	}

	if b == ' ' || b == '\t' || b == '\r' {
		return WhiteSpace, nil
	}

	return Letter, nil
}

// Skip the next byte b, if found
func (r *ByteReader) SkipByte(b byte) error {
	for {
		next, err := r.NextByte()

		if err != nil {
			return err
		}

		if b == next {
			return nil
		}
	}
}
