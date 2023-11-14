package testingex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// See good testing table article:
// https://dave.cheney.net/2019/05/07/prefer-table-driven-tests
func equal[T comparable](a, b T) bool {
	return a == b
}

func TestEqual(t *testing.T) {
	tests := map[string]struct {
		inputA int
		inputB int
		want   bool
	}{
		"a = b": {inputA: 1, inputB: 1, want: true},
		"a > b": {inputA: 2, inputB: 1, want: false},
		"a < b": {inputA: 1, inputB: 2, want: false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := equal[int](tc.inputA, tc.inputB)
			assert.Equal(t, tc.want, got)
		})
	}
}
