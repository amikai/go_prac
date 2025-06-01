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
	for range 1000000 {
		wq.Submit(func() {
			atomic.AddUint64(&total, 1)
		})
	}
	wq.Wait()
	assert.Equal(t, uint64(1000000), total)
}

func TestWorkerQueueWait(t *testing.T) {
	wq := NewWorkerQueue(16, 0)
	for range 100 {
		wq.Submit(func() {
			time.Sleep(100 * time.Millisecond)
		})
	}
	for range 100 {
		go func() {
			wq.Wait()
		}()
	}
	wq.Wait()
}
