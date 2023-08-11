package goroutineex

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWorkerQueue(t *testing.T) {
	wq := NewWorkerQueue(16, 0)
	var total uint64
	for i := 0; i < 1000000; i++ {
		wq.Submit(func() {
			atomic.AddUint64(&total, 1)
		})
	}
	wq.Wait()
	assert.Equal(t, uint64(1000000), total)
}

func TestWorkerQueueWait(t *testing.T) {
	wq := NewWorkerQueue(16, 0)
	for i := 0; i < 100; i++ {
		wq.Submit(func() {
			time.Sleep(100 * time.Millisecond)
		})
	}
	for i := 0; i < 100; i++ {
		go func() {
			wq.Wait()
		}()
	}
	wq.Wait()
}
