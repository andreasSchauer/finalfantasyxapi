// Package serviceverify adapsoadsokp
package serviceverify

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/service/servicetypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
)

type ParamsVerifyer struct {
	paramsRules servicetypes.ParamsRules
	userValues  map[ts.FieldName]any
}

func NewParamsVerifyer(paramsRules servicetypes.ParamsRules, userValues map[ts.FieldName]any) ParamsVerifyer {
	return ParamsVerifyer{
		paramsRules: paramsRules,
		userValues:  userValues,
	}
}

func (v ParamsVerifyer) hasVal(fieldName ts.FieldName) bool {
	jsonVal, ok := v.userValues[fieldName]
	return ok && jsonVal != nil
}

func (v ParamsVerifyer) hasNoVal(fieldName ts.FieldName) bool {
	return !v.hasVal(fieldName)
}

func (v ParamsVerifyer) CheckFields() error {
	for fieldName := range v.userValues {
		_, exists := v.paramsRules[fieldName]
		if !exists {
			return errs.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid field: '%s'", fieldName), nil)
		}
	}

	return nil
}

func (v ParamsVerifyer) newFieldVerifyer[T any](val T, fieldName ts.FieldName) fieldVerifyer[T] {
	return fieldVerifyer[T]{
		fieldName:    fieldName,
		val:          val,
		doc:          v.paramsRules[fieldName].Doc,
		valIsPresent: v.hasVal(fieldName),
		paramsRules:  v.paramsRules,
		userValues:   v.userValues,
	}
}

func (v ParamsVerifyer) field[T any](val T, fieldName ts.FieldName) (T, error) {
	var zero T
	fv := v.newFieldVerifyer(val, fieldName)
	fns := []func() error{
		fv.required,
		fv.conflictsWith,
		fv.requiredOr,
		fv.requiresAll,
		fv.requiresOne,
		fv.allowedIDs,
		fv.allowedIDValPairs,
	}

	for _, fn := range fns {
		err := fn()
		if err != nil {
			return zero, err
		}
	}

	return fv.assignDefaultVal(), nil
}

func (v ParamsVerifyer) fieldPtr[T any](ptr *T, fieldName ts.FieldName) (*T, error) {
	ptr, err := v.field(ptr, fieldName)
	if err != nil {
		return nil, err
	}

	if ptr == nil {
		return ptr, nil
	}

	*ptr, err = v.field(*ptr, fieldName)
	if err != nil {
		return nil, err
	}

	return ptr, nil
}

func (v ParamsVerifyer) obj[T any](obj T, fieldName ts.FieldName, verifyFn verifyFn[T]) (T, error) {
	var zero T

	obj, err := v.field(obj, fieldName)
	if err != nil {
		return zero, err
	}

	if verifyFn == nil {
		return obj, nil
	}

	errStart := fmt.Sprintf("object at field '%s'", fieldName)

	cv, err := v.objChildVerifyer(fieldName)
	if err != nil {
		return zero, errs.CascadeHTTPError(errStart, err)
	}

	obj, err = cv.objFieldWrapper(obj, fieldName, verifyFn)
	if err != nil {
		return zero, errs.CascadeHTTPError(errStart, err)
	}

	return obj, nil
}

func (v ParamsVerifyer) objPtr[T any](objPtr *T, fieldName ts.FieldName, verifyFn verifyFn[T]) (*T, error) {
	objPtr, err := v.field(objPtr, fieldName)
	if err != nil {
		return nil, err
	}

	if objPtr == nil || verifyFn == nil {
		return objPtr, nil
	}

	*objPtr, err = v.obj(*objPtr, fieldName, verifyFn)
	if err != nil {
		return nil, err
	}

	return objPtr, nil
}

// array verifies, if an array-field fulfills all its conditions. the basic checks as well as a max length check are performed on the array itself. after that, a type-specific verification function can be performed on each of the array's items.
func (v ParamsVerifyer) array[T any](arr []T, fieldName ts.FieldName, verifyFn verifyFn[T]) ([]T, error) {
	arr, err := v.arrayField(arr, fieldName)
	if errors.Is(err, errs.ErrNullArray) {
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
		return nil, errs.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't find value slice of field '%s'", fieldName), nil)
	}

	for i, item := range arr {
		errStart := fmt.Sprintf("index %d of array at field '%s'", i, fieldName)

		cv, err := v.arrChildVerifyer(valSlice[i], fieldName)
		if err != nil {
			return nil, errs.CascadeHTTPError(errStart, err)
		}

		arr[i], err = cv.arrItemWrapper(item, fieldName, verifyFn)
		if err != nil {
			return nil, errs.CascadeHTTPError(errStart, err)
		}
	}

	return arr, nil
}

func (v ParamsVerifyer) idArray(arr []int32, fieldName ts.FieldName) ([]int32, error) {
	arr, err := v.arrayField(arr, fieldName)
	if errors.Is(err, errs.ErrNullArray) {
		return arr, nil
	}
	if err != nil {
		return nil, err
	}

	isDupe := make(map[int32]bool)

	for i, val := range arr {
		errStart := fmt.Sprintf("index %d of array at field '%s'", i, fieldName)

		if isDupe[val] {
			return nil, errs.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%s: array of field '%s' can't have duplicate ids", errStart, fieldName), nil)
		}

		arr[i], err = vfArrayChildID(v, val, fieldName)
		if err != nil {
			return nil, errs.CascadeHTTPError(errStart, err)
		}

		isDupe[val] = true
	}

	return arr, nil
}

func (v ParamsVerifyer) objChildVerifyer(fieldName ts.FieldName) (ParamsVerifyer, error) {
	childUserVals, ok := v.userValues[fieldName].(map[ts.FieldName]any)
	if !ok {
		return ParamsVerifyer{}, errs.NewHTTPError(http.StatusInternalServerError, "couldn't find value map", nil)
	}

	cv := NewParamsVerifyer(v.paramsRules[fieldName].Children, childUserVals)

	return cv, nil
}

func (v ParamsVerifyer) arrChildVerifyer(rawItem any, fieldName ts.FieldName) (ParamsVerifyer, error) {
	childUserVals, ok := rawItem.(map[ts.FieldName]any)
	if !ok {
		return ParamsVerifyer{}, errs.NewHTTPError(http.StatusInternalServerError, "couldn't find value map", nil)
	}

	cv := NewParamsVerifyer(v.paramsRules[fieldName].Children, childUserVals)

	return cv, nil
}
