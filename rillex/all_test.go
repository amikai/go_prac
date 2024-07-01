package rillex

import (
	"testing"

	"github.com/destel/rill"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAllIsPositive(t *testing.T) {
	numbers := rill.FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil)
	ok, err := rill.All(numbers, concLimit, func(x int) (bool, error) {
		return x > 0, nil
	})
	require.NoError(t, err)
	assert.True(t, ok)
}

func TestAllIsPositive2(t *testing.T) {
	numbers := rill.FromSlice([]int{-1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil)
	ok, err := rill.All(numbers, concLimit, func(x int) (bool, error) {
		return x > 0, nil
	})
	require.NoError(t, err)
	assert.False(t, ok)
}
