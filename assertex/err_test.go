package assertex

import (
	"errors"
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

func TestErrAs(t *testing.T) {
	// TODO:
}

func TestErrIs(t *testing.T) {
	// TODO:

}
