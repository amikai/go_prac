package main

import (
	"fmt"
	"time"
)

func doSomething() {
	time.Sleep(2 * time.Second)
}

func main() {
	ch := make(chan struct{}, 1)
	go func() {
		doSomething()
		ch <- struct{}{}
	}()
	select {
	case <-ch:
		fmt.Printf("doSomething")
	case <-time.After(3 * time.Second):
		fmt.Printf("after 3 second")
	}
}
