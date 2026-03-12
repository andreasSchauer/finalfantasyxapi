package api

import (
	"net/http"
	"testing"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func TestGetItemAbility(t *testing.T) {
	tests := []expItemAbility{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/item-abilities/70",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "item ability with provided id '70' doesn't exist. max id: 69.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/item-abilities/2",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"battle interactions": 1,
				},
			},
			expUnique:           newExpUnique(2, "hi-potion"),
			rank:                h.GetInt32Ptr(2),
			appearsInHelpBar:    true,
			canCopyCat:          true,
			untypedAbility:      294,
			item:                2,
			category:            1,
			CanUseOutsideBattle: true,
			battleInteractions: []expBattleInteraction{
				{
					target:         "single-target",
					abilityRange:   h.GetInt32Ptr(3),
					hitAmount:      1,
					shatterRate:    0,
					basedOnUserAtk: false,
					darkable:       false,
					silenceable:    false,
					reflectable:    false,
					accuracy: expAccuracy{
						accSource:   "rate",
						hitChance:   h.GetInt32Ptr(255),
						accModifier: nil,
					},
					damage: &expDamage{
						damageCalc: []expAbilityDamage{
							{
								attackType:     2,
								targetStat:     1,
								damageType:     3,
								damageFormula:  8,
								damageConstant: 20,
							},
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
					statChanges:               []expStatChange{},
					modifierChanges:           []expModChange{},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/item-abilities/48",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"battle interactions":                               1,
					"battle interactions 0 inflicted status conditions": 1,
				},
			},
			expUnique:           newExpUnique(48, "gold hourglass"),
			rank:                h.GetInt32Ptr(2),
			appearsInHelpBar:    true,
			canCopyCat:          true,
			untypedAbility:      340,
			item:                48,
			category:            2,
			CanUseOutsideBattle: false,
			battleInteractions: []expBattleInteraction{
				{
					target:         "target-party",
					abilityRange:   h.GetInt32Ptr(3),
					hitAmount:      1,
					shatterRate:    70,
					basedOnUserAtk: false,
					darkable:       false,
					silenceable:    false,
					reflectable:    false,
					accuracy: expAccuracy{
						accSource:   "rate",
						hitChance:   h.GetInt32Ptr(255),
						accModifier: nil,
					},
					damage: &expDamage{
						damageCalc: []expAbilityDamage{
							{
								attackType:     1,
								targetStat:     1,
								damageType:     3,
								damageFormula:  9,
								damageConstant: 20,
							},
						},
						critical:        nil,
						criticalPlusVal: nil,
						isPiercing:      true,
						breakDmgLmt:     nil,
						element:         nil,
					},
					inflictedDelay: &expInflictedDelay{
						ctbAttackType:  "attack",
						delayType:      "tick-speed-based",
						damageConstant: 24,
						delayStrength:  "weak",
					},
					inflictedStatusConditions: []int32{15},
					removedStatusConditions:   []int32{},
					copiedStatusConditions:    []int32{},
					statChanges:               []expStatChange{},
					modifierChanges:           []expModChange{},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/item-abilities/62",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"battle interactions": 1,
				},
			},
			expUnique:           newExpUnique(62, "soul spring"),
			rank:                h.GetInt32Ptr(2),
			appearsInHelpBar:    true,
			canCopyCat:          true,
			untypedAbility:      354,
			item:                62,
			category:            2,
			CanUseOutsideBattle: false,
			battleInteractions: []expBattleInteraction{
				{
					target:         "single-target",
					abilityRange:   h.GetInt32Ptr(3),
					hitAmount:      1,
					shatterRate:    10,
					basedOnUserAtk: false,
					darkable:       false,
					silenceable:    false,
					reflectable:    false,
					accuracy: expAccuracy{
						accSource:   "rate",
						hitChance:   h.GetInt32Ptr(255),
						accModifier: nil,
					},
					damage: &expDamage{
						damageCalc: []expAbilityDamage{
							{
								attackType:     3,
								targetStat:     1,
								damageType:     3,
								damageFormula:  9,
								damageConstant: 30,
							},
							{
								attackType:     3,
								targetStat:     2,
								damageType:     3,
								damageFormula:  9,
								damageConstant: 30,
							},
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
					statChanges:               []expStatChange{},
					modifierChanges:           []expModChange{},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/item-abilities/66",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"battle interactions":                1,
					"battle interactions 0 stat changes": 1,
				},
			},
			expUnique:           newExpUnique(66, "stamina tonic"),
			rank:                h.GetInt32Ptr(2),
			appearsInHelpBar:    true,
			canCopyCat:          true,
			untypedAbility:      358,
			item:                66,
			category:            3,
			CanUseOutsideBattle: false,
			battleInteractions: []expBattleInteraction{
				{
					target:         "all-allies",
					abilityRange:   nil,
					hitAmount:      1,
					shatterRate:    0,
					basedOnUserAtk: false,
					darkable:       false,
					silenceable:    false,
					reflectable:    false,
					accuracy: expAccuracy{
						accSource:   "rate",
						hitChance:   h.GetInt32Ptr(255),
						accModifier: nil,
					},
					damage:                    nil,
					inflictedDelay:            nil,
					inflictedStatusConditions: []int32{},
					removedStatusConditions:   []int32{},
					copiedStatusConditions:    []int32{},
					statChanges: []expStatChange{
						{
							stat:            1,
							calculationType: "multiply-highest",
							value:           2,
						},
					},
					modifierChanges: []expModChange{},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/item-abilities/69",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"battle interactions":               1,
					"battle interactions 0 mod changes": 1,
				},
			},
			expUnique:           newExpUnique(69, "three stars"),
			rank:                h.GetInt32Ptr(2),
			appearsInHelpBar:    true,
			canCopyCat:          true,
			untypedAbility:      361,
			item:                69,
			category:            3,
			CanUseOutsideBattle: false,
			battleInteractions: []expBattleInteraction{
				{
					target:         "all-allies",
					abilityRange:   nil,
					hitAmount:      1,
					shatterRate:    0,
					basedOnUserAtk: false,
					darkable:       false,
					silenceable:    false,
					reflectable:    false,
					accuracy: expAccuracy{
						accSource:   "rate",
						hitChance:   h.GetInt32Ptr(255),
						accModifier: nil,
					},
					damage:                    nil,
					inflictedDelay:            nil,
					inflictedStatusConditions: []int32{},
					removedStatusConditions:   []int32{},
					copiedStatusConditions:    []int32{},
					statChanges:               []expStatChange{},
					modifierChanges: []expModChange{
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

	testSingleResources(t, tests, "GetItemAbility", testCfg.HandleItemAbilities, compareItemAbilities)
}

func TestRetrieveItemAbilities(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/item-abilities?limit=max",
				expectedStatus: http.StatusOK,
			},
			count:    69,
			previous: nil,
			next:     nil,
			results:  []int32{1, 13, 27, 38, 39, 42, 55, 69},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/item-abilities?category=2&limit=max",
				expectedStatus: http.StatusOK,
			},
			count:    34,
			previous: nil,
			next:     nil,
			results:  []int32{24, 27, 35, 43, 50, 60},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/item-abilities?attack_type=heal&outside_battle=false",
				expectedStatus: http.StatusOK,
			},
			count:    5,
			previous: nil,
			next:     nil,
			results:  []int32{9, 10, 21, 22, 23},
		},
	}

	testIdList(t, tests, testCfg.e.itemAbilities.endpoint, "RetrieveItemAbilities", testCfg.HandleItemAbilities, compareAPIResourceLists[NamedApiResourceList])
}
