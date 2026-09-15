package api

import (
	"fmt"
	"net/http"
	"reflect"
	"slices"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func vfExistingFields(valueMap map[FieldName]any, valTree ValidationTree) error {
	for fieldName := range valueMap {
		_, exists := valTree[fieldName]
		if !exists {
			return newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid field: '%s'", fieldName), nil)
		}
	}

	return nil
}

func basicFieldChecks[T any](val T, fieldName FieldName, valueMap map[FieldName]any, valTree ValidationTree) (T, error) {
	var zero T
	doc := valTree[fieldName].Doc
	valIsPresent := hasVal(val)

	err := vfRequired(fieldName, doc, valIsPresent)
	if err != nil {
		return zero, err
	}

	err = vfConflictsWith(fieldName, doc, valIsPresent, valueMap, valTree)
	if err != nil {
		return zero, err
	}

	err = vfRequiredOr(fieldName, doc, valIsPresent, valueMap)
	if err != nil {
		return zero, err
	}

	err = vfRequiresAll(fieldName, doc, valIsPresent, valueMap)
	if err != nil {
		return zero, err
	}

	err = vfRequiresOne(fieldName, doc, valIsPresent, valueMap)
	if err != nil {
		return zero, err
	}

	err = vfAllowedIDs(val, fieldName, doc, valIsPresent, valueMap)
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

func vfConflictsWith(fieldName FieldName, doc FieldDoc, valIsPresent bool, valueMap map[FieldName]any, valTree ValidationTree) error {
	if doc.ConflictsWith == nil {
		return nil
	}

	for _, conflictingField := range doc.ConflictsWith {
		if valIsPresent && hasVal(valueMap[conflictingField]) {
			return newHTTPError(http.StatusBadRequest, fmt.Sprintf("field '%s' can't be used in combination with field '%s'.", fieldName, valTree[conflictingField].Doc.Field), nil)
		}
	}

	return nil
}

func vfRequiredOr(fieldName FieldName, doc FieldDoc, valIsPresent bool, valueMap map[FieldName]any) error {
	if doc.RequiredOr == nil || valIsPresent {
		return nil
	}

	for _, option := range doc.RequiredOr {
		if hasVal(valueMap[option]) {
			return nil
		}
	}

	requiredFields := append(doc.RequiredOr, fieldName)

	return newHTTPError(http.StatusBadRequest, fmt.Sprintf("at least one of these fields must have a value: %s.", formatPfnSlice(requiredFields)), nil)
}

func vfRequiresAll(fieldName FieldName, doc FieldDoc, valIsPresent bool, valueMap map[FieldName]any) error {
	if doc.RequiresAll == nil || !valIsPresent {
		return nil
	}

	for _, field := range doc.RequiresAll {
		if !hasVal(valueMap[field]) {
			return newHTTPError(http.StatusBadRequest, fmt.Sprintf("field '%s' requires the following fields to have a value: %s", fieldName, formatPfnSlice(doc.RequiresAll)), nil)
		}
	}

	return nil
}

func vfRequiresOne(fieldName FieldName, doc FieldDoc, valIsPresent bool, valueMap map[FieldName]any) error {
	if doc.RequiresOne == nil || !valIsPresent {
		return nil
	}

	for _, field := range doc.RequiresOne {
		if hasVal(valueMap[field]) {
			return nil
		}
	}

	return newHTTPError(http.StatusBadRequest, fmt.Sprintf("field '%s' requires one of the following fields to have a value: %s", fieldName, formatPfnSlice(doc.RequiresOne)), nil)
}

func vfAllowedIDs(val any, fieldName FieldName, doc FieldDoc, valIsPresent bool, valueMap map[FieldName]any) error {
	if !valIsPresent || len(doc.AllowedIDs) == 0 || valIsPointer(val) {
		return nil
	}

	idRaw, ok := valueMap[pfnID]
	if !ok {
		return nil
	}
	id := int32(idRaw.(float64))

	if !slices.Contains(doc.AllowedIDs, id) {
		return newHTTPError(http.StatusBadRequest, fmt.Sprintf("field '%s' can only be used with the following ids: %s", fieldName, h.FormatIntSlice(doc.AllowedIDs)), nil)
	}

	return nil
}

// sideNote: a pointer is completely optional, so it will never have a default value.
// otherwise it wouldn't be a pointer
func assignDefaultVal[T any](val T, valIsPresent bool, doc FieldDoc) T {
	if valIsPresent || !hasVal(doc.DefaultVal) {
		return val
	}

	typedDefault, ok := doc.DefaultVal.(T)
	if ok {
		return typedDefault
	}

	return val
}

func hasVal(val any) bool {
	if val == nil {
		return false
	}

	switch t := val.(type) {
	case string:
		return t != ""

	case int32:
		return t != 0

	case int:
		return t != 0

	case float64:
		return t != 0

	case bool:
		return t
	}

	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.Pointer:
		return !v.IsNil()

	case reflect.Slice:
		return !(v.IsNil() || v.Len() == 0)

	case reflect.Struct:
		return !v.IsZero()
	}

	return false
}

func valIsPointer(val any) bool {
	return reflect.ValueOf(val).Kind() == reflect.Pointer

}
