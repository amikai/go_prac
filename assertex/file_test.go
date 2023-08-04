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
	//  The directory t.TempDir() is automatically removed by Cleanup when the test and all its subtests complete.
	f, err := os.CreateTemp(t.TempDir(), "dummy")
	assert.NoError(t, err)
	defer f.Close()
	assert.FileExists(t, f.Name())
}
