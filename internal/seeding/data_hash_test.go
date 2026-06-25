package seeding

import (
	"reflect"
	"testing"
)

func TestCombineFields(t *testing.T) {
	tests := []struct {
		input   []any
		exp 	string
	}{
		{
			input:    	[]any{"sleep buster", 15, nil, false, nil, nil, 7.5, "phoenix down", 3, true, nil},
			exp: 	  	"sleep buster|15|NULL|false|NULL|NULL|7.5|phoenix down|3|true|NULL",
		},
		{
			input:    	[]any{},
			exp: 		"",
		},
		{
			input:    	[]any{nil, nil, nil, nil},
			exp: 		"NULL|NULL|NULL|NULL",
		},
	}

	for i, tc := range tests {
		got := combineFields(tc.input)
		
		if !reflect.DeepEqual(got, tc.exp) {
			t.Errorf("combineFields: Testcase %d: input: %v. expected: %v, got: %v", i+1, tc.input, tc.exp, got)
		}
	}
}
