package main

import (
	"bytes"
	"encoding/json"
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