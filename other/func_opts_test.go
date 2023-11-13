package other

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputerFuncOpts(t *testing.T) {
	c := NewComputer(WithMemory("2G"), WithCPUNum(4))
	assert.Equal(t, 4, c.CPUNum)
	assert.Equal(t, "2G", c.Memory)
}
