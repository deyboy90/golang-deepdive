package errorsdemo

/**

// The error built-in interface type is the conventional interface for
// representing an error condition, with the nil value representing no error.
// Any concrete value that implements this interface can be used as an error value.
type error interface {
	Error() string
}

// New returns an error that formats as the given text.
// Each call to New returns a distinct error value even if the text is identical.
func New(text string) error {
	return &errorString{text}
}

// errorString is a trivial implementation of error.
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

**/

import (
	"errors"
	"fmt"
)

var (
	// error variables start with Err and are exported
	ErrBadRequest = errors.New("Bad request")
	ErrPageMoved  = errors.New("Page moved")
)

func webcall(b bool) error {
	if !b {
		return ErrBadRequest
	}
	return ErrPageMoved

}

func BasicErrorsDemo() {
	if err := webcall(false); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Success")
}
