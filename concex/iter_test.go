package concex

import (
	"testing"

	"github.com/sourcegraph/conc/iter"
	"github.com/stretchr/testify/assert"
)

func TestIter(t *testing.T) {
	out := make(chan string)
	go func() {
		iter.ForEach(searchURLs, func(url *string) {
			result, err := fakeSearch(*url)
			if err != nil {
				return
			}
			out <- result
		})
		close(out)
	}()

	results := []string{}
	for o := range out {
		results = append(results, o)
	}
	assert.ElementsMatch(t, expSearchResult, results)
}
