package servicetypes

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/errs"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
)

func ReadJSONRequest[P any](data []byte, params P, errMsg string) (P, map[ts.FieldName]any, error) {
	var zero P

	err := json.Unmarshal(data, &params)
	if err != nil {
		return zero, nil, errs.NewHTTPError(http.StatusBadRequest, errMsg, err)
	}

	var rawMap map[string]any
	_ = json.Unmarshal(data, &rawMap)

	userValues := toUserValueMap(rawMap)

	return params, userValues, nil
}

func GetParamsJSONBody[P any](r *http.Request) (P, map[ts.FieldName]any, error) {
	var zero P
	var params P

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return zero, nil, errs.NewHTTPError(http.StatusInternalServerError, "failed to read request body", err)
	}

	return ReadJSONRequest(bodyBytes, params, "invalid or corrupt payload in request body")
}
