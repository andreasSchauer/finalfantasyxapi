package seeding

import (
	"reflect"
	"testing"
)

func TestCombineFields(t *testing.T) {
	tests := []struct {
		input		[]any
		expected    string
	}{
		{
			input: 		[]any{"sleep buster", 15, nil, false, nil, nil, 7.5, "phoenix down", 3, true, nil},
			expected: "sleep buster|15|NULL|false|NULL|NULL|7.5|phoenix down|3|true|NULL",
		},
		{
			input:		[]any{},
			expected:	"",
		},
		{
			input:		[]any{nil, nil, nil, nil},
			expected:	"NULL|NULL|NULL|NULL",
		},
	}

	for i, tc := range tests {
		actual := combineFields(tc.input)
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("combineFields: Testcase %d: input: %v. expected: %v, actual: %v", i+1, tc.input, tc.expected, actual)
		}
	}
}