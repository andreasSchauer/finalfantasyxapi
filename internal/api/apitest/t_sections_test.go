package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
)

func TestSubsections(t *testing.T) {
	t.Parallel()
	tests := []testGeneral{
		{
			requestURL:     "/api/treasures/1/simple?availability=post",
			expectedStatus: http.StatusBadRequest,
			handler:        testSrv.HandleTreasures,
			expectedErr:    "endpoint /treasures doesn't have any subsections",
		},
		{
			requestURL:     "/api/sublocations/1/simple?limit=2",
			expectedStatus: http.StatusBadRequest,
			handler:        testSrv.HandleSublocations,
			expectedErr:    "query parameters can't be used in combination with subsections",
		},
		{
			requestURL:     "/api/sublocations/1/areas?limit=2",
			expectedStatus: http.StatusBadRequest,
			handler:        testSrv.HandleSublocations,
			expectedErr:    "query parameters can't be used in combination with subsections",
		},
		{
			requestURL:     "/api/sublocations/simple?ids=1-5&offset=3",
			expectedStatus: http.StatusBadRequest,
			handler:        testSrv.HandleSublocations,
			expectedErr:    "parameter 'ids' can't be combined with other parameters",
		},
		{
			requestURL:     "/api/locations/1/simple",
			expectedStatus: http.StatusOK,
			handler:        testSrv.HandleLocations,
		},
		{
			requestURL:     "/api/sublocations/1/simple",
			expectedStatus: http.StatusOK,
			handler:        testSrv.HandleSublocations,
		},
		{
			requestURL:     "/api/areas/1/simple",
			expectedStatus: http.StatusOK,
			handler:        testSrv.HandleAreas,
		},
		{
			requestURL:     "/api/monster-formations/1/simple",
			expectedStatus: http.StatusOK,
			handler:        testSrv.HandleMonsterFormations,
		},
		{
			requestURL:     "/api/shops/1/simple",
			expectedStatus: http.StatusOK,
			handler:        testSrv.HandleShops,
		},
		{
			requestURL:     "/api/monsters/1/simple",
			expectedStatus: http.StatusOK,
			handler:        testSrv.HandleMonsters,
		},
		{
			requestURL:     "/api/abilities/1/simple",
			expectedStatus: http.StatusOK,
			handler:        testSrv.HandleAbilities,
		},
		{
			requestURL:     "/api/player-abilities/1/simple",
			expectedStatus: http.StatusOK,
			handler:        testSrv.HandlePlayerAbilities,
		},
		{
			requestURL:     "/api/overdrive-abilities/1/simple",
			expectedStatus: http.StatusOK,
			handler:        testSrv.HandleOverdriveAbilities,
		},
		{
			requestURL:     "/api/item-abilities/1/simple",
			expectedStatus: http.StatusOK,
			handler:        testSrv.HandleItemAbilities,
		},
		{
			requestURL:     "/api/trigger-commands/1/simple",
			expectedStatus: http.StatusOK,
			handler:        testSrv.HandleTriggerCommands,
		},
		{
			requestURL:     "/api/misc-abilities/1/simple",
			expectedStatus: http.StatusOK,
			handler:        testSrv.HandleMiscAbilities,
		},
		{
			requestURL:     "/api/enemy-abilities/1/simple",
			expectedStatus: http.StatusOK,
			handler:        testSrv.HandleEnemyAbilities,
		},
		{
			requestURL:     "/api/overdrives/1/simple",
			expectedStatus: http.StatusOK,
			handler:        testSrv.HandleOverdrives,
		},
		{
			requestURL:     "/api/mixes/1/simple",
			expectedStatus: http.StatusOK,
			handler:        testSrv.HandleMixes,
		},
		{
			requestURL:     "/api/auto-abilities/1/simple",
			expectedStatus: http.StatusOK,
			handler:        testSrv.HandleAutoAbilities,
		},
	}

	testStatusses(t, tests, "Subsections", nil)
}

func TestSections(t *testing.T) {
	t.Parallel()
	tests := []ListNames{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/sections",
				expectedStatus: http.StatusOK,
				endpoint:       apilookups.E.Areas.Endpoint,
				handler:        testSrv.HandleAreas,
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
				endpoint:       apilookups.E.Sublocations.Endpoint,
				handler:        testSrv.HandleSublocations,
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
				endpoint:       apilookups.E.Locations.Endpoint,
				handler:        testSrv.HandleLocations,
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
				endpoint:       apilookups.E.Aeons.Endpoint,
				handler:        testSrv.HandleAeons,
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
				endpoint:       apilookups.E.Characters.Endpoint,
				handler:        testSrv.HandleCharacters,
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
				endpoint:       apilookups.E.CharacterClasses.Endpoint,
				handler:        testSrv.HandleCharacterClasses,
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
				endpoint:       apilookups.E.ArenaCreations.Endpoint,
				handler:        testSrv.HandleArenaCreations,
			},
			count:   0,
			results: []ts.NamedParam{},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/blitzball-prizes/sections",
				expectedStatus: http.StatusOK,
				endpoint:       apilookups.E.BlitzballPrizes.Endpoint,
				handler:        testSrv.HandleBlitzballPrizes,
			},
			count:   0,
			results: []ts.NamedParam{},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/fmvs/sections",
				expectedStatus: http.StatusOK,
				endpoint:       apilookups.E.FMVs.Endpoint,
				handler:        testSrv.HandleFMVs,
			},
			count:   0,
			results: []ts.NamedParam{},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/sections",
				expectedStatus: http.StatusOK,
				endpoint:       apilookups.E.Monsters.Endpoint,
				handler:        testSrv.HandleMonsters,
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
				endpoint:       apilookups.E.MonsterFormations.Endpoint,
				handler:        testSrv.HandleMonsterFormations,
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
				endpoint:       apilookups.E.OverdriveModes.Endpoint,
				handler:        testSrv.HandleOverdriveModes,
			},
			count:   0,
			results: []ts.NamedParam{},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/shops/sections",
				expectedStatus: http.StatusOK,
				endpoint:       apilookups.E.Shops.Endpoint,
				handler:        testSrv.HandleShops,
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
				endpoint:       apilookups.E.Songs.Endpoint,
				handler:        testSrv.HandleSongs,
			},
			count:   0,
			results: []ts.NamedParam{},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sidequests/sections",
				expectedStatus: http.StatusOK,
				endpoint:       apilookups.E.Sidequests.Endpoint,
				handler:        testSrv.HandleSidequests,
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
				endpoint:       apilookups.E.Subquests.Endpoint,
				handler:        testSrv.HandleSubquests,
			},
			count:   0,
			results: []ts.NamedParam{},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/treasures/sections",
				expectedStatus: http.StatusOK,
				endpoint:       apilookups.E.Treasures.Endpoint,
				handler:        testSrv.HandleTreasures,
			},
			count:   0,
			results: []ts.NamedParam{},
		},
	}

	testNameList(t, tests, "Sections", compareSectionLists)
}
