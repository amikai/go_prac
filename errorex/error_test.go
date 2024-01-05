package errorex

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var errExampleA = errors.New("errExampleA")
var errExampleB = errors.New("errExampleA")
var errExampleC = errors.New("errExampleA")

func TestWrapMultiError(t *testing.T) {
	wrapsErr := fmt.Errorf("%w: %w: %w", errExampleA, errExampleB, errExampleC)

	assert.True(t, errors.Is(wrapsErr, errExampleA))
	assert.True(t, errors.Is(wrapsErr, errExampleB))
	assert.True(t, errors.Is(wrapsErr, errExampleC))
}

func TestJoin(t *testing.T) {
	joinErr := errors.Join(errExampleA, errExampleB, errExampleC)

	assert.True(t, errors.Is(joinErr, errExampleA))
	assert.True(t, errors.Is(joinErr, errExampleB))
	assert.True(t, errors.Is(joinErr, errExampleC))
}

// TODO: add error.As example
