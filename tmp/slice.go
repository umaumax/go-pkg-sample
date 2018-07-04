package main

import (
	"fmt"
)

func main() {
	str := []rune("ABCDEFGHIJKLMN")
	a := str[:4]
	b := str[4:]
	fmt.Printf("a:[%s], b:[%s]\n", string(a), string(b))
	//	NOTE capacityに余裕がある場合には、意図した動作に貼らないので注意
	a = append(a, 'Z')
	fmt.Printf("a=append(a,'Z')\n")
	fmt.Printf("a:[%s], b:[%s]\n", string(a), string(b))
}
