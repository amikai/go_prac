package builtin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type set[T comparable] = map[T]struct{}

func TestSet(t *testing.T) {
	s := set[string]{
		"e1": struct{}{},
		"e2": struct{}{},
	}

	exp := map[string]struct{}{
		"e1": {},
		"e2": {},
	}

	assert.Equal(t, exp, s)
}
