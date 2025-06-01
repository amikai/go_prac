package testingex

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTEnv(t *testing.T) {
	t.Setenv("foo", "bar")
	assert.Equal(t, "bar", os.Getenv("foo"))
}
