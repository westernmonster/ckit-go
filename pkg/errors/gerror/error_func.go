package gerror

import (
	"ckit-go/pkg/errors/ecode"
	"fmt"
	"runtime"
)

// stack represents a stack of program counters.
type stack []uintptr

const (
	// maxStackDepth marks the max stack depth for error back traces.
	maxStackDepth = 64
)

func New(text string) error {
	return &Error{
		stack: callers(),
		text:  text,
		code:  ecode.Nil,
	}
}

func Newf(format string, args ...interface{}) error {
	return &Error{
		stack: callers(),
		text:  fmt.Sprintf(format, args...),
		code:  ecode.Nil,
	}
}

// NewSkip creates and returns an error which is formatted from given text.
// The parameter `skip` specifies the stack callers skipped amount.
func NewSkip(skip int, text string) error {
	return &Error{
		stack: callers(skip),
		text:  text,
		code:  ecode.Nil,
	}
}

// NewSkipf returns an error that formats as the given format and args.
// The parameter `skip` specifies the stack callers skipped amount.
func NewSkipf(skip int, format string, args ...interface{}) error {
	return &Error{
		stack: callers(skip),
		text:  fmt.Sprintf(format, args...),
		code:  ecode.Nil,
	}
}

// Wrap wraps error with text. It returns nil if given err is nil.
// Note that it does not lose the error code of wrapped error, as it inherits the error code from it.
func Wrap(err error, text string) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(),
		text:  text,
		code:  Code(err),
	}
}

// Cause returns the root cause error of `err`.
func Cause(err error) error {
	if err == nil {
		return nil
	}
	if e, ok := err.(ICause); ok {
		return e.Cause()
	}
	if e, ok := err.(IUnwrap); ok {
		return Cause(e.Unwrap())
	}
	return err
}

// Stack returns the stack callers as string.
// It returns the error string directly if the `err` does not support stacks.
func Stack(err error) string {
	if err == nil {
		return ""
	}
	if e, ok := err.(IStack); ok {
		return e.Stack()
	}
	return err.Error()
}

// Current creates and returns the current level error.
// It returns nil if current level error is nil.
func Current(err error) error {
	if err == nil {
		return nil
	}
	if e, ok := err.(ICurrent); ok {
		return e.Current()
	}
	return err
}

// Unwrap returns the next level error.
// It returns nil if current level error or the next level error is nil.
func Unwrap(err error) error {
	if err == nil {
		return nil
	}
	if e, ok := err.(IUnwrap); ok {
		return e.Unwrap()
	}
	return nil
}

// HasStack checks and reports whether `err` implemented interface `gerror.IStack`.
func HasStack(err error) bool {
	_, ok := err.(IStack)
	return ok
}

// Equal reports whether current error `err` equals to error `target`.
// Please note that, in default comparison logic for `Error`,
// the errors are considered the same if both the `code` and `text` of them are the same.
func Equal(err, target error) bool {
	if err == target {
		return true
	}
	if e, ok := err.(IEqual); ok {
		return e.Equal(target)
	}
	if e, ok := target.(IEqual); ok {
		return e.Equal(err)
	}
	return false
}

// Is reports whether current error `err` has error `target` in its chaining errors.
// It is just for implements for stdlib errors.Is from Go version 1.17.
func Is(err, target error) bool {
	if e, ok := err.(IIs); ok {
		return e.Is(target)
	}
	return false
}

// HasError is alias of Is, which more easily understanding semantics.
func HasError(err, target error) bool {
	return Is(err, target)
}

// callers returns the stack callers.
// Note that it here just retrieves the caller memory address array not the caller information.
func callers(skip ...int) stack {
	var (
		pcs [maxStackDepth]uintptr
		n   = 3
	)
	if len(skip) > 0 {
		n += skip[0]
	}
	return pcs[:runtime.Callers(n, pcs[:])]
}
