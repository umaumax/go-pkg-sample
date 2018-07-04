package main

import (
	"fmt"
	"log"
	"os/user"
)

func main() {
	//	current user
	u, err := user.Current()
	//	u, err = user.Lookup("user")
	//	u, err = user.Lookup("root")
	u, err = user.LookupId("0")
	u, err = user.LookupId("1")
	if err != nil {
		log.Fatalln(err)
	}
	/*	----
		20
		/Users/user
		501
		user
		user
	*/
	/*	sudo
		0
		/var/root
		0
		System Administrator
		root
	*/
	fmt.Println("primary group id:", u.Gid)
	fmt.Println("home dir", u.HomeDir)
	fmt.Println("user id", u.Uid)
	fmt.Println("name", u.Name)
	fmt.Println("user name", u.Username)
}
