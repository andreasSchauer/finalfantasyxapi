package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)


func getParamsStateURL[P any](r *http.Request, queryLookup map[QueryParamName]QueryParam) (P, map[FieldName]any, error) {
	var zero P
	var params P

	state, err := checkEmptyQuery(r, queryLookup[qpnState])
	if err != nil {
		return zero, nil, err
	}

	return readJsonRequest([]byte(state), params, fmt.Sprintf("invalid or corrupt payload for parameter '%s'", qpnState))
}

func getParamsJsonBody[P any](r *http.Request) (P, map[FieldName]any, error) {
	var zero P
	var params P

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return zero, nil, newHTTPError(http.StatusInternalServerError, "failed to read request body", err)
	}

	return readJsonRequest(bodyBytes, params, "invalid or corrupt payload in request body")
}

func readJsonRequest[P any](data []byte, params P, errMsg string) (P, map[FieldName]any, error) {
	var zero P

	err := json.Unmarshal(data, &params)
	if err != nil {
		return zero, nil, newHTTPError(http.StatusBadRequest, errMsg, err)
	}

	var rawMap map[string]any
	_ = json.Unmarshal(data, &rawMap)

	valueMap := toValueMap(rawMap)

	return params, valueMap, nil
}

func toValueMap(input map[string]any) map[FieldName]any {
	if input == nil {
		return nil
	}

	output := make(map[FieldName]any, len(input))

	for k, v := range input {
		typedKey := FieldName(k)

		switch value := v.(type) {
		case map[string]any:
			output[typedKey] = toValueMap(value)

		case []any:
			output[typedKey] = toValueSlice(value)

		default:
			output[typedKey] = value
		}
	}

	return output
}

func toValueSlice(input []any) []any {
	output := make([]any, len(input))

	for i, v := range input {
		switch value := v.(type) {
		case map[string]any:
			output[i] = toValueMap(value)

		case []any:
			output[i] = toValueSlice(value)

		default:
			output[i] = value
		}
	}

	return output
}