package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

var (
	taskName string
)

func init() {
	flag.StringVar(&taskName, "task", "A", "task name you want to exec")
}

func main() {
	flag.Parse()

	//	shellで実行するときには、複数の過程を経ているかもしれないので注意
	//	ちなみに、echo $!は毎回動作が異なる
	//	次の場合にも異なる
	//	a=$!
	//	b=$!
	//	echo $a, $b
	fmt.Println("process id", os.Getpid())
	fmt.Println("parent process id", os.Getppid())

	fmt.Printf("task name is %s\n", taskName)
	switch taskName {
	case "A":
		cmd := exec.Command(os.Args[0], "-task", "B")

		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	case "B":
		go func() {
			for _ = range time.Tick(500 * time.Millisecond) {
				ret := isParentProcessRunning()
				if !ret {
					os.Exit(1)
				}
			}
		}()
		cnt := 0
		for _ = range time.Tick(500 * time.Millisecond) {
			fmt.Println(cnt)
			//			os.Stdout.Sync()
			cnt++
		}
	default:
		log.Fatalf("unknown task name : %s", taskName)
	}
	fmt.Println("fin!")
}

func isParentProcessRunning() bool {
	//	NOTE 毎回PPIDを取得すると意図した動作とならない
	//	->本来の親プロセスがkillされると、PIDが1となる(1は特殊なPID)
	ppid := os.Getppid()
	if ppid == 1 {
		return false
	}
	process, err := os.FindProcess(ppid)
	if err != nil {
		return false
	}
	_ = process
	//	NOTE 本来は親プロセスと考えられるプロセスに関して、
	//	特定の実行コマンドで動いているかの確認や、
	//	特定のsignalを送って適切な返答があるかを確認するべき
	return true
}
