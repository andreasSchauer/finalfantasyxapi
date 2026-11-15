package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apitesttypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/sections/sectiontypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

func TestGetPlayerAbility(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.PlayerAbility{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/player-abilities/103",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "player ability with provided id '103' doesn't exist. max id: 102",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/player-abilities/10?ability_user=shiva",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"monsters":            0,
					"battle interactions": 1,
				},
			},
			NameVer:          apitesttypes.NewExpNameVer(10, "delay buster", 0),
			Rank:             h.GetInt32Ptr(8),
			AppearsInHelpBar: true,
			CanCopyCat:       true,
			UntypedAbility:   10,
			Topmenu:          nil,
			Submenu:          h.GetInt32Ptr(2),
			OpenSubmenu:      nil,
			StdChar:          h.GetInt32Ptr(1),
			ExpChar:          h.GetInt32Ptr(1),
			Monsters:         []int32{},
			BattleInteractions: []apitesttypes.BattleInteraction{
				{
					Target:         "single-target",
					AbilityRange:   h.GetInt32Ptr(1),
					HitAmount:      1,
					ShatterRate:    30,
					BasedOnUserAtk: true,
					Darkable:       true,
					Silenceable:    false,
					Reflectable:    false,
					Accuracy: apitesttypes.Accuracy{
						AccSource:   "accuracy",
						HitChance:   nil,
						AccModifier: h.GetFloat32Ptr(2.5),
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
						IsPiercing:      false,
						BreakDmgLmt:     h.GetStrPtr("auto-ability"),
						Element:         nil,
					},
					InflictedDelay: &apitesttypes.InflictedDelay{
						DelayStrength: "strong",
					},
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
				RequestURL:     "/api/player-abilities/10?ability_user=auron&celestial_weapon=true",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"monsters":            0,
					"battle interactions": 1,
				},
			},
			NameVer:          apitesttypes.NewExpNameVer(10, "delay buster", 0),
			Rank:             h.GetInt32Ptr(8),
			AppearsInHelpBar: true,
			CanCopyCat:       true,
			UntypedAbility:   10,
			Topmenu:          nil,
			Submenu:          h.GetInt32Ptr(2),
			OpenSubmenu:      nil,
			StdChar:          h.GetInt32Ptr(1),
			ExpChar:          h.GetInt32Ptr(1),
			Monsters:         []int32{},
			BattleInteractions: []apitesttypes.BattleInteraction{
				{
					Target:         "single-target",
					AbilityRange:   h.GetInt32Ptr(1),
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
							DamageFormula:  database.DamageFormulaClstlHpLow,
							DamageConstant: 16,
						},
						Critical:        h.GetStrPtr("crit-equip"),
						CriticalPlusVal: nil,
						IsPiercing:      false,
						BreakDmgLmt:     h.GetStrPtr("auto-ability"),
						Element:         nil,
					},
					InflictedDelay: &apitesttypes.InflictedDelay{
						DelayStrength: "strong",
					},
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
				RequestURL:     "/api/player-abilities/24",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"battle interactions": 0,
				},
			},
			NameVer:            apitesttypes.NewExpNameVer(24, "use", 0),
			Rank:               h.GetInt32Ptr(2),
			AppearsInHelpBar:   false,
			CanCopyCat:         true,
			UntypedAbility:     24,
			Topmenu:            nil,
			Submenu:            h.GetInt32Ptr(3),
			OpenSubmenu:        h.GetInt32Ptr(9),
			StdChar:            h.GetInt32Ptr(7),
			ExpChar:            h.GetInt32Ptr(7),
			Monsters:           []int32{},
			BattleInteractions: []apitesttypes.BattleInteraction{},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/player-abilities/27",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"monsters":                           0,
					"battle interactions":                1,
					"battle interactions 0 stat changes": 1,
					"battle interactions 0 mod changes":  1,
				},
			},
			NameVer:          apitesttypes.NewExpNameVer(27, "cheer", 0),
			Rank:             h.GetInt32Ptr(2),
			AppearsInHelpBar: true,
			CanCopyCat:       true,
			UntypedAbility:   27,
			Topmenu:          nil,
			Submenu:          h.GetInt32Ptr(3),
			OpenSubmenu:      nil,
			StdChar:          h.GetInt32Ptr(1),
			ExpChar:          h.GetInt32Ptr(1),
			Monsters:         []int32{},
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
					InflictedStatusConditions: []int32{},
					RemovedStatusConditions:   []int32{},
					CopiedStatusConditions:    []int32{},
					StatChanges: []apitesttypes.StatChange{
						{
							Stat:            3,
							CalculationType: "added-value",
							Value:           1,
						},
					},
					ModifierChanges: []apitesttypes.ModChange{
						{
							Modifier:        3,
							CalculationType: "added-value",
							Value:           -0.066,
						},
					},
				},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/player-abilities/47",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"monsters":            7,
					"battle interactions": 1,
				},
			},
			NameVer:          apitesttypes.NewExpNameVer(47, "curaga", 0),
			Rank:             h.GetInt32Ptr(3),
			AppearsInHelpBar: true,
			CanCopyCat:       true,
			UntypedAbility:   47,
			Topmenu:          nil,
			Submenu:          h.GetInt32Ptr(4),
			OpenSubmenu:      nil,
			StdChar:          h.GetInt32Ptr(2),
			ExpChar:          h.GetInt32Ptr(2),
			Monsters:         []int32{187, 195, 223, 226, 272, 275, 287},
			BattleInteractions: []apitesttypes.BattleInteraction{
				{
					Target:         "single-target",
					AbilityRange:   h.GetInt32Ptr(4),
					HitAmount:      1,
					ShatterRate:    0,
					BasedOnUserAtk: false,
					Darkable:       false,
					Silenceable:    true,
					Reflectable:    true,
					Accuracy: apitesttypes.Accuracy{
						AccSource:   "rate",
						HitChance:   h.GetInt32Ptr(255),
						AccModifier: nil,
					},
					Damage: &apitesttypes.Damage{
						AttackType: database.AttackTypeHeal,
						DamageType: database.DamageTypeMagical,
						Hp: &apitesttypes.ClassDamage{
							DamageFormula:  database.DamageFormulaHealing,
							DamageConstant: 80,
						},
						Critical:        nil,
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

	testSingleResources(t, tests, "GetPlayerAbility", testSrv.HandlePlayerAbilities)
}

func TestGetMultiplePlayerAbilities(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.ListIDs{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/player-abilities/wakizashi",
				ExpectedStatus: http.StatusMultipleChoices,
			},
			Count:   2,
			Results: []int32{96, 97},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/player-abilities/auto-life",
				ExpectedStatus: http.StatusMultipleChoices,
			},
			Count:   2,
			Results: []int32{66, 67},
		},
	}

	testIDList(t, tests, apilookups.E.PlayerAbilities.Endpoint, "GetMultiplePlayerAbilities", testSrv.HandlePlayerAbilities, compareAPIResourceLists[apiresource.Named])
}

