package api

import (
	"errors"
	"fmt"
	"net/http"
)

type verifyFn[T any] func(paramsVerifyer, T, FieldName) (T, error)

type paramsVerifyer struct {
	cfg         *Config
	paramsRules ParamsRules
	userValues  map[FieldName]any
}

func createParamsVerifyer(cfg *Config, paramsRules ParamsRules, userValues map[FieldName]any) paramsVerifyer {
	return paramsVerifyer{
		cfg:         cfg,
		paramsRules: paramsRules,
		userValues:  userValues,
	}
}

func (v paramsVerifyer) vfIntField(val int32, fieldName FieldName) (int32, error) {
	return v.vfFieldWrapper(val, fieldName, vfInt)
}

func (v paramsVerifyer) vfIDField(val int32, fieldName FieldName) (int32, error) {
	return v.vfFieldWrapper(val, fieldName, vfID)
}

func (v paramsVerifyer) vfEnumField(val string, fieldName FieldName) (string, error) {
	return v.vfFieldWrapper(val, fieldName, vfEnum)
}

func (v paramsVerifyer) vfArrayField[T any](arr []T, fieldName FieldName) ([]T, error) {
	return v.vfFieldWrapper(arr, fieldName, vfArray)
}

func (v paramsVerifyer) vfFieldWrapper[T any](val T, fieldName FieldName, verifyFn verifyFn[T]) (T, error) {
	var zero T

	val, err := v.vfField(val, fieldName)
	if err != nil {
		return zero, err
	}

	if verifyFn != nil {
		val, err = verifyFn(v, val, fieldName)
		if err != nil {
			return zero, err
		}
	}

	return val, nil
}

// verifies, if a pointer-field fulfills all its conditions. the basic checks are performed on the pointer itself. after that, the pointer is safely dereferenced and a type-specific verification function can be performed on the pointer's typed value. a pointer will never have a default value since it's optional.
func (v paramsVerifyer) verifyParamFieldPtr[T any](ptr *T, fieldName FieldName, verifyFn verifyFn[T]) (*T, error) {
	ptr, err := v.vfField(ptr, fieldName)
	if err != nil {
		return nil, err
	}

	if ptr == nil || verifyFn == nil {
		return ptr, nil
	}

	*ptr, err = v.vfFieldWrapper(*ptr, fieldName, verifyFn)
	if err != nil {
		return nil, err
	}

	return ptr, nil
}

func (v paramsVerifyer) verifyParamFieldObj[T any](obj T, fieldName FieldName, verifyFn verifyFn[T]) (T, error) {
	var zero T

	obj, err := v.vfField(obj, fieldName)
	if err != nil {
		return zero, err
	}

	if verifyFn == nil {
		return obj, nil
	}

	errStart := fmt.Sprintf("object at field '%s'", fieldName)

	cv, err := v.childObj(fieldName)
	if err != nil {
		return zero, cascadeHTTPError(errStart, err)
	}

	obj, err = cv.vfFieldWrapper(obj, fieldName, verifyFn)
	if err != nil {
		return zero, cascadeHTTPError(errStart, err)
	}

	return obj, nil
}

func (v paramsVerifyer) verifyParamFieldObjPtr[T any](objPtr *T, fieldName FieldName, verifyFn verifyFn[T]) (*T, error) {
	objPtr, err := v.vfField(objPtr, fieldName)
	if err != nil {
		return nil, err
	}

	if objPtr == nil || verifyFn == nil {
		return objPtr, nil
	}

	*objPtr, err = v.verifyParamFieldObj(*objPtr, fieldName, verifyFn)
	if err != nil {
		return nil, err
	}

	return objPtr, nil
}

// verifies, if an array-field fulfills all its conditions. the basic checks as well as a max length check are performed on the array itself. after that, a type-specific verification function can be performed on each of the array's items.
func (v paramsVerifyer) verifyParamFieldArr[T any](arr []T, fieldName FieldName, verifyFn verifyFn[T]) ([]T, error) {
	arr, err := v.vfArrayField(arr, fieldName)
	if errors.Is(err, errNullArray) {
		return arr, nil
	}
	if err != nil {
		return nil, err
	}

	if verifyFn == nil {
		return arr, nil
	}

	valSlice, ok := v.userValues[fieldName].([]any)
	if !ok {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't find value slice of field '%s'", fieldName), nil)
	}

	for i, item := range arr {
		errStart := fmt.Sprintf("index %d of array at field '%s'", i, fieldName)

		cv, err := v.childArr(valSlice[i], fieldName)
		if err != nil {
			return nil, cascadeHTTPError(errStart, err)
		}

		arr[i], err = verifyFn(cv, item, fieldName)
		if err != nil {
			return nil, cascadeHTTPError(errStart, err)
		}
	}

	return arr, nil
}

func (v paramsVerifyer) verifyParamFieldIDArr(arr []int32, fieldName FieldName) ([]int32, error) {
	arr, err := v.vfArrayField(arr, fieldName)
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

		arr[i], err = vfArrayChildID(v, val, fieldName)
		if err != nil {
			return nil, cascadeHTTPError(errStart, err)
		}

		isDupe[val] = true
	}

	return arr, nil
}

func (v paramsVerifyer) childObj(fieldName FieldName) (paramsVerifyer, error) {
	childUserVals, ok := v.userValues[fieldName].(map[FieldName]any)
	if !ok {
		return paramsVerifyer{}, newHTTPError(http.StatusInternalServerError, "couldn't find value map", nil)
	}

	cv := createParamsVerifyer(v.cfg, v.paramsRules[fieldName].Children, childUserVals)

	return cv, nil
}

func (v paramsVerifyer) childArr(rawItem any, fieldName FieldName) (paramsVerifyer, error) {
	childUserVals, ok := rawItem.(map[FieldName]any)
	if !ok {
		return paramsVerifyer{}, newHTTPError(http.StatusInternalServerError, "couldn't find value map", nil)
	}

	cv := createParamsVerifyer(v.cfg, v.paramsRules[fieldName].Children, childUserVals)

	return cv, nil
}
