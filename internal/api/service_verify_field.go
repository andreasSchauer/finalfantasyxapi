package api

import (
	"fmt"
	"net/http"
	"reflect"
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

type fieldVerifyer[T any] struct {
	fieldName    	ts.FieldName
	val          	T
	doc          	FieldDoc
	valIsPresent 	bool
	paramsVerifyer
}

func (v fieldVerifyer[T]) required() error {
	if !v.valIsPresent && v.doc.Required {
		return errs.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("field '%s' can't be empty", v.fieldName), nil)
	}

	return nil
}

func (v fieldVerifyer[T]) conflictsWith() error {
	if v.doc.ConflictsWith == nil {
		return nil
	}

	for _, conflictingField := range v.doc.ConflictsWith {
		if v.valIsPresent && v.hasVal(conflictingField) {
			return errs.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("field '%s' can't be used in combination with field '%s'", v.fieldName, v.paramsRules[conflictingField].Doc.Field), nil)
		}
	}

	return nil
}

func (v fieldVerifyer[T]) requiredOr() error {
	if v.doc.RequiredOr == nil || v.valIsPresent {
		return nil
	}

	if slices.ContainsFunc(v.doc.RequiredOr, v.hasVal) {
		return nil
	}

	requiredFields := append(v.doc.RequiredOr, v.fieldName)

	return errs.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("at least one of these fields must have a value: %s", ts.FormatPfnSlice(requiredFields)), nil)
}

func (v fieldVerifyer[T]) requiresAll() error {
	if v.doc.RequiresAll == nil || !v.valIsPresent {
		return nil
	}

	if slices.ContainsFunc(v.doc.RequiresAll, v.hasNoVal) {
		return errs.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("field '%s' requires the following fields to have a value: %s", v.fieldName, ts.FormatPfnSlice(v.doc.RequiresAll)), nil)
	}

	return nil
}

func (v fieldVerifyer[T]) requiresOne() error {
	if v.doc.RequiresOne == nil || !v.valIsPresent {
		return nil
	}

	if slices.ContainsFunc(v.doc.RequiresOne, v.hasVal) {
		return nil
	}

	return errs.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("field '%s' requires one of the following fields to have a value: %s", v.fieldName, ts.FormatPfnSlice(v.doc.RequiresOne)), nil)
}

func (v fieldVerifyer[T]) allowedIDs() error {
	if !v.valIsPresent || len(v.doc.AllowedIDs) == 0 || v.valIsPointer() {
		return nil
	}

	id, err := getInt32FromUserVals(ts.PfnUnitID, v.userValues)
	if err != nil {
		return nil
	}

	if !slices.Contains(v.doc.AllowedIDs, id) {
		return errs.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("field '%s' can only be used with the following ids: %s", v.fieldName, h.FormatIntSlice(v.doc.AllowedIDs)), nil)
	}

	return nil
}

func (v fieldVerifyer[T]) allowedIDValPairs() error {
	if !v.valIsPresent || len(v.doc.AllowedIDValPairs) == 0 || v.valIsPointer() {
		return nil
	}

	if v.doc.ParentIDField == nil {
		return errs.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't find parent id field for field '%s'", v.fieldName), nil)
	}

	id, _ := getInt32FromUserVals(*v.doc.ParentIDField, v.userValues)
	value, err := getInt32FromUserVals(v.fieldName, v.userValues)
	if err != nil {
		return nil
	}

	for _, entry := range v.doc.AllowedIDValPairs {
		if id != entry.ID {
			continue
		}

		if value > entry.MaxVal || value <= 0 {
			return errs.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("provided value '%d' used for '%s' is out of range for id '%d'. max val: %d. allowed vals: %s", value, v.fieldName, id, entry.MaxVal, h.IntSliceToString(entry.Values)), nil)
		}

		if !slices.Contains(entry.Values, value) {
			return errs.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("provided value '%d' used for '%s' is not available for id '%d'. allowed vals: %s", value, v.fieldName, id, h.IntSliceToString(entry.Values)), nil)
		}

		return nil
	}

	return errs.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("field '%s' can only be used with the following id-value pairs: %s", v.fieldName, formatAllowedIDValPairs(v.doc.AllowedIDValPairs)), nil)
}

func (v fieldVerifyer[T]) assignDefaultVal() T {
	if v.valIsPresent || !v.hasDefaultVal() {
		return v.val
	}

	typedDefault, ok := v.doc.DefaultVal.(T)
	if ok {
		return typedDefault
	}

	return v.val
}

func (v fieldVerifyer[T]) hasDefaultVal() bool {
	return v.doc.DefaultVal != nil
}

func (v fieldVerifyer[T]) valIsPointer() bool {
	return reflect.ValueOf(v.val).Kind() == reflect.Pointer
}