func TestRetrievePlayerAbilities(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.ListIDs{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/player-abilities?limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   102,
			Results: []int32{1, 22, 57, 78, 99, 102},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/player-abilities?category=2&help_bar=false",
				ExpectedStatus: http.StatusOK,
			},
			Count:    4,
			Previous: nil,
			Next:     nil,
			Results:  []int32{24, 40, 42, 43},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/player-abilities?related_stat=3&darkable=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:    1,
			Previous: nil,
			Next:     nil,
			Results:  []int32{11},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/player-abilities?copycat=true&std_sg=3&user_atk=false",
				ExpectedStatus: http.StatusOK,
			},
			Count:    3,
			Previous: nil,
			Next:     nil,
			Results:  []int32{28, 84, 85},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/player-abilities?status_inflict=none&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:    54,
			Previous: nil,
			Next:     nil,
			Results:  []int32{20, 24, 28, 36, 39, 43, 53, 76, 86, 96, 102},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/player-abilities?silenceable=true&reflectable=false&user=3",
				ExpectedStatus: http.StatusOK,
			},
			Count:    2,
			Previous: nil,
			Next:     nil,
			Results:  []int32{82, 87},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/player-abilities?mp=20&exp_sg=2",
				ExpectedStatus: http.StatusOK,
			},
			Count:    1,
			Previous: nil,
			Next:     nil,
			Results:  []int32{47},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/player-abilities?mp_min=20&learn_item=59",
				ExpectedStatus: http.StatusOK,
			},
			Count:    1,
			Previous: nil,
			Next:     nil,
			Results:  []int32{64},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/player-abilities?mp_max=10&element=3",
				ExpectedStatus: http.StatusOK,
			},
			Count:    2,
			Previous: nil,
			Next:     nil,
			Results:  []int32{71, 75},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/player-abilities?element=fire,ice",
				ExpectedStatus: http.StatusOK,
			},
			Count:    6,
			Previous: nil,
			Next:     nil,
			Results:  []int32{69, 72, 73, 76, 77, 80},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/player-abilities?element=none&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:    89,
			Previous: nil,
			Next:     nil,
			Results:  []int32{1, 12, 24, 34, 43, 55, 67, 95, 102},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/player-abilities?delay=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:    3,
			Previous: nil,
			Next:     nil,
			Results:  []int32{9, 10, 88},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/player-abilities?outside_battle=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:    3,
			Previous: nil,
			Next:     nil,
			Results:  []int32{45, 46, 47},
		},
	}

	testIDList(t, tests, apilookups.E.PlayerAbilities.Endpoint, "RetrievePlayerAbilities", testSrv.HandlePlayerAbilities, compareAPIResourceLists[apiresource.Named])
}

func TestSubsectionPlayerAbilities(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.ListIDs{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/characters/2/default-abilities",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleCharacters,
			},
			Count:          2,
			ParentResource: h.GetStrPtr("/characters/2"),
			Results:        []int32{45, 53},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/characters/2/std-sg-abilities",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleCharacters,
			},
			Count:          18,
			ParentResource: h.GetStrPtr("/characters/2"),
			Results:        []int32{26, 42, 47, 53, 60, 62, 66},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/characters/2/exp-sg-abilities",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleCharacters,
			},
			Count:          19,
			ParentResource: h.GetStrPtr("/characters/2"),
			Results:        []int32{14, 45, 49, 54, 60, 64, 65},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/aeons/1/default-abilities",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleAeons,
			},
			Count:          5,
			ParentResource: h.GetStrPtr("/aeons/1"),
			Results:        []int32{69, 70, 71, 72, 88},
		},
	}

	testIDList(t, tests, apilookups.E.PlayerAbilities.Endpoint, "SubsectionPlayerAbilities", nil, compareSimpleResourceLists[apiresource.Named, sectiontypes.PlayerAbilitySimple])
}
