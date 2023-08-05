package assertex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYamlEq(t *testing.T) {
	// not in same order
	assert.YAMLEq(t, `
- name: amikai-chuang
  city: Taipei
  age: 18
`,
		`
- age: 18
  name: amikai-chuang
  city: Taipei
`)

}
