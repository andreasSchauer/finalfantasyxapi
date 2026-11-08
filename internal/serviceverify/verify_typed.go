package serviceverify

import ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"

func (v ParamsVerifyer) Int(val int32, fieldName ts.FieldName) (int32, error) {
	return v.fieldWrapper(val, fieldName, vfInt)
}

func (v ParamsVerifyer) ID(val int32, fieldName ts.FieldName) (int32, error) {
	return v.fieldWrapper(val, fieldName, vfID)
}

func (v ParamsVerifyer) String(val string, fieldName ts.FieldName) (string, error) {
	return v.field(val, fieldName)
}

func (v ParamsVerifyer) Bool(val bool, fieldName ts.FieldName) (bool, error) {
	return v.field(val, fieldName)
}

func (v ParamsVerifyer) Enum(val string, fieldName ts.FieldName) (string, error) {
	return v.fieldWrapper(val, fieldName, vfEnum)
}

func (v ParamsVerifyer) IntPtr(val *int32, fieldName ts.FieldName) (*int32, error) {
	return v.fieldPtrWrapper(val, fieldName, vfInt)
}

func (v ParamsVerifyer) IDPtr(val *int32, fieldName ts.FieldName) (*int32, error) {
	return v.fieldPtrWrapper(val, fieldName, vfID)
}

func (v ParamsVerifyer) StringPtr(val *string, fieldName ts.FieldName) (*string, error) {
	return v.fieldPtr(val, fieldName)
}

func (v ParamsVerifyer) BoolPtr(val *bool, fieldName ts.FieldName) (*bool, error) {
	return v.fieldPtr(val, fieldName)
}

func (v ParamsVerifyer) EnumPtr(val *string, fieldName ts.FieldName) (*string, error) {
	return v.fieldPtrWrapper(val, fieldName, vfEnum)
}

func (v ParamsVerifyer) arrayField[T any](arr []T, fieldName ts.FieldName) ([]T, error) {
	return v.fieldWrapper(arr, fieldName, vfArray)
}
