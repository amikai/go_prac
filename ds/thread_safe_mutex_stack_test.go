package ds

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreadSafeMutexStack(t *testing.T) {
	s := NewThreadSafeMutexStack[int]()

	_, err := s.Pop()
	assert.ErrorIs(t, err, ErrStackEmpty)

	wg := sync.WaitGroup{}
	for i := 0; i < 1024; i++ {
		i := i
		wg.Add(1)
		go func() {
			s.Push(i)
			wg.Done()
		}()
	}
	wg.Wait()

	collected := make([]int, 1024)
	for i := 1023; i >= 0; i-- {
		i := i
		wg.Add(1)
		go func() {
			ele, err := s.Pop()
			collected[i] = ele
			assert.NoError(t, err)
			wg.Done()
		}()
	}
	wg.Wait()

	_, err = s.Pop()
	assert.ErrorIs(t, err, ErrStackEmpty)
	assert.True(t, s.Empty())

	exp := make([]int, 1024)
	for i := 0; i < 1024; i++ {
		exp[i] = i
	}
	assert.ElementsMatch(t, exp, collected)
}
