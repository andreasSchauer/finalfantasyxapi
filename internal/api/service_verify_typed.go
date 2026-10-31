package api

import (
	"fmt"
	"net/http"
	"slices"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func vfID(v paramsVerifyer, val int32, fieldName FieldName) (int32, error) {
	doc := v.paramsRules[fieldName].Doc

	return v.vfIntRange(val, doc.MinVal, doc.MaxVal, fieldName, "id", false)
}

func vfInt(v paramsVerifyer, val int32, fieldName FieldName) (int32, error) {
	doc := v.paramsRules[fieldName].Doc

	return v.vfIntRange(val, doc.MinVal, doc.MaxVal, fieldName, "value", false)
}

func vfArrayChildID(v paramsVerifyer, val int32, fieldName FieldName) (int32, error) {
	doc := v.paramsRules[fieldName].Doc

	return v.vfIntRange(val, doc.ChildMinVal, doc.ChildMaxVal, fieldName, "id", true)
}

func (v paramsVerifyer) vfIntRange(val int32, minVal, maxVal *int32, fieldName FieldName, intType string, isArray bool) (int32, error) {
	errStart := getFieldRangeErrStart(val, fieldName, intType, isArray)

	if minVal != nil && maxVal != nil && (val < *minVal || val > *maxVal) {
		return 0, newHTTPError(http.StatusBadRequest, fmt.Sprintf("%s allowed range: %d to %d.", errStart, *minVal, *maxVal), nil)
	}

	if minVal != nil && val < *minVal {
		return 0, newHTTPError(http.StatusBadRequest, fmt.Sprintf("%s lowest allowed value: %d.", errStart, *minVal), nil)
	}

	if maxVal != nil && val > *maxVal {
		return 0, newHTTPError(http.StatusBadRequest, fmt.Sprintf("%s highest allowed value: %d.", errStart, *maxVal), nil)
	}

	return val, nil
}

func vfEnum(v paramsVerifyer, typedVal string, fieldName FieldName) (string, error) {
	doc := v.paramsRules[fieldName].Doc

	if doc.EnumValues == nil {
		return typedVal, nil
	}

	if !slices.Contains(doc.EnumValues, typedVal) {
		return "", newHTTPError(http.StatusBadRequest, fmt.Sprintf("value '%s' of field '%s' is not a valid enum value. allowed values: %s.", typedVal, fieldName, h.FormatStringSlice(doc.EnumValues)), nil)
	}

	return typedVal, nil
}

func vfArray[T any](v paramsVerifyer, arr []T, fieldName FieldName) ([]T, error) {
	if arr == nil {
		return arr, errNullArray
	}

	doc := v.paramsRules[fieldName].Doc

	if len(arr) == 0 {
		return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("empty array used for field '%s'. the use of empty arrays is not supported by this API. omit the field, or use 'null' instead.", fieldName), nil)
	}

	if doc.MaxArrayLen == nil {
		return arr, nil
	}

	if len(arr) > *doc.MaxArrayLen {
		return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("too many items in array of field '%s'. got: %d, maximum allowed: %d.", fieldName, len(arr), *doc.MaxArrayLen), nil)
	}

	return arr, nil
}
