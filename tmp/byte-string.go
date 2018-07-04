package main

import (
	"fmt"
)

/*
rune=127
[127]
[127]
[127]
rune=128
[194 128]
[194 128]
[128]
rune=255
[195 191]
[195 191]
[255]
rune=256
[196 128]
[0]
[0]
*/
func main() {
	for _, r := range []rune{0x7f, 0x80, 0xff, 0x100} {
		fmt.Printf("rune=%v\n", r)
		{
			str := ""
			str += string(r)
			fmt.Printf("%v\n", []byte(str))
		}
		{
			str := ""
			str += string(byte(r))
			fmt.Printf("%v\n", []byte(str))
		}
		{
			str := ""
			str = string(append([]byte(str), byte(r)))
			fmt.Printf("%v\n", []byte(str))
		}
	}
}
