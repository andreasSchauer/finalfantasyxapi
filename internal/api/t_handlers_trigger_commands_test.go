package api

import (
	"net/http"
	"testing"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func TestGetTriggerCommand(t *testing.T) {
	tests := []expTriggerCommand{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/trigger-commands/13",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "trigger command with provided id '13' doesn't exist. max id: 12.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/trigger-commands/4",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"battle interactions":                1,
					"battle interactions 0 stat changes": 1,
					"used by":                            3,
					"monster formations":                 3,
				},
			},
			expNameVer:        newExpNameVer(4, "talk", 1),
			rank:              h.GetInt32Ptr(3),
			appearsInHelpBar:  true,
			canCopyCat:        false,
			untypedAbility:    365,
			topmenu:           h.GetInt32Ptr(3),
			usedBy:            []int32{5, 9, 10},
			monsterFormations: []int32{137, 192, 235},
			battleInteractions: []expBattleInteraction{
				{
					target:         "self",
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
							stat:            3,
							calculationType: "added-value",
							value:           10,
						},
					},
					modifierChanges: []expModChange{},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/trigger-commands/6",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"battle interactions":               1,
					"battle interactions 0 mod changes": 1,
					"used by":                           1,
					"monster formations":                1,
				},
			},
			expNameVer:        newExpNameVer(6, "talk", 3),
			rank:              h.GetInt32Ptr(3),
			appearsInHelpBar:  true,
			canCopyCat:        false,
			untypedAbility:    367,
			topmenu:           h.GetInt32Ptr(3),
			usedBy:            []int32{5},
			monsterFormations: []int32{270},
			battleInteractions: []expBattleInteraction{
				{
					target:         "single-enemy",
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
							modifier:        18,
							calculationType: "set-value",
							value:           1,
						},
					},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/trigger-commands/12?ability_user=wakka",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"battle interactions": 1,
					"used by":             3,
					"monster formations":  1,
				},
			},
			expNameVer:        newExpNameVer(12, "struggle", 0),
			rank:              h.GetInt32Ptr(3),
			appearsInHelpBar:  true,
			canCopyCat:        false,
			untypedAbility:    373,
			topmenu:           h.GetInt32Ptr(3),
			usedBy:            []int32{5, 7, 11},
			monsterFormations: []int32{7},
			battleInteractions: []expBattleInteraction{
				{
					target:         "single-enemy",
					abilityRange:   h.GetInt32Ptr(4),
					hitAmount:      1,
					shatterRate:    0,
					basedOnUserAtk: true,
					darkable:       true,
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
								damageType:     1,
								damageFormula:  1,
								damageConstant: 16,
							},
						},
						critical:        h.GetStrPtr("crit+weapon%"),
						criticalPlusVal: nil,
						isPiercing:      true,
						breakDmgLmt:     h.GetStrPtr("auto-ability"),
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
	}

	testSingleResources(t, tests, "GetTriggerCommand", testCfg.HandleTriggerCommands, compareTriggerCommands)
}

func TestGetMultipleTriggerCommands(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/trigger-commands/talk",
				expectedStatus: http.StatusMultipleChoices,
			},
			count:   3,
			results: []int32{4, 5, 6},
		},
	}

	testIdList(t, tests, testCfg.e.triggerCommands.endpoint, "GetMultipleTriggerCommands", testCfg.HandleTriggerCommands, compareAPIResourceLists[NamedApiResourceList])
}

func TestRetrieveTriggerCommands(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/trigger-commands?limit=max",
				expectedStatus: http.StatusOK,
			},
			count:    12,
			previous: nil,
			next:     nil,
			results:  []int32{1, 4, 7, 9, 11, 12},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/trigger-commands?related_stat=6",
				expectedStatus: http.StatusOK,
			},
			count:    1,
			previous: nil,
			next:     nil,
			results:  []int32{5},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/trigger-commands?user=7",
				expectedStatus: http.StatusOK,
			},
			count:    3,
			previous: nil,
			next:     nil,
			results:  []int32{5, 10, 12},
		},
	}

	testIdList(t, tests, testCfg.e.triggerCommands.endpoint, "RetrieveTriggerCommands", testCfg.HandleTriggerCommands, compareAPIResourceLists[NamedApiResourceList])
}
