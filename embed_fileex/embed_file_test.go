package embedfileex

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed file.txt
var s string

//go:embed file.txt
var b []byte

func TestEmbedString(t *testing.T) {
	assert.Equal(t, "hello world\n", s)
}

func TestEmbedBytes(t *testing.T) {
	assert.Equal(t, []byte("hello world\n"), b)
}
