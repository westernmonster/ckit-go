package gerror

import "ckit-go/pkg/errors/ecode"

type IIs interface {
	Error() string
	Is(target error) bool
}

type IEqual interface {
	Error() string
	Equal(target error) bool
}

type ICode interface {
	Error() string
	Code() ecode.Code
}

type IStack interface {
	Error() string
	Stack() string
}

type ICause interface {
	Error() string
	Cause() error
}

type ICurrent interface {
	Error() string
	Current() error
}

type IUnwrap interface {
	Error() string
	Unwrap() error
}

const (
	// commaSeparatorSpace is the comma separator with space.
	commaSeparatorSpace = ", "
)
