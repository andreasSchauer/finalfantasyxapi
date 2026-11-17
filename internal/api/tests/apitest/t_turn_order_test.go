package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/service/servicetypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/tests/apitesttypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"

	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
)

func TestTurnOrder(t *testing.T) {
	t.Parallel()
	requestURL := "/api/turn-order"

	tests := []apitesttypes.TurnOrderResponse{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     requestURL,
				Method:         http.MethodPost,
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "at least one of these fields must have a value: 'party', 'mons', 'formation_id'",
				RequestBody:    map[ts.FieldName]any{},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     requestURL,
				Method:         http.MethodPost,
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "value '0' of field 'turns_amt' is out of range. allowed range: 1 to 1000",
				RequestBody: map[ts.FieldName]any{
					ts.PfnTurnsAmt: 0,
				},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     requestURL,
				Method:         http.MethodPost,
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "id '0' of field 'formation_id' is out of range. allowed range: 1 to 331",
				RequestBody: map[ts.FieldName]any{
					ts.PfnFormationID: 0,
				},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     requestURL,
				Method:         http.MethodPost,
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "empty array used for field 'party'. the use of empty arrays is not supported by this api. omit the field, or use 'null' instead",
				RequestBody: map[ts.FieldName]any{
					ts.PfnParty: []map[ts.FieldName]any{},
				},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     requestURL,
				Method:         http.MethodPost,
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "empty array used for field 'mons'. the use of empty arrays is not supported by this api. omit the field, or use 'null' instead",
				RequestBody: map[ts.FieldName]any{
					ts.PfnMons: []map[ts.FieldName]any{},
				},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     requestURL,
				Method:         http.MethodPost,
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "too many items in array of field 'party'. got: 4, maximum allowed: 3",
				RequestBody: map[ts.FieldName]any{
					ts.PfnParty: []map[ts.FieldName]any{{}, {}, {}, {}},
				},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     requestURL,
				Method:         http.MethodPost,
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "field 'formation_id' can't be used in combination with field 'mons'",
				RequestBody: map[ts.FieldName]any{
					ts.PfnFormationID: 0,
					ts.PfnMons:        []map[ts.FieldName]any{{}, {}, {}, {}},
				},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     requestURL,
				Method:         http.MethodPost,
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "index 0 of array at field 'party': id '0' of field 'unit_id' is out of range. allowed range: 1 to 18",
				RequestBody: map[ts.FieldName]any{
					ts.PfnParty: []map[ts.FieldName]any{
						{
							ts.PfnUnitID: 0,
						},
					},
				},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     requestURL,
				Method:         http.MethodPost,
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "index 1 of array at field 'party': field 'unit_id' can't be empty",
				RequestBody: map[ts.FieldName]any{
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
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     requestURL,
				Method:         http.MethodPost,
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "index 1 of array at field 'party': id '0' of field 'unit_id' is out of range. allowed range: 1 to 18",
				RequestBody: map[ts.FieldName]any{
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
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     requestURL,
				Method:         http.MethodPost,
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "index 1 of array at field 'party': field 'agility' can't be empty",
				RequestBody: map[ts.FieldName]any{
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
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     requestURL,
				Method:         http.MethodPost,
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "field 'party': duplicate unit id found at index 1. each party member can only appear once",
				RequestBody: map[ts.FieldName]any{
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
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     requestURL,
				Method:         http.MethodPost,
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "index 0 of array at field 'party': value 'hapfte' of field 'status' is not a valid enum value. allowed values: 'haste', 'slow', 'auto-haste'",
				RequestBody: map[ts.FieldName]any{
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
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     requestURL,
				Method:         http.MethodPost,
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "index 1 of array at field 'party': invalid field: 'text'",
				RequestBody: map[ts.FieldName]any{
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
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     requestURL,
				Method:         http.MethodPost,
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "index 0 of array at field 'mons': at least one of these fields must have a value: 'monster_id', 'name'",
				RequestBody: map[ts.FieldName]any{
					ts.PfnMons: []map[ts.FieldName]any{
						{
							ts.PfnAgility: 25,
						},
					},
				},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     requestURL,
				Method:         http.MethodPost,
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "index 0 of array at field 'mons': field 'name' requires the following fields to have a value: 'agility'",
				RequestBody: map[ts.FieldName]any{
					ts.PfnMons: []map[ts.FieldName]any{
						{
							ts.PfnName: "frank",
						},
					},
				},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     requestURL,
				Method:         http.MethodPost,
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "index 1 of array at field 'mons': field 'name' can't be used in combination with field 'monster_id'",
				RequestBody: map[ts.FieldName]any{
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
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     requestURL,
				Method:         http.MethodPost,
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "index 1 of array at field 'mons': field 'alt_state' requires the following fields to have a value: 'monster_id'",
				RequestBody: map[ts.FieldName]any{
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
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     requestURL,
				Method:         http.MethodPost,
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "index 2 of array at field 'mons': field 'alt_state' can only be used with the following id-value pairs: id: 32, vals: 2, 3 | id: 139, vals: 1, 2, 3 | id: 160, vals: 1 | id: 198, vals: 1 | id: 285, vals: 1 | id: 287, vals: 2",
				RequestBody: map[ts.FieldName]any{
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
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     requestURL,
				Method:         http.MethodPost,
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "index 0 of array at field 'mons': provided value '1' used for 'alt_state' is not available for id '32'. allowed vals: 2, 3",
				RequestBody: map[ts.FieldName]any{
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
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     requestURL,
				Method:         http.MethodPost,
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "index 0 of array at field 'mons': provided value '0' used for 'alt_state' is out of range for id '32'. max val: 3. allowed vals: 2, 3",
				RequestBody: map[ts.FieldName]any{
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
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     requestURL,
				Method:         http.MethodPost,
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "index 0 of array at field 'mons': provided value '4' used for 'alt_state' is out of range for id '32'. max val: 3. allowed vals: 2, 3",
				RequestBody: map[ts.FieldName]any{
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
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     requestURL,
				Method:         http.MethodPost,
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "field 'mons': duplicate monster found at index 1. exact duplicate monsters are not allowed",
				RequestBody: map[ts.FieldName]any{
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
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     requestURL,
				Method:         http.MethodPost,
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "monster 'penance' is immune to 'slow'",
				RequestBody: map[ts.FieldName]any{
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
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     requestURL,
				Method:         http.MethodPost,
				ExpectedStatus: http.StatusOK,
				ExpLengths: map[string]int{
					"turn order":     30,
					"player party":   1,
					"opponent party": 0,
				},
				RequestBody: map[ts.FieldName]any{
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
						TickSpeed: 6,
						MinICV:    h.GetInt32Ptr(14),
						MaxICV:    h.GetInt32Ptr(18),
						FirstStrike:   false,
						Status:        nil,
						TurnsReceived: 30,
						TurnsPercent:  100,
					},
				},
				OpponentParty: []servicetypes.Participant{},
			},
			TestTurnOrder: []apitesttypes.TestBattleTurn{
				{
					Index:       0,
					Turn:        1,
					CurrentTick: 16,
					Name:        "tidus",
					Party:       servicetypes.BattlePartyPlayer,
				},
				{
					Index:       4,
					Turn:        5,
					CurrentTick: 88,
					Name:        "tidus",
					Party:       servicetypes.BattlePartyPlayer,
				},
				{
					Index:       29,
					Turn:        30,
					CurrentTick: 538,
					Name:        "tidus",
					Party:       servicetypes.BattlePartyPlayer,
				},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     requestURL,
				Method:         http.MethodPost,
				ExpectedStatus: http.StatusOK,
				ExpLengths: map[string]int{
					"turn order":     20,
					"player party":   1,
					"opponent party": 1,
				},
				RequestBody: map[ts.FieldName]any{
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
						TickSpeed: 3,
						MinICV:    h.GetInt32Ptr(14),
						MaxICV:    h.GetInt32Ptr(18),
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
						TickSpeed: 32,
						MinICV:    h.GetInt32Ptr(48),
						MaxICV:    h.GetInt32Ptr(53),
						FirstStrike:   false,
						Status:        h.GetStrPtr(string(database.HasteStatusSlow)),
						TurnsReceived: 2,
						TurnsPercent:  10,
					},
				},
			},
			TestTurnOrder: []apitesttypes.TestBattleTurn{
				{
					Index:       0,
					Turn:        1,
					CurrentTick: 18,
					Name:        "tidus",
					Party:       servicetypes.BattlePartyPlayer,
				},
				{
					Index:       4,
					Turn:        5,
					CurrentTick: 48,
					Name:        "yellow element",
					Party:       servicetypes.BattlePartyOpponent,
				},
				{
					Index:       10,
					Turn:        11,
					CurrentTick: 99,
					Name:        "tidus",
					Party:       servicetypes.BattlePartyPlayer,
				},
				{
					Index:       16,
					Turn:        17,
					CurrentTick: 144,
					Name:        "yellow element",
					Party:       servicetypes.BattlePartyOpponent,
				},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     requestURL,
				Method:         http.MethodPost,
				ExpectedStatus: http.StatusOK,
				ExpLengths: map[string]int{
					"turn order":     10,
					"player party":   1,
					"opponent party": 1,
				},
				RequestBody: map[ts.FieldName]any{
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
						TickSpeed: 3,
						MinICV:    h.GetInt32Ptr(7),
						MaxICV:    h.GetInt32Ptr(9),
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
						TickSpeed: 10,
						MinICV:    h.GetInt32Ptr(15),
						MaxICV:    h.GetInt32Ptr(16),
						FirstStrike:   false,
						Status:        h.GetStrPtr(string(database.HasteStatusSlow)),
						TurnsReceived: 2,
						TurnsPercent:  20,
					},
				},
			},
			TestTurnOrder: []apitesttypes.TestBattleTurn{
				{
					Index:       0,
					Turn:        1,
					CurrentTick: 7,
					Name:        "tidus",
					Party:       servicetypes.BattlePartyPlayer,
				},
				{
					Index:       1,
					Turn:        2,
					CurrentTick: 16,
					Name:        "cactuar king",
					Party:       servicetypes.BattlePartyOpponent,
				},
				{
					Index:       2,
					Turn:        3,
					CurrentTick: 16,
					Name:        "tidus",
					Party:       servicetypes.BattlePartyPlayer,
				},
				{
					Index:       3,
					Turn:        4,
					CurrentTick: 25,
					Name:        "tidus",
					Party:       servicetypes.BattlePartyPlayer,
				},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     requestURL,
				Method:         http.MethodPost,
				ExpectedStatus: http.StatusOK,
				ExpLengths: map[string]int{
					"turn order":     10,
					"player party":   1,
					"opponent party": 2,
				},
				RequestBody: map[ts.FieldName]any{
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
						TickSpeed: 3,
						MinICV:    h.GetInt32Ptr(0),
						MaxICV:    h.GetInt32Ptr(0),
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
						TickSpeed: 12,
						MinICV:    h.GetInt32Ptr(-1),
						MaxICV:    h.GetInt32Ptr(-1),
						FirstStrike:   true,
						Status:        h.GetStrPtr(string(database.HasteStatusSlow)),
						TurnsReceived: 1,
						TurnsPercent:  10,
					},
					{
						Name:    "frank (2)",
						Agility: 80,
						TickSpeed: 2,
						MinICV:    h.GetInt32Ptr(15),
						MaxICV:    h.GetInt32Ptr(16),
						FirstStrike:   false,
						Status:        h.GetStrPtr(string(database.HasteStatusHaste)),
						TurnsReceived: 4,
						TurnsPercent:  40,
					},
				},
			},
			TestTurnOrder: []apitesttypes.TestBattleTurn{
				{
					Index:       0,
					Turn:        1,
					CurrentTick: 0,
					Name:        "frank (1)",
					Party:       servicetypes.BattlePartyOpponent,
				},
				{
					Index:       1,
					Turn:        2,
					CurrentTick: 0,
					Name:        "tidus",
					Party:       servicetypes.BattlePartyPlayer,
				},
				{
					Index:       2,
					Turn:        3,
					CurrentTick: 9,
					Name:        "tidus",
					Party:       servicetypes.BattlePartyPlayer,
				},
				{
					Index:       3,
					Turn:        4,
					CurrentTick: 16,
					Name:        "frank (2)",
					Party:       servicetypes.BattlePartyOpponent,
				},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     requestURL,
				Method:         http.MethodPost,
				ExpectedStatus: http.StatusOK,
				ExpLengths: map[string]int{
					"turn order":     10,
					"player party":   3,
					"opponent party": 1,
				},
				RequestBody: map[ts.FieldName]any{
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
						TickSpeed: 1,
						MinICV:    h.GetInt32Ptr(4),
						MaxICV:    h.GetInt32Ptr(4),
						FirstStrike:   false,
						Status:        h.GetStrPtr(string(database.HasteStatusAutoHaste)),
						TurnsReceived: 3,
						TurnsPercent:  30,
					},
					{
						Name:    "tidus",
						Agility: 255,
						TickSpeed: 1,
						MinICV:    h.GetInt32Ptr(4),
						MaxICV:    h.GetInt32Ptr(4),
						FirstStrike:   false,
						Status:        h.GetStrPtr(string(database.HasteStatusAutoHaste)),
						TurnsReceived: 3,
						TurnsPercent:  30,
					},
					{
						Name:    "rikku",
						Agility: 255,
						TickSpeed: 1,
						MinICV:    h.GetInt32Ptr(4),
						MaxICV:    h.GetInt32Ptr(4),
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
						TickSpeed: 3,
						MinICV:    h.GetInt32Ptr(0),
						MaxICV:    h.GetInt32Ptr(0),
						FirstStrike:   false,
						TurnsReceived: 2,
						TurnsPercent:  20,
					},
				},
			},
			TestTurnOrder: []apitesttypes.TestBattleTurn{
				{
					Index:       0,
					Turn:        1,
					CurrentTick: 0,
					Name:        "frank",
					Party:       servicetypes.BattlePartyOpponent,
				},
				{
					Index:       1,
					Turn:        2,
					CurrentTick: 4,
					Name:        "tidus",
					Party:       servicetypes.BattlePartyPlayer,
				},
				{
					Index:       2,
					Turn:        3,
					CurrentTick: 4,
					Name:        "kimahri",
					Party:       servicetypes.BattlePartyPlayer,
				},
				{
					Index:       3,
					Turn:        4,
					CurrentTick: 4,
					Name:        "rikku",
					Party:       servicetypes.BattlePartyPlayer,
				},
				{
					Index:       4,
					Turn:        5,
					CurrentTick: 7,
					Name:        "tidus",
					Party:       servicetypes.BattlePartyPlayer,
				},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     requestURL,
				Method:         http.MethodPost,
				ExpectedStatus: http.StatusOK,
				ExpLengths: map[string]int{
					"turn order":     10,
					"player party":   1,
					"opponent party": 2,
				},
				RequestBody: map[ts.FieldName]any{
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
						TickSpeed: 14,
						MinICV:    h.GetInt32Ptr(41),
						MaxICV:    h.GetInt32Ptr(42),
						FirstStrike:   false,
						TurnsReceived: 3,
						TurnsPercent:  30,
					},
				},
				OpponentParty: []servicetypes.Participant{
					{
						Name:    "basilisk",
						Agility: 9,
						TickSpeed: 15,
						MinICV:    h.GetInt32Ptr(45),
						MaxICV:    h.GetInt32Ptr(50),
						FirstStrike:   false,
						TurnsReceived: 3,
						TurnsPercent:  30,
					},
					{
						Name:    "bite bug",
						Agility: 15,
						TickSpeed: 12,
						MinICV:    h.GetInt32Ptr(36),
						MaxICV:    h.GetInt32Ptr(40),
						FirstStrike:   false,
						TurnsReceived: 4,
						TurnsPercent:  40,
					},
				},
			},
			TestTurnOrder: []apitesttypes.TestBattleTurn{
				{
					Index:       0,
					Turn:        1,
					CurrentTick: 38,
					Name:        "bite bug",
					Party:       servicetypes.BattlePartyOpponent,
				},
				{
					Index:       1,
					Turn:        2,
					CurrentTick: 41,
					Name:        "anima",
					Party:       servicetypes.BattlePartyPlayer,
				},
				{
					Index:       2,
					Turn:        3,
					CurrentTick: 47,
					Name:        "basilisk",
					Party:       servicetypes.BattlePartyOpponent,
				},
				{
					Index:       7,
					Turn:        8,
					CurrentTick: 125,
					Name:        "anima",
					Party:       servicetypes.BattlePartyPlayer,
				},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     requestURL,
				Method:         http.MethodPost,
				ExpectedStatus: http.StatusOK,
				ExpLengths: map[string]int{
					"turn order":     10,
					"player party":   1,
					"opponent party": 2,
				},
				RequestBody: map[ts.FieldName]any{
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
						TickSpeed: 14,
						MinICV:    h.GetInt32Ptr(0),
						MaxICV:    h.GetInt32Ptr(0),
						FirstStrike:   false,
						TurnsReceived: 5,
						TurnsPercent:  50,
					},
				},
				OpponentParty: []servicetypes.Participant{
					{
						Name:    "basilisk",
						Agility: 9,
						TickSpeed: 15,
						MinICV:    h.GetInt32Ptr(135),
						MaxICV:    h.GetInt32Ptr(135),
						FirstStrike:   false,
						TurnsReceived: 2,
						TurnsPercent:  20,
					},
					{
						Name:    "bite bug",
						Agility: 15,
						TickSpeed: 12,
						MinICV:    h.GetInt32Ptr(108),
						MaxICV:    h.GetInt32Ptr(108),
						FirstStrike:   false,
						TurnsReceived: 3,
						TurnsPercent:  30,
					},
				},
			},
			TestTurnOrder: []apitesttypes.TestBattleTurn{
				{
					Index:       0,
					Turn:        1,
					CurrentTick: 0,
					Name:        "anima",
					Party:       servicetypes.BattlePartyPlayer,
				},
				{
					Index:       1,
					Turn:        2,
					CurrentTick: 42,
					Name:        "anima",
					Party:       servicetypes.BattlePartyPlayer,
				},
				{
					Index:       2,
					Turn:        3,
					CurrentTick: 84,
					Name:        "anima",
					Party:       servicetypes.BattlePartyPlayer,
				},
				{
					Index:       3,
					Turn:        4,
					CurrentTick: 108,
					Name:        "bite bug",
					Party:       servicetypes.BattlePartyOpponent,
				},
				{
					Index:       5,
					Turn:        6,
					CurrentTick: 135,
					Name:        "basilisk",
					Party:       servicetypes.BattlePartyOpponent,
				},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     requestURL,
				Method:         http.MethodPost,
				ExpectedStatus: http.StatusOK,
				ExpLengths: map[string]int{
					"turn order":     10,
					"player party":   1,
					"opponent party": 1,
				},
				RequestBody: map[ts.FieldName]any{
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
						TickSpeed: 9,
						MinICV:    h.GetInt32Ptr(0),
						MaxICV:    h.GetInt32Ptr(0),
						FirstStrike:   false,
						TurnsReceived: 5,
						TurnsPercent:  50,
					},
				},
				OpponentParty: []servicetypes.Participant{
					{
						Name:    "sin - 1",
						Agility: 6,
						TickSpeed: 8,
						MinICV:    h.GetInt32Ptr(0),
						MaxICV:    h.GetInt32Ptr(0),
						FirstStrike:   false,
						Status:        h.GetStrPtr(string(database.HasteStatusHaste)),
						TurnsReceived: 5,
						TurnsPercent:  50,
					},
				},
			},
			TestTurnOrder: []apitesttypes.TestBattleTurn{
				{
					Index:       0,
					Turn:        1,
					CurrentTick: 0,
					Name:        "yuna",
					Party:       servicetypes.BattlePartyPlayer,
				},
				{
					Index:       1,
					Turn:        2,
					CurrentTick: 0,
					Name:        "sin - 1",
					Party:       servicetypes.BattlePartyOpponent,
				},
				{
					Index:       2,
					Turn:        3,
					CurrentTick: 24,
					Name:        "sin - 1",
					Party:       servicetypes.BattlePartyOpponent,
				},
				{
					Index:       3,
					Turn:        4,
					CurrentTick: 27,
					Name:        "yuna",
					Party:       servicetypes.BattlePartyPlayer,
				},
				{
					Index:       8,
					Turn:        9,
					CurrentTick: 96,
					Name:        "sin - 1",
					Party:       servicetypes.BattlePartyOpponent,
				},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     requestURL,
				Method:         http.MethodPost,
				ExpectedStatus: http.StatusOK,
				ExpLengths: map[string]int{
					"turn order":     10,
					"player party":   0,
					"opponent party": 3,
				},
				RequestBody: map[ts.FieldName]any{
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
						TickSpeed: 8,
						MinICV:    h.GetInt32Ptr(24),
						MaxICV:    h.GetInt32Ptr(26),
						FirstStrike:   false,
						TurnsReceived: 4,
						TurnsPercent:  40,
					},
					{
						Name:    "oblitzerator (2)",
						Agility: 16,
						TickSpeed: 12,
						MinICV:    h.GetInt32Ptr(36),
						MaxICV:    h.GetInt32Ptr(40),
						FirstStrike:   false,
						AltState:      h.GetInt32Ptr(2),
						TurnsReceived: 2,
						TurnsPercent:  20,
					},
					{
						Name:    "oblitzerator (3)",
						Agility: 26,
						TickSpeed: 9,
						MinICV:    h.GetInt32Ptr(27),
						MaxICV:    h.GetInt32Ptr(30),
						FirstStrike:   false,
						AltState:      h.GetInt32Ptr(3),
						TurnsReceived: 4,
						TurnsPercent:  40,
					},
				},
			},
			TestTurnOrder: []apitesttypes.TestBattleTurn{
				{
					Index:       0,
					Turn:        1,
					CurrentTick: 25,
					Name:        "oblitzerator (1)",
					Party:       servicetypes.BattlePartyOpponent,
				},
				{
					Index:       1,
					Turn:        2,
					CurrentTick: 28,
					Name:        "oblitzerator (3)",
					Party:       servicetypes.BattlePartyOpponent,
				},
				{
					Index:       2,
					Turn:        3,
					CurrentTick: 38,
					Name:        "oblitzerator (2)",
					Party:       servicetypes.BattlePartyOpponent,
				},
				{
					Index:       3,
					Turn:        4,
					CurrentTick: 49,
					Name:        "oblitzerator (1)",
					Party:       servicetypes.BattlePartyOpponent,
				},
				{
					Index:       7,
					Turn:        8,
					CurrentTick: 82,
					Name:        "oblitzerator (3)",
					Party:       servicetypes.BattlePartyOpponent,
				},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     requestURL,
				Method:         http.MethodPost,
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"turn order": true,
				},
				ExpLengths: map[string]int{
					"turn order":     30,
					"player party":   0,
					"opponent party": 1,
				},
				RequestBody: map[ts.FieldName]any{
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
						TickSpeed: 6,
						MinICV:    h.GetInt32Ptr(18),
						MaxICV:    h.GetInt32Ptr(20),
						FirstStrike:   false,
						TurnsReceived: 30,
						TurnsPercent:  100,
					},
				},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     requestURL,
				Method:         http.MethodPost,
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"turn order": true,
				},
				ExpLengths: map[string]int{
					"turn order":     30,
					"player party":   0,
					"opponent party": 1,
				},
				RequestBody: map[ts.FieldName]any{
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
						TickSpeed: 6,
						MinICV:    h.GetInt32Ptr(12),
						MaxICV:    h.GetInt32Ptr(13),
						FirstStrike:   false,
						Status:        h.GetStrPtr(string(database.HasteStatusSlow)),
						TurnsReceived: 30,
						TurnsPercent:  100,
					},
				},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     requestURL,
				Method:         http.MethodPost,
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"turn order": true,
				},
				ExpLengths: map[string]int{
					"turn order":     30,
					"player party":   0,
					"opponent party": 2,
				},
				RequestBody: map[ts.FieldName]any{
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
						TickSpeed: 3,
						MinICV:    h.GetInt32Ptr(0),
						MaxICV:    h.GetInt32Ptr(0),
						FirstStrike:   false,
						TurnsReceived: 30,
						TurnsPercent:  100,
					},
					{
						Name:    "dark koma inu",
						Agility: 0,
						TickSpeed: 0,
						MinICV:    nil,
						MaxICV:    nil,
						FirstStrike:   false,
						TurnsReceived: 0,
						TurnsPercent:  0,
					},
				},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     requestURL,
				Method:         http.MethodPost,
				ExpectedStatus: http.StatusOK,
				ExpLengths: map[string]int{
					"turn order":     0,
					"player party":   1,
					"opponent party": 0,
				},
				RequestBody: map[ts.FieldName]any{
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
						TickSpeed: 0,
						MinICV:    nil,
						MaxICV:    nil,
						FirstStrike:   false,
						Status:        nil,
						TurnsReceived: 0,
						TurnsPercent:  0,
					},
				},
				OpponentParty: []servicetypes.Participant{},
			},
			TestTurnOrder: []apitesttypes.TestBattleTurn{},
		},
	}

	testServiceResponsesPost(t, tests, "TurnOrder", testSrv.HandleTurnOrderPost)
}
