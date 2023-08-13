package goroutineex

import (
	"sync"
)

type WorkerQueue struct {
	workerNum int
	jobQ      chan func()
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

	wq := &WorkerQueue{
		workerNum: workerNum,
		jobQ:      make(chan func(), queueSize),
		wg:        &sync.WaitGroup{},
	}

	worker := func() {
		defer wq.wg.Done()
		for job := range wq.jobQ {
			job()
		}
	}

	for i := 0; i < wq.workerNum; i++ {
		wq.wg.Add(1)
		go worker()
	}
	return wq

}

func (wq *WorkerQueue) Submit(f func()) {
	if f != nil {
		wq.jobQ <- f
	}
}

func (wq *WorkerQueue) Wait() {
	wq.closeOnce.Do(func() {
		close(wq.jobQ)
	})
	wq.wg.Wait()
}
