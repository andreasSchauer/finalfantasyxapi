package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func TestSubsections(t *testing.T) {
	t.Parallel()
	tests := []testGeneral{
		{
			requestURL:     "/api/treasures/1/simple?availability=post",
			expectedStatus: http.StatusBadRequest,
			handler:        testCfg.HandleTreasures,
			expectedErr:    "endpoint /treasures doesn't have any subsections",
		},
		{
			requestURL:     "/api/sublocations/1/simple?limit=2",
			expectedStatus: http.StatusBadRequest,
			handler:        testCfg.HandleSublocations,
			expectedErr:    "query parameters can't be used in combination with subsections",
		},
		{
			requestURL:     "/api/sublocations/1/areas?limit=2",
			expectedStatus: http.StatusBadRequest,
			handler:        testCfg.HandleSublocations,
			expectedErr:    "query parameters can't be used in combination with subsections",
		},
		{
			requestURL:     "/api/sublocations/simple?ids=1-5&offset=3",
			expectedStatus: http.StatusBadRequest,
			handler:        testCfg.HandleSublocations,
			expectedErr:    "parameter 'ids' can't be combined with other parameters",
		},
		{
			requestURL:     "/api/locations/1/simple",
			expectedStatus: http.StatusOK,
			handler:        testCfg.HandleLocations,
		},
		{
			requestURL:     "/api/sublocations/1/simple",
			expectedStatus: http.StatusOK,
			handler:        testCfg.HandleSublocations,
		},
		{
			requestURL:     "/api/areas/1/simple",
			expectedStatus: http.StatusOK,
			handler:        testCfg.HandleAreas,
		},
		{
			requestURL:     "/api/monster-formations/1/simple",
			expectedStatus: http.StatusOK,
			handler:        testCfg.HandleMonsterFormations,
		},
		{
			requestURL:     "/api/shops/1/simple",
			expectedStatus: http.StatusOK,
			handler:        testCfg.HandleShops,
		},
		{
			requestURL:     "/api/monsters/1/simple",
			expectedStatus: http.StatusOK,
			handler:        testCfg.HandleMonsters,
		},
		{
			requestURL:     "/api/abilities/1/simple",
			expectedStatus: http.StatusOK,
			handler:        testCfg.HandleAbilities,
		},
		{
			requestURL:     "/api/player-abilities/1/simple",
			expectedStatus: http.StatusOK,
			handler:        testCfg.HandlePlayerAbilities,
		},
		{
			requestURL:     "/api/overdrive-abilities/1/simple",
			expectedStatus: http.StatusOK,
			handler:        testCfg.HandleOverdriveAbilities,
		},
		{
			requestURL:     "/api/item-abilities/1/simple",
			expectedStatus: http.StatusOK,
			handler:        testCfg.HandleItemAbilities,
		},
		{
			requestURL:     "/api/trigger-commands/1/simple",
			expectedStatus: http.StatusOK,
			handler:        testCfg.HandleTriggerCommands,
		},
		{
			requestURL:     "/api/misc-abilities/1/simple",
			expectedStatus: http.StatusOK,
			handler:        testCfg.HandleMiscAbilities,
		},
		{
			requestURL:     "/api/enemy-abilities/1/simple",
			expectedStatus: http.StatusOK,
			handler:        testCfg.HandleEnemyAbilities,
		},
		{
			requestURL:     "/api/overdrives/1/simple",
			expectedStatus: http.StatusOK,
			handler:        testCfg.HandleOverdrives,
		},
		{
			requestURL:     "/api/mixes/1/simple",
			expectedStatus: http.StatusOK,
			handler:        testCfg.HandleMixes,
		},
		{
			requestURL:     "/api/auto-abilities/1/simple",
			expectedStatus: http.StatusOK,
			handler:        testCfg.HandleAutoAbilities,
		},
	}

	testStatusses(t, tests, "Subsections", nil)
}

