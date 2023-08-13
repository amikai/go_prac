package ds

import (
	"errors"
	"slices"
)

// TODO: implement a linked stack

var ErrStackEmpty = errors.New("stack is empty")

type SliceStack[T any] struct {
	stack []T
	top   int
}

var defaultSliceStackSize = 0

func NewSliceStack[T any]() *SliceStack[T] {
	return &SliceStack[T]{
		stack: make([]T, defaultSliceStackSize),
		top:   -1,
	}
}

func (s *SliceStack[T]) Push(val T) {
	s.top++
	if s.top == len(s.stack) {
		s.stack = append(s.stack, val)
	} else {
		s.stack[s.top] = val
	}
}

func (s *SliceStack[T]) Peek() (T, error) {
	var ret T
	if s.top == -1 {
		return ret, ErrStackEmpty
	}
	return s.stack[s.top], nil
}

func (s *SliceStack[T]) Pop() (T, error) {
	var ret T
	if s.top == -1 {
		return ret, ErrStackEmpty
	}

	ret = s.stack[s.top]
	s.top--

	if (s.top+1)*2 < len(s.stack) {
		slices.Delete(s.stack, s.top+1, len(s.stack)-1)
	}
	return ret, nil
}

func (s *SliceStack[T]) Empty() bool {
	return s.top == -1
}
