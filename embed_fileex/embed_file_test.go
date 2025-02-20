package embedfileex

import (
	"embed"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:embed file.txt
var s string

//go:embed file.txt
var b []byte

//go:embed file.txt amikai.txt
var fs embed.FS

func TestEmbedString(t *testing.T) {
	assert.Equal(t, "hello world\n", s)
}

func TestEmbedBytes(t *testing.T) {
	assert.Equal(t, []byte("hello world\n"), b)
}

func TestEmbeddedFs(t *testing.T) {
	b, err := fs.ReadFile("file.txt")
	require.NoError(t, err)
	assert.Equal(t, []byte("hello world\n"), b)

	b, err = fs.ReadFile("amikai.txt")
	require.NoError(t, err)
	assert.Equal(t, []byte("Hi, Amikai\n"), b)
}
