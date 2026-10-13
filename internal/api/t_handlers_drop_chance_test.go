package api

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
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
				expectedErr:    "field 'monster_id' can't be empty.",
				requestBody: map[FieldName]any{
					pfnMinEmptySlots: 2,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "id '10' of field 'character_id' is out of range. allowed range: 1 to 7.",
				requestBody: map[FieldName]any{
					pfnMonsterID: 13,
					pfnCharacterID: 10,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "id '0' of field 'character_id' is out of range. allowed range: 1 to 7.",
				requestBody: map[FieldName]any{
					pfnMonsterID: 13,
					pfnCharacterID: 0,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "array of field 'party_members' can't have duplicate ids.",
				requestBody: map[FieldName]any{
					pfnMonsterID: 13,
					pfnPartyMembers: []any{2,2,2,2,2},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "character must be in the party to be eligible for receiving equipment.",
				requestBody: map[FieldName]any{
					pfnMonsterID: 13,
					pfnCharacterID: 4,
					pfnPartyMembers: []any{1,2,3},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "too many items in array of field 'party_members'. got: 8, maximum allowed: 7.",
				requestBody: map[FieldName]any{
					pfnMonsterID: 13,
					pfnCharacterID: 4,
					pfnPartyMembers: []any{1,2,3,4,5,6,7,8},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "index 0 of field 'auto_abilities': id '131' is out of range. allowed range: 1 to 130.",
				requestBody: map[FieldName]any{
					pfnMonsterID: 13,
					pfnAutoAbilities: []any{131},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "too many items in array of field 'auto_abilities'. got: 5, maximum allowed: 4.",
				requestBody: map[FieldName]any{
					pfnMonsterID: 13,
					pfnAutoAbilities: []any{13,14,15,16,17},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "value '5' of field 'min_empty_slots' is out of range. allowed range: 0 to 4.",
				requestBody: map[FieldName]any{
					pfnMonsterID: 13,
					pfnMinEmptySlots: 5,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "value '5' of field 'total_slots' is out of range. allowed range: 0 to 4.",
				requestBody: map[FieldName]any{
					pfnMonsterID: 13,
					pfnTotalSlots: 5,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "value '-1' of field 'min_empty_slots' is out of range. allowed range: 0 to 4.",
				requestBody: map[FieldName]any{
					pfnMonsterID: 13,
					pfnMinEmptySlots: -1,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "can't discern equipment's type.",
				requestBody: map[FieldName]any{
					pfnMonsterID: 13,
					pfnMinEmptySlots: 2,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "the amount of auto-abilities and empty slots combined can't exceed the total amount of slots.",
				requestBody: map[FieldName]any{
					pfnMonsterID: 13,
					pfnAutoAbilities: []any{2},
					pfnTotalSlots: 0,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "the amount of auto-abilities and empty slots combined can't exceed 4.",
				requestBody: map[FieldName]any{
					pfnMonsterID: 13,
					pfnAutoAbilities: []any{2, 7, 22},
					pfnMinEmptySlots: 2,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "weapon- and armor-abilities can't be combined.",
				requestBody: map[FieldName]any{
					pfnMonsterID: 13,
					pfnAutoAbilities: []any{2, 129},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "weapon- and armor-abilities can't be combined.",
				requestBody: map[FieldName]any{
					pfnMonsterID: 13,
					pfnAutoAbilities: []any{129},
					pfnEquipType: string(database.EquipTypeWeapon),
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "value 'weaponf' of field 'equip_type' is not a valid enum value. allowed values: 'weapon', 'armor'.",
				requestBody: map[FieldName]any{
					pfnMonsterID: 13,
					pfnEquipType: "weaponf",
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "monster 'sinscale - 1 (dream zanarkand)' doesn't drop equipment.",
				requestBody: map[FieldName]any{
					pfnMonsterID: 1,
					pfnAutoAbilities: []any{2},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "water flan - 1 (tutorial) doesn't drop auto-ability 'zombiestrike' (22). it only drops the following auto-abilities: 'piercing' (2), 'waterstrike' (7), 'distill mana' (26), 'water ward' (61), 'waterproof' (62).",
				requestBody: map[FieldName]any{
					pfnMonsterID: 13,
					pfnAutoAbilities: []any{2, 7, 22},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "water flan - 1 (tutorial) doesn't drop auto-ability 'piercing' (2) for lulu.",
				requestBody: map[FieldName]any{
					pfnMonsterID: 13,
					pfnCharacterID: 4,
					pfnAutoAbilities: []any{2, 7},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid field: 'text'.",
				requestBody: map[FieldName]any{
					pfnText: "this won't work.",
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusOK,
				requestBody: map[FieldName]any{
					pfnMonsterID: 27,
					pfnPartyMembers: []any{1,2,3,4,5},
					pfnAutoAbilities: []any{2},
					pfnTotalSlots: 1,
				},
			},
			DropChanceResponse: DropChanceResponse{
				TotalDropChances: TotalDropChances{
					WithFinBlow: 0.7059,
					NoFinBlow: 0.2647,
				},
				EquipmentChances: EquipmentChances{
					MonsterDrop: 4.7059,
					MatchFinBlow: 15,
					MatchNoFinBlow: 5.625,
				},
				CharacterChances: CharacterChances{
					EligibleChars: 1,
					AnyCharFinBlow: 40,
					AnyCharNoFinBlow: 15,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusOK,
				requestBody: map[FieldName]any{
					pfnMonsterID: 305,
					pfnAutoAbilities: []any{51},
				},
			},
			DropChanceResponse: DropChanceResponse{
				TotalDropChances: TotalDropChances{
					WithFinBlow: 50,
					NoFinBlow: 50,
				},
				EquipmentChances: EquipmentChances{
					MonsterDrop: 100,
					MatchFinBlow: 50,
					MatchNoFinBlow: 50,
				},
				CharacterChances: CharacterChances{
					EligibleChars: 7,
					AnyCharFinBlow: 100,
					AnyCharNoFinBlow: 100,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method: 		http.MethodPost,
				expectedStatus: http.StatusOK,
				requestBody: map[FieldName]any{
					pfnMonsterID: 305,
					pfnAutoAbilities: []any{51},
					pfnCharacterID: 1,
				},
			},
			DropChanceResponse: DropChanceResponse{
				TotalDropChances: TotalDropChances{
					WithFinBlow: 17.8571,
					NoFinBlow: 5.3571,
				},
				EquipmentChances: EquipmentChances{
					MonsterDrop: 100,
					MatchFinBlow: 17.8571,
					MatchNoFinBlow: 5.3571,
				},
				CharacterChances: CharacterChances{
					EligibleChars: 7,
					AnyCharFinBlow: 100,
					AnyCharNoFinBlow: 100,
					CharFinBlow: h.GetFloat64Ptr(35.7143),
					CharNoFinBlow: h.GetFloat64Ptr(10.7143),
				},
			},
		},
	}

	testServiceResponses(t, tests, "DropChance", testCfg.HandleDropChanceGet, testCfg.HandleDropChancePost, compareDropChanceResponses)
}