package fsex

import (
	"io/fs"
	"testing"
	"testing/fstest"

	"github.com/stretchr/testify/assert"
)

var mfs = fstest.MapFS{
	"file1": &fstest.MapFile{
		Data: []byte("file1 content"),
	},
	"file2": &fstest.MapFile{
		Data: []byte("file2 content"),
	},
}

func TestMapfs(t *testing.T) {
	b, err := fs.ReadFile(mfs, "file1")
	assert.NoError(t, err)
	assert.Equal(t, b, []byte("file1 content"))

	b, err = fs.ReadFile(mfs, "file2")
	assert.NoError(t, err)
	assert.Equal(t, b, []byte("file2 content"))
}
