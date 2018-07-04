package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

var (
	recursiveFlag bool
	sameFlag      bool
	root          string
)

func init() {
	flag.BoolVar(&recursiveFlag, "r", false, "recursive test flag")
	flag.BoolVar(&sameFlag, "same", false, "same dir open flag")
	flag.StringVar(&root, "root", ".", "root dir")
}

func main() {
	flag.Parse()

	fsw, _ := fsnotify.NewWatcher()
	n := 100000
	cnt := 0

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-fsw.Events:
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}
				eventPath := filepath.Clean(event.Name)
				info, err := os.Stat(eventPath)
				if err != nil {
					log.Println("error:", err)
					break
				}
				if !info.IsDir() {
					break
				}
				err = fsw.Add(eventPath)
				if err != nil {
					log.Println("error:", err)
				}
			case err := <-fsw.Errors:
				log.Println("error:", err)
			}
		}
	}()

	if recursiveFlag {
		_ = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if info == nil {
				return err
			}
			if !info.IsDir() {
				return nil
			}
			err = fsw.Add(path)
			if err != nil {
				fmt.Println(err)
				return err
			}
			cnt++
			fmt.Println(cnt, path)
			return nil
		})
	}

	if sameFlag {
		//	NOTE 同一ファイルをいくら追加してもよい
		for i := 0; i < n; i++ {
			err := fsw.Add(".")
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println(i)
		}
	}

	<-done
}
