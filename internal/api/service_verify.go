package api

// custom logic like the 10 monsters for /turn-order will always be done in the type's own verify function directly



// verifies, if a field fulfills all its conditions. assigns a default value, if the field is empty. if the field is of a certain type, a type-specific verification function can be performed on the typed value.
func verifyParamField[T any](val T, fieldName string, values map[string]any, lookup map[string]FieldDoc, verifyFn func(T, string, FieldDoc) error) (T, error) {
	var zero T
	rules := lookup[fieldName]
	
	val, err := basicFieldChecks(val, fieldName, values, lookup)
	if err != nil {
		return zero, err
	}
	
	if verifyFn != nil {
		err = verifyFn(val, fieldName, rules)
		if err != nil {
			return zero, err
		}
	}

	return val, nil
}

// verifies, if a pointer-field fulfills all its conditions. the basic checks are performed on the pointer itself. after that, the pointer is safely dereferenced and a type-specific verification function can be performed on the pointer's typed value. a pointer will never have a default value since it's optional.
func verifyParamFieldPtr[T any](ptr *T, fieldName string, values map[string]any, lookup map[string]FieldDoc, verifyFn func(T, string, FieldDoc) error) (*T, error) {
	rules := lookup[fieldName]
	
	ptr, err := basicFieldChecks(ptr, fieldName, values, lookup)
	if err != nil {
		return nil, err
	}

	if ptr == nil {
		return ptr, nil
	}
	
	if verifyFn != nil {
		err = verifyFn(*ptr, fieldName, rules)
		if err != nil {
			return nil, err
		}
	}

	return ptr, nil
}

// verifies, if an array-field fulfills all its conditions. the basic checks as well as a max length check are performed on the array itself. after that, a type-specific verification function can be performed on each of the array's items.
func verifyParamFieldArr[T any](arr []T, fieldName string, values map[string]any, lookup map[string]FieldDoc, verifyFn func(T, string, FieldDoc) error) ([]T, error) {
	rules := lookup[fieldName]
	
	arr, err := basicFieldChecks(arr, fieldName, values, lookup)
	if err != nil {
		return nil, err
	}

	err = vfArray(arr, fieldName, rules)
	if err != nil {
		return nil, err
	}
	
	if verifyFn != nil {
		for _, item := range arr {
			err = verifyFn(item, fieldName, rules)
			if err != nil {
				return nil, err
			}
		}
	}

	return arr, nil
}