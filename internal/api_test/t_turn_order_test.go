package apiTest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func TestTurnOrder(t *testing.T) {
	t.Parallel()
	requestURL := "/api/turn-order"

	tests := []expTurnOrderResponse{
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "at least one of these fields must have a value: 'party', 'mons', 'formation_id'.",
				requestBody:    map[api.FieldName]any{},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "value '0' of field 'turns_amt' is out of range. allowed range: 1 to 1000.",
				requestBody: map[api.FieldName]any{
					api.PfnTurnsAmt: 0,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "id '0' of field 'formation_id' is out of range. allowed range: 1 to 331.",
				requestBody: map[api.FieldName]any{
					api.PfnFormationID: 0,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "empty array used for field 'party'. the use of empty arrays is not supported by this API. omit the field, or use 'null' instead.",
				requestBody: map[api.FieldName]any{
					api.PfnParty: []map[api.FieldName]any{},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "empty array used for field 'mons'. the use of empty arrays is not supported by this API. omit the field, or use 'null' instead.",
				requestBody: map[api.FieldName]any{
					api.PfnMons: []map[api.FieldName]any{},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "too many items in array of field 'party'. got: 4, maximum allowed: 3.",
				requestBody: map[api.FieldName]any{
					api.PfnParty: []map[api.FieldName]any{{}, {}, {}, {}},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "field 'formation_id' can't be used in combination with field 'mons'.",
				requestBody: map[api.FieldName]any{
					api.PfnFormationID: 0,
					api.PfnMons:        []map[api.FieldName]any{{}, {}, {}, {}},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "index 0 of field 'party': id '0' of field 'unit_id' is out of range. allowed range: 1 to 18.",
				requestBody: map[api.FieldName]any{
					api.PfnParty: []map[api.FieldName]any{
						{
							api.PfnUnitID: 0,
						},
					},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "index 1 of field 'party': field 'unit_id' can't be empty.",
				requestBody: map[api.FieldName]any{
					api.PfnParty: []map[api.FieldName]any{
						{
							api.PfnUnitID:  1,
							api.PfnAgility: 0,
						},
						{
							api.PfnAgility: 0,
						},
					},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "index 1 of field 'party': id '0' of field 'unit_id' is out of range. allowed range: 1 to 18.",
				requestBody: map[api.FieldName]any{
					api.PfnParty: []map[api.FieldName]any{
						{
							api.PfnUnitID:  1,
							api.PfnAgility: 0,
						},
						{
							api.PfnUnitID:  0,
							api.PfnAgility: 0,
						},
					},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "index 1 of field 'party': field 'agility' can't be empty.",
				requestBody: map[api.FieldName]any{
					api.PfnParty: []map[api.FieldName]any{
						{
							api.PfnUnitID:  1,
							api.PfnAgility: 255,
						},
						{
							api.PfnUnitID: 2,
						},
					},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "field 'party': duplicate unit id found at index 1. each party member can only appear once.",
				requestBody: map[api.FieldName]any{
					api.PfnParty: []map[api.FieldName]any{
						{
							api.PfnUnitID:  1,
							api.PfnAgility: 255,
						},
						{
							api.PfnUnitID:  1,
							api.PfnAgility: 25,
						},
					},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "index 0 of field 'party': value 'hapfte' of field 'status' is not a valid enum value. allowed values: 'haste', 'slow', 'auto-haste'.",
				requestBody: map[api.FieldName]any{
					api.PfnParty: []map[api.FieldName]any{
						{
							api.PfnUnitID:  1,
							api.PfnAgility: 255,
							api.PfnStatus:  "hapfte",
						},
						{
							api.PfnUnitID:  1,
							api.PfnAgility: 25,
						},
					},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "index 1 of field 'party': invalid field: 'text'.",
				requestBody: map[api.FieldName]any{
					api.PfnParty: []map[api.FieldName]any{
						{
							api.PfnUnitID:  1,
							api.PfnAgility: 255,
						},
						{
							api.PfnText:    1,
							api.PfnAgility: 25,
						},
					},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "index 0 of field 'mons': at least one of these fields must have a value: 'monster_id', 'name'.",
				requestBody: map[api.FieldName]any{
					api.PfnMons: []map[api.FieldName]any{
						{
							api.PfnAgility: 25,
						},
					},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "index 0 of field 'mons': field 'name' requires the following fields to have a value: 'agility'.",
				requestBody: map[api.FieldName]any{
					api.PfnMons: []map[api.FieldName]any{
						{
							api.PfnName: "frank",
						},
					},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "index 1 of field 'mons': field 'name' can't be used in combination with field 'monster_id'.",
				requestBody: map[api.FieldName]any{
					api.PfnMons: []map[api.FieldName]any{
						{
							api.PfnMonsterID: 23,
						},
						{
							api.PfnName:      "frank",
							api.PfnMonsterID: 23,
						},
					},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "index 1 of field 'mons': field 'alt_state' requires the following fields to have a value: 'monster_id'.",
				requestBody: map[api.FieldName]any{
					api.PfnMons: []map[api.FieldName]any{
						{
							api.PfnMonsterID: 23,
						},
						{
							api.PfnName:     "frank",
							api.PfnAltState: 2,
						},
					},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "index 2 of field 'mons': field 'alt_state' can only be used with the following id-value pairs: id: 32, vals: 2, 3 | id: 139, vals: 1, 2, 3 | id: 160, vals: 1 | id: 198, vals: 1 | id: 285, vals: 1 | id: 287, vals: 2.",
				requestBody: map[api.FieldName]any{
					api.PfnMons: []map[api.FieldName]any{
						{
							api.PfnMonsterID: 23,
						},
						{
							api.PfnMonsterID: 15,
						},
						{
							api.PfnMonsterID: 12,
							api.PfnAltState:  2,
						},
						{
							api.PfnMonsterID: 15,
						},
					},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "index 0 of field 'mons': provided value '1' used for 'alt_state' is not available for id '32'. allowed vals: 2, 3.",
				requestBody: map[api.FieldName]any{
					api.PfnMons: []map[api.FieldName]any{
						{
							api.PfnMonsterID: 32,
							api.PfnAltState:  1,
						},
					},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "index 0 of field 'mons': provided value '0' used for 'alt_state' is out of range for id '32'. max val: 3. allowed vals: 2, 3.",
				requestBody: map[api.FieldName]any{
					api.PfnMons: []map[api.FieldName]any{
						{
							api.PfnMonsterID: 32,
							api.PfnAltState:  0,
						},
					},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "index 0 of field 'mons': provided value '4' used for 'alt_state' is out of range for id '32'. max val: 3. allowed vals: 2, 3.",
				requestBody: map[api.FieldName]any{
					api.PfnMons: []map[api.FieldName]any{
						{
							api.PfnMonsterID: 32,
							api.PfnAltState:  4,
						},
					},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "field 'mons': duplicate monster found at index 1. exact duplicate monsters are not allowed.",
				requestBody: map[api.FieldName]any{
					api.PfnMons: []map[api.FieldName]any{
						{
							api.PfnMonsterID: 32,
						},
						{
							api.PfnMonsterID: 32,
						},
					},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "monster 'penance' is immune to 'slow'.",
				requestBody: map[api.FieldName]any{
					api.PfnMons: []map[api.FieldName]any{
						{
							api.PfnMonsterID: 305,
							api.PfnStatus:    string(database.HasteStatusSlow),
						},
					},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"turn order":     30,
					"player party":   1,
					"opponent party": 0,
				},
				requestBody: map[api.FieldName]any{
					api.PfnParty: []map[api.FieldName]any{
						{
							api.PfnUnitID:  1,
							api.PfnAgility: 50,
						},
					},
				},
			},
			TurnOrderResponse: api.TurnOrderResponse{
				TurnsAmt:      30,
				IgnFirstTurn:  false,
				IgnImmunities: false,
				BattleStart:   string(database.BattleStartNormal),
				RNG:           string(database.TurnOrderRngMedian),
				PlayerParty: []api.Participant{
					{
						Name:    "tidus",
						Agility: 50,
						AgilityVals: api.AgilityVals{
							TickSpeed: 6,
							MinICV:    h.GetInt32Ptr(14),
							MaxICV:    h.GetInt32Ptr(18),
						},
						FirstStrike:   false,
						Status:        nil,
						TurnsReceived: 30,
						TurnsPercent:  100,
					},
				},
				OpponentParty: []api.Participant{},
			},
			testTurnOrder: []testBattleTurn{
				{
					index:       0,
					turn:        1,
					currentTick: 16,
					name:        "tidus",
					party:       api.BattlePartyPlayer,
				},
				{
					index:       4,
					turn:        5,
					currentTick: 88,
					name:        "tidus",
					party:       api.BattlePartyPlayer,
				},
				{
					index:       29,
					turn:        30,
					currentTick: 538,
					name:        "tidus",
					party:       api.BattlePartyPlayer,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"turn order":     20,
					"player party":   1,
					"opponent party": 1,
				},
				requestBody: map[api.FieldName]any{
					api.PfnTurnsAmt: 20,
					api.PfnRNG:      string(database.TurnOrderRngWorst),
					api.PfnParty: []map[api.FieldName]any{
						{
							api.PfnUnitID:  1,
							api.PfnAgility: 50,
							api.PfnStatus:  string(database.HasteStatusHaste),
						},
					},
					api.PfnMons: []map[api.FieldName]any{
						{
							api.PfnMonsterID: 27,
							api.PfnStatus:    string(database.HasteStatusSlow),
						},
					},
				},
			},
			TurnOrderResponse: api.TurnOrderResponse{
				TurnsAmt:      20,
				IgnFirstTurn:  false,
				IgnImmunities: false,
				BattleStart:   string(database.BattleStartNormal),
				RNG:           string(database.TurnOrderRngWorst),
				PlayerParty: []api.Participant{
					{
						Name:    "tidus",
						Agility: 50,
						AgilityVals: api.AgilityVals{
							TickSpeed: 3,
							MinICV:    h.GetInt32Ptr(14),
							MaxICV:    h.GetInt32Ptr(18),
						},
						FirstStrike:   false,
						Status:        h.GetStrPtr(string(database.HasteStatusHaste)),
						TurnsReceived: 18,
						TurnsPercent:  90,
					},
				},
				OpponentParty: []api.Participant{
					{
						Name:    "yellow element",
						Agility: 5,
						AgilityVals: api.AgilityVals{
							TickSpeed: 32,
							MinICV:    h.GetInt32Ptr(48),
							MaxICV:    h.GetInt32Ptr(53),
						},
						FirstStrike:   false,
						Status:        h.GetStrPtr(string(database.HasteStatusSlow)),
						TurnsReceived: 2,
						TurnsPercent:  10,
					},
				},
			},
			testTurnOrder: []testBattleTurn{
				{
					index:       0,
					turn:        1,
					currentTick: 18,
					name:        "tidus",
					party:       api.BattlePartyPlayer,
				},
				{
					index:       4,
					turn:        5,
					currentTick: 48,
					name:        "yellow element",
					party:       api.BattlePartyOpponent,
				},
				{
					index:       10,
					turn:        11,
					currentTick: 99,
					name:        "tidus",
					party:       api.BattlePartyPlayer,
				},
				{
					index:       16,
					turn:        17,
					currentTick: 144,
					name:        "yellow element",
					party:       api.BattlePartyOpponent,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"turn order":     10,
					"player party":   1,
					"opponent party": 1,
				},
				requestBody: map[api.FieldName]any{
					api.PfnTurnsAmt:      10,
					api.PfnIgnImmunities: true,
					api.PfnRNG:           string(database.TurnOrderRngBest),
					api.PfnParty: []map[api.FieldName]any{
						{
							api.PfnUnitID:  1,
							api.PfnAgility: 50,
							api.PfnStatus:  string(database.HasteStatusAutoHaste),
						},
					},
					api.PfnMons: []map[api.FieldName]any{
						{
							api.PfnMonsterID: 261,
							api.PfnStatus:    string(database.HasteStatusSlow),
						},
					},
				},
			},
			TurnOrderResponse: api.TurnOrderResponse{
				TurnsAmt:      10,
				IgnFirstTurn:  false,
				IgnImmunities: true,
				BattleStart:   string(database.BattleStartNormal),
				RNG:           string(database.TurnOrderRngBest),
				PlayerParty: []api.Participant{
					{
						Name:    "tidus",
						Agility: 50,
						AgilityVals: api.AgilityVals{
							TickSpeed: 3,
							MinICV:    h.GetInt32Ptr(7),
							MaxICV:    h.GetInt32Ptr(9),
						},
						FirstStrike:   false,
						Status:        h.GetStrPtr(string(database.HasteStatusAutoHaste)),
						TurnsReceived: 8,
						TurnsPercent:  80,
					},
				},
				OpponentParty: []api.Participant{
					{
						Name:    "cactuar king",
						Agility: 80,
						AgilityVals: api.AgilityVals{
							TickSpeed: 10,
							MinICV:    h.GetInt32Ptr(15),
							MaxICV:    h.GetInt32Ptr(16),
						},
						FirstStrike:   false,
						Status:        h.GetStrPtr(string(database.HasteStatusSlow)),
						TurnsReceived: 2,
						TurnsPercent:  20,
					},
				},
			},
			testTurnOrder: []testBattleTurn{
				{
					index:       0,
					turn:        1,
					currentTick: 7,
					name:        "tidus",
					party:       api.BattlePartyPlayer,
				},
				{
					index:       1,
					turn:        2,
					currentTick: 16,
					name:        "cactuar king",
					party:       api.BattlePartyOpponent,
				},
				{
					index:       2,
					turn:        3,
					currentTick: 16,
					name:        "tidus",
					party:       api.BattlePartyPlayer,
				},
				{
					index:       3,
					turn:        4,
					currentTick: 25,
					name:        "tidus",
					party:       api.BattlePartyPlayer,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"turn order":     10,
					"player party":   1,
					"opponent party": 2,
				},
				requestBody: map[api.FieldName]any{
					api.PfnTurnsAmt:      10,
					api.PfnIgnImmunities: true,
					api.PfnRNG:           string(database.TurnOrderRngBest),
					api.PfnParty: []map[api.FieldName]any{
						{
							api.PfnUnitID:      1,
							api.PfnAgility:     50,
							api.PfnFirstStrike: true,
							api.PfnStatus:      string(database.HasteStatusAutoHaste),
						},
					},
					api.PfnMons: []map[api.FieldName]any{
						{
							api.PfnName:        "frank",
							api.PfnAgility:     50,
							api.PfnFirstStrike: true,
							api.PfnStatus:      string(database.HasteStatusSlow),
						},
						{
							api.PfnName:    "frank",
							api.PfnAgility: 80,
							api.PfnStatus:  string(database.HasteStatusHaste),
						},
					},
				},
			},
			TurnOrderResponse: api.TurnOrderResponse{
				TurnsAmt:      10,
				IgnFirstTurn:  false,
				IgnImmunities: true,
				BattleStart:   string(database.BattleStartNormal),
				RNG:           string(database.TurnOrderRngBest),
				PlayerParty: []api.Participant{
					{
						Name:    "tidus",
						Agility: 50,
						AgilityVals: api.AgilityVals{
							TickSpeed: 3,
							MinICV:    h.GetInt32Ptr(0),
							MaxICV:    h.GetInt32Ptr(0),
						},
						FirstStrike:   true,
						Status:        h.GetStrPtr(string(database.HasteStatusAutoHaste)),
						TurnsReceived: 5,
						TurnsPercent:  50,
					},
				},
				OpponentParty: []api.Participant{
					{
						Name:    "frank (1)",
						Agility: 50,
						AgilityVals: api.AgilityVals{
							TickSpeed: 12,
							MinICV:    h.GetInt32Ptr(-1),
							MaxICV:    h.GetInt32Ptr(-1),
						},
						FirstStrike:   true,
						Status:        h.GetStrPtr(string(database.HasteStatusSlow)),
						TurnsReceived: 1,
						TurnsPercent:  10,
					},
					{
						Name:    "frank (2)",
						Agility: 80,
						AgilityVals: api.AgilityVals{
							TickSpeed: 2,
							MinICV:    h.GetInt32Ptr(15),
							MaxICV:    h.GetInt32Ptr(16),
						},
						FirstStrike:   false,
						Status:        h.GetStrPtr(string(database.HasteStatusHaste)),
						TurnsReceived: 4,
						TurnsPercent:  40,
					},
				},
			},
			testTurnOrder: []testBattleTurn{
				{
					index:       0,
					turn:        1,
					currentTick: 0,
					name:        "frank (1)",
					party:       api.BattlePartyOpponent,
				},
				{
					index:       1,
					turn:        2,
					currentTick: 0,
					name:        "tidus",
					party:       api.BattlePartyPlayer,
				},
				{
					index:       2,
					turn:        3,
					currentTick: 9,
					name:        "tidus",
					party:       api.BattlePartyPlayer,
				},
				{
					index:       3,
					turn:        4,
					currentTick: 16,
					name:        "frank (2)",
					party:       api.BattlePartyOpponent,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"turn order":     10,
					"player party":   3,
					"opponent party": 1,
				},
				requestBody: map[api.FieldName]any{
					api.PfnTurnsAmt:    10,
					api.PfnBattleStart: string(database.BattleStartAmbush),
					api.PfnParty: []map[api.FieldName]any{
						{
							api.PfnUnitID:  5,
							api.PfnAgility: 255,
							api.PfnStatus:  string(database.HasteStatusAutoHaste),
						},
						{
							api.PfnUnitID:  1,
							api.PfnAgility: 255,
							api.PfnStatus:  string(database.HasteStatusAutoHaste),
						},
						{
							api.PfnUnitID:  7,
							api.PfnAgility: 255,
							api.PfnStatus:  string(database.HasteStatusAutoHaste),
						},
					},
					api.PfnMons: []map[api.FieldName]any{
						{
							api.PfnName:    "frank",
							api.PfnAgility: 255,
						},
					},
				},
			},
			TurnOrderResponse: api.TurnOrderResponse{
				TurnsAmt:      10,
				IgnFirstTurn:  false,
				IgnImmunities: false,
				BattleStart:   string(database.BattleStartAmbush),
				RNG:           string(database.TurnOrderRngMedian),
				PlayerParty: []api.Participant{
					{
						Name:    "kimahri",
						Agility: 255,
						AgilityVals: api.AgilityVals{
							TickSpeed: 1,
							MinICV:    h.GetInt32Ptr(4),
							MaxICV:    h.GetInt32Ptr(4),
						},
						FirstStrike:   false,
						Status:        h.GetStrPtr(string(database.HasteStatusAutoHaste)),
						TurnsReceived: 3,
						TurnsPercent:  30,
					},
					{
						Name:    "tidus",
						Agility: 255,
						AgilityVals: api.AgilityVals{
							TickSpeed: 1,
							MinICV:    h.GetInt32Ptr(4),
							MaxICV:    h.GetInt32Ptr(4),
						},
						FirstStrike:   false,
						Status:        h.GetStrPtr(string(database.HasteStatusAutoHaste)),
						TurnsReceived: 3,
						TurnsPercent:  30,
					},
					{
						Name:    "rikku",
						Agility: 255,
						AgilityVals: api.AgilityVals{
							TickSpeed: 1,
							MinICV:    h.GetInt32Ptr(4),
							MaxICV:    h.GetInt32Ptr(4),
						},
						FirstStrike:   false,
						Status:        h.GetStrPtr(string(database.HasteStatusAutoHaste)),
						TurnsReceived: 2,
						TurnsPercent:  20,
					},
				},
				OpponentParty: []api.Participant{
					{
						Name:    "frank",
						Agility: 255,
						AgilityVals: api.AgilityVals{
							TickSpeed: 3,
							MinICV:    h.GetInt32Ptr(0),
							MaxICV:    h.GetInt32Ptr(0),
						},
						FirstStrike:   false,
						TurnsReceived: 2,
						TurnsPercent:  20,
					},
				},
			},
			testTurnOrder: []testBattleTurn{
				{
					index:       0,
					turn:        1,
					currentTick: 0,
					name:        "frank",
					party:       api.BattlePartyOpponent,
				},
				{
					index:       1,
					turn:        2,
					currentTick: 4,
					name:        "tidus",
					party:       api.BattlePartyPlayer,
				},
				{
					index:       2,
					turn:        3,
					currentTick: 4,
					name:        "kimahri",
					party:       api.BattlePartyPlayer,
				},
				{
					index:       3,
					turn:        4,
					currentTick: 4,
					name:        "rikku",
					party:       api.BattlePartyPlayer,
				},
				{
					index:       4,
					turn:        5,
					currentTick: 7,
					name:        "tidus",
					party:       api.BattlePartyPlayer,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"turn order":     10,
					"player party":   1,
					"opponent party": 2,
				},
				requestBody: map[api.FieldName]any{
					api.PfnTurnsAmt:    10,
					api.PfnBattleStart: string(database.BattleStartNormal),
					api.PfnFormationID: 88,
					api.PfnParty: []map[api.FieldName]any{
						{
							api.PfnUnitID:  14,
							api.PfnAgility: 10,
						},
					},
				},
			},
			TurnOrderResponse: api.TurnOrderResponse{
				TurnsAmt:      10,
				IgnFirstTurn:  false,
				IgnImmunities: false,
				BattleStart:   string(database.BattleStartNormal),
				RNG:           string(database.TurnOrderRngMedian),
				PlayerParty: []api.Participant{
					{
						Name:    "anima",
						Agility: 10,
						AgilityVals: api.AgilityVals{
							TickSpeed: 14,
							MinICV:    h.GetInt32Ptr(41),
							MaxICV:    h.GetInt32Ptr(42),
						},
						FirstStrike:   false,
						TurnsReceived: 3,
						TurnsPercent:  30,
					},
				},
				OpponentParty: []api.Participant{
					{
						Name:    "basilisk",
						Agility: 9,
						AgilityVals: api.AgilityVals{
							TickSpeed: 15,
							MinICV:    h.GetInt32Ptr(45),
							MaxICV:    h.GetInt32Ptr(50),
						},
						FirstStrike:   false,
						TurnsReceived: 3,
						TurnsPercent:  30,
					},
					{
						Name:    "bite bug",
						Agility: 15,
						AgilityVals: api.AgilityVals{
							TickSpeed: 12,
							MinICV:    h.GetInt32Ptr(36),
							MaxICV:    h.GetInt32Ptr(40),
						},
						FirstStrike:   false,
						TurnsReceived: 4,
						TurnsPercent:  40,
					},
				},
			},
			testTurnOrder: []testBattleTurn{
				{
					index:       0,
					turn:        1,
					currentTick: 38,
					name:        "bite bug",
					party:       api.BattlePartyOpponent,
				},
				{
					index:       1,
					turn:        2,
					currentTick: 41,
					name:        "anima",
					party:       api.BattlePartyPlayer,
				},
				{
					index:       2,
					turn:        3,
					currentTick: 47,
					name:        "basilisk",
					party:       api.BattlePartyOpponent,
				},
				{
					index:       7,
					turn:        8,
					currentTick: 125,
					name:        "anima",
					party:       api.BattlePartyPlayer,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"turn order":     10,
					"player party":   1,
					"opponent party": 2,
				},
				requestBody: map[api.FieldName]any{
					api.PfnTurnsAmt:    10,
					api.PfnBattleStart: string(database.BattleStartPreemptive),
					api.PfnFormationID: 88,
					api.PfnParty: []map[api.FieldName]any{
						{
							api.PfnUnitID:  14,
							api.PfnAgility: 10,
						},
					},
				},
			},
			TurnOrderResponse: api.TurnOrderResponse{
				TurnsAmt:      10,
				IgnFirstTurn:  false,
				IgnImmunities: false,
				BattleStart:   string(database.BattleStartPreemptive),
				RNG:           string(database.TurnOrderRngMedian),
				PlayerParty: []api.Participant{
					{
						Name:    "anima",
						Agility: 10,
						AgilityVals: api.AgilityVals{
							TickSpeed: 14,
							MinICV:    h.GetInt32Ptr(0),
							MaxICV:    h.GetInt32Ptr(0),
						},
						FirstStrike:   false,
						TurnsReceived: 5,
						TurnsPercent:  50,
					},
				},
				OpponentParty: []api.Participant{
					{
						Name:    "basilisk",
						Agility: 9,
						AgilityVals: api.AgilityVals{
							TickSpeed: 15,
							MinICV:    h.GetInt32Ptr(135),
							MaxICV:    h.GetInt32Ptr(135),
						},
						FirstStrike:   false,
						TurnsReceived: 2,
						TurnsPercent:  20,
					},
					{
						Name:    "bite bug",
						Agility: 15,
						AgilityVals: api.AgilityVals{
							TickSpeed: 12,
							MinICV:    h.GetInt32Ptr(108),
							MaxICV:    h.GetInt32Ptr(108),
						},
						FirstStrike:   false,
						TurnsReceived: 3,
						TurnsPercent:  30,
					},
				},
			},
			testTurnOrder: []testBattleTurn{
				{
					index:       0,
					turn:        1,
					currentTick: 0,
					name:        "anima",
					party:       api.BattlePartyPlayer,
				},
				{
					index:       1,
					turn:        2,
					currentTick: 42,
					name:        "anima",
					party:       api.BattlePartyPlayer,
				},
				{
					index:       2,
					turn:        3,
					currentTick: 84,
					name:        "anima",
					party:       api.BattlePartyPlayer,
				},
				{
					index:       3,
					turn:        4,
					currentTick: 108,
					name:        "bite bug",
					party:       api.BattlePartyOpponent,
				},
				{
					index:       5,
					turn:        6,
					currentTick: 135,
					name:        "basilisk",
					party:       api.BattlePartyOpponent,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"turn order":     10,
					"player party":   1,
					"opponent party": 1,
				},
				requestBody: map[api.FieldName]any{
					api.PfnTurnsAmt:      10,
					api.PfnIgnFirstTurn:  true,
					api.PfnIgnImmunities: true,
					api.PfnParty: []map[api.FieldName]any{
						{
							api.PfnUnitID:  2,
							api.PfnAgility: 24,
						},
					},
					api.PfnMons: []map[api.FieldName]any{
						{
							api.PfnMonsterID: 19,
							api.PfnStatus:    string(database.HasteStatusHaste),
						},
					},
				},
			},
			TurnOrderResponse: api.TurnOrderResponse{
				TurnsAmt:      10,
				IgnFirstTurn:  true,
				IgnImmunities: true,
				BattleStart:   string(database.BattleStartNormal),
				RNG:           string(database.TurnOrderRngMedian),
				PlayerParty: []api.Participant{
					{
						Name:    "yuna",
						Agility: 24,
						AgilityVals: api.AgilityVals{
							TickSpeed: 9,
							MinICV:    h.GetInt32Ptr(0),
							MaxICV:    h.GetInt32Ptr(0),
						},
						FirstStrike:   false,
						TurnsReceived: 5,
						TurnsPercent:  50,
					},
				},
				OpponentParty: []api.Participant{
					{
						Name:    "sin - 1",
						Agility: 6,
						AgilityVals: api.AgilityVals{
							TickSpeed: 8,
							MinICV:    h.GetInt32Ptr(0),
							MaxICV:    h.GetInt32Ptr(0),
						},
						FirstStrike:   false,
						Status:        h.GetStrPtr(string(database.HasteStatusHaste)),
						TurnsReceived: 5,
						TurnsPercent:  50,
					},
				},
			},
			testTurnOrder: []testBattleTurn{
				{
					index:       0,
					turn:        1,
					currentTick: 0,
					name:        "yuna",
					party:       api.BattlePartyPlayer,
				},
				{
					index:       1,
					turn:        2,
					currentTick: 0,
					name:        "sin - 1",
					party:       api.BattlePartyOpponent,
				},
				{
					index:       2,
					turn:        3,
					currentTick: 24,
					name:        "sin - 1",
					party:       api.BattlePartyOpponent,
				},
				{
					index:       3,
					turn:        4,
					currentTick: 27,
					name:        "yuna",
					party:       api.BattlePartyPlayer,
				},
				{
					index:       8,
					turn:        9,
					currentTick: 96,
					name:        "sin - 1",
					party:       api.BattlePartyOpponent,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"turn order":     10,
					"player party":   0,
					"opponent party": 3,
				},
				requestBody: map[api.FieldName]any{
					api.PfnTurnsAmt: 10,
					api.PfnMons: []map[api.FieldName]any{
						{
							api.PfnMonsterID: 32,
						},
						{
							api.PfnMonsterID: 32,
							api.PfnAltState:  2,
						},
						{
							api.PfnMonsterID: 32,
							api.PfnAltState:  3,
						},
					},
				},
			},
			TurnOrderResponse: api.TurnOrderResponse{
				TurnsAmt:      10,
				IgnFirstTurn:  false,
				IgnImmunities: false,
				BattleStart:   string(database.BattleStartNormal),
				RNG:           string(database.TurnOrderRngMedian),
				PlayerParty:   []api.Participant{},
				OpponentParty: []api.Participant{
					{
						Name:    "oblitzerator (1)",
						Agility: 30,
						AgilityVals: api.AgilityVals{
							TickSpeed: 8,
							MinICV:    h.GetInt32Ptr(24),
							MaxICV:    h.GetInt32Ptr(26),
						},
						FirstStrike:   false,
						TurnsReceived: 4,
						TurnsPercent:  40,
					},
					{
						Name:    "oblitzerator (2)",
						Agility: 16,
						AgilityVals: api.AgilityVals{
							TickSpeed: 12,
							MinICV:    h.GetInt32Ptr(36),
							MaxICV:    h.GetInt32Ptr(40),
						},
						FirstStrike:   false,
						AltState:      h.GetInt32Ptr(2),
						TurnsReceived: 2,
						TurnsPercent:  20,
					},
					{
						Name:    "oblitzerator (3)",
						Agility: 26,
						AgilityVals: api.AgilityVals{
							TickSpeed: 9,
							MinICV:    h.GetInt32Ptr(27),
							MaxICV:    h.GetInt32Ptr(30),
						},
						FirstStrike:   false,
						AltState:      h.GetInt32Ptr(3),
						TurnsReceived: 4,
						TurnsPercent:  40,
					},
				},
			},
			testTurnOrder: []testBattleTurn{
				{
					index:       0,
					turn:        1,
					currentTick: 25,
					name:        "oblitzerator (1)",
					party:       api.BattlePartyOpponent,
				},
				{
					index:       1,
					turn:        2,
					currentTick: 28,
					name:        "oblitzerator (3)",
					party:       api.BattlePartyOpponent,
				},
				{
					index:       2,
					turn:        3,
					currentTick: 38,
					name:        "oblitzerator (2)",
					party:       api.BattlePartyOpponent,
				},
				{
					index:       3,
					turn:        4,
					currentTick: 49,
					name:        "oblitzerator (1)",
					party:       api.BattlePartyOpponent,
				},
				{
					index:       7,
					turn:        8,
					currentTick: 82,
					name:        "oblitzerator (3)",
					party:       api.BattlePartyOpponent,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"turn order": true,
				},
				expLengths: map[string]int{
					"turn order":     30,
					"player party":   0,
					"opponent party": 1,
				},
				requestBody: map[api.FieldName]any{
					api.PfnMons: []map[api.FieldName]any{
						{
							api.PfnMonsterID:   93,
							api.PfnAgility:     55,
							api.PfnFirstStrike: false,
						},
					},
				},
			},
			TurnOrderResponse: api.TurnOrderResponse{
				TurnsAmt:      30,
				IgnFirstTurn:  false,
				IgnImmunities: false,
				BattleStart:   string(database.BattleStartNormal),
				RNG:           string(database.TurnOrderRngMedian),
				PlayerParty:   []api.Participant{},
				OpponentParty: []api.Participant{
					{
						Name:    "seymour",
						Agility: 55,
						AgilityVals: api.AgilityVals{
							TickSpeed: 6,
							MinICV:    h.GetInt32Ptr(18),
							MaxICV:    h.GetInt32Ptr(20),
						},
						FirstStrike:   false,
						TurnsReceived: 30,
						TurnsPercent:  100,
					},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"turn order": true,
				},
				expLengths: map[string]int{
					"turn order":     30,
					"player party":   0,
					"opponent party": 1,
				},
				requestBody: map[api.FieldName]any{
					api.PfnIgnImmunities: true,
					api.PfnMons: []map[api.FieldName]any{
						{
							api.PfnMonsterID: 306,
							api.PfnStatus:    string(database.HasteStatusSlow),
						},
					},
				},
			},
			TurnOrderResponse: api.TurnOrderResponse{
				TurnsAmt:      30,
				IgnFirstTurn:  false,
				IgnImmunities: true,
				BattleStart:   string(database.BattleStartNormal),
				RNG:           string(database.TurnOrderRngMedian),
				PlayerParty:   []api.Participant{},
				OpponentParty: []api.Participant{
					{
						Name:    "left arm",
						Agility: 255,
						AgilityVals: api.AgilityVals{
							TickSpeed: 6,
							MinICV:    h.GetInt32Ptr(12),
							MaxICV:    h.GetInt32Ptr(13),
						},
						FirstStrike:   false,
						Status:        h.GetStrPtr(string(database.HasteStatusSlow)),
						TurnsReceived: 30,
						TurnsPercent:  100,
					},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"turn order": true,
				},
				expLengths: map[string]int{
					"turn order":     30,
					"player party":   0,
					"opponent party": 2,
				},
				requestBody: map[api.FieldName]any{
					api.PfnFormationID: 226,
				},
			},
			TurnOrderResponse: api.TurnOrderResponse{
				TurnsAmt:      30,
				IgnFirstTurn:  false,
				IgnImmunities: false,
				BattleStart:   string(database.BattleStartAmbush),
				RNG:           string(database.TurnOrderRngMedian),
				PlayerParty:   []api.Participant{},
				OpponentParty: []api.Participant{
					{
						Name:    "dark yojimbo",
						Agility: 243,
						AgilityVals: api.AgilityVals{
							TickSpeed: 3,
							MinICV:    h.GetInt32Ptr(0),
							MaxICV:    h.GetInt32Ptr(0),
						},
						FirstStrike:   false,
						TurnsReceived: 30,
						TurnsPercent:  100,
					},
					{
						Name:    "dark koma inu",
						Agility: 0,
						AgilityVals: api.AgilityVals{
							TickSpeed: 0,
							MinICV:    nil,
							MaxICV:    nil,
						},
						FirstStrike:   false,
						TurnsReceived: 0,
						TurnsPercent:  0,
					},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"turn order":     0,
					"player party":   1,
					"opponent party": 0,
				},
				requestBody: map[api.FieldName]any{
					api.PfnParty: []map[api.FieldName]any{
						{
							api.PfnUnitID:  1,
							api.PfnAgility: 0,
						},
					},
				},
			},
			TurnOrderResponse: api.TurnOrderResponse{
				TurnsAmt:      30,
				IgnFirstTurn:  false,
				IgnImmunities: false,
				BattleStart:   string(database.BattleStartNormal),
				RNG:           string(database.TurnOrderRngMedian),
				PlayerParty: []api.Participant{
					{
						Name:    "tidus",
						Agility: 0,
						AgilityVals: api.AgilityVals{
							TickSpeed: 0,
							MinICV:    nil,
							MaxICV:    nil,
						},
						FirstStrike:   false,
						Status:        nil,
						TurnsReceived: 0,
						TurnsPercent:  0,
					},
				},
				OpponentParty: []api.Participant{},
			},
			testTurnOrder: []testBattleTurn{},
		},
	}

	testServiceResponsesPost(t, tests, "TurnOrder", testCfg.HandleTurnOrderPost, compareTurnOrderResponses)
}
