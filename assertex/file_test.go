package assertex

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirExists(t *testing.T) {
	// root folder must exist
	assert.DirExists(t, "/")
	assert.NoDirExists(t, "/123")
}

func TestFileExists(t *testing.T) {
	f, err := os.CreateTemp("", "dummy")
	assert.NoError(t, err)
	defer os.Remove(f.Name())
	defer f.Close()
	assert.FileExists(t, f.Name())
}
