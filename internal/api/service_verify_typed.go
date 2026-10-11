package api

import (
	"fmt"
	"net/http"
	"slices"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func vfId(_ *Config, val int32, fieldName FieldName, _ map[FieldName]any, valTree ValidationTree) (int32, error) {
	doc := valTree[fieldName].Doc

	return vfIntRange(val, doc.MinVal, doc.MaxVal, fieldName, "id", false)
}

func vfInt(_ *Config, val int32, fieldName FieldName, _ map[FieldName]any, valTree ValidationTree) (int32, error) {
	doc := valTree[fieldName].Doc

	return vfIntRange(val, doc.MinVal, doc.MaxVal, fieldName, "value", false)
}

func vfArrayChildId(_ *Config, val int32, fieldName FieldName, _ map[FieldName]any, valTree ValidationTree) (int32, error) {
	doc := valTree[fieldName].Doc

	return vfIntRange(val, doc.ChildMinVal, doc.ChildMaxVal, fieldName, "id", true)
}

func vfIntRange(val int32, minVal, maxVal *int32, fieldName FieldName, intType string, isArray bool) (int32, error) {
	errStart := getRangeErrStart(val, fieldName, intType, isArray)

	if minVal != nil && maxVal != nil && (val < *minVal || val > *maxVal) {
		return 0, newHTTPError(http.StatusBadRequest, fmt.Sprintf("%s. allowed range: %d to %d.", errStart, *minVal, *maxVal), nil)
	}

	if minVal != nil && val < *minVal {
		return 0, newHTTPError(http.StatusBadRequest, fmt.Sprintf("%s. lowest allowed value: %d.", errStart, *minVal), nil)
	}

	if maxVal != nil && val > *maxVal {
		return 0, newHTTPError(http.StatusBadRequest, fmt.Sprintf("%s. highest allowed value: %d.", errStart, *maxVal), nil)
	}

	return val, nil
}

func getRangeErrStart(val int32, fieldName FieldName, intType string, isArray bool) string {
	errStart := fmt.Sprintf("%s '%d'", intType, val)

	if isArray {
		errStart += " inside array"
	}

	errStart += fmt.Sprintf(" of field '%s' is out of range", fieldName)

	return errStart
}

func vfEnum(_ *Config, typedVal string, fieldName FieldName, _ map[FieldName]any, valTree ValidationTree) (string, error) {
	doc := valTree[fieldName].Doc

	if doc.EnumValues == nil {
		return typedVal, nil
	}

	if !slices.Contains(doc.EnumValues, typedVal) {
		return "", newHTTPError(http.StatusBadRequest, fmt.Sprintf("value '%s' of field '%s' is not a valid enum value. allowed values: %s.", typedVal, fieldName, h.FormatStringSlice(doc.EnumValues)), nil)
	}

	return typedVal, nil
}

func vfArray[T any](_ *Config, arr []T, fieldName FieldName, _ map[FieldName]any, valTree ValidationTree) ([]T, error) {
	if arr == nil {
		return arr, errNullArray
	}

	doc := valTree[fieldName].Doc

	if len(arr) == 0 {
		return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("empty array used for field '%s'. the use of empty arrays is not supported by this API. omit the field, or use null instead.", fieldName), nil)
	}

	if doc.MaxArrayLen == nil {
		return arr, nil
	}

	if len(arr) > *doc.MaxArrayLen {
		return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("too many items in '%s' array. got: %d, maximum allowed: %d.", fieldName, len(arr), *doc.MaxArrayLen), nil)
	}

	return arr, nil
}
