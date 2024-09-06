package builtin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	t.Run("min(1,2,3)", func(t *testing.T) {
		got := min(1, 2, 3)
		want := 1
		assert.Equal(t, want, got)
	})

	t.Run("min(1.0,2.0,3.0)", func(t *testing.T) {
		got := min(1.0, 2.0, 3.0)
		want := 1.0
		assert.Equal(t, want, got)
	})

	t.Run("min(1.0,2.0,3)", func(t *testing.T) {
		got := min(1.0, 2.0, 3)
		want := 1.0
		assert.Equal(t, want, got)
	})

	t.Run("min(1,2.0,3.0)", func(t *testing.T) {
		got := min(1, 2.0, 3.0)
		want := 1.0
		assert.Equal(t, want, got)
	})
}

func TestMax(t *testing.T) {
	t.Run("max(1, 2, 3)", func(t *testing.T) {
		got := max(1, 2, 3)
		want := 3
		assert.Equal(t, want, got)
	})

	t.Run("max(1.0, 2.0, 3.0)", func(t *testing.T) {
		got := max(1.0, 2.0, 3.0)
		want := 3.0
		assert.Equal(t, want, got)
	})

	t.Run("max(1.0,2.0,3)", func(t *testing.T) {
		got := max(1.0, 2.0, 3)
		want := 3.0
		assert.Equal(t, want, got)
	})

	t.Run("max(1,2.0,3.0)", func(t *testing.T) {
		got := max(1, 2.0, 3.0)
		want := 3.0
		assert.Equal(t, want, got)
	})
}
