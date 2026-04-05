package api

import (
	"net/http"
	"testing"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func TestGetPlayerAbility(t *testing.T) {
	tests := []expPlayerAbility{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/player-abilities/103",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "player ability with provided id '103' doesn't exist. max id: 102.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/player-abilities/10?ability_user=shiva",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"monsters":            0,
					"battle interactions": 1,
				},
			},
			expNameVer:       newExpNameVer(10, "delay buster", 0),
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
			battleInteractions: []expBattleInteraction{
				{
					target:         "single-target",
					abilityRange:   h.GetInt32Ptr(1),
					hitAmount:      1,
					shatterRate:    30,
					basedOnUserAtk: true,
					darkable:       true,
					silenceable:    false,
					reflectable:    false,
					accuracy: expAccuracy{
						accSource:   "accuracy",
						hitChance:   nil,
						accModifier: h.GetFloat32Ptr(2.5),
					},
					damage: &expDamage{
						damageCalc: []expAbilityDamage{
							{
								attackType:     1,
								targetStat:     1,
								damageType:     1,
								damageFormula:  1,
								damageConstant: 14,
							},
						},
						critical:        h.GetStrPtr("crit+weapon%"),
						criticalPlusVal: nil,
						isPiercing:      false,
						breakDmgLmt:     h.GetStrPtr("auto-ability"),
						element:         nil,
					},
					inflictedDelay: &expInflictedDelay{
						ctbAttackType:  "attack",
						delayType:      "tick-speed-based",
						damageConstant: 48,
						delayStrength:  "strong",
					},
					inflictedStatusConditions: []int32{},
					removedStatusConditions:   []int32{},
					copiedStatusConditions:    []int32{},
					statChanges:               []expStatChange{},
					modifierChanges:           []expModChange{},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/player-abilities/24",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"battle interactions": 0,
				},
			},
			expNameVer:         newExpNameVer(24, "use", 0),
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
			battleInteractions: []expBattleInteraction{},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/player-abilities/27",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"monsters":                           0,
					"battle interactions":                1,
					"battle interactions 0 stat changes": 1,
					"battle interactions 0 mod changes":  1,
				},
			},
			expNameVer:       newExpNameVer(27, "cheer", 0),
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
			battleInteractions: []expBattleInteraction{
				{
					target:         "all-allies",
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
							value:           1,
						},
					},
					modifierChanges: []expModChange{
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
			testGeneral: testGeneral{
				requestURL:     "/api/player-abilities/47",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"monsters":            7,
					"battle interactions": 1,
				},
			},
			expNameVer:       newExpNameVer(47, "curaga", 0),
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
			battleInteractions: []expBattleInteraction{
				{
					target:         "single-target",
					abilityRange:   h.GetInt32Ptr(4),
					hitAmount:      1,
					shatterRate:    0,
					basedOnUserAtk: false,
					darkable:       false,
					silenceable:    true,
					reflectable:    true,
					accuracy: expAccuracy{
						accSource:   "rate",
						hitChance:   h.GetInt32Ptr(255),
						accModifier: nil,
					},
					damage: &expDamage{
						damageCalc: []expAbilityDamage{
							{
								attackType:     2,
								targetStat:     1,
								damageType:     2,
								damageFormula:  7,
								damageConstant: 80,
							},
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
					statChanges:               []expStatChange{},
					modifierChanges:           []expModChange{},
				},
			},
		},
	}

	testSingleResources(t, tests, "GetPlayerAbility", testCfg.HandlePlayerAbilities, comparePlayerAbilities)
}

func TestGetMultiplePlayerAbilities(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/player-abilities/wakizashi",
				expectedStatus: http.StatusMultipleChoices,
			},
			count:   2,
			results: []int32{96, 97},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/player-abilities/auto-life",
				expectedStatus: http.StatusMultipleChoices,
			},
			count:   2,
			results: []int32{66, 67},
		},
	}

	testIdList(t, tests, testCfg.e.playerAbilities.endpoint, "GetMultiplePlayerAbilities", testCfg.HandlePlayerAbilities, compareAPIResourceLists[NamedApiResourceList])
}

