package builtin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// See the SPEC https://go.dev/ref/spec#Comparison_operators

func TestPointerComp(t *testing.T) {
	a := 1
	pa := &a
	pb := &a
	// pa and pb point to same variable
	assert.True(t, pa == pb)
}

func TestPointerNilComp(t *testing.T) {
	var pa *int
	var pb *int
	// pa and pb are nil
	assert.True(t, pa == pb)
}

func TestChComp(t *testing.T) {
	cha := make(chan int)
	chb := cha
	// Two channel values are equal if they were created by the same call to make or if both have value nil.
	assert.True(t, cha == chb)
}

func TestChNilComp(t *testing.T) {
	var cha chan int
	var chb chan int
	// Two channel values are equal if they were created by the same call to make or if both have value nil.
	assert.True(t, cha == chb)
}

// Slice, map, and function types are not comparable.
// However, as a special case, a slice, map, or function value may be compared
// to the predeclared identifier nil.
func TestSliceMapFuncNil(t *testing.T) {
	var nilS []int
	assert.Nil(t, nilS)

	var nilF func()
	assert.Nil(t, nilF)

	var nilM map[int]int
	assert.Nil(t, nilM)
}
