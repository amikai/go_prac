package iterex

import (
	"iter"
	"slices"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://github.com/golang/go/issues/61898
func Filter[V any](f func(V) bool, seq iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq {
			if f(v) && !yield(v) {
				return
			}
		}
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
	}
}

func TestMap(t *testing.T) {
	seqInt := slices.Values([]int{0, 1, 2})
	seqStr := Map(func(v int) string { return strconv.Itoa(v) }, seqInt)
	assert.Equal(t, []string{"0", "1", "2"}, slices.Collect(seqStr))
}
