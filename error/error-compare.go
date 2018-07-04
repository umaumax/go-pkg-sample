package main

import (
	"errors"
	"fmt"
)

var (
	MyError = errors.New("MyError")
)

func main() {
	e0 := MyError
	e1 := errors.New("MyError")

	if e0 == MyError {
		fmt.Println("e0 == MyError")
	} else {
		fmt.Println("e0 != MyError")
	}
	if e0 == errors.New("MyError") {
		fmt.Println(`e0 == errors.New("MyError")`)
	} else {
		fmt.Println(`e0 != errors.New("MyError")`)
	}
	if e1 == MyError {
		fmt.Println("e1 == MyError")
	} else {
		fmt.Println("e1 != MyError")
	}
	if e1 == errors.New("MyError") {
		fmt.Println(`e1 == errors.New("MyError")`)
	} else {
		fmt.Println(`e1 != errors.New("MyError")`)
	}
}
