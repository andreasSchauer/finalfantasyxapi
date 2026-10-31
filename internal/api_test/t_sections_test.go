package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api"
)

func TestSubsections(t *testing.T) {
	t.Parallel()
	tests := []testGeneral{
		{
			requestURL:     "/api/treasures/1/simple?availability=post",
			expectedStatus: http.StatusBadRequest,
			handler:        testCfg.HandleTreasures,
			expectedErr:    "endpoint /treasures doesn't have any subsections.",
		},
		{
			requestURL:     "/api/sublocations/1/simple?limit=2",
			expectedStatus: http.StatusBadRequest,
			handler:        testCfg.HandleSublocations,
			expectedErr:    "query parameters can't be used in combination with subsections.",
		},
		{
			requestURL:     "/api/sublocations/1/areas?limit=2",
			expectedStatus: http.StatusBadRequest,
			handler:        testCfg.HandleSublocations,
			expectedErr:    "query parameters can't be used in combination with subsections.",
		},
		{
			requestURL:     "/api/sublocations/simple?ids=1-5&offset=3",
			expectedStatus: http.StatusBadRequest,
			handler:        testCfg.HandleSublocations,
			expectedErr:    "parameter 'ids' can't be combined with other parameters.",
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
			results: api.SnsToNamedParams([]api.SectionName{
				api.SnSimple,
				api.SnConnected,
				api.SnMonsters,
				api.SnMonsterFormations,
				api.SnSongs,
				api.SnTreasures,
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
			results: api.SnsToNamedParams([]api.SectionName{
				api.SnSimple,
				api.SnAreas,
				api.SnConnected,
				api.SnMonsters,
				api.SnMonsterFormations,
				api.SnShops,
				api.SnSongs,
				api.SnTreasures,
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
			results: api.SnsToNamedParams([]api.SectionName{
				api.SnSimple,
				api.SnAreas,
				api.SnConnected,
				api.SnMonsters,
				api.SnMonsterFormations,
				api.SnShops,
				api.SnSongs,
				api.SnSublocations,
				api.SnTreasures,
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
			results: api.SnsToNamedParams([]api.SectionName{
				api.SnDefaultAbilities,
				api.SnOverdriveAbilities,
				api.SnOverdrives,
				api.SnStats,
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
			results: api.SnsToNamedParams([]api.SectionName{
				api.SnDefaultAbilities,
				api.SnStdSgAbilities,
				api.SnExpSgAbilities,
				api.SnOverdriveAbilities,
				api.SnOverdrives,
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
			results: api.SnsToNamedParams([]api.SectionName{
				api.SnDefaultAbilities,
				api.SnLearnableOverdrives,
				api.SnDefaultOverdrives,
				api.SnLearnableOverdrives,
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
			results: []api.NamedParam{},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/blitzball-prizes/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.BlitzballPrizes.Endpoint,
				handler:        testCfg.HandleBlitzballPrizes,
			},
			count:   0,
			results: []api.NamedParam{},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/fmvs/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.FMVs.Endpoint,
				handler:        testCfg.HandleFMVs,
			},
			count:   0,
			results: []api.NamedParam{},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.Monsters.Endpoint,
				handler:        testCfg.HandleMonsters,
			},
			count: 4,
			results: api.SnsToNamedParams([]api.SectionName{
				api.SnSimple,
				api.SnAbilities,
				api.SnAreas,
				api.SnMonsterFormations,
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
			results: api.SnsToNamedParams([]api.SectionName{
				api.SnSimple,
				api.SnMonsters,
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
			results: []api.NamedParam{},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/shops/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.Shops.Endpoint,
				handler:        testCfg.HandleShops,
			},
			count: 1,
			results: api.SnsToNamedParams([]api.SectionName{
				api.SnSimple,
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
			results: []api.NamedParam{},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sidequests/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.Sidequests.Endpoint,
				handler:        testCfg.HandleSidequests,
			},
			count: 1,
			results: api.SnsToNamedParams([]api.SectionName{
				api.SnSubquests,
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
			results: []api.NamedParam{},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/treasures/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.Treasures.Endpoint,
				handler:        testCfg.HandleTreasures,
			},
			count:   0,
			results: []api.NamedParam{},
		},
	}

	testNameList(t, tests, "", "Sections", nil, compareSectionLists)
}
