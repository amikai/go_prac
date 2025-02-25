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
}

func BenchCompare(b *testing.B) {
	// TODO: compare string interning compare and pure string compare
}
