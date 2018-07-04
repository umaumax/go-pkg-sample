package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	src := []byte(`</form>
<div align="right">

    1～10件（15件中）

    <div>
        <table width="100%" cellspacing="0" cellpadding="0" border="0" style="margin-top:2px;"></table>
        <div align="center">

            ページ 1  `)
	fmt.Println(Count(src))
}

func Count(src []byte) (start, end, max int, err error) {
	countReg := regexp.MustCompile(`(\d+)～(\d+)件（(\d+)件中）`)
	dst := countReg.FindSubmatch(src)
	if len(dst) != 4 {
		err = fmt.Errorf("not match, match num is %d, required 4", len(dst))
		return
	}
	start, _ = strconv.Atoi(string(dst[1]))
	end, _ = strconv.Atoi(string(dst[2]))
	max, _ = strconv.Atoi(string(dst[3]))
	return
}
