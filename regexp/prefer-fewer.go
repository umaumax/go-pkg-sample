package main

import (
	"fmt"
	"regexp"
)

func main() {
	//	x*?            zero or more x, prefer fewer
	//	x+?            one or more x, prefer fewer
	type SSTuple struct {
		A, B string
	}
	m := map[SSTuple][]string{
		SSTuple{`.*A`, `.*?A`}: {
			"aaAbbA", "aaA",
			"", "",
		},
	}
	for tuple, datas := range m {
		pt1, pt2 := tuple.A, tuple.B
		reg := regexp.MustCompile(pt2)
		for i := 0; i < len(datas); i += 2 {
			question := datas[i]
			answer := datas[i+1]
			result := reg.FindString(question)
			if result != answer {
				fmt.Println("regexp(origin)    noooooooo!", pt2, question, answer, result)
			}
			result = fewerReg(pt1, question)
			if result != answer {
				fmt.Println("regexp(my-method) noooooooo!", pt1, question, answer, result)
			}
		}
	}
}

//	最短一致(all)を最長一致で実現
//	(部分的な最長一致はできない)
//	最長一致をした結果を後ろから1文字削って最長一致を行う
func fewerReg(pattern, text string) (result string) {
	reg := regexp.MustCompile(pattern)
	for {
		newText := reg.FindString(text)
		if newText == "" {
			return
		}
		result = newText
		text = string([]rune(newText)[:len(newText)-1])
	}
	return
}
