package goroutineex

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func doSomething() {
	time.Sleep(2 * time.Second)
}

func TestTimeAfter(t *testing.T) {
	ch := make(chan struct{}, 1)
	var final string

	go func() {
		doSomething()
		ch <- struct{}{}
	}()
	select {
	case <-ch:
		final = "ch"
	case <-time.After(3 * time.Second):
		final = "time.After"
	}
	assert.Equal(t, "ch", final)
}
