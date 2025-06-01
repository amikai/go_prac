package rillex

import (
	"testing"

	"github.com/destel/rill"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReduce(t *testing.T) {
	numbers := rill.FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil)
	sum, ok, err := rill.Reduce(numbers, concLimit, func(x int, y int) (int, error) {
		return x + y, nil
	})
	require.NoError(t, err)
	assert.True(t, ok)
	assert.Equal(t, 55, sum)
}
