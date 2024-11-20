package errs

import (
	"errors"
	"fmt"
	"runtime"
)

// ErrCode represents an error code in the system.
type ErrCode struct {
	value int
}

// Value returns the integer value of the error code.
func (ec ErrCode) Value() int {
	return ec.value
}

// String returns the string representation of the error code.
func (ec ErrCode) String() string {
	return codeNames[ec]
}

// UnmarshalText implement the unmarshal interface for JSON conversions.
func (ec *ErrCode) UnmarshalText(data []byte) error {
	errName := string(data)

	v, exists := codeNumbers[errName]
	if !exists {
		return fmt.Errorf("err code %q does not exist", errName)
	}

	*ec = v

	return nil
}

// MarshalText implement the marshal interface for JSON conversions.
func (ec ErrCode) MarshalText() ([]byte, error) {
	return []byte(ec.String()), nil
}

// Equal provides support for the go-cmp package and testing.
func (ec ErrCode) Equal(ec2 ErrCode) bool {
	return ec.value == ec2.value
}

// =============================================================================

// Error represents an error in the system.
type Error struct {
	Code     ErrCode `json:"code"`
	Message  string  `json:"message"`
	FuncName string  `json:"-"`
	FileName string  `json:"-"`
}

// New constructs an error based on an app error.
func New(code ErrCode, err error) *Error {
	pc, filename, line, _ := runtime.Caller(1)

	return &Error{
		Code:     code,
		Message:  err.Error(),
		FuncName: runtime.FuncForPC(pc).Name(),
		FileName: fmt.Sprintf("%s:%d", filename, line),
	}
}

// Newf constructs an error based on a error message.
func Newf(code ErrCode, format string, v ...any) *Error {
	pc, filename, line, _ := runtime.Caller(1)

	return &Error{
		Code:     code,
		Message:  fmt.Sprintf(format, v...),
		FuncName: runtime.FuncForPC(pc).Name(),
		FileName: fmt.Sprintf("%s:%d", filename, line),
	}
}

// NewError checks for an Error in the error interface value. If it doesn't
// exist, will create one from the error.
func NewError(err error) *Error {
	var errsErr *Error
	if errors.As(err, &errsErr) {
		return errsErr
	}

	return New(Internal, err)
}

// Error implements the error interface.
func (e *Error) Error() string {
	return e.Message
}
