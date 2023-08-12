package goroutineex

import (
	"fmt"
	"testing"
	"time"
)

var workTime = 500 * time.Millisecond
var tick = 100 * time.Millisecond

type workerFunc func(chan struct{})

func doWork(done chan struct{}) {
	time.Sleep(workTime)
	done <- struct{}{}
}

func SelectTimeTick(work workerFunc) {
	done := make(chan struct{}, 1)
	ticker := time.NewTicker(tick)
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

func TestSelectTimeTick(t *testing.T) {
	SelectTimeTick(doWork)
}
