package rillex

import (
	"testing"

	"github.com/destel/rill"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOneIsPositive(t *testing.T) {
	numbers := rill.FromSlice([]int{-1, -2, -3, 4}, nil)
	ok, err := rill.Any(numbers, concLimit, func(x int) (bool, error) {
		return x > 0, nil
	})
	require.NoError(t, err)
	assert.True(t, ok)
}

func TestNoOneIsPositive(t *testing.T) {
	numbers := rill.FromSlice([]int{-1, -2, -3, -4}, nil)
	ok, err := rill.Any(numbers, concLimit, func(x int) (bool, error) {
		return x > 0, nil
	})
	require.NoError(t, err)
	assert.False(t, ok)
}
