package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
)

func TestSubsections(t *testing.T) {
	t.Parallel()
	tests := []TestGeneral{
		{
			RequestURL:     "/api/treasures/1/simple?availability=post",
			ExpectedStatus: http.StatusBadRequest,
			Handler:        testSrv.HandleTreasures,
			ExpectedErr:    "endpoint /treasures doesn't have any subsections",
		},
		{
			RequestURL:     "/api/sublocations/1/simple?limit=2",
			ExpectedStatus: http.StatusBadRequest,
			Handler:        testSrv.HandleSublocations,
			ExpectedErr:    "query parameters can't be used in combination with subsections",
		},
		{
			RequestURL:     "/api/sublocations/1/areas?limit=2",
			ExpectedStatus: http.StatusBadRequest,
			Handler:        testSrv.HandleSublocations,
			ExpectedErr:    "query parameters can't be used in combination with subsections",
		},
		{
			RequestURL:     "/api/sublocations/simple?ids=1-5&offset=3",
			ExpectedStatus: http.StatusBadRequest,
			Handler:        testSrv.HandleSublocations,
			ExpectedErr:    "parameter 'ids' can't be combined with other parameters",
		},
		{
			RequestURL:     "/api/locations/1/simple",
			ExpectedStatus: http.StatusOK,
			Handler:        testSrv.HandleLocations,
		},
		{
			RequestURL:     "/api/sublocations/1/simple",
			ExpectedStatus: http.StatusOK,
			Handler:        testSrv.HandleSublocations,
		},
		{
			RequestURL:     "/api/areas/1/simple",
			ExpectedStatus: http.StatusOK,
			Handler:        testSrv.HandleAreas,
		},
		{
			RequestURL:     "/api/monster-formations/1/simple",
			ExpectedStatus: http.StatusOK,
			Handler:        testSrv.HandleMonsterFormations,
		},
		{
			RequestURL:     "/api/shops/1/simple",
			ExpectedStatus: http.StatusOK,
			Handler:        testSrv.HandleShops,
		},
		{
			RequestURL:     "/api/monsters/1/simple",
			ExpectedStatus: http.StatusOK,
			Handler:        testSrv.HandleMonsters,
		},
		{
			RequestURL:     "/api/abilities/1/simple",
			ExpectedStatus: http.StatusOK,
			Handler:        testSrv.HandleAbilities,
		},
		{
			RequestURL:     "/api/player-abilities/1/simple",
			ExpectedStatus: http.StatusOK,
			Handler:        testSrv.HandlePlayerAbilities,
		},
		{
			RequestURL:     "/api/overdrive-abilities/1/simple",
			ExpectedStatus: http.StatusOK,
			Handler:        testSrv.HandleOverdriveAbilities,
		},
		{
			RequestURL:     "/api/item-abilities/1/simple",
			ExpectedStatus: http.StatusOK,
			Handler:        testSrv.HandleItemAbilities,
		},
		{
			RequestURL:     "/api/trigger-commands/1/simple",
			ExpectedStatus: http.StatusOK,
			Handler:        testSrv.HandleTriggerCommands,
		},
		{
			RequestURL:     "/api/misc-abilities/1/simple",
			ExpectedStatus: http.StatusOK,
			Handler:        testSrv.HandleMiscAbilities,
		},
		{
			RequestURL:     "/api/enemy-abilities/1/simple",
			ExpectedStatus: http.StatusOK,
			Handler:        testSrv.HandleEnemyAbilities,
		},
		{
			RequestURL:     "/api/overdrives/1/simple",
			ExpectedStatus: http.StatusOK,
			Handler:        testSrv.HandleOverdrives,
		},
		{
			RequestURL:     "/api/mixes/1/simple",
			ExpectedStatus: http.StatusOK,
			Handler:        testSrv.HandleMixes,
		},
		{
			RequestURL:     "/api/auto-abilities/1/simple",
			ExpectedStatus: http.StatusOK,
			Handler:        testSrv.HandleAutoAbilities,
		},
	}

	testStatusses(t, tests, "Subsections", nil)
}

