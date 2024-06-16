package iterex

import (
	"fmt"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleSliceAll() {
	arr := []int{0, 1, 2, 3, 4}
	for i, v := range slices.All(arr) {
		fmt.Printf("%d, %d\n", i, v)
	}

	// Output:
	// 0, 0
	// 1, 1
	// 2, 2
	// 3, 3
	// 4, 4
}

func ExampleSliceAll2() {
	arr := []int{0, 1, 2, 3, 4}
	slices.All(arr)(func(i int, v int) bool {
		fmt.Printf("%d, %d\n", i, v)
		return true
	})

	// Output:
	// 0, 0
	// 1, 1
	// 2, 2
	// 3, 3
	// 4, 4
}

func TestAppendSeq(t *testing.T) {
	arr := []int{0, 1, 2}
	seq := slices.Values([]int{3, 4, 5})
	arr = slices.AppendSeq(arr, seq)

	assert.Equal(t, []int{0, 1, 2, 3, 4, 5}, arr)
}

func ExampleSliceBackward() {
	arr := []int{0, 1, 2, 3, 4}
	for i, v := range slices.Backward(arr) {
		fmt.Printf("%d, %d\n", i, v)
	}

	// Output:
	// 4, 4
	// 3, 3
	// 2, 2
	// 1, 1
	// 0, 0
}

func ExampleSliceCollect() {
	seq := slices.Values([]int{0, 1, 2})
	// convert seq to slice
	arr := slices.Collect(seq)
	for i, v := range arr {
		fmt.Printf("%d, %d\n", i, v)
	}

	// Output:
	// 0, 0
	// 1, 1
	// 2, 2
}

func ExampleSliceSorted() {
	seq := slices.Values([]int{0, 2, 1})
	// convert seq to slice
	arr := slices.Sorted(seq)
	for i, v := range arr {
		fmt.Printf("%d, %d\n", i, v)
	}

	// Output:
	// 0, 0
	// 1, 1
	// 2, 2
}
