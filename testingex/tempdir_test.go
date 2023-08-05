package testingex

import "testing"

func TestTempDir(t *testing.T) {
	dir := t.TempDir()
	t.Logf("tempdir: %s", dir)
}
