package main

//	Go - sync.RWMutexを使おう - Qiita
//	see also: http://qiita.com/y_matsuwitter/items/36565a3a53ac52732cae
import (
	"sync"
	"time"
)

var mu sync.RWMutex
var data map[string]string

func main() {
	data = map[string]string{"hoge": "fuga"}
	mu = sync.RWMutex{}
	go read()
	go read()
	go write()
	go read()
	time.Sleep(10 * time.Second)
}

func read() {
	println("read_start")
	mu.RLock()
	defer mu.RUnlock()
	time.Sleep(1 * time.Second)
	println("read_complete", data["hoge"])
}

func write() {
	println("write_start")
	mu.Lock()
	defer mu.Unlock()
	time.Sleep(2 * time.Second)
	data["hoge"] = "piyo"
	println("write_complete")
}
