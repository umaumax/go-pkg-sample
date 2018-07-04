package main

import (
	"fmt"
	"log"
	"math/big"
)

func main() {
	ScanInt()
	ScanRat()
	SetStringInt()
	SetStringRat()
}

func ScanInt() {
	i := new(big.Int)
	_, err := fmt.Sscan("18446744073709551617", i)
	if err != nil {
		log.Println("error scanning value:", err)
	} else {
		fmt.Println(i)
	}
}
func ScanRat() {
	r := new(big.Rat)
	_, err := fmt.Sscan("1.5000", r)
	if err != nil {
		log.Println("error scanning value:", err)
	} else {
		fmt.Println(r)
	}
}
func SetStringInt() {
	i := new(big.Int)
	i.SetString("644", 8) // octal
	fmt.Println(i)
}
func SetStringRat() {
	r := new(big.Rat)
	r.SetString("355/113")
	fmt.Println(r.FloatString(3))
}
