package ecode

type Code interface {
	Code() int

	Message() string

	Details() interface{}
}

func New(code int, message string, details interface{}) Code {
	return localCode{
		code:    code,
		message: message,
		details: details,
	}
}

func WithCode(code Code, details interface{}) Code {
	return localCode{
		code:    code.Code(),
		message: code.Message(),
		details: details,
	}
}
