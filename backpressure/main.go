package main

import (
	"fmt"
	"time"
)

func main() {
	work := make(chan int)

	// 10 worker
	for i := 0; i < 5; i++ {
		go doWork(i, work)
	}

	// generate work
	for i := 0; i < 100; i++ {
		work <- i
	}
	close(work)

	time.Sleep(3 * time.Second)
}

func doWork(id int, work <-chan int) {
	fmt.Printf("Worker %d started.\n", id)
	for i := range work {
		fmt.Printf("Work on %d by %d\n", i, id)
		time.Sleep(time.Duration(100) * time.Millisecond)
	}
	fmt.Printf("Worker %d finsihed.\n", id)
}
