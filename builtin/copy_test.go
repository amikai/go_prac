package builtin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopyToOther(t *testing.T) {
	s := make([]int, 3)
	// copy min(len(dst), len(src)) elements from src to dst
	n := copy(s, []int{0, 1, 2, 3})

	assert.Equal(t, 3, n)
	assert.Equal(t, []int{0, 1, 2}, s)
}

func TestCopyFromStringToByteSlice(t *testing.T) {
	b := make([]byte, 5)
	n := copy(b, "Hello, world!")

	assert.Equal(t, 5, n)
	assert.Equal(t, []byte("Hello"), b)
}

func TestCopyToItSelf(t *testing.T) {
	s := []int{0, 1, 2}
	n := copy(s, s[1:]) // n == 2, s == []int{1, 2, 2}

	assert.Equal(t, 2, n)
	assert.Equal(t, []int{1, 2, 2}, s)
}
