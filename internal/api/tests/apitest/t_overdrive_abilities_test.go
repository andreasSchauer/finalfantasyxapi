package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/sections/sectiontypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/tests/apitesttypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

func TestGetOverdriveAbility(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.OverdriveAbility{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/overdrive-abilities/191",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "overdrive ability with provided id '191' doesn't exist. max id: 190",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/overdrive-abilities/20",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"overdrives":          2,
					"battle interactions": 1,
				},
			},
			NameVer:          apitesttypes.NewExpNameVer(20, "thunder shot", 1),
			Rank:             h.GetInt32Ptr(4),
			UntypedAbility:   122,
			OverdriveCommand: h.GetInt32Ptr(3),
			Overdrives:       []int32{13, 16},
			BattleInteractions: []apitesttypes.BattleInteraction{
				{
					Target:         "random-enemy",
					AbilityRange:   h.GetInt32Ptr(3),
					HitAmount:      1,
					ShatterRate:    100,
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
							DamageFormula:  database.DamageFormulaStrVsDef,
							DamageConstant: 34,
						},
						Critical:        nil,
						CriticalPlusVal: nil,
						IsPiercing:      true,
						BreakDmgLmt:     h.GetStrPtr("auto-ability"),
						Element:         h.GetInt32Ptr(2),
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
				RequestURL:     "/api/overdrive-abilities/159",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"overdrives":          1,
					"battle interactions": 1,
					"battle interactions 0 inflicted status conditions": 1,
					"battle interactions 0 removed status conditions":   10,
				},
			},
			NameVer:          apitesttypes.NewExpNameVer(159, "final elixir", 0),
			Rank:             h.GetInt32Ptr(6),
			UntypedAbility:   261,
			OverdriveCommand: h.GetInt32Ptr(7),
			Overdrives:       []int32{93},
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
					Damage: &apitesttypes.Damage{
						AttackType: database.AttackTypeHeal,
						DamageType: database.DamageTypeSpecial,
						Hp: &apitesttypes.ClassDamage{
							DamageFormula:  database.DamageFormulaPercentageMax,
							DamageConstant: 16,
						},
						Mp: &apitesttypes.ClassDamage{
							DamageFormula:  database.DamageFormulaPercentageMax,
							DamageConstant: 16,
						},
						Critical:        nil,
						CriticalPlusVal: nil,
						IsPiercing:      true,
						BreakDmgLmt:     h.GetStrPtr("always"),
						Element:         nil,
					},
					InflictedDelay:            nil,
					InflictedStatusConditions: []int32{40},
					RemovedStatusConditions:   []int32{1, 2, 3, 4, 9, 10, 12, 13, 14, 16},
					CopiedStatusConditions:    []int32{},
					StatChanges:               []apitesttypes.StatChange{},
					ModifierChanges:           []apitesttypes.ModChange{},
				},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/overdrive-abilities/163",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"battle interactions":                               1,
					"battle interactions 0 inflicted status conditions": 4,
					"battle interactions 0 stat changes":                2,
					"battle interactions 0 mod changes":                 4,
				},
			},
			NameVer:          apitesttypes.NewExpNameVer(163, "ultra nulall", 0),
			Rank:             h.GetInt32Ptr(6),
			UntypedAbility:   265,
			OverdriveCommand: h.GetInt32Ptr(7),
			Overdrives:       []int32{97},
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
					InflictedStatusConditions: []int32{22, 23, 24, 25},
					RemovedStatusConditions:   []int32{},
					CopiedStatusConditions:    []int32{},
					StatChanges: []apitesttypes.StatChange{
						{
							Stat:            3,
							CalculationType: "added-value",
							Value:           5,
						},
						{
							Stat:            5,
							CalculationType: "added-value",
							Value:           5,
						},
					},
					ModifierChanges: []apitesttypes.ModChange{
						{
							Modifier:        3,
							CalculationType: "added-value",
							Value:           -0.33,
						},
						{
							Modifier:        4,
							CalculationType: "added-value",
							Value:           -0.33,
						},
						{
							Modifier:        6,
							CalculationType: "added-percentage",
							Value:           50,
						},
						{
							Modifier:        7,
							CalculationType: "added-percentage",
							Value:           50,
						},
					},
				},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/overdrive-abilities/7",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"overdrives":          1,
					"battle interactions": 2,
				},
			},
			NameVer:          apitesttypes.NewExpNameVer(7, "blitz ace", 1),
			Rank:             h.GetInt32Ptr(7),
			UntypedAbility:   109,
			OverdriveCommand: h.GetInt32Ptr(1),
			Overdrives:       []int32{4},
			BattleInteractions: []apitesttypes.BattleInteraction{
				{
					Target:         "single-enemy",
					AbilityRange:   h.GetInt32Ptr(1),
					HitAmount:      8,
					ShatterRate:    100,
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
							DamageFormula:  database.DamageFormulaStrVsDef,
							DamageConstant: 4,
						},
						Critical:        h.GetStrPtr("crit"),
						CriticalPlusVal: nil,
						IsPiercing:      true,
						BreakDmgLmt:     h.GetStrPtr("auto-ability"),
						Element:         nil,
					},
					InflictedDelay:            nil,
					InflictedStatusConditions: []int32{},
					RemovedStatusConditions:   []int32{},
					CopiedStatusConditions:    []int32{},
					StatChanges:               []apitesttypes.StatChange{},
					ModifierChanges:           []apitesttypes.ModChange{},
				},
				{
					Target:         "single-enemy",
					AbilityRange:   h.GetInt32Ptr(1),
					HitAmount:      1,
					ShatterRate:    100,
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
							DamageFormula:  database.DamageFormulaStrVsDef,
							DamageConstant: 24,
						},
						Critical:        h.GetStrPtr("crit"),
						CriticalPlusVal: nil,
						IsPiercing:      true,
						BreakDmgLmt:     h.GetStrPtr("auto-ability"),
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
	}

	testSingleResources(t, tests, "GetOverdriveAbility", testSrv.HandleOverdriveAbilities)
}

