package assertex

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpBodyContains(t *testing.T) {
	// simple case
	method := http.MethodGet
	var urlPath string
	var qValues url.Values
	assert.HTTPBodyContains(t,
		func(w http.ResponseWriter, _ *http.Request) {
			fmt.Fprintf(w, "hello, world")
		},
		method, urlPath, qValues, "hello")

	// test with method, url path and query string
	method = http.MethodGet
	urlPath = "/test/path"
	qValues, err := url.ParseQuery("testkey=testval")
	assert.NoError(t, err)
	assert.HTTPBodyContains(t,
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method == method && r.URL.Path == urlPath && r.URL.Query().Get("testkey") == "testval" {
				fmt.Fprintf(w, "hello, world")
			}
		},
		method, urlPath, qValues, "hello,")
}

func TestHttpStatusCode(t *testing.T) {
	// simple case
	method := http.MethodGet
	var urlPath string
	var qValues url.Values
	assert.HTTPStatusCode(t,
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotImplemented)
		},
		method, urlPath, qValues, http.StatusNotImplemented)

	// test with method, url path and query string
	method = http.MethodGet
	urlPath = "/test/path"
	qValues, err := url.ParseQuery("testkey=testval")
	assert.NoError(t, err)
	assert.HTTPStatusCode(t,
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method == method && r.URL.Path == urlPath && r.URL.Query().Get("testkey") == "testval" {
				w.WriteHeader(http.StatusNotImplemented)
			}
		},
		method, urlPath, qValues, http.StatusNotImplemented)
}
