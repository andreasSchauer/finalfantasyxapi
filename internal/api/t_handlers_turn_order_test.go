package api

import (
	"net/http"
	"testing"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func TestTurnOrder(t *testing.T) {
	t.Parallel()
	requestURL := "/api/turn-order"

	tests := []expTurnOrderResponse{
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "at least one of these fields must have a value: 'party', 'mons', 'formation'.",
				requestBody: TurnOrderParams{},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "value '0' of field 'turns_amt' is out of range. allowed range: 1 to 1000.",
				requestBody: TurnOrderParams{
					TurnsAmt: 0,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "id '0' of field 'formation' is out of range. allowed range: 1 to 331.",
				requestBody: TurnOrderParams{
					Formation: h.GetInt32Ptr(0),
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "empty array used for field 'party'. the use of empty arrays is not supported by this API. omit the field, or use null instead.",
				requestBody: TurnOrderParams{
					Party: []turnOrderParty{},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "empty array used for field 'mons'. the use of empty arrays is not supported by this API. omit the field, or use null instead.",
				requestBody: TurnOrderParams{
					Mons: []turnOrderMon{},
				},
			},
		},
		/*
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "",
				requestBody: TurnOrderParams{

				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "",
				requestBody: TurnOrderParams{

				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "",
				requestBody: TurnOrderParams{

				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "",
				requestBody: TurnOrderParams{

				},
			},
		},
		*/
		/*
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusOK,
				requestBody: TurnOrderParams{

				},
			},
			TurnOrderResponse: TurnOrderResponse{
				
			},
		},
		*/
	}

	testServiceResponses(t, tests, "TurnOrder", testCfg.HandleTurnOrderGet, testCfg.HandleTurnOrderPost, compareTurnOrderResponses)
}