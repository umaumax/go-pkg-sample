package main

import (
	"flag"
	"log"
	"net"
	"sync"
	"time"
)

var (
	connectN   int
	clientN    int
	remoteAddr string
)

func init() {
	flag.IntVar(&clientN, "c", 1, "client num")
	flag.IntVar(&connectN, "n", 10, "connect num")
	flag.StringVar(&remoteAddr, "r", ":1234", "remote addr")
}

func main() {
	flag.Parse()

	{
		//		remoteAddr := "www.example.com:443"
		remoteAddr := ":7777"
		remoteTCP, err := net.ResolveTCPAddr("tcp", remoteAddr)
		if err != nil {
			log.Fatalln(err)
		}
		conn, err := net.DialTCP("tcp", nil, remoteTCP)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("sleep.")
		time.Sleep(5 * time.Second)
		log.Println("connect.")
		conn.SetDeadline(time.Now().Add(3 * time.Second))
		b := make([]byte, 1024)
		n, err := conn.Read(b)
		_ = n
		if err != nil {
			log.Fatalln(err)
		}
		err = conn.Close()
		if err != nil {
			log.Println(err)
		}
		log.Println("fin.")
	}
	return

	remoteTCP, err := net.ResolveTCPAddr("tcp", remoteAddr)
	if err != nil {
		log.Fatalln(err)
	}

	var wg sync.WaitGroup
	wg.Add(clientN)
	for i := 0; i < clientN; i++ {
		go func() {
			for i := 0; i < connectN; i++ {
				conn, err := net.DialTCP("tcp", nil, remoteTCP)
				if err != nil {
					log.Println(err)
					i--
					continue
				}
				conn.SetDeadline(time.Now().Add(5 * time.Second))
				b := make([]byte, 1024)
				n, err := conn.Read(b)
				_ = n
				if err != nil {
					log.Fatalln(err)
				}
				err = conn.Close()
				if err != nil {
					log.Println(err)
				}
			}
			wg.Done()
			log.Println("go func fin.")
		}()
	}
	wg.Wait()
	log.Println("fin.")
}
