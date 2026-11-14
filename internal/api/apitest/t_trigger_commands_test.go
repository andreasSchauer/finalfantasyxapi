package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

func TestGetTriggerCommand(t *testing.T) {
	t.Parallel()
	tests := []TriggerCommand{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/trigger-commands/13",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "trigger command with provided id '13' doesn't exist. max id: 12",
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
					"monster-formations":                 3,
				},
			},
			NameVer:           newExpNameVer(4, "talk", 1),
			rank:              h.GetInt32Ptr(3),
			appearsInHelpBar:  true,
			canCopyCat:        false,
			untypedAbility:    365,
			topmenu:           h.GetInt32Ptr(3),
			usedBy:            []int32{5, 9, 10},
			monsterFormations: []int32{137, 192, 235},
			battleInteractions: []BattleInteraction{
				{
					target:         "self",
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
							stat:            3,
							calculationType: "added-value",
							value:           10,
						},
					},
					modifierChanges: []ModChange{},
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
					"monster-formations":                1,
				},
			},
			NameVer:           newExpNameVer(6, "talk", 3),
			rank:              h.GetInt32Ptr(3),
			appearsInHelpBar:  true,
			canCopyCat:        false,
			untypedAbility:    367,
			topmenu:           h.GetInt32Ptr(3),
			usedBy:            []int32{5},
			monsterFormations: []int32{270},
			battleInteractions: []BattleInteraction{
				{
					target:         "single-enemy",
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
				requestURL:     "/api/trigger-commands/12?ability_user=wakka&celestial_weapon=false",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"battle interactions": 1,
					"used by":             3,
					"monster-formations":  1,
				},
			},
			NameVer:           newExpNameVer(12, "struggle", 0),
			rank:              h.GetInt32Ptr(3),
			appearsInHelpBar:  true,
			canCopyCat:        false,
			untypedAbility:    373,
			topmenu:           h.GetInt32Ptr(3),
			usedBy:            []int32{5, 7, 11},
			monsterFormations: []int32{7},
			battleInteractions: []BattleInteraction{
				{
					target:         "single-enemy",
					abilityRange:   h.GetInt32Ptr(4),
					hitAmount:      1,
					shatterRate:    0,
					basedOnUserAtk: true,
					darkable:       true,
					silenceable:    false,
					reflectable:    false,
					accuracy: Accuracy{
						accSource:   "rate",
						hitChance:   h.GetInt32Ptr(255),
						accModifier: nil,
					},
					damage: &Damage{
						attackType: database.AttackTypeAttack,
						damageType: database.DamageTypePhysical,
						hp: &ClassDamage{
							damageFormula:  database.DamageFormulaStrVsDef,
							damageConstant: 16,
						},
						critical:        h.GetStrPtr("crit-equip"),
						criticalPlusVal: nil,
						isPiercing:      true,
						breakDmgLmt:     h.GetStrPtr("auto-ability"),
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
	}

	testSingleResources(t, tests, "GetTriggerCommand", testSrv.HandleTriggerCommands)
}

func TestGetMultipleTriggerCommands(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/trigger-commands/talk",
				expectedStatus: http.StatusMultipleChoices,
			},
			count:   3,
			results: []int32{4, 5, 6},
		},
	}

	testIDList(t, tests, apilookups.E.TriggerCommands.Endpoint, "GetMultipleTriggerCommands", testSrv.HandleTriggerCommands, compareAPIResourceLists[apiresource.Named])
}

func TestRetrieveTriggerCommands(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
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

	testIDList(t, tests, apilookups.E.TriggerCommands.Endpoint, "RetrieveTriggerCommands", testSrv.HandleTriggerCommands, compareAPIResourceLists[apiresource.Named])
}
