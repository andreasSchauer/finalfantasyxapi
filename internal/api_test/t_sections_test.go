package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func TestSubsections(t *testing.T) {
	t.Parallel()
	tests := []testGeneral{
		{
			requestURL:     "/api/treasures/1/simple?availability=post",
			expectedStatus: http.StatusBadRequest,
			handler:        testSrv.Cfg.HandleTreasures,
			expectedErr:    "endpoint /treasures doesn't have any subsections",
		},
		{
			requestURL:     "/api/sublocations/1/simple?limit=2",
			expectedStatus: http.StatusBadRequest,
			handler:        testSrv.Cfg.HandleSublocations,
			expectedErr:    "query parameters can't be used in combination with subsections",
		},
		{
			requestURL:     "/api/sublocations/1/areas?limit=2",
			expectedStatus: http.StatusBadRequest,
			handler:        testSrv.Cfg.HandleSublocations,
			expectedErr:    "query parameters can't be used in combination with subsections",
		},
		{
			requestURL:     "/api/sublocations/simple?ids=1-5&offset=3",
			expectedStatus: http.StatusBadRequest,
			handler:        testSrv.Cfg.HandleSublocations,
			expectedErr:    "parameter 'ids' can't be combined with other parameters",
		},
		{
			requestURL:     "/api/locations/1/simple",
			expectedStatus: http.StatusOK,
			handler:        testSrv.Cfg.HandleLocations,
		},
		{
			requestURL:     "/api/sublocations/1/simple",
			expectedStatus: http.StatusOK,
			handler:        testSrv.Cfg.HandleSublocations,
		},
		{
			requestURL:     "/api/areas/1/simple",
			expectedStatus: http.StatusOK,
			handler:        testSrv.Cfg.HandleAreas,
		},
		{
			requestURL:     "/api/monster-formations/1/simple",
			expectedStatus: http.StatusOK,
			handler:        testSrv.Cfg.HandleMonsterFormations,
		},
		{
			requestURL:     "/api/shops/1/simple",
			expectedStatus: http.StatusOK,
			handler:        testSrv.Cfg.HandleShops,
		},
		{
			requestURL:     "/api/monsters/1/simple",
			expectedStatus: http.StatusOK,
			handler:        testSrv.Cfg.HandleMonsters,
		},
		{
			requestURL:     "/api/abilities/1/simple",
			expectedStatus: http.StatusOK,
			handler:        testSrv.Cfg.HandleAbilities,
		},
		{
			requestURL:     "/api/player-abilities/1/simple",
			expectedStatus: http.StatusOK,
			handler:        testSrv.Cfg.HandlePlayerAbilities,
		},
		{
			requestURL:     "/api/overdrive-abilities/1/simple",
			expectedStatus: http.StatusOK,
			handler:        testSrv.Cfg.HandleOverdriveAbilities,
		},
		{
			requestURL:     "/api/item-abilities/1/simple",
			expectedStatus: http.StatusOK,
			handler:        testSrv.Cfg.HandleItemAbilities,
		},
		{
			requestURL:     "/api/trigger-commands/1/simple",
			expectedStatus: http.StatusOK,
			handler:        testSrv.Cfg.HandleTriggerCommands,
		},
		{
			requestURL:     "/api/misc-abilities/1/simple",
			expectedStatus: http.StatusOK,
			handler:        testSrv.Cfg.HandleMiscAbilities,
		},
		{
			requestURL:     "/api/enemy-abilities/1/simple",
			expectedStatus: http.StatusOK,
			handler:        testSrv.Cfg.HandleEnemyAbilities,
		},
		{
			requestURL:     "/api/overdrives/1/simple",
			expectedStatus: http.StatusOK,
			handler:        testSrv.Cfg.HandleOverdrives,
		},
		{
			requestURL:     "/api/mixes/1/simple",
			expectedStatus: http.StatusOK,
			handler:        testSrv.Cfg.HandleMixes,
		},
		{
			requestURL:     "/api/auto-abilities/1/simple",
			expectedStatus: http.StatusOK,
			handler:        testSrv.Cfg.HandleAutoAbilities,
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
				endpoint:       api.E.Areas.Endpoint,
				handler:        testSrv.Cfg.HandleAreas,
			},
			count: 6,
			results: ts.SnsToNamedParams([]ts.SectionName{
				ts.SnSimple,
				ts.SnConnected,
				ts.SnMonsters,
				ts.SnMonsterFormations,
				ts.SnSongs,
				ts.SnTreasures,
			}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations/sections",
				expectedStatus: http.StatusOK,
				endpoint:       api.E.Sublocations.Endpoint,
				handler:        testSrv.Cfg.HandleSublocations,
			},
			count: 8,
			results: ts.SnsToNamedParams([]ts.SectionName{
				ts.SnSimple,
				ts.SnAreas,
				ts.SnConnected,
				ts.SnMonsters,
				ts.SnMonsterFormations,
				ts.SnShops,
				ts.SnSongs,
				ts.SnTreasures,
			}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations/sections",
				expectedStatus: http.StatusOK,
				endpoint:       api.E.Locations.Endpoint,
				handler:        testSrv.Cfg.HandleLocations,
			},
			count: 9,
			results: ts.SnsToNamedParams([]ts.SectionName{
				ts.SnSimple,
				ts.SnAreas,
				ts.SnConnected,
				ts.SnMonsters,
				ts.SnMonsterFormations,
				ts.SnShops,
				ts.SnSongs,
				ts.SnSublocations,
				ts.SnTreasures,
			}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/aeons/sections",
				expectedStatus: http.StatusOK,
				endpoint:       api.E.Aeons.Endpoint,
				handler:        testSrv.Cfg.HandleAeons,
			},
			count: 4,
			results: ts.SnsToNamedParams([]ts.SectionName{
				ts.SnDefaultAbilities,
				ts.SnOverdriveAbilities,
				ts.SnOverdrives,
				ts.SnStats,
			}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/characters/sections",
				expectedStatus: http.StatusOK,
				endpoint:       api.E.Characters.Endpoint,
				handler:        testSrv.Cfg.HandleCharacters,
			},
			count: 5,
			results: ts.SnsToNamedParams([]ts.SectionName{
				ts.SnDefaultAbilities,
				ts.SnStdSgAbilities,
				ts.SnExpSgAbilities,
				ts.SnOverdriveAbilities,
				ts.SnOverdrives,
			}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/character-classes/sections",
				expectedStatus: http.StatusOK,
				endpoint:       api.E.CharacterClasses.Endpoint,
				handler:        testSrv.Cfg.HandleCharacterClasses,
			},
			count: 4,
			results: ts.SnsToNamedParams([]ts.SectionName{
				ts.SnDefaultAbilities,
				ts.SnLearnableOverdrives,
				ts.SnDefaultOverdrives,
				ts.SnLearnableOverdrives,
			}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/arena-creations/sections",
				expectedStatus: http.StatusOK,
				endpoint:       api.E.ArenaCreations.Endpoint,
				handler:        testSrv.Cfg.HandleArenaCreations,
			},
			count:   0,
			results: []ts.NamedParam{},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/blitzball-prizes/sections",
				expectedStatus: http.StatusOK,
				endpoint:       api.E.BlitzballPrizes.Endpoint,
				handler:        testSrv.Cfg.HandleBlitzballPrizes,
			},
			count:   0,
			results: []ts.NamedParam{},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/fmvs/sections",
				expectedStatus: http.StatusOK,
				endpoint:       api.E.FMVs.Endpoint,
				handler:        testSrv.Cfg.HandleFMVs,
			},
			count:   0,
			results: []ts.NamedParam{},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/sections",
				expectedStatus: http.StatusOK,
				endpoint:       api.E.Monsters.Endpoint,
				handler:        testSrv.Cfg.HandleMonsters,
			},
			count: 4,
			results: ts.SnsToNamedParams([]ts.SectionName{
				ts.SnSimple,
				ts.SnAbilities,
				ts.SnAreas,
				ts.SnMonsterFormations,
			}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monster-formations/sections",
				expectedStatus: http.StatusOK,
				endpoint:       api.E.MonsterFormations.Endpoint,
				handler:        testSrv.Cfg.HandleMonsterFormations,
			},
			count: 2,
			results: ts.SnsToNamedParams([]ts.SectionName{
				ts.SnSimple,
				ts.SnMonsters,
			}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-modes/sections",
				expectedStatus: http.StatusOK,
				endpoint:       api.E.OverdriveModes.Endpoint,
				handler:        testSrv.Cfg.HandleOverdriveModes,
			},
			count:   0,
			results: []ts.NamedParam{},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/shops/sections",
				expectedStatus: http.StatusOK,
				endpoint:       api.E.Shops.Endpoint,
				handler:        testSrv.Cfg.HandleShops,
			},
			count: 1,
			results: ts.SnsToNamedParams([]ts.SectionName{
				ts.SnSimple,
			}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/songs/sections",
				expectedStatus: http.StatusOK,
				endpoint:       api.E.Songs.Endpoint,
				handler:        testSrv.Cfg.HandleSongs,
			},
			count:   0,
			results: []ts.NamedParam{},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sidequests/sections",
				expectedStatus: http.StatusOK,
				endpoint:       api.E.Sidequests.Endpoint,
				handler:        testSrv.Cfg.HandleSidequests,
			},
			count: 1,
			results: ts.SnsToNamedParams([]ts.SectionName{
				ts.SnSubquests,
			}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/subquests/sections",
				expectedStatus: http.StatusOK,
				endpoint:       api.E.Subquests.Endpoint,
				handler:        testSrv.Cfg.HandleSubquests,
			},
			count:   0,
			results: []ts.NamedParam{},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/treasures/sections",
				expectedStatus: http.StatusOK,
				endpoint:       api.E.Treasures.Endpoint,
				handler:        testSrv.Cfg.HandleTreasures,
			},
			count:   0,
			results: []ts.NamedParam{},
		},
	}

	testNameList(t, tests, "Sections", compareSectionLists)
}
