package builtin

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// also see
// 1. https://github.com/golang/go/wiki/SliceTricks
// 2. https://ueokande.github.io/go-slice-tricks/

func TestSliceCompare(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	assert.True(t, slices.Equal(a, b))
}

func TestSliceCompareFunc(t *testing.T) {
	pa := []People{
		{
			Name: "John",
			Age:  20,
		},
		{
			Name: "Mary",
			Age:  30,
		},
		{
			Name: "Bob",
			Age:  40,
		},
	}
	pb := []People{
		{
			Name: "John",
			Age:  20,
		},
		{
			Name: "Mary",
			Age:  30,
		},
		{
			Name: "Bob",
			Age:  40,
		},
	}
	assert.True(t, slices.EqualFunc(pa, pb, func(a, b People) bool {
		return a.Equal(b)
	}))
}

func TestSliceCompact(t *testing.T) {
	input := []int{1, 1, 1, 1, 2, 3, 4, 4, 4, 4}
	// Notice: compact modify the input
	got := slices.Compact(input)
	exp := []int{1, 2, 3, 4}

	assert.Equal(t, exp, got)
}

func TestSliceMin(t *testing.T) {
	input := []int{3, 2, 1}

	got := slices.Min(input)
	exp := 1
	assert.Equal(t, exp, got)
}

func TestSliceMinEmpty(t *testing.T) {
	var input []int
	assert.Panics(t, func() { _ = slices.Min(input) })
}

func TestSliceMax(t *testing.T) {
	input := []int{3, 2, 1}

	got := slices.Max(input)
	exp := 3
	assert.Equal(t, exp, got)
}

func TestSliceMaxEmpty(t *testing.T) {
	var input []int
	assert.Panics(t, func() { _ = slices.Max(input) })
}

func TestSliceClone(t *testing.T) {
	s := []int{1, 2, 3}
	sc := slices.Clone(s)
	assert.Equal(t, s, sc)
	// check the slice is not same
	assert.NotSame(t, &s, &sc)
}

func TestSliceContain(t *testing.T) {
	s := []int{1, 2, 3}
	assert.True(t, slices.Contains(s, 1))
	assert.True(t, slices.Contains(s, 2))
	assert.False(t, slices.Contains(s, 5))
}

func TestSliceDelete(t *testing.T) {
	s := []int{1, 2, 3}
	assert.Equal(t, []int{1, 3}, slices.Delete(s, 1, 2))
}

func TestSliceConcat(t *testing.T) {
	var s0 []int
	s1 := []int{1, 2}
	s2 := []int{3, 4}
	s3 := []int{5, 6}
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, slices.Concat(s0, s1, s2, s3))
}
