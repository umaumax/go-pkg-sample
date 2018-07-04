package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var (
	root string
)

func init() {
	flag.StringVar(&root, "r", ".", "root")
}
func main() {
	//	rootにはディレクトリだけではなく、ファイルも指定可能
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		fmt.Println(path, info.Name(), err)
		return nil
	})
	if err != nil {
		log.Fatalln(err)
	}
}
