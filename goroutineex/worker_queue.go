package goroutineex

import (
	"sync"
)

type WorkerQueue struct {
	workerNum int
	queue     chan func()
	wg        *sync.WaitGroup
	closeOnce sync.Once
}

func NewWorkerQueue(workerNum int, queueSize int) *WorkerQueue {
	if workerNum <= 0 {
		panic("worker number should > 0")
	}
	if queueSize < 0 {
		panic("queue size should >= 0")
	}

	queue := make(chan func(), queueSize)
	wg := sync.WaitGroup{}
	for i := 0; i < workerNum; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for f := range queue {
				f()
			}
		}()
	}

	return &WorkerQueue{
		workerNum: workerNum,
		queue:     queue,
		wg:        &wg,
	}
}

func (wq *WorkerQueue) Submit(f func()) {
	if f != nil {
		wq.queue <- f
	}
}

func (wq *WorkerQueue) Wait() {
	wq.closeOnce.Do(func() {
		close(wq.queue)
	})
	wq.wg.Wait()
}
