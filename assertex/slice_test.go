package assertex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceEmptyOrNil(t *testing.T) {
	var nilSlice []any
	assert.Empty(t, nilSlice)

	emptySlice := []any{}
	assert.Empty(t, emptySlice)

	nonEmptySlice := []any{1}
	assert.NotEmpty(t, nonEmptySlice)
}

func TestSliceContain_(t *testing.T) {
	s := []any{1, "a", 1.0, []int{1, 2, 3}, struct{}{}}
	assert.Contains(t, s, 1)
	assert.Contains(t, s, "a")
	assert.Contains(t, s, 1.0)
	assert.Contains(t, s, []int{1, 2, 3})
	assert.Contains(t, s, struct{}{})
}

func TestSliceEmelemtMatch(t *testing.T) {
	s1 := []any{1, 1.0, "a"}
	s2 := []any{"a", 1.0, 1}
	assert.ElementsMatch(t, s1, s2)
}

func TestSliceSubset(t *testing.T) {
	assert.Subset(t, []int{1, 2, 3}, []int{1, 3})
	assert.Subset(t, []int{1, 2, 3}, []int{3})
	assert.NotSubset(t, []int{1, 2, 3}, []int{3, 4})
}

func TestSliceLen(t *testing.T) {
	s := make([]any, 10)
	assert.Len(t, s, 10)
}
