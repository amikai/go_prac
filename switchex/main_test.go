package switchex

import (
	"testing"
)

func TestTypeSelectSwitch(t *testing.T) {
	type testcase struct {
		input interface{}
		exp   string
	}

	testcases := []testcase{
		{
			input: true,
			exp:   "bool",
		},
		{
			input: 1,
			exp:   "int",
		},
		{
			input: 1.0,
			exp:   "other",
		},
	}

	for _, tt := range testcases {
		if got := typeSwtich(tt.input); got != tt.exp {
			t.Errorf(`typeSelectSwitch(%v) got %s, expect %s`, tt.input, got, tt.exp)
		}
	}
}
