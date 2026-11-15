package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/sections/sectiontypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

func TestGetOverdriveAbility(t *testing.T) {
	t.Parallel()
	tests := []OverdriveAbility{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/overdrive-abilities/191",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "overdrive ability with provided id '191' doesn't exist. max id: 190",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/overdrive-abilities/20",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"overdrives":          2,
					"battle interactions": 1,
				},
			},
			NameVer:          newExpNameVer(20, "thunder shot", 1),
			rank:             h.GetInt32Ptr(4),
			untypedAbility:   122,
			overdriveCommand: h.GetInt32Ptr(3),
			overdrives:       []int32{13, 16},
			battleInteractions: []BattleInteraction{
				{
					target:         "random-enemy",
					abilityRange:   h.GetInt32Ptr(3),
					hitAmount:      1,
					shatterRate:    100,
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
							damageFormula:  database.DamageFormulaStrVsDef,
							damageConstant: 34,
						},
						critical:        nil,
						criticalPlusVal: nil,
						isPiercing:      true,
						breakDmgLmt:     h.GetStrPtr("auto-ability"),
						element:         h.GetInt32Ptr(2),
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
			NameVer:          newExpNameVer(159, "final elixir", 0),
			rank:             h.GetInt32Ptr(6),
			untypedAbility:   261,
			overdriveCommand: h.GetInt32Ptr(7),
			overdrives:       []int32{93},
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
					damage: &Damage{
						attackType: database.AttackTypeHeal,
						damageType: database.DamageTypeSpecial,
						hp: &ClassDamage{
							damageFormula:  database.DamageFormulaPercentageMax,
							damageConstant: 16,
						},
						mp: &ClassDamage{
							damageFormula:  database.DamageFormulaPercentageMax,
							damageConstant: 16,
						},
						critical:        nil,
						criticalPlusVal: nil,
						isPiercing:      true,
						breakDmgLmt:     h.GetStrPtr("always"),
						element:         nil,
					},
					inflictedDelay:            nil,
					inflictedStatusConditions: []int32{40},
					removedStatusConditions:   []int32{1, 2, 3, 4, 9, 10, 12, 13, 14, 16},
					copiedStatusConditions:    []int32{},
					statChanges:               []StatChange{},
					modifierChanges:           []ModChange{},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
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
			NameVer:          newExpNameVer(163, "ultra nulall", 0),
			rank:             h.GetInt32Ptr(6),
			untypedAbility:   265,
			overdriveCommand: h.GetInt32Ptr(7),
			overdrives:       []int32{97},
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
					inflictedStatusConditions: []int32{22, 23, 24, 25},
					removedStatusConditions:   []int32{},
					copiedStatusConditions:    []int32{},
					statChanges: []StatChange{
						{
							stat:            3,
							calculationType: "added-value",
							value:           5,
						},
						{
							stat:            5,
							calculationType: "added-value",
							value:           5,
						},
					},
					modifierChanges: []ModChange{
						{
							modifier:        3,
							calculationType: "added-value",
							value:           -0.33,
						},
						{
							modifier:        4,
							calculationType: "added-value",
							value:           -0.33,
						},
						{
							modifier:        6,
							calculationType: "added-percentage",
							value:           50,
						},
						{
							modifier:        7,
							calculationType: "added-percentage",
							value:           50,
						},
					},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/overdrive-abilities/7",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"overdrives":          1,
					"battle interactions": 2,
				},
			},
			NameVer:          newExpNameVer(7, "blitz ace", 1),
			rank:             h.GetInt32Ptr(7),
			untypedAbility:   109,
			overdriveCommand: h.GetInt32Ptr(1),
			overdrives:       []int32{4},
			battleInteractions: []BattleInteraction{
				{
					target:         "single-enemy",
					abilityRange:   h.GetInt32Ptr(1),
					hitAmount:      8,
					shatterRate:    100,
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
							damageFormula:  database.DamageFormulaStrVsDef,
							damageConstant: 4,
						},
						critical:        h.GetStrPtr("crit"),
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
				{
					target:         "single-enemy",
					abilityRange:   h.GetInt32Ptr(1),
					hitAmount:      1,
					shatterRate:    100,
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
							damageFormula:  database.DamageFormulaStrVsDef,
							damageConstant: 24,
						},
						critical:        h.GetStrPtr("crit"),
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

	testSingleResources(t, tests, "GetOverdriveAbility", testSrv.HandleOverdriveAbilities)
}

func TestGetMultipleOverdriveAbilities(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/overdrive-abilities/grand_summon",
				ExpectedStatus: http.StatusMultipleChoices,
			},
			Count:   8,
			Results: []int32{9, 10, 11, 12, 13, 14, 15, 16},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/overdrive-abilities/dragon_fang",
				ExpectedStatus: http.StatusMultipleChoices,
			},
			Count:   3,
			Results: []int32{107, 108, 109},
		},
		{
			TestGeneral: TestGeneral{
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
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/overdrive-abilities?limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:    190,
			Previous: nil,
			Next:     nil,
			Results:  []int32{1, 26, 83, 99, 176, 190},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/overdrive-abilities?rank=6&can_crit=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:    14,
			Previous: nil,
			Next:     nil,
			Results:  []int32{113, 115, 120, 123, 126, 127, 190},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/overdrive-abilities?status_remove=10",
				ExpectedStatus: http.StatusOK,
			},
			Count:    4,
			Previous: nil,
			Next:     nil,
			Results:  []int32{152, 153, 158, 159},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/overdrive-abilities?stat_changes=true&mod_changes=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:    4,
			Previous: nil,
			Next:     nil,
			Results:  []int32{162, 163, 170, 173},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/overdrive-abilities?user=5",
				ExpectedStatus: http.StatusOK,
			},
			Count:    8,
			Previous: nil,
			Next:     nil,
			Results:  []int32{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			TestGeneral: TestGeneral{
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
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/overdrives/14/overdrive-abilities",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleOverdrives,
			},
			Count:          6,
			ParentResource: h.GetStrPtr("/overdrives/14"),
			Results:        []int32{26, 27, 28, 29, 30, 31},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/overdrive-commands/1/overdrive-abilities",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleOverdriveCommands,
			},
			Count:          8,
			ParentResource: h.GetStrPtr("/overdrive-commands/1"),
			Results:        []int32{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/characters/1/overdrive-abilities",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleCharacters,
			},
			Count:          8,
			ParentResource: h.GetStrPtr("/characters/1"),
			Results:        []int32{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			TestGeneral: TestGeneral{
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
