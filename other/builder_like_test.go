package other

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputerBuilderLike(t *testing.T) {
	e := NewEmployee().WithID("ID-001").WithAge(18).WithName("amikai")
	assert.Equal(t, "ID-001", e.ID)
	assert.Equal(t, 18, e.Age)
	assert.Equal(t, "amikai", e.Name)
}
