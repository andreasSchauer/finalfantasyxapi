package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/tests/apitesttypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

func TestGetItemAbility(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.ItemAbility{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/item-abilities/70",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "item ability with provided id '70' doesn't exist. max id: 69",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/item-abilities/2",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"battle interactions": 1,
				},
			},
			Unique:              apitesttypes.NewExpUnique(2, "hi-potion"),
			Rank:                h.GetInt32Ptr(2),
			AppearsInHelpBar:    true,
			CanCopyCat:          true,
			UntypedAbility:      294,
			Item:                2,
			Category:            database.ItemCategoryHealing,
			CanUseOutsideBattle: true,
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
						AttackType: database.AttackTypeHeal,
						DamageType: database.DamageTypeSpecial,
						Hp: &apitesttypes.ClassDamage{
							DamageFormula:  database.DamageFormulaSpecialNoVar,
							DamageConstant: 20,
						},
						Critical:        nil,
						CriticalPlusVal: nil,
						IsPiercing:      true,
						BreakDmgLmt:     nil,
						Element:         nil,
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
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/item-abilities/48",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"battle interactions":                               1,
					"battle interactions 0 inflicted status conditions": 1,
				},
			},
			Unique:              apitesttypes.NewExpUnique(48, "gold hourglass"),
			Rank:                h.GetInt32Ptr(2),
			AppearsInHelpBar:    true,
			CanCopyCat:          true,
			UntypedAbility:      340,
			Item:                48,
			Category:            database.ItemCategoryOffensive,
			CanUseOutsideBattle: false,
			BattleInteractions: []apitesttypes.BattleInteraction{
				{
					Target:         "target-party",
					AbilityRange:   h.GetInt32Ptr(3),
					HitAmount:      1,
					ShatterRate:    70,
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
						DamageType: database.DamageTypeSpecial,
						Hp: &apitesttypes.ClassDamage{
							DamageFormula:  database.DamageFormulaSpecialVar,
							DamageConstant: 20,
						},
						Critical:        nil,
						CriticalPlusVal: nil,
						IsPiercing:      true,
						BreakDmgLmt:     nil,
						Element:         nil,
					},
					InflictedDelay: &apitesttypes.InflictedDelay{
						DelayStrength: "weak",
					},
					InflictedStatusConditions: []int32{14},
					RemovedStatusConditions:   []int32{},
					CopiedStatusConditions:    []int32{},
					StatChanges:               []apitesttypes.StatChange{},
					ModifierChanges:           []apitesttypes.ModChange{},
				},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/item-abilities/62",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"battle interactions": 1,
				},
			},
			Unique:              apitesttypes.NewExpUnique(62, "soul spring"),
			Rank:                h.GetInt32Ptr(2),
			AppearsInHelpBar:    true,
			CanCopyCat:          true,
			UntypedAbility:      354,
			Item:                62,
			Category:            database.ItemCategoryOffensive,
			CanUseOutsideBattle: false,
			BattleInteractions: []apitesttypes.BattleInteraction{
				{
					Target:         "single-target",
					AbilityRange:   h.GetInt32Ptr(3),
					HitAmount:      1,
					ShatterRate:    10,
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
						AttackType: database.AttackTypeAbsorb,
						DamageType: database.DamageTypeSpecial,
						Hp: &apitesttypes.ClassDamage{
							DamageFormula:  database.DamageFormulaSpecialVar,
							DamageConstant: 30,
						},
						Mp: &apitesttypes.ClassDamage{
							DamageFormula:  database.DamageFormulaSpecialVar,
							DamageConstant: 30,
						},
						Critical:        nil,
						CriticalPlusVal: nil,
						IsPiercing:      true,
						BreakDmgLmt:     nil,
						Element:         nil,
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
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/item-abilities/66",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"battle interactions":                1,
					"battle interactions 0 stat changes": 1,
				},
			},
			Unique:              apitesttypes.NewExpUnique(66, "stamina tonic"),
			Rank:                h.GetInt32Ptr(2),
			AppearsInHelpBar:    true,
			CanCopyCat:          true,
			UntypedAbility:      358,
			Item:                66,
			Category:            database.ItemCategorySupport,
			CanUseOutsideBattle: false,
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
					Damage:                    nil,
					InflictedDelay:            nil,
					InflictedStatusConditions: []int32{},
					RemovedStatusConditions:   []int32{},
					CopiedStatusConditions:    []int32{},
					StatChanges: []apitesttypes.StatChange{
						{
							Stat:            1,
							CalculationType: "multiply-highest",
							Value:           2,
						},
					},
					ModifierChanges: []apitesttypes.ModChange{},
				},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/item-abilities/69",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"battle interactions":               1,
					"battle interactions 0 mod changes": 1,
				},
			},
			Unique:              apitesttypes.NewExpUnique(69, "three stars"),
			Rank:                h.GetInt32Ptr(2),
			AppearsInHelpBar:    true,
			CanCopyCat:          true,
			UntypedAbility:      361,
			Item:                69,
			Category:            database.ItemCategorySupport,
			CanUseOutsideBattle: false,
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
					Damage:                    nil,
					InflictedDelay:            nil,
					InflictedStatusConditions: []int32{},
					RemovedStatusConditions:   []int32{},
					CopiedStatusConditions:    []int32{},
					StatChanges:               []apitesttypes.StatChange{},
					ModifierChanges: []apitesttypes.ModChange{
						{
							Modifier:        16,
							CalculationType: "set-value",
							Value:           0,
						},
					},
				},
			},
		},
	}

	testSingleResources(t, tests, "GetItemAbility", testSrv.HandleItemAbilities)
}

func TestRetrieveItemAbilities(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.ListIDs{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/item-abilities?limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:    69,
			Previous: nil,
			Next:     nil,
			Results:  []int32{1, 13, 27, 38, 39, 42, 55, 69},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/item-abilities?category=2&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:    34,
			Previous: nil,
			Next:     nil,
			Results:  []int32{24, 27, 35, 43, 50, 60},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/item-abilities?attack_type=heal&outside_battle=false",
				ExpectedStatus: http.StatusOK,
			},
			Count:    7,
			Previous: nil,
			Next:     nil,
			Results:  []int32{9, 10, 21, 22, 23, 54, 55},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/item-abilities?related_stat=1",
				ExpectedStatus: http.StatusOK,
			},
			Count:    2,
			Previous: nil,
			Next:     nil,
			Results:  []int32{64, 66},
		},
	}

	testIDList(t, tests, apilookups.E.ItemAbilities.Endpoint, "RetrieveItemAbilities", testSrv.HandleItemAbilities, compareAPIResourceLists[apiresource.Named])
}
