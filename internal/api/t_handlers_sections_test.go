package api

import (
	"net/http"
	"testing"
)

func TestSections(t *testing.T) {
	t.Parallel()
	tests := []expListNames{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.areas.endpoint,
				handler:        testCfg.HandleAreas,
			},
			count: 6,
			results: snsToNamedParams([]SectionName{
				snSimple,
				snConnected,
				snMonsters,
				snMonsterFormations,
				snSongs,
				snTreasures,
			}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.sublocations.endpoint,
				handler:        testCfg.HandleSublocations,
			},
			count: 8,
			results: snsToNamedParams([]SectionName{
				snSimple,
				snAreas,
				snConnected,
				snMonsters,
				snMonsterFormations,
				snShops,
				snSongs,
				snTreasures,
			}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.locations.endpoint,
				handler:        testCfg.HandleLocations,
			},
			count: 9,
			results: snsToNamedParams([]SectionName{
				snSimple,
				snAreas,
				snConnected,
				snMonsters,
				snMonsterFormations,
				snShops,
				snSongs,
				snSublocations,
				snTreasures,
			}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/aeons/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.aeons.endpoint,
				handler:        testCfg.HandleAeons,
			},
			count: 4,
			results: snsToNamedParams([]SectionName{
				snDefaultAbilities,
				snOverdriveAbilities,
				snOverdrives,
				snStats,
			}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/characters/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.characters.endpoint,
				handler:        testCfg.HandleCharacters,
			},
			count: 5,
			results: snsToNamedParams([]SectionName{
				snDefaultAbilities,
				snStdSgAbilities,
				snExpSgAbilities,
				snOverdriveAbilities,
				snOverdrives,
			}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/character-classes/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.characterClasses.endpoint,
				handler:        testCfg.HandleCharacterClasses,
			},
			count: 4,
			results: snsToNamedParams([]SectionName{
				snDefaultAbilities,
				snLearnableOverdrives,
				snDefaultOverdrives,
				snLearnableOverdrives,
			}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/arena-creations/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.arenaCreations.endpoint,
				handler:        testCfg.HandleArenaCreations,
			},
			count:   0,
			results: []NamedParam{},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/blitzball-prizes/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.blitzballPrizes.endpoint,
				handler:        testCfg.HandleBlitzballPrizes,
			},
			count:   0,
			results: []NamedParam{},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/fmvs/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.fmvs.endpoint,
				handler:        testCfg.HandleFMVs,
			},
			count:   0,
			results: []NamedParam{},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.monsters.endpoint,
				handler:        testCfg.HandleMonsters,
			},
			count: 4,
			results: snsToNamedParams([]SectionName{
				snSimple,
				snAbilities,
				snAreas,
				snMonsterFormations,
			}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monster-formations/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.monsterFormations.endpoint,
				handler:        testCfg.HandleMonsterFormations,
			},
			count: 2,
			results: snsToNamedParams([]SectionName{
				snSimple,
				snMonsters,
			}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-modes/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.overdriveModes.endpoint,
				handler:        testCfg.HandleOverdriveModes,
			},
			count:   0,
			results: []NamedParam{},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/shops/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.shops.endpoint,
				handler:        testCfg.HandleShops,
			},
			count: 1,
			results: snsToNamedParams([]SectionName{
				snSimple,
			}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/songs/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.songs.endpoint,
				handler:        testCfg.HandleSongs,
			},
			count:   0,
			results: []NamedParam{},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sidequests/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.sidequests.endpoint,
				handler:        testCfg.HandleSidequests,
			},
			count: 1,
			results: snsToNamedParams([]SectionName{
				snSubquests,
			}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/subquests/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.subquests.endpoint,
				handler:        testCfg.HandleSubquests,
			},
			count:   0,
			results: []NamedParam{},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/treasures/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.treasures.endpoint,
				handler:        testCfg.HandleTreasures,
			},
			count:   0,
			results: []NamedParam{},
		},
	}

	testNameList(t, tests, "", "Sections", nil, compareSectionLists)
}
