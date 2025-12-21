package main

import (
	"reflect"
	"testing"
)

// checks if two values of a field are equal
func compare(t *testing.T, testName, fieldName string, exp, got any, dontCheck map[string]bool) {
	t.Helper()

	if dontCheck != nil && dontCheck[fieldName] {
		return
	}

	switch e := exp.(type) {

	case int:
		g, ok := got.(int)
		if !ok {
			t.Fatalf("%s: %s type mismatch: expected int32, got %T", testName, fieldName, got)
		}
		compInt(t, testName, fieldName, e, g)

	case int32:
		g, ok := got.(int32)
		if !ok {
			t.Fatalf("%s: %s type mismatch: expected int32, got %T", testName, fieldName, got)
		}
		compInt32(t, testName, fieldName, e, g)

	case float32:
		g, ok := got.(float32)
		if !ok {
			t.Fatalf("%s: %s type mismatch: expected float32, got %T", testName, fieldName, got)
		}
		compFloat32(t, testName, fieldName, e, g)

	case string:
		g, ok := got.(string)
		if !ok {
			t.Fatalf("%s: %s type mismatch: expected string, got %T", testName, fieldName, got)
		}
		compString(t, testName, fieldName, e, g)

	case *int32:
		g, ok := got.(*int32)
		if !ok {
			t.Fatalf("%s: %s type mismatch: expected *int32, got %T", testName, fieldName, got)
		}
		compInt32Ptr(t, testName, fieldName, e, g)

	case *float32:
		g, ok := got.(*float32)
		if !ok {
			t.Fatalf("%s: %s type mismatch: expected *float32, got %T", testName, fieldName, got)
		}
		compFloat32Ptr(t, testName, fieldName, e, g)

	case *string:
		g, ok := got.(*string)
		if !ok {
			t.Fatalf("%s: %s type mismatch: expected *string, got %T", testName, fieldName, got)
		}
		compStringPtr(t, testName, fieldName, e, g)

	default:
		t.Fatalf("%s: unsupported type for %s: %T", testName, fieldName, exp)
	}
}

func compInt(t *testing.T, testName, fieldName string, exp, got int) {
	if exp != got {
		t.Fatalf("%s: expected %s %d, got %d", testName, fieldName, exp, got)
	}
}

func compInt32(t *testing.T, testName, fieldName string, exp, got int32) {
	if exp != got {
		t.Fatalf("%s: expected %s %d, got %d", testName, fieldName, exp, got)
	}
}

func compFloat32(t *testing.T, testName, fieldName string, exp, got float32) {
	if exp != got {
		t.Fatalf("%s: expected %s %.2f, got %.2f", testName, fieldName, exp, got)
	}
}

func compString(t *testing.T, testName, fieldName, exp, got string) {
	if exp != "" && exp != got {
		t.Fatalf("%s: expected %s %s, got %s", testName, fieldName, exp, got)
	}
}

func compInt32Ptr(t *testing.T, testName, fieldName string, exp, got *int32) {
	compPtr(t, testName, fieldName, exp, got, compInt32)
}

func compFloat32Ptr(t *testing.T, testName, fieldName string, exp, got *float32) {
	compPtr(t, testName, fieldName, exp, got, compFloat32)
}

func compStringPtr(t *testing.T, testName, fieldName string, exp, got *string) {
	compPtr(t, testName, fieldName, exp, got, compString)
}

func bothPtrsPresent[T, U any](t *testing.T, testName, fieldName string, exp *T, got *U) bool {
	switch {
	case exp == nil && got == nil:
		return false
	case exp == nil && got != nil:
		t.Fatalf("%s: expected %s nil, got %v", testName, fieldName, *got)
		return false
	case exp != nil && got == nil:
		t.Fatalf("%s: expected %s %v, got nil", testName, fieldName, *exp)
		return false
	default:
		return true
	}
}

func compPtr[T any](t *testing.T, testName, fieldName string, exp, got *T, compFunc func(*testing.T, string, string, T, T)) {
	if !bothPtrsPresent(t, testName, fieldName, exp, got) {
		return
	}
	compFunc(t, testName, fieldName, *exp, *got)
}

func testStructSlices[T any](t *testing.T, testName, fieldName string, exp, got []T) {
	switch {
	case exp == nil && got == nil:
		return
	case exp == nil && got != nil:
		t.Fatalf("%s: expected %s nil, got %v", testName, fieldName, got)
	case exp != nil && got == nil:
		t.Fatalf("%s: expected %s %v, got nil", testName, fieldName, exp)
	}
}

func compStructSlices[T any](t *testing.T, testName, fieldName string, expItems, gotItems []T, expLengths map[string]int) {
	expLen, ok := expLengths[fieldName]
	if !ok {
		return
	}
	compare(t, testName, fieldName+" length", expLen, len(gotItems), nil)

	testStructSlices(t, testName, fieldName, expItems, gotItems)

	for i, item := range expItems {
		compStructs(t, testName, "bribe chances", item, gotItems[i])
	}
}

func compStructs[T any](t *testing.T, testName, fieldName string, exp, got T) {
	t.Helper()

	if !reflect.DeepEqual(exp, got) {
		t.Fatalf("%s: expected %s %v, got %v", testName, fieldName, exp, got)
	}
}

func compStructPtrs[T any](t *testing.T, testName, fieldName string, exp, got *T, dontCheck map[string]bool) {
	t.Helper()

	if dontCheck != nil && dontCheck[fieldName] {
		return
	}

	compPtr(t, testName, fieldName, exp, got, compStructs)
}

// checks if two not-nullable apiResources are equal
func compAPIResources[T HasAPIResource](t *testing.T, cfg *Config, testName, fieldName, expPath string, gotRes T, dontCheck map[string]bool) {
	t.Helper()

	if dontCheck != nil && dontCheck[fieldName] {
		return
	}

	expURL := cfg.completeURL(expPath)
	gotURL := gotRes.getAPIResource().getURL()

	compare(t, testName, fieldName, expURL, gotURL, nil)
}

// don't know if I really need this function to have its own switch
// checks if two optional apiResources are equal
func compResourcePtrs[T HasAPIResource](t *testing.T, cfg *Config, testName, fieldName string, expPathPtr *string, gotResPtr *T, dontCheck map[string]bool) {
	t.Helper()

	if dontCheck != nil && dontCheck[fieldName] {
		return
	}

	switch {
	case expPathPtr == nil && gotResPtr == nil:
		return
	case expPathPtr == nil && gotResPtr != nil:
		res := *gotResPtr
		gotURL := res.getAPIResource().getURL()
		t.Fatalf("%s: expected nil for %s, but got %s", testName, fieldName, gotURL)
	case expPathPtr != nil && gotResPtr == nil:
		t.Fatalf("%s: expected %s %v, got nil", testName, fieldName, *expPathPtr)
	default:
		gotRes := *gotResPtr
		expPath := *expPathPtr

		compAPIResources(t, cfg, testName, fieldName, expPath, gotRes, dontCheck)
	}
}

func compPageURL(t *testing.T, cfg *Config, testName, fieldName string, expPathPtr, gotURLPtr *string, dontCheck map[string]bool) {
	if dontCheck != nil && dontCheck[fieldName] {
		return
	}

	var expURLPtr *string

	if expPathPtr != nil {
		expPath := *expPathPtr
		expURL := cfg.completeURL(expPath)
		expURLPtr = &expURL
	}

	compare(t, testName, fieldName, expURLPtr, gotURLPtr, nil)
}
