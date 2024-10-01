package sortex

import (
	"slices"
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

	upperBoundFunc := func(s []int, target int) (int, func(i int) bool) {
		return len(s), func(i int) bool {
			return s[i] > target
		}
	}
	// x = 3, p = 4
	// 1 3 3 3 9
	//         ^
	p := sort.Search(upperBoundFunc(arr, 3))
	assert.Equal(t, 4, p)

	// x = 9, p = 5
	// 1 3 3 3 9
	//           ^
	p = sort.Search(upperBoundFunc(arr, 9))
	assert.Equal(t, 5, p)

	// x = 0 (not in v), p = 0
	// 1 3 3 3 9
	// ^
	p = sort.Search(upperBoundFunc(arr, 0))
	assert.Equal(t, 0, p)

	// x = 10 (not in v), p = 5
	// 1 3 3 3 9
	//           ^
	p = sort.Search(upperBoundFunc(arr, 10))
	assert.Equal(t, 5, p)
}

// sort.Find is more easy to used than sort.Search
// sort.Find will check the target will exist in the slice or not, but sort.Search will not
func TestSortFindLowerBound(t *testing.T) {
	arr := []int{1, 3, 3, 3, 9}

	// Find uses binary search to find and return the smallest index i in [0, n)
	// at which cmp(i) <= 0. If there is no such index i, Find returns i = n.
	lowerBoundFunc := func(a []int, target int) (int, func(int) int) {
		return len(a), func(i int) int {
			return target - a[i]
		}
	}
	// x = 3, p = 1
	// 1 3 3 3 9
	//   ^
	p, exist := sort.Find(lowerBoundFunc(arr, 3))
	assert.Equal(t, 1, p)
	assert.True(t, exist)

	// x = 4 (not in v), p = 4
	// 1 3 3 3 9
	//         ^
	p, exist = sort.Find(lowerBoundFunc(arr, 4))
	assert.Equal(t, 4, p)
	assert.False(t, exist)

	// x = 0 (not in v), p = 0
	// 1 3 3 3 9
	// ^
	p, exist = sort.Find(lowerBoundFunc(arr, 0))
	assert.Equal(t, 0, p)
	assert.False(t, exist)

	// x = 10 (not in v), p = 5
	// 1 3 3 3 9
	//           ^
	p, exist = sort.Find(lowerBoundFunc(arr, 10))
	assert.Equal(t, 5, p)
	assert.False(t, exist)
}

// slice.BinarySearch work likes lower_bound in C++
func TestSliceBinarySearch(t *testing.T) {
	arr := []int{1, 3, 3, 3, 9}

	// x = 3, p = 1
	// 1 3 3 3 9
	//   ^
	p, exist := slices.BinarySearch(arr, 3)
	assert.Equal(t, 1, p)
	assert.True(t, exist)

	// x = 4 (not in v), p = 4
	// 1 3 3 3 9
	//         ^
	p, exist = slices.BinarySearch(arr, 4)
	assert.Equal(t, 4, p)
	assert.False(t, exist)

	// x = 0 (not in v), p = 0
	// 1 3 3 3 9
	// ^
	p, exist = slices.BinarySearch(arr, 0)
	assert.Equal(t, 0, p)
	assert.False(t, exist)

	// x = 10 (not in v), p = 5
	// 1 3 3 3 9
	//           ^
	p, exist = slices.BinarySearch(arr, 10)
	assert.Equal(t, 5, p)
	assert.False(t, exist)
}
