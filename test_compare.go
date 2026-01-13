package main

import (
	"reflect"
)

// checks if two values of a field are equal
func compare(test test, fieldName string, exp, got any) {
	t := test.t
	testName := test.name
	dontCheck := test.dontCheck
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
		compInt(test, fieldName, e, g)

	case int32:
		g, ok := got.(int32)
		if !ok {
			t.Fatalf("%s: %s type mismatch: expected int32, got %T", testName, fieldName, got)
		}
		compInt32(test, fieldName, e, g)

	case float32:
		g, ok := got.(float32)
		if !ok {
			t.Fatalf("%s: %s type mismatch: expected float32, got %T", testName, fieldName, got)
		}
		compFloat32(test, fieldName, e, g)

	case string:
		g, ok := got.(string)
		if !ok {
			t.Fatalf("%s: %s type mismatch: expected string, got %T", testName, fieldName, got)
		}
		compString(test, fieldName, e, g)

	case bool:
		g, ok := got.(bool)
		if !ok {
			t.Fatalf("%s: %s type mismatch: expected bool, got %T", testName, fieldName, got)
		}
		compBool(test, fieldName, e, g)

	case *int32:
		g, ok := got.(*int32)
		if !ok {
			t.Fatalf("%s: %s type mismatch: expected *int32, got %T", testName, fieldName, got)
		}
		compInt32Ptr(test, fieldName, e, g)

	case *float32:
		g, ok := got.(*float32)
		if !ok {
			t.Fatalf("%s: %s type mismatch: expected *float32, got %T", testName, fieldName, got)
		}
		compFloat32Ptr(test, fieldName, e, g)

	case *string:
		g, ok := got.(*string)
		if !ok {
			t.Fatalf("%s: %s type mismatch: expected *string, got %T", testName, fieldName, got)
		}
		compStringPtr(test, fieldName, e, g)

	default:
		t.Fatalf("%s: unsupported type for %s: %T", testName, fieldName, exp)
	}
}

func compInt(test test, fieldName string, exp, got int) {
	if exp != got {
		test.t.Fatalf("%s: expected %s %d, got %d", test.name, fieldName, exp, got)
	}
}

func compInt32(test test, fieldName string, exp, got int32) {
	if exp != got {
		test.t.Fatalf("%s: expected %s %d, got %d", test.name, fieldName, exp, got)
	}
}

func compFloat32(test test, fieldName string, exp, got float32) {
	if exp != got {
		test.t.Fatalf("%s: expected %s %.2f, got %.2f", test.name, fieldName, exp, got)
	}
}

func compString(test test, fieldName, exp, got string) {
	if exp != "" && exp != got {
		test.t.Fatalf("%s: expected %s %s, got %s", test.name, fieldName, exp, got)
	}
}

func compBool(test test, fieldName string, exp, got bool) {
	if exp != got {
		test.t.Fatalf("%s: expected %s %t, got %t", test.name, fieldName, exp, got)
	}
}

func compInt32Ptr(test test, fieldName string, exp, got *int32) {
	compPtr(test, fieldName, exp, got, compInt32)
}

func compFloat32Ptr(test test, fieldName string, exp, got *float32) {
	compPtr(test, fieldName, exp, got, compFloat32)
}

func compStringPtr(test test, fieldName string, exp, got *string) {
	compPtr(test, fieldName, exp, got, compString)
}

func bothPtrsPresent[T, U any](test test, fieldName string, exp *T, got *U) bool {
	switch {
	case exp == nil && got == nil:
		return false
	case exp == nil && got != nil:
		test.t.Fatalf("%s: expected %s nil, got %v", test.name, fieldName, *got)
		return false
	case exp != nil && got == nil:
		test.t.Fatalf("%s: expected %s %v, got nil", test.name, fieldName, *exp)
		return false
	default:
		return true
	}
}

