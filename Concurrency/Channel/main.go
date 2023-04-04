package main

import "fmt"

func main() {
	c := make(chan string, 4)
	c <- "Hello"
	c <- "World"
	
	v := <- c
	fmt.Println(v)
	fmt.Println(<-c)
	
}