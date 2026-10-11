package api

import (
	"net/http"
	"testing"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func TestDropChance(t *testing.T) {
	t.Parallel()
	requestURL := "/api/drop-chance"
	
	tests := []expDropChanceResponse{
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "field 'monster' can't be empty.",
				requestBody: DropChanceParams{
					MinEmptySlots: h.GetInt32Ptr(2),
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "id '10' of field 'character' is out of range. allowed range: 1 to 7.",
				requestBody: DropChanceParams{
					Monster: 13,
					Character: h.GetInt32Ptr(10),
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "id '0' of field 'character' is out of range. allowed range: 1 to 7.",
				requestBody: DropChanceParams{
					Monster: 13,
					Character: h.GetInt32Ptr(0),
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "array of field 'party_members' can't have duplicate ids.",
				requestBody: DropChanceParams{
					Monster: 13,
					PartyMembers: []int32{2,2,2,2,2},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "character must be in the party to be eligible for receiving equipment.",
				requestBody: DropChanceParams{
					Monster: 13,
					Character: h.GetInt32Ptr(4),
					PartyMembers: []int32{1,2,3},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "too many items in 'party_members' array. got: 8, maximum allowed: 7.",
				requestBody: DropChanceParams{
					Monster: 13,
					Character: h.GetInt32Ptr(4),
					PartyMembers: []int32{1,2,3,4,5,6,7,8},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "id '131' inside array of field 'auto_abilities' is out of range. allowed range: 1 to 130.",
				requestBody: DropChanceParams{
					Monster: 13,
					AutoAbilities: []int32{131},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "too many items in 'auto_abilities' array. got: 5, maximum allowed: 4.",
				requestBody: DropChanceParams{
					Monster: 13,
					AutoAbilities: []int32{13,14,15,16,17},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "value '5' of field 'min_empty_slots' is out of range. allowed range: 0 to 4.",
				requestBody: DropChanceParams{
					Monster: 13,
					MinEmptySlots: h.GetInt32Ptr(5),
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "value '5' of field 'total_slots' is out of range. allowed range: 0 to 4.",
				requestBody: DropChanceParams{
					Monster: 13,
					TotalSlots: h.GetInt32Ptr(5),
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "value '-1' of field 'min_empty_slots' is out of range. allowed range: 0 to 4.",
				requestBody: DropChanceParams{
					Monster: 13,
					MinEmptySlots: h.GetInt32Ptr(-1),
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "can't discern equipment's type.",
				requestBody: DropChanceParams{
					Monster: 13,
					MinEmptySlots: h.GetInt32Ptr(2),
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "the amount of auto-abilities and empty slots combined can't exceed the total amount of slots.",
				requestBody: DropChanceParams{
					Monster: 13,
					AutoAbilities: []int32{2},
					TotalSlots: h.GetInt32Ptr(0),
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "the amount of auto-abilities and empty slots combined can't exceed 4.",
				requestBody: DropChanceParams{
					Monster: 13,
					AutoAbilities: []int32{2, 7, 22},
					MinEmptySlots: h.GetInt32Ptr(2),
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "weapon- and armor-abilities can't be combined.",
				requestBody: DropChanceParams{
					Monster: 13,
					AutoAbilities: []int32{2, 129},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "monster 'sinscale - 1 (dream zanarkand)' doesn't drop equipment.",
				requestBody: DropChanceParams{
					Monster: 1,
					AutoAbilities: []int32{2},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "water flan - 1 (tutorial) doesn't drop auto-ability 'zombiestrike' (22). it only drops the following auto-abilities: 'piercing' (2), 'waterstrike' (7), 'distill mana' (26), 'water ward' (61), 'waterproof' (62).",
				requestBody: DropChanceParams{
					Monster: 13,
					AutoAbilities: []int32{2, 7, 22},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "water flan - 1 (tutorial) doesn't drop auto-ability 'piercing' (2) for lulu.",
				requestBody: DropChanceParams{
					Monster: 13,
					Character: h.GetInt32Ptr(4),
					AutoAbilities: []int32{2, 7},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid field: 'text'.",
				requestBody: AlBhedParams{
					Text: "this won't work.",
				},
			},
		},
		/*
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusOK,
				requestBody: DropChanceParams{

				},
			},
			DropChanceResponse: DropChanceResponse{
				
			},
		},	
		*/
	}

	testServiceResponses(t, tests, "DropChance", testCfg.HandleDropChanceGet, testCfg.HandleDropChancePost, compareDropChanceResponses)
}