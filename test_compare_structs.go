package main

import (
	"errors"
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

// checks if two struct pointers with the same type are equal. ptr presence is checked via compPtr
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
	err := sliceBasicChecks(test, fieldName, exp, got)
	if errors.Is(err, errIgnoredField) {
		return
	}

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

// checks if a slice of test structs is equal to the slice of original structs, meaning it contains every entry of the original struct in the correct order. uses a compare function. if the slice is not nullable, it needs to be explicitely ignored.
// []testItemAmount == []itemAmount
func compTestStructSlices[E, G any](test test, fieldName string, exp []E, got []G, compFn func(test, E, G)) {
	test.t.Helper()
	err := sliceBasicChecks(test, fieldName, exp, got)
	if errors.Is(err, errIgnoredField) {
		return
	}

	for i := range exp {
		compFn(test, exp[i], got[i])
	}
}

// checks if all stated testStructs with index are present in the gotStruct slice by using a testStruct's targetIndex.
func checkTestStructsInSlice[E testStructIdx, G any](test test, fieldName string, exp []E, got []G, compFn func(test, E, G)) {
	test.t.Helper()
	err := sliceBasicChecks(test, fieldName, exp, got)
	if errors.Is(err, errIgnoredField) {
		return
	}

	for _, item := range exp {
		i := item.GetIndex()
		compFn(test, item, got[i])
	}
}

// checks the length of a slice, whether it should be ignored (returns errIgnoredField in that case), and if the exp and got slice are present
func sliceBasicChecks[E, G any](test test, fieldName string, exp []E, got []G) error {
	compLength(test, fieldName, len(got))

	dontCheck := test.dontCheck
	if dontCheck != nil && dontCheck[fieldName] {
		return errIgnoredField
	}

	bothSlicesPresent(test, fieldName, exp, got)
	return nil
}
