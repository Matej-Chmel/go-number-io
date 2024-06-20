package internal

import (
	"errors"
	"fmt"
)

const (
	CodeBadFormat int = iota
	CodeLetter
)

var (
	ErrBreak        = errors.New("break")
	ErrNewLine      = errors.New("newline")
	ErrNewLineValue = errors.New("newline with value")
	ErrEOFValue     = errors.New("EOF with value")
)

type CustomError struct {
	Code    int
	Message string
}

func (c *CustomError) Error() string {
	return c.Message
}

func NewCustomError(code int, format string, a ...any) *CustomError {
	return &CustomError{
		Code:    code,
		Message: fmt.Sprintf(format, a...),
	}
}

func isErrorCode(err error, code int) bool {
	if c, ok := err.(*CustomError); ok {
		return c.Code == code
	}

	return false
}
