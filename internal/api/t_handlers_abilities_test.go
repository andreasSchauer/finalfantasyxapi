package api

import (
	"net/http"
	"testing"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func TestGetAbility(t *testing.T) {
	tests := []expAbility{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/abilities/841",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "ability with provided id '841' doesn't exist. max id: 840.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/abilities/737",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{

				},
				expLengths: map[string]int{
					"monsters": 2,
					"battle interactions": 2,
					"battle interactions 1 inflicted status conditions": 1,
				},
			},
			expNameVer: newExpNameVer(737, "self destruct", 1),
			rank: h.GetInt32Ptr(3),
			appearsInHelpBar: true,
			canCopyCat: false,
			abilityType: 6,
			typedAbility: "/enemy-abilities/340",
			monsters: []int32{38, 176},
			battleInteractions: []expBattleInteraction{
				{
					target: "single-target",
					abilityRange: h.GetInt32Ptr(3),
					hitAmount: 1,
					shatterRate: 0,
					basedOnPhysAtk: false,
					darkable: false,
					silenceable: false,
					reflectable: false,
					accuracy: expAccuracy{
						accSource: "rate",
						hitChance: h.GetInt32Ptr(255),
						accModifier: nil,
					},
					damage: &expDamage{
						damageCalc: []expAbilityDamage{
							{
								attackType: 1,
								targetStat: 1,
								damageType: 1,
								damageFormula: 8,
								damageConstant: 12,
							},
						},
						critical: nil,
						isPiercing: true,
						breakDmgLmt: nil,
						element: h.GetInt32Ptr(1),
					},
					inflictedDelay: nil,
					inflictedStatusConditions: []int32{},
					removedStatusConditions: []int32{},
					copiedStatusConditions: []int32{},
					statChanges: []expStatChange{},
					modifierChanges: []expModChange{},
				},
				{
					target: "self",
					abilityRange: nil,
					hitAmount: 1,
					shatterRate: 0,
					basedOnPhysAtk: false,
					darkable: false,
					silenceable: false,
					reflectable: false,
					accuracy: expAccuracy{
						accSource: "rate",
						hitChance: h.GetInt32Ptr(255),
						accModifier: nil,
					},
					damage: nil,
					inflictedDelay: nil,
					inflictedStatusConditions: []int32{8},
					removedStatusConditions: []int32{},
					copiedStatusConditions: []int32{},
					statChanges: []expStatChange{},
					modifierChanges: []expModChange{},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/abilities/269",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{

				},
				expLengths: map[string]int{
					"monsters": 0,
					"battle interactions": 1,
					"battle interactions 0 inflicted status conditions": 5,
				},
			},
			expNameVer: newExpNameVer(269, "hyper mighty g", 0),
			rank: h.GetInt32Ptr(6),
			appearsInHelpBar: true,
			canCopyCat: false,
			abilityType: 2,
			typedAbility: "/overdrive-abilities/167",
			monsters: []int32{},
			battleInteractions: []expBattleInteraction{
				{
					target: "all-allies",
					abilityRange: nil,
					hitAmount: 1,
					shatterRate: 0,
					basedOnPhysAtk: false,
					darkable: false,
					silenceable: false,
					reflectable: false,
					accuracy: expAccuracy{
						accSource: "rate",
						hitChance: h.GetInt32Ptr(255),
						accModifier: nil,
					},
					damage: nil,
					inflictedDelay: nil,
					inflictedStatusConditions: []int32{30, 27, 22, 29, 31},
					removedStatusConditions: []int32{},
					copiedStatusConditions: []int32{},
					statChanges: []expStatChange{},
					modifierChanges: []expModChange{},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/abilities/72",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{

				},
				expLengths: map[string]int{
					"monsters": 6,
					"battle interactions": 1,
				},
			},
			expNameVer: newExpNameVer(72, "blizzard", 0),
			rank: h.GetInt32Ptr(3),
			appearsInHelpBar: true,
			canCopyCat: true,
			abilityType: 1,
			typedAbility: "/player-abilities/72",
			monsters: []int32{45, 47, 65, 86, 94, 167},
			battleInteractions: []expBattleInteraction{
				{
					target: "single-target",
					abilityRange: h.GetInt32Ptr(4),
					hitAmount: 1,
					shatterRate: 10,
					basedOnPhysAtk: false,
					darkable: false,
					silenceable: true,
					reflectable: true,
					accuracy: expAccuracy{
						accSource: "rate",
						hitChance: h.GetInt32Ptr(255),
						accModifier: nil,
					},
					damage: &expDamage{
						damageCalc: []expAbilityDamage{
							{
								attackType: 1,
								targetStat: 1,
								damageType: 2,
								damageFormula: 3,
								damageConstant: 12,
							},
						},
						critical: nil,
						isPiercing: true,
						breakDmgLmt: h.GetStrPtr("auto-ability"),
						element: h.GetInt32Ptr(4),
					},
					inflictedDelay: nil,
					inflictedStatusConditions: []int32{},
					removedStatusConditions: []int32{},
					copiedStatusConditions: []int32{},
					statChanges: []expStatChange{},
					modifierChanges: []expModChange{},
				},
			},
		},
	}

	testSingleResources(t, tests, "GetAbility", testCfg.HandleAbilities, compareAbilities)
}

func TestGetMultipleAbilities(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/guado_guardian",
				expectedStatus: http.StatusMultipleChoices,
			},
			count:   3,
			results: []int32{94, 96, 113},
		},
	}

	testIdList(t, tests, testCfg.e.abilities.endpoint, "GetMultipleAbilities", testCfg.HandleAbilities, compareAPIResourceLists[NamedApiResourceList])
}

func TestRetrieveAbilities(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?species=wywrm",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid enum value 'wywrm' used for parameter 'species'. use /api/monsters/parameters to see allowed values.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?limit=max",
				expectedStatus: http.StatusOK,
			},
			count:    307,
			previous: nil,
			next:     nil,
			results:  []int32{1, 175, 238, 307},
		},
	}

	testIdList(t, tests, testCfg.e.abilities.endpoint, "RetrieveAbilities", testCfg.HandleAbilities, compareAPIResourceLists[NamedApiResourceList])
}