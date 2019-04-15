package main

import (
	"errors"
	"fmt"
)

func main() {

	// 新建error
	err := errors.New("test error")
	fmt.Println(err)
}