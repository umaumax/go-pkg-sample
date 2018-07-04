package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var (
	outFilePath string
)

func init() {
	flag.StringVar(&outFilePath, "o", "tmp.zip", "out zip filepath")
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("You forget input args...")
		return
	}

	f, err := os.Create(outFilePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	zw := zip.NewWriter(f)

	err = func() (err error) {
		var fi *os.File
		var fw io.Writer
		for i, name := range args {
			log.Println(i, name)
			fi, err = os.Open(name)
			if err != nil {
				return
			}
			defer fi.Close()

			header := &zip.FileHeader{
				Name:         name,
				Method:       zip.Store,
				ModifiedTime: uint16(time.Now().UnixNano()),
				ModifiedDate: uint16(time.Now().UnixNano()),
			}

			fw, err = zw.CreateHeader(header)
			if err != nil {
				return
			}

			_, err = io.Copy(fw, fi)
			if err != nil {
				return
			}
		}
		return
	}()
	if err != nil {
		log.Fatalln(err)
	}

	if err := zw.Close(); err != nil {
		log.Fatalln(err)
	}
}
