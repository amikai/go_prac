package main

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
	"github.com/go-chi/chi/v5"
)

// GreetingOutput represents the greeting operation response.

func main() {
	// Create a new router & API.
	router := chi.NewMux()
	api := humachi.New(router, huma.DefaultConfig("My API", "1.0.0"))
	huma.Get(api, GreetingAPIPath, GreetHandler)

	http.ListenAndServe("127.0.0.1:8888", router)
}
