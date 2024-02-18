package concex

import (
	"runtime"
	"testing"

	"github.com/sourcegraph/conc/stream"
	"github.com/stretchr/testify/assert"
)

func TestStream(t *testing.T) {
	// Stream is used to execute a stream of tasks concurrently while maintaining the order of the results.
	// Callbacks are called in the same order that tasks are submitted.
	s := stream.New().WithMaxGoroutines(runtime.NumCPU())
	var got []string
	for _, url := range searchURLs {
		s.Go(func() stream.Callback {
			result, _ := fakeSearch(url)
			return func() { got = append(got, result) }
		})
	}
	// Wait will not return until all tasks and callbacks have been run.
	s.Wait()

	assert.Equal(t, expSearchResult, got)
}
