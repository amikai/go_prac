package concex

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	set "github.com/deckarep/golang-set/v2"
)

const fakeResultPrefix = "Fake Result: "

func fakeSearch(url string) (string, error) {
	n := rand.Intn(100)
	time.Sleep(time.Duration(n) * time.Millisecond)
	return fakeResultPrefix + url, nil
}

func FakeSearchMustErr(url string) (string, error) {
	return "", searchErrs[url]
}

func FakeSearchErrCtxWithDuration(ctx context.Context, url string, t time.Duration) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-time.After(t):
		return "", searchErrs[url]
	}
}

func fakeSearchCtx(_ context.Context, url string) (string, error) {
	return fakeResultPrefix + url, nil
}

func FakeSearchCtxWithDuration(ctx context.Context, url string, t time.Duration) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-time.After(t):
		return fakeResultPrefix + url, nil
	}
}

var SearchURLs = []string{
	"https://www.google.com",
	"https://www.yahoo.com",
	"https://www.amazon.com",
	"https://www.apple.com",
	"https://www.facebook.com",
	"https://www.merriam-webster.com",
	"https://www.pinterest.com",
	"https://www.yelp.com",
	"https://www.microsoft.com",
	"https://www.walmart.com",
}

var searchErrs = map[string]error{
	"https://www.google.com":          errors.New("network failed when searching https://www.google.com"),
	"https://www.yahoo.com":           errors.New("network failed when searching https://www.yahoo.com"),
	"https://www.amazon.com":          errors.New("network failed when searching https://www.amazon.com"),
	"https://www.apple.com":           errors.New("network failed when searching https://www.apple.com"),
	"https://www.facebook.com":        errors.New("network failed when searching https://www.facebook.com"),
	"https://www.merriam-webster.com": errors.New("network failed when searching https://www.merriam-webster.com"),
	"https://www.pinterest.com":       errors.New("network failed when searching https://www.pinterest.com"),
	"https://www.yelp.com":            errors.New("network failed when searching https://www.yelp.com"),
	"https://www.microsoft.com":       errors.New("network failed when searching https://www.microsoft.com"),
	"https://www.walmart.com":         errors.New("network failed when searching https://www.walmart.com"),
}

// SearchErr return the correspond error for each url
// if len(errs) > 0, then wrap errs to new error
func SearchErr(url string, errs ...error) error {
	retErr := searchErrs[url]
	if len(errs) == 0 {
		return retErr
	}
	for _, err := range errs {
		retErr = fmt.Errorf("%w: %w", retErr, err)
	}
	return retErr
}

func RandomFailURLIndex() int {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	return random.Intn(len(SearchURLs))
}

var expSearchResult = []string{
	fakeResultPrefix + "https://www.google.com",
	fakeResultPrefix + "https://www.yahoo.com",
	fakeResultPrefix + "https://www.amazon.com",
	fakeResultPrefix + "https://www.apple.com",
	fakeResultPrefix + "https://www.facebook.com",
	fakeResultPrefix + "https://www.merriam-webster.com",
	fakeResultPrefix + "https://www.pinterest.com",
	fakeResultPrefix + "https://www.yelp.com",
	fakeResultPrefix + "https://www.microsoft.com",
	fakeResultPrefix + "https://www.walmart.com",
}

// the expSearchResultSet is thread safe
var expSearchResultSet = set.NewSet(expSearchResult...)
