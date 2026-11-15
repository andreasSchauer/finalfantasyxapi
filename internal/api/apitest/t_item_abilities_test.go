package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

func TestGetItemAbility(t *testing.T) {
	t.Parallel()
	tests := []ItemAbility{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/item-abilities/70",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "item ability with provided id '70' doesn't exist. max id: 69",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/item-abilities/2",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"battle interactions": 1,
				},
			},
			Unique:              newExpUnique(2, "hi-potion"),
			rank:                h.GetInt32Ptr(2),
			appearsInHelpBar:    true,
			canCopyCat:          true,
			untypedAbility:      294,
			item:                2,
			category:            database.ItemCategoryHealing,
			CanUseOutsideBattle: true,
			battleInteractions: []BattleInteraction{
				{
					target:         "single-target",
					abilityRange:   h.GetInt32Ptr(3),
					hitAmount:      1,
					shatterRate:    0,
					basedOnUserAtk: false,
					darkable:       false,
					silenceable:    false,
					reflectable:    false,
					accuracy: Accuracy{
						accSource:   "rate",
						hitChance:   h.GetInt32Ptr(255),
						accModifier: nil,
					},
					damage: &Damage{
						attackType: database.AttackTypeHeal,
						damageType: database.DamageTypeSpecial,
						hp: &ClassDamage{
							damageFormula:  database.DamageFormulaSpecialNoVar,
							damageConstant: 20,
						},
						critical:        nil,
						criticalPlusVal: nil,
						isPiercing:      true,
						breakDmgLmt:     nil,
						element:         nil,
					},
					inflictedDelay:            nil,
					inflictedStatusConditions: []int32{},
					removedStatusConditions:   []int32{},
					copiedStatusConditions:    []int32{},
					statChanges:               []StatChange{},
					modifierChanges:           []ModChange{},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/item-abilities/48",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"battle interactions":                               1,
					"battle interactions 0 inflicted status conditions": 1,
				},
			},
			Unique:              newExpUnique(48, "gold hourglass"),
			rank:                h.GetInt32Ptr(2),
			appearsInHelpBar:    true,
			canCopyCat:          true,
			untypedAbility:      340,
			item:                48,
			category:            database.ItemCategoryOffensive,
			CanUseOutsideBattle: false,
			battleInteractions: []BattleInteraction{
				{
					target:         "target-party",
					abilityRange:   h.GetInt32Ptr(3),
					hitAmount:      1,
					shatterRate:    70,
					basedOnUserAtk: false,
					darkable:       false,
					silenceable:    false,
					reflectable:    false,
					accuracy: Accuracy{
						accSource:   "rate",
						hitChance:   h.GetInt32Ptr(255),
						accModifier: nil,
					},
					damage: &Damage{
						attackType: database.AttackTypeAttack,
						damageType: database.DamageTypeSpecial,
						hp: &ClassDamage{
							damageFormula:  database.DamageFormulaSpecialVar,
							damageConstant: 20,
						},
						critical:        nil,
						criticalPlusVal: nil,
						isPiercing:      true,
						breakDmgLmt:     nil,
						element:         nil,
					},
					inflictedDelay: &InflictedDelay{
						delayStrength: "weak",
					},
					inflictedStatusConditions: []int32{14},
					removedStatusConditions:   []int32{},
					copiedStatusConditions:    []int32{},
					statChanges:               []StatChange{},
					modifierChanges:           []ModChange{},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/item-abilities/62",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"battle interactions": 1,
				},
			},
			Unique:              newExpUnique(62, "soul spring"),
			rank:                h.GetInt32Ptr(2),
			appearsInHelpBar:    true,
			canCopyCat:          true,
			untypedAbility:      354,
			item:                62,
			category:            database.ItemCategoryOffensive,
			CanUseOutsideBattle: false,
			battleInteractions: []BattleInteraction{
				{
					target:         "single-target",
					abilityRange:   h.GetInt32Ptr(3),
					hitAmount:      1,
					shatterRate:    10,
					basedOnUserAtk: false,
					darkable:       false,
					silenceable:    false,
					reflectable:    false,
					accuracy: Accuracy{
						accSource:   "rate",
						hitChance:   h.GetInt32Ptr(255),
						accModifier: nil,
					},
					damage: &Damage{
						attackType: database.AttackTypeAbsorb,
						damageType: database.DamageTypeSpecial,
						hp: &ClassDamage{
							damageFormula:  database.DamageFormulaSpecialVar,
							damageConstant: 30,
						},
						mp: &ClassDamage{
							damageFormula:  database.DamageFormulaSpecialVar,
							damageConstant: 30,
						},
						critical:        nil,
						criticalPlusVal: nil,
						isPiercing:      true,
						breakDmgLmt:     nil,
						element:         nil,
					},
					inflictedDelay:            nil,
					inflictedStatusConditions: []int32{},
					removedStatusConditions:   []int32{},
					copiedStatusConditions:    []int32{},
					statChanges:               []StatChange{},
					modifierChanges:           []ModChange{},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/item-abilities/66",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"battle interactions":                1,
					"battle interactions 0 stat changes": 1,
				},
			},
			Unique:              newExpUnique(66, "stamina tonic"),
			rank:                h.GetInt32Ptr(2),
			appearsInHelpBar:    true,
			canCopyCat:          true,
			untypedAbility:      358,
			item:                66,
			category:            database.ItemCategorySupport,
			CanUseOutsideBattle: false,
			battleInteractions: []BattleInteraction{
				{
					target:         "all-allies",
					abilityRange:   nil,
					hitAmount:      1,
					shatterRate:    0,
					basedOnUserAtk: false,
					darkable:       false,
					silenceable:    false,
					reflectable:    false,
					accuracy: Accuracy{
						accSource:   "rate",
						hitChance:   h.GetInt32Ptr(255),
						accModifier: nil,
					},
					damage:                    nil,
					inflictedDelay:            nil,
					inflictedStatusConditions: []int32{},
					removedStatusConditions:   []int32{},
					copiedStatusConditions:    []int32{},
					statChanges: []StatChange{
						{
							stat:            1,
							calculationType: "multiply-highest",
							value:           2,
						},
					},
					modifierChanges: []ModChange{},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/item-abilities/69",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"battle interactions":               1,
					"battle interactions 0 mod changes": 1,
				},
			},
			Unique:              newExpUnique(69, "three stars"),
			rank:                h.GetInt32Ptr(2),
			appearsInHelpBar:    true,
			canCopyCat:          true,
			untypedAbility:      361,
			item:                69,
			category:            database.ItemCategorySupport,
			CanUseOutsideBattle: false,
			battleInteractions: []BattleInteraction{
				{
					target:         "all-allies",
					abilityRange:   nil,
					hitAmount:      1,
					shatterRate:    0,
					basedOnUserAtk: false,
					darkable:       false,
					silenceable:    false,
					reflectable:    false,
					accuracy: Accuracy{
						accSource:   "rate",
						hitChance:   h.GetInt32Ptr(255),
						accModifier: nil,
					},
					damage:                    nil,
					inflictedDelay:            nil,
					inflictedStatusConditions: []int32{},
					removedStatusConditions:   []int32{},
					copiedStatusConditions:    []int32{},
					statChanges:               []StatChange{},
					modifierChanges: []ModChange{
						{
							modifier:        16,
							calculationType: "set-value",
							value:           0,
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
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/item-abilities?limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:    69,
			Previous: nil,
			Next:     nil,
			Results:  []int32{1, 13, 27, 38, 39, 42, 55, 69},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/item-abilities?category=2&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:    34,
			Previous: nil,
			Next:     nil,
			Results:  []int32{24, 27, 35, 43, 50, 60},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/item-abilities?attack_type=heal&outside_battle=false",
				ExpectedStatus: http.StatusOK,
			},
			Count:    7,
			Previous: nil,
			Next:     nil,
			Results:  []int32{9, 10, 21, 22, 23, 54, 55},
		},
		{
			TestGeneral: TestGeneral{
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
