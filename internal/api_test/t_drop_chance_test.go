package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func TestDropChance(t *testing.T) {
	t.Parallel()
	requestURL := "/api/drop-chance"

	tests := []expDropChanceResponse{
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "field 'monster_id' can't be empty",
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMinEmptySlots: 2,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "id '10' of field 'character_id' is out of range. allowed range: 1 to 7",
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:     13,
					typestrings.PfnCharacterID:   10,
					typestrings.PfnMinEmptySlots: 1,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "id '0' of field 'character_id' is out of range. allowed range: 1 to 7",
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:     13,
					typestrings.PfnCharacterID:   0,
					typestrings.PfnMinEmptySlots: 1,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "array of field 'party_members' can't have duplicate ids",
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:     13,
					typestrings.PfnPartyMembers:  []any{2, 2, 2, 2, 2},
					typestrings.PfnMinEmptySlots: 1,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "character must be in the party to be eligible for receiving equipment",
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:     13,
					typestrings.PfnCharacterID:   4,
					typestrings.PfnPartyMembers:  []any{1, 2, 3},
					typestrings.PfnMinEmptySlots: 1,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "too many items in array of field 'party_members'. got: 8, maximum allowed: 7",
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:     13,
					typestrings.PfnCharacterID:   4,
					typestrings.PfnPartyMembers:  []any{1, 2, 3, 4, 5, 6, 7, 8},
					typestrings.PfnMinEmptySlots: 1,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "index 0 of array at field 'auto_abilities': id '131' is out of range. allowed range: 1 to 130",
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:     13,
					typestrings.PfnAutoAbilities: []any{131},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "too many items in array of field 'auto_abilities'. got: 5, maximum allowed: 4",
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:     13,
					typestrings.PfnAutoAbilities: []any{13, 14, 15, 16, 17},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "value '5' of field 'min_empty_slots' is out of range. allowed range: 0 to 4",
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:     13,
					typestrings.PfnMinEmptySlots: 5,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "value '5' of field 'total_slots' is out of range. allowed range: 0 to 4",
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:  13,
					typestrings.PfnTotalSlots: 5,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "value '-1' of field 'min_empty_slots' is out of range. allowed range: 0 to 4",
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:     13,
					typestrings.PfnMinEmptySlots: -1,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "can't discern equipment's type. use 'auto_abilities' to discern the type, or set it manually by using 'equip_type'",
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:     13,
					typestrings.PfnMinEmptySlots: 2,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "the amount of auto-abilities and empty slots combined (1) can't exceed the given total amount of slots (0)",
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:     13,
					typestrings.PfnAutoAbilities: []any{2},
					typestrings.PfnTotalSlots:    0,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "the amount of auto-abilities and empty slots combined can't exceed 4",
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:     13,
					typestrings.PfnAutoAbilities: []any{2, 7, 22},
					typestrings.PfnMinEmptySlots: 2,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "weapon- and armor-abilities can't be combined",
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:     13,
					typestrings.PfnAutoAbilities: []any{2, 129},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "weapon- and armor-abilities can't be combined",
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:     13,
					typestrings.PfnAutoAbilities: []any{2},
					typestrings.PfnEquipType:     string(database.EquipTypeArmor),
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "weapon- and armor-abilities can't be combined",
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:     13,
					typestrings.PfnAutoAbilities: []any{129},
					typestrings.PfnEquipType:     string(database.EquipTypeWeapon),
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "value 'weaponf' of field 'equip_type' is not a valid enum value. allowed values: 'weapon', 'armor'",
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:     13,
					typestrings.PfnMinEmptySlots: 1,
					typestrings.PfnEquipType:     "weaponf",
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "monster 'sinscale - 1 (dream zanarkand)' doesn't drop equipment",
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:     1,
					typestrings.PfnAutoAbilities: []any{2},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "water flan - 1 (tutorial) doesn't drop auto-ability 'zombiestrike' (22). it only drops the following auto-abilities: 'piercing' (2), 'waterstrike' (7), 'distill mana' (26), 'water ward' (61), 'waterproof' (62)",
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:     13,
					typestrings.PfnAutoAbilities: []any{7, 22},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "kottos doesn't drop auto-ability 'piercing' (2). it only drops the following auto-abilities: 'counterattack' (37), 'evade & counter' (38), 'magic counter' (39), 'hp +30%' (114), 'hp +20%' (113), 'hp +10%' (112)",
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:     258,
					typestrings.PfnAutoAbilities: []any{2},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "at least one of these fields must have a value: 'min_empty_slots', 'total_slots', 'auto_abilities'",
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID: 13,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid field: 'text'",
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnText: "this won't work",
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:     27,
					typestrings.PfnPartyMembers:  []any{1, 2, 3, 4, 5},
					typestrings.PfnAutoAbilities: []any{2},
					typestrings.PfnTotalSlots:    1,
				},
			},
			DropChanceResponse: api.DropChanceResponse{
				DropChanceFinBlow:   0.7059,
				DropChanceNoFinBlow: 0.2647,
				WantedEquipment: api.WantedEquipment{
					MinSlots:         1,
					MaxSlots:         1,
					ReqAbilities:     []string{"piercing"},
					ReqEmptySlots:    nil,
					MinOptionalSlots: 0,
					MaxOptionalSlots: 0,
					OptionalSlotsAre: string(api.OptionalSlotAllEmpty),
				},
				EquipmentMatchChances: api.EquipmentMatchChances{
					FinBlow:   15,
					NoFinBlow: 5.625,
				},
				CharacterChances: api.CharacterChances{
					EligibleChars: []string{"kimahri"},
					AnyFinBlow:    40,
					AnyNoFinBlow:  15,
				},
				MonsterChances: api.MonsterDropChances{
					MonsterID:           27,
					MonsterName:         "yellow element",
					EquipmentDropChance: 4.7059,
					SlotsChances: []api.ChanceAmount{
						{
							Amount: 1,
							Chance: 75,
						},
						{
							Amount: 2,
							Chance: 25,
						},
					},
					ShotsChances: []api.ChanceAmount{
						{
							Amount: 0,
							Chance: 50,
						},
						{
							Amount: 1,
							Chance: 50,
						},
					},
					AutoAbilities: api.DroppedAutoAbilities{
						Forced: &api.EqAbilitySimple{
							ID:         2,
							Name:       "piercing",
							Characters: []string{"kimahri", "auron"},
						},
						Dropped: []api.EqAbilityChance{
							{
								EqAbilitySimple: api.EqAbilitySimple{
									ID:         6,
									Name:       "lightningstrike",
									Characters: []string{},
								},
								Chance: 57.1429,
							},
							{
								EqAbilitySimple: api.EqAbilitySimple{
									ID:         26,
									Name:       "distill mana",
									Characters: []string{},
								},
								Chance: 42.8571,
							},
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
				dontCheck: map[string]bool{
					"monster drop chances": true,
				},
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:     27,
					typestrings.PfnAutoAbilities: []any{2},
					typestrings.PfnCharacterID:   5,
				},
			},
			DropChanceResponse: api.DropChanceResponse{
				DropChanceFinBlow:   0.7353,
				DropChanceNoFinBlow: 0.2206,
				WantedEquipment: api.WantedEquipment{
					MinSlots:         1,
					MaxSlots:         2,
					ReqAbilities:     []string{"piercing"},
					ReqEmptySlots:    nil,
					MinOptionalSlots: 0,
					MaxOptionalSlots: 1,
					OptionalSlotsAre: string(api.OptionalSlotAllEmpty),
				},
				EquipmentMatchChances: api.EquipmentMatchChances{
					FinBlow:   15.625,
					NoFinBlow: 4.6875,
				},
				CharacterChances: api.CharacterChances{
					EligibleChars: []string{"kimahri", "auron"},
					AnyFinBlow:    46.4286,
					AnyNoFinBlow:  21.4286,
					SpecificChar: &api.SpecificCharChances{
						FinBlow:    35.7143,
						NoFinBlow:  10.7143,
						IsEligible: true,
					},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:     50,
					typestrings.PfnAutoAbilities: []any{9},
					typestrings.PfnCharacterID:   5,
					typestrings.PfnDecimals:      true,
				},
			},
			DropChanceResponse: api.DropChanceResponse{
				DropChanceFinBlow:   0,
				DropChanceNoFinBlow: 0,
				WantedEquipment: api.WantedEquipment{
					MinSlots:         1,
					MaxSlots:         2,
					ReqAbilities:     []string{"darktouch"},
					ReqEmptySlots:    nil,
					MinOptionalSlots: 0,
					MaxOptionalSlots: 1,
					OptionalSlotsAre: string(api.OptionalSlotAllEmpty),
				},
				EquipmentMatchChances: api.EquipmentMatchChances{
					FinBlow:   0,
					NoFinBlow: 0,
				},
				CharacterChances: api.CharacterChances{
					EligibleChars: []string{"tidus", "yuna", "wakka", "lulu", "rikku"},
					AnyFinBlow:    0.785714,
					AnyNoFinBlow:  0.535714,
					SpecificChar: &api.SpecificCharChances{
						FinBlow:    0.357143,
						NoFinBlow:  0.107143,
						IsEligible: false,
					},
				},
				MonsterChances: api.MonsterDropChances{
					MonsterID:           50,
					MonsterName:         "garuda - 3 (regular encounter)",
					EquipmentDropChance: 0.501961,
					SlotsChances: []api.ChanceAmount{
						{
							Amount: 1,
							Chance: 0.5,
						},
						{
							Amount: 2,
							Chance: 0.5,
						},
					},
					ShotsChances: []api.ChanceAmount{
						{
							Amount: 1,
							Chance: 0.75,
						},
						{
							Amount: 2,
							Chance: 0.25,
						},
					},
					AutoAbilities: api.DroppedAutoAbilities{
						Forced: &api.EqAbilitySimple{
							ID:         2,
							Name:       "piercing",
							Characters: []string{"kimahri", "auron"},
						},
						Dropped: []api.EqAbilityChance{
							{
								EqAbilitySimple: api.EqAbilitySimple{
									ID:         9,
									Name:       "darktouch",
									Characters: []string{},
								},
								Chance: 0.285714,
							},
							{
								EqAbilitySimple: api.EqAbilitySimple{
									ID:         29,
									Name:       "strength +3%",
									Characters: []string{"tidus", "wakka", "kimahri", "auron", "rikku"},
								},
								Chance: 0.714286,
							},
							{
								EqAbilitySimple: api.EqAbilitySimple{
									ID:         33,
									Name:       "magic +3%",
									Characters: []string{"yuna", "lulu"},
								},
								Chance: 0.714286,
							},
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
				dontCheck: map[string]bool{
					"monster drop chances": true,
					"wanted equipment":     true,
				},
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:     50,
					typestrings.PfnAutoAbilities: []any{9},
					typestrings.PfnCharacterID:   5,
				},
			},
			DropChanceResponse: api.DropChanceResponse{
				DropChanceFinBlow:   0,
				DropChanceNoFinBlow: 0,
				EquipmentMatchChances: api.EquipmentMatchChances{
					FinBlow:   0,
					NoFinBlow: 0,
				},
				CharacterChances: api.CharacterChances{
					EligibleChars: []string{"tidus", "yuna", "wakka", "lulu", "rikku"},
					AnyFinBlow:    78.5714,
					AnyNoFinBlow:  53.5714,
					SpecificChar: &api.SpecificCharChances{
						FinBlow:    35.7143,
						NoFinBlow:  10.7143,
						IsEligible: false,
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
					"character chances":    true,
					"monster drop chances": true,
				},
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:         50,
					typestrings.PfnAutoAbilities:     []any{9},
					typestrings.PfnOptionalAbilities: true,
					typestrings.PfnCharacterID:       5,
				},
			},
			DropChanceResponse: api.DropChanceResponse{
				DropChanceFinBlow:   1.2805,
				DropChanceNoFinBlow: 0.3842,
				WantedEquipment: api.WantedEquipment{
					MinSlots:         1,
					MaxSlots:         2,
					ReqAbilities:     []string{"darktouch"},
					ReqEmptySlots:    nil,
					MinOptionalSlots: 0,
					MaxOptionalSlots: 1,
					OptionalSlotsAre: string(api.OptionalSlotAbilityOrEmpty),
				},
				EquipmentMatchChances: api.EquipmentMatchChances{
					FinBlow:   2.551,
					NoFinBlow: 0.7653,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"monster drop chances": true,
					"wanted equipment":     true,
				},
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:     50,
					typestrings.PfnAutoAbilities: []any{9},
					typestrings.PfnMinEmptySlots: 1,
					typestrings.PfnCharacterID:   5,
				},
			},
			DropChanceResponse: api.DropChanceResponse{
				DropChanceFinBlow:   0,
				DropChanceNoFinBlow: 0,
				EquipmentMatchChances: api.EquipmentMatchChances{
					FinBlow:   0,
					NoFinBlow: 0,
				},
				CharacterChances: api.CharacterChances{
					EligibleChars: []string{"tidus", "yuna", "wakka", "lulu", "rikku"},
					AnyFinBlow:    78.5714,
					AnyNoFinBlow:  53.5714,
					SpecificChar: &api.SpecificCharChances{
						IsEligible: false,
						FinBlow:    35.7143,
						NoFinBlow:  10.7143,
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
					"character chances":    true,
					"monster drop chances": true,
					"wanted equipment":     true,
				},
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:     50,
					typestrings.PfnAutoAbilities: []any{9},
					typestrings.PfnTotalSlots:    2,
					typestrings.PfnCharacterID:   5,
				},
			},
			DropChanceResponse: api.DropChanceResponse{
				DropChanceFinBlow:   0,
				DropChanceNoFinBlow: 0,
				EquipmentMatchChances: api.EquipmentMatchChances{
					FinBlow:   0,
					NoFinBlow: 0,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"monster drop chances": true,
					"wanted equipment":     true,
				},
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:         50,
					typestrings.PfnAutoAbilities:     []any{2},
					typestrings.PfnOptionalAbilities: true,
					typestrings.PfnTotalSlots:        2,
					typestrings.PfnCharacterID:       5,
				},
			},
			DropChanceResponse: api.DropChanceResponse{
				DropChanceFinBlow:   4.4818,
				DropChanceNoFinBlow: 1.3445,
				EquipmentMatchChances: api.EquipmentMatchChances{
					FinBlow:   8.9286,
					NoFinBlow: 2.6786,
				},
				CharacterChances: api.CharacterChances{
					EligibleChars: []string{"kimahri", "auron"},
					AnyFinBlow:    46.4286,
					AnyNoFinBlow:  21.4286,
					SpecificChar: &api.SpecificCharChances{
						FinBlow:    35.7143,
						NoFinBlow:  10.7143,
						IsEligible: true,
					},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:         258,
					typestrings.PfnAutoAbilities:     []any{38},
					typestrings.PfnOptionalAbilities: true,
				},
			},
			DropChanceResponse: api.DropChanceResponse{
				DropChanceFinBlow:   0,
				DropChanceNoFinBlow: 0,
				WantedEquipment: api.WantedEquipment{
					MinSlots:         2,
					MaxSlots:         3,
					ReqAbilities:     []string{"evade & counter"},
					ReqEmptySlots:    nil,
					MinOptionalSlots: 1,
					MaxOptionalSlots: 2,
					OptionalSlotsAre: string(api.OptionalSlotAbilityOrEmpty),
				},
				EquipmentMatchChances: api.EquipmentMatchChances{
					FinBlow:   0,
					NoFinBlow: 0,
				},
				CharacterChances: api.CharacterChances{
					EligibleChars: []string{"tidus", "yuna", "wakka", "lulu", "kimahri", "auron", "rikku"},
					AnyFinBlow:    100,
					AnyNoFinBlow:  100,
				},
				MonsterChances: api.MonsterDropChances{
					MonsterID:           258,
					MonsterName:         "kottos",
					EquipmentDropChance: 100,
					SlotsChances: []api.ChanceAmount{
						{
							Amount: 2,
							Chance: 50,
						},
						{
							Amount: 3,
							Chance: 50,
						},
					},
					ShotsChances: []api.ChanceAmount{
						{
							Amount: 4,
							Chance: 50,
						},
						{
							Amount: 5,
							Chance: 50,
						},
					},
					AutoAbilities: api.DroppedAutoAbilities{
						Forced: &api.EqAbilitySimple{
							ID:         37,
							Name:       "counterattack",
							Characters: []string{},
						},
						Dropped: []api.EqAbilityChance{
							{
								EqAbilitySimple: api.EqAbilitySimple{
									ID:         37,
									Name:       "counterattack",
									Characters: []string{},
								},
								Chance: 28.5714,
							},
							{
								EqAbilitySimple: api.EqAbilitySimple{
									ID:         38,
									Name:       "evade & counter",
									Characters: []string{},
								},
								Chance: 42.8571,
							},
							{
								EqAbilitySimple: api.EqAbilitySimple{
									ID:         39,
									Name:       "magic counter",
									Characters: []string{},
								},
								Chance: 28.5714,
							},
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
				dontCheck: map[string]bool{
					"monster drop chances": true,
					"wanted equipment":     true,
				},
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:         258,
					typestrings.PfnAutoAbilities:     []any{38},
					typestrings.PfnOptionalAbilities: false,
				},
			},
			DropChanceResponse: api.DropChanceResponse{
				DropChanceFinBlow:   0,
				DropChanceNoFinBlow: 0,
				EquipmentMatchChances: api.EquipmentMatchChances{
					FinBlow:   0,
					NoFinBlow: 0,
				},
				CharacterChances: api.CharacterChances{
					EligibleChars: []string{},
					AnyFinBlow:    0,
					AnyNoFinBlow:  0,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"monster drop chances": true,
					"wanted equipment":     true,
				},
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:         258,
					typestrings.PfnAutoAbilities:     []any{37},
					typestrings.PfnOptionalAbilities: true,
				},
			},
			DropChanceResponse: api.DropChanceResponse{
				DropChanceFinBlow:   50,
				DropChanceNoFinBlow: 50,
				EquipmentMatchChances: api.EquipmentMatchChances{
					FinBlow:   50,
					NoFinBlow: 50,
				},
				CharacterChances: api.CharacterChances{
					EligibleChars: []string{"tidus", "yuna", "wakka", "lulu", "kimahri", "auron", "rikku"},
					AnyFinBlow:    100,
					AnyNoFinBlow:  100,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"monster drop chances": true,
					"wanted equipment":     true,
				},
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:     243,
					typestrings.PfnAutoAbilities: []any{30, 19},
				},
			},
			DropChanceResponse: api.DropChanceResponse{
				DropChanceFinBlow:   0.2744,
				DropChanceNoFinBlow: 0.1543,
				EquipmentMatchChances: api.EquipmentMatchChances{
					FinBlow:   1.1662,
					NoFinBlow: 0.656,
				},
				CharacterChances: api.CharacterChances{
					EligibleChars: []string{"tidus", "wakka", "rikku"},
					AnyFinBlow:    57.1429,
					AnyNoFinBlow:  32.1429,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"monster drop chances": true,
					"wanted equipment":     true,
				},
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:     243,
					typestrings.PfnPartyMembers:  []any{1, 3, 7},
					typestrings.PfnAutoAbilities: []any{30, 19},
				},
			},
			DropChanceResponse: api.DropChanceResponse{
				DropChanceFinBlow:   0.4802,
				DropChanceNoFinBlow: 0.4802,
				EquipmentMatchChances: api.EquipmentMatchChances{
					FinBlow:   2.0408,
					NoFinBlow: 2.0408,
				},
				CharacterChances: api.CharacterChances{
					EligibleChars: []string{"tidus", "wakka", "rikku"},
					AnyFinBlow:    100,
					AnyNoFinBlow:  100,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"monster drop chances": true,
					"wanted equipment":     true,
				},
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:     243,
					typestrings.PfnAutoAbilities: []any{34, 19},
				},
			},
			DropChanceResponse: api.DropChanceResponse{
				DropChanceFinBlow:   0.2229,
				DropChanceNoFinBlow: 0.1029,
				EquipmentMatchChances: api.EquipmentMatchChances{
					FinBlow:   0.9475,
					NoFinBlow: 0.4373,
				},
				CharacterChances: api.CharacterChances{
					EligibleChars: []string{"yuna", "lulu"},
					AnyFinBlow:    46.4286,
					AnyNoFinBlow:  21.4286,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"monster drop chances": true,
					"wanted equipment":     true,
				},
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:         243,
					typestrings.PfnAutoAbilities:     []any{30, 19},
					typestrings.PfnCharacterID:       5,
					typestrings.PfnOptionalAbilities: false,
				},
			},
			DropChanceResponse: api.DropChanceResponse{
				DropChanceFinBlow:   0,
				DropChanceNoFinBlow: 0,
				EquipmentMatchChances: api.EquipmentMatchChances{
					FinBlow:   0,
					NoFinBlow: 0,
				},
				CharacterChances: api.CharacterChances{
					EligibleChars: []string{"tidus", "wakka", "rikku"},
					AnyFinBlow:    57.1429,
					AnyNoFinBlow:  32.1429,
					SpecificChar: &api.SpecificCharChances{
						FinBlow:    35.7143,
						NoFinBlow:  10.7143,
						IsEligible: false,
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
					"monster drop chances": true,
					"wanted equipment":     true,
				},
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:         243,
					typestrings.PfnAutoAbilities:     []any{30, 19},
					typestrings.PfnCharacterID:       3,
					typestrings.PfnOptionalAbilities: false,
				},
			},
			DropChanceResponse: api.DropChanceResponse{
				DropChanceFinBlow:   0.1715,
				DropChanceNoFinBlow: 0.0514,
				EquipmentMatchChances: api.EquipmentMatchChances{
					FinBlow:   0.7289,
					NoFinBlow: 0.2187,
				},
				CharacterChances: api.CharacterChances{
					EligibleChars: []string{"tidus", "wakka", "rikku"},
					AnyFinBlow:    57.1429,
					AnyNoFinBlow:  32.1429,
					SpecificChar: &api.SpecificCharChances{
						FinBlow:    35.7143,
						NoFinBlow:  10.7143,
						IsEligible: true,
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
					"monster drop chances": true,
				},
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:  247,
					typestrings.PfnTotalSlots: 4,
					typestrings.PfnEquipType:  string(database.EquipTypeWeapon),
				},
			},
			DropChanceResponse: api.DropChanceResponse{
				DropChanceFinBlow:   6.5076,
				DropChanceNoFinBlow: 4.437,
				WantedEquipment: api.WantedEquipment{
					MinSlots:         4,
					MaxSlots:         4,
					ReqAbilities:     []string{},
					ReqEmptySlots:    nil,
					MinOptionalSlots: 4,
					MaxOptionalSlots: 4,
					OptionalSlotsAre: string(api.OptionalSlotAllEmpty),
				},
				EquipmentMatchChances: api.EquipmentMatchChances{
					FinBlow:   12.9643,
					NoFinBlow: 8.8393,
				},
				CharacterChances: api.CharacterChances{
					EligibleChars: []string{"tidus", "yuna", "wakka", "lulu", "rikku"},
					AnyFinBlow:    78.5714,
					AnyNoFinBlow:  53.5714,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"monster drop chances": true,
				},
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:     247,
					typestrings.PfnMinEmptySlots: 4,
					typestrings.PfnEquipType:     string(database.EquipTypeWeapon),
				},
			},
			DropChanceResponse: api.DropChanceResponse{
				DropChanceFinBlow:   6.5076,
				DropChanceNoFinBlow: 4.437,
				WantedEquipment: api.WantedEquipment{
					MinSlots:         4,
					MaxSlots:         4,
					ReqAbilities:     []string{},
					ReqEmptySlots:    h.GetInt32Ptr(4),
					MinOptionalSlots: 0,
					MaxOptionalSlots: 0,
					OptionalSlotsAre: string(api.OptionalSlotAllEmpty),
				},
				EquipmentMatchChances: api.EquipmentMatchChances{
					FinBlow:   12.9643,
					NoFinBlow: 8.8393,
				},
				CharacterChances: api.CharacterChances{
					EligibleChars: []string{"tidus", "yuna", "wakka", "lulu", "rikku"},
					AnyFinBlow:    78.5714,
					AnyNoFinBlow:  53.5714,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:     305,
					typestrings.PfnAutoAbilities: []any{129},
				},
			},
			DropChanceResponse: api.DropChanceResponse{
				DropChanceFinBlow:   50,
				DropChanceNoFinBlow: 50,
				WantedEquipment: api.WantedEquipment{
					MinSlots:         4,
					MaxSlots:         4,
					ReqAbilities:     []string{"ribbon"},
					ReqEmptySlots:    nil,
					MinOptionalSlots: 3,
					MaxOptionalSlots: 3,
					OptionalSlotsAre: string(api.OptionalSlotAllEmpty),
				},
				EquipmentMatchChances: api.EquipmentMatchChances{
					FinBlow:   50,
					NoFinBlow: 50,
				},
				CharacterChances: api.CharacterChances{
					EligibleChars: []string{"tidus", "yuna", "wakka", "lulu", "kimahri", "auron", "rikku"},
					AnyFinBlow:    100,
					AnyNoFinBlow:  100,
				},
				MonsterChances: api.MonsterDropChances{
					MonsterID:           305,
					MonsterName:         "penance",
					EquipmentDropChance: 100,
					SlotsChances: []api.ChanceAmount{
						{
							Amount: 4,
							Chance: 100,
						},
					},
					ShotsChances: []api.ChanceAmount{
						{
							Amount: 1,
							Chance: 100,
						},
					},
					AutoAbilities: api.DroppedAutoAbilities{
						Dropped: []api.EqAbilityChance{
							{
								EqAbilitySimple: api.EqAbilitySimple{
									ID:         129,
									Name:       "ribbon",
									Characters: []string{},
								},
								Chance: 100,
							},
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
				dontCheck: map[string]bool{
					"wanted equipment": true,
				},
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:     305,
					typestrings.PfnAutoAbilities: []any{51},
				},
			},
			DropChanceResponse: api.DropChanceResponse{
				DropChanceFinBlow:   50,
				DropChanceNoFinBlow: 50,
				EquipmentMatchChances: api.EquipmentMatchChances{
					FinBlow:   50,
					NoFinBlow: 50,
				},
				CharacterChances: api.CharacterChances{
					EligibleChars: []string{"tidus", "yuna", "wakka", "lulu", "kimahri", "auron", "rikku"},
					AnyFinBlow:    100,
					AnyNoFinBlow:  100,
				},
				MonsterChances: api.MonsterDropChances{
					MonsterID:           305,
					MonsterName:         "penance",
					EquipmentDropChance: 100,
					SlotsChances: []api.ChanceAmount{
						{
							Amount: 4,
							Chance: 100,
						},
					},
					ShotsChances: []api.ChanceAmount{
						{
							Amount: 1,
							Chance: 100,
						},
					},
					AutoAbilities: api.DroppedAutoAbilities{
						Dropped: []api.EqAbilityChance{
							{
								EqAbilitySimple: api.EqAbilitySimple{
									ID:         51,
									Name:       "break damage limit",
									Characters: []string{},
								},
								Chance: 100,
							},
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
				dontCheck: map[string]bool{
					"monster drop chances": true,
					"wanted equipment":     true,
				},
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:     305,
					typestrings.PfnAutoAbilities: []any{51},
					typestrings.PfnCharacterID:   1,
				},
			},
			DropChanceResponse: api.DropChanceResponse{
				DropChanceFinBlow:   17.8571,
				DropChanceNoFinBlow: 5.3571,
				EquipmentMatchChances: api.EquipmentMatchChances{
					FinBlow:   17.8571,
					NoFinBlow: 5.3571,
				},
				CharacterChances: api.CharacterChances{
					EligibleChars: []string{"tidus", "yuna", "wakka", "lulu", "kimahri", "auron", "rikku"},
					AnyFinBlow:    100,
					AnyNoFinBlow:  100,
					SpecificChar: &api.SpecificCharChances{
						FinBlow:    35.7143,
						NoFinBlow:  10.7143,
						IsEligible: true,
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
					"monster drop chances": true,
				},
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:         13,
					typestrings.PfnAutoAbilities:     []any{61},
					typestrings.PfnOptionalAbilities: true,
				},
			},
			DropChanceResponse: api.DropChanceResponse{
				DropChanceFinBlow:   2.0168,
				DropChanceNoFinBlow: 2.0168,
				WantedEquipment: api.WantedEquipment{
					MinSlots:         1,
					MaxSlots:         2,
					ReqAbilities:     []string{"water ward"},
					ReqEmptySlots:    nil,
					MinOptionalSlots: 0,
					MaxOptionalSlots: 1,
					OptionalSlotsAre: string(api.OptionalSlotAbilityOrEmpty),
				},
				EquipmentMatchChances: api.EquipmentMatchChances{
					FinBlow:   42.8571,
					NoFinBlow: 42.8571,
				},
				CharacterChances: api.CharacterChances{
					EligibleChars: []string{"tidus", "yuna", "wakka", "lulu", "kimahri", "auron", "rikku"},
					AnyFinBlow:    100,
					AnyNoFinBlow:  100,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"monster drop chances": true,
					"wanted equipment":     true,
				},
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:         13,
					typestrings.PfnAutoAbilities:     []any{62},
					typestrings.PfnOptionalAbilities: true,
				},
			},
			DropChanceResponse: api.DropChanceResponse{
				DropChanceFinBlow:   0.3361,
				DropChanceNoFinBlow: 0.3361,
				EquipmentMatchChances: api.EquipmentMatchChances{
					FinBlow:   7.1429,
					NoFinBlow: 7.1429,
				},
				CharacterChances: api.CharacterChances{
					EligibleChars: []string{"tidus", "yuna", "wakka", "lulu", "kimahri", "auron", "rikku"},
					AnyFinBlow:    100,
					AnyNoFinBlow:  100,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"monster drop chances": true,
					"wanted equipment":     true,
				},
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:         13,
					typestrings.PfnAutoAbilities:     []any{61},
					typestrings.PfnOptionalAbilities: false,
				},
			},
			DropChanceResponse: api.DropChanceResponse{
				DropChanceFinBlow:   2.0168,
				DropChanceNoFinBlow: 2.0168,
				EquipmentMatchChances: api.EquipmentMatchChances{
					FinBlow:   42.8571,
					NoFinBlow: 42.8571,
				},
				CharacterChances: api.CharacterChances{
					EligibleChars: []string{"tidus", "yuna", "wakka", "lulu", "kimahri", "auron", "rikku"},
					AnyFinBlow:    100,
					AnyNoFinBlow:  100,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"monster drop chances": true,
					"wanted equipment":     true,
				},
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:         13,
					typestrings.PfnAutoAbilities:     []any{62},
					typestrings.PfnOptionalAbilities: false,
				},
			},
			DropChanceResponse: api.DropChanceResponse{
				DropChanceFinBlow:   0.3361,
				DropChanceNoFinBlow: 0.3361,
				EquipmentMatchChances: api.EquipmentMatchChances{
					FinBlow:   7.1429,
					NoFinBlow: 7.1429,
				},
				CharacterChances: api.CharacterChances{
					EligibleChars: []string{"tidus", "yuna", "wakka", "lulu", "kimahri", "auron", "rikku"},
					AnyFinBlow:    100,
					AnyNoFinBlow:  100,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"monster drop chances": true,
					"wanted equipment":     true,
				},
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:     13,
					typestrings.PfnAutoAbilities: []any{62, 61},
				},
			},
			DropChanceResponse: api.DropChanceResponse{
				DropChanceFinBlow:   0,
				DropChanceNoFinBlow: 0,
				EquipmentMatchChances: api.EquipmentMatchChances{
					FinBlow:   0,
					NoFinBlow: 0,
				},
				CharacterChances: api.CharacterChances{
					EligibleChars: []string{"tidus", "yuna", "wakka", "lulu", "kimahri", "auron", "rikku"},
					AnyFinBlow:    100,
					AnyNoFinBlow:  100,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"monster drop chances": true,
				},
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:         256,
					typestrings.PfnAutoAbilities:     []any{113, 112},
					typestrings.PfnOptionalAbilities: false,
				},
			},
			DropChanceResponse: api.DropChanceResponse{
				DropChanceFinBlow:   0,
				DropChanceNoFinBlow: 0,
				WantedEquipment: api.WantedEquipment{
					MinSlots:         2,
					MaxSlots:         4,
					ReqAbilities:     []string{"hp +20%", "hp +10%"},
					ReqEmptySlots:    nil,
					MinOptionalSlots: 0,
					MaxOptionalSlots: 2,
					OptionalSlotsAre: string(api.OptionalSlotAllEmpty),
				},
				EquipmentMatchChances: api.EquipmentMatchChances{
					FinBlow:   0,
					NoFinBlow: 0,
				},
				CharacterChances: api.CharacterChances{
					EligibleChars: []string{},
					AnyFinBlow:    0,
					AnyNoFinBlow:  0,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"monster drop chances": true,
					"wanted equipment":     true,
				},
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:         256,
					typestrings.PfnAutoAbilities:     []any{113, 112},
					typestrings.PfnOptionalAbilities: true,
				},
			},
			DropChanceResponse: api.DropChanceResponse{
				DropChanceFinBlow:   31.7126,
				DropChanceNoFinBlow: 31.7126,
				EquipmentMatchChances: api.EquipmentMatchChances{
					FinBlow:   31.7126,
					NoFinBlow: 31.7126,
				},
				CharacterChances: api.CharacterChances{
					EligibleChars: []string{"tidus", "yuna", "wakka", "lulu", "kimahri", "auron", "rikku"},
					AnyFinBlow:    100,
					AnyNoFinBlow:  100,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"monster drop chances": true,
				},
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:         191,
					typestrings.PfnMinEmptySlots:     4,
					typestrings.PfnEquipType:         string(database.EquipTypeWeapon),
					typestrings.PfnOptionalAbilities: false,
				},
			},
			DropChanceResponse: api.DropChanceResponse{
				DropChanceFinBlow:   0.2311,
				DropChanceNoFinBlow: 0.1576,
				WantedEquipment: api.WantedEquipment{
					MinSlots:         4,
					MaxSlots:         4,
					ReqAbilities:     []string{},
					ReqEmptySlots:    h.GetInt32Ptr(4),
					MinOptionalSlots: 0,
					MaxOptionalSlots: 0,
					OptionalSlotsAre: string(api.OptionalSlotAllEmpty),
				},
				EquipmentMatchChances: api.EquipmentMatchChances{
					FinBlow:   1.9643,
					NoFinBlow: 1.3393,
				},
				CharacterChances: api.CharacterChances{
					EligibleChars: []string{"tidus", "yuna", "wakka", "lulu", "rikku"},
					AnyFinBlow:    78.5714,
					AnyNoFinBlow:  53.5714,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"monster drop chances": true,
				},
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:         191,
					typestrings.PfnMinEmptySlots:     4,
					typestrings.PfnEquipType:         string(database.EquipTypeWeapon),
					typestrings.PfnOptionalAbilities: true,
				},
			},
			DropChanceResponse: api.DropChanceResponse{
				DropChanceFinBlow:   0.2311,
				DropChanceNoFinBlow: 0.1576,
				WantedEquipment: api.WantedEquipment{
					MinSlots:         4,
					MaxSlots:         4,
					ReqAbilities:     []string{},
					ReqEmptySlots:    h.GetInt32Ptr(4),
					MinOptionalSlots: 0,
					MaxOptionalSlots: 0,
					OptionalSlotsAre: string(api.OptionalSlotAbilityOrEmpty),
				},
				EquipmentMatchChances: api.EquipmentMatchChances{
					FinBlow:   1.9643,
					NoFinBlow: 1.3393,
				},
				CharacterChances: api.CharacterChances{
					EligibleChars: []string{"tidus", "yuna", "wakka", "lulu", "rikku"},
					AnyFinBlow:    78.5714,
					AnyNoFinBlow:  53.5714,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"monster drop chances": true,
				},
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:         191,
					typestrings.PfnTotalSlots:        4,
					typestrings.PfnEquipType:         string(database.EquipTypeWeapon),
					typestrings.PfnOptionalAbilities: false,
				},
			},
			DropChanceResponse: api.DropChanceResponse{
				DropChanceFinBlow:   0.2311,
				DropChanceNoFinBlow: 0.1576,
				WantedEquipment: api.WantedEquipment{
					MinSlots:         4,
					MaxSlots:         4,
					ReqAbilities:     []string{},
					ReqEmptySlots:    nil,
					MinOptionalSlots: 4,
					MaxOptionalSlots: 4,
					OptionalSlotsAre: string(api.OptionalSlotAllEmpty),
				},
				EquipmentMatchChances: api.EquipmentMatchChances{
					FinBlow:   1.9643,
					NoFinBlow: 1.3393,
				},
				CharacterChances: api.CharacterChances{
					EligibleChars: []string{"tidus", "yuna", "wakka", "lulu", "rikku"},
					AnyFinBlow:    78.5714,
					AnyNoFinBlow:  53.5714,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"monster drop chances": true,
				},
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID:         191,
					typestrings.PfnTotalSlots:        4,
					typestrings.PfnEquipType:         string(database.EquipTypeWeapon),
					typestrings.PfnOptionalAbilities: true,
				},
			},
			DropChanceResponse: api.DropChanceResponse{
				DropChanceFinBlow:   0.5882,
				DropChanceNoFinBlow: 0.5882,
				WantedEquipment: api.WantedEquipment{
					MinSlots:         4,
					MaxSlots:         4,
					ReqAbilities:     []string{},
					ReqEmptySlots:    nil,
					MinOptionalSlots: 4,
					MaxOptionalSlots: 4,
					OptionalSlotsAre: string(api.OptionalSlotAbilityOrEmpty),
				},
				EquipmentMatchChances: api.EquipmentMatchChances{
					FinBlow:   5,
					NoFinBlow: 5,
				},
				CharacterChances: api.CharacterChances{
					EligibleChars: []string{"tidus", "yuna", "wakka", "lulu", "kimahri", "auron", "rikku"},
					AnyFinBlow:    100,
					AnyNoFinBlow:  100,
				},
			},
		},
	}

	testServiceResponsesPost(t, tests, "DropChance", testCfg.HandleDropChancePost, compareDropChanceResponses)
}
