package concex

import (
	"testing"

	"github.com/sourcegraph/conc"
	"github.com/stretchr/testify/assert"
)

func TestConcWg(t *testing.T) {
	results := make([]string, len(searchURLs))
	wg := conc.NewWaitGroup()

	for i, url := range searchURLs {
		i, url := i, url
		wg.Go(func() {
			result, err := fakeSearch(url)
			if err != nil {
				return
			}
			results[i] = result
		})
	}
	wg.Wait()
	assert.Equal(t, expSearchResult, results)
}
