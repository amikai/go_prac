package concex

import (
	"runtime"
	"testing"

	"github.com/sourcegraph/conc/stream"
	"github.com/stretchr/testify/assert"
)

func TestStream(t *testing.T) {
	out := make(chan string)

	// Stream is used to execute a stream of tasks concurrently while maintaining the order of the results.
	// Callbacks are called in the same order that tasks are submitted.
	s := stream.New().WithMaxGoroutines(runtime.NumCPU())
	go func() {
		for _, url := range searchURLs {
			url := url
			s.Go(func() stream.Callback {
				result, _ := fakeSearch(url)
				return func() { out <- result }
			})
		}
		// Wait will not return until all tasks and callbacks have been run.
		s.Wait()
		close(out)
	}()

	results := []string{}
	for o := range out {
		results = append(results, o)
	}
	assert.Equal(t, expSearchResult, results)

}
