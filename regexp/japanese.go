package main

import (
	"fmt"
	"regexp"
)

var (
	janReg = regexp.MustCompile(`^[\p{Hiragana}\p{Katakana}\p{Han}]+$`)
)

func main() {
	inputs := []string{
		"日本語",
		"English",
		"0123456789",
		"ABC",
	}
	for i, v := range inputs {
		fmt.Println(i, v, janReg.MatchString(v))
	}
}
