package api

import (
	"net/http"
	"testing"
)

func TestTurnOrder(t *testing.T) {
	t.Parallel()
	tests := []expTurnOrderResponse{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/turn-order",
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "",
				requestBody: TurnOrderParams{

				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/turn-order",
				method: 		http.MethodPost,
				expectedStatus: http.StatusOK,
				requestBody: TurnOrderParams{

				},
			},
			TurnOrderResponse: TurnOrderResponse{
				
			},
		},	
	}

	testServiceResponses(t, tests, "TurnOrder", testCfg.HandleTurnOrderGet, testCfg.HandleTurnOrderPost, compareTurnOrderResponses)
}