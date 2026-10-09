package api

import (
	"net/http"
	"testing"
)

func TestDropChance(t *testing.T) {
	t.Parallel()
	tests := []expDropChanceResponse{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/drop-chance",
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "",
				requestBody: DropChanceParams{

				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/drop-chance",
				method: 		http.MethodPost,
				expectedStatus: http.StatusOK,
				requestBody: DropChanceParams{

				},
			},
			DropChanceResponse: DropChanceResponse{
				
			},
		},	
	}

	testServiceResponses(t, tests, "DropChance", testCfg.HandleDropChanceGet, testCfg.HandleDropChancePost, compareDropChanceResponses)
}