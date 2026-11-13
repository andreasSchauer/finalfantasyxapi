package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/service/servicetypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"

	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
)

func TestTurnOrder(t *testing.T) {
	t.Parallel()
	requestURL := "/api/turn-order"

	tests := []TurnOrderResponse{
		{
			TestGeneral: TestGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "at least one of these fields must have a value: 'party', 'mons', 'formation_id'",
				requestBody:    map[ts.FieldName]any{},
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "value '0' of field 'turns_amt' is out of range. allowed range: 1 to 1000",
				requestBody: map[ts.FieldName]any{
					ts.PfnTurnsAmt: 0,
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "id '0' of field 'formation_id' is out of range. allowed range: 1 to 331",
				requestBody: map[ts.FieldName]any{
					ts.PfnFormationID: 0,
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "empty array used for field 'party'. the use of empty arrays is not supported by this api. omit the field, or use 'null' instead",
				requestBody: map[ts.FieldName]any{
					ts.PfnParty: []map[ts.FieldName]any{},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "empty array used for field 'mons'. the use of empty arrays is not supported by this api. omit the field, or use 'null' instead",
				requestBody: map[ts.FieldName]any{
					ts.PfnMons: []map[ts.FieldName]any{},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "too many items in array of field 'party'. got: 4, maximum allowed: 3",
				requestBody: map[ts.FieldName]any{
					ts.PfnParty: []map[ts.FieldName]any{{}, {}, {}, {}},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "field 'formation_id' can't be used in combination with field 'mons'",
				requestBody: map[ts.FieldName]any{
					ts.PfnFormationID: 0,
					ts.PfnMons:        []map[ts.FieldName]any{{}, {}, {}, {}},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "index 0 of array at field 'party': id '0' of field 'unit_id' is out of range. allowed range: 1 to 18",
				requestBody: map[ts.FieldName]any{
					ts.PfnParty: []map[ts.FieldName]any{
						{
							ts.PfnUnitID: 0,
						},
					},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "index 1 of array at field 'party': field 'unit_id' can't be empty",
				requestBody: map[ts.FieldName]any{
					ts.PfnParty: []map[ts.FieldName]any{
						{
							ts.PfnUnitID:  1,
							ts.PfnAgility: 0,
						},
						{
							ts.PfnAgility: 0,
						},
					},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "index 1 of array at field 'party': id '0' of field 'unit_id' is out of range. allowed range: 1 to 18",
				requestBody: map[ts.FieldName]any{
					ts.PfnParty: []map[ts.FieldName]any{
						{
							ts.PfnUnitID:  1,
							ts.PfnAgility: 0,
						},
						{
							ts.PfnUnitID:  0,
							ts.PfnAgility: 0,
						},
					},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "index 1 of array at field 'party': field 'agility' can't be empty",
				requestBody: map[ts.FieldName]any{
					ts.PfnParty: []map[ts.FieldName]any{
						{
							ts.PfnUnitID:  1,
							ts.PfnAgility: 255,
						},
						{
							ts.PfnUnitID: 2,
						},
					},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "field 'party': duplicate unit id found at index 1. each party member can only appear once",
				requestBody: map[ts.FieldName]any{
					ts.PfnParty: []map[ts.FieldName]any{
						{
							ts.PfnUnitID:  1,
							ts.PfnAgility: 255,
						},
						{
							ts.PfnUnitID:  1,
							ts.PfnAgility: 25,
						},
					},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "index 0 of array at field 'party': value 'hapfte' of field 'status' is not a valid enum value. allowed values: 'haste', 'slow', 'auto-haste'",
				requestBody: map[ts.FieldName]any{
					ts.PfnParty: []map[ts.FieldName]any{
						{
							ts.PfnUnitID:  1,
							ts.PfnAgility: 255,
							ts.PfnStatus:  "hapfte",
						},
						{
							ts.PfnUnitID:  1,
							ts.PfnAgility: 25,
						},
					},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "index 1 of array at field 'party': invalid field: 'text'",
				requestBody: map[ts.FieldName]any{
					ts.PfnParty: []map[ts.FieldName]any{
						{
							ts.PfnUnitID:  1,
							ts.PfnAgility: 255,
						},
						{
							ts.PfnText:    1,
							ts.PfnAgility: 25,
						},
					},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "index 0 of array at field 'mons': at least one of these fields must have a value: 'monster_id', 'name'",
				requestBody: map[ts.FieldName]any{
					ts.PfnMons: []map[ts.FieldName]any{
						{
							ts.PfnAgility: 25,
						},
					},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "index 0 of array at field 'mons': field 'name' requires the following fields to have a value: 'agility'",
				requestBody: map[ts.FieldName]any{
					ts.PfnMons: []map[ts.FieldName]any{
						{
							ts.PfnName: "frank",
						},
					},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "index 1 of array at field 'mons': field 'name' can't be used in combination with field 'monster_id'",
				requestBody: map[ts.FieldName]any{
					ts.PfnMons: []map[ts.FieldName]any{
						{
							ts.PfnMonsterID: 23,
						},
						{
							ts.PfnName:      "frank",
							ts.PfnMonsterID: 23,
						},
					},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "index 1 of array at field 'mons': field 'alt_state' requires the following fields to have a value: 'monster_id'",
				requestBody: map[ts.FieldName]any{
					ts.PfnMons: []map[ts.FieldName]any{
						{
							ts.PfnMonsterID: 23,
						},
						{
							ts.PfnName:     "frank",
							ts.PfnAltState: 2,
						},
					},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "index 2 of array at field 'mons': field 'alt_state' can only be used with the following id-value pairs: id: 32, vals: 2, 3 | id: 139, vals: 1, 2, 3 | id: 160, vals: 1 | id: 198, vals: 1 | id: 285, vals: 1 | id: 287, vals: 2",
				requestBody: map[ts.FieldName]any{
					ts.PfnMons: []map[ts.FieldName]any{
						{
							ts.PfnMonsterID: 23,
						},
						{
							ts.PfnMonsterID: 15,
						},
						{
							ts.PfnMonsterID: 12,
							ts.PfnAltState:  2,
						},
						{
							ts.PfnMonsterID: 15,
						},
					},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "index 0 of array at field 'mons': provided value '1' used for 'alt_state' is not available for id '32'. allowed vals: 2, 3",
				requestBody: map[ts.FieldName]any{
					ts.PfnMons: []map[ts.FieldName]any{
						{
							ts.PfnMonsterID: 32,
							ts.PfnAltState:  1,
						},
					},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "index 0 of array at field 'mons': provided value '0' used for 'alt_state' is out of range for id '32'. max val: 3. allowed vals: 2, 3",
				requestBody: map[ts.FieldName]any{
					ts.PfnMons: []map[ts.FieldName]any{
						{
							ts.PfnMonsterID: 32,
							ts.PfnAltState:  0,
						},
					},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "index 0 of array at field 'mons': provided value '4' used for 'alt_state' is out of range for id '32'. max val: 3. allowed vals: 2, 3",
				requestBody: map[ts.FieldName]any{
					ts.PfnMons: []map[ts.FieldName]any{
						{
							ts.PfnMonsterID: 32,
							ts.PfnAltState:  4,
						},
					},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "field 'mons': duplicate monster found at index 1. exact duplicate monsters are not allowed",
				requestBody: map[ts.FieldName]any{
					ts.PfnMons: []map[ts.FieldName]any{
						{
							ts.PfnMonsterID: 32,
						},
						{
							ts.PfnMonsterID: 32,
						},
					},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "monster 'penance' is immune to 'slow'",
				requestBody: map[ts.FieldName]any{
					ts.PfnMons: []map[ts.FieldName]any{
						{
							ts.PfnMonsterID: 305,
							ts.PfnStatus:    string(database.HasteStatusSlow),
						},
					},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"turn order":     30,
					"player party":   1,
					"opponent party": 0,
				},
				requestBody: map[ts.FieldName]any{
					ts.PfnParty: []map[ts.FieldName]any{
						{
							ts.PfnUnitID:  1,
							ts.PfnAgility: 50,
						},
					},
				},
			},
			TurnOrderResponse: servicetypes.TurnOrderResponse{
				TurnsAmt:      30,
				IgnFirstTurn:  false,
				IgnImmunities: false,
				BattleStart:   string(database.BattleStartNormal),
				RNG:           string(database.TurnOrderRngMedian),
				PlayerParty: []servicetypes.Participant{
					{
						Name:    "tidus",
						Agility: 50,
						AgilityVals: servicetypes.AgilityVals{
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
				OpponentParty: []servicetypes.Participant{},
			},
			testTurnOrder: []TestBattleTurn{
				{
					index:       0,
					turn:        1,
					currentTick: 16,
					name:        "tidus",
					party:       servicetypes.BattlePartyPlayer,
				},
				{
					index:       4,
					turn:        5,
					currentTick: 88,
					name:        "tidus",
					party:       servicetypes.BattlePartyPlayer,
				},
				{
					index:       29,
					turn:        30,
					currentTick: 538,
					name:        "tidus",
					party:       servicetypes.BattlePartyPlayer,
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"turn order":     20,
					"player party":   1,
					"opponent party": 1,
				},
				requestBody: map[ts.FieldName]any{
					ts.PfnTurnsAmt: 20,
					ts.PfnRNG:      string(database.TurnOrderRngWorst),
					ts.PfnParty: []map[ts.FieldName]any{
						{
							ts.PfnUnitID:  1,
							ts.PfnAgility: 50,
							ts.PfnStatus:  string(database.HasteStatusHaste),
						},
					},
					ts.PfnMons: []map[ts.FieldName]any{
						{
							ts.PfnMonsterID: 27,
							ts.PfnStatus:    string(database.HasteStatusSlow),
						},
					},
				},
			},
			TurnOrderResponse: servicetypes.TurnOrderResponse{
				TurnsAmt:      20,
				IgnFirstTurn:  false,
				IgnImmunities: false,
				BattleStart:   string(database.BattleStartNormal),
				RNG:           string(database.TurnOrderRngWorst),
				PlayerParty: []servicetypes.Participant{
					{
						Name:    "tidus",
						Agility: 50,
						AgilityVals: servicetypes.AgilityVals{
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
				OpponentParty: []servicetypes.Participant{
					{
						Name:    "yellow element",
						Agility: 5,
						AgilityVals: servicetypes.AgilityVals{
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
			testTurnOrder: []TestBattleTurn{
				{
					index:       0,
					turn:        1,
					currentTick: 18,
					name:        "tidus",
					party:       servicetypes.BattlePartyPlayer,
				},
				{
					index:       4,
					turn:        5,
					currentTick: 48,
					name:        "yellow element",
					party:       servicetypes.BattlePartyOpponent,
				},
				{
					index:       10,
					turn:        11,
					currentTick: 99,
					name:        "tidus",
					party:       servicetypes.BattlePartyPlayer,
				},
				{
					index:       16,
					turn:        17,
					currentTick: 144,
					name:        "yellow element",
					party:       servicetypes.BattlePartyOpponent,
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"turn order":     10,
					"player party":   1,
					"opponent party": 1,
				},
				requestBody: map[ts.FieldName]any{
					ts.PfnTurnsAmt:      10,
					ts.PfnIgnImmunities: true,
					ts.PfnRNG:           string(database.TurnOrderRngBest),
					ts.PfnParty: []map[ts.FieldName]any{
						{
							ts.PfnUnitID:  1,
							ts.PfnAgility: 50,
							ts.PfnStatus:  string(database.HasteStatusAutoHaste),
						},
					},
					ts.PfnMons: []map[ts.FieldName]any{
						{
							ts.PfnMonsterID: 261,
							ts.PfnStatus:    string(database.HasteStatusSlow),
						},
					},
				},
			},
			TurnOrderResponse: servicetypes.TurnOrderResponse{
				TurnsAmt:      10,
				IgnFirstTurn:  false,
				IgnImmunities: true,
				BattleStart:   string(database.BattleStartNormal),
				RNG:           string(database.TurnOrderRngBest),
				PlayerParty: []servicetypes.Participant{
					{
						Name:    "tidus",
						Agility: 50,
						AgilityVals: servicetypes.AgilityVals{
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
				OpponentParty: []servicetypes.Participant{
					{
						Name:    "cactuar king",
						Agility: 80,
						AgilityVals: servicetypes.AgilityVals{
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
			testTurnOrder: []TestBattleTurn{
				{
					index:       0,
					turn:        1,
					currentTick: 7,
					name:        "tidus",
					party:       servicetypes.BattlePartyPlayer,
				},
				{
					index:       1,
					turn:        2,
					currentTick: 16,
					name:        "cactuar king",
					party:       servicetypes.BattlePartyOpponent,
				},
				{
					index:       2,
					turn:        3,
					currentTick: 16,
					name:        "tidus",
					party:       servicetypes.BattlePartyPlayer,
				},
				{
					index:       3,
					turn:        4,
					currentTick: 25,
					name:        "tidus",
					party:       servicetypes.BattlePartyPlayer,
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"turn order":     10,
					"player party":   1,
					"opponent party": 2,
				},
				requestBody: map[ts.FieldName]any{
					ts.PfnTurnsAmt:      10,
					ts.PfnIgnImmunities: true,
					ts.PfnRNG:           string(database.TurnOrderRngBest),
					ts.PfnParty: []map[ts.FieldName]any{
						{
							ts.PfnUnitID:      1,
							ts.PfnAgility:     50,
							ts.PfnFirstStrike: true,
							ts.PfnStatus:      string(database.HasteStatusAutoHaste),
						},
					},
					ts.PfnMons: []map[ts.FieldName]any{
						{
							ts.PfnName:        "frank",
							ts.PfnAgility:     50,
							ts.PfnFirstStrike: true,
							ts.PfnStatus:      string(database.HasteStatusSlow),
						},
						{
							ts.PfnName:    "frank",
							ts.PfnAgility: 80,
							ts.PfnStatus:  string(database.HasteStatusHaste),
						},
					},
				},
			},
			TurnOrderResponse: servicetypes.TurnOrderResponse{
				TurnsAmt:      10,
				IgnFirstTurn:  false,
				IgnImmunities: true,
				BattleStart:   string(database.BattleStartNormal),
				RNG:           string(database.TurnOrderRngBest),
				PlayerParty: []servicetypes.Participant{
					{
						Name:    "tidus",
						Agility: 50,
						AgilityVals: servicetypes.AgilityVals{
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
				OpponentParty: []servicetypes.Participant{
					{
						Name:    "frank (1)",
						Agility: 50,
						AgilityVals: servicetypes.AgilityVals{
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
						AgilityVals: servicetypes.AgilityVals{
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
			testTurnOrder: []TestBattleTurn{
				{
					index:       0,
					turn:        1,
					currentTick: 0,
					name:        "frank (1)",
					party:       servicetypes.BattlePartyOpponent,
				},
				{
					index:       1,
					turn:        2,
					currentTick: 0,
					name:        "tidus",
					party:       servicetypes.BattlePartyPlayer,
				},
				{
					index:       2,
					turn:        3,
					currentTick: 9,
					name:        "tidus",
					party:       servicetypes.BattlePartyPlayer,
				},
				{
					index:       3,
					turn:        4,
					currentTick: 16,
					name:        "frank (2)",
					party:       servicetypes.BattlePartyOpponent,
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"turn order":     10,
					"player party":   3,
					"opponent party": 1,
				},
				requestBody: map[ts.FieldName]any{
					ts.PfnTurnsAmt:    10,
					ts.PfnBattleStart: string(database.BattleStartAmbush),
					ts.PfnParty: []map[ts.FieldName]any{
						{
							ts.PfnUnitID:  5,
							ts.PfnAgility: 255,
							ts.PfnStatus:  string(database.HasteStatusAutoHaste),
						},
						{
							ts.PfnUnitID:  1,
							ts.PfnAgility: 255,
							ts.PfnStatus:  string(database.HasteStatusAutoHaste),
						},
						{
							ts.PfnUnitID:  7,
							ts.PfnAgility: 255,
							ts.PfnStatus:  string(database.HasteStatusAutoHaste),
						},
					},
					ts.PfnMons: []map[ts.FieldName]any{
						{
							ts.PfnName:    "frank",
							ts.PfnAgility: 255,
						},
					},
				},
			},
			TurnOrderResponse: servicetypes.TurnOrderResponse{
				TurnsAmt:      10,
				IgnFirstTurn:  false,
				IgnImmunities: false,
				BattleStart:   string(database.BattleStartAmbush),
				RNG:           string(database.TurnOrderRngMedian),
				PlayerParty: []servicetypes.Participant{
					{
						Name:    "kimahri",
						Agility: 255,
						AgilityVals: servicetypes.AgilityVals{
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
						AgilityVals: servicetypes.AgilityVals{
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
						AgilityVals: servicetypes.AgilityVals{
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
				OpponentParty: []servicetypes.Participant{
					{
						Name:    "frank",
						Agility: 255,
						AgilityVals: servicetypes.AgilityVals{
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
			testTurnOrder: []TestBattleTurn{
				{
					index:       0,
					turn:        1,
					currentTick: 0,
					name:        "frank",
					party:       servicetypes.BattlePartyOpponent,
				},
				{
					index:       1,
					turn:        2,
					currentTick: 4,
					name:        "tidus",
					party:       servicetypes.BattlePartyPlayer,
				},
				{
					index:       2,
					turn:        3,
					currentTick: 4,
					name:        "kimahri",
					party:       servicetypes.BattlePartyPlayer,
				},
				{
					index:       3,
					turn:        4,
					currentTick: 4,
					name:        "rikku",
					party:       servicetypes.BattlePartyPlayer,
				},
				{
					index:       4,
					turn:        5,
					currentTick: 7,
					name:        "tidus",
					party:       servicetypes.BattlePartyPlayer,
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"turn order":     10,
					"player party":   1,
					"opponent party": 2,
				},
				requestBody: map[ts.FieldName]any{
					ts.PfnTurnsAmt:    10,
					ts.PfnBattleStart: string(database.BattleStartNormal),
					ts.PfnFormationID: 88,
					ts.PfnParty: []map[ts.FieldName]any{
						{
							ts.PfnUnitID:  14,
							ts.PfnAgility: 10,
						},
					},
				},
			},
			TurnOrderResponse: servicetypes.TurnOrderResponse{
				TurnsAmt:      10,
				IgnFirstTurn:  false,
				IgnImmunities: false,
				BattleStart:   string(database.BattleStartNormal),
				RNG:           string(database.TurnOrderRngMedian),
				PlayerParty: []servicetypes.Participant{
					{
						Name:    "anima",
						Agility: 10,
						AgilityVals: servicetypes.AgilityVals{
							TickSpeed: 14,
							MinICV:    h.GetInt32Ptr(41),
							MaxICV:    h.GetInt32Ptr(42),
						},
						FirstStrike:   false,
						TurnsReceived: 3,
						TurnsPercent:  30,
					},
				},
				OpponentParty: []servicetypes.Participant{
					{
						Name:    "basilisk",
						Agility: 9,
						AgilityVals: servicetypes.AgilityVals{
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
						AgilityVals: servicetypes.AgilityVals{
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
			testTurnOrder: []TestBattleTurn{
				{
					index:       0,
					turn:        1,
					currentTick: 38,
					name:        "bite bug",
					party:       servicetypes.BattlePartyOpponent,
				},
				{
					index:       1,
					turn:        2,
					currentTick: 41,
					name:        "anima",
					party:       servicetypes.BattlePartyPlayer,
				},
				{
					index:       2,
					turn:        3,
					currentTick: 47,
					name:        "basilisk",
					party:       servicetypes.BattlePartyOpponent,
				},
				{
					index:       7,
					turn:        8,
					currentTick: 125,
					name:        "anima",
					party:       servicetypes.BattlePartyPlayer,
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"turn order":     10,
					"player party":   1,
					"opponent party": 2,
				},
				requestBody: map[ts.FieldName]any{
					ts.PfnTurnsAmt:    10,
					ts.PfnBattleStart: string(database.BattleStartPreemptive),
					ts.PfnFormationID: 88,
					ts.PfnParty: []map[ts.FieldName]any{
						{
							ts.PfnUnitID:  14,
							ts.PfnAgility: 10,
						},
					},
				},
			},
			TurnOrderResponse: servicetypes.TurnOrderResponse{
				TurnsAmt:      10,
				IgnFirstTurn:  false,
				IgnImmunities: false,
				BattleStart:   string(database.BattleStartPreemptive),
				RNG:           string(database.TurnOrderRngMedian),
				PlayerParty: []servicetypes.Participant{
					{
						Name:    "anima",
						Agility: 10,
						AgilityVals: servicetypes.AgilityVals{
							TickSpeed: 14,
							MinICV:    h.GetInt32Ptr(0),
							MaxICV:    h.GetInt32Ptr(0),
						},
						FirstStrike:   false,
						TurnsReceived: 5,
						TurnsPercent:  50,
					},
				},
				OpponentParty: []servicetypes.Participant{
					{
						Name:    "basilisk",
						Agility: 9,
						AgilityVals: servicetypes.AgilityVals{
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
						AgilityVals: servicetypes.AgilityVals{
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
			testTurnOrder: []TestBattleTurn{
				{
					index:       0,
					turn:        1,
					currentTick: 0,
					name:        "anima",
					party:       servicetypes.BattlePartyPlayer,
				},
				{
					index:       1,
					turn:        2,
					currentTick: 42,
					name:        "anima",
					party:       servicetypes.BattlePartyPlayer,
				},
				{
					index:       2,
					turn:        3,
					currentTick: 84,
					name:        "anima",
					party:       servicetypes.BattlePartyPlayer,
				},
				{
					index:       3,
					turn:        4,
					currentTick: 108,
					name:        "bite bug",
					party:       servicetypes.BattlePartyOpponent,
				},
				{
					index:       5,
					turn:        6,
					currentTick: 135,
					name:        "basilisk",
					party:       servicetypes.BattlePartyOpponent,
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"turn order":     10,
					"player party":   1,
					"opponent party": 1,
				},
				requestBody: map[ts.FieldName]any{
					ts.PfnTurnsAmt:      10,
					ts.PfnIgnFirstTurn:  true,
					ts.PfnIgnImmunities: true,
					ts.PfnParty: []map[ts.FieldName]any{
						{
							ts.PfnUnitID:  2,
							ts.PfnAgility: 24,
						},
					},
					ts.PfnMons: []map[ts.FieldName]any{
						{
							ts.PfnMonsterID: 19,
							ts.PfnStatus:    string(database.HasteStatusHaste),
						},
					},
				},
			},
			TurnOrderResponse: servicetypes.TurnOrderResponse{
				TurnsAmt:      10,
				IgnFirstTurn:  true,
				IgnImmunities: true,
				BattleStart:   string(database.BattleStartNormal),
				RNG:           string(database.TurnOrderRngMedian),
				PlayerParty: []servicetypes.Participant{
					{
						Name:    "yuna",
						Agility: 24,
						AgilityVals: servicetypes.AgilityVals{
							TickSpeed: 9,
							MinICV:    h.GetInt32Ptr(0),
							MaxICV:    h.GetInt32Ptr(0),
						},
						FirstStrike:   false,
						TurnsReceived: 5,
						TurnsPercent:  50,
					},
				},
				OpponentParty: []servicetypes.Participant{
					{
						Name:    "sin - 1",
						Agility: 6,
						AgilityVals: servicetypes.AgilityVals{
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
			testTurnOrder: []TestBattleTurn{
				{
					index:       0,
					turn:        1,
					currentTick: 0,
					name:        "yuna",
					party:       servicetypes.BattlePartyPlayer,
				},
				{
					index:       1,
					turn:        2,
					currentTick: 0,
					name:        "sin - 1",
					party:       servicetypes.BattlePartyOpponent,
				},
				{
					index:       2,
					turn:        3,
					currentTick: 24,
					name:        "sin - 1",
					party:       servicetypes.BattlePartyOpponent,
				},
				{
					index:       3,
					turn:        4,
					currentTick: 27,
					name:        "yuna",
					party:       servicetypes.BattlePartyPlayer,
				},
				{
					index:       8,
					turn:        9,
					currentTick: 96,
					name:        "sin - 1",
					party:       servicetypes.BattlePartyOpponent,
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"turn order":     10,
					"player party":   0,
					"opponent party": 3,
				},
				requestBody: map[ts.FieldName]any{
					ts.PfnTurnsAmt: 10,
					ts.PfnMons: []map[ts.FieldName]any{
						{
							ts.PfnMonsterID: 32,
						},
						{
							ts.PfnMonsterID: 32,
							ts.PfnAltState:  2,
						},
						{
							ts.PfnMonsterID: 32,
							ts.PfnAltState:  3,
						},
					},
				},
			},
			TurnOrderResponse: servicetypes.TurnOrderResponse{
				TurnsAmt:      10,
				IgnFirstTurn:  false,
				IgnImmunities: false,
				BattleStart:   string(database.BattleStartNormal),
				RNG:           string(database.TurnOrderRngMedian),
				PlayerParty:   []servicetypes.Participant{},
				OpponentParty: []servicetypes.Participant{
					{
						Name:    "oblitzerator (1)",
						Agility: 30,
						AgilityVals: servicetypes.AgilityVals{
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
						AgilityVals: servicetypes.AgilityVals{
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
						AgilityVals: servicetypes.AgilityVals{
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
			testTurnOrder: []TestBattleTurn{
				{
					index:       0,
					turn:        1,
					currentTick: 25,
					name:        "oblitzerator (1)",
					party:       servicetypes.BattlePartyOpponent,
				},
				{
					index:       1,
					turn:        2,
					currentTick: 28,
					name:        "oblitzerator (3)",
					party:       servicetypes.BattlePartyOpponent,
				},
				{
					index:       2,
					turn:        3,
					currentTick: 38,
					name:        "oblitzerator (2)",
					party:       servicetypes.BattlePartyOpponent,
				},
				{
					index:       3,
					turn:        4,
					currentTick: 49,
					name:        "oblitzerator (1)",
					party:       servicetypes.BattlePartyOpponent,
				},
				{
					index:       7,
					turn:        8,
					currentTick: 82,
					name:        "oblitzerator (3)",
					party:       servicetypes.BattlePartyOpponent,
				},
			},
		},
		{
			TestGeneral: TestGeneral{
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
				requestBody: map[ts.FieldName]any{
					ts.PfnMons: []map[ts.FieldName]any{
						{
							ts.PfnMonsterID:   93,
							ts.PfnAgility:     55,
							ts.PfnFirstStrike: false,
						},
					},
				},
			},
			TurnOrderResponse: servicetypes.TurnOrderResponse{
				TurnsAmt:      30,
				IgnFirstTurn:  false,
				IgnImmunities: false,
				BattleStart:   string(database.BattleStartNormal),
				RNG:           string(database.TurnOrderRngMedian),
				PlayerParty:   []servicetypes.Participant{},
				OpponentParty: []servicetypes.Participant{
					{
						Name:    "seymour",
						Agility: 55,
						AgilityVals: servicetypes.AgilityVals{
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
			TestGeneral: TestGeneral{
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
				requestBody: map[ts.FieldName]any{
					ts.PfnIgnImmunities: true,
					ts.PfnMons: []map[ts.FieldName]any{
						{
							ts.PfnMonsterID: 306,
							ts.PfnStatus:    string(database.HasteStatusSlow),
						},
					},
				},
			},
			TurnOrderResponse: servicetypes.TurnOrderResponse{
				TurnsAmt:      30,
				IgnFirstTurn:  false,
				IgnImmunities: true,
				BattleStart:   string(database.BattleStartNormal),
				RNG:           string(database.TurnOrderRngMedian),
				PlayerParty:   []servicetypes.Participant{},
				OpponentParty: []servicetypes.Participant{
					{
						Name:    "left arm",
						Agility: 255,
						AgilityVals: servicetypes.AgilityVals{
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
			TestGeneral: TestGeneral{
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
				requestBody: map[ts.FieldName]any{
					ts.PfnFormationID: 226,
				},
			},
			TurnOrderResponse: servicetypes.TurnOrderResponse{
				TurnsAmt:      30,
				IgnFirstTurn:  false,
				IgnImmunities: false,
				BattleStart:   string(database.BattleStartAmbush),
				RNG:           string(database.TurnOrderRngMedian),
				PlayerParty:   []servicetypes.Participant{},
				OpponentParty: []servicetypes.Participant{
					{
						Name:    "dark yojimbo",
						Agility: 243,
						AgilityVals: servicetypes.AgilityVals{
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
						AgilityVals: servicetypes.AgilityVals{
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
			TestGeneral: TestGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"turn order":     0,
					"player party":   1,
					"opponent party": 0,
				},
				requestBody: map[ts.FieldName]any{
					ts.PfnParty: []map[ts.FieldName]any{
						{
							ts.PfnUnitID:  1,
							ts.PfnAgility: 0,
						},
					},
				},
			},
			TurnOrderResponse: servicetypes.TurnOrderResponse{
				TurnsAmt:      30,
				IgnFirstTurn:  false,
				IgnImmunities: false,
				BattleStart:   string(database.BattleStartNormal),
				RNG:           string(database.TurnOrderRngMedian),
				PlayerParty: []servicetypes.Participant{
					{
						Name:    "tidus",
						Agility: 0,
						AgilityVals: servicetypes.AgilityVals{
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
				OpponentParty: []servicetypes.Participant{},
			},
			testTurnOrder: []TestBattleTurn{},
		},
	}

	testServiceResponsesPost(t, tests, "TurnOrder", testSrv.HandleTurnOrderPost, compareTurnOrderResponses)
}
