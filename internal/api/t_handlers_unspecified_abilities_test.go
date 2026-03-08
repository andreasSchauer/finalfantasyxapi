package api

import (
	"net/http"
	"testing"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func TestGetUnspecifiedAbility(t *testing.T) {
	tests := []expUnspecifiedAbility{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/unspecified-abilities/25",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "unspecified ability with provided id '25' doesn't exist. max id: 24.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/unspecified-abilities/2?user=character:cindy",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "unknown character 'cindy' used for parameter 'user'.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/unspecified-abilities/2?user=character:wakka",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid input for parameter 'user': character 'wakka' can't learn unspecified ability 'attack - 2'",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/unspecified-abilities/2?user=aeon:we",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "unknown aeon 'we' used for parameter 'user'.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/unspecified-abilities/1?user=character:wakka",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{

				},
				expLengths: map[string]int{
					"battle interactions": 1,
				},
			},
			expNameVer: newExpNameVer(1, "attack", 1),
			rank: h.GetInt32Ptr(3),
			appearsInHelpBar: false,
			canCopyCat: true,
			untypedAbility: 374,
			topmenu: h.GetInt32Ptr(1),
			submenu: nil,
			openSubmenu: nil,
			battleInteractions: []expBattleInteraction{
				{
					target: "single-target",
					abilityRange: h.GetInt32Ptr(4),
					hitAmount: 1,
					shatterRate: 30,
					basedOnPhysAtk: true,
					darkable: true,
					silenceable: false,
					reflectable: false,
					accuracy: expAccuracy{
						accSource: "accuracy",
						hitChance: nil,
						accModifier: h.GetFloat32Ptr(1),
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
						critical: h.GetStrPtr("crit+weapon%"),
						criticalPlusVal: nil,
						isPiercing: false,
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
		{
			testGeneral: testGeneral{
				requestURL:     "/api/unspecified-abilities/2?user=aeon:cindy",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{

				},
				expLengths: map[string]int{
					"battle interactions": 1,
				},
			},
			expNameVer: newExpNameVer(2, "attack", 2),
			rank: h.GetInt32Ptr(3),
			appearsInHelpBar: false,
			canCopyCat: false,
			untypedAbility: 375,
			topmenu: h.GetInt32Ptr(1),
			submenu: nil,
			openSubmenu: nil,
			battleInteractions: []expBattleInteraction{
				{
					target: "single-target",
					abilityRange: h.GetInt32Ptr(1),
					hitAmount: 1,
					shatterRate: 20,
					basedOnPhysAtk: true,
					darkable: true,
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
								damageConstant: 14,
							},
						},
						critical: h.GetStrPtr("crit+weapon%"),
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
		{
			testGeneral: testGeneral{
				requestURL:     "/api/unspecified-abilities/10",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{

				},
				expLengths: map[string]int{
					"battle interactions": 0,
				},
			},
			expNameVer: newExpNameVer(10, "switch", 1),
			rank: h.GetInt32Ptr(0),
			appearsInHelpBar: false,
			canCopyCat: false,
			untypedAbility: 383,
			topmenu: nil,
			submenu: h.GetInt32Ptr(13),
			openSubmenu: nil,
			battleInteractions: []expBattleInteraction{},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/unspecified-abilities/4",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{

				},
				expLengths: map[string]int{
					"battle interactions": 0,
				},
			},
			expNameVer: newExpNameVer(4, "weapon", 0),
			rank: h.GetInt32Ptr(1),
			appearsInHelpBar: false,
			canCopyCat: false,
			untypedAbility: 377,
			topmenu: h.GetInt32Ptr(2),
			submenu: nil,
			openSubmenu: h.GetInt32Ptr(7),
			battleInteractions: []expBattleInteraction{},
		},
		
	}

	testSingleResources(t, tests, "GetUnspecifiedAbility", testCfg.HandleUnspecifiedAbilities, compareUnspecifiedAbilities)
}

func TestGetMultipleUnspecifiedAbilities(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/unspecified-abilities/attack",
				expectedStatus: http.StatusMultipleChoices,
			},
			count:   2,
			results: []int32{1, 2},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/unspecified-abilities/switch",
				expectedStatus: http.StatusMultipleChoices,
			},
			count:   7,
			results: []int32{10, 11, 12, 13, 14, 15, 16},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/unspecified-abilities/summon",
				expectedStatus: http.StatusMultipleChoices,
			},
			count:   8,
			results: []int32{17, 18, 19, 20, 21, 22, 23, 24},
		},
	}

	testIdList(t, tests, testCfg.e.unspecifiedAbilities.endpoint, "GetMultipleUnspecifiedAbilities", testCfg.HandleUnspecifiedAbilities, compareAPIResourceLists[NamedApiResourceList])
}

func TestRetrieveUnspecifiedAbilities(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/unspecified-abilities?limit=max",
				expectedStatus: http.StatusOK,
			},
			count:    24,
			previous: nil,
			next:     nil,
			results:  []int32{1, 8, 13, 17, 18, 24},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/unspecified-abilities?rank=1",
				expectedStatus: http.StatusOK,
			},
			count:    3,
			previous: nil,
			next:     nil,
			results:  []int32{4, 5, 6},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/unspecified-abilities?help_bar=true",
				expectedStatus: http.StatusOK,
			},
			count:    3,
			previous: nil,
			next:     nil,
			results:  []int32{7, 8, 9},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/unspecified-abilities?char_class=characters&copycat=true",
				expectedStatus: http.StatusOK,
			},
			count:    3,
			previous: nil,
			next:     nil,
			results:  []int32{1, 3, 6},
		},
	}

	testIdList(t, tests, testCfg.e.unspecifiedAbilities.endpoint, "RetrieveUnspecifiedAbilities", testCfg.HandleUnspecifiedAbilities, compareAPIResourceLists[NamedApiResourceList])
}