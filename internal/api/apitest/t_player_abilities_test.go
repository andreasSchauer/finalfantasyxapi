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

func TestGetPlayerAbility(t *testing.T) {
	t.Parallel()
	tests := []PlayerAbility{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/player-abilities/103",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "player ability with provided id '103' doesn't exist. max id: 102",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/player-abilities/10?ability_user=shiva",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"monsters":            0,
					"battle interactions": 1,
				},
			},
			NameVer:          newExpNameVer(10, "delay buster", 0),
			rank:             h.GetInt32Ptr(8),
			appearsInHelpBar: true,
			canCopyCat:       true,
			untypedAbility:   10,
			topmenu:          nil,
			submenu:          h.GetInt32Ptr(2),
			openSubmenu:      nil,
			stdChar:          h.GetInt32Ptr(1),
			expChar:          h.GetInt32Ptr(1),
			monsters:         []int32{},
			battleInteractions: []BattleInteraction{
				{
					target:         "single-target",
					abilityRange:   h.GetInt32Ptr(1),
					hitAmount:      1,
					shatterRate:    30,
					basedOnUserAtk: true,
					darkable:       true,
					silenceable:    false,
					reflectable:    false,
					accuracy: Accuracy{
						accSource:   "accuracy",
						hitChance:   nil,
						accModifier: h.GetFloat32Ptr(2.5),
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
						isPiercing:      false,
						breakDmgLmt:     h.GetStrPtr("auto-ability"),
						element:         nil,
					},
					inflictedDelay: &InflictedDelay{
						delayStrength: "strong",
					},
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
				RequestURL:     "/api/player-abilities/10?ability_user=auron&celestial_weapon=true",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"monsters":            0,
					"battle interactions": 1,
				},
			},
			NameVer:          newExpNameVer(10, "delay buster", 0),
			rank:             h.GetInt32Ptr(8),
			appearsInHelpBar: true,
			canCopyCat:       true,
			untypedAbility:   10,
			topmenu:          nil,
			submenu:          h.GetInt32Ptr(2),
			openSubmenu:      nil,
			stdChar:          h.GetInt32Ptr(1),
			expChar:          h.GetInt32Ptr(1),
			monsters:         []int32{},
			battleInteractions: []BattleInteraction{
				{
					target:         "single-target",
					abilityRange:   h.GetInt32Ptr(1),
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
							damageFormula:  database.DamageFormulaClstlHpLow,
							damageConstant: 16,
						},
						critical:        h.GetStrPtr("crit-equip"),
						criticalPlusVal: nil,
						isPiercing:      false,
						breakDmgLmt:     h.GetStrPtr("auto-ability"),
						element:         nil,
					},
					inflictedDelay: &InflictedDelay{
						delayStrength: "strong",
					},
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
				RequestURL:     "/api/player-abilities/24",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"battle interactions": 0,
				},
			},
			NameVer:            newExpNameVer(24, "use", 0),
			rank:               h.GetInt32Ptr(2),
			appearsInHelpBar:   false,
			canCopyCat:         true,
			untypedAbility:     24,
			topmenu:            nil,
			submenu:            h.GetInt32Ptr(3),
			openSubmenu:        h.GetInt32Ptr(9),
			stdChar:            h.GetInt32Ptr(7),
			expChar:            h.GetInt32Ptr(7),
			monsters:           []int32{},
			battleInteractions: []BattleInteraction{},
		},
		{
			TestGeneral: TestGeneral{
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
			NameVer:          newExpNameVer(27, "cheer", 0),
			rank:             h.GetInt32Ptr(2),
			appearsInHelpBar: true,
			canCopyCat:       true,
			untypedAbility:   27,
			topmenu:          nil,
			submenu:          h.GetInt32Ptr(3),
			openSubmenu:      nil,
			stdChar:          h.GetInt32Ptr(1),
			expChar:          h.GetInt32Ptr(1),
			monsters:         []int32{},
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
					inflictedStatusConditions: []int32{},
					removedStatusConditions:   []int32{},
					copiedStatusConditions:    []int32{},
					statChanges: []StatChange{
						{
							stat:            3,
							calculationType: "added-value",
							value:           1,
						},
					},
					modifierChanges: []ModChange{
						{
							modifier:        3,
							calculationType: "added-value",
							value:           -0.066,
						},
					},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/player-abilities/47",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"monsters":            7,
					"battle interactions": 1,
				},
			},
			NameVer:          newExpNameVer(47, "curaga", 0),
			rank:             h.GetInt32Ptr(3),
			appearsInHelpBar: true,
			canCopyCat:       true,
			untypedAbility:   47,
			topmenu:          nil,
			submenu:          h.GetInt32Ptr(4),
			openSubmenu:      nil,
			stdChar:          h.GetInt32Ptr(2),
			expChar:          h.GetInt32Ptr(2),
			monsters:         []int32{187, 195, 223, 226, 272, 275, 287},
			battleInteractions: []BattleInteraction{
				{
					target:         "single-target",
					abilityRange:   h.GetInt32Ptr(4),
					hitAmount:      1,
					shatterRate:    0,
					basedOnUserAtk: false,
					darkable:       false,
					silenceable:    true,
					reflectable:    true,
					accuracy: Accuracy{
						accSource:   "rate",
						hitChance:   h.GetInt32Ptr(255),
						accModifier: nil,
					},
					damage: &Damage{
						attackType: database.AttackTypeHeal,
						damageType: database.DamageTypeMagical,
						hp: &ClassDamage{
							damageFormula:  database.DamageFormulaHealing,
							damageConstant: 80,
						},
						critical:        nil,
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

	testSingleResources(t, tests, "GetPlayerAbility", testSrv.HandlePlayerAbilities)
}

func TestGetMultiplePlayerAbilities(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/player-abilities/wakizashi",
				ExpectedStatus: http.StatusMultipleChoices,
			},
			Count:   2,
			Results: []int32{96, 97},
		},
		{
			TestGeneral: TestGeneral{
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
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/player-abilities?limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   102,
			Results: []int32{1, 22, 57, 78, 99, 102},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/player-abilities?category=2&help_bar=false",
				ExpectedStatus: http.StatusOK,
			},
			Count:    4,
			Previous: nil,
			Next:     nil,
			Results:  []int32{24, 40, 42, 43},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/player-abilities?related_stat=3&darkable=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:    1,
			Previous: nil,
			Next:     nil,
			Results:  []int32{11},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/player-abilities?copycat=true&std_sg=3&user_atk=false",
				ExpectedStatus: http.StatusOK,
			},
			Count:    3,
			Previous: nil,
			Next:     nil,
			Results:  []int32{28, 84, 85},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/player-abilities?status_inflict=none&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:    54,
			Previous: nil,
			Next:     nil,
			Results:  []int32{20, 24, 28, 36, 39, 43, 53, 76, 86, 96, 102},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/player-abilities?silenceable=true&reflectable=false&user=3",
				ExpectedStatus: http.StatusOK,
			},
			Count:    2,
			Previous: nil,
			Next:     nil,
			Results:  []int32{82, 87},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/player-abilities?mp=20&exp_sg=2",
				ExpectedStatus: http.StatusOK,
			},
			Count:    1,
			Previous: nil,
			Next:     nil,
			Results:  []int32{47},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/player-abilities?mp_min=20&learn_item=59",
				ExpectedStatus: http.StatusOK,
			},
			Count:    1,
			Previous: nil,
			Next:     nil,
			Results:  []int32{64},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/player-abilities?mp_max=10&element=3",
				ExpectedStatus: http.StatusOK,
			},
			Count:    2,
			Previous: nil,
			Next:     nil,
			Results:  []int32{71, 75},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/player-abilities?element=fire,ice",
				ExpectedStatus: http.StatusOK,
			},
			Count:    6,
			Previous: nil,
			Next:     nil,
			Results:  []int32{69, 72, 73, 76, 77, 80},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/player-abilities?element=none&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:    89,
			Previous: nil,
			Next:     nil,
			Results:  []int32{1, 12, 24, 34, 43, 55, 67, 95, 102},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/player-abilities?delay=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:    3,
			Previous: nil,
			Next:     nil,
			Results:  []int32{9, 10, 88},
		},
		{
			TestGeneral: TestGeneral{
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
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/characters/2/default-abilities",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleCharacters,
			},
			Count:          2,
			ParentResource: h.GetStrPtr("/characters/2"),
			Results:        []int32{45, 53},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/characters/2/std-sg-abilities",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleCharacters,
			},
			Count:          18,
			ParentResource: h.GetStrPtr("/characters/2"),
			Results:        []int32{26, 42, 47, 53, 60, 62, 66},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/characters/2/exp-sg-abilities",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleCharacters,
			},
			Count:          19,
			ParentResource: h.GetStrPtr("/characters/2"),
			Results:        []int32{14, 45, 49, 54, 60, 64, 65},
		},
		{
			TestGeneral: TestGeneral{
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
