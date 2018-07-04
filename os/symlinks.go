package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var ()

func main() {
	flag.Parse()

	args := flag.Args()
	for _, name := range args {
		fi, err := os.Lstat(name)
		if err != nil {
			log.Fatalln(err)
		}

		if fi.Mode()&os.ModeSymlink == os.ModeSymlink {
			realPath, err := os.Readlink(name)
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Printf("%s is a symbolic file pointing to %s\n", name, realPath)
		} else {
			fmt.Printf("%s is a not symbolic file\n", name)
		}
	}
}
