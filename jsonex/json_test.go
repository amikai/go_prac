package jsonex

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Company struct {
	Employees []*Employee `json:"employees"`
	Name      string      `json:"name"`
}

type Employee struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func TestMarshal(t *testing.T) {
	data := []byte(`
{
	"employees": [{
			"id": "001",
			"name": "John"
		},
		{
			"id": "002",
			"name": "Mary"
		},
		{
			"id": "003",
			"name": "Bob"
		}
	],
	"name": "ABCCompany"
}`)

	expected := Company{
		Employees: []*Employee{
			{ID: "001", Name: "John"},
			{ID: "002", Name: "Mary"},
			{ID: "003", Name: "Bob"},
		},
		Name: "ABCCompany",
	}

	var got Company
	err := json.Unmarshal(data, &got)
	assert.NoError(t, err)
	assert.Equal(t, got, expected)
}
