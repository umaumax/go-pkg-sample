package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	{
		t, err := time.Parse(
			"2006-01-02 15:04:05 -0700", // スキャンフォーマット
			"2013-06-19 21:54:23 +0900") // パースしたい文字列
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(t)
	}
	{
		t, err := time.Parse(
			"2006-01-02", // スキャンフォーマット
			"1999-01-01") // パースしたい文字列
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(t)
		fmt.Println(t.Format("20060102"))
	}
}
