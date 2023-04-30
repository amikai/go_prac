package httpex

import (
	"bytes"
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

func TestHttpHandlerEndToEnd(t *testing.T) {
	var uppercaseHandler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		assert.NoError(t, err)

		_, err = w.Write(bytes.ToUpper(body))
		assert.NoError(t, err)
	}
	ts := httptest.NewServer(uppercaseHandler)
	defer ts.Close()
	client := ts.Client()

	resp, err := client.Post(ts.URL+"/upper", "application/text", bytes.NewReader([]byte("abc123XYZ")))
	assert.NoError(t, err)

	respBody, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)

	assert.Equal(t, "ABC123XYZ", string(respBody))
}
