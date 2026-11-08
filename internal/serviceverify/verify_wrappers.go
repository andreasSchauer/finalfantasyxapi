package serviceverify

import ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"

func (v ParamsVerifyer) fieldWrapper[T any](val T, fieldName ts.FieldName, verifyFn verifyFn[T]) (T, error) {
	var zero T

	val, err := v.field(val, fieldName)
	if err != nil {
		return zero, err
	}

	return verifyFn(v, val, fieldName)
}

// verifies, if a pointer-field fulfills all its conditions. the basic checks are performed on the pointer itself. after that, the pointer is safely dereferenced and a type-specific verification function can be performed on the pointer's typed value. a pointer will never have a default value since it's optional.
func (v ParamsVerifyer) fieldPtrWrapper[T any](ptr *T, fieldName ts.FieldName, verifyFn verifyFn[T]) (*T, error) {
	ptr, err := v.field(ptr, fieldName)
	if err != nil {
		return nil, err
	}

	if ptr == nil || verifyFn == nil {
		return ptr, nil
	}

	*ptr, err = v.fieldWrapper(*ptr, fieldName, verifyFn)
	if err != nil {
		return nil, err
	}

	return ptr, nil
}

func (v ParamsVerifyer) objFieldWrapper[T any](val T, fieldName ts.FieldName, verifyFn verifyFn[T]) (T, error) {
	var zero T

	err := v.CheckFields()
	if err != nil {
		return zero, err
	}

	return v.fieldWrapper(val, fieldName, verifyFn)
}

func (v ParamsVerifyer) arrItemWrapper[T any](val T, fieldName ts.FieldName, verifyFn verifyFn[T]) (T, error) {
	var zero T

	err := v.CheckFields()
	if err != nil {
		return zero, err
	}

	return verifyFn(v, val, fieldName)
}
