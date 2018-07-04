package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	//  -hオプション用文言
	flag.Usage = func() {
		//  os.Args[0]  実行時コマンド名
		fmt.Fprintf(os.Stderr, `
Usage of %s:
   %s [OPTIONS] ARGS...
Options`+"\n", os.Args[0], os.Args[0])
		flag.PrintDefaults()
	}

	//  コマンドラインオプションをコマンドライン引数の後に配置するためにFlagSet型を用いる
	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	//  フラグ設定
	var (
		// opt1 = flag.String("opt1", "default-value", "First string option")
		// opt2 = flag.String("opt2", "default-value", "Second string option")
		opt1 = fs.String("opt1", "default-value", "First string option")
		opt2 = fs.String("opt2", "default-value", "Second string option")
	)

	var arg1 string
	var args []string
	//  フラグ設定後にパース
	if len(os.Args) == 1 {
		flag.Parse()
		args = flag.Args()
	} else {
		//  第一引数を取る場合
		n := 1
		arg1 = os.Args[n]
		fs.Parse(os.Args[n+1:])
		args = fs.Args()
	}

	fmt.Println("arg1:", arg1)
	fmt.Println("opt1:", *opt1)
	fmt.Println("opt2:", *opt2)
	fmt.Println("args:", args)
}
