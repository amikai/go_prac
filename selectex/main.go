package main

import (
	"fmt"
	"time"
)

var workTime = 5 * time.Second

type workerFunc func(chan struct{})

func doWork(done chan struct{}) {
	time.Sleep(workTime)
	done <- struct{}{}
}

func workWithProgres(work workerFunc) {
	done := make(chan struct{}, 1)
	ticker := time.NewTicker(1 * time.Second)
	go work(done)
	for {
		select {
		case <-done:
			fmt.Println("done")
			return
		case t := <-ticker.C:
			fmt.Println(t)
		}
	}
}

func main() {
	workWithProgres(doWork)
}
