package main

import (
	"fmt"
	"unsafe"
)

func main() {
	v := 123
	//	unsafe.Pointer -> int
	tmp := int(uintptr(unsafe.Pointer(&v)))
	fmt.Printf("%X\n", tmp)
}
