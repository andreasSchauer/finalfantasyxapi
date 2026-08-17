package api

import (
	"bytes"
	"encoding/json"
	"fmt"
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


func convertStateParam[R any](r *http.Request, queryLookup map[QueryParamName]QueryParam, response R)(R, error) {
	var zero R

	state, err := checkEmptyQuery(r, queryLookup[qpnState])
	if err != nil {
		return zero, err
	}

	err = json.Unmarshal([]byte(state), &response)
	if err != nil {
		return zero, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid or corrupt payload for parameter '%s'", qpnState), err)
	}

	return response, nil
}

func decodeJsonBody[R any](r *http.Request, params R) (R, error) {
	var zero R
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&params)
	if err != nil {
		return zero, newHTTPError(http.StatusBadRequest, "invalid or corrupt payload in request body", err)
	}

	return params, nil
}