package ecode

type Code interface {
	Code() int

	Message() string

	Details() interface{}
}
