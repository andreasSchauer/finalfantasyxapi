package api

import (
	"fmt"
	"net/http"
	"reflect"
)

func basicFieldChecks[T any](val T, fieldName FieldName, values map[FieldName]any, lookup map[FieldName]FieldDoc) (T, error) {
	var zero T
	rules := lookup[fieldName]
	valIsPresent := hasVal(val)

	err := vfRequired(fieldName, rules, valIsPresent)
	if err != nil {
		return zero, err
	}

	err = vfConflictsWith(fieldName, rules, valIsPresent, values, lookup)
	if err != nil {
		return zero, err
	}

	err = vfRequiredOr(fieldName, rules, valIsPresent)
	if err != nil {
		return zero, err
	}

	return assignDefaultVal(val, valIsPresent, rules), nil
}

func vfRequired(fieldName FieldName, rules FieldDoc, valIsPresent bool) error {
	if !valIsPresent && rules.Required {
		return newHTTPError(http.StatusBadRequest, fmt.Sprintf("field '%s' can't be empty.", fieldName), nil)
	}

	return nil
}

func vfConflictsWith(fieldName FieldName, rules FieldDoc, valIsPresent bool, values map[FieldName]any, lookup map[FieldName]FieldDoc) error {
	if rules.ConflictsWith == nil {
		return nil
	}

	for _, conflictingField := range rules.ConflictsWith {
		if valIsPresent && hasVal(values[conflictingField]) {
			return newHTTPError(http.StatusBadRequest, fmt.Sprintf("field '%s' can't be used in combination with field '%s'.", fieldName, lookup[conflictingField].Field), nil)
		}
	}

	return nil
}

func vfRequiredOr(fieldName FieldName, rules FieldDoc, valIsPresent bool) error {
	if rules.RequiredOr == nil || valIsPresent {
		return nil
	}
	
	var onePresent bool

	for _, option := range rules.RequiredOr {
		if hasVal(option) {
			onePresent = true
			break
		}
	}

	if !onePresent {
		requiredFields := append(rules.RequiredOr, fieldName)
		return newHTTPError(http.StatusBadRequest, fmt.Sprintf("at least one of these fields must have a value: '%s'.", formatPfnSlice(requiredFields)), nil)
	}

	return nil
}

// sideNote: a pointer is completely optional, so it will never have a default value.
// otherwise it wouldn't be a pointer
func assignDefaultVal[T any](val T, valIsPresent bool, rules FieldDoc) T {
	if valIsPresent || !hasVal(rules.DefaultVal) {
		return val
	}

	typedDefault, ok := rules.DefaultVal.(T)
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
