package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	//	"time"
)

var (
	addr string

	closeFlag bool
)

func init() {
	flag.StringVar(&addr, "p", ":7777", "server addr")
}

func main() {
	flag.Parse()

	localTcp, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}
	localListener, err := net.ListenTCP("tcp", localTcp)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("losten port", addr)

	for cnt := 0; ; cnt++ {
		log.Printf("Local   Accepting(%d)\n", cnt)
		localConn, err := localListener.AcceptTCP()
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("Local   Accept(%d)\n", cnt)
		go func(cnt int, conn *net.TCPConn) {
			go func() {
				io.Copy(os.Stdout, conn)
				conn.CloseRead()
				log.Println("close read", cnt)
			}()
			go func() {
				conn.Write([]byte("ok"))
				conn.CloseWrite()
				log.Println("close write", cnt)
			}()
		}(cnt, localConn)
	}
}
