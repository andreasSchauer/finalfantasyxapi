package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apitesttypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/sections/sectiontypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

func TestGetAbility(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.Ability{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/abilities/841",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "ability with provided id '841' doesn't exist. max id: 840",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/abilities/736",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"monsters":            2,
					"battle interactions": 2,
					"battle interactions 1 inflicted status conditions": 1,
				},
			},
			NameVer:          apitesttypes.NewExpNameVer(736, "self destruct", 1),
			Rank:             h.GetInt32Ptr(3),
			AppearsInHelpBar: true,
			CanCopyCat:       false,
			AbilityType:      database.AbilityTypeEnemyAbility,
			TypedAbility:     "/enemy-abilities/339",
			Monsters:         []int32{38, 176},
			BattleInteractions: []apitesttypes.BattleInteraction{
				{
					Target:         "single-target",
					AbilityRange:   h.GetInt32Ptr(3),
					HitAmount:      1,
					ShatterRate:    0,
					BasedOnUserAtk: false,
					Darkable:       false,
					Silenceable:    false,
					Reflectable:    false,
					Accuracy: apitesttypes.Accuracy{
						AccSource:   "rate",
						HitChance:   h.GetInt32Ptr(255),
						AccModifier: nil,
					},
					Damage: &apitesttypes.Damage{
						AttackType: database.AttackTypeAttack,
						DamageType: database.DamageTypePhysical,
						Hp: &apitesttypes.ClassDamage{
							DamageFormula:  database.DamageFormulaSpecialNoVar,
							DamageConstant: 12,
						},
						Critical:    nil,
						IsPiercing:  true,
						BreakDmgLmt: nil,
						Element:     h.GetInt32Ptr(1),
					},
					InflictedDelay:            nil,
					InflictedStatusConditions: []int32{},
					RemovedStatusConditions:   []int32{},
					CopiedStatusConditions:    []int32{},
					StatChanges:               []apitesttypes.StatChange{},
					ModifierChanges:           []apitesttypes.ModChange{},
				},
				{
					Target:         "self",
					AbilityRange:   nil,
					HitAmount:      1,
					ShatterRate:    0,
					BasedOnUserAtk: false,
					Darkable:       false,
					Silenceable:    false,
					Reflectable:    false,
					Accuracy: apitesttypes.Accuracy{
						AccSource:   "rate",
						HitChance:   h.GetInt32Ptr(255),
						AccModifier: nil,
					},
					Damage:                    nil,
					InflictedDelay:            nil,
					InflictedStatusConditions: []int32{7},
					RemovedStatusConditions:   []int32{},
					CopiedStatusConditions:    []int32{},
					StatChanges:               []apitesttypes.StatChange{},
					ModifierChanges:           []apitesttypes.ModChange{},
				},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/abilities/269",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"monsters":            0,
					"battle interactions": 1,
					"battle interactions 0 inflicted status conditions": 5,
				},
			},
			NameVer:          apitesttypes.NewExpNameVer(269, "hyper mighty g", 0),
			Rank:             h.GetInt32Ptr(6),
			AppearsInHelpBar: true,
			CanCopyCat:       false,
			AbilityType:      database.AbilityTypeOverdriveAbility,
			TypedAbility:     "/overdrive-abilities/167",
			Monsters:         []int32{},
			BattleInteractions: []apitesttypes.BattleInteraction{
				{
					Target:         "all-allies",
					AbilityRange:   nil,
					HitAmount:      1,
					ShatterRate:    0,
					BasedOnUserAtk: false,
					Darkable:       false,
					Silenceable:    false,
					Reflectable:    false,
					Accuracy: apitesttypes.Accuracy{
						AccSource:   "rate",
						HitChance:   h.GetInt32Ptr(255),
						AccModifier: nil,
					},
					Damage: &apitesttypes.Damage{
						AttackType: database.AttackTypeHeal,
						DamageType: database.DamageTypeSpecial,
						Ctb: &apitesttypes.ClassDamage{
							DamageFormula:  database.DamageFormulaCtbCurrent,
							DamageConstant: 8,
						},
						Critical:    nil,
						IsPiercing:  true,
						BreakDmgLmt: nil,
						Element:     nil,
					},
					InflictedDelay:            nil,
					InflictedStatusConditions: []int32{29, 26, 21, 28, 30},
					RemovedStatusConditions:   []int32{},
					CopiedStatusConditions:    []int32{},
					StatChanges:               []apitesttypes.StatChange{},
					ModifierChanges:           []apitesttypes.ModChange{},
				},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/abilities/72",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"monsters":            6,
					"battle interactions": 1,
				},
			},
			NameVer:          apitesttypes.NewExpNameVer(72, "blizzard", 0),
			Rank:             h.GetInt32Ptr(3),
			AppearsInHelpBar: true,
			CanCopyCat:       true,
			AbilityType:      database.AbilityTypePlayerAbility,
			TypedAbility:     "/player-abilities/72",
			Monsters:         []int32{45, 47, 65, 86, 94, 167},
			BattleInteractions: []apitesttypes.BattleInteraction{
				{
					Target:         "single-target",
					AbilityRange:   h.GetInt32Ptr(4),
					HitAmount:      1,
					ShatterRate:    10,
					BasedOnUserAtk: false,
					Darkable:       false,
					Silenceable:    true,
					Reflectable:    true,
					Accuracy: apitesttypes.Accuracy{
						AccSource:   "rate",
						HitChance:   h.GetInt32Ptr(255),
						AccModifier: nil,
					},
					Damage: &apitesttypes.Damage{
						AttackType: database.AttackTypeAttack,
						DamageType: database.DamageTypeMagical,
						Hp: &apitesttypes.ClassDamage{
							DamageFormula:  database.DamageFormulaMagVsMdf,
							DamageConstant: 12,
						},
						Critical:    nil,
						IsPiercing:  true,
						BreakDmgLmt: h.GetStrPtr("auto-ability"),
						Element:     h.GetInt32Ptr(4),
					},
					InflictedDelay:            nil,
					InflictedStatusConditions: []int32{},
					RemovedStatusConditions:   []int32{},
					CopiedStatusConditions:    []int32{},
					StatChanges:               []apitesttypes.StatChange{},
					ModifierChanges:           []apitesttypes.ModChange{},
				},
			},
		},
	}

	testSingleResources(t, tests, "GetAbility", testSrv.HandleAbilities)
}

