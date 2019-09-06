package main

import (
	"errors"
	"fmt"
)

type CustomError string

func (e CustomError) Error() string { return string(e) }
func (e CustomError) Unwrap() error { return errors.New("unwrap error") }

func main() {

	// 新建error
	err := errors.New("test error")
	fmt.Println(err)

	ce := CustomError("test error")
	fmt.Println(errors.Is(err, ce))

	var e error
	fmt.Println(errors.As(err, &e), e)

	fmt.Println(errors.Unwrap(ce))
}
