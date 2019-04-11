package main

import (
	"fmt"
)

func init() {
	fmt.Print("Content-Type: text/plain;charset=utf-8\n\n")
}

func main() {
	fmt.Println("This is gocgi test")
}