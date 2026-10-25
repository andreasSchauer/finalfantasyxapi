package apiTest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api"
)

func TestReadJsonRequest_Malformed(t *testing.T) {
	malformedData := []byte(`{ "name": "Bob", broken_json_here `)

	type MockPayload struct {
		Name string `json:"name"`
	}

	_, _, err := api.ReadJsonRequest(malformedData, MockPayload{}, "custom decode error")
	if err == nil {
		t.Fatal("expected an error for malformed JSON, got nil")
	}

	httpErr, ok := err.(api.HttpError)
	if !ok || httpErr.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %v", err)
	}
}
