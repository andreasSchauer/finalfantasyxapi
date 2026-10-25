package api

import (
	"net/http"
	"testing"

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
				requestBody:    map[FieldName]any{},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
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
				method:         http.MethodPost,
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
				method:         http.MethodPost,
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
				method:         http.MethodPost,
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
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "too many items in array of field 'party'. got: 4, maximum allowed: 3.",
				requestBody: map[FieldName]any{
					pfnParty: []map[FieldName]any{{}, {}, {}, {}},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "field 'formation_id' can't be used in combination with field 'mons'.",
				requestBody: map[FieldName]any{
					pfnFormationID: 0,
					pfnMons:        []map[FieldName]any{{}, {}, {}, {}},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
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
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "index 1 of field 'party': field 'unit_id' can't be empty.",
				requestBody: map[FieldName]any{
					pfnParty: []map[FieldName]any{
						{
							pfnUnitID:  1,
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
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "index 1 of field 'party': id '0' of field 'unit_id' is out of range. allowed range: 1 to 18.",
				requestBody: map[FieldName]any{
					pfnParty: []map[FieldName]any{
						{
							pfnUnitID:  1,
							pfnAgility: 0,
						},
						{
							pfnUnitID:  0,
							pfnAgility: 0,
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
				requestBody: map[FieldName]any{
					pfnParty: []map[FieldName]any{
						{
							pfnUnitID:  1,
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
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "field 'party': duplicate unit id found at index 1. each party member can only appear once.",
				requestBody: map[FieldName]any{
					pfnParty: []map[FieldName]any{
						{
							pfnUnitID:  1,
							pfnAgility: 255,
						},
						{
							pfnUnitID:  1,
							pfnAgility: 25,
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
				requestBody: map[FieldName]any{
					pfnParty: []map[FieldName]any{
						{
							pfnUnitID:  1,
							pfnAgility: 255,
							pfnStatus:  "hapfte",
						},
						{
							pfnUnitID:  1,
							pfnAgility: 25,
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
				requestBody: map[FieldName]any{
					pfnParty: []map[FieldName]any{
						{
							pfnUnitID:  1,
							pfnAgility: 255,
						},
						{
							pfnText:    1,
							pfnAgility: 25,
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
				method:         http.MethodPost,
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
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "index 1 of field 'mons': field 'name' can't be used in combination with field 'monster_id'.",
				requestBody: map[FieldName]any{
					pfnMons: []map[FieldName]any{
						{
							pfnMonsterID: 23,
						},
						{
							pfnName:      "frank",
							pfnMonsterID: 23,
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
				requestBody: map[FieldName]any{
					pfnMons: []map[FieldName]any{
						{
							pfnMonsterID: 23,
						},
						{
							pfnName:     "frank",
							pfnAltState: 2,
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
							pfnAltState:  2,
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
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "index 0 of field 'mons': provided value '1' used for 'alt_state' is not available for id '32'. allowed vals: 2, 3.",
				requestBody: map[FieldName]any{
					pfnMons: []map[FieldName]any{
						{
							pfnMonsterID: 32,
							pfnAltState:  1,
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
				requestBody: map[FieldName]any{
					pfnMons: []map[FieldName]any{
						{
							pfnMonsterID: 32,
							pfnAltState:  0,
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
				requestBody: map[FieldName]any{
					pfnMons: []map[FieldName]any{
						{
							pfnMonsterID: 32,
							pfnAltState:  4,
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
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "monster 'penance' is immune to 'slow'.",
				requestBody: map[FieldName]any{
					pfnMons: []map[FieldName]any{
						{
							pfnMonsterID: 305,
							pfnStatus:    string(database.HasteStatusSlow),
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
				requestBody: map[FieldName]any{
					pfnParty: []map[FieldName]any{
						{
							pfnUnitID:  1,
							pfnAgility: 50,
						},
					},
				},
			},
			TurnOrderResponse: TurnOrderResponse{
				TurnsAmt:      30,
				IgnFirstTurn:  false,
				IgnImmunities: false,
				BattleStart:   string(database.BattleStartNormal),
				RNG:           string(database.TurnOrderRngMedian),
				PlayerParty: []Participant{
					{
						Name:    "tidus",
						Agility: 50,
						AgilityVals: AgilityVals{
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
				OpponentParty: []Participant{},
			},
			testTurnOrder: []testBattleTurn{
				{
					index:       0,
					turn:        1,
					currentTick: 16,
					name:        "tidus",
					party:       battlePartyPlayer,
				},
				{
					index:       4,
					turn:        5,
					currentTick: 88,
					name:        "tidus",
					party:       battlePartyPlayer,
				},
				{
					index:       29,
					turn:        30,
					currentTick: 538,
					name:        "tidus",
					party:       battlePartyPlayer,
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
				requestBody: map[FieldName]any{
					pfnTurnsAmt: 20,
					pfnRNG:      string(database.TurnOrderRngWorst),
					pfnParty: []map[FieldName]any{
						{
							pfnUnitID:  1,
							pfnAgility: 50,
							pfnStatus:  string(database.HasteStatusHaste),
						},
					},
					pfnMons: []map[FieldName]any{
						{
							pfnMonsterID: 27,
							pfnStatus:    string(database.HasteStatusSlow),
						},
					},
				},
			},
			TurnOrderResponse: TurnOrderResponse{
				TurnsAmt:      20,
				IgnFirstTurn:  false,
				IgnImmunities: false,
				BattleStart:   string(database.BattleStartNormal),
				RNG:           string(database.TurnOrderRngWorst),
				PlayerParty: []Participant{
					{
						Name:    "tidus",
						Agility: 50,
						AgilityVals: AgilityVals{
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
				OpponentParty: []Participant{
					{
						Name:    "yellow element",
						Agility: 5,
						AgilityVals: AgilityVals{
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
					party:       battlePartyPlayer,
				},
				{
					index:       4,
					turn:        5,
					currentTick: 48,
					name:        "yellow element",
					party:       battlePartyOpponent,
				},
				{
					index:       10,
					turn:        11,
					currentTick: 99,
					name:        "tidus",
					party:       battlePartyPlayer,
				},
				{
					index:       16,
					turn:        17,
					currentTick: 144,
					name:        "yellow element",
					party:       battlePartyOpponent,
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
				requestBody: map[FieldName]any{
					pfnTurnsAmt:      10,
					pfnIgnImmunities: true,
					pfnRNG:           string(database.TurnOrderRngBest),
					pfnParty: []map[FieldName]any{
						{
							pfnUnitID:  1,
							pfnAgility: 50,
							pfnStatus:  string(database.HasteStatusAutoHaste),
						},
					},
					pfnMons: []map[FieldName]any{
						{
							pfnMonsterID: 261,
							pfnStatus:    string(database.HasteStatusSlow),
						},
					},
				},
			},
			TurnOrderResponse: TurnOrderResponse{
				TurnsAmt:      10,
				IgnFirstTurn:  false,
				IgnImmunities: true,
				BattleStart:   string(database.BattleStartNormal),
				RNG:           string(database.TurnOrderRngBest),
				PlayerParty: []Participant{
					{
						Name:    "tidus",
						Agility: 50,
						AgilityVals: AgilityVals{
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
				OpponentParty: []Participant{
					{
						Name:    "cactuar king",
						Agility: 80,
						AgilityVals: AgilityVals{
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
					party:       battlePartyPlayer,
				},
				{
					index:       1,
					turn:        2,
					currentTick: 16,
					name:        "cactuar king",
					party:       battlePartyOpponent,
				},
				{
					index:       2,
					turn:        3,
					currentTick: 16,
					name:        "tidus",
					party:       battlePartyPlayer,
				},
				{
					index:       3,
					turn:        4,
					currentTick: 25,
					name:        "tidus",
					party:       battlePartyPlayer,
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
				requestBody: map[FieldName]any{
					pfnTurnsAmt:      10,
					pfnIgnImmunities: true,
					pfnRNG:           string(database.TurnOrderRngBest),
					pfnParty: []map[FieldName]any{
						{
							pfnUnitID:      1,
							pfnAgility:     50,
							pfnFirstStrike: true,
							pfnStatus:      string(database.HasteStatusAutoHaste),
						},
					},
					pfnMons: []map[FieldName]any{
						{
							pfnName:        "frank",
							pfnAgility:     50,
							pfnFirstStrike: true,
							pfnStatus:      string(database.HasteStatusSlow),
						},
						{
							pfnName:    "frank",
							pfnAgility: 80,
							pfnStatus:  string(database.HasteStatusHaste),
						},
					},
				},
			},
			TurnOrderResponse: TurnOrderResponse{
				TurnsAmt:      10,
				IgnFirstTurn:  false,
				IgnImmunities: true,
				BattleStart:   string(database.BattleStartNormal),
				RNG:           string(database.TurnOrderRngBest),
				PlayerParty: []Participant{
					{
						Name:    "tidus",
						Agility: 50,
						AgilityVals: AgilityVals{
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
				OpponentParty: []Participant{
					{
						Name:    "frank (1)",
						Agility: 50,
						AgilityVals: AgilityVals{
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
						AgilityVals: AgilityVals{
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
					party:       battlePartyOpponent,
				},
				{
					index:       1,
					turn:        2,
					currentTick: 0,
					name:        "tidus",
					party:       battlePartyPlayer,
				},
				{
					index:       2,
					turn:        3,
					currentTick: 9,
					name:        "tidus",
					party:       battlePartyPlayer,
				},
				{
					index:       3,
					turn:        4,
					currentTick: 16,
					name:        "frank (2)",
					party:       battlePartyOpponent,
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
				requestBody: map[FieldName]any{
					pfnTurnsAmt:    10,
					pfnBattleStart: string(database.BattleStartAmbush),
					pfnParty: []map[FieldName]any{
						{
							pfnUnitID:  5,
							pfnAgility: 255,
							pfnStatus:  string(database.HasteStatusAutoHaste),
						},
						{
							pfnUnitID:  1,
							pfnAgility: 255,
							pfnStatus:  string(database.HasteStatusAutoHaste),
						},
						{
							pfnUnitID:  7,
							pfnAgility: 255,
							pfnStatus:  string(database.HasteStatusAutoHaste),
						},
					},
					pfnMons: []map[FieldName]any{
						{
							pfnName:    "frank",
							pfnAgility: 255,
						},
					},
				},
			},
			TurnOrderResponse: TurnOrderResponse{
				TurnsAmt:      10,
				IgnFirstTurn:  false,
				IgnImmunities: false,
				BattleStart:   string(database.BattleStartAmbush),
				RNG:           string(database.TurnOrderRngMedian),
				PlayerParty: []Participant{
					{
						Name:    "kimahri",
						Agility: 255,
						AgilityVals: AgilityVals{
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
						AgilityVals: AgilityVals{
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
						AgilityVals: AgilityVals{
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
				OpponentParty: []Participant{
					{
						Name:    "frank",
						Agility: 255,
						AgilityVals: AgilityVals{
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
					party:       battlePartyOpponent,
				},
				{
					index:       1,
					turn:        2,
					currentTick: 4,
					name:        "tidus",
					party:       battlePartyPlayer,
				},
				{
					index:       2,
					turn:        3,
					currentTick: 4,
					name:        "kimahri",
					party:       battlePartyPlayer,
				},
				{
					index:       3,
					turn:        4,
					currentTick: 4,
					name:        "rikku",
					party:       battlePartyPlayer,
				},
				{
					index:       4,
					turn:        5,
					currentTick: 7,
					name:        "tidus",
					party:       battlePartyPlayer,
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
				requestBody: map[FieldName]any{
					pfnTurnsAmt:    10,
					pfnBattleStart: string(database.BattleStartNormal),
					pfnFormationID: 88,
					pfnParty: []map[FieldName]any{
						{
							pfnUnitID:  14,
							pfnAgility: 10,
						},
					},
				},
			},
			TurnOrderResponse: TurnOrderResponse{
				TurnsAmt:      10,
				IgnFirstTurn:  false,
				IgnImmunities: false,
				BattleStart:   string(database.BattleStartNormal),
				RNG:           string(database.TurnOrderRngMedian),
				PlayerParty: []Participant{
					{
						Name:    "anima",
						Agility: 10,
						AgilityVals: AgilityVals{
							TickSpeed: 14,
							MinICV:    h.GetInt32Ptr(41),
							MaxICV:    h.GetInt32Ptr(42),
						},
						FirstStrike:   false,
						TurnsReceived: 3,
						TurnsPercent:  30,
					},
				},
				OpponentParty: []Participant{
					{
						Name:    "basilisk",
						Agility: 9,
						AgilityVals: AgilityVals{
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
						AgilityVals: AgilityVals{
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
					party:       battlePartyOpponent,
				},
				{
					index:       1,
					turn:        2,
					currentTick: 41,
					name:        "anima",
					party:       battlePartyPlayer,
				},
				{
					index:       2,
					turn:        3,
					currentTick: 47,
					name:        "basilisk",
					party:       battlePartyOpponent,
				},
				{
					index:       7,
					turn:        8,
					currentTick: 125,
					name:        "anima",
					party:       battlePartyPlayer,
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
				requestBody: map[FieldName]any{
					pfnTurnsAmt:    10,
					pfnBattleStart: string(database.BattleStartPreemptive),
					pfnFormationID: 88,
					pfnParty: []map[FieldName]any{
						{
							pfnUnitID:  14,
							pfnAgility: 10,
						},
					},
				},
			},
			TurnOrderResponse: TurnOrderResponse{
				TurnsAmt:      10,
				IgnFirstTurn:  false,
				IgnImmunities: false,
				BattleStart:   string(database.BattleStartPreemptive),
				RNG:           string(database.TurnOrderRngMedian),
				PlayerParty: []Participant{
					{
						Name:    "anima",
						Agility: 10,
						AgilityVals: AgilityVals{
							TickSpeed: 14,
							MinICV:    h.GetInt32Ptr(0),
							MaxICV:    h.GetInt32Ptr(0),
						},
						FirstStrike:   false,
						TurnsReceived: 5,
						TurnsPercent:  50,
					},
				},
				OpponentParty: []Participant{
					{
						Name:    "basilisk",
						Agility: 9,
						AgilityVals: AgilityVals{
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
						AgilityVals: AgilityVals{
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
					party:       battlePartyPlayer,
				},
				{
					index:       1,
					turn:        2,
					currentTick: 42,
					name:        "anima",
					party:       battlePartyPlayer,
				},
				{
					index:       2,
					turn:        3,
					currentTick: 84,
					name:        "anima",
					party:       battlePartyPlayer,
				},
				{
					index:       3,
					turn:        4,
					currentTick: 108,
					name:        "bite bug",
					party:       battlePartyOpponent,
				},
				{
					index:       5,
					turn:        6,
					currentTick: 135,
					name:        "basilisk",
					party:       battlePartyOpponent,
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
				requestBody: map[FieldName]any{
					pfnTurnsAmt:      10,
					pfnIgnFirstTurn:  true,
					pfnIgnImmunities: true,
					pfnParty: []map[FieldName]any{
						{
							pfnUnitID:  2,
							pfnAgility: 24,
						},
					},
					pfnMons: []map[FieldName]any{
						{
							pfnMonsterID: 19,
							pfnStatus:    string(database.HasteStatusHaste),
						},
					},
				},
			},
			TurnOrderResponse: TurnOrderResponse{
				TurnsAmt:      10,
				IgnFirstTurn:  true,
				IgnImmunities: true,
				BattleStart:   string(database.BattleStartNormal),
				RNG:           string(database.TurnOrderRngMedian),
				PlayerParty: []Participant{
					{
						Name:    "yuna",
						Agility: 24,
						AgilityVals: AgilityVals{
							TickSpeed: 9,
							MinICV:    h.GetInt32Ptr(0),
							MaxICV:    h.GetInt32Ptr(0),
						},
						FirstStrike:   false,
						TurnsReceived: 5,
						TurnsPercent:  50,
					},
				},
				OpponentParty: []Participant{
					{
						Name:    "sin - 1",
						Agility: 6,
						AgilityVals: AgilityVals{
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
					party:       battlePartyPlayer,
				},
				{
					index:       1,
					turn:        2,
					currentTick: 0,
					name:        "sin - 1",
					party:       battlePartyOpponent,
				},
				{
					index:       2,
					turn:        3,
					currentTick: 24,
					name:        "sin - 1",
					party:       battlePartyOpponent,
				},
				{
					index:       3,
					turn:        4,
					currentTick: 27,
					name:        "yuna",
					party:       battlePartyPlayer,
				},
				{
					index:       8,
					turn:        9,
					currentTick: 96,
					name:        "sin - 1",
					party:       battlePartyOpponent,
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
				requestBody: map[FieldName]any{
					pfnTurnsAmt: 10,
					pfnMons: []map[FieldName]any{
						{
							pfnMonsterID: 32,
						},
						{
							pfnMonsterID: 32,
							pfnAltState:  2,
						},
						{
							pfnMonsterID: 32,
							pfnAltState:  3,
						},
					},
				},
			},
			TurnOrderResponse: TurnOrderResponse{
				TurnsAmt:      10,
				IgnFirstTurn:  false,
				IgnImmunities: false,
				BattleStart:   string(database.BattleStartNormal),
				RNG:           string(database.TurnOrderRngMedian),
				PlayerParty:   []Participant{},
				OpponentParty: []Participant{
					{
						Name:    "oblitzerator (1)",
						Agility: 30,
						AgilityVals: AgilityVals{
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
						AgilityVals: AgilityVals{
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
						AgilityVals: AgilityVals{
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
					party:       battlePartyOpponent,
				},
				{
					index:       1,
					turn:        2,
					currentTick: 28,
					name:        "oblitzerator (3)",
					party:       battlePartyOpponent,
				},
				{
					index:       2,
					turn:        3,
					currentTick: 38,
					name:        "oblitzerator (2)",
					party:       battlePartyOpponent,
				},
				{
					index:       3,
					turn:        4,
					currentTick: 49,
					name:        "oblitzerator (1)",
					party:       battlePartyOpponent,
				},
				{
					index:       7,
					turn:        8,
					currentTick: 82,
					name:        "oblitzerator (3)",
					party:       battlePartyOpponent,
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
				requestBody: map[FieldName]any{
					pfnMons: []map[FieldName]any{
						{
							pfnMonsterID:   93,
							pfnAgility:     55,
							pfnFirstStrike: false,
						},
					},
				},
			},
			TurnOrderResponse: TurnOrderResponse{
				TurnsAmt:      30,
				IgnFirstTurn:  false,
				IgnImmunities: false,
				BattleStart:   string(database.BattleStartNormal),
				RNG:           string(database.TurnOrderRngMedian),
				PlayerParty:   []Participant{},
				OpponentParty: []Participant{
					{
						Name:    "seymour",
						Agility: 55,
						AgilityVals: AgilityVals{
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
				requestBody: map[FieldName]any{
					pfnIgnImmunities: true,
					pfnMons: []map[FieldName]any{
						{
							pfnMonsterID: 306,
							pfnStatus:    string(database.HasteStatusSlow),
						},
					},
				},
			},
			TurnOrderResponse: TurnOrderResponse{
				TurnsAmt:      30,
				IgnFirstTurn:  false,
				IgnImmunities: true,
				BattleStart:   string(database.BattleStartNormal),
				RNG:           string(database.TurnOrderRngMedian),
				PlayerParty:   []Participant{},
				OpponentParty: []Participant{
					{
						Name:    "left arm",
						Agility: 255,
						AgilityVals: AgilityVals{
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
				requestBody: map[FieldName]any{
					pfnFormationID: 226,
				},
			},
			TurnOrderResponse: TurnOrderResponse{
				TurnsAmt:      30,
				IgnFirstTurn:  false,
				IgnImmunities: false,
				BattleStart:   string(database.BattleStartAmbush),
				RNG:           string(database.TurnOrderRngMedian),
				PlayerParty:   []Participant{},
				OpponentParty: []Participant{
					{
						Name:    "dark yojimbo",
						Agility: 243,
						AgilityVals: AgilityVals{
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
						AgilityVals: AgilityVals{
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
				requestBody: map[FieldName]any{
					pfnParty: []map[FieldName]any{
						{
							pfnUnitID:  1,
							pfnAgility: 0,
						},
					},
				},
			},
			TurnOrderResponse: TurnOrderResponse{
				TurnsAmt:      30,
				IgnFirstTurn:  false,
				IgnImmunities: false,
				BattleStart:   string(database.BattleStartNormal),
				RNG:           string(database.TurnOrderRngMedian),
				PlayerParty: []Participant{
					{
						Name:    "tidus",
						Agility: 0,
						AgilityVals: AgilityVals{
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
				OpponentParty: []Participant{},
			},
			testTurnOrder: []testBattleTurn{},
		},
	}

	testServiceResponsesPost(t, tests, "TurnOrder", testCfg.HandleTurnOrderPost, compareTurnOrderResponses)
}
