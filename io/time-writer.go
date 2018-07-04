package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

//	FYI:
//	Golang のオフィシャルが提供するインタフェースまとめAdd Star
//	http://mattn.kaoriya.net/software/lang/go/20140501172821.htm

// io.Writer
//type Writer interface {
//    Write(p []byte) (n int, err error)
//}

//	Time Wrapper

type TimeWriter struct {
	io.Writer
}

func (t *TimeWriter) Write(p []byte) (n int, err error) {
	var n1, n2 int
	n1, err = t.Writer.Write([]byte(fmt.Sprintf("%s ", time.Now())))
	if err != nil {
		return
	}
	n2, err = t.Writer.Write(p)
	n = n1 + n2
	//	if you return length of p []byte
	//	invalid WriteString count
	n = n2
	return
}

func main() {
	w := &TimeWriter{os.Stdout}
	strs := []string{
		"sample\n",
		"test\n",
		"piyopiyo\n",
	}
	for _, str := range strs {
		r := strings.NewReader(str)
		io.Copy(w, r)
	}
}
