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
	for range 1000000 {
		limiter.Go(func() {
			atomic.AddUint64(&total, 1)
		})
	}
	limiter.Wait()

	assert.Equal(t, uint64(1000000), total)
}

func TestLimiterWait(t *testing.T) {
	limiter := NewLimiter(16)
	for range 100 {
		limiter.Go(func() {
			time.Sleep(100 * time.Millisecond)
		})
	}
	for range 100 {
		go func() {
			limiter.Wait()
		}()
	}
	limiter.Wait()
}
