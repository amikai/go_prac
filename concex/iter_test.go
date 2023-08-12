package concex

import (
	"runtime"
	"testing"

	"github.com/sourcegraph/conc/iter"
	"github.com/stretchr/testify/assert"
)

func TestIter(t *testing.T) {
	iter.ForEach(searchURLs, func(url *string) {
		_, _ = fakeSearch(*url)
	})
}

func TestIterator(t *testing.T) {
	iterator := iter.Iterator[string]{MaxGoroutines: runtime.NumCPU()}
	iterator.ForEach(searchURLs, func(url *string) {
		_, _ = fakeSearch(*url)
	})
}

func TestMap(t *testing.T) {
	results := iter.Map(searchURLs, func(url *string) string {
		result, err := fakeSearch(*url)
		if err != nil {
			return ""
		}
		return result
	})
	assert.ElementsMatch(t, expSearchResult, results)
}

func TestMapErr(t *testing.T) {
	_, retErr := iter.MapErr(searchURLs, func(url *string) (string, error) {
		result, err := fakeSearchMustErr(*url)
		if err != nil {
			return "", err
		}
		return result, nil
	})

	for _, url := range searchURLs {
		assert.ErrorIs(t, retErr, searchErr(url))
	}
}

func TestMapperMap(t *testing.T) {
	mapper := iter.Mapper[string, string]{MaxGoroutines: runtime.NumCPU()}
	results := mapper.Map(searchURLs, func(url *string) string {
		result, err := fakeSearch(*url)
		if err != nil {
			return ""
		}
		return result
	})
	assert.ElementsMatch(t, expSearchResult, results)
}

func TestMapperMapErr(t *testing.T) {
	mapper := iter.Mapper[string, string]{MaxGoroutines: runtime.NumCPU()}
	_, retErr := mapper.MapErr(searchURLs, func(url *string) (string, error) {
		result, err := fakeSearchMustErr(*url)
		if err != nil {
			return "", err
		}
		return result, nil
	})

	for _, url := range searchURLs {
		assert.ErrorIs(t, retErr, searchErr(url))
	}
}
