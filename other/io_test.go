package other

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestLimiter(t *testing.T) {
	r := bytes.NewReader([]byte("ABCDEFG"))
	lr := io.LimitReader(r, 4)

	got, err := io.ReadAll(lr)
	if err != nil {
		t.Fatal(err)
	}

	want := []byte("ABCD")
	if !bytes.Equal(want, got) {
		t.Errorf("got != want")
	}
}

func TestMultiReader(t *testing.T) {
	r1 := strings.NewReader("ABC")
	r2 := strings.NewReader("DEF")
	r3 := strings.NewReader("GHI")

	mr := io.MultiReader(r1, r2, r3)
	got, err := io.ReadAll(mr)
	if err != nil {
		t.Fatal(err)
	}

	want := []byte("ABCDEFGHI")
	if !bytes.Equal(want, got) {
		t.Errorf("got != want")
	}
}

func TestMultiWriter(t *testing.T) {
	w1 := &strings.Builder{}
	w2 := &strings.Builder{}
	w3 := &strings.Builder{}

	mw := io.MultiWriter(w1, w2, w3)
	_, err := io.WriteString(mw, "XYZ")
	if err != nil {
		t.Fatal(err)
	}

	want := "XYZ"
	if want != w1.String() {
		t.Errorf("want = %s, got = %s", want, w1.String())
	}
	if want != w2.String() {
		t.Errorf("want = %s, got = %s", want, w1.String())
	}
	if want != w3.String() {
		t.Errorf("want = %s, got = %s", want, w1.String())
	}
}

func TestTeeReader(t *testing.T) {
	r := strings.NewReader("ABC")
	w := &strings.Builder{}

	rt := io.TeeReader(r, w)
	wt := &strings.Builder{}
	_, err := io.Copy(wt, rt)
	if err != nil {
		t.Fatal(err)
	}

	want := "ABC"
	if want != w.String() {
		t.Errorf("want = %s, got = %s", want, w.String())
	}
	if want != wt.String() {
		t.Errorf("want = %s, got = %s", want, wt.String())
	}
}

func TestPipe(t *testing.T) {
	r, w := io.Pipe()
	go func() {
		_, _ = io.WriteString(w, "ABC")
		w.Close()
	}()
	got, err := io.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}
	want := []byte("ABC")
	if !bytes.Equal(want, got) {
		t.Errorf("got != want")
	}
}
