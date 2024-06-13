package rw

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPipeline(t *testing.T) {
	pr, pw := io.Pipe()

	concurrency := 10
	input := []byte{0, 1, 2, 3, 4, 5, 6, 7, 9, 10, 11, 12, 13, 14, 15}
	for i := 0; i < concurrency; i++ {
		go func() {
			_, err := pw.Write(input)
			assert.NoError(t, err)
		}()
	}

	got := make([]byte, len(input))
	exp := []byte{0, 1, 2, 3, 4, 5, 6, 7, 9, 10, 11, 12, 13, 14, 15}

	for i := 0; i < concurrency; i++ {
		_, err := pr.Read(got)
		assert.NoError(t, err)
		assert.Equal(t, exp, got)
	}
}
