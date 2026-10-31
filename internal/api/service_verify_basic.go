package api

import (
	"fmt"
	"net/http"
	"slices"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)


type fieldVerifyer[T any] struct {
	fieldName		FieldName
	val 			T
	doc 			FieldDoc
	valIsPresent	bool
	pv				paramsVerifyer
}

func (v paramsVerifyer) createFieldVerifyer[T any](val T, fieldName FieldName) fieldVerifyer[T] {
	return fieldVerifyer[T]{
		fieldName: 		fieldName,
		val: 			val,
		doc: 			v.paramsRules[fieldName].Doc,
		valIsPresent: 	hasVal(fieldName, v.userValues),
		pv: 			v,
	}
}

func (v paramsVerifyer) vfExistingFields() error {
	for fieldName := range v.userValues {
		_, exists := v.paramsRules[fieldName]
		if !exists {
			return newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid field: '%s'.", fieldName), nil)
		}
	}

	return nil
}

func (v paramsVerifyer) vfField[T any](val T, fieldName FieldName) (T, error) {
	var zero T
	fv := v.createFieldVerifyer(val, fieldName)

	err := fv.vfRequired()
	if err != nil {
		return zero, err
	}

	err = fv.vfConflictsWith()
	if err != nil {
		return zero, err
	}

	err = fv.vfRequiredOr()
	if err != nil {
		return zero, err
	}

	err = fv.vfRequiresAll()
	if err != nil {
		return zero, err
	}

	err = fv.vfRequiresOne()
	if err != nil {
		return zero, err
	}

	err = fv.vfAllowedIDs()
	if err != nil {
		return zero, err
	}

	err = fv.vfAllowedIDValPairs()
	if err != nil {
		return zero, err
	}

	return fv.assignDefaultVal(), nil
}

func (v fieldVerifyer[T]) vfRequired() error {
	if !v.valIsPresent && v.doc.Required {
		return newHTTPError(http.StatusBadRequest, fmt.Sprintf("field '%s' can't be empty.", v.fieldName), nil)
	}

	return nil
}

func (v fieldVerifyer[T]) vfConflictsWith() error {
	if v.doc.ConflictsWith == nil {
		return nil
	}

	for _, conflictingField := range v.doc.ConflictsWith {
		if v.valIsPresent && hasVal(conflictingField, v.pv.userValues) {
			return newHTTPError(http.StatusBadRequest, fmt.Sprintf("field '%s' can't be used in combination with field '%s'.", v.fieldName, v.pv.paramsRules[conflictingField].Doc.Field), nil)
		}
	}

	return nil
}

func (v fieldVerifyer[T]) vfRequiredOr() error {
	if v.doc.RequiredOr == nil || v.valIsPresent {
		return nil
	}

	for _, option := range v.doc.RequiredOr {
		if hasVal(option, v.pv.userValues) {
			return nil
		}
	}

	requiredFields := append(v.doc.RequiredOr, v.fieldName)

	return newHTTPError(http.StatusBadRequest, fmt.Sprintf("at least one of these fields must have a value: %s.", formatPfnSlice(requiredFields)), nil)
}

func (v fieldVerifyer[T]) vfRequiresAll() error {
	if v.doc.RequiresAll == nil || !v.valIsPresent {
		return nil
	}

	for _, field := range v.doc.RequiresAll {
		if !hasVal(field, v.pv.userValues) {
			return newHTTPError(http.StatusBadRequest, fmt.Sprintf("field '%s' requires the following fields to have a value: %s.", v.fieldName, formatPfnSlice(v.doc.RequiresAll)), nil)
		}
	}

	return nil
}

func (v fieldVerifyer[T]) vfRequiresOne() error {
	if v.doc.RequiresOne == nil || !v.valIsPresent {
		return nil
	}

	for _, field := range v.doc.RequiresOne {
		if hasVal(field, v.pv.userValues) {
			return nil
		}
	}

	return newHTTPError(http.StatusBadRequest, fmt.Sprintf("field '%s' requires one of the following fields to have a value: %s.", v.fieldName, formatPfnSlice(v.doc.RequiresOne)), nil)
}

func (v fieldVerifyer[T]) vfAllowedIDs() error {
	if !v.valIsPresent || len(v.doc.AllowedIDs) == 0 || valIsPointer(v.val) {
		return nil
	}

	id, err := getInt32FromUserVals(PfnUnitID, v.pv.userValues)
	if err != nil {
		return nil
	}

	if !slices.Contains(v.doc.AllowedIDs, id) {
		return newHTTPError(http.StatusBadRequest, fmt.Sprintf("field '%s' can only be used with the following ids: %s.", v.fieldName, h.FormatIntSlice(v.doc.AllowedIDs)), nil)
	}

	return nil
}

func (v fieldVerifyer[T]) vfAllowedIDValPairs() error {
	if !v.valIsPresent || len(v.doc.AllowedIDValPairs) == 0 || valIsPointer(v.val) {
		return nil
	}

	if v.doc.ParentIDField == nil {
		return newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't find parent id field for field '%s'", v.fieldName), nil)
	}

	id, _ := getInt32FromUserVals(*v.doc.ParentIDField, v.pv.userValues)
	value, err := getInt32FromUserVals(v.fieldName, v.pv.userValues)
	if err != nil {
		return nil
	}

	for _, entry := range v.doc.AllowedIDValPairs {
		if id != entry.ID {
			continue
		}

		if value > entry.MaxVal || value <= 0 {
			return newHTTPError(http.StatusBadRequest, fmt.Sprintf("provided value '%d' used for '%s' is out of range for id '%d'. max val: %d. allowed vals: %s.", value, v.fieldName, id, entry.MaxVal, h.IntSliceToString(entry.Values)), nil)
		}

		if !slices.Contains(entry.Values, value) {
			return newHTTPError(http.StatusBadRequest, fmt.Sprintf("provided value '%d' used for '%s' is not available for id '%d'. allowed vals: %s.", value, v.fieldName, id, h.IntSliceToString(entry.Values)), nil)
		}

		return nil
	}

	return newHTTPError(http.StatusBadRequest, fmt.Sprintf("field '%s' can only be used with the following id-value pairs: %s.", v.fieldName, formatAllowedIDValPairs(v.doc.AllowedIDValPairs)), nil)
}


func (v fieldVerifyer[T]) assignDefaultVal() T {
	if v.valIsPresent || !hasDefaultVal(v.doc) {
		return v.val
	}

	typedDefault, ok := v.doc.DefaultVal.(T)
	if ok {
		return typedDefault
	}

	return v.val
}