package api

import (
	"fmt"
	"net/http"
	"slices"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func vfExistingFields(userValues map[FieldName]any, paramsRules ParamsRules) error {
	for fieldName := range userValues {
		_, exists := paramsRules[fieldName]
		if !exists {
			return newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid field: '%s'.", fieldName), nil)
		}
	}

	return nil
}

func basicFieldChecks[T any](val T, fieldName FieldName, userValues map[FieldName]any, paramsRules ParamsRules) (T, error) {
	var zero T
	doc := paramsRules[fieldName].Doc
	valIsPresent := hasVal(fieldName, userValues)

	err := vfRequired(fieldName, doc, valIsPresent)
	if err != nil {
		return zero, err
	}

	err = vfConflictsWith(fieldName, doc, valIsPresent, userValues, paramsRules)
	if err != nil {
		return zero, err
	}

	err = vfRequiredOr(fieldName, doc, valIsPresent, userValues)
	if err != nil {
		return zero, err
	}

	err = vfRequiresAll(fieldName, doc, valIsPresent, userValues)
	if err != nil {
		return zero, err
	}

	err = vfRequiresOne(fieldName, doc, valIsPresent, userValues)
	if err != nil {
		return zero, err
	}

	err = vfAllowedIDs(val, fieldName, doc, valIsPresent, userValues)
	if err != nil {
		return zero, err
	}

	err = vfAllowedIdValPairs(val, fieldName, doc, valIsPresent, userValues)
	if err != nil {
		return zero, err
	}

	return assignDefaultVal(val, valIsPresent, doc), nil
}

func vfRequired(fieldName FieldName, doc FieldDoc, valIsPresent bool) error {
	if !valIsPresent && doc.Required {
		return newHTTPError(http.StatusBadRequest, fmt.Sprintf("field '%s' can't be empty.", fieldName), nil)
	}

	return nil
}

func vfConflictsWith(fieldName FieldName, doc FieldDoc, valIsPresent bool, userValues map[FieldName]any, paramsRules ParamsRules) error {
	if doc.ConflictsWith == nil {
		return nil
	}

	for _, conflictingField := range doc.ConflictsWith {
		if valIsPresent && hasVal(conflictingField, userValues) {
			return newHTTPError(http.StatusBadRequest, fmt.Sprintf("field '%s' can't be used in combination with field '%s'.", fieldName, paramsRules[conflictingField].Doc.Field), nil)
		}
	}

	return nil
}

func vfRequiredOr(fieldName FieldName, doc FieldDoc, valIsPresent bool, userValues map[FieldName]any) error {
	if doc.RequiredOr == nil || valIsPresent {
		return nil
	}

	for _, option := range doc.RequiredOr {
		if hasVal(option, userValues) {
			return nil
		}
	}

	requiredFields := append(doc.RequiredOr, fieldName)

	return newHTTPError(http.StatusBadRequest, fmt.Sprintf("at least one of these fields must have a value: %s.", formatPfnSlice(requiredFields)), nil)
}

func vfRequiresAll(fieldName FieldName, doc FieldDoc, valIsPresent bool, userValues map[FieldName]any) error {
	if doc.RequiresAll == nil || !valIsPresent {
		return nil
	}

	for _, field := range doc.RequiresAll {
		if !hasVal(field, userValues) {
			return newHTTPError(http.StatusBadRequest, fmt.Sprintf("field '%s' requires the following fields to have a value: %s.", fieldName, formatPfnSlice(doc.RequiresAll)), nil)
		}
	}

	return nil
}

func vfRequiresOne(fieldName FieldName, doc FieldDoc, valIsPresent bool, userValues map[FieldName]any) error {
	if doc.RequiresOne == nil || !valIsPresent {
		return nil
	}

	for _, field := range doc.RequiresOne {
		if hasVal(field, userValues) {
			return nil
		}
	}

	return newHTTPError(http.StatusBadRequest, fmt.Sprintf("field '%s' requires one of the following fields to have a value: %s.", fieldName, formatPfnSlice(doc.RequiresOne)), nil)
}

func vfAllowedIDs(val any, fieldName FieldName, doc FieldDoc, valIsPresent bool, userValues map[FieldName]any) error {
	if !valIsPresent || len(doc.AllowedIDs) == 0 || valIsPointer(val) {
		return nil
	}

	id, err := getInt32FromUserVals(pfnUnitID, userValues)
	if err != nil {
		return nil
	}

	if !slices.Contains(doc.AllowedIDs, id) {
		return newHTTPError(http.StatusBadRequest, fmt.Sprintf("field '%s' can only be used with the following ids: %s.", fieldName, h.FormatIntSlice(doc.AllowedIDs)), nil)
	}

	return nil
}

func vfAllowedIdValPairs(val any, fieldName FieldName, doc FieldDoc, valIsPresent bool, userValues map[FieldName]any) error {
	if !valIsPresent || len(doc.AllowedIdValPairs) == 0 || valIsPointer(val) {
		return nil
	}

	if doc.ParentIdField == nil {
		return newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't find parent id field for field '%s'", fieldName), nil)
	}

	id, _ := getInt32FromUserVals(*doc.ParentIdField, userValues)
	value, err := getInt32FromUserVals(fieldName, userValues)
	if err != nil {
		return nil
	}

	for _, entry := range doc.AllowedIdValPairs {
		if id != entry.ID {
			continue
		}

		if value > entry.MaxVal || value <= 0 {
			return newHTTPError(http.StatusBadRequest, fmt.Sprintf("provided value '%d' used for '%s' is out of range for id '%d'. max val: %d. allowed vals: %s.", value, fieldName, id, entry.MaxVal, h.IntSliceToString(entry.Values)), nil)
		}

		if !slices.Contains(entry.Values, value) {
			return newHTTPError(http.StatusBadRequest, fmt.Sprintf("provided value '%d' used for '%s' is not available for id '%d'. allowed vals: %s.", value, fieldName, id, h.IntSliceToString(entry.Values)), nil)
		}

		return nil
	}

	return newHTTPError(http.StatusBadRequest, fmt.Sprintf("field '%s' can only be used with the following id-value pairs: %s.", fieldName, formatAllowedIdValPairs(doc.AllowedIdValPairs)), nil)
}

// sideNote: a pointer is completely optional, so it will never have a default value.
// otherwise it wouldn't be a pointer
func assignDefaultVal[T any](val T, valIsPresent bool, doc FieldDoc) T {
	if valIsPresent || !hasDefaultVal(doc) {
		return val
	}

	typedDefault, ok := doc.DefaultVal.(T)
	if ok {
		return typedDefault
	}

	return val
}
