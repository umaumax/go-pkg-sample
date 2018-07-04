package main

import (
	"bytes"
	"io"
	"os"
)

func main() {
	readNTimes(os.Stdin)
}

//	io.Readerから同じ内容を複数回読み込むサンプル
//	r1とr2を交互に読み込むことで期待通りのサンプルとなる
//	通常、1つのbytes.Bufferを用いている時にはwriteしてreadした後は読み捨てられる
func readNTimes(r io.Reader) {
	var w1 bytes.Buffer
	var w2 bytes.Buffer
	io.Copy(&w1, r)

	r1 := io.TeeReader(&w1, &w2)
	r2 := io.TeeReader(&w2, &w1)

	getReader := func() io.Reader {
		tmp := r1
		r1, r2 = r2, r1
		return tmp
	}

	io.Copy(os.Stdout, getReader())
	io.Copy(os.Stdout, getReader())
	io.Copy(os.Stdout, getReader())
	io.Copy(os.Stdout, getReader())
}
