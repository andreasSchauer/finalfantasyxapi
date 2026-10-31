package api

import (
	"encoding/json"
	"io"
	"net/http"
)

func getParamsJSONBody[P any](r *http.Request) (P, map[FieldName]any, error) {
	var zero P
	var params P

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return zero, nil, newHTTPError(http.StatusInternalServerError, "failed to read request body.", err)
	}

	return ReadJSONRequest(bodyBytes, params, "invalid or corrupt payload in request body.")
}

func ReadJSONRequest[P any](data []byte, params P, errMsg string) (P, map[FieldName]any, error) {
	var zero P

	err := json.Unmarshal(data, &params)
	if err != nil {
		return zero, nil, newHTTPError(http.StatusBadRequest, errMsg, err)
	}

	var rawMap map[string]any
	_ = json.Unmarshal(data, &rawMap)

	userValues := toUserValueMap(rawMap)

	return params, userValues, nil
}
