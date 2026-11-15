package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

func TestGetMiscAbility(t *testing.T) {
	t.Parallel()
	tests := []MiscAbility{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/misc-abilities/25",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "misc ability with provided id '25' doesn't exist. max id: 24",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/misc-abilities/2?ability_user=wakka",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "invalid input for parameter 'ability_user': player unit 'wakka' can't learn misc ability 'attack - 2'",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/misc-abilities/2?ability_user=we",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "unknown player unit 'we' used for parameter 'ability_user'",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/misc-abilities/2?bomb_wpn=true",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "invalid usage of parameter 'bomb_wpn'. when using parameter 'bomb_wpn', the following parameter(s) must be present: 'ability_user'",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/misc-abilities/1?ability_user=8&celestial_weapon=true",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "character seymour doesn't have a celestial weapon",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/misc-abilities/1?ability_user=9&celestial_weapon=true",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "aeons don't use celestial weapons",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/misc-abilities/1?ability_user=wakka&celestial_weapon=true",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"battle interactions": 1,
				},
			},
			NameVer:          newExpNameVer(1, "attack", 1),
			rank:             h.GetInt32Ptr(3),
			appearsInHelpBar: false,
			canCopyCat:       true,
			untypedAbility:   374,
			topmenu:          h.GetInt32Ptr(1),
			submenu:          nil,
			openSubmenu:      nil,
			battleInteractions: []BattleInteraction{
				{
					target:         "single-target",
					abilityRange:   h.GetInt32Ptr(4),
					hitAmount:      1,
					shatterRate:    30,
					basedOnUserAtk: true,
					darkable:       true,
					silenceable:    false,
					reflectable:    false,
					accuracy: Accuracy{
						accSource:   "accuracy",
						hitChance:   nil,
						accModifier: h.GetFloat32Ptr(1),
					},
					damage: &Damage{
						attackType: database.AttackTypeAttack,
						damageType: database.DamageTypePhysical,
						hp: &ClassDamage{
							damageFormula:  database.DamageFormulaClstlHpHigh,
							damageConstant: 16,
						},
						critical:        h.GetStrPtr("crit-equip"),
						criticalPlusVal: nil,
						isPiercing:      false,
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
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/misc-abilities/1?ability_user=3&bomb_wpn=true",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"battle interactions": 1,
				},
			},
			NameVer:          newExpNameVer(1, "attack", 1),
			rank:             h.GetInt32Ptr(3),
			appearsInHelpBar: false,
			canCopyCat:       true,
			untypedAbility:   374,
			topmenu:          h.GetInt32Ptr(1),
			submenu:          nil,
			openSubmenu:      nil,
			battleInteractions: []BattleInteraction{
				{
					target:         "single-target",
					abilityRange:   h.GetInt32Ptr(4),
					hitAmount:      1,
					shatterRate:    30,
					basedOnUserAtk: true,
					darkable:       true,
					silenceable:    false,
					reflectable:    false,
					accuracy: Accuracy{
						accSource:   "accuracy",
						hitChance:   nil,
						accModifier: h.GetFloat32Ptr(1),
					},
					damage: &Damage{
						attackType: database.AttackTypeAttack,
						damageType: database.DamageTypePhysical,
						hp: &ClassDamage{
							damageFormula:  database.DamageFormulaStrVsDef,
							damageConstant: 18,
						},
						critical:        h.GetStrPtr("crit-equip"),
						criticalPlusVal: nil,
						isPiercing:      false,
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
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/misc-abilities/2?ability_user=16&bomb_wpn=true",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"battle interactions": 1,
				},
			},
			NameVer:          newExpNameVer(2, "attack", 2),
			rank:             h.GetInt32Ptr(3),
			appearsInHelpBar: false,
			canCopyCat:       false,
			untypedAbility:   375,
			topmenu:          h.GetInt32Ptr(1),
			submenu:          nil,
			openSubmenu:      nil,
			battleInteractions: []BattleInteraction{
				{
					target:         "single-target",
					abilityRange:   h.GetInt32Ptr(1),
					hitAmount:      1,
					shatterRate:    20,
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
							damageConstant: 14,
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
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/misc-abilities/10",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"battle interactions": 0,
				},
			},
			NameVer:            newExpNameVer(10, "switch", 1),
			rank:               h.GetInt32Ptr(0),
			appearsInHelpBar:   false,
			canCopyCat:         false,
			untypedAbility:     383,
			topmenu:            nil,
			submenu:            h.GetInt32Ptr(13),
			openSubmenu:        nil,
			battleInteractions: []BattleInteraction{},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/misc-abilities/4",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"battle interactions": 0,
				},
			},
			NameVer:            newExpNameVer(4, "weapon", 0),
			rank:               h.GetInt32Ptr(1),
			appearsInHelpBar:   false,
			canCopyCat:         false,
			untypedAbility:     377,
			topmenu:            h.GetInt32Ptr(2),
			submenu:            nil,
			openSubmenu:        h.GetInt32Ptr(7),
			battleInteractions: []BattleInteraction{},
		},
	}

	testSingleResources(t, tests, "GetMiscAbility", testSrv.HandleMiscAbilities)
}

func TestGetMultipleMiscAbilities(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/misc-abilities/attack",
				ExpectedStatus: http.StatusMultipleChoices,
			},
			Count:   2,
			Results: []int32{1, 2},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/misc-abilities/switch",
				ExpectedStatus: http.StatusMultipleChoices,
			},
			Count:   7,
			Results: []int32{10, 11, 12, 13, 14, 15, 16},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/misc-abilities/summon",
				ExpectedStatus: http.StatusMultipleChoices,
			},
			Count:   8,
			Results: []int32{17, 18, 19, 20, 21, 22, 23, 24},
		},
	}

	testIDList(t, tests, apilookups.E.MiscAbilities.Endpoint, "GetMultipleMiscAbilities", testSrv.HandleMiscAbilities, compareAPIResourceLists[apiresource.Named])
}

func TestRetrieveMiscAbilities(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/misc-abilities?limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:    24,
			Previous: nil,
			Next:     nil,
			Results:  []int32{1, 8, 13, 17, 18, 24},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/misc-abilities?rank=1",
				ExpectedStatus: http.StatusOK,
			},
			Count:    3,
			Previous: nil,
			Next:     nil,
			Results:  []int32{4, 5, 6},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/misc-abilities?help_bar=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:    3,
			Previous: nil,
			Next:     nil,
			Results:  []int32{7, 8, 9},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/misc-abilities?user=characters&copycat=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:    3,
			Previous: nil,
			Next:     nil,
			Results:  []int32{1, 3, 6},
		},
	}

	testIDList(t, tests, apilookups.E.MiscAbilities.Endpoint, "RetrieveMiscAbilities", testSrv.HandleMiscAbilities, compareAPIResourceLists[apiresource.Named])
}
