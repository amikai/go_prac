package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSiceQueue(t *testing.T) {
	q := NewSliceQueue[int]()

	_, err := q.Dequeue()
	assert.ErrorIs(t, err, ErrQueueEmpty)

	for i := 0; i < 1024; i++ {
		q.Enqueue(i)
		top, err := q.Peek()
		assert.Equal(t, 0, top)
		assert.NoError(t, err)
	}

	for i := 0; i < 1024; i++ {
		top, err := q.Dequeue()
		assert.Equal(t, i, top)
		assert.NoError(t, err)
	}
	_, err = q.Dequeue()
	assert.ErrorIs(t, err, ErrQueueEmpty)
	assert.True(t, q.Empty())
}
