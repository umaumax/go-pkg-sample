package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
	"time"
)

var (
	exitFlag bool
	killFlag bool
)

func init() {
	flag.BoolVar(&exitFlag, "exit", false, "exit flag while exec")
	flag.BoolVar(&killFlag, "kill", false, "send kill signal to me while exec")
}

//	わかりやすい記事
//	[Ctrl+Cとkill -SIGINTの違いからLinuxプロセスグループを理解する | ギークを目指して](http://equj65.net/tech/linuxprocessgroup/)
func main() {
	flag.Parse()

	//	pgrep, ps aux ...
	//	実行中のgolangプロセスを
	//	kill(SIGHUP, SIGINT) os.Exit() : sleepプロセスはそのまま残る
	//	Ctrl+C : sleepプロセスは消える(実行terminalに依存するらしい)
	cmd := exec.Command("sleep", "100")

	run := func() {
		var out bytes.Buffer
		cmd.Stdout = &out

		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(out.String())
	}

	if killFlag {
		go run()
		time.Sleep(time.Second)
		pid := os.Getpid()
		syscall.Kill(pid, syscall.SIGINT)
		return
	}
	if killFlag {
		go run()
		time.Sleep(time.Second)
		pid := os.Getpid()
		syscall.Kill(pid, syscall.SIGHUP)
		return
	}
	if exitFlag {
		go run()
		time.Sleep(time.Second)
		os.Exit(0)
		return
	}
	run()
}