func TestRetrievePlayerAbilities(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/player-abilities?limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   102,
			results: []int32{1, 22, 57, 78, 99, 102},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/player-abilities?category=2&help_bar=false",
				expectedStatus: http.StatusOK,
			},
			count:    4,
			previous: nil,
			next:     nil,
			results:  []int32{24, 40, 42, 43},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/player-abilities?related_stat=3&darkable=true",
				expectedStatus: http.StatusOK,
			},
			count:    1,
			previous: nil,
			next:     nil,
			results:  []int32{11},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/player-abilities?copycat=true&std_sg=3&user_atk=false",
				expectedStatus: http.StatusOK,
			},
			count:    3,
			previous: nil,
			next:     nil,
			results:  []int32{28, 84, 85},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/player-abilities?status_inflict=none&limit=max",
				expectedStatus: http.StatusOK,
			},
			count:    50,
			previous: nil,
			next:     nil,
			results:  []int32{20, 28, 36, 39, 53, 76, 86, 96, 102},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/player-abilities?silenceable=true&reflectable=false&user=3",
				expectedStatus: http.StatusOK,
			},
			count:    2,
			previous: nil,
			next:     nil,
			results:  []int32{82, 87},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/player-abilities?mp=20&exp_sg=2",
				expectedStatus: http.StatusOK,
			},
			count:    1,
			previous: nil,
			next:     nil,
			results:  []int32{47},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/player-abilities?mp_min=20&learn_item=59",
				expectedStatus: http.StatusOK,
			},
			count:    1,
			previous: nil,
			next:     nil,
			results:  []int32{64},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/player-abilities?mp_max=10&element=3",
				expectedStatus: http.StatusOK,
			},
			count:    2,
			previous: nil,
			next:     nil,
			results:  []int32{71, 75},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/player-abilities?element=fire,ice",
				expectedStatus: http.StatusOK,
			},
			count:    6,
			previous: nil,
			next:     nil,
			results:  []int32{69, 72, 73, 76, 77, 80},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/player-abilities?element=none&limit=max",
				expectedStatus: http.StatusOK,
			},
			count:    89,
			previous: nil,
			next:     nil,
			results:  []int32{1, 12, 24, 34, 43, 55, 67, 95, 102},
		},
	}

	testIdList(t, tests, testCfg.e.playerAbilities.endpoint, "RetrievePlayerAbilities", testCfg.HandlePlayerAbilities, compareAPIResourceLists[NamedApiResourceList])
}


func TestSubsectionPlayerAbilities(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/characters/2/default-abilities",
				expectedStatus: http.StatusOK,
				handler:        testCfg.HandleCharacters,
			},
			count:          2,
			parentResource: h.GetStrPtr("/characters/2"),
			results:        []int32{45, 53},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/characters/2/std-sg-abilities",
				expectedStatus: http.StatusOK,
				handler:        testCfg.HandleCharacters,
			},
			count:          18,
			parentResource: h.GetStrPtr("/characters/2"),
			results:        []int32{26, 42, 47, 53, 60, 62, 66},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/characters/2/exp-sg-abilities",
				expectedStatus: http.StatusOK,
				handler:        testCfg.HandleCharacters,
			},
			count:          19,
			parentResource: h.GetStrPtr("/characters/2"),
			results:        []int32{14, 45, 49, 54, 60, 64, 65},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/aeons/1/default-abilities",
				expectedStatus: http.StatusOK,
				handler:        testCfg.HandleAeons,
			},
			count:          5,
			parentResource: h.GetStrPtr("/aeons/1"),
			results:        []int32{69, 70, 71, 72, 88},
		},
	}

	testIdList(t, tests, testCfg.e.playerAbilities.endpoint, "SubsectionPlayerAbilities", nil, compareSimpleResourceLists[NamedAPIResource, PlayerAbilitySimple])
}
