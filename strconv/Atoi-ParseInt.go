package main

import (
	"fmt"
	"strconv"
)

func main() {
	strs := []string{"8", "8)"}
	for i, str := range strs {
		{
			v, err := strconv.ParseInt(str, 0, 64)
			fmt.Println(i, "ParseInt", v, err)
		}
		{
			v, err := strconv.Atoi(str)
			fmt.Println(i, v, err)
		}
	}
}
