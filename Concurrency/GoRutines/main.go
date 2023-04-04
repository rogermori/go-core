package main

import (
	"fmt"
	"time"
)

func main() {
	go say("go")
	go say("for it")
	var in string
	fmt.Scanln(&in)
}

func say(message string){
	for i:= 0 ; i < 10; i++ {
		time.Sleep(100 * time.Millisecond) // perform a big task (e.g HTTP Request, etc, etc)
		fmt.Println(message)
	}
}