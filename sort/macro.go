package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	var datas []byte
	var err error

	if datas, err = ioutil.ReadAll(os.Stdin); err != nil {
		log.Fatalln(err)
	}

	replaceMap := map[string]string{
		"PackageName": "intkv",
		"KeyType":     "string",
		"ValueType":   "int",
	}
	replaceSlice := make([]string, len(replaceMap)*2)
	i := 0
	for k, v := range replaceMap {
		replaceSlice[i] = "?" + k + "?"
		replaceSlice[i+1] = v
		i += 2
	}
	replacer := strings.NewReplacer(replaceSlice...)

	prefix := ""
	output := func(s string) {
		fmt.Printf("%s%s\n", prefix, s)
	}

	macroBuf := make([]string, 1000)
	macroBufIndex := 0
	skipIndex := 0
	for _, line := range regexp.MustCompile("\r\n|\n\r|\n|\r").Split(string(datas), -1) {
		if skipIndex > 0 {
			skipIndex--
			continue
		}
		if strings.HasPrefix(strings.TrimSpace(line), "///") {
			line = strings.TrimSpace(line)[3:]
			if line == "#OFF" {
				prefix = "//"
				continue
			}
			if line == "#ON" {
				prefix = ""
				continue
			}
			macroBuf[macroBufIndex] = line
			macroBufIndex++
		} else if macroBufIndex > 0 {
			for _, macroLine := range macroBuf[:macroBufIndex] {
				macroLine = replacer.Replace(macroLine)
				output(fmt.Sprint(macroLine, "//	replaced"))
			}
			skipIndex = macroBufIndex - 1
			macroBufIndex = 0
		} else {
			output(line)
		}
	}
}
