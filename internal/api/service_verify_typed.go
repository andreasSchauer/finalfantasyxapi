package api

func (v paramsVerifyer) int(val int32, fieldName FieldName) (int32, error) {
	return v.fieldWrapper(val, fieldName, vfInt)
}

func (v paramsVerifyer) id(val int32, fieldName FieldName) (int32, error) {
	return v.fieldWrapper(val, fieldName, vfID)
}

func (v paramsVerifyer) enum(val string, fieldName FieldName) (string, error) {
	return v.fieldWrapper(val, fieldName, vfEnum)
}

func (v paramsVerifyer) intPtr(val *int32, fieldName FieldName) (*int32, error) {
	return v.fieldPtrWrapper(val, fieldName, vfInt)
}

func (v paramsVerifyer) idPtr(val *int32, fieldName FieldName) (*int32, error) {
	return v.fieldPtrWrapper(val, fieldName, vfID)
}

func (v paramsVerifyer) enumPtr(val *string, fieldName FieldName) (*string, error) {
	return v.fieldPtrWrapper(val, fieldName, vfEnum)
}

func (v paramsVerifyer) arrayField[T any](arr []T, fieldName FieldName) ([]T, error) {
	return v.fieldWrapper(arr, fieldName, vfArray)
}