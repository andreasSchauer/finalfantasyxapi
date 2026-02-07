package main

import (
	"reflect"
)

// checks if two standard-type values or pointers of a field are equal
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
			t.Fatalf("%s: %s type mismatch: expected int, got %T", testName, fieldName, got)
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

// can be used for any direct pointer comparisons
func compPtr[T any](test test, fieldName string, exp, got *T, compFunc func(test, string, T, T)) {
	if !bothPtrsPresent(test, fieldName, exp, got) {
		return
	}
	compFunc(test, fieldName, *exp, *got)
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

func compLength(test test, fieldName string, gotLen int) {
	expLen, ok := test.expLengths[fieldName]
	if ok {
		compare(test, fieldName+" length", expLen, gotLen)
	}
}

// checks if two structs with the same type are equal.
// itemAmount == itemAmount
func compStructs[T any](test test, fieldName string, exp, got T) {
	test.t.Helper()

	if !reflect.DeepEqual(exp, got) {
		test.t.Fatalf("%s: expected %s %v, got %v", test.name, fieldName, exp, got)
	}
}

// checks if two struct pointers with the same type are equal.
// itemAmount == itemAmount (both ptrs)
func compStructPtrs[T any](test test, fieldName string, exp, got *T) {
	test.t.Helper()

	if test.dontCheck != nil && test.dontCheck[fieldName] {
		return
	}

	compPtr(test, fieldName, exp, got, compStructs)
}


// checks if two same-typed struct slices are equal. if the slice is not nullable, it needs to be explicitely ignored.
// []itemAmount == []itemAmount
func compStructSlices[T any](test test, fieldName string, exp, got []T) {
	checkStructSlices(test, fieldName, exp, got)

	for i := range exp {
		compStructs(test, fieldName, exp[i], got[i])
	}
}

// checks if a slice of test structs is equal to the slice of original structs. uses a compare function. if the slice is not nullable, it needs to be explicitely ignored.
// []testItemAmount == []itemAmount
func compTestStructSlices[E, G any](test test, fieldName string, exp []E, got []G, compFn func(test, E, G)) {
	checkStructSlices(test, fieldName, exp, got)

	for i := range exp {
		compFn(test, exp[i], got[i])
	}
}

func checkStructSlices[E, G any](test test, fieldName string, exp []E, got []G) {
	compLength(test, fieldName, len(got))

	dontCheck := test.dontCheck
	if dontCheck != nil && dontCheck[fieldName] {
		return
	}

	if !bothStructSlicesPresent(test, fieldName, exp, got) {
		return
	}
}

// takes an expected path and checks, if it matches the URL of target struct's API Resource.
// /endpoint/23 == struct with api resource linking to host/api/endpoint/23
func compPathApiResource[T HasAPIResource](test test, fieldName, expPath string, gotRes T) {
	test.t.Helper()

	if test.dontCheck != nil && test.dontCheck[fieldName] {
		return
	}

	expURL := completeTestURL(test.cfg, expPath)
	gotURL := gotRes.GetAPIResource().GetURL()

	compare(test, fieldName, expURL, gotURL)
}

// takes an expected path ptr and checks, if it matches the URL of target struct ptr's API Resource.
// /endpoint/23 == struct with api resource linking to host/api/endpoint/23 (both ptrs)
func compPathApiResourcePtrs[T HasAPIResource](test test, fieldName string, expPathPtr *string, gotResPtr *T) {
	test.t.Helper()

	if test.dontCheck != nil && test.dontCheck[fieldName] {
		return
	}

	if !bothResourcePtrsPresent(test, fieldName, expPathPtr, gotResPtr) {
		return
	}

	compPathApiResource(test, fieldName, *expPathPtr, *gotResPtr)
}

// takes an id, assembles the url, and checks, if it matches the URL of target struct's API Resource.
// 23 == struct with api resource linking to host/api/endpoint/23
func compIdApiResource[T HasAPIResource](test test, fieldName, endpoint string, expID int32, gotRes T) {
	expPath := completeTestPath(endpoint, expID)
	compPathApiResource(test, fieldName, expPath, gotRes)
}

// takes an id ptr, assembles the url, and checks, if it matches the URL of target struct ptr's API Resource.
// 23 == struct with api resource linking to host/api/endpoint/23 (both ptrs)
func compIdApiResourcePtrs[T HasAPIResource](test test, fieldName, endpoint string, expIDPtr *int32, gotResPtr *T) {
	test.t.Helper()

	var expPathPtr *string

	if expIDPtr == nil {
		expPathPtr = nil
	} else {
		expPath := completeTestPath(endpoint, *expIDPtr)
		expPathPtr = &expPath
	}

	compPathApiResourcePtrs(test, fieldName, expPathPtr, gotResPtr)
}

// compares the expected pagination path with the got url
func compPageURL(test test, fieldName string, expPathPtr, gotURLPtr *string) {
	if test.dontCheck != nil && test.dontCheck[fieldName] {
		return
	}

	var expURLPtr *string

	if expPathPtr != nil {
		expPath := *expPathPtr
		expURL := completeTestURL(test.cfg, expPath)
		expURLPtr = &expURL
	}

	compare(test, fieldName, expURLPtr, gotURLPtr)
}
