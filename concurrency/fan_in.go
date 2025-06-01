package goroutineex

import (
	"sync"
)

func Merge[T any](cs ...chan T) chan T {
	var wg sync.WaitGroup
	out := make(chan T)

	send := func(ch chan T) {
		for item := range ch {
			out <- item
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, ch := range cs {
		go send(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
