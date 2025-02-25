package main

import (
	"context"
	"fmt"
)

// TODO: move api here
const GreetingAPIPath = "/greeting/{name}"

type GreetingOutput struct {
	Body struct {
		Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
	}
}

type GreetingInput struct {
	Name string `path:"name" maxLength:"30" example:"world" doc:"Name to greet"`
}

func GreetHandler(ctx context.Context, input *GreetingInput,
) (*GreetingOutput, error) {
	resp := &GreetingOutput{}
	resp.Body.Message = fmt.Sprintf("Hello, %s!", input.Name)
	return resp, nil
}