func compPtr[T any](test test, fieldName string, exp, got *T, compFunc func(test, string, T, T)) {
	if !bothPtrsPresent(test, fieldName, exp, got) {
		return
	}
	compFunc(test, fieldName, *exp, *got)
}

func testStructSlices[T any](test test, fieldName string, exp, got []T) {
	switch {
	case exp == nil && got == nil:
		return
	case exp == nil && got != nil:
		test.t.Fatalf("%s: expected %s nil, got %v", test.name, fieldName, got)
	case exp != nil && got == nil:
		test.t.Fatalf("%s: expected %s %v, got nil", test.name, fieldName, exp)
	}
}

func compStructSlices[T any](test test, fieldName string, expItems, gotItems []T) {
	expLen, ok := test.expLengths[fieldName]
	if !ok {
		return
	}
	compare(test, fieldName + " length", expLen, len(gotItems))

	testStructSlices(test, fieldName, expItems, gotItems)

	for i, item := range expItems {
		compStructs(test, "bribe chances", item, gotItems[i])
	}
}

func compStructs[T any](test test, fieldName string, exp, got T) {
	test.t.Helper()

	if !reflect.DeepEqual(exp, got) {
		test.t.Fatalf("%s: expected %s %v, got %v", test.name, fieldName, exp, got)
	}
}

func compStructPtrs[T any](test test, fieldName string, exp, got *T) {
	test.t.Helper()

	if test.dontCheck != nil && test.dontCheck[fieldName] {
		return
	}

	compPtr(test, fieldName, exp, got, compStructs)
}

// checks if two not-nullable apiResources are equal
func compAPIResources[T HasAPIResource](test test, fieldName, expPath string, gotRes T) {
	test.t.Helper()

	if test.dontCheck != nil && test.dontCheck[fieldName] {
		return
	}

	expURL := test.cfg.completeTestURL(expPath)
	gotURL := gotRes.GetAPIResource().GetURL()

	compare(test, fieldName, expURL, gotURL)
}

func compAPIResourcesFromID[T HasAPIResource](test test, fieldName, endpoint string, expID int32, gotRes T) {
	expPath := completeTestPath(endpoint, expID)
	compAPIResources(test, fieldName, expPath, gotRes)
}

// don't know if I really need this function to have its own switch
// checks if two optional apiResources are equal
func compResourcePtrs[T HasAPIResource](test test, fieldName string, expPathPtr *string, gotResPtr *T) {
	test.t.Helper()

	if test.dontCheck != nil && test.dontCheck[fieldName] {
		return
	}

	switch {
	case expPathPtr == nil && gotResPtr == nil:
		return
	case expPathPtr == nil && gotResPtr != nil:
		res := *gotResPtr
		gotURL := res.GetAPIResource().GetURL()
		test.t.Fatalf("%s: expected nil for %s, but got %s", test.name, fieldName, gotURL)
	case expPathPtr != nil && gotResPtr == nil:
		test.t.Fatalf("%s: expected %s %v, got nil", test.name, fieldName, *expPathPtr)
	default:
		gotRes := *gotResPtr
		expPath := *expPathPtr

		compAPIResources(test, fieldName, expPath, gotRes)
	}
}

func compResPtrsFromID[T HasAPIResource](test test, fieldName, endpoint string, expIDPtr *int32, gotResPtr *T) {
	var expPathPtr *string

	if expIDPtr == nil {
		expPathPtr = nil
	} else {
		expPath := completeTestPath(endpoint, *expIDPtr)
		expPathPtr = &expPath
	}

	compResourcePtrs(test, fieldName, expPathPtr, gotResPtr)
}

func compPageURL(test test, fieldName string, expPathPtr, gotURLPtr *string) {
	if test.dontCheck != nil && test.dontCheck[fieldName] {
		return
	}

	var expURLPtr *string

	if expPathPtr != nil {
		expPath := *expPathPtr
		expURL := test.cfg.completeTestURL(expPath)
		expURLPtr = &expURL
	}

	compare(test, fieldName, expURLPtr, gotURLPtr)
}
