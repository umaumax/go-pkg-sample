package main

import (
	"fmt"
	"os"
)

//	vim cmd
// :GoIferr
func main() {
	fmt.Printf("Hello world\n")

	_ = open("")
}

//	1行開けている場合には挿入する
func open(name string) (err error) {
	f, err := os.Open("")

	defer f.Close()
	return err
}

//	1行開けない場合には挿入しない
// func open(name string) (err error) {
// 	f, err := os.Open("")
// 	defer f.Close()
// 	return err
// }
