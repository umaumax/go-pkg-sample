package main

import (
	"fmt"
	"strings"
)

func main() {
	strs := []string{
		" ", "\t", "\n", "\n\n",
	}
	for _, str := range strs {
		fmt.Printf("[%s]->[%s]\n", str, strings.TrimSpace(str))
	}
}
