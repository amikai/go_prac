package rillex

import (
	"sync"
	"testing"

	"github.com/destel/rill"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSplit2(t *testing.T) {
	rch := rill.FromChan(gen(5), nil)
	rchTrue, rchFalse := rill.Split2(rch, concLimit, func(x int) (bool, error) {
		if x == 1 || x == 3 {
			return true, nil
		}
		return false, nil
	})

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		retTrue, err := rill.ToSlice(rchTrue)
		require.NoError(t, err)
		assert.ElementsMatch(t, []int{1, 3}, retTrue)
	}()

	retFalse, err := rill.ToSlice(rchFalse)
	require.NoError(t, err)
	assert.ElementsMatch(t, []int{2, 4, 5}, retFalse)

	wg.Wait()
}

func TestOrderSplit2(t *testing.T) {
	rch := rill.FromChan(gen(5), nil)
	rchTrue, rchFalse := rill.OrderedSplit2(rch, concLimit, func(x int) (bool, error) {
		if x == 1 || x == 3 {
			return true, nil
		}
		return false, nil
	})

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		retTrue, err := rill.ToSlice(rchTrue)
		require.NoError(t, err)
		assert.Equal(t, []int{1, 3}, retTrue)
	}()

	retFalse, err := rill.ToSlice(rchFalse)
	require.NoError(t, err)
	assert.Equal(t, []int{2, 4, 5}, retFalse)

	wg.Wait()
}
