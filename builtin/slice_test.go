package builtin

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slices"
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
