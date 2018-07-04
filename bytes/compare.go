package main

import (
	"bytes"
	"fmt"
)

func main() {
	a := []byte("abc")
	b := []byte("ABC")
	cmp := bytes.Compare(a, b)
	fmt.Println(a, b, cmp)
}
