package api

import (
	"net/http"
	"testing"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func TestGetEnemyAbility(t *testing.T) {
	tests := []expEnemyAbility{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/enemy-abilities/444",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "enemy ability with provided id '444' doesn't exist. max id: 443.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/enemy-abilities/56",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{

				},
				expLengths: map[string]int{
					"monsters": 1,
					"battle interactions": 1,
					"battle interactions 0 removed status conditions": 10,
				},
			},
			expNameVer: newExpNameVer(56, "attack", 40),
			rank: h.GetInt32Ptr(1),
			appearsInHelpBar: false,
			canCopyCat: false,
			untypedAbility: 453,
			monsters: []int32{297},
			battleInteractions: []expBattleInteraction{
				{
					target: "single-target",
					abilityRange: h.GetInt32Ptr(1),
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
								damageFormula: 1,
								damageConstant: 16,
							},
						},
						critical: h.GetStrPtr("crit+ability%"),
						criticalPlusVal: h.GetInt32Ptr(25),
						isPiercing: false,
						breakDmgLmt: h.GetStrPtr("always"),
						element: nil,
					},
					inflictedDelay: nil,
					inflictedStatusConditions: []int32{},
					removedStatusConditions: []int32{22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
					copiedStatusConditions: []int32{},
					statChanges: []expStatChange{},
					modifierChanges: []expModChange{},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/enemy-abilities/228",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{

				},
				expLengths: map[string]int{
					"monsters": 1,
					"battle interactions": 1,
				},
			},
			expNameVer: newExpNameVer(228, "karma", 2),
			rank: h.GetInt32Ptr(3),
			appearsInHelpBar: true,
			canCopyCat: false,
			untypedAbility: 625,
			monsters: []int32{265},
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
								damageType: 3,
								damageFormula: 12,
								damageConstant: 100,
							},
						},
						critical: nil,
						criticalPlusVal: nil,
						isPiercing: false,
						breakDmgLmt: h.GetStrPtr("always"),
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
		{
			testGeneral: testGeneral{
				requestURL:     "/api/enemy-abilities/394",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{

				},
				expLengths: map[string]int{
					"monsters": 1,
					"battle interactions": 1,
					"battle interactions 0 removed status conditions": 9,
					"battle interactions 0 copied status conditions": 18,
				},
			},
			expNameVer: newExpNameVer(394, "swallow", 2),
			rank: h.GetInt32Ptr(3),
			appearsInHelpBar: false,
			canCopyCat: false,
			untypedAbility: 791,
			monsters: []int32{227},
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
					damage: nil,
					inflictedDelay: nil,
					inflictedStatusConditions: []int32{},
					removedStatusConditions: []int32{4, 11, 13, 15, 17, 18, 19, 20, 21},
					copiedStatusConditions: []int32{4, 13, 15, 17, 21, 22, 28, 29, 25, 26},
					statChanges: []expStatChange{},
					modifierChanges: []expModChange{},
				},
			},
		},
	}

	testSingleResources(t, tests, "GetEnemyAbility", testCfg.HandleEnemyAbilities, compareEnemyAbilities)
}

func TestGetMultipleEnemyAbilities(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/enemy-abilities/attack?limit=max",
				expectedStatus: http.StatusMultipleChoices,
			},
			count:   43,
			results: []int32{17, 22, 24, 32, 37, 39, 45, 54, 55, 59},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/enemy-abilities/breath",
				expectedStatus: http.StatusMultipleChoices,
			},
			count:   6,
			results: []int32{88, 89, 90, 91, 92, 93},
		},
	}

	testIdList(t, tests, testCfg.e.enemyAbilities.endpoint, "GetMultipleEnemyAbilities", testCfg.HandleEnemyAbilities, compareAPIResourceLists[NamedApiResourceList])
}

func TestRetrieveEnemyAbilities(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/enemy-abilities?limit=max",
				expectedStatus: http.StatusOK,
			},
			count:    443,
			previous: nil,
			next:     nil,
			results:  []int32{1, 223, 18, 45, 177, 386, 443},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/enemy-abilities?target_type=single-target&rank=3&damage_type=magical",
				expectedStatus: http.StatusOK,
			},
			count:    20,
			previous: nil,
			next:     nil,
			results:  []int32{12, 111, 150, 180, 258, 380, 435},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/enemy-abilities?damage_formula=10",
				expectedStatus: http.StatusOK,
			},
			count:    16,
			previous: nil,
			next:     nil,
			results:  []int32{126, 133, 155, 247, 402},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/enemy-abilities?monster=244",
				expectedStatus: http.StatusOK,
			},
			count:    4,
			previous: nil,
			next:     nil,
			results:  []int32{46, 227, 427, 429},
		},
	}

	testIdList(t, tests, testCfg.e.enemyAbilities.endpoint, "RetrieveEnemyAbilities", testCfg.HandleEnemyAbilities, compareAPIResourceLists[NamedApiResourceList])
}