func TestSections(t *testing.T) {
	t.Parallel()
	tests := []ListNames{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas/sections",
				ExpectedStatus: http.StatusOK,
				Endpoint:       apilookups.E.Areas.Endpoint,
				Handler:        testSrv.HandleAreas,
			},
			Count: 6,
			Results: ts.SnsToNamedParams([]ts.SectionName{
				ts.SnSimple,
				ts.SnConnected,
				ts.SnMonsters,
				ts.SnMonsterFormations,
				ts.SnSongs,
				ts.SnTreasures,
			}),
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sublocations/sections",
				ExpectedStatus: http.StatusOK,
				Endpoint:       apilookups.E.Sublocations.Endpoint,
				Handler:        testSrv.HandleSublocations,
			},
			Count: 8,
			Results: ts.SnsToNamedParams([]ts.SectionName{
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
			TestGeneral: TestGeneral{
				RequestURL:     "/api/locations/sections",
				ExpectedStatus: http.StatusOK,
				Endpoint:       apilookups.E.Locations.Endpoint,
				Handler:        testSrv.HandleLocations,
			},
			Count: 9,
			Results: ts.SnsToNamedParams([]ts.SectionName{
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
			TestGeneral: TestGeneral{
				RequestURL:     "/api/aeons/sections",
				ExpectedStatus: http.StatusOK,
				Endpoint:       apilookups.E.Aeons.Endpoint,
				Handler:        testSrv.HandleAeons,
			},
			Count: 4,
			Results: ts.SnsToNamedParams([]ts.SectionName{
				ts.SnDefaultAbilities,
				ts.SnOverdriveAbilities,
				ts.SnOverdrives,
				ts.SnStats,
			}),
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/characters/sections",
				ExpectedStatus: http.StatusOK,
				Endpoint:       apilookups.E.Characters.Endpoint,
				Handler:        testSrv.HandleCharacters,
			},
			Count: 5,
			Results: ts.SnsToNamedParams([]ts.SectionName{
				ts.SnDefaultAbilities,
				ts.SnStdSgAbilities,
				ts.SnExpSgAbilities,
				ts.SnOverdriveAbilities,
				ts.SnOverdrives,
			}),
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/character-classes/sections",
				ExpectedStatus: http.StatusOK,
				Endpoint:       apilookups.E.CharacterClasses.Endpoint,
				Handler:        testSrv.HandleCharacterClasses,
			},
			Count: 4,
			Results: ts.SnsToNamedParams([]ts.SectionName{
				ts.SnDefaultAbilities,
				ts.SnLearnableOverdrives,
				ts.SnDefaultOverdrives,
				ts.SnLearnableOverdrives,
			}),
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/arena-creations/sections",
				ExpectedStatus: http.StatusOK,
				Endpoint:       apilookups.E.ArenaCreations.Endpoint,
				Handler:        testSrv.HandleArenaCreations,
			},
			Count:   0,
			Results: []ts.NamedParam{},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/blitzball-prizes/sections",
				ExpectedStatus: http.StatusOK,
				Endpoint:       apilookups.E.BlitzballPrizes.Endpoint,
				Handler:        testSrv.HandleBlitzballPrizes,
			},
			Count:   0,
			Results: []ts.NamedParam{},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/fmvs/sections",
				ExpectedStatus: http.StatusOK,
				Endpoint:       apilookups.E.FMVs.Endpoint,
				Handler:        testSrv.HandleFMVs,
			},
			Count:   0,
			Results: []ts.NamedParam{},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/monsters/sections",
				ExpectedStatus: http.StatusOK,
				Endpoint:       apilookups.E.Monsters.Endpoint,
				Handler:        testSrv.HandleMonsters,
			},
			Count: 4,
			Results: ts.SnsToNamedParams([]ts.SectionName{
				ts.SnSimple,
				ts.SnAbilities,
				ts.SnAreas,
				ts.SnMonsterFormations,
			}),
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/monster-formations/sections",
				ExpectedStatus: http.StatusOK,
				Endpoint:       apilookups.E.MonsterFormations.Endpoint,
				Handler:        testSrv.HandleMonsterFormations,
			},
			Count: 2,
			Results: ts.SnsToNamedParams([]ts.SectionName{
				ts.SnSimple,
				ts.SnMonsters,
			}),
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/overdrive-modes/sections",
				ExpectedStatus: http.StatusOK,
				Endpoint:       apilookups.E.OverdriveModes.Endpoint,
				Handler:        testSrv.HandleOverdriveModes,
			},
			Count:   0,
			Results: []ts.NamedParam{},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/shops/sections",
				ExpectedStatus: http.StatusOK,
				Endpoint:       apilookups.E.Shops.Endpoint,
				Handler:        testSrv.HandleShops,
			},
			Count: 1,
			Results: ts.SnsToNamedParams([]ts.SectionName{
				ts.SnSimple,
			}),
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/songs/sections",
				ExpectedStatus: http.StatusOK,
				Endpoint:       apilookups.E.Songs.Endpoint,
				Handler:        testSrv.HandleSongs,
			},
			Count:   0,
			Results: []ts.NamedParam{},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sidequests/sections",
				ExpectedStatus: http.StatusOK,
				Endpoint:       apilookups.E.Sidequests.Endpoint,
				Handler:        testSrv.HandleSidequests,
			},
			Count: 1,
			Results: ts.SnsToNamedParams([]ts.SectionName{
				ts.SnSubquests,
			}),
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/subquests/sections",
				ExpectedStatus: http.StatusOK,
				Endpoint:       apilookups.E.Subquests.Endpoint,
				Handler:        testSrv.HandleSubquests,
			},
			Count:   0,
			Results: []ts.NamedParam{},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/treasures/sections",
				ExpectedStatus: http.StatusOK,
				Endpoint:       apilookups.E.Treasures.Endpoint,
				Handler:        testSrv.HandleTreasures,
			},
			Count:   0,
			Results: []ts.NamedParam{},
		},
	}

	testNameList(t, tests, "Sections", compareSectionLists)
}
