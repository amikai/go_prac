package goroutineex

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLimiter(t *testing.T) {
	limiter := NewLimiter(16)
	var total uint64
	for i := 0; i < 1000000; i++ {
		limiter.Go(func() {
			atomic.AddUint64(&total, 1)
		})
	}
	limiter.Wait()

	assert.Equal(t, uint64(1000000), total)
}

func TestLimiterWait(t *testing.T) {
	limiter := NewLimiter(16)
	for i := 0; i < 100; i++ {
		limiter.Go(func() {
			time.Sleep(100 * time.Millisecond)
		})
	}
	for i := 0; i < 100; i++ {
		go func() {
			limiter.Wait()
		}()
	}
	limiter.Wait()
}
