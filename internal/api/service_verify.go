package api

import (
	"fmt"
	"net/http"
)

type verifyFn[T any] func (*Config, T, FieldName, map[FieldName]any, ValidationTree) (T, error)


// verifies, if a field fulfills all its conditions. assigns a default value, if the field is empty. if the field is of a certain type, a type-specific verification function can be performed on the typed value.
func verifyParamField[T any](cfg *Config, val T, fieldName FieldName, valueMap map[FieldName]any, valTree ValidationTree, verifyFn verifyFn[T]) (T, error) {
	var zero T
	
	val, err := basicFieldChecks(val, fieldName, valueMap, valTree)
	if err != nil {
		return zero, err
	}
	
	if verifyFn != nil {
		val, err = verifyFn(cfg, val, fieldName, valueMap, valTree)
		if err != nil {
			return zero, err
		}
	}

	return val, nil
}

// verifies, if a pointer-field fulfills all its conditions. the basic checks are performed on the pointer itself. after that, the pointer is safely dereferenced and a type-specific verification function can be performed on the pointer's typed value. a pointer will never have a default value since it's optional.
func verifyParamFieldPtr[T any](cfg *Config, ptr *T, fieldName FieldName, valueMap map[FieldName]any, valTree ValidationTree, verifyFn verifyFn[T]) (*T, error) {
	ptr, err := basicFieldChecks(ptr, fieldName, valueMap, valTree)
	if err != nil {
		return nil, err
	}

	if ptr == nil || verifyFn == nil {
		return ptr, nil
	}
	
	// do these inputs also work for struct pointers?
	*ptr, err = verifyFn(cfg, *ptr, fieldName, valueMap, valTree)
	if err != nil {
		return nil, err
	}

	return ptr, nil
}

// verifies, if an array-field fulfills all its conditions. the basic checks as well as a max length check are performed on the array itself. after that, a type-specific verification function can be performed on each of the array's items.
func verifyParamFieldArr[T any](cfg *Config, arr []T, fieldName FieldName, valueMap map[FieldName]any, valTree ValidationTree, verifyFn verifyFn[T]) ([]T, error) {	
	arr, err := basicFieldChecks(arr, fieldName, valueMap, valTree)
	if err != nil {
		return nil, err
	}

	if arr == nil {
		return arr, nil
	}

	arr, err = vfArray(cfg, arr, fieldName, valueMap, valTree)
	if err != nil {
		return nil, err
	}

	if verifyFn == nil {
		return arr, nil
	}
	
	valSlice, ok := valueMap[fieldName].([]any)
	if !ok {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't find value slice of field '%s'", fieldName), nil)
	}

	for i, item := range arr {
		rawItem := valSlice[i]
		
		valMap, ok := rawItem.(map[FieldName]any)
		if !ok {
			return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("expected value map at index %d of field '%s'", i, fieldName), nil)
		}

		arr[i], err = verifyParamField(cfg, item, fieldName, valMap, valTree[fieldName].Children, verifyFn)
		if err != nil {
			return nil, err
		}
	}

	return arr, nil
}