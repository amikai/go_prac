package rillex

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"iter"
	"slices"
	"strings"
	"testing"

	"github.com/destel/rill"
	"github.com/stretchr/testify/assert"
)

type Person struct {
	Name   string `json:"name,omitempty"`
	Age    int    `json:"age,omitempty"`
	Status string `json:"status,omitempty"`
}

// The demographic data is formatted as JSON lines, where each line represents
// an individual's data. The CensusDataStream reads each line one at a time, immediately
// processes the information by incrementing each person's age by one and
// changes their status to "done". After processing, each line is then
// immediately output to w. This means that there is no waiting for all lines to
// be processed before outputting; each line is handled individually and
// sequentially.
//
// For example, given the input:
// {"name":"Mary", "Age": 18}
// {"name":"John", "Age": 20}
//
// The output will be generated line by line, as follows:
// {"name":"Mary", "Age": 19, "status":"done"}
// {"name":"John", "Age": 21, "status":"done"}
func CensusDataStream(ctx context.Context, r io.Reader, w io.Writer) {
	lines := rill.FromSeq2(lines(ctx, r))
	people := rill.OrderedMap(lines, 4, func(b []byte) (*Person, error) {
		var p Person
		err := json.Unmarshal(b, &p)
		if err != nil {
			return nil, fmt.Errorf("unmarshall byte: %w", err)
		}
		p.Age++
		p.Status = "done"
		return &p, nil
	})

	dec := json.NewEncoder(w)
	for p, err := range rill.ToSeq2(people) {
		if err != nil {
			// WARN: Handle error properly in production code.
			continue
		}

		err = dec.Encode(&p)
		if err != nil {
			// WARN: Handle error properly in production code.
			continue
		}
	}
}

func TestCensusDataStream(t *testing.T) {
	r := strings.NewReader(`{"name":"Mary", "age": 18}
{"name":"John", "age": 20}`)
	w := &strings.Builder{}
	CensusDataStream(t.Context(), r, w)
	assert.Equal(t, `{"name":"Mary","age":19,"status":"done"}
{"name":"John","age":21,"status":"done"}
`, w.String())
}

// Lines returns a sequence of lines read from the input reader. Each line has
// the trailing newline removed and is allocated in a new memory block for safe
// use in multiple goroutines.
func lines(ctx context.Context, input io.Reader) iter.Seq2[[]byte, error] {
	scanner := bufio.NewScanner(input)
	scanner.Buffer(make([]byte, 8*1024), 16*1024*1024)
	return func(yield func([]byte, error) bool) {
		for scanner.Scan() {
			if ctx.Err() != nil {
				// The returned boolean from yield is ignored since we will
				// return immediately.
				yield(nil, ctx.Err())
				return
			}
			// The scanner's buffer is reused, so the line must be copied for
			// safe concurrent access. Note that the bytes returned by the
			// scanner do not include the trailing newline.
			if !yield(slices.Clone(scanner.Bytes()), nil) {
				return
			}
		}
		if err := scanner.Err(); err != nil {
			// The returned boolean from yield is ignored since we will return
			// immediately.
			yield(nil, err)
		}
	}
}
