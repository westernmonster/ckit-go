package main

import (
	"ckit-go/pkg/errors/gerror"
	"fmt"
	"reflect"
)

func main() {
	err := gerror.New("")

	rvExpect := reflect.ValueOf(nil)

	fmt.Println(err.String(), string(rvExpect))
}
