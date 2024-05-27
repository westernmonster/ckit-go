package gerror

import (
	"ckit-go/pkg/errors/ecode"
	"errors"
	"fmt"
	"io"
)

type Error struct {
	error error      // Wrapped error
	stack stack      //  records the stack information when this error is created or wrapped.
	text  string     // custom error text when error is created, might be empty when its code is not nil
	code  ecode.Code // error code if necessary
}

func (p *Error) Error() string {
	if p == nil {
		return ""
	}

	errStr := p.text
	if errStr == "" && p.code != nil {
		errStr = p.code.Message()
	}

	if p.error != nil {
		if errStr != "" {
			errStr += ": "
		}

		errStr += p.error.Error()
	}

	return errStr
}

// Cause returns the root cause error
func (p *Error) Cause() error {
	if p == nil {
		return nil
	}
	lookup := p
	for lookup != nil {
		if lookup.error != nil {
			if e, ok := lookup.error.(*Error); ok {
				lookup = e
			} else if e, ok := lookup.error.(ICause); ok {
				return e.Cause()
			} else {
				return lookup.error
			}
		} else {
			return errors.New(lookup.text)
		}
	}

	return nil
}

func (p *Error) Current() error {
	if p == nil {
		return nil
	}

	return &Error{
		error: nil,
		stack: p.stack,
		text:  p.text,
		code:  p.code,
	}
}

func (p *Error) Unwrap() error {
	if p == nil {
		return nil
	}

	return p.error
}

// Equal reports whether current error `err` equals to error `target`.
// Please note that, in default comparison for `Error`,
// the errors are considered the same if both the `code` and `text` of them are the same.
func (err *Error) Equal(target error) bool {
	if err == target {
		return true
	}
	// Code should be the same.
	// Note that if both errors have `nil` code, they are also considered equal.
	if err.code != Code(target) {
		return false
	}
	// Text should be the same.
	if err.text != fmt.Sprintf(`%-s`, target) {
		return false
	}
	return true
}

// Is reports whether current error `err` has error `target` in its chaining errors.
// It is just for implements for stdlib errors.Is from Go version 1.17.
func (err *Error) Is(target error) bool {
	if Equal(err, target) {
		return true
	}
	nextErr := err.Unwrap()
	if nextErr == nil {
		return false
	}
	if Equal(nextErr, target) {
		return true
	}
	if e, ok := nextErr.(IIs); ok {
		return e.Is(target)
	}
	return false
}

// Code returns the error code.
// It returns Nil if it has no error code.
func (err *Error) Code() ecode.Code {
	if err == nil {
		return ecode.Nil
	}
	if err.code == ecode.Nil {
		return Code(err.Unwrap())
	}
	return err.code
}

// SetCode updates the internal code with given code.
func (err *Error) SetCode(code ecode.Code) {
	if err == nil {
		return
	}
	err.code = code
}

// Format formats the frame according to the fmt.Formatter interface.
//
// %v, %s   : Print all the error string;
// %-v, %-s : Print current level error string;
// %+s      : Print full stack error list;
// %+v      : Print the error string and full stack error list
func (err *Error) Format(s fmt.State, verb rune) {
	switch verb {
	case 's', 'v':
		switch {
		case s.Flag('-'):
			if err.text != "" {
				_, _ = io.WriteString(s, err.text)
			} else {
				_, _ = io.WriteString(s, err.Error())
			}
		case s.Flag('+'):
			if verb == 's' {
				_, _ = io.WriteString(s, err.Stack())
			} else {
				_, _ = io.WriteString(s, err.Error()+"\n"+err.Stack())
			}
		default:
			_, _ = io.WriteString(s, err.Error())
		}
	}
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
// Note that do not use pointer as its receiver here.
func (err Error) MarshalJSON() ([]byte, error) {
	return []byte(`"` + err.Error() + `"`), nil
}
