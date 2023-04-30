package assertex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyMap(t *testing.T) {
	var nilMap map[string]string
	assert.Empty(t, nilMap)

	emptyMap := map[string]string{}
	assert.Empty(t, emptyMap)
}

func TestEmptySlice(t *testing.T) {
	var nilSlice []string
	assert.Empty(t, nilSlice)

	emptySlice := []string{}
	assert.Empty(t, emptySlice)
}

func TestEmptyChan(t *testing.T) {
	var nilChan chan int
	assert.Empty(t, nilChan)

	emptyChan := make(chan int)
	assert.Empty(t, emptyChan)
}

func TestNilPointer(t *testing.T) {
	var nilP *int
	assert.Empty(t, nilP)
}
