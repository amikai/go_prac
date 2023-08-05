package assertex

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPanic(t *testing.T) {
	assert.Panics(t, func() { panic("world crash") })
}

func TestPanicWithError(t *testing.T) {
	err := errors.New("world crash")
	assert.PanicsWithError(t, err.Error(), func() { panic(err) })
}

func TestPanicWithValue(t *testing.T) {
	exp := "world crash"
	assert.PanicsWithValue(t, exp, func() { panic(exp) })
}
