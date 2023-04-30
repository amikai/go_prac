package assertex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonEq(t *testing.T) {
	// not in same order
	assert.JSONEq(t,
		`{"hello": "world", "foo": "bar"}`,
		`{"foo": "bar", "hello": "world"}`)

	// ugly format
	assert.JSONEq(t,
		`{"hello": "world", "foo": "bar"}`,
		`{
"hello": "world", 
"foo":             "bar"    
	}`)
}
