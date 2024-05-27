package ecode

import "fmt"

type localCode struct {
	code    int
	message string
	details interface{}
}

func (p localCode) Code() int {
	return p.code
}

func (p localCode) Message() string {
	return p.message
}

func (p localCode) Details() interface{} {
	return p.details
}

func (p localCode) String() string {
	if p.details == nil {
		return fmt.Sprintf("%d:%s %v", p.code, p.message, p.details)
	}

	if p.message != "" {
		return fmt.Sprintf("%d:%s", p.code, p.message)
	}

	return fmt.Sprintf("%d", p.code)
}
