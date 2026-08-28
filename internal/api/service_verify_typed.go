package api

import (
	"fmt"
	"net/http"
	"slices"
)

func vfIntId(_ *Config, val int32, fieldName FieldName, _ map[FieldName]any, valTree ValidationTree) (int32, error) {
	doc := valTree[fieldName].Doc

	if (doc.MinVal != nil && val < *doc.MinVal) ||
		(doc.MaxVal != nil && val > *doc.MaxVal) {
		return 0, newHTTPError(http.StatusBadRequest, fmt.Sprintf("value '%d' of field '%s' is out of range. allowed range: %d to %d.", val, fieldName, *doc.MinVal, *doc.MaxVal), nil)
	}

	return val, nil
}

func vfEnum(_ *Config, typedVal string, fieldName FieldName, _ map[FieldName]any, valTree ValidationTree) (string, error) {
	doc := valTree[fieldName].Doc

	if doc.EnumValues == nil {
		return typedVal, nil
	}

	if !slices.Contains(doc.EnumValues, typedVal) {
		return "", newHTTPError(http.StatusBadRequest, fmt.Sprintf("value '%s' of field '%s' is not a valid enum value.", typedVal, fieldName), nil)
	}

	return typedVal, nil
}

func vfArray[T any](_ *Config, arr []T, fieldName FieldName, _ map[FieldName]any, valTree ValidationTree) ([]T, error) {
	doc := valTree[fieldName].Doc

	if doc.MaxArrayLen == nil {
		return arr, nil
	}

	if len(arr) > *doc.MaxArrayLen {
		return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("too many items in '%s' array. got: %d, maximum allowed: %d", fieldName, len(arr), *doc.MaxArrayLen), nil)
	}

	return arr, nil
}
