package loop

import (
	"fmt"
	"testing"
)

func TestUint64LoopSum(t *testing.T) {
	type testcase struct {
		input uint64
		exp   uint64
	}

	testcases := []testcase{
		{
			input: 100,
			exp:   5050,
		},
		{
			input: 3,
			exp:   6,
		},
		{
			input: 0,
			exp:   0,
		},
	}

	for i, tt := range testcases {
		t.Run(fmt.Sprintf("cStyleLoopSum testcase %d", i), func(t *testing.T) {
			if got := cStyleLoopSum(tt.input); got != tt.exp {
				t.Errorf(`cStyleLoopSum(%d) got %d, expect %d`, tt.input, got, tt.exp)
			}
		})

		t.Run(fmt.Sprintf("whileLoopSum testcase %d", i), func(t *testing.T) {
			if got := whileLoopSum(tt.input); got != tt.exp {
				t.Errorf(`whileLoopSum(%d) got %d, expect %d`, tt.input, got, tt.exp)
			}
		})

		t.Run(fmt.Sprintf("infiniteLoopSum testcase %d", i), func(t *testing.T) {
			if got := infiniteLoopSum(tt.input); got != tt.exp {
				t.Errorf(`infiniteLoopSum(%d) got %d, expect %d`, tt.input, got, tt.exp)
			}
		})

	}
}

func TestUint32LoopSum(t *testing.T) {
	type testcase struct {
		input uint32
		exp   uint32
	}

	testcases := []testcase{
		{
			input: 100,
			exp:   5050,
		},
		{
			input: 3,
			exp:   6,
		},
		{
			input: 0,
			exp:   0,
		},
	}

	for i, tt := range testcases {
		t.Run(fmt.Sprintf("cStyleLoopSum testcase %d", i), func(t *testing.T) {
			if got := cStyleLoopSum(tt.input); got != tt.exp {
				t.Errorf(`cStyleLoopSum(%d) got %d, expect %d`, tt.input, got, tt.exp)
			}
		})

		t.Run(fmt.Sprintf("whileLoopSum testcase %d", i), func(t *testing.T) {
			if got := whileLoopSum(tt.input); got != tt.exp {
				t.Errorf(`whileLoopSum(%d) got %d, expect %d`, tt.input, got, tt.exp)
			}
		})

		t.Run(fmt.Sprintf("infiniteLoopSum testcase %d", i), func(t *testing.T) {
			if got := infiniteLoopSum(tt.input); got != tt.exp {
				t.Errorf(`infiniteLoopSum(%d) got %d, expect %d`, tt.input, got, tt.exp)
			}
		})

	}
}

func TestUint32ForEachArraySum(t *testing.T) {
	type testcase struct {
		input []uint32
		exp   uint32
	}

	testcases := []testcase{
		{
			input: []uint32{9, 8, 6, 13},
			exp:   36,
		},
		{
			input: []uint32{100, 101, 102},
			exp:   303,
		},
	}

	for i, tt := range testcases {
		t.Run(fmt.Sprintf("forEachArraySum testcase %d", i), func(t *testing.T) {
			if got := forEachArraySum(tt.input); got != tt.exp {
				t.Errorf(`forEachArraySum(%d) got %d, expect %d`, tt.input, got, tt.exp)
			}
		})
	}
}

func TestUint64ForEachArraySum(t *testing.T) {
	type testcase struct {
		input []uint64
		exp   uint64
	}

	testcases := []testcase{
		{
			input: []uint64{9, 8, 6, 13},
			exp:   36,
		},
		{
			input: []uint64{100, 101, 102},
			exp:   303,
		},
	}

	for i, tt := range testcases {
		t.Run(fmt.Sprintf("forEachArraySum testcase %d", i), func(t *testing.T) {
			if got := forEachArraySum(tt.input); got != tt.exp {
				t.Errorf(`forEachArraySum(%d) got %d, expect %d`, tt.input, got, tt.exp)
			}
		})
	}
}
