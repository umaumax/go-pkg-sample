package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println(os.Getegid())
	fmt.Println(os.Geteuid())
	fmt.Println(os.Getgid())
	fmt.Println(os.Getgroups())
	fmt.Println(os.Getpagesize())
	fmt.Println(os.Getpid())
	fmt.Println(os.Getppid())
	fmt.Println(os.Getuid())
	fmt.Println(os.Getwd())
	fmt.Println(os.Hostname())
	fmt.Println(exec.LookPath("ls"))
	fmt.Println(exec.LookPath("go"))
	fmt.Println(exec.LookPath("doc.go"))
	fmt.Println(exec.LookPath("main.go"))
}
