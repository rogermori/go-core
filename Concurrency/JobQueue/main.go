package main

import (
	"fmt"
	"time"
)

func main() {
	jobcount := 100
	workercount := 5
	jobs := make(chan int)
	results := make(chan int)
	
	go func(){
		for i:= 0; i < jobcount ; i++ {
			jobs <- i
		}
	}()
	
	for i := 1; i <= workercount; i++ {
		go worker(i,jobs,results)
	}
	
	for i := 1; i < jobcount; i++ {
		r := <- results
		fmt.Printf("Result %v: job %v\n",i,r)
	}
	
}

func worker(id int, jobs, results chan int){
	for jobid := range jobs {
		fmt.Printf("Worker %v: job %v\n",id,jobid)
		time.Sleep( 1 * time.Second)
		results <- jobid
	}
}