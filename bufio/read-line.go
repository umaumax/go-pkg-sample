package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	path := "read-line.go"

	f, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	i := 0
	for scanner.Scan() {
		i++
		line := scanner.Text()
		fmt.Println(i, line)
	}
	if err = scanner.Err(); err != nil {
		log.Fatalln(err)
	}
}
