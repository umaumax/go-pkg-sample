package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"os"
)

//	go run tee-writer.go a.txt b.txt
//	input:sample
//	input:Ctrl+D
//
//	stdout,a.txt,b.txt:"sample"
var (
	appendOutputFlag bool
)

func init() {
	flag.BoolVar(&appendOutputFlag, "a", false, "Append the output to the files rather than overwriting them.")
}

func main() {
	flag.Parse()

	var err error
	bufferedStdin := bufio.NewReader(os.Stdin)
	bufferedStdout := bufio.NewWriter(os.Stdout)

	var bufferedWriters = []*bufio.Writer{bufferedStdout}
	var writers = []io.Writer{bufferedStdout}

	fileOpenFlag := os.O_CREATE | os.O_RDWR
	if appendOutputFlag {
		fileOpenFlag |= os.O_APPEND
	}

	for _, teeTo := range flag.Args() {
		var w *os.File
		w, err = os.OpenFile(teeTo, fileOpenFlag, 0644)

		if err != nil {
			log.Fatal(err)
		}
		//	NOTE for文中であるが、main関数終了まで開いているひつようがあるので、OK
		defer w.Close()

		bufwriter := bufio.NewWriter(w)
		writers = append(writers, bufwriter)
		bufferedWriters = append(bufferedWriters, bufwriter)
	}

	teeWriter := io.MultiWriter(writers...)

	_, _ = io.Copy(teeWriter, bufferedStdin)
	for _, bw := range bufferedWriters {
		err = bw.Flush()
		if err != nil {
			log.Fatal(err)
		}
	}
}
