package serviceverify

import ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"

func (v ParamsVerifyer) int(val int32, fieldName ts.FieldName) (int32, error) {
	return v.fieldWrapper(val, fieldName, vfInt)
}

func (v ParamsVerifyer) id(val int32, fieldName ts.FieldName) (int32, error) {
	return v.fieldWrapper(val, fieldName, vfID)
}

func (v ParamsVerifyer) string(val string, fieldName ts.FieldName) (string, error) {
	return v.field(val, fieldName)
}

func (v ParamsVerifyer) bool(val bool, fieldName ts.FieldName) (bool, error) {
	return v.field(val, fieldName)
}

func (v ParamsVerifyer) enum(val string, fieldName ts.FieldName) (string, error) {
	return v.fieldWrapper(val, fieldName, vfEnum)
}

func (v ParamsVerifyer) intPtr(val *int32, fieldName ts.FieldName) (*int32, error) {
	return v.fieldPtrWrapper(val, fieldName, vfInt)
}

func (v ParamsVerifyer) idPtr(val *int32, fieldName ts.FieldName) (*int32, error) {
	return v.fieldPtrWrapper(val, fieldName, vfID)
}

func (v ParamsVerifyer) stringPtr(val *string, fieldName ts.FieldName) (*string, error) {
	return v.fieldPtr(val, fieldName)
}

func (v ParamsVerifyer) boolPtr(val *bool, fieldName ts.FieldName) (*bool, error) {
	return v.fieldPtr(val, fieldName)
}

func (v ParamsVerifyer) enumPtr(val *string, fieldName ts.FieldName) (*string, error) {
	return v.fieldPtrWrapper(val, fieldName, vfEnum)
}

func (v ParamsVerifyer) arrayField[T any](arr []T, fieldName ts.FieldName) ([]T, error) {
	return v.fieldWrapper(arr, fieldName, vfArray)
}