func TestSections(t *testing.T) {
	t.Parallel()
	tests := []expListNames{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.Areas.Endpoint,
				handler:        testCfg.HandleAreas,
			},
			count: 6,
			results: typestrings.SnsToNamedParams([]typestrings.SectionName{
				typestrings.SnSimple,
				typestrings.SnConnected,
				typestrings.SnMonsters,
				typestrings.SnMonsterFormations,
				typestrings.SnSongs,
				typestrings.SnTreasures,
			}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.Sublocations.Endpoint,
				handler:        testCfg.HandleSublocations,
			},
			count: 8,
			results: typestrings.SnsToNamedParams([]typestrings.SectionName{
				typestrings.SnSimple,
				typestrings.SnAreas,
				typestrings.SnConnected,
				typestrings.SnMonsters,
				typestrings.SnMonsterFormations,
				typestrings.SnShops,
				typestrings.SnSongs,
				typestrings.SnTreasures,
			}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.Locations.Endpoint,
				handler:        testCfg.HandleLocations,
			},
			count: 9,
			results: typestrings.SnsToNamedParams([]typestrings.SectionName{
				typestrings.SnSimple,
				typestrings.SnAreas,
				typestrings.SnConnected,
				typestrings.SnMonsters,
				typestrings.SnMonsterFormations,
				typestrings.SnShops,
				typestrings.SnSongs,
				typestrings.SnSublocations,
				typestrings.SnTreasures,
			}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/aeons/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.Aeons.Endpoint,
				handler:        testCfg.HandleAeons,
			},
			count: 4,
			results: typestrings.SnsToNamedParams([]typestrings.SectionName{
				typestrings.SnDefaultAbilities,
				typestrings.SnOverdriveAbilities,
				typestrings.SnOverdrives,
				typestrings.SnStats,
			}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/characters/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.Characters.Endpoint,
				handler:        testCfg.HandleCharacters,
			},
			count: 5,
			results: typestrings.SnsToNamedParams([]typestrings.SectionName{
				typestrings.SnDefaultAbilities,
				typestrings.SnStdSgAbilities,
				typestrings.SnExpSgAbilities,
				typestrings.SnOverdriveAbilities,
				typestrings.SnOverdrives,
			}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/character-classes/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.CharacterClasses.Endpoint,
				handler:        testCfg.HandleCharacterClasses,
			},
			count: 4,
			results: typestrings.SnsToNamedParams([]typestrings.SectionName{
				typestrings.SnDefaultAbilities,
				typestrings.SnLearnableOverdrives,
				typestrings.SnDefaultOverdrives,
				typestrings.SnLearnableOverdrives,
			}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/arena-creations/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.ArenaCreations.Endpoint,
				handler:        testCfg.HandleArenaCreations,
			},
			count:   0,
			results: []typestrings.NamedParam{},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/blitzball-prizes/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.BlitzballPrizes.Endpoint,
				handler:        testCfg.HandleBlitzballPrizes,
			},
			count:   0,
			results: []typestrings.NamedParam{},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/fmvs/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.FMVs.Endpoint,
				handler:        testCfg.HandleFMVs,
			},
			count:   0,
			results: []typestrings.NamedParam{},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.Monsters.Endpoint,
				handler:        testCfg.HandleMonsters,
			},
			count: 4,
			results: typestrings.SnsToNamedParams([]typestrings.SectionName{
				typestrings.SnSimple,
				typestrings.SnAbilities,
				typestrings.SnAreas,
				typestrings.SnMonsterFormations,
			}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monster-formations/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.MonsterFormations.Endpoint,
				handler:        testCfg.HandleMonsterFormations,
			},
			count: 2,
			results: typestrings.SnsToNamedParams([]typestrings.SectionName{
				typestrings.SnSimple,
				typestrings.SnMonsters,
			}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-modes/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.OverdriveModes.Endpoint,
				handler:        testCfg.HandleOverdriveModes,
			},
			count:   0,
			results: []typestrings.NamedParam{},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/shops/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.Shops.Endpoint,
				handler:        testCfg.HandleShops,
			},
			count: 1,
			results: typestrings.SnsToNamedParams([]typestrings.SectionName{
				typestrings.SnSimple,
			}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/songs/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.Songs.Endpoint,
				handler:        testCfg.HandleSongs,
			},
			count:   0,
			results: []typestrings.NamedParam{},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sidequests/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.Sidequests.Endpoint,
				handler:        testCfg.HandleSidequests,
			},
			count: 1,
			results: typestrings.SnsToNamedParams([]typestrings.SectionName{
				typestrings.SnSubquests,
			}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/subquests/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.Subquests.Endpoint,
				handler:        testCfg.HandleSubquests,
			},
			count:   0,
			results: []typestrings.NamedParam{},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/treasures/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.Treasures.Endpoint,
				handler:        testCfg.HandleTreasures,
			},
			count:   0,
			results: []typestrings.NamedParam{},
		},
	}

	testNameList(t, tests, "Sections", compareSectionLists)
}
