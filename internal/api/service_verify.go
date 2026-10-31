package api

import (
	"errors"
	"fmt"
	"net/http"
)

type verifyFn[T any] func(*Config, T, FieldName, map[FieldName]any, ParamsRules) (T, error)

// verifies, if a field fulfills all its conditions. assigns a default value, if the field is empty. if the field is of a certain type, a type-specific verification function can be performed on the typed value.
func verifyParamField[T any](cfg *Config, val T, fieldName FieldName, userValues map[FieldName]any, paramsRules ParamsRules, verifyFn verifyFn[T]) (T, error) {
	var zero T

	val, err := basicFieldChecks(val, fieldName, userValues, paramsRules)
	if err != nil {
		return zero, err
	}

	if verifyFn != nil {
		val, err = verifyFn(cfg, val, fieldName, userValues, paramsRules)
		if err != nil {
			return zero, err
		}
	}

	return val, nil
}

// verifies, if a pointer-field fulfills all its conditions. the basic checks are performed on the pointer itself. after that, the pointer is safely dereferenced and a type-specific verification function can be performed on the pointer's typed value. a pointer will never have a default value since it's optional.
func verifyParamFieldPtr[T any](cfg *Config, ptr *T, fieldName FieldName, userValues map[FieldName]any, paramsRules ParamsRules, verifyFn verifyFn[T]) (*T, error) {
	ptr, err := basicFieldChecks(ptr, fieldName, userValues, paramsRules)
	if err != nil {
		return nil, err
	}

	if ptr == nil || verifyFn == nil {
		return ptr, nil
	}

	*ptr, err = verifyParamField(cfg, *ptr, fieldName, userValues, paramsRules, verifyFn)
	if err != nil {
		return nil, err
	}

	return ptr, nil
}

func verifyParamFieldObj[T any](cfg *Config, obj T, fieldName FieldName, userValues map[FieldName]any, paramsRules ParamsRules, verifyFn verifyFn[T]) (T, error) {
	var zero T

	obj, err := basicFieldChecks(obj, fieldName, userValues, paramsRules)
	if err != nil {
		return zero, err
	}

	if verifyFn == nil {
		return obj, nil
	}

	childValMap, ok := userValues[fieldName].(map[FieldName]any)
	if !ok {
		return zero, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't find value map of field '%s'", fieldName), nil)
	}

	// connect fields with dots somehow
	errStart := fmt.Sprintf("object at field '%s'", fieldName)

	obj, err = verifyParamField(cfg, obj, fieldName, childValMap, paramsRules[fieldName].Children, verifyFn)
	if err != nil {
		return zero, cascadeHTTPError(errStart, err)
	}

	return obj, nil
}

func verifyParamFieldObjPtr[T any](cfg *Config, objPtr *T, fieldName FieldName, userValues map[FieldName]any, paramsRules ParamsRules, verifyFn verifyFn[T]) (*T, error) {
	objPtr, err := basicFieldChecks(objPtr, fieldName, userValues, paramsRules)
	if err != nil {
		return nil, err
	}

	if objPtr == nil || verifyFn == nil {
		return objPtr, nil
	}

	*objPtr, err = verifyParamFieldObj(cfg, *objPtr, fieldName, userValues, paramsRules, verifyFn)
	if err != nil {
		return nil, err
	}

	return objPtr, nil
}

// verifies, if an array-field fulfills all its conditions. the basic checks as well as a max length check are performed on the array itself. after that, a type-specific verification function can be performed on each of the array's items.
func verifyParamFieldArr[T any](cfg *Config, arr []T, fieldName FieldName, userValues map[FieldName]any, paramsRules ParamsRules, verifyFn verifyFn[T]) ([]T, error) {
	arr, err := basicFieldChecks(arr, fieldName, userValues, paramsRules)
	if err != nil {
		return nil, err
	}

	arr, err = vfArray(cfg, arr, fieldName, userValues, paramsRules)
	if errors.Is(err, errNullArray) {
		return arr, nil
	}
	if err != nil {
		return nil, err
	}

	if verifyFn == nil {
		return arr, nil
	}

	valSlice, ok := userValues[fieldName].([]any)
	if !ok {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't find value slice of field '%s'", fieldName), nil)
	}

	for i, item := range arr {
		errStart := fmt.Sprintf("index %d of array at field '%s'", i, fieldName)
		rawItem := valSlice[i]

		childValMap, ok := rawItem.(map[FieldName]any)
		if !ok {
			return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("%s: couldn't find value map", errStart), nil)
		}

		arr[i], err = verifyFn(cfg, item, fieldName, childValMap, paramsRules[fieldName].Children)
		if err != nil {
			return nil, cascadeHTTPError(errStart, err)
		}
	}

	return arr, nil
}

func verifyParamFieldIDArr(cfg *Config, arr []int32, fieldName FieldName, userValues map[FieldName]any, paramsRules ParamsRules) ([]int32, error) {
	arr, err := basicFieldChecks(arr, fieldName, userValues, paramsRules)
	if err != nil {
		return nil, err
	}

	arr, err = vfArray(cfg, arr, fieldName, userValues, paramsRules)
	if errors.Is(err, errNullArray) {
		return arr, nil
	}
	if err != nil {
		return nil, err
	}

	isDupe := make(map[int32]bool)

	for i, val := range arr {
		errStart := fmt.Sprintf("index %d of array at field '%s'", i, fieldName)

		if isDupe[val] {
			return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("%s: array of field '%s' can't have duplicate ids.", errStart, fieldName), nil)
		}

		arr[i], err = vfArrayChildID(nil, val, fieldName, nil, paramsRules)
		if err != nil {
			return nil, cascadeHTTPError(errStart, err)
		}

		isDupe[val] = true
	}

	return arr, nil
}
