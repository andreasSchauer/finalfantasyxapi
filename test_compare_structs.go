package main

import (
	"reflect"
)

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
	test.t.Helper()
	checkStructSlices(test, fieldName, exp, got)

	for i := range exp {
		compStructs(test, fieldName, exp[i], got[i])
	}
}

// checks if a pointer to a test struct is equal to a pointer to the original struct. uses a compare function.
// testItemAmount == itemAmount (both Ptrs)
func compTestStructPtrs[E, G any](test test, fieldName string, exp *E, got *G, compFn func(test, E, G)) {
	test.t.Helper()
	
	if test.dontCheck != nil && test.dontCheck[fieldName] {
		return
	}
	
	if !bothPtrsPresent(test, fieldName, exp, got) {
		return
	}

	compFn(test, *exp, *got)
}

// checks if a slice of test structs is equal to the slice of original structs. uses a compare function. if the slice is not nullable, it needs to be explicitely ignored.
// []testItemAmount == []itemAmount
func compTestStructSlices[E, G any](test test, fieldName string, exp []E, got []G, compFn func(test, E, G)) {
	test.t.Helper()
	checkStructSlices(test, fieldName, exp, got)

	for i := range exp {
		compFn(test, exp[i], got[i])
	}
}

func checkStructSlices[E, G any](test test, fieldName string, exp []E, got []G) {
	test.t.Helper()
	compLength(test, fieldName, len(got))

	dontCheck := test.dontCheck
	if dontCheck != nil && dontCheck[fieldName] {
		return
	}

	if !bothStructSlicesPresent(test, fieldName, exp, got) {
		return
	}
}
