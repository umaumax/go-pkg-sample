package main

import (
	"fmt"
	"strings"
)

func main() {
	texts := []string{
		"This is an apple!",
		"Book!",
	}
	tmp := strings.Join(texts, string(0x0))
	out := strings.Split(tmp, string(0x0))
	fmt.Printf(strings.Join(out, "@"))
}
