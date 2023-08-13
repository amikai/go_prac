package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSiceStack(t *testing.T) {
	s := NewSliceStack[int]()

	_, err := s.Pop()
	assert.ErrorIs(t, err, ErrStackEmpty)

	for i := 0; i < 1024; i++ {
		s.Push(i)
		top, err := s.Peek()
		assert.Equal(t, i, top)
		assert.NoError(t, err)
	}

	for i := 1023; i >= 0; i-- {
		top, err := s.Pop()
		assert.Equal(t, i, top)
		assert.NoError(t, err)
	}
	_, err = s.Pop()
	assert.ErrorIs(t, err, ErrStackEmpty)
	assert.True(t, s.Empty())
}