func TestRetrieveAbilities(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.ListIDs{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/abilities?type=7",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "provided id '7' used for parameter 'type' doesn't exist. max id: 6",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/abilities?type=asd",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "invalid enum value 'asd' used for parameter 'type'. use /api/abilities/parameters to see allowed values",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/abilities?limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:    840,
			Previous: nil,
			Next:     nil,
			Results:  []int32{1, 133, 299, 333, 555, 666, 840},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/abilities?monster=86",
				ExpectedStatus: http.StatusOK,
			},
			Count:    11,
			Previous: nil,
			Next:     nil,
			Results:  []int32{69, 70, 71, 72, 415, 481, 548, 565, 705, 804, 831},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/abilities?type=2&help_bar=false&rank=5&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:    8,
			Previous: nil,
			Next:     nil,
			Results:  []int32{111, 112, 113, 114, 115, 116, 117, 118},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/abilities?type=3&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:    69,
			Previous: nil,
			Next:     nil,
			Results:  []int32{293, 295, 310, 337, 344, 352, 361},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/abilities?delay=true&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:    33,
			Previous: nil,
			Next:     nil,
			Results:  []int32{9, 10, 88, 138, 340, 546, 709, 803},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/abilities?status_inflict=4&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:    39,
			Previous: nil,
			Next:     nil,
			Results:  []int32{3, 6, 8, 134, 226, 333, 422, 708, 752},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/abilities?bdl=true&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:    281,
			Previous: nil,
			Next:     nil,
			Results:  []int32{1, 5, 65, 89, 128, 170, 205, 399, 449, 530, 623, 830},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/abilities?copycat=true&target_type=single-enemy&user_atk=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   2,
			Results: []int32{20, 21},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/abilities?darkable=true&attack_type=absorb",
				ExpectedStatus: http.StatusOK,
			},
			Count:   2,
			Results: []int32{538, 539},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/abilities?silenceable=true&reflectable=true&damage_type=magical&element=1",
				ExpectedStatus: http.StatusOK,
			},
			Count:   5,
			Results: []int32{69, 73, 77, 565, 665},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/abilities?damage_formula=special-var&can_crit=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   12,
			Results: []int32{220, 225, 229, 328, 329},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/abilities?status_remove=4",
				ExpectedStatus: http.StatusOK,
			},
			Count:   12,
			Results: []int32{53, 254, 260, 308, 672, 790},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/abilities?stat_changes=true&mod_changes=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   6,
			Results: []int32{27, 29, 264, 265, 272, 275},
		},
	}

	testIDList(t, tests, apilookups.E.Abilities.Endpoint, "RetrieveAbilities", testSrv.HandleAbilities, compareAPIResourceLists[apiresource.Named])
}

func TestSubsectionAbilities(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.ListIDs{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/character-classes/1/default-abilities",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleCharacterClasses,
			},
			Count:          5,
			ParentResource: h.GetStrPtr("/character-classes/1"),
			Results:        []int32{374, 376, 377, 378, 379},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/character-classes/1/learnable-abilities",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleCharacterClasses,
			},
			Count:          90,
			ParentResource: h.GetStrPtr("/character-classes/1"),
			Results:        []int32{1, 13, 24, 53, 66, 78, 374, 379},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/submenus/2/abilities",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleSubmenus,
			},
			Count:          22,
			ParentResource: h.GetStrPtr("/submenus/2"),
			Results:        []int32{1, 5, 8, 13, 17, 19, 22},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/topmenus/2/abilities",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleTopmenus,
			},
			Count:          6,
			ParentResource: h.GetStrPtr("/topmenus/2"),
			Results:        []int32{377, 378, 379, 380, 381, 382},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/185/abilities",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleMonsters,
			},
			Count:          6,
			ParentResource: h.GetStrPtr("/monsters/185"),
			Results:        []int32{61, 62, 460, 531, 573, 630},
		},
	}

	testIDList(t, tests, apilookups.E.Abilities.Endpoint, "SubsectionAbilities", nil, compareSimpleResourceLists[apiresource.Named, sectiontypes.AbilitySimple])
}
