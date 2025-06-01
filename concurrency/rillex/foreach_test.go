package rillex

import (
	"context"
	"testing"
	"time"

	"github.com/destel/rill"
	"github.com/stretchr/testify/assert"

	"github.com/amikai/go_prac/concurrency/concex"
)

const concLimit = 8

func TestForEach(t *testing.T) {
	ctx, cancel := context.WithCancel(t.Context())
	defer cancel()

	urls := rill.FromSlice(concex.SearchURLs, nil)
	err := rill.ForEach(urls, concLimit, func(url string) error {
		_, err := concex.FakeSearchCtxWithDuration(ctx, url, time.Second)
		return err
	})
	assert.NoError(t, err)
}

func TestForEachErr(t *testing.T) {
	ctx, cancel := context.WithCancel(t.Context())
	defer cancel()

	failIndex := concex.RandomFailURLIndex()

	// See https://github.com/destel/rill/issues/19#issuecomment-2151875524 for detail explanation
	urls := rill.FromSlice(concex.SearchURLs, nil)
	// terminate early on the first error
	// for each does not guarantee the order of the results
	err := rill.ForEach(urls, concLimit, func(url string) error {
		searchFunc := concex.FakeSearchCtxWithDuration
		if url == concex.SearchURLs[failIndex] {
			searchFunc = concex.FakeSearchErrCtxWithDuration
		}
		_, err := searchFunc(ctx, url, time.Second)
		return err
	})
	// enconter first error, then cancel the tasks
	assert.ErrorIs(t, err, concex.SearchErr(concex.SearchURLs[failIndex]))
}
