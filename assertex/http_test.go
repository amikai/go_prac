package assertex

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpHandlerRspBodyContains(t *testing.T) {
	// simple case
	t.Run("simple case", func(t *testing.T) {
		method := http.MethodGet
		var urlPath string
		var qValues url.Values
		assert.HTTPBodyContains(t,
			func(w http.ResponseWriter, _ *http.Request) {
				fmt.Fprintf(w, "hello, world")
			},
			method, urlPath, qValues, "hello")
	})

	t.Run("complex case", func(t *testing.T) {
		method := http.MethodGet
		urlPath := "/test/path"
		qValues, err := url.ParseQuery("testkey=testval")
		assert.NoError(t, err)
		assert.HTTPBodyContains(t,
			func(w http.ResponseWriter, r *http.Request) {
				if r.Method == method && r.URL.Path == urlPath && r.URL.Query().Get("testkey") == "testval" {
					fmt.Fprintf(w, "hello, world")
				}
			},
			method, urlPath, qValues, "hello,")

	})
}

func TestHttpHandlerRspStatusCode(t *testing.T) {
	// simple case
	t.Run("simple case", func(t *testing.T) {
		method := http.MethodGet
		var urlPath string
		var qValues url.Values
		assert.HTTPStatusCode(t,
			func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(http.StatusTeapot)
			},
			method, urlPath, qValues, http.StatusTeapot)
	})

	t.Run("complex case", func(t *testing.T) {
		// test with method, url path and query string
		method := http.MethodGet
		urlPath := "/test/path"
		qValues, err := url.ParseQuery("testkey=testval")
		assert.NoError(t, err)
		assert.HTTPStatusCode(t,
			func(w http.ResponseWriter, r *http.Request) {
				if r.Method == method && r.URL.Path == urlPath && r.URL.Query().Get("testkey") == "testval" {
					w.WriteHeader(http.StatusTeapot)
				}
			},
			method, urlPath, qValues, http.StatusTeapot)
	})
}

func TestHttpHandlerRspSuccess(t *testing.T) {
	// simple case
	t.Run("simple case", func(t *testing.T) {
		method := http.MethodGet
		var urlPath string
		var qValues url.Values
		assert.HTTPSuccess(t,
			func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(http.StatusOK)
			},
			method, urlPath, qValues)
	})

	t.Run("complex case", func(t *testing.T) {
		// test with method, url path and query string
		method := http.MethodGet
		urlPath := "/test/path"
		qValues, err := url.ParseQuery("testkey=testval")
		assert.NoError(t, err)
		assert.HTTPSuccess(t,
			func(w http.ResponseWriter, r *http.Request) {
				if r.Method == method && r.URL.Path == urlPath && r.URL.Query().Get("testkey") == "testval" {
					w.WriteHeader(http.StatusOK)
				}
			},
			method, urlPath, qValues)
	})
}

func TestHttpHandlerRedirect(t *testing.T) {
	// 308 will case assert failure
	for rspStatus := 300; rspStatus <= 307; rspStatus++ {
		t.Run(fmt.Sprintf("3xx redirect status code: %d", rspStatus), func(t *testing.T) {
			method := http.MethodGet
			var urlPath string
			var qValues url.Values
			assert.HTTPRedirect(t,
				func(w http.ResponseWriter, _ *http.Request) {
					w.WriteHeader(rspStatus)
				},
				method, urlPath, qValues)
		})
	}
}

func TestHttpHandlerError(t *testing.T) {
	// from 400 to 511
	for rspStatus := http.StatusBadRequest; rspStatus <= http.StatusNetworkAuthenticationRequired; rspStatus++ {
		t.Run(fmt.Sprintf("http error: %d", rspStatus), func(t *testing.T) {
			method := http.MethodGet
			var urlPath string
			var qValues url.Values
			assert.HTTPError(t,
				func(w http.ResponseWriter, _ *http.Request) {
					w.WriteHeader(rspStatus)
				},
				method, urlPath, qValues)
		})
	}
}

func TestHttpBodyUtils(t *testing.T) {
	method := http.MethodGet
	var urlPath string
	var qValues url.Values

	want := "hello"
	got := assert.HTTPBody(
		func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("hello"))
		},
		method, urlPath, qValues)
	assert.Equal(t, want, got)
}
