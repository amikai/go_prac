package rillex

import (
	"testing"

	"github.com/destel/rill"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func gen(n int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 1; i <= n; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func TestFilter(t *testing.T) {
	rch := rill.FromChan(gen(10), nil)
	rch = rill.Filter(rch, concLimit, func(x int) (bool, error) {
		if x%2 == 0 {
			return true, nil
		}
		return false, nil
	})
	ret, err := rill.ToSlice(rch)

	require.NoError(t, err)
	assert.ElementsMatch(t, []int{2, 4, 6, 8, 10}, ret)
}

func TestOrderFilter(t *testing.T) {
	rch := rill.FromChan(gen(10), nil)
	rch = rill.OrderedFilter(rch, concLimit, func(x int) (bool, error) {
		if x%2 == 0 {
			return true, nil
		}
		return false, nil
	})
	ret, err := rill.ToSlice(rch)

	require.NoError(t, err)
	assert.Equal(t, []int{2, 4, 6, 8, 10}, ret)
}
