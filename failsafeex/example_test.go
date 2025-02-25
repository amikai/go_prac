package failsafeex

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/failsafe-go/failsafe-go"
	"github.com/failsafe-go/failsafe-go/circuitbreaker"
	"github.com/failsafe-go/failsafe-go/failsafehttp"
	"github.com/failsafe-go/failsafe-go/timeout"
)

func Example_httpClientRoundTripper() {
	timeout := timeout.With[*http.Response](10 * time.Second)
	retry := failsafehttp.RetryPolicyBuilder().
		WithBackoff(time.Second, 30*time.Second).
		WithMaxRetries(3).
		Build()

	circuitBreaker := circuitbreaker.Builder[*http.Response]().
		HandleIf(func(response *http.Response, err error) bool {
			return response != nil && response.StatusCode == 429
		}).
		WithDelayFunc(failsafehttp.DelayFunc).
		Build()

	polices := []failsafe.Policy[*http.Response]{timeout, retry, circuitBreaker}
	rt := failsafehttp.NewRoundTripper(http.DefaultTransport, polices...)

	req, err := http.NewRequestWithContext(context.Background(), "GET", "http://example.com", nil)
	if err != nil {
		log.Fatal("fail to create request")
	}

	c := http.Client{
		Transport: rt,
	}
	rsp, err := c.Do(req)
	if err != nil {
		log.Fatal("fail to send request")
	}
	defer rsp.Body.Close()
	b, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatal("fail to read response body")
	}
	fmt.Println(b)
}
