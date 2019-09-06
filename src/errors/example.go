package main

import (
	"errors"
	"fmt"
)

type CustomError string

func (e CustomError) Error() string { return string(e) }
func (e CustomError) Unwrap() error { return errors.New("unwrap error") }
func (e CustomError) As(err error) bool { return true }

func main() {

	// 新建error
	err := errors.New("test error")
	fmt.Println(err)

	ce := CustomError("test error1")
	fmt.Println(errors.Is(err, ce))

	var e error
	fmt.Println(errors.As(ce, &e), e)

	fmt.Println(errors.Unwrap(ce))
}
