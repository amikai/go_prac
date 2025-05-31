package interning

import (
	"testing"
	"unique"

	"github.com/stretchr/testify/assert"
)

func TestCompareHandle(t *testing.T) {
	s1 := unique.Make("abcdefghijklmnoppqrstuvwxyz")
	s2 := unique.Make("abcdefghijklmnoppqrstuvwxyz")
	assert.True(t, s1 == s2)

	assert.Equal(t, "abcdefghijklmnoppqrstuvwxyz", s1.Value())
	assert.Equal(t, "abcdefghijklmnoppqrstuvwxyz", s2.Value())
}

func BenchCompare(b *testing.B) {
	// TODO: compare string interning compare and pure string compare
}
