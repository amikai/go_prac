package syncex

import (
	"context"
	"fmt"
	"net/url"
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

const (
	fakeResultPrefix = "Fake Result: "
)

func fakeSearch(ctx context.Context, rawURL string) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("parse url %s: %w", u, err)
	}
	d := u.Query().Get("delay")
	delay, err := time.ParseDuration(d)
	if err != nil {
		return "", fmt.Errorf("parse duration %s: %w", delay, err)
	}

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-time.After(delay):
		return fakeResultPrefix + rawURL, nil
	}
}

// ConcSearch searches the web for results by URLs and allocates the results in
// the same order. ConcSearch does its best effort to search the web for
// results; it returns partial search results even if an error occurs.
func ConcSearch(ctx context.Context, urls []string, concLimit int) ([]string, error) {
	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(concLimit)

	results := make([]string, len(urls))
	for i, url := range urls {
		g.Go(func() error {
			result, err := fakeSearch(ctx, url)
			if err != nil {
				return err
			}
			results[i] = result
			return nil
		})
	}
	err := g.Wait()
	if err != nil {
		return results, err
	}

	return results, nil
}

func TestConcSearchLongRunWebsite(t *testing.T) {
	synctest.Run(func() {
		urls := []string{
			"https://www.example1.com?delay=1s",
			"https://www.example2.com?delay=1s",
			"https://www.example3.com?delay=1s",
			"https://www.example4.com?delay=2s",
			"https://www.example5.com?delay=1s",
			"https://www.example6.com?delay=1s",
			"https://www.example7.com?delay=1s",
			"https://www.example8.com?delay=1s",
			"https://www.example9.com?delay=1s",
		}
		want := []string{
			fakeResultPrefix + "https://www.example1.com?delay=1s",
			fakeResultPrefix + "https://www.example2.com?delay=1s",
			fakeResultPrefix + "https://www.example3.com?delay=1s",
			// context is done, https://www.example4.com?delay=2s takes too long to get result
			"",
			fakeResultPrefix + "https://www.example5.com?delay=1s",
			fakeResultPrefix + "https://www.example6.com?delay=1s",
			fakeResultPrefix + "https://www.example7.com?delay=1s",
			fakeResultPrefix + "https://www.example8.com?delay=1s",
			// context is done, we don't have enough time to search https://www.example9.com?delay=1s
			"",
		}

		concLimit := 8
		start := time.Now()
		ctx, cancel := context.WithTimeout(t.Context(), 1*time.Second)
		defer cancel()
		got, err := ConcSearch(ctx, urls, concLimit)

		assert.Equal(t, want, got)
		assert.ErrorIs(t, err, context.DeadlineExceeded)
		assert.Equal(t, time.Duration(len(urls)/concLimit)*time.Second, time.Since(start))
	})
}

func TestConcSearch(t *testing.T) {
	synctest.Run(func() {
		urls := []string{
			"https://www.example1.com?delay=1s",
			"https://www.example2.com?delay=1s",
			"https://www.example3.com?delay=1s",
			"https://www.example4.com?delay=1s",
			"https://www.example5.com?delay=1s",
			"https://www.example6.com?delay=1s",
			"https://www.example7.com?delay=1s",
			"https://www.example8.com?delay=1s",
		}
		want := []string{
			fakeResultPrefix + "https://www.example1.com?delay=1s",
			fakeResultPrefix + "https://www.example2.com?delay=1s",
			fakeResultPrefix + "https://www.example3.com?delay=1s",
			fakeResultPrefix + "https://www.example4.com?delay=1s",
			fakeResultPrefix + "https://www.example5.com?delay=1s",
			fakeResultPrefix + "https://www.example6.com?delay=1s",
			fakeResultPrefix + "https://www.example7.com?delay=1s",
			fakeResultPrefix + "https://www.example8.com?delay=1s",
		}

		concLimit := 2
		start := time.Now()
		got, err := ConcSearch(t.Context(), urls, concLimit)

		require.NoError(t, err)
		assert.Equal(t, want, got)
		assert.Equal(t, time.Duration(len(urls)/concLimit)*time.Second, time.Since(start))
	})
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
