package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	
	go func() {
		fmt.Println("Running ... 1")
		time.Sleep(250 * time.Millisecond)
		done <- true
	}()
	
	go func() {
		fmt.Println("Running ... 2")
		time.Sleep(1000 * time.Millisecond)
		done <- true
	}()
	
	waiton(done)
	waiton(done)
	
}

func waiton(done chan bool) {
	select {
	case <- done:
		fmt.Println("Finished")
	case <-	time.After(500 * time.Millisecond) :
		fmt.Println("Timeout")
	}
}
