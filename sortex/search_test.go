package sortex

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Use like CPP lower_bound
func TestSortSearchLowerBound(t *testing.T) {
	arr := []int{1, 3, 3, 3, 9}

	// Search uses binary search to find and return the smallest index i in [0, n) at which f(i) is true,
	lowerBoundFunc := func(s []int, target int) (int, func(i int) bool) {
		return len(s), func(i int) bool {
			return s[i] >= target
		}
	}
	// x = 3, p = 1
	// 1 3 3 3 9
	//   ^
	p := sort.Search(lowerBoundFunc(arr, 3))
	assert.Equal(t, 1, p)

	// x = 4 (not in v), p = 4
	// 1 3 3 3 9
	//         ^
	p = sort.Search(lowerBoundFunc(arr, 4))
	assert.Equal(t, 4, p)

	// x = 0 (not in v), p = 0
	// 1 3 3 3 9
	// ^
	p = sort.Search(lowerBoundFunc(arr, 0))
	assert.Equal(t, 0, p)

	// x = 10 (not in v), p = 5
	// 1 3 3 3 9
	//           ^
	p = sort.Search(lowerBoundFunc(arr, 10))
	assert.Equal(t, 5, p)
}

// Use like CPP upper_bound
func TestSortSearchUpperBound(t *testing.T) {
	arr := []int{1, 3, 3, 3, 9}

	lowerBoundFunc := func(s []int, target int) (int, func(i int) bool) {
		return len(s), func(i int) bool {
			return s[i] > target
		}
	}
	// x = 3, p = 4
	// 1 3 3 3 9
	//         ^
	p := sort.Search(lowerBoundFunc(arr, 3))
	assert.Equal(t, 4, p)

	// x = 9, p = 5
	// 1 3 3 3 9
	//           ^
	p = sort.Search(lowerBoundFunc(arr, 9))
	assert.Equal(t, 5, p)

	// x = 0 (not in v), p = 0
	// 1 3 3 3 9
	// ^
	p = sort.Search(lowerBoundFunc(arr, 0))
	assert.Equal(t, 0, p)

	// x = 10 (not in v), p = 5
	// 1 3 3 3 9
	//           ^
	p = sort.Search(lowerBoundFunc(arr, 10))
	assert.Equal(t, 5, p)
}
