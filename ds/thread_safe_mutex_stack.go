package ds

import (
	"slices"
	"sync"
)

type ThreadSafeMutexStack[T any] struct {
	stack []T
	top   int
	rwmu  *sync.RWMutex
}

var defaultThreadSafeMutexStackSize = 0

func NewThreadSafeMutexStack[T any]() *ThreadSafeMutexStack[T] {
	return &ThreadSafeMutexStack[T]{
		stack: make([]T, defaultThreadSafeMutexStackSize),
		top:   -1,
		rwmu:  &sync.RWMutex{},
	}
}

func (s *ThreadSafeMutexStack[T]) Push(val T) {
	s.rwmu.Lock()
	defer s.rwmu.Unlock()
	s.top++
	if s.top == len(s.stack) {
		s.stack = append(s.stack, val)
	} else {
		s.stack[s.top] = val
	}
}

func (s *ThreadSafeMutexStack[T]) Peek() (T, error) {
	var ret T
	s.rwmu.RLock()
	defer s.rwmu.RUnlock()
	if s.top == -1 {
		return ret, ErrStackEmpty
	}
	return s.stack[s.top], nil
}

func (s *ThreadSafeMutexStack[T]) Pop() (T, error) {
	var ret T
	s.rwmu.RLock()
	if s.top == -1 {
		s.rwmu.RUnlock()
		return ret, ErrStackEmpty
	}
	s.rwmu.RUnlock()

	s.rwmu.Lock()
	defer s.rwmu.Unlock()

	ret = s.stack[s.top]
	s.top--
	if (s.top+1)*2 < len(s.stack) {
		s.stack = slices.Delete(s.stack, s.top+1, len(s.stack)-1)
	}
	return ret, nil
}

func (s *ThreadSafeMutexStack[T]) Empty() bool {
	s.rwmu.RLock()
	defer s.rwmu.RUnlock()
	return s.top == -1
}
