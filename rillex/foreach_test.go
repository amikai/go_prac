package rillex

import (
	"context"
	"testing"
	"time"

	"github.com/destel/rill"
	"github.com/stretchr/testify/assert"

	"github.com/amikai/go_prac/concex"
)

const concLimit = 8

func TestForEach(t *testing.T) {
	ctx := context.Background()
	urls := rill.FromSlice(concex.SearchURLs, nil)
	err := rill.ForEach(urls, concLimit, func(url string) error {
		_, err := concex.FakeSearchCtxWithDuration(ctx, url, time.Second)
		return err
	})
	assert.NoError(t, err)
}

func TestForEachErr(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	failIndex := concex.RandomFailURLIndex()

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
	if err != nil {
		cancel()
	}
	assert.ErrorIs(t, err, concex.SearchErr(concex.SearchURLs[failIndex]))
}
