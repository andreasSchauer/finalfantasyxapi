package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/tests/apitesttypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

func TestGetMiscAbility(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.MiscAbility{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/misc-abilities/25",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "misc ability with provided id '25' doesn't exist. max id: 24",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/misc-abilities/2?ability_user=wakka",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "invalid input for parameter 'ability_user': player unit 'wakka' can't learn misc ability 'attack - 2'",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/misc-abilities/2?ability_user=we",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "unknown player unit 'we' used for parameter 'ability_user'",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/misc-abilities/2?bomb_wpn=true",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "invalid usage of parameter 'bomb_wpn'. when using parameter 'bomb_wpn', the following parameter(s) must be present: 'ability_user'",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/misc-abilities/1?ability_user=8&celestial_weapon=true",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "character seymour doesn't have a celestial weapon",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/misc-abilities/1?ability_user=9&celestial_weapon=true",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "aeons don't use celestial weapons",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/misc-abilities/1?ability_user=wakka&celestial_weapon=true",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"battle interactions": 1,
				},
			},
			NameVer:          apitesttypes.NewExpNameVer(1, "attack", 1),
			Rank:             h.GetInt32Ptr(3),
			AppearsInHelpBar: false,
			CanCopyCat:       true,
			UntypedAbility:   374,
			Topmenu:          h.GetInt32Ptr(1),
			Submenu:          nil,
			OpenSubmenu:      nil,
			BattleInteractions: []apitesttypes.BattleInteraction{
				{
					Target:         "single-target",
					AbilityRange:   h.GetInt32Ptr(4),
					HitAmount:      1,
					ShatterRate:    30,
					BasedOnUserAtk: true,
					Darkable:       true,
					Silenceable:    false,
					Reflectable:    false,
					Accuracy: apitesttypes.Accuracy{
						AccSource:   "accuracy",
						HitChance:   nil,
						AccModifier: h.GetFloat32Ptr(1),
					},
					Damage: &apitesttypes.Damage{
						AttackType: database.AttackTypeAttack,
						DamageType: database.DamageTypePhysical,
						Hp: &apitesttypes.ClassDamage{
							DamageFormula:  database.DamageFormulaClstlHpHigh,
							DamageConstant: 16,
						},
						Critical:        h.GetStrPtr("crit-equip"),
						CriticalPlusVal: nil,
						IsPiercing:      false,
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
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/misc-abilities/1?ability_user=3&bomb_wpn=true",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"battle interactions": 1,
				},
			},
			NameVer:          apitesttypes.NewExpNameVer(1, "attack", 1),
			Rank:             h.GetInt32Ptr(3),
			AppearsInHelpBar: false,
			CanCopyCat:       true,
			UntypedAbility:   374,
			Topmenu:          h.GetInt32Ptr(1),
			Submenu:          nil,
			OpenSubmenu:      nil,
			BattleInteractions: []apitesttypes.BattleInteraction{
				{
					Target:         "single-target",
					AbilityRange:   h.GetInt32Ptr(4),
					HitAmount:      1,
					ShatterRate:    30,
					BasedOnUserAtk: true,
					Darkable:       true,
					Silenceable:    false,
					Reflectable:    false,
					Accuracy: apitesttypes.Accuracy{
						AccSource:   "accuracy",
						HitChance:   nil,
						AccModifier: h.GetFloat32Ptr(1),
					},
					Damage: &apitesttypes.Damage{
						AttackType: database.AttackTypeAttack,
						DamageType: database.DamageTypePhysical,
						Hp: &apitesttypes.ClassDamage{
							DamageFormula:  database.DamageFormulaStrVsDef,
							DamageConstant: 18,
						},
						Critical:        h.GetStrPtr("crit-equip"),
						CriticalPlusVal: nil,
						IsPiercing:      false,
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
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/misc-abilities/2?ability_user=16&bomb_wpn=true",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"battle interactions": 1,
				},
			},
			NameVer:          apitesttypes.NewExpNameVer(2, "attack", 2),
			Rank:             h.GetInt32Ptr(3),
			AppearsInHelpBar: false,
			CanCopyCat:       false,
			UntypedAbility:   375,
			Topmenu:          h.GetInt32Ptr(1),
			Submenu:          nil,
			OpenSubmenu:      nil,
			BattleInteractions: []apitesttypes.BattleInteraction{
				{
					Target:         "single-target",
					AbilityRange:   h.GetInt32Ptr(1),
					HitAmount:      1,
					ShatterRate:    20,
					BasedOnUserAtk: true,
					Darkable:       true,
					Silenceable:    false,
					Reflectable:    false,
					Accuracy: apitesttypes.Accuracy{
						AccSource:   "rate",
						HitChance:   h.GetInt32Ptr(255),
						AccModifier: nil,
					},
					Damage: &apitesttypes.Damage{
						AttackType: database.AttackTypeAttack,
						DamageType: database.DamageTypePhysical,
						Hp: &apitesttypes.ClassDamage{
							DamageFormula:  database.DamageFormulaStrVsDef,
							DamageConstant: 14,
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
					StatChanges:               []apitesttypes.StatChange{},
					ModifierChanges:           []apitesttypes.ModChange{},
				},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/misc-abilities/10",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"battle interactions": 0,
				},
			},
			NameVer:            apitesttypes.NewExpNameVer(10, "switch", 1),
			Rank:               h.GetInt32Ptr(0),
			AppearsInHelpBar:   false,
			CanCopyCat:         false,
			UntypedAbility:     383,
			Topmenu:            nil,
			Submenu:            h.GetInt32Ptr(13),
			OpenSubmenu:        nil,
			BattleInteractions: []apitesttypes.BattleInteraction{},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/misc-abilities/4",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"battle interactions": 0,
				},
			},
			NameVer:            apitesttypes.NewExpNameVer(4, "weapon", 0),
			Rank:               h.GetInt32Ptr(1),
			AppearsInHelpBar:   false,
			CanCopyCat:         false,
			UntypedAbility:     377,
			Topmenu:            h.GetInt32Ptr(2),
			Submenu:            nil,
			OpenSubmenu:        h.GetInt32Ptr(7),
			BattleInteractions: []apitesttypes.BattleInteraction{},
		},
	}

	testSingleResources(t, tests, "GetMiscAbility", testSrv.HandleMiscAbilities)
}

func TestGetMultipleMiscAbilities(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.ListIDs{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/misc-abilities/attack",
				ExpectedStatus: http.StatusMultipleChoices,
			},
			Count:   2,
			Results: []int32{1, 2},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/misc-abilities/switch",
				ExpectedStatus: http.StatusMultipleChoices,
			},
			Count:   7,
			Results: []int32{10, 11, 12, 13, 14, 15, 16},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
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
	tests := []apitesttypes.ListIDs{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/misc-abilities?limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:    24,
			Previous: nil,
			Next:     nil,
			Results:  []int32{1, 8, 13, 17, 18, 24},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/misc-abilities?rank=1",
				ExpectedStatus: http.StatusOK,
			},
			Count:    3,
			Previous: nil,
			Next:     nil,
			Results:  []int32{4, 5, 6},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/misc-abilities?help_bar=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:    3,
			Previous: nil,
			Next:     nil,
			Results:  []int32{7, 8, 9},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
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
