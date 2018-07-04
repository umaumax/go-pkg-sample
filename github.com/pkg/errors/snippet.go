package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	err := fmt.Errorf("This is sample error!")
	err = errors.Wrap(err, "This is additional message!")
	fmt.Printf("%s\n", err)
	err = errors.New("errors.New sample")
	fmt.Printf("%s\n", err)
	fmt.Println("errors.Print")
	errors.Print(err)
}
