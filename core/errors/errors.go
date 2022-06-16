package errors

import (
	"fmt"

	"golang.org/x/xerrors"
)

type coreError struct {
	err   error
	frame xerrors.Frame
}

// NewCoreError is contructor to set skip param.
func NewCoreError(err error, skip int) error {
	return &coreError{
		err:   err,
		frame: xerrors.Caller(skip),
	}
}

func (e *coreError) Error() string {
	return e.err.Error()
}

func (e *coreError) Format(f fmt.State, c rune) {
	xerrors.FormatError(e, f, c)
}

func (e *coreError) FormatError(p xerrors.Printer) error {
	e.frame.Format(p)

	return e.Unwrap()
}

// Unwrap returns the result of calling the Unwrap method on err.
func (e *coreError) Unwrap() error {
	return e.err
}

// Wrap is constructor to wrap error object.
func Wrap(err error) error {
	const skip = 1

	return &coreError{
		err:   err,
		frame: xerrors.Caller(skip),
	}
}
