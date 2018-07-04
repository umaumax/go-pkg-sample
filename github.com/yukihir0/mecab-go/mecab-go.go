package main

import (
	"flag"
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/yukihir0/mecab-go"
)

func main() {
	flag.Parse()
	input := "すもももももももものうち"
	if flag.NArg() > 0 {
		input = flag.Arg(0)
	}
	fmt.Println("input", input)

	args := mecab.NewArgs()
	//	args.DicDir = "/usr/local/Cellar/mecab/0.996/lib/mecab/dic/mecab-ipadic-neologd"
	mecab.Initialize(args)

	nodes, err := mecab.Parse(input)
	if err != nil {
		panic(err)
	}
	pp.Println(nodes)
}
