package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

var (
	remoteAddr string
	localAddr  string

	closeFlag bool
)

func init() {
	flag.StringVar(&remoteAddr, "r", ":8080", "remote addr")
	flag.StringVar(&localAddr, "l", ":1234", "local addr")
	flag.BoolVar(&closeFlag, "close", true, "r/w close flag")
}

func main() {
	flag.Parse()

	//	remoteTcp, err := net.ResolveTCPAddr("tcp", remoteAddr)
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//	remoteListener, err := net.ListenTCP("tcp", remoteTcp)
	//	if err != nil {
	//		log.Fatalln(err)
	//	}

	localTcp, err := net.ResolveTCPAddr("tcp", localAddr)
	if err != nil {
		log.Fatalln(err)
	}
	localListener, err := net.ListenTCP("tcp", localTcp)
	if err != nil {
		log.Fatalln(err)
	}

	//	controlConn, err := remoteListener.AcceptTCP()
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//	log.Printf("Control Accept\n")
	//

	fmt.Println("losten port", localAddr)

	for cnt := 0; ; cnt++ {
		log.Printf("Local   Accepting(%d)\n", cnt)
		localConn, err := localListener.AcceptTCP()
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("Local   Accept(%d)\n", cnt)
		go func(conn net.Conn) {
			if closeFlag {
				conn.Write([]byte("ok"))
				//				err := conn.Close()
				//				if err != nil {
				//					log.Fatalln(err)
				//				}
				err1 := localConn.CloseRead()
				err2 := localConn.CloseWrite()
				//			err3 := localConn.Close()
				fmt.Println(err1, err2)
				//			fmt.Println(err1, err2, err3)
				//			fmt.Println(err3)
			}
		}(localConn)
	}

	//
	//		log.Printf("Control Writing(%d)\n", cnt)
	//		controlConn.SetWriteDeadline(time.Now().Add(3600 * time.Second))
	//		controlConn.Write([]byte("CONNECT"))
	//		log.Printf("Control Write(%d)\n", cnt)
	//
	//		log.Printf("Remote  Accepting(%d)\n", cnt)
	//		remoteConn, err := remoteListener.AcceptTCP()
	//		if err != nil {
	//			log.Fatalln(err)
	//		}
	//		log.Printf("Remote  Accept(%d)\n", cnt)
	//
	//		log.Printf("Control Reading(%d)\n", cnt)
	//		controlConn.SetReadDeadline(time.Now().Add(3600 * time.Second))
	//		messageBuf := make([]byte, 10)
	//		messageLen, err := controlConn.Read(messageBuf)
	//		if err != nil {
	//			//	fmt.Fprintf(os.Stderr, "Control Read(%d): %s\n", cnt, err.Error())
	//			log.Printf("Control Read(%d): %s\n", cnt, err.Error())
	//			break
	//		}
	//
	//		message := string(messageBuf[:messageLen])
	//		//		fmt.Fprintf(os.Stderr, "Control Read(%d): %s\n", cnt, message)
	//		log.Printf("Control Read(%d): %s\n", cnt, message)
	//
	//		remote := remoteConn
	//		local := localConn
	//
	//		go func(cnt int) {
	//			var remlog *os.File
	//			//			defer local.Close()
	//			defer func() {
	//				//	log.Println("remote CloseRead" , remote.CloseRead())
	//				//	log.Println("local  CloseWrite", local.CloseWrite())
	//				//	log.Println("remote log Close" , remlog.Close())
	//				remote.CloseRead()
	//				local.CloseWrite()
	//				remlog.Close()
	//			}()
	//			//			//remote.SetReadTimeout(120*1E9)
	//			//			//	remote側のプログラム WriteClose()
	//			//			//	このプログラム ReadClose()に相当
	//			//			//	するまで処理をする
	//			//			//	もしくは write するときにエラーが起きるまで
	//			remlog, _ = os.Create(fmt.Sprintf("%04d.res", cnt))
	//			io.Copy(local, io.TeeReader(remote, newTimeWriter(remlog, noTime)))
	//			log.Printf("Read    Close(%d)\n", cnt)
	//		}(cnt)
	//		go func(cnt int) {
	//			var loclog *os.File
	//			//			defer remote.Close()
	//			defer func() {
	//				//	log.Println("remote CloseWrite", remote.CloseWrite())
	//				//	log.Println("local  CloseRead" , local.CloseRead())
	//				//	log.Println("local  log Close" , loclog.Close())
	//				remote.CloseWrite()
	//				local.CloseRead()
	//				loclog.Close()
	//			}()
	//			//local.SetReadTimeout(120*1E9)
	//			//			//	local側のプログラム WriteClose()
	//			//			//	このプログラム ReadClose()に相当
	//			//			//	するまで処理をする(このときremote側は関係ない)
	//			//			//	もしくは write するときにエラーが起きるまで
	//			//	この処理が終了ということは dst が書き込めないか src が読み込めないか
	//			loclog, _ = os.Create(fmt.Sprintf("%04d.req", cnt))
	//			io.Copy(remote, io.TeeReader(local, newTimeWriter(loclog, noTime)))
	//			log.Printf("Write   Close(%d)\n", cnt)
	//		}(cnt)
	//	}
}
