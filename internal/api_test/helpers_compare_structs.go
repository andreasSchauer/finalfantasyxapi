package apitest

import (
	"errors"
	"fmt"
	"reflect"
)

// checks if two structs with the same type are equal.
// itemAmount == itemAmount
func (test test) compStructs[T any](fieldName string, exp, got T) {
	test.t.Helper()

	if test.dontCheck != nil && test.dontCheck[fieldName] {
		return
	}

	if !reflect.DeepEqual(exp, got) {
		test.t.Errorf("expected %s %v, got %v\n\n", fieldName, exp, got)
	}
}

func (test test) compStructsWithFn[T any](fieldName string, exp, got T, compFn compFnEqual[T]) {
	test.t.Helper()

	if test.dontCheck != nil && test.dontCheck[fieldName] {
		return
	}

	compFn(test, fieldName, exp, got)
}

func (test test) compStructPtrsWithFn[T any](fieldName string, exp, got *T, compFn compFnEqual[T]) {
	test.t.Helper()

	if test.dontCheck != nil && test.dontCheck[fieldName] {
		return
	}

	if !test.presencePtrs(fieldName, exp, got) {
		return
	}

	compFn(test, fieldName, *exp, *got)
}

func (test test) compStructSlicesWithFn[T any](fieldName string, exp, got []T, compFn compFnEqual[T]) {
	test.t.Helper()
	err := test.checkSliceBasic(fieldName, exp, got)
	if errors.Is(err, errIgnoredField) {
		return
	}

	for i := range exp {
		compFn(test, fieldName, exp[i], got[i])
	}
}

// checks if two struct pointers with the same type are equal. ptr presence is checked via compPtr
// itemAmount == itemAmount (both ptrs)
func (test test) compStructPtrs[T any](fieldName string, exp, got *T) {
	test.t.Helper()

	if test.dontCheck != nil && test.dontCheck[fieldName] {
		return
	}

	test.compPtr(fieldName, exp, got, test.compStructs)
}

// checks if two same-typed struct slices are equal. if the slice is not nullable, it needs to be explicitely ignored.
// []itemAmount == []itemAmount
func (test test) compStructSlices[T any](fieldName string, exp, got []T) {
	test.t.Helper()
	err := test.checkSliceBasic(fieldName, exp, got)
	if errors.Is(err, errIgnoredField) {
		return
	}

	for i := range exp {
		test.compStructs(fieldName, exp[i], got[i])
	}
}

// checks if a pointer to a test struct is equal to a pointer to the original struct. uses a compare function.
// testItemAmount == itemAmount (both Ptrs)
func (test test) compTestStructPtrs[E, G any](fieldName string, exp *E, got *G, compFn compFnDifferent[E, G]) {
	test.t.Helper()

	if test.dontCheck != nil && test.dontCheck[fieldName] {
		return
	}

	if !test.presencePtrs(fieldName, exp, got) {
		return
	}

	compFn(test, fieldName, *exp, *got)
}

// checks if a slice of test structs is equal to the slice of original structs, meaning it contains every entry of the original struct in the correct order. uses a compare function. if the slice is not nullable, it needs to be explicitely ignored.
// []testItemAmount == []itemAmount
func (test test) compTestStructSlices[E, G any](fieldName string, exp []E, got []G, compFn compFnDifferent[E, G]) {
	test.t.Helper()
	err := test.checkSliceBasic(fieldName, exp, got)
	if errors.Is(err, errIgnoredField) {
		return
	}

	for i := range exp {
		subName := fmt.Sprintf("%s %d", fieldName, i)
		compFn(test, subName, exp[i], got[i])
	}
}

// checks if all stated testStructs with index are present in the gotStruct slice by using a testStruct's targetIndex.
func (test test) checkTestStructsInSlice[E testStructIdx, G any](fieldName string, exp []E, got []G, compFn compFnDifferent[E, G]) {
	test.t.Helper()
	err := test.checkSliceBasic(fieldName, exp, got)
	if errors.Is(err, errIgnoredField) {
		return
	}

	for _, item := range exp {
		i := item.GetIndex()
		subName := fmt.Sprintf("%s %d", fieldName, i)
		compFn(test, subName, item, got[i])
	}
}

// checks the length of a slice, whether it should be ignored (returns errIgnoredField in that case), and if the exp and got slice are present
func (test test) checkSliceBasic[E, G any](fieldName string, exp []E, got []G) error {
	test.t.Helper()
	test.compLength(fieldName, len(got))

	dontCheck := test.dontCheck
	if dontCheck != nil && dontCheck[fieldName] {
		return errIgnoredField
	}

	test.presenceSlices(fieldName, exp, got)
	return nil
}