func TestGetMultipleOverdriveAbilities(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.ListIDs{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/overdrive-abilities/grand_summon",
				ExpectedStatus: http.StatusMultipleChoices,
			},
			Count:   8,
			Results: []int32{9, 10, 11, 12, 13, 14, 15, 16},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/overdrive-abilities/dragon_fang",
				ExpectedStatus: http.StatusMultipleChoices,
			},
			Count:   3,
			Results: []int32{107, 108, 109},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/overdrive-abilities/fire_fury",
				ExpectedStatus: http.StatusMultipleChoices,
			},
			Count:   3,
			Results: []int32{40, 41, 42},
		},
	}

	testIDList(t, tests, apilookups.E.OverdriveAbilities.Endpoint, "GetMultipleOverdriveAbilities", testSrv.HandleOverdriveAbilities, compareAPIResourceLists[apiresource.Named])
}

func TestRetrieveOverdriveAbilities(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.ListIDs{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/overdrive-abilities?limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:    190,
			Previous: nil,
			Next:     nil,
			Results:  []int32{1, 26, 83, 99, 176, 190},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/overdrive-abilities?rank=6&can_crit=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:    14,
			Previous: nil,
			Next:     nil,
			Results:  []int32{113, 115, 120, 123, 126, 127, 190},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/overdrive-abilities?status_remove=10",
				ExpectedStatus: http.StatusOK,
			},
			Count:    4,
			Previous: nil,
			Next:     nil,
			Results:  []int32{152, 153, 158, 159},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/overdrive-abilities?stat_changes=true&mod_changes=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:    4,
			Previous: nil,
			Next:     nil,
			Results:  []int32{162, 163, 170, 173},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/overdrive-abilities?user=5",
				ExpectedStatus: http.StatusOK,
			},
			Count:    8,
			Previous: nil,
			Next:     nil,
			Results:  []int32{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/overdrive-abilities?related_stat=5",
				ExpectedStatus: http.StatusOK,
			},
			Count:    3,
			Previous: nil,
			Next:     nil,
			Results:  []int32{162, 163, 173},
		},
	}

	testIDList(t, tests, apilookups.E.OverdriveAbilities.Endpoint, "RetrieveOverdriveAbilities", testSrv.HandleOverdriveAbilities, compareAPIResourceLists[apiresource.Named])
}

func TestSubsectionOverdriveAbilities(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.ListIDs{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/overdrives/14/overdrive-abilities",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleOverdrives,
			},
			Count:          6,
			ParentResource: h.GetStrPtr("/overdrives/14"),
			Results:        []int32{26, 27, 28, 29, 30, 31},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/overdrive-commands/1/overdrive-abilities",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleOverdriveCommands,
			},
			Count:          8,
			ParentResource: h.GetStrPtr("/overdrive-commands/1"),
			Results:        []int32{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/characters/1/overdrive-abilities",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleCharacters,
			},
			Count:          8,
			ParentResource: h.GetStrPtr("/characters/1"),
			Results:        []int32{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/aeons/1/overdrive-abilities",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleAeons,
			},
			Count:          2,
			ParentResource: h.GetStrPtr("/aeons/1"),
			Results:        []int32{183, 184},
		},
	}

	testIDList(t, tests, apilookups.E.OverdriveAbilities.Endpoint, "SubsectionOverdriveAbilities", nil, compareSimpleResourceLists[apiresource.Named, sectiontypes.AbilitySimple])
}
