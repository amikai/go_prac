package iterex

import (
	"fmt"
	"iter"
	"slices"
	"strconv"
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

// https://github.com/golang/go/issues/61898
func Filter[V any](f func(V) bool, seq iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq {
			if f(v) && !yield(v) {
				return
			}
		}
		return
	}
}

func TestFilter(t *testing.T) {
	seq := slices.Values([]int{-1, -2, 0, 1, 2})
	positiveSeq := Filter(func(v int) bool { return v > 0 }, seq)
	assert.Equal(t, []int{1, 2}, slices.Collect(positiveSeq))
}

// https://github.com/golang/go/issues/61898
func Map[In, Out any](f func(In) Out, seq iter.Seq[In]) iter.Seq[Out] {
	return func(yield func(Out) bool) {
		for in := range seq {
			if !yield(f(in)) {
				return
			}
		}
		return
	}
}

func TestMap(t *testing.T) {
	seqInt := slices.Values([]int{0, 1, 2})
	seqStr := Map(func(v int) string { return strconv.Itoa(v) }, seqInt)
	assert.Equal(t, []string{"0", "1", "2"}, slices.Collect(seqStr))
}
