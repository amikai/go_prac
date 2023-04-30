package httpex

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpHandler(t *testing.T) {
	contentTypeKey := "Conten-Type"
	contentTypeVal := "text/html; charset=utf-8"
	statusCode := http.StatusAccepted
	bodyContent := "hello world"

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(contentTypeKey, contentTypeVal)
		w.WriteHeader(statusCode)
		io.WriteString(w, bodyContent)
	}

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	assert.Equal(t, resp.StatusCode, statusCode)
	assert.Equal(t, resp.Header.Get(contentTypeKey), contentTypeVal)
	assert.Equal(t, body, []byte(bodyContent))
}

// TODO: httptest.Server
