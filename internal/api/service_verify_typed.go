package api

import (
	"fmt"
	"net/http"
)


func vfIntId(val int32, fieldName string, rules FieldDoc) error {
	if 	(rules.MinVal != nil && val < *rules.MinVal) ||
		(rules.MaxVal != nil && val > *rules.MaxVal) {
		return newHTTPError(http.StatusBadRequest, fmt.Sprintf("value '%d' of field '%s' is out of range. allowed range: %d to %d.", val, fieldName, *rules.MinVal, *rules.MaxVal), nil)
	}

	return nil
}

func vfEnum(typedVal string, fieldName string, rules FieldDoc) error {
	if rules.EnumValues == nil {
		return nil
	}

	var isValidEnum bool

	for _, enumVal := range rules.EnumValues {
		if typedVal == enumVal {
			isValidEnum = true
		}
	}

	if !isValidEnum {
		return newHTTPError(http.StatusBadRequest, fmt.Sprintf("value '%s' of field '%s' is not a valid enum value.", typedVal, fieldName), nil)
	}

	return nil
}

func vfArray[T any](arr []T, fieldName string, rules FieldDoc) error {
	if rules.MaxArrayLen == nil {
		return nil
	}

	if len(arr) > *rules.MaxArrayLen {
		return newHTTPError(http.StatusBadRequest, fmt.Sprintf("too many items in '%s' array. got: %d, maximum allowed: %d", fieldName, len(arr), *rules.MaxArrayLen), nil)
	}

	return nil
}