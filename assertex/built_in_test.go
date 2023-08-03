package assertex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrueFalse(t *testing.T) {
	assert.True(t, true)
	assert.False(t, false)
}

func TestGreaterLess(t *testing.T) {
	// compare should be the same type
	assert.Greater(t, 2, 1)
	assert.GreaterOrEqual(t, 2, 1)

	assert.Less(t, 1, 2)
	assert.LessOrEqual(t, 2, 2)
}

func TestPosiNeg(t *testing.T) {
	assert.Positive(t, 1)
	assert.Positive(t, 1.0)

	assert.Negative(t, -1)
	assert.Negative(t, -1.0)
}

func TestNil(t *testing.T) {
	assert.Nil(t, nil)
}

func TestPointerReferenceObject(t *testing.T) {
	a := 1
	b := 1
	x := &a
	y := &a
	z := &b

	// asserts that two pointers reference the same object
	// Both x, y reference to a
	assert.Same(t, x, y)

	assert.NotSame(t, x, z)
}
