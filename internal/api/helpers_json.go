package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string, err error) {
	if err != nil {
		log.Println(err)
	}

	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}

	type errorResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, code, errorResponse{
		Error: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}

	dat = bytes.ReplaceAll(dat, []byte(`\u0026`), []byte("&"))
	dat = bytes.ReplaceAll(dat, []byte(`\u003c`), []byte("<"))
	dat = bytes.ReplaceAll(dat, []byte(`\u003e`), []byte(">"))

	w.WriteHeader(code)
	w.Write(dat)
}


func convertStateParam[R any](r *http.Request, queryLookup map[QueryParamName]QueryParam, params R)(R, map[string]any, error) {
	var zero R

	state, err := checkEmptyQuery(r, queryLookup[qpnState])
	if err != nil {
		return zero, nil, err
	}

	return readJsonRequest([]byte(state), params, fmt.Sprintf("invalid or corrupt payload for parameter '%s'", qpnState))
}

func decodeJsonBody[R any](r *http.Request, params R) (R, map[string]any, error) {
	var zero R

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return zero, nil, newHTTPError(http.StatusInternalServerError, "failed to read request body", err)
	}

	return readJsonRequest(bodyBytes, params, "invalid or corrupt payload in request body")
}

func readJsonRequest[R any](data []byte, params R, errMsg string) (R, map[string]any, error) {
	var zero R

	err := json.Unmarshal(data, &params)
	if err != nil {
		return zero, nil, newHTTPError(http.StatusBadRequest, errMsg, err)
	}

	var payloadMap map[string]any
	_ = json.Unmarshal(data, &payloadMap)

	return params, payloadMap, nil
}