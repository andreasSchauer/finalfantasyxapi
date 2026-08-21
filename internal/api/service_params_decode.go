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

	var payloadMap map[FieldName]any
	_ = json.Unmarshal(data, &payloadMap)

	return params, payloadMap, nil
}
