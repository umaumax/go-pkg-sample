package main

import (
	"fmt"
	"regexp"
)

/*
try on it
https://regex-golang.appspot.com/assets/html/index.html
but
	str := regexp.MustCompile(".+?").FindString("nya-")
	fmt.Println(str)
	return
	これはサイトだとnya-,本来はn
docment
https://code.google.com/p/re2/wiki/Syntax
*/

func main() {
	//	ニコニコ動画のコメント改変サンプル
	//	後方検索でないので修正
	//chatReg := regexp.MustCompile(">(.*?)</chat>")
	chatReg := regexp.MustCompile(">([^>]*?)</chat>")
	src := []byte(`<?xml version="1.0" encoding="UTF-8"?><packet><thread resultcode="0" thread="1326108478" last_res="223" ticket="0x355f0000" revision="1" server_time="1425733939"/><leaf thread="1326108478" count="139"/><leaf thread="1326108478" leaf="1" count="84"/><view_counter video="22465" id="sm16654069" mylist="1302"/><global_num_res thread="1326108478" num_res="223"/><thread resultcode="0" thread="1326108478" last_res="223" ticket="0x355f0000" revision="1" server_time="1425733939"/><chat thread="1326108478" no="1" vpos="2245" date="1326108738" mail="184" user_id="dFvy1h5EQQhlG0ch42uCL0VTiR4" premium="1" anonymity="1">うぽつですー！</chat><chat thread="1326108478" no="2" vpos="1209" date="1326109339" mail="184" user_id="qu13t2q5uVP0thV5ffKtz7SPIfc" premium="1" anonymity="1">ビーマイビー！</chat></packet>`)
	dst := chatReg.ReplaceAllFunc(src, func(src []byte) (dst []byte) {
		comment := string(src[1 : len(src)-len("</chat>")])
		fmt.Println(comment)
		comment = "こんにちは～！"
		dst = []byte(">" + comment + "</chat>")
		return
	})
	fmt.Println()
	fmt.Println(string(dst))
}
