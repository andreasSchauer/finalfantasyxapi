package servicetypes

import (
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func toUserValueMap(input map[string]any) map[ts.FieldName]any {
	if input == nil {
		return nil
	}

	output := make(map[ts.FieldName]any, len(input))

	for k, v := range input {
		typedKey := ts.FieldName(k)

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
