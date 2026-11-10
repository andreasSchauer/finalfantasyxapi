package helpers

import (
	"reflect"
	"testing"
)


func TestNameToString(t *testing.T) {
	tests := []struct {
		name	string
		version	*int32
		spec	*string
		exp		string
	}{
		{
			name:		"fred",
			version: 	nil,
			spec: 		nil,
			exp:		"fred",
		},
		{
			name:		"harold",
			version: 	GetInt32Ptr(1),
			spec: 		GetStrPtr("the one and only"),
			exp:		"harold - 1 (the one and only)",
		},
		{
			name:		"hank",
			version: 	GetInt32Ptr(2),
			spec: 		nil,
			exp:		"hank - 2",
		},
	}

	for i, tc := range tests {
		got := NameToString(tc.name, tc.version, tc.spec)

		if !reflect.DeepEqual(got, tc.exp) {
			t.Errorf("NameToString: Testcase %d: name: %s, version %s, spec %v. expected: %v, got: %v", i+1, tc.name, FormatInt32Ptr(tc.version), DerefStringPtr(tc.spec), tc.exp, got)
		}
	}
}

func TestNameAmountToString(t *testing.T) {
	tests := []struct {
		name	string
		version	*int32
		amt		int32
		exp		string
	}{
		{
			name:		"fred",
			version: 	nil,
			amt: 		2,
			exp:		"fred x2",
		},
		{
			name:		"harold",
			version: 	GetInt32Ptr(1),
			amt: 		1,
			exp:		"harold - 1",
		},
	}

	for i, tc := range tests {
		got := NameAmountString(tc.name, tc.version, nil, tc.amt)

		if !reflect.DeepEqual(got, tc.exp) {
			t.Errorf("NameToString: Testcase %d: name: %s, version %s, amt %d. expected: %v, got: %v", i+1, tc.name, FormatInt32Ptr(tc.version), tc.amt, tc.exp, got)
		}
	}
}


func TestGetMapKeyStr(t *testing.T) {
	tests := []struct {
		input	map[string]int
		exp		string
	}{
		{
			input: map[string]int{
				"f": 0,
				"a": 1,
				"t": 2,
				"d": 3,
			},
			exp: "'a', 'd', 'f', 't'",
		},
		{
			input: map[string]int{
				"f": 0,
			},
			exp: "'f'",
		},
		{
			input: map[string]int{},
			exp: "",
		},
		{
			input: nil,
			exp: "",
		},
	}

	for i, tc := range tests {
		got := GetMapKeyStr(tc.input)

		if !reflect.DeepEqual(got, tc.exp) {
			t.Errorf("NameToString: Testcase %d: input: %v. expected: %v, got: %v", i+1, tc.input, tc.exp, got)
		}
	}
}


func TestFormatStringSlice(t *testing.T) {
	tests := []struct {
		input	[]string
		exp		string
	}{
		{
			input: []string{"a", "b", "c", "d"},
			exp: "'a', 'b', 'c', 'd'",
		},
		{
			input: []string{"a"},
			exp: "'a'",
		},
		{
			input: []string{},
			exp: "",
		},
		{
			input: nil,
			exp: "",
		},
	}

	for i, tc := range tests {
		got := FormatStringSlice(tc.input)

		if !reflect.DeepEqual(got, tc.exp) {
			t.Errorf("NameToString: Testcase %d: input: %v. expected: %v, got: %v", i+1, tc.input, tc.exp, got)
		}
	}
}


func TestFormatIntSlice(t *testing.T) {
	tests := []struct {
		input	[]int32
		exp		string
	}{
		{
			input: []int32{1, 2, 3, 4},
			exp: "'1', '2', '3', '4'",
		},
		{
			input: []int32{1},
			exp: "'1'",
		},
		{
			input: []int32{},
			exp: "",
		},
		{
			input: nil,
			exp: "",
		},
	}

	for i, tc := range tests {
		got := FormatIntSlice(tc.input)

		if !reflect.DeepEqual(got, tc.exp) {
			t.Errorf("NameToString: Testcase %d: input: %v. expected: %v, got: %v", i+1, tc.input, tc.exp, got)
		}
	}
}