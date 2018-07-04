package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("$1=%s\n", os.Args[0])
	panic("panic!")
}
