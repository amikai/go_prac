package assertex

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErr(t *testing.T) {
	// test empty err
	var err error
	assert.NoError(t, err)

	// assert has err
	err = errors.New("example err")
	assert.Error(t, err)
}

func TestErrIs(t *testing.T) {
	var cause = errors.New("cause")
	var wrap = fmt.Errorf("wrap: %w", cause)
	assert.ErrorIs(t, wrap, cause)
}

type exampleErr struct {
	msg string
}

func (e exampleErr) Error() string {
	return e.msg
}

func TestErrAs(t *testing.T) {
	err := exampleErr{"just an example"}
	var exErr exampleErr
	assert.ErrorAs(t, err, &exErr)
}
