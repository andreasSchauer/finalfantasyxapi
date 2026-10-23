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
				method:         http.MethodPost,
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
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "id '10' of field 'character_id' is out of range. allowed range: 1 to 7.",
				requestBody: map[FieldName]any{
					pfnMonsterID:     13,
					pfnCharacterID:   10,
					pfnMinEmptySlots: 1,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "id '0' of field 'character_id' is out of range. allowed range: 1 to 7.",
				requestBody: map[FieldName]any{
					pfnMonsterID:     13,
					pfnCharacterID:   0,
					pfnMinEmptySlots: 1,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "array of field 'party_members' can't have duplicate ids.",
				requestBody: map[FieldName]any{
					pfnMonsterID:     13,
					pfnPartyMembers:  []any{2, 2, 2, 2, 2},
					pfnMinEmptySlots: 1,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "character must be in the party to be eligible for receiving equipment.",
				requestBody: map[FieldName]any{
					pfnMonsterID:     13,
					pfnCharacterID:   4,
					pfnPartyMembers:  []any{1, 2, 3},
					pfnMinEmptySlots: 1,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "too many items in array of field 'party_members'. got: 8, maximum allowed: 7.",
				requestBody: map[FieldName]any{
					pfnMonsterID:     13,
					pfnCharacterID:   4,
					pfnPartyMembers:  []any{1, 2, 3, 4, 5, 6, 7, 8},
					pfnMinEmptySlots: 1,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "index 0 of field 'auto_abilities': id '131' is out of range. allowed range: 1 to 130.",
				requestBody: map[FieldName]any{
					pfnMonsterID:     13,
					pfnAutoAbilities: []any{131},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "too many items in array of field 'auto_abilities'. got: 5, maximum allowed: 4.",
				requestBody: map[FieldName]any{
					pfnMonsterID:     13,
					pfnAutoAbilities: []any{13, 14, 15, 16, 17},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "value '5' of field 'min_empty_slots' is out of range. allowed range: 0 to 4.",
				requestBody: map[FieldName]any{
					pfnMonsterID:     13,
					pfnMinEmptySlots: 5,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "value '5' of field 'total_slots' is out of range. allowed range: 0 to 4.",
				requestBody: map[FieldName]any{
					pfnMonsterID:  13,
					pfnTotalSlots: 5,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "value '-1' of field 'min_empty_slots' is out of range. allowed range: 0 to 4.",
				requestBody: map[FieldName]any{
					pfnMonsterID:     13,
					pfnMinEmptySlots: -1,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "can't discern equipment's type. use 'auto_abilities' to discern the type, or set it manually by using 'equip_type'.",
				requestBody: map[FieldName]any{
					pfnMonsterID:     13,
					pfnMinEmptySlots: 2,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "the amount of auto-abilities and empty slots combined (1) can't exceed the given total amount of slots (0).",
				requestBody: map[FieldName]any{
					pfnMonsterID:     13,
					pfnAutoAbilities: []any{2},
					pfnTotalSlots:    0,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "the amount of auto-abilities and empty slots combined can't exceed 4.",
				requestBody: map[FieldName]any{
					pfnMonsterID:     13,
					pfnAutoAbilities: []any{2, 7, 22},
					pfnMinEmptySlots: 2,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "weapon- and armor-abilities can't be combined.",
				requestBody: map[FieldName]any{
					pfnMonsterID:     13,
					pfnAutoAbilities: []any{2, 129},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "weapon- and armor-abilities can't be combined.",
				requestBody: map[FieldName]any{
					pfnMonsterID:     13,
					pfnAutoAbilities: []any{2},
					pfnEquipType:     string(database.EquipTypeArmor),
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "weapon- and armor-abilities can't be combined.",
				requestBody: map[FieldName]any{
					pfnMonsterID:     13,
					pfnAutoAbilities: []any{129},
					pfnEquipType:     string(database.EquipTypeWeapon),
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "value 'weaponf' of field 'equip_type' is not a valid enum value. allowed values: 'weapon', 'armor'.",
				requestBody: map[FieldName]any{
					pfnMonsterID:     13,
					pfnMinEmptySlots: 1,
					pfnEquipType:     "weaponf",
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "monster 'sinscale - 1 (dream zanarkand)' doesn't drop equipment.",
				requestBody: map[FieldName]any{
					pfnMonsterID:     1,
					pfnAutoAbilities: []any{2},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "water flan - 1 (tutorial) doesn't drop auto-ability 'zombiestrike' (22). it only drops the following auto-abilities: 'piercing' (2), 'waterstrike' (7), 'distill mana' (26), 'water ward' (61), 'waterproof' (62).",
				requestBody: map[FieldName]any{
					pfnMonsterID:     13,
					pfnAutoAbilities: []any{7, 22},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "kottos doesn't drop auto-ability 'piercing' (2). it only drops the following auto-abilities: 'counterattack' (37), 'evade & counter' (38), 'magic counter' (39), 'hp +30%' (114), 'hp +20%' (113), 'hp +10%' (112).",
				requestBody: map[FieldName]any{
					pfnMonsterID:     258,
					pfnAutoAbilities: []any{2},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "at least one of these fields must have a value: 'min_empty_slots', 'total_slots', 'auto_abilities'.",
				requestBody: map[FieldName]any{
					pfnMonsterID: 13,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
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
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				requestBody: map[FieldName]any{
					pfnMonsterID:     27,
					pfnPartyMembers:  []any{1, 2, 3, 4, 5},
					pfnAutoAbilities: []any{2},
					pfnTotalSlots:    1,
				},
			},
			DropChanceResponse: DropChanceResponse{
				DropChanceFinBlow:   0.7059,
				DropChanceNoFinBlow: 0.2647,
				WantedEquipment: WantedEquipment{
					MinSlots:         1,
					MaxSlots:         1,
					ReqAbilities:     []string{"piercing"},
					ReqEmptySlots:    nil,
					MinOptionalSlots: 0,
					MaxOptionalSlots: 0,
					OptionalSlotsAre: string(optionalSlotAllEmpty),
				},
				EquipmentMatchChances: EquipmentMatchChances{
					FinBlow:   15,
					NoFinBlow: 5.625,
				},
				CharacterChances: CharacterChances{
					EligibleChars: []string{"kimahri"},
					AnyFinBlow:    40,
					AnyNoFinBlow:  15,
				},
				MonsterChances: MonsterDropChances{
					MonsterID:           27,
					MonsterName:         "yellow element",
					EquipmentDropChance: 4.7059,
					SlotsChances: []ChanceAmount{
						{
							Amount: 1,
							Chance: 75,
						},
						{
							Amount: 2,
							Chance: 25,
						},
					},
					ShotsChances: []ChanceAmount{
						{
							Amount: 0,
							Chance: 50,
						},
						{
							Amount: 1,
							Chance: 50,
						},
					},
					AutoAbilities: DroppedAutoAbilities{
						Forced: &EqAbilitySimple{
							ID:         2,
							Name:       "piercing",
							Characters: []string{"kimahri", "auron"},
						},
						Dropped: []EqAbilityChance{
							{
								EqAbilitySimple: EqAbilitySimple{
									ID:         6,
									Name:       "lightningstrike",
									Characters: []string{},
								},
								Chance: 57.1429,
							},
							{
								EqAbilitySimple: EqAbilitySimple{
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
				requestBody: map[FieldName]any{
					pfnMonsterID:     27,
					pfnAutoAbilities: []any{2},
					pfnCharacterID:   5,
				},
			},
			DropChanceResponse: DropChanceResponse{
				DropChanceFinBlow:   0.7353,
				DropChanceNoFinBlow: 0.2206,
				WantedEquipment: WantedEquipment{
					MinSlots:         1,
					MaxSlots:         2,
					ReqAbilities:     []string{"piercing"},
					ReqEmptySlots:    nil,
					MinOptionalSlots: 0,
					MaxOptionalSlots: 1,
					OptionalSlotsAre: string(optionalSlotAllEmpty),
				},
				EquipmentMatchChances: EquipmentMatchChances{
					FinBlow:   15.625,
					NoFinBlow: 4.6875,
				},
				CharacterChances: CharacterChances{
					EligibleChars: []string{"kimahri", "auron"},
					AnyFinBlow:    46.4286,
					AnyNoFinBlow:  21.4286,
					SpecificChar: &SpecificCharChances{
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
				requestBody: map[FieldName]any{
					pfnMonsterID:     50,
					pfnAutoAbilities: []any{9},
					pfnCharacterID:   5,
					pfnDecimals:      true,
				},
			},
			DropChanceResponse: DropChanceResponse{
				DropChanceFinBlow:   0,
				DropChanceNoFinBlow: 0,
				WantedEquipment: WantedEquipment{
					MinSlots:         1,
					MaxSlots:         2,
					ReqAbilities:     []string{"darktouch"},
					ReqEmptySlots:    nil,
					MinOptionalSlots: 0,
					MaxOptionalSlots: 1,
					OptionalSlotsAre: string(optionalSlotAllEmpty),
				},
				EquipmentMatchChances: EquipmentMatchChances{
					FinBlow:   0,
					NoFinBlow: 0,
				},
				CharacterChances: CharacterChances{
					EligibleChars: []string{"tidus", "yuna", "wakka", "lulu", "rikku"},
					AnyFinBlow:    0.785714,
					AnyNoFinBlow:  0.535714,
					SpecificChar: &SpecificCharChances{
						FinBlow:    0.357143,
						NoFinBlow:  0.107143,
						IsEligible: false,
					},
				},
				MonsterChances: MonsterDropChances{
					MonsterID:           50,
					MonsterName:         "garuda - 3 (regular encounter)",
					EquipmentDropChance: 0.501961,
					SlotsChances: []ChanceAmount{
						{
							Amount: 1,
							Chance: 0.5,
						},
						{
							Amount: 2,
							Chance: 0.5,
						},
					},
					ShotsChances: []ChanceAmount{
						{
							Amount: 1,
							Chance: 0.75,
						},
						{
							Amount: 2,
							Chance: 0.25,
						},
					},
					AutoAbilities: DroppedAutoAbilities{
						Forced: &EqAbilitySimple{
							ID:         2,
							Name:       "piercing",
							Characters: []string{"kimahri", "auron"},
						},
						Dropped: []EqAbilityChance{
							{
								EqAbilitySimple: EqAbilitySimple{
									ID:         9,
									Name:       "darktouch",
									Characters: []string{},
								},
								Chance: 0.285714,
							},
							{
								EqAbilitySimple: EqAbilitySimple{
									ID:         29,
									Name:       "strength +3%",
									Characters: []string{"tidus", "wakka", "kimahri", "auron", "rikku"},
								},
								Chance: 0.714286,
							},
							{
								EqAbilitySimple: EqAbilitySimple{
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
				requestBody: map[FieldName]any{
					pfnMonsterID:     50,
					pfnAutoAbilities: []any{9},
					pfnCharacterID:   5,
				},
			},
			DropChanceResponse: DropChanceResponse{
				DropChanceFinBlow:   0,
				DropChanceNoFinBlow: 0,
				EquipmentMatchChances: EquipmentMatchChances{
					FinBlow:   0,
					NoFinBlow: 0,
				},
				CharacterChances: CharacterChances{
					EligibleChars: []string{"tidus", "yuna", "wakka", "lulu", "rikku"},
					AnyFinBlow:    78.5714,
					AnyNoFinBlow:  53.5714,
					SpecificChar: &SpecificCharChances{
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
				requestBody: map[FieldName]any{
					pfnMonsterID:         50,
					pfnAutoAbilities:     []any{9},
					pfnOptionalAbilities: true,
					pfnCharacterID:       5,
				},
			},
			DropChanceResponse: DropChanceResponse{
				DropChanceFinBlow:   1.2805,
				DropChanceNoFinBlow: 0.3842,
				WantedEquipment: WantedEquipment{
					MinSlots:         1,
					MaxSlots:         2,
					ReqAbilities:     []string{"darktouch"},
					ReqEmptySlots:    nil,
					MinOptionalSlots: 0,
					MaxOptionalSlots: 1,
					OptionalSlotsAre: string(optionalSlotAbilityOrEmpty),
				},
				EquipmentMatchChances: EquipmentMatchChances{
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
				requestBody: map[FieldName]any{
					pfnMonsterID:     50,
					pfnAutoAbilities: []any{9},
					pfnMinEmptySlots: 1,
					pfnCharacterID:   5,
				},
			},
			DropChanceResponse: DropChanceResponse{
				DropChanceFinBlow:   0,
				DropChanceNoFinBlow: 0,
				EquipmentMatchChances: EquipmentMatchChances{
					FinBlow:   0,
					NoFinBlow: 0,
				},
				CharacterChances: CharacterChances{
					EligibleChars: []string{"tidus", "yuna", "wakka", "lulu", "rikku"},
					AnyFinBlow:    78.5714,
					AnyNoFinBlow:  53.5714,
					SpecificChar: &SpecificCharChances{
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
				requestBody: map[FieldName]any{
					pfnMonsterID:     50,
					pfnAutoAbilities: []any{9},
					pfnTotalSlots:    2,
					pfnCharacterID:   5,
				},
			},
			DropChanceResponse: DropChanceResponse{
				DropChanceFinBlow:   0,
				DropChanceNoFinBlow: 0,
				EquipmentMatchChances: EquipmentMatchChances{
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
				requestBody: map[FieldName]any{
					pfnMonsterID:         50,
					pfnAutoAbilities:     []any{2},
					pfnOptionalAbilities: true,
					pfnTotalSlots:        2,
					pfnCharacterID:       5,
				},
			},
			DropChanceResponse: DropChanceResponse{
				DropChanceFinBlow:   4.4818,
				DropChanceNoFinBlow: 1.3445,
				EquipmentMatchChances: EquipmentMatchChances{
					FinBlow:   8.9286,
					NoFinBlow: 2.6786,
				},
				CharacterChances: CharacterChances{
					EligibleChars: []string{"kimahri", "auron"},
					AnyFinBlow:    46.4286,
					AnyNoFinBlow:  21.4286,
					SpecificChar: &SpecificCharChances{
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
				requestBody: map[FieldName]any{
					pfnMonsterID:         258,
					pfnAutoAbilities:     []any{38},
					pfnOptionalAbilities: true,
				},
			},
			DropChanceResponse: DropChanceResponse{
				DropChanceFinBlow:   0,
				DropChanceNoFinBlow: 0,
				WantedEquipment: WantedEquipment{
					MinSlots:         2,
					MaxSlots:         3,
					ReqAbilities:     []string{"evade & counter"},
					ReqEmptySlots:    nil,
					MinOptionalSlots: 1,
					MaxOptionalSlots: 2,
					OptionalSlotsAre: string(optionalSlotAbilityOrEmpty),
				},
				EquipmentMatchChances: EquipmentMatchChances{
					FinBlow:   0,
					NoFinBlow: 0,
				},
				CharacterChances: CharacterChances{
					EligibleChars: []string{"tidus", "yuna", "wakka", "lulu", "kimahri", "auron", "rikku"},
					AnyFinBlow:    100,
					AnyNoFinBlow:  100,
				},
				MonsterChances: MonsterDropChances{
					MonsterID:           258,
					MonsterName:         "kottos",
					EquipmentDropChance: 100,
					SlotsChances: []ChanceAmount{
						{
							Amount: 2,
							Chance: 50,
						},
						{
							Amount: 3,
							Chance: 50,
						},
					},
					ShotsChances: []ChanceAmount{
						{
							Amount: 4,
							Chance: 50,
						},
						{
							Amount: 5,
							Chance: 50,
						},
					},
					AutoAbilities: DroppedAutoAbilities{
						Forced: &EqAbilitySimple{
							ID:         37,
							Name:       "counterattack",
							Characters: []string{},
						},
						Dropped: []EqAbilityChance{
							{
								EqAbilitySimple: EqAbilitySimple{
									ID:         37,
									Name:       "counterattack",
									Characters: []string{},
								},
								Chance: 28.5714,
							},
							{
								EqAbilitySimple: EqAbilitySimple{
									ID:         38,
									Name:       "evade & counter",
									Characters: []string{},
								},
								Chance: 42.8571,
							},
							{
								EqAbilitySimple: EqAbilitySimple{
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
				requestBody: map[FieldName]any{
					pfnMonsterID:         258,
					pfnAutoAbilities:     []any{38},
					pfnOptionalAbilities: false,
				},
			},
			DropChanceResponse: DropChanceResponse{
				DropChanceFinBlow:   0,
				DropChanceNoFinBlow: 0,
				EquipmentMatchChances: EquipmentMatchChances{
					FinBlow:   0,
					NoFinBlow: 0,
				},
				CharacterChances: CharacterChances{
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
				requestBody: map[FieldName]any{
					pfnMonsterID:         258,
					pfnAutoAbilities:     []any{37},
					pfnOptionalAbilities: true,
				},
			},
			DropChanceResponse: DropChanceResponse{
				DropChanceFinBlow:   50,
				DropChanceNoFinBlow: 50,
				EquipmentMatchChances: EquipmentMatchChances{
					FinBlow:   50,
					NoFinBlow: 50,
				},
				CharacterChances: CharacterChances{
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
				requestBody: map[FieldName]any{
					pfnMonsterID:     243,
					pfnAutoAbilities: []any{30, 19},
				},
			},
			DropChanceResponse: DropChanceResponse{
				DropChanceFinBlow:   0.2744,
				DropChanceNoFinBlow: 0.1543,
				EquipmentMatchChances: EquipmentMatchChances{
					FinBlow:   1.1662,
					NoFinBlow: 0.656,
				},
				CharacterChances: CharacterChances{
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
				requestBody: map[FieldName]any{
					pfnMonsterID:     243,
					pfnPartyMembers:  []any{1, 3, 7},
					pfnAutoAbilities: []any{30, 19},
				},
			},
			DropChanceResponse: DropChanceResponse{
				DropChanceFinBlow:   0.4802,
				DropChanceNoFinBlow: 0.4802,
				EquipmentMatchChances: EquipmentMatchChances{
					FinBlow:   2.0408,
					NoFinBlow: 2.0408,
				},
				CharacterChances: CharacterChances{
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
				requestBody: map[FieldName]any{
					pfnMonsterID:     243,
					pfnAutoAbilities: []any{34, 19},
				},
			},
			DropChanceResponse: DropChanceResponse{
				DropChanceFinBlow:   0.2229,
				DropChanceNoFinBlow: 0.1029,
				EquipmentMatchChances: EquipmentMatchChances{
					FinBlow:   0.9475,
					NoFinBlow: 0.4373,
				},
				CharacterChances: CharacterChances{
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
				requestBody: map[FieldName]any{
					pfnMonsterID:         243,
					pfnAutoAbilities:     []any{30, 19},
					pfnCharacterID:       5,
					pfnOptionalAbilities: false,
				},
			},
			DropChanceResponse: DropChanceResponse{
				DropChanceFinBlow:   0,
				DropChanceNoFinBlow: 0,
				EquipmentMatchChances: EquipmentMatchChances{
					FinBlow:   0,
					NoFinBlow: 0,
				},
				CharacterChances: CharacterChances{
					EligibleChars: []string{"tidus", "wakka", "rikku"},
					AnyFinBlow:    57.1429,
					AnyNoFinBlow:  32.1429,
					SpecificChar: &SpecificCharChances{
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
				requestBody: map[FieldName]any{
					pfnMonsterID:         243,
					pfnAutoAbilities:     []any{30, 19},
					pfnCharacterID:       3,
					pfnOptionalAbilities: false,
				},
			},
			DropChanceResponse: DropChanceResponse{
				DropChanceFinBlow:   0.1715,
				DropChanceNoFinBlow: 0.0514,
				EquipmentMatchChances: EquipmentMatchChances{
					FinBlow:   0.7289,
					NoFinBlow: 0.2187,
				},
				CharacterChances: CharacterChances{
					EligibleChars: []string{"tidus", "wakka", "rikku"},
					AnyFinBlow:    57.1429,
					AnyNoFinBlow:  32.1429,
					SpecificChar: &SpecificCharChances{
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
				requestBody: map[FieldName]any{
					pfnMonsterID:  247,
					pfnTotalSlots: 4,
					pfnEquipType:  string(database.EquipTypeWeapon),
				},
			},
			DropChanceResponse: DropChanceResponse{
				DropChanceFinBlow:   6.5076,
				DropChanceNoFinBlow: 4.437,
				WantedEquipment: WantedEquipment{
					MinSlots:         4,
					MaxSlots:         4,
					ReqAbilities:     []string{},
					ReqEmptySlots:    nil,
					MinOptionalSlots: 4,
					MaxOptionalSlots: 4,
					OptionalSlotsAre: string(optionalSlotAllEmpty),
				},
				EquipmentMatchChances: EquipmentMatchChances{
					FinBlow:   12.9643,
					NoFinBlow: 8.8393,
				},
				CharacterChances: CharacterChances{
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
				requestBody: map[FieldName]any{
					pfnMonsterID:     247,
					pfnMinEmptySlots: 4,
					pfnEquipType:     string(database.EquipTypeWeapon),
				},
			},
			DropChanceResponse: DropChanceResponse{
				DropChanceFinBlow:   6.5076,
				DropChanceNoFinBlow: 4.437,
				WantedEquipment: WantedEquipment{
					MinSlots:         4,
					MaxSlots:         4,
					ReqAbilities:     []string{},
					ReqEmptySlots:    h.GetInt32Ptr(4),
					MinOptionalSlots: 0,
					MaxOptionalSlots: 0,
					OptionalSlotsAre: string(optionalSlotAllEmpty),
				},
				EquipmentMatchChances: EquipmentMatchChances{
					FinBlow:   12.9643,
					NoFinBlow: 8.8393,
				},
				CharacterChances: CharacterChances{
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
				requestBody: map[FieldName]any{
					pfnMonsterID:     305,
					pfnAutoAbilities: []any{129},
				},
			},
			DropChanceResponse: DropChanceResponse{
				DropChanceFinBlow:   50,
				DropChanceNoFinBlow: 50,
				WantedEquipment: WantedEquipment{
					MinSlots:         4,
					MaxSlots:         4,
					ReqAbilities:     []string{"ribbon"},
					ReqEmptySlots:    nil,
					MinOptionalSlots: 3,
					MaxOptionalSlots: 3,
					OptionalSlotsAre: string(optionalSlotAllEmpty),
				},
				EquipmentMatchChances: EquipmentMatchChances{
					FinBlow:   50,
					NoFinBlow: 50,
				},
				CharacterChances: CharacterChances{
					EligibleChars: []string{"tidus", "yuna", "wakka", "lulu", "kimahri", "auron", "rikku"},
					AnyFinBlow:    100,
					AnyNoFinBlow:  100,
				},
				MonsterChances: MonsterDropChances{
					MonsterID:           305,
					MonsterName:         "penance",
					EquipmentDropChance: 100,
					SlotsChances: []ChanceAmount{
						{
							Amount: 4,
							Chance: 100,
						},
					},
					ShotsChances: []ChanceAmount{
						{
							Amount: 1,
							Chance: 100,
						},
					},
					AutoAbilities: DroppedAutoAbilities{
						Dropped: []EqAbilityChance{
							{
								EqAbilitySimple: EqAbilitySimple{
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
				requestBody: map[FieldName]any{
					pfnMonsterID:     305,
					pfnAutoAbilities: []any{51},
				},
			},
			DropChanceResponse: DropChanceResponse{
				DropChanceFinBlow:   50,
				DropChanceNoFinBlow: 50,
				EquipmentMatchChances: EquipmentMatchChances{
					FinBlow:   50,
					NoFinBlow: 50,
				},
				CharacterChances: CharacterChances{
					EligibleChars: []string{"tidus", "yuna", "wakka", "lulu", "kimahri", "auron", "rikku"},
					AnyFinBlow:    100,
					AnyNoFinBlow:  100,
				},
				MonsterChances: MonsterDropChances{
					MonsterID:           305,
					MonsterName:         "penance",
					EquipmentDropChance: 100,
					SlotsChances: []ChanceAmount{
						{
							Amount: 4,
							Chance: 100,
						},
					},
					ShotsChances: []ChanceAmount{
						{
							Amount: 1,
							Chance: 100,
						},
					},
					AutoAbilities: DroppedAutoAbilities{
						Dropped: []EqAbilityChance{
							{
								EqAbilitySimple: EqAbilitySimple{
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
				requestBody: map[FieldName]any{
					pfnMonsterID:     305,
					pfnAutoAbilities: []any{51},
					pfnCharacterID:   1,
				},
			},
			DropChanceResponse: DropChanceResponse{
				DropChanceFinBlow:   17.8571,
				DropChanceNoFinBlow: 5.3571,
				EquipmentMatchChances: EquipmentMatchChances{
					FinBlow:   17.8571,
					NoFinBlow: 5.3571,
				},
				CharacterChances: CharacterChances{
					EligibleChars: []string{"tidus", "yuna", "wakka", "lulu", "kimahri", "auron", "rikku"},
					AnyFinBlow:    100,
					AnyNoFinBlow:  100,
					SpecificChar: &SpecificCharChances{
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
				requestBody: map[FieldName]any{
					pfnMonsterID:         13,
					pfnAutoAbilities:     []any{61},
					pfnOptionalAbilities: true,
				},
			},
			DropChanceResponse: DropChanceResponse{
				DropChanceFinBlow:   2.0168,
				DropChanceNoFinBlow: 2.0168,
				WantedEquipment: WantedEquipment{
					MinSlots:         1,
					MaxSlots:         2,
					ReqAbilities:     []string{"water ward"},
					ReqEmptySlots:    nil,
					MinOptionalSlots: 0,
					MaxOptionalSlots: 1,
					OptionalSlotsAre: string(optionalSlotAbilityOrEmpty),
				},
				EquipmentMatchChances: EquipmentMatchChances{
					FinBlow:   42.8571,
					NoFinBlow: 42.8571,
				},
				CharacterChances: CharacterChances{
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
				requestBody: map[FieldName]any{
					pfnMonsterID:         13,
					pfnAutoAbilities:     []any{62},
					pfnOptionalAbilities: true,
				},
			},
			DropChanceResponse: DropChanceResponse{
				DropChanceFinBlow:   0.3361,
				DropChanceNoFinBlow: 0.3361,
				EquipmentMatchChances: EquipmentMatchChances{
					FinBlow:   7.1429,
					NoFinBlow: 7.1429,
				},
				CharacterChances: CharacterChances{
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
				requestBody: map[FieldName]any{
					pfnMonsterID:         13,
					pfnAutoAbilities:     []any{61},
					pfnOptionalAbilities: false,
				},
			},
			DropChanceResponse: DropChanceResponse{
				DropChanceFinBlow:   2.0168,
				DropChanceNoFinBlow: 2.0168,
				EquipmentMatchChances: EquipmentMatchChances{
					FinBlow:   42.8571,
					NoFinBlow: 42.8571,
				},
				CharacterChances: CharacterChances{
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
				requestBody: map[FieldName]any{
					pfnMonsterID:         13,
					pfnAutoAbilities:     []any{62},
					pfnOptionalAbilities: false,
				},
			},
			DropChanceResponse: DropChanceResponse{
				DropChanceFinBlow:   0.3361,
				DropChanceNoFinBlow: 0.3361,
				EquipmentMatchChances: EquipmentMatchChances{
					FinBlow:   7.1429,
					NoFinBlow: 7.1429,
				},
				CharacterChances: CharacterChances{
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
				requestBody: map[FieldName]any{
					pfnMonsterID:     13,
					pfnAutoAbilities: []any{62, 61},
				},
			},
			DropChanceResponse: DropChanceResponse{
				DropChanceFinBlow:   0,
				DropChanceNoFinBlow: 0,
				EquipmentMatchChances: EquipmentMatchChances{
					FinBlow:   0,
					NoFinBlow: 0,
				},
				CharacterChances: CharacterChances{
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
				requestBody: map[FieldName]any{
					pfnMonsterID:         256,
					pfnAutoAbilities:     []any{113, 112},
					pfnOptionalAbilities: false,
				},
			},
			DropChanceResponse: DropChanceResponse{
				DropChanceFinBlow:   0,
				DropChanceNoFinBlow: 0,
				WantedEquipment: WantedEquipment{
					MinSlots:         2,
					MaxSlots:         4,
					ReqAbilities:     []string{"hp +20%", "hp +10%"},
					ReqEmptySlots:    nil,
					MinOptionalSlots: 0,
					MaxOptionalSlots: 2,
					OptionalSlotsAre: string(optionalSlotAllEmpty),
				},
				EquipmentMatchChances: EquipmentMatchChances{
					FinBlow:   0,
					NoFinBlow: 0,
				},
				CharacterChances: CharacterChances{
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
				requestBody: map[FieldName]any{
					pfnMonsterID:         256,
					pfnAutoAbilities:     []any{113, 112},
					pfnOptionalAbilities: true,
				},
			},
			DropChanceResponse: DropChanceResponse{
				DropChanceFinBlow:   31.7126,
				DropChanceNoFinBlow: 31.7126,
				EquipmentMatchChances: EquipmentMatchChances{
					FinBlow:   31.7126,
					NoFinBlow: 31.7126,
				},
				CharacterChances: CharacterChances{
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
				requestBody: map[FieldName]any{
					pfnMonsterID:         191,
					pfnMinEmptySlots:     4,
					pfnEquipType:         string(database.EquipTypeWeapon),
					pfnOptionalAbilities: false,
				},
			},
			DropChanceResponse: DropChanceResponse{
				DropChanceFinBlow:   0.2311,
				DropChanceNoFinBlow: 0.1576,
				WantedEquipment: WantedEquipment{
					MinSlots:         4,
					MaxSlots:         4,
					ReqAbilities:     []string{},
					ReqEmptySlots:    h.GetInt32Ptr(4),
					MinOptionalSlots: 0,
					MaxOptionalSlots: 0,
					OptionalSlotsAre: string(optionalSlotAllEmpty),
				},
				EquipmentMatchChances: EquipmentMatchChances{
					FinBlow:   1.9643,
					NoFinBlow: 1.3393,
				},
				CharacterChances: CharacterChances{
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
				requestBody: map[FieldName]any{
					pfnMonsterID:         191,
					pfnMinEmptySlots:     4,
					pfnEquipType:         string(database.EquipTypeWeapon),
					pfnOptionalAbilities: true,
				},
			},
			DropChanceResponse: DropChanceResponse{
				DropChanceFinBlow:   0.2311,
				DropChanceNoFinBlow: 0.1576,
				WantedEquipment: WantedEquipment{
					MinSlots:         4,
					MaxSlots:         4,
					ReqAbilities:     []string{},
					ReqEmptySlots:    h.GetInt32Ptr(4),
					MinOptionalSlots: 0,
					MaxOptionalSlots: 0,
					OptionalSlotsAre: string(optionalSlotAbilityOrEmpty),
				},
				EquipmentMatchChances: EquipmentMatchChances{
					FinBlow:   1.9643,
					NoFinBlow: 1.3393,
				},
				CharacterChances: CharacterChances{
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
				requestBody: map[FieldName]any{
					pfnMonsterID:         191,
					pfnTotalSlots:        4,
					pfnEquipType:         string(database.EquipTypeWeapon),
					pfnOptionalAbilities: false,
				},
			},
			DropChanceResponse: DropChanceResponse{
				DropChanceFinBlow:   0.2311,
				DropChanceNoFinBlow: 0.1576,
				WantedEquipment: WantedEquipment{
					MinSlots:         4,
					MaxSlots:         4,
					ReqAbilities:     []string{},
					ReqEmptySlots:    nil,
					MinOptionalSlots: 4,
					MaxOptionalSlots: 4,
					OptionalSlotsAre: string(optionalSlotAllEmpty),
				},
				EquipmentMatchChances: EquipmentMatchChances{
					FinBlow:   1.9643,
					NoFinBlow: 1.3393,
				},
				CharacterChances: CharacterChances{
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
				requestBody: map[FieldName]any{
					pfnMonsterID:         191,
					pfnTotalSlots:        4,
					pfnEquipType:         string(database.EquipTypeWeapon),
					pfnOptionalAbilities: true,
				},
			},
			DropChanceResponse: DropChanceResponse{
				DropChanceFinBlow:   0.5882,
				DropChanceNoFinBlow: 0.5882,
				WantedEquipment: WantedEquipment{
					MinSlots:         4,
					MaxSlots:         4,
					ReqAbilities:     []string{},
					ReqEmptySlots:    nil,
					MinOptionalSlots: 4,
					MaxOptionalSlots: 4,
					OptionalSlotsAre: string(optionalSlotAbilityOrEmpty),
				},
				EquipmentMatchChances: EquipmentMatchChances{
					FinBlow:   5,
					NoFinBlow: 5,
				},
				CharacterChances: CharacterChances{
					EligibleChars: []string{"tidus", "yuna", "wakka", "lulu", "kimahri", "auron", "rikku"},
					AnyFinBlow:    100,
					AnyNoFinBlow:  100,
				},
			},
		},
	}

	testServiceResponses(t, tests, "DropChance", testCfg.HandleDropChanceGet, testCfg.HandleDropChancePost, compareDropChanceResponses)
}
