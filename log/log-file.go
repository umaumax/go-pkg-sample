package main

import (
	"log"
	"os"
)

func main() {
	//	追記モード
	file, err := os.OpenFile("test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()
	log.SetOutput(file)
	log.Println("LogSample")

	logger := log.New(file, "[Tag]", log.LstdFlags)
	logger.Print("Test")
}
