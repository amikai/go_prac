package testingex

import "testing"

func TestCleanup(t *testing.T) {
	t.Cleanup(func() { t.Log("third") })
	defer t.Log("second")
	t.Log("first")
}
