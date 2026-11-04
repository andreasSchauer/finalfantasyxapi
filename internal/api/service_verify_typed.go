package api

import ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"

func (v paramsVerifyer) int(val int32, fieldName ts.FieldName) (int32, error) {
	return v.fieldWrapper(val, fieldName, vfInt)
}

func (v paramsVerifyer) id(val int32, fieldName ts.FieldName) (int32, error) {
	return v.fieldWrapper(val, fieldName, vfID)
}

func (v paramsVerifyer) enum(val string, fieldName ts.FieldName) (string, error) {
	return v.fieldWrapper(val, fieldName, vfEnum)
}

func (v paramsVerifyer) intPtr(val *int32, fieldName ts.FieldName) (*int32, error) {
	return v.fieldPtrWrapper(val, fieldName, vfInt)
}

func (v paramsVerifyer) idPtr(val *int32, fieldName ts.FieldName) (*int32, error) {
	return v.fieldPtrWrapper(val, fieldName, vfID)
}

func (v paramsVerifyer) enumPtr(val *string, fieldName ts.FieldName) (*string, error) {
	return v.fieldPtrWrapper(val, fieldName, vfEnum)
}

func (v paramsVerifyer) arrayField[T any](arr []T, fieldName ts.FieldName) ([]T, error) {
	return v.fieldWrapper(arr, fieldName, vfArray)
}
