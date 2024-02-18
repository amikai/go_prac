package syncex

import (
	"context"
	"runtime"
	"sync/atomic"
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

func TestErrGroupSetLimit(t *testing.T) {
	g, _ := errgroup.WithContext(context.Background())
	g.SetLimit(runtime.NumCPU())
	t.Logf("errgroup SetLimit(%d)", runtime.NumCPU())

	var total uint64

	for i := 0; i < 1000000; i++ {
		g.Go(func() error {
			atomic.AddUint64(&total, 1)
			return nil
		})
	}
	err := g.Wait()
	assert.NoError(t, err)
	assert.Equal(t, uint64(1000000), total)
}
