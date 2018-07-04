package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	//	If base == 0, the base is implied by the string's prefix: base 16 for "0x", base 8 for "0", and base 10 otherwise.
	fmt.Println(strconv.ParseInt("123.0", 0, 64))
	fmt.Println(strconv.ParseInt("077", 0, 64))
	fmt.Println(strconv.ParseInt("099", 0, 64))
	fmt.Println(strconv.ParseInt("-077", 0, 64))
	fmt.Println(strconv.ParseInt("0xff", 0, 64))
	fmt.Println(strconv.ParseInt("0XFF", 0, 64))
	fmt.Println(strconv.ParseInt("0xBadFace", 0, 64))
	fmt.Println(strconv.ParseInt("123456789", 0, 64))
	fmt.Println(strconv.ParseInt("123456789987654321123456789", 0, 64))
	fmt.Println(strconv.ParseFloat("123", 64))
	fmt.Println(strconv.ParseFloat("123e+1", 64))
	fmt.Println(strconv.ParseFloat("123e", 64))
	fmt.Println(strconv.ParseFloat("0123", 64))
	fmt.Println(strconv.ParseFloat("0x123", 64))
	fmt.Println(strconv.ParseFloat("123.", 64))
	fmt.Println(strconv.ParseFloat(".123", 64))
	fmt.Println(strconv.ParseFloat("-.123", 64))
	fmt.Println(strconv.ParseFloat("-.", 64))
	fmt.Println(strconv.ParseFloat("-.1e1", 64))
	fmt.Println(strconv.ParseFloat("-.1e0", 64))
	fmt.Println(strconv.ParseFloat("1.2.3", 64))
	fmt.Println(strconv.ParseFloat("123e4", 64))
	fmt.Println(strconv.ParseFloat("123e-4", 64))
	fmt.Println(strconv.ParseFloat("123.0", 64))
	fmt.Println(strconv.ParseFloat("123e400", 64))
	fmt.Println(math.Mod(3.55, 2))
	fmt.Println(math.Mod(30.54, 2))
	fmt.Println(math.Mod(3.56, 1))
	fmt.Println(math.Mod(4.55, 3))
	fmt.Println(math.Mod(6.54, 3))
}
