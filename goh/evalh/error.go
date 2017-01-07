package evalh

import (
	"fmt"
	"go/token"
)

type Error struct {
	msg string
	pos token.Position
}

func (err *Error) Error() string {
	return err.pos.String() + ": " + err.msg
}

type intError string

func (err *intError) Error() string {
	return string(err)
}

func newIntError(msg string) *intError {
	return &intError(msg)
}
func newIntErrorf(format string, a ...interface{}) *intError {
	return newIntError(fmt.Sprintf(format, a...))
}
