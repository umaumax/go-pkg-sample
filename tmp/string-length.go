package main

import (
	"fmt"
)

func main() {
	str := "An 🍎!"
	//	len(string) is len([]byte)
	fmt.Println(str, "length is", len(str))
}
