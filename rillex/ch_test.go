package rillex

import (
	"sync/atomic"
	"testing"

	"github.com/destel/rill"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMerge(t *testing.T) {
	ch0 := make(chan int)
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		ch0 <- 0
		close(ch0)
	}()
	go func() {
		ch1 <- 1
		close(ch1)
	}()
	go func() {
		ch2 <- 2
		close(ch2)
	}()

	ch := rill.Merge(ch0, ch1, ch2)
	rch := rill.FromChan(ch, nil)
	ret, err := rill.ToSlice(rch)
	require.NoError(t, err)
	assert.ElementsMatch(t, []int{0, 1, 2}, ret)
}

func TestDrain(t *testing.T) {
	ch := make(chan struct{})
	b := &atomic.Bool{}
	go func() {
		ch <- struct{}{}
		close(ch)
		b.Store(true)
	}()
	rill.Drain(ch)
	assert.True(t, b.Load())
}

func TestDrainNB(t *testing.T) {
	ch := make(chan struct{})
	defer rill.DrainNB(ch)
	go func() {
		ch <- struct{}{}
		ch <- struct{}{}
		close(ch)
	}()
	<-ch
}
