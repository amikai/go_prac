// The contents of this file are entirely experimental, unverified, and not robust.
package iterex

import (
	"iter"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"
)

func gen(n int) chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 1; i <= n; i++ {
			ch <- i
		}
	}()
	return ch
}

func All[Ch chan E, E any](ch Ch) iter.Seq[E] {
	return func(yield func(E) bool) {
		defer func() {
			// drain the channel
			for range ch {
			}
		}()
		for v := range ch {
			if !yield(v) {
				return
			}
		}
	}
}

func TestChAll(t *testing.T) {
	defer goleak.VerifyNone(t)
	ch := gen(5)
	s := slices.Collect(All(ch))
	assert.ElementsMatch(t, []int{1, 2, 3, 4, 5}, s)
}

func TestChAllLoopBreak(t *testing.T) {
	defer goleak.VerifyNone(t)
	ch := gen(5)
	for v := range All(ch) {
		if v == 3 {
			break
		}
	}
}
