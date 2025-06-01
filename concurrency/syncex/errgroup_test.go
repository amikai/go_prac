package syncex

import (
	"sync/atomic"
	"testing"
	"testing/synctest"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	g, _ := errgroup.WithContext(t.Context())

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
	require.NoError(t, err)

	exp := []string{
		fakeResultPrefix + "https://www.google.com",
		fakeResultPrefix + "https://www.yahoo.com",
		fakeResultPrefix + "https://www.amazon.com",
	}
	assert.Equal(t, exp, results)
}

func TestConcAddOnePerSec(t *testing.T) {
	synctest.Run(func() {
		n := 100000
		concLimit := 8

		start := time.Now()
		total := ConcAddOnePerSec(n, concLimit)

		assert.Equal(t, n, total)
		assert.Equal(t, time.Duration(n/concLimit)*time.Second, time.Since(start))
	})
}

func ConcAddOnePerSec(num, concLimit int) int {
	g := &errgroup.Group{}
	g.SetLimit(concLimit)
	var total uint64
	for range num {
		g.Go(func() error {
			time.Sleep(1 * time.Second)
			atomic.AddUint64(&total, 1)
			return nil
		})
	}
	_ = g.Wait()
	return int(total)
}
