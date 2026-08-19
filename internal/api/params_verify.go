package api

import (
	"fmt"
	"net/http"
	"reflect"
)

// custom logic like the 10 monsters for /turn-order will be done in the verifyTurnOrderParams function directly. these are simply generic checks that apply to all fields
// make smaller helper functions out of these
// how to handle pointers and slices?
func verifyParamField[T any](val T, fieldName string, values map[string]any, lookup map[string]FieldDoc) (T, error) {
	var zero T
	rules := lookup[fieldName]
	valIsZero := isZero(val)
	
	// required
	if valIsZero && rules.Required {
		return zero, newHTTPError(http.StatusBadRequest, fmt.Sprintf("field '%s' can't be empty.", fieldName), nil)
	}

	// conflictsWith
	if rules.ConflictsWith != nil {
		for _, conflictingField := range rules.ConflictsWith {
			if !valIsZero && !isZero(values[conflictingField]) {
				return zero, newHTTPError(http.StatusBadRequest, fmt.Sprintf("field '%s' can't be used in combination with field '%s'.", fieldName, lookup[conflictingField].Field), nil)
			}
		}
	}

	// default val
	if valIsZero && !isZero(rules.DefaultVal) {
		typedDefault, ok := rules.DefaultVal.(T)
		if ok {
			val = typedDefault
		}
	}
	
	switch typedVal := any(val).(type) {
	// min/max val
	case int32:
		if 	(rules.MinVal != nil && typedVal < *rules.MinVal) ||
			(rules.MaxVal != nil && typedVal > *rules.MaxVal) {
			return zero, newHTTPError(http.StatusBadRequest, fmt.Sprintf("value '%d' of field '%s' is out of range.", typedVal, fieldName), nil)
		}
		
	// enum values
	case string:
		if rules.EnumValues != nil {
			var isValidEnum bool

			for _, enumVal := range rules.EnumValues {
				if typedVal == enumVal {
					isValidEnum = true
				}
			}

			if !isValidEnum {
				return zero, newHTTPError(http.StatusBadRequest, fmt.Sprintf("value '%s' of field '%s' is not a valid enum value.", typedVal, fieldName), nil)
			}
		}
	}

	// child properties? or do slices separately?
	// with the fieldDocSliceToMap func, I can easily make a function just for the child properties, I think
	// will worry about that, when I have params with actual slices


	return val, nil
}

// to helpers?
func isZero(val any) bool {
	if val == nil {
		return true
	}

	switch t := val.(type) {
	case string:
		return t == ""
	
	case int32:
		return t == 0

	case bool:
		return !t
	}

	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.Pointer:
		return v.IsNil()

	case reflect.Slice:
		return v.IsNil() || v.Len() == 0
	}

	return false
}
