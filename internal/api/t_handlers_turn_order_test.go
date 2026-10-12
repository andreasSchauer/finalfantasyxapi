package api

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
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
				expectedErr:    "at least one of these fields must have a value: 'party', 'mons', 'formation_id'.",
				requestBody: map[FieldName]any{},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "value '0' of field 'turns_amt' is out of range. allowed range: 1 to 1000.",
				requestBody: map[FieldName]any{
					pfnTurnsAmt: 0,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "id '0' of field 'formation_id' is out of range. allowed range: 1 to 331.",
				requestBody: map[FieldName]any{
					pfnFormationID: 0,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "empty array used for field 'party'. the use of empty arrays is not supported by this API. omit the field, or use 'null' instead.",
				requestBody: map[FieldName]any{
					pfnParty: []map[FieldName]any{},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "empty array used for field 'mons'. the use of empty arrays is not supported by this API. omit the field, or use 'null' instead.",
				requestBody: map[FieldName]any{
					pfnMons: []map[FieldName]any{},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "too many items in array of field 'party'. got: 4, maximum allowed: 3.",
				requestBody: map[FieldName]any{
					pfnParty: []map[FieldName]any{{},{},{},{}},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "field 'formation_id' can't be used in combination with field 'mons'.",
				requestBody: map[FieldName]any{
					pfnFormationID: 0,
					pfnMons: []map[FieldName]any{{},{},{},{}},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "index 0 of field 'party': id '0' of field 'unit_id' is out of range. allowed range: 1 to 18.",
				requestBody: map[FieldName]any{
					pfnParty: []map[FieldName]any{
						{
							pfnUnitID: 0,
						},
					},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "index 1 of field 'party': field 'unit_id' can't be empty.",
				requestBody: map[FieldName]any{
					pfnParty: []map[FieldName]any{
						{
							pfnUnitID: 1,
							pfnAgility: 0,
						},
						{
							pfnAgility: 0,
						},
					},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "index 1 of field 'party': id '0' of field 'unit_id' is out of range. allowed range: 1 to 18.",
				requestBody: map[FieldName]any{
					pfnParty: []map[FieldName]any{
						{
							pfnUnitID: 1,
							pfnAgility: 0,
						},
						{
							pfnUnitID: 0,
							pfnAgility: 0,
						},
					},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "index 1 of field 'party': field 'agility' can't be empty.",
				requestBody: map[FieldName]any{
					pfnParty: []map[FieldName]any{
						{
							pfnUnitID: 1,
							pfnAgility: 255,
						},
						{
							pfnUnitID: 2,
						},
					},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "field 'party': duplicate unit id found at index 1. each party member can only appear once.",
				requestBody: map[FieldName]any{
					pfnParty: []map[FieldName]any{
						{
							pfnUnitID: 1,
							pfnAgility: 255,
						},
						{
							pfnUnitID: 1,
							pfnAgility: 25,
						},
					},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "index 0 of field 'party': value 'hapfte' of field 'status' is not a valid enum value. allowed values: 'haste', 'slow', 'auto-haste'.",
				requestBody: map[FieldName]any{
					pfnParty: []map[FieldName]any{
						{
							pfnUnitID: 1,
							pfnAgility: 255,
							pfnStatus: "hapfte",
						},
						{
							pfnUnitID: 1,
							pfnAgility: 25,
						},
					},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "index 1 of field 'party': invalid field: 'text'.",
				requestBody: map[FieldName]any{
					pfnParty: []map[FieldName]any{
						{
							pfnUnitID: 1,
							pfnAgility: 255,
						},
						{
							pfnText: 1,
							pfnAgility: 25,
						},
					},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "index 0 of field 'mons': at least one of these fields must have a value: 'monster_id', 'name'.",
				requestBody: map[FieldName]any{
					pfnMons: []map[FieldName]any{
						{
							pfnAgility: 25,
						},
					},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "index 0 of field 'mons': field 'name' requires the following fields to have a value: 'agility'.",
				requestBody: map[FieldName]any{
					pfnMons: []map[FieldName]any{
						{
							pfnName: "frank",
						},
					},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "index 1 of field 'mons': field 'name' can't be used in combination with field 'monster_id'.",
				requestBody: map[FieldName]any{
					pfnMons: []map[FieldName]any{
						{
							pfnMonsterID: 23,
						},
						{
							pfnName: "frank",
							pfnMonsterID: 23,
						},
					},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "index 1 of field 'mons': field 'alt_state' requires the following fields to have a value: 'monster_id'.",
				requestBody: map[FieldName]any{
					pfnMons: []map[FieldName]any{
						{
							pfnMonsterID: 23,
						},
						{
							pfnName: "frank",
							pfnAltState: 2,
						},
					},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "index 2 of field 'mons': field 'alt_state' can only be used with the following id-value pairs: id: 32, vals: 2, 3 | id: 139, vals: 1, 2, 3 | id: 160, vals: 1 | id: 198, vals: 1 | id: 285, vals: 1 | id: 287, vals: 2.",
				requestBody: map[FieldName]any{
					pfnMons: []map[FieldName]any{
						{
							pfnMonsterID: 23,
						},
						{
							pfnMonsterID: 15,
						},
						{
							pfnMonsterID: 12,
							pfnAltState: 2,
						},
						{
							pfnMonsterID: 15,
						},
					},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "index 0 of field 'mons': provided value '1' used for 'alt_state' is not available for id '32'. allowed vals: 2, 3.",
				requestBody: map[FieldName]any{
					pfnMons: []map[FieldName]any{
						{
							pfnMonsterID: 32,
							pfnAltState: 1,
						},
					},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "index 0 of field 'mons': provided value '0' used for 'alt_state' is out of range for id '32'. max val: 3. allowed vals: 2, 3.",
				requestBody: map[FieldName]any{
					pfnMons: []map[FieldName]any{
						{
							pfnMonsterID: 32,
							pfnAltState: 0,
						},
					},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "index 0 of field 'mons': provided value '4' used for 'alt_state' is out of range for id '32'. max val: 3. allowed vals: 2, 3.",
				requestBody: map[FieldName]any{
					pfnMons: []map[FieldName]any{
						{
							pfnMonsterID: 32,
							pfnAltState: 4,
						},
					},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "field 'mons': duplicate monster found at index 1. exact duplicate monsters are not allowed.",
				requestBody: map[FieldName]any{
					pfnMons: []map[FieldName]any{
						{
							pfnMonsterID: 32,
						},
						{
							pfnMonsterID: 32,
						},
					},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "monster 'penance' is immune to 'slow'.",
				requestBody: map[FieldName]any{
					pfnMons: []map[FieldName]any{
						{
							pfnMonsterID: 305,
							pfnStatus: string(database.HasteStatusSlow),
						},
					},
				},
			},
		},
		/*
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusOK,
				requestBody: map[FieldName]any{

				},
			},
			TurnOrderResponse: TurnOrderResponse{
				
			},
		},
		*/
	}

	testServiceResponses(t, tests, "TurnOrder", testCfg.HandleTurnOrderGet, testCfg.HandleTurnOrderPost, compareTurnOrderResponses)
}