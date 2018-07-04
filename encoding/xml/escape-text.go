package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

func main() {
	s := "<詳細は別添「サンラビン点滴静注用の調製方法」を参照>"
	err := xml.EscapeText(os.Stdout, []byte(s))
	if err != nil {
		fmt.Println(err)
	}
}
