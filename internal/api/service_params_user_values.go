package api


func toUserValueMap(input map[string]any) map[FieldName]any {
	if input == nil {
		return nil
	}

	output := make(map[FieldName]any, len(input))

	for k, v := range input {
		typedKey := FieldName(k)

		switch value := v.(type) {
		case map[string]any:
			output[typedKey] = toUserValueMap(value)

		case []any:
			output[typedKey] = toUserValueSlice(value)

		default:
			output[typedKey] = value
		}
	}

	return output
}

func toUserValueSlice(input []any) []any {
	output := make([]any, len(input))

	for i, v := range input {
		switch value := v.(type) {
		case map[string]any:
			output[i] = toUserValueMap(value)

		case []any:
			output[i] = toUserValueSlice(value)

		default:
			output[i] = value
		}
	}

	return output
}

func getInt32FromUserVals(fieldName FieldName, userValues map[FieldName]any) (int32, error) {
	raw, ok := userValues[fieldName]
	if !ok {
		return 0, errNoVal
	}
	intVal := int32(raw.(float64))

	return intVal, nil
}