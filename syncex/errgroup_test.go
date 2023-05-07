package syncex

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"
)

// There are good examples on error group document.
// https://pkg.go.dev/golang.org/x/sync/errgroup

const fakeResultPrefix = "Fake Result: "

func fakeSearch(url string) (string, error) {
	return fakeResultPrefix + url, nil
}

func TestErrGroup(t *testing.T) {
	urls := []string{
		"https://www.google.com",
		"https://www.yahoo.com",
		"https://www.amazon.com",
	}
	results := make([]string, len(urls))
	g, _ := errgroup.WithContext(context.Background())

	for i, url := range urls {
		i, url := i, url
		g.Go(func() error {
			result, err := fakeSearch(url)
			if err != nil {
				return err
			}
			results[i] = result
			return nil
		})
	}
	err := g.Wait()
	assert.NoError(t, err)

	exp := []string{
		fakeResultPrefix + "https://www.google.com",
		fakeResultPrefix + "https://www.yahoo.com",
		fakeResultPrefix + "https://www.amazon.com",
	}
	assert.Equal(t, exp, results)
}
