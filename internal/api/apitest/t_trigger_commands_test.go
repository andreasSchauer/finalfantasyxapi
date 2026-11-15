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
			TestGeneral: TestGeneral{
				RequestURL:     "/api/trigger-commands/13",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "trigger command with provided id '13' doesn't exist. max id: 12",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/trigger-commands/4",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"battle interactions":                1,
					"battle interactions 0 stat changes": 1,
					"used by":                            3,
					"monster-formations":                 3,
				},
			},
			NameVer:           NewExpNameVer(4, "talk", 1),
			Rank:              h.GetInt32Ptr(3),
			AppearsInHelpBar:  true,
			CanCopyCat:        false,
			UntypedAbility:    365,
			Topmenu:           h.GetInt32Ptr(3),
			UsedBy:            []int32{5, 9, 10},
			MonsterFormations: []int32{137, 192, 235},
			BattleInteractions: []BattleInteraction{
				{
					Target:         "self",
					AbilityRange:   nil,
					HitAmount:      1,
					ShatterRate:    0,
					BasedOnUserAtk: false,
					Darkable:       false,
					Silenceable:    false,
					Reflectable:    false,
					Accuracy: Accuracy{
						AccSource:   "rate",
						HitChance:   h.GetInt32Ptr(255),
						AccModifier: nil,
					},
					Damage:                    nil,
					InflictedDelay:            nil,
					InflictedStatusConditions: []int32{},
					RemovedStatusConditions:   []int32{},
					CopiedStatusConditions:    []int32{},
					StatChanges: []StatChange{
						{
							Stat:            3,
							CalculationType: "added-value",
							Value:           10,
						},
					},
					ModifierChanges: []ModChange{},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/trigger-commands/6",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"battle interactions":               1,
					"battle interactions 0 mod changes": 1,
					"used by":                           1,
					"monster-formations":                1,
				},
			},
			NameVer:           NewExpNameVer(6, "talk", 3),
			Rank:              h.GetInt32Ptr(3),
			AppearsInHelpBar:  true,
			CanCopyCat:        false,
			UntypedAbility:    367,
			Topmenu:           h.GetInt32Ptr(3),
			UsedBy:            []int32{5},
			MonsterFormations: []int32{270},
			BattleInteractions: []BattleInteraction{
				{
					Target:         "single-enemy",
					AbilityRange:   nil,
					HitAmount:      1,
					ShatterRate:    0,
					BasedOnUserAtk: false,
					Darkable:       false,
					Silenceable:    false,
					Reflectable:    false,
					Accuracy: Accuracy{
						AccSource:   "rate",
						HitChance:   h.GetInt32Ptr(255),
						AccModifier: nil,
					},
					Damage:                    nil,
					InflictedDelay:            nil,
					InflictedStatusConditions: []int32{},
					RemovedStatusConditions:   []int32{},
					CopiedStatusConditions:    []int32{},
					StatChanges:               []StatChange{},
					ModifierChanges: []ModChange{
						{
							Modifier:        18,
							CalculationType: "set-value",
							Value:           1,
						},
					},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/trigger-commands/12?ability_user=wakka&celestial_weapon=false",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"battle interactions": 1,
					"used by":             3,
					"monster-formations":  1,
				},
			},
			NameVer:           NewExpNameVer(12, "struggle", 0),
			Rank:              h.GetInt32Ptr(3),
			AppearsInHelpBar:  true,
			CanCopyCat:        false,
			UntypedAbility:    373,
			Topmenu:           h.GetInt32Ptr(3),
			UsedBy:            []int32{5, 7, 11},
			MonsterFormations: []int32{7},
			BattleInteractions: []BattleInteraction{
				{
					Target:         "single-enemy",
					AbilityRange:   h.GetInt32Ptr(4),
					HitAmount:      1,
					ShatterRate:    0,
					BasedOnUserAtk: true,
					Darkable:       true,
					Silenceable:    false,
					Reflectable:    false,
					Accuracy: Accuracy{
						AccSource:   "rate",
						HitChance:   h.GetInt32Ptr(255),
						AccModifier: nil,
					},
					Damage: &Damage{
						AttackType: database.AttackTypeAttack,
						DamageType: database.DamageTypePhysical,
						Hp: &ClassDamage{
							DamageFormula:  database.DamageFormulaStrVsDef,
							DamageConstant: 16,
						},
						Critical:        h.GetStrPtr("crit-equip"),
						CriticalPlusVal: nil,
						IsPiercing:      true,
						BreakDmgLmt:     h.GetStrPtr("auto-ability"),
						Element:         nil,
					},
					InflictedDelay:            nil,
					InflictedStatusConditions: []int32{},
					RemovedStatusConditions:   []int32{},
					CopiedStatusConditions:    []int32{},
					StatChanges:               []StatChange{},
					ModifierChanges:           []ModChange{},
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
			TestGeneral: TestGeneral{
				RequestURL:     "/api/trigger-commands/talk",
				ExpectedStatus: http.StatusMultipleChoices,
			},
			Count:   3,
			Results: []int32{4, 5, 6},
		},
	}

	testIDList(t, tests, apilookups.E.TriggerCommands.Endpoint, "GetMultipleTriggerCommands", testSrv.HandleTriggerCommands, compareAPIResourceLists[apiresource.Named])
}

func TestRetrieveTriggerCommands(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/trigger-commands?limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:    12,
			Previous: nil,
			Next:     nil,
			Results:  []int32{1, 4, 7, 9, 11, 12},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/trigger-commands?related_stat=6",
				ExpectedStatus: http.StatusOK,
			},
			Count:    1,
			Previous: nil,
			Next:     nil,
			Results:  []int32{5},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/trigger-commands?user=7",
				ExpectedStatus: http.StatusOK,
			},
			Count:    3,
			Previous: nil,
			Next:     nil,
			Results:  []int32{5, 10, 12},
		},
	}

	testIDList(t, tests, apilookups.E.TriggerCommands.Endpoint, "RetrieveTriggerCommands", testSrv.HandleTriggerCommands, compareAPIResourceLists[apiresource.Named])
}
