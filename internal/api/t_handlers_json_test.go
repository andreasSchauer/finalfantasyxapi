package api

import (
	"net/http"
	"testing"
)

func TestReadJsonRequest_Malformed(t *testing.T) {
	malformedData := []byte(`{ "name": "Bob", broken_json_here `)
	
	type MockPayload struct {
		Name string `json:"name"`
	}

	_, _, err := readJsonRequest(malformedData, MockPayload{}, "custom decode error")
	if err == nil {
		t.Fatal("expected an error for malformed JSON, got nil")
	}

	httpErr, ok := err.(httpError)
	if !ok || httpErr.code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %v", err)
	}
}