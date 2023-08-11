package goroutineex

import (
	"sync"
)

type Limiter struct {
	maxConcurrency int
	limiter        chan struct{}
	queue          chan func()
	wg             *sync.WaitGroup
}

func NewLimiter(maxConcurrency int) *Limiter {
	l := &Limiter{
		maxConcurrency: maxConcurrency,
		limiter:        make(chan struct{}, maxConcurrency),
		queue:          make(chan func(), 1024),
		wg:             &sync.WaitGroup{},
	}

	return l
}

func (l *Limiter) Go(f func()) {
	l.limiter <- struct{}{}
	l.wg.Add(1)
	go func() {
		defer func() {
			<-l.limiter
			l.wg.Done()
		}()
		f()
	}()
}

func (l *Limiter) Wait() {
	l.wg.Wait()
}
