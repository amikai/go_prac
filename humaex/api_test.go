package main

import (
	"testing"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/humatest"
	"github.com/stretchr/testify/assert"
)

func TestGetGreeting(t *testing.T) {
	_, api := humatest.New(t)
	huma.Get(api, GreetingAPIPath, GreetHandler)

	resp := api.Get("/greeting/world")
	assert.JSONEq(t, `{"message":"Hello, world!"}`, resp.Body.String())
}
