package apitesttypes

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/errs"
)

// checks if two structs with the same type are equal.
// itemAmount == itemAmount
func (test Test) compStructs[T any](fieldName string, exp, got T) {
	test.T.Helper()

	if test.DontCheck != nil && test.DontCheck[fieldName] {
		return
	}

	if !reflect.DeepEqual(exp, got) {
		test.T.Errorf("expected %s %v, got %v\n\n", fieldName, exp, got)
	}
}

// checks if two struct pointers with the same type are equal. ptr presence is checked via compPtr
// itemAmount == itemAmount (both ptrs)
func (test Test) compStructPtrs[T any](fieldName string, exp, got *T) {
	test.T.Helper()

	if test.DontCheck != nil && test.DontCheck[fieldName] {
		return
	}

	test.compPtr(fieldName, exp, got, test.compStructs)
}

// checks if two same-typed struct slices are equal. if the slice is not nullable, it needs to be explicitely ignored.
// []itemAmount == []itemAmount
func (test Test) compStructSlices[T any](fieldName string, exp, got []T) {
	test.T.Helper()
	err := test.checkSliceBasic(fieldName, exp, got)
	if errors.Is(err, errs.ErrIgnoredField) {
		return
	}

	for i := range exp {
		test.compStructs(fieldName, exp[i], got[i])
	}
}

// checks if a pointer to a test struct is equal to a pointer to the original struct. uses a compare function.
// testItemAmount == itemAmount (both Ptrs)
func (test Test) compTestStructPtrs[E testStruct[G], G any](fieldName string, exp *E, got *G) {
	test.T.Helper()

	if test.DontCheck != nil && test.DontCheck[fieldName] {
		return
	}

	if !test.presencePtrs(fieldName, exp, got) {
		return
	}

	e := *exp
	e.Compare(test, fieldName, *got)
}

// checks if a slice of test structs is equal to the slice of original structs, meaning it contains every entry of the original struct in the correct order. uses a compare function. if the slice is not nullable, it needs to be explicitely ignored.
// []testItemAmount == []itemAmount
func (test Test) compTestStructSlices[E testStruct[G], G any](fieldName string, exp []E, got []G) {
	test.T.Helper()
	err := test.checkSliceBasic(fieldName, exp, got)
	if errors.Is(err, errs.ErrIgnoredField) {
		return
	}

	for i := range exp {
		subName := fmt.Sprintf("%s %d", fieldName, i)
		exp[i].Compare(test, subName, got[i])
	}
}

// checks if all stated testStructs with index are present in the gotStruct slice by using a testStruct's targetIndex.
func (test Test) checkTestStructsInSlice[E testStructIdx[G], G any](fieldName string, exp []E, got []G) {
	test.T.Helper()
	err := test.checkSliceBasic(fieldName, exp, got)
	if errors.Is(err, errs.ErrIgnoredField) {
		return
	}

	for _, item := range exp {
		i := item.GetIndex()
		subName := fmt.Sprintf("%s %d", fieldName, i)
		item.Compare(test, subName, got[i])
	}
}

// checks the length of a slice, whether it should be ignored (returns errIgnoredField in that case), and if the exp and got slice are present
func (test Test) checkSliceBasic[E, G any](fieldName string, exp []E, got []G) error {
	test.T.Helper()
	test.compLength(fieldName, len(got))

	dontCheck := test.DontCheck
	if dontCheck != nil && dontCheck[fieldName] {
		return errs.ErrIgnoredField
	}

	test.presenceSlices(fieldName, exp, got)
	return nil
}
