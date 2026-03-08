package api

import (
	"net/http"
	"testing"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func TestGetOverdriveAbility(t *testing.T) {
	tests := []expOverdriveAbility{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-abilities/191",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "overdrive ability with provided id '191' doesn't exist. max id: 190.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-abilities/20",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{

				},
				expLengths: map[string]int{
					"overdrives": 2,
					"battle interactions": 1,
				},
			},
			expNameVer: newExpNameVer(20, "thunder shot", 1),
			rank: h.GetInt32Ptr(4),
			untypedAbility: 122,
			overdriveCommand: h.GetInt32Ptr(3),
			overdrives: []int32{13, 16},
			battleInteractions: []expBattleInteraction{
				{
					target: "random-enemy",
					abilityRange: h.GetInt32Ptr(3),
					hitAmount: 1,
					shatterRate: 100,
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
								damageType: 3,
								damageFormula: 1,
								damageConstant: 34,
							},
						},
						critical: nil,
						criticalPlusVal: nil,
						isPiercing: true,
						breakDmgLmt: h.GetStrPtr("auto-ability"),
						element: h.GetInt32Ptr(2),
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
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-abilities/159",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{

				},
				expLengths: map[string]int{
					"overdrives": 1,
					"battle interactions": 1,
					"battle interactions 0 inflicted status conditions": 1,
					"battle interactions 0 removed status conditions": 10,
				},
			},
			expNameVer: newExpNameVer(159, "final elixir", 0),
			rank: h.GetInt32Ptr(6),
			untypedAbility: 261,
			overdriveCommand: h.GetInt32Ptr(7),
			overdrives: []int32{93},
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
					damage: &expDamage{
						damageCalc: []expAbilityDamage{
							{
								attackType: 2,
								targetStat: 1,
								damageType: 3,
								damageFormula: 6,
								damageConstant: 16,
							},
							{
								attackType: 2,
								targetStat: 2,
								damageType: 3,
								damageFormula: 6,
								damageConstant: 16,
							},
						},
						critical: nil,
						criticalPlusVal: nil,
						isPiercing: true,
						breakDmgLmt: h.GetStrPtr("always"),
						element: nil,
					},
					inflictedDelay: nil,
					inflictedStatusConditions: []int32{41},
					removedStatusConditions: []int32{1, 2, 3, 4, 10, 11, 13, 14, 15, 17},
					copiedStatusConditions: []int32{},
					statChanges: []expStatChange{},
					modifierChanges: []expModChange{},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-abilities/163",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{

				},
				expLengths: map[string]int{
					"battle interactions": 1,
					"battle interactions 0 inflicted status conditions": 4,
					"battle interactions 0 stat changes": 2,
					"battle interactions 0 mod changes": 4,
				},
			},
			expNameVer: newExpNameVer(163, "ultra nulall", 0),
			rank: h.GetInt32Ptr(6),
			untypedAbility: 265,
			overdriveCommand: h.GetInt32Ptr(7),
			overdrives: []int32{97},
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
					inflictedStatusConditions: []int32{23, 24, 25, 26},
					removedStatusConditions: []int32{},
					copiedStatusConditions: []int32{},
					statChanges: []expStatChange{
						{
							stat: 3,
							calculationType: "added-value",
							value: 5,
						},
						{
							stat: 5,
							calculationType: "added-value",
							value: 5,
						},
					},
					modifierChanges: []expModChange{
						{
							modifier: 3,
							calculationType: "added-value",
							value: -0.33,
						},
						{
							modifier: 4,
							calculationType: "added-value",
							value: -0.33,
						},
						{
							modifier: 6,
							calculationType: "added-percentage",
							value: 50,
						},
						{
							modifier: 7,
							calculationType: "added-percentage",
							value: 50,
						},
					},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-abilities/7",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{

				},
				expLengths: map[string]int{
					"overdrives": 1,
					"battle interactions": 2,
				},
			},
			expNameVer: newExpNameVer(7, "blitz ace", 1),
			rank: h.GetInt32Ptr(7),
			untypedAbility: 109,
			overdriveCommand: h.GetInt32Ptr(1),
			overdrives: []int32{4},
			battleInteractions: []expBattleInteraction{
				{
					target: "single-enemy",
					abilityRange: h.GetInt32Ptr(1),
					hitAmount: 8,
					shatterRate: 100,
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
								damageType: 3,
								damageFormula: 1,
								damageConstant: 4,
							},
						},
						critical: h.GetStrPtr("crit"),
						criticalPlusVal: nil,
						isPiercing: true,
						breakDmgLmt: h.GetStrPtr("auto-ability"),
						element: nil,
					},
					inflictedDelay: nil,
					inflictedStatusConditions: []int32{},
					removedStatusConditions: []int32{},
					copiedStatusConditions: []int32{},
					statChanges: []expStatChange{},
					modifierChanges: []expModChange{},
				},
				{
					target: "single-enemy",
					abilityRange: h.GetInt32Ptr(1),
					hitAmount: 1,
					shatterRate: 100,
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
								damageType: 3,
								damageFormula: 1,
								damageConstant: 24,
							},
						},
						critical: h.GetStrPtr("crit"),
						criticalPlusVal: nil,
						isPiercing: true,
						breakDmgLmt: h.GetStrPtr("auto-ability"),
						element: nil,
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

	testSingleResources(t, tests, "GetOverdriveAbility", testCfg.HandleOverdriveAbilities, compareOverdriveAbilities)
}

func TestGetMultipleOverdriveAbilities(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-abilities/grand_summon",
				expectedStatus: http.StatusMultipleChoices,
			},
			count:   8,
			results: []int32{9, 10, 11, 12, 13, 14, 15, 16},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-abilities/dragon_fang",
				expectedStatus: http.StatusMultipleChoices,
			},
			count:   3,
			results: []int32{107, 108, 109},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-abilities/fire_fury",
				expectedStatus: http.StatusMultipleChoices,
			},
			count:   3,
			results: []int32{40, 41, 42},
		},
	}

	testIdList(t, tests, testCfg.e.overdriveAbilities.endpoint, "GetMultipleOverdriveAbilities", testCfg.HandleOverdriveAbilities, compareAPIResourceLists[NamedApiResourceList])
}

func TestRetrieveOverdriveAbilities(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-abilities?limit=max",
				expectedStatus: http.StatusOK,
			},
			count:    190,
			previous: nil,
			next:     nil,
			results:  []int32{1, 26, 83, 99, 176, 190},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-abilities?rank=6&can_crit=true",
				expectedStatus: http.StatusOK,
			},
			count:    14,
			previous: nil,
			next:     nil,
			results:  []int32{113, 115, 120, 123, 126, 127, 190},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-abilities?status_remove=10",
				expectedStatus: http.StatusOK,
			},
			count:    4,
			previous: nil,
			next:     nil,
			results:  []int32{152, 153, 158, 159},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-abilities?stat_changes=true&mod_changes=true",
				expectedStatus: http.StatusOK,
			},
			count:    4,
			previous: nil,
			next:     nil,
			results:  []int32{162, 163, 170, 173},
		},
	}

	testIdList(t, tests, testCfg.e.overdriveAbilities.endpoint, "RetrieveOverdriveAbilities", testCfg.HandleOverdriveAbilities, compareAPIResourceLists[NamedApiResourceList])
}