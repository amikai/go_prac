package concex

import (
	"context"
	"math/rand"
	"runtime"
	"testing"
	"time"

	"github.com/sourcegraph/conc/pool"
	"github.com/stretchr/testify/assert"
)

func TestResultPool(t *testing.T) {
	var p *pool.ResultPool[string] = pool.NewWithResults[string]().WithMaxGoroutines(runtime.NumCPU())

	for _, url := range SearchURLs {
		p.Go(func() string {
			result, _ := fakeSearch(url)
			return result
		})
	}
	got := p.Wait()
	// The order of the results is not guaranteed to be the same as the order the tasks were submitted.
	assert.ElementsMatch(t, expSearchResult, got)
}

func TestResultErrPool(t *testing.T) {
	var p *pool.ResultErrorPool[string] = pool.NewWithResults[string]().WithMaxGoroutines(runtime.NumCPU()).WithErrors()
	failedIndex := rand.Intn(len(SearchURLs))
	failedIndex2 := rand.Intn(len(SearchURLs))
	for i, url := range SearchURLs {
		p.Go(func() (string, error) {
			if i == failedIndex || i == failedIndex2 {
				_, err := FakeSearchMustErr(url)
				return "", err
			}
			res, _ := fakeSearch(url)
			return res, nil
		})
	}
	got, err := p.Wait()

	var exp []string
	// Wait() return value ignore the failed results
	for i, url := range SearchURLs {
		if i == failedIndex || i == failedIndex2 {
			continue
		}
		s, _ := fakeSearch(url)
		exp = append(exp, s)
	}

	assert.ErrorIs(t, err, SearchErr(SearchURLs[failedIndex]))
	assert.ErrorIs(t, err, SearchErr(SearchURLs[failedIndex2]))
	// The order of the results is not guaranteed to be the same as the order the tasks were submitted.
	assert.ElementsMatch(t, exp, got)
}

func TestResultErrPoolCollected(t *testing.T) {
	var p *pool.ResultErrorPool[string] = pool.NewWithResults[string]().WithMaxGoroutines(runtime.NumCPU()).WithErrors().WithCollectErrored()
	failedIndex := rand.Intn(len(SearchURLs))
	failedIndex2 := rand.Intn(len(SearchURLs))
	for i, url := range SearchURLs {
		p.Go(func() (string, error) {
			if i == failedIndex || i == failedIndex2 {
				_, err := FakeSearchMustErr(url)
				return "", err
			}
			res, _ := fakeSearch(url)
			return res, nil
		})
	}
	got, err := p.Wait()

	var exp []string
	// Wait() return value still collect the failed results
	for i, url := range SearchURLs {
		var s string
		if i != failedIndex && i != failedIndex2 {
			s, _ = fakeSearch(url)
		}
		exp = append(exp, s)
	}

	assert.ErrorIs(t, err, SearchErr(SearchURLs[failedIndex]))
	assert.ErrorIs(t, err, SearchErr(SearchURLs[failedIndex2]))
	// The order of the results is not guaranteed to be the same as the order the tasks were submitted.
	assert.ElementsMatch(t, exp, got)
}

func TestResultContextPoolCancelOnError(t *testing.T) {
	var p *pool.ResultContextPool[string] = pool.NewWithResults[string]().WithMaxGoroutines(runtime.NumCPU()).WithContext(t.Context()).WithCancelOnError()
	failedIndex := rand.Intn(len(SearchURLs))
	for i, url := range SearchURLs {
		p.Go(func(ctx context.Context) (string, error) {
			if i == failedIndex {
				return FakeSearchMustErr(url)
			}
			return fakeSearchCtx(ctx, url)
		})
	}
	got, err := p.Wait()
	assert.Error(t, err)
	assert.ErrorIs(t, err, SearchErr(SearchURLs[failedIndex]))
	// WithCancelOnError() make pool cancel all task when one task raise error, so the result must be subset of expected result.
	// These results was finished before canceling.
	assert.Subset(t, expSearchResult, got)
}

func TestResultContextPoolCancelOnErrorTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(t.Context(), 100*time.Millisecond)
	defer cancel()
	failedIndex := rand.Intn(len(SearchURLs))
	var p *pool.ResultContextPool[string] = pool.NewWithResults[string]().WithMaxGoroutines(runtime.NumCPU()).WithContext(ctx)
	for i, url := range SearchURLs {
		i, url := i, url
		p.Go(func(ctx context.Context) (string, error) {
			if i == failedIndex {
				return FakeSearchMustErr(url)
			}
			return fakeSearchCtx(ctx, url)
		})
	}
	got, err := p.Wait()
	assert.Error(t, err)
	assert.ErrorIs(t, err, SearchErr(SearchURLs[failedIndex]))
	// WithCancelOnError() make pool cancel all task when one task raise error, so the result must be subset of expected result.
	// These results was finished before canceling.
	assert.Subset(t, expSearchResult, got)
}
