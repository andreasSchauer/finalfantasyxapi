package api

import (
	"net/http"
	"testing"
)

func TestGetCharacterClass(t *testing.T) {
	tests := []expCharacterClass{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/character-classes/23",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "character class with provided id '23' doesn't exist. max id: 22.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/character-classes/1",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"members":              7,
					"default abilities":    5,
					"learnable abilities":  85,
					"default overdrives":   0,
					"learnable overdrives": 0,
					"submenus":             11,
				},
			},
			expUnique: expUnique{
				id:   1,
				name: "characters",
			},
			category: "group",
			members: []string{
				"/characters/1",
				"/characters/2",
				"/characters/3",
				"/characters/4",
				"/characters/5",
				"/characters/6",
				"/characters/7",
			},
			defaultAbilities: []string{
				"/other-abilities/1",
				"/other-abilities/3",
				"/other-abilities/4",
				"/other-abilities/5",
				"/other-abilities/6",
			},
			learnableAbilities: []string{
				"/player-abilities/1",
				"/player-abilities/12",
				"/player-abilities/27",
				"/player-abilities/29",
				"/player-abilities/35",
				"/player-abilities/51",
				"/player-abilities/69",
				"/player-abilities/87",
			},
			defaultOverdrives:   []int32{},
			learnableOverdrives: []int32{},
			submenus:            []int32{2, 3, 4, 5, 6, 7, 8, 9, 11, 13, 15},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/character-classes/3",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"members":              6,
					"default abilities":    3,
					"learnable abilities":  67,
					"default overdrives":   0,
					"learnable overdrives": 0,
					"submenus":             4,
				},
			},
			expUnique: expUnique{
				id:   3,
				name: "standard-aeons",
			},
			category: "group",
			members: []string{
				"/aeons/1",
				"/aeons/2",
				"/aeons/3",
				"/aeons/4",
				"/aeons/5",
				"/aeons/6",
			},
			defaultAbilities: []string{
				"/other-abilities/2",
				"/other-abilities/7",
				"/other-abilities/8",
			},
			learnableAbilities: []string{
				"/player-abilities/3",
				"/player-abilities/14",
				"/player-abilities/19",
				"/player-abilities/29",
				"/player-abilities/50",
				"/player-abilities/57",
				"/player-abilities/64",
				"/player-abilities/85",
			},
			defaultOverdrives:   []int32{},
			learnableOverdrives: []int32{},
			submenus:            []int32{2, 3, 4, 5},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/character-classes/4",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"members":              3,
					"default abilities":    2,
					"learnable abilities":  0,
					"default overdrives":   1,
					"learnable overdrives": 0,
					"submenus":             0,
				},
			},
			expUnique: expUnique{
				id:   4,
				name: "magus-sisters",
			},
			category: "group",
			members: []string{
				"/aeons/8",
				"/aeons/9",
				"/aeons/10",
			},
			defaultAbilities: []string{
				"/player-abilities/102",
				"/other-abilities/2",
			},
			learnableAbilities:  []string{},
			defaultOverdrives:   []int32{124},
			learnableOverdrives: []int32{},
			submenus:            []int32{},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/character-classes/6",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"members":              1,
					"default abilities":    9,
					"learnable abilities":  7,
					"default overdrives":   1,
					"learnable overdrives": 7,
					"submenus":             1,
				},
			},
			expUnique: expUnique{
				id:   6,
				name: "yuna",
			},
			category: "character",
			members: []string{
				"/characters/2",
			},
			defaultAbilities: []string{
				"/player-abilities/45",
				"/player-abilities/53",
				"/other-abilities/10",
				"/other-abilities/12",
				"/other-abilities/13",
				"/other-abilities/14",
				"/other-abilities/15",
				"/other-abilities/16",
				"/other-abilities/17",
			},
			learnableAbilities: []string{
				"/other-abilities/18",
				"/other-abilities/19",
				"/other-abilities/20",
				"/other-abilities/21",
				"/other-abilities/22",
				"/other-abilities/23",
				"/other-abilities/24",
			},
			defaultOverdrives:   []int32{5},
			learnableOverdrives: []int32{6, 7, 8, 9, 10, 11, 12},
			submenus:            []int32{1},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/character-classes/8",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"members":              1,
					"default abilities":    10,
					"learnable abilities":  0,
					"default overdrives":   4,
					"learnable overdrives": 15,
					"submenus":             1,
				},
			},
			expUnique: expUnique{
				id:   8,
				name: "lulu",
			},
			category: "character",
			members: []string{
				"/characters/4",
			},
			defaultAbilities: []string{
				"/player-abilities/69",
				"/player-abilities/70",
				"/player-abilities/71",
				"/player-abilities/72",
				"/other-abilities/10",
				"/other-abilities/11",
				"/other-abilities/12",
				"/other-abilities/14",
				"/other-abilities/15",
				"/other-abilities/16",
			},
			learnableAbilities:  []string{},
			defaultOverdrives:   []int32{17, 18, 19, 20},
			learnableOverdrives: []int32{21, 26, 30, 32, 34, 35},
			submenus:            []int32{14},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/character-classes/19",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"members":              1,
					"default abilities":    5,
					"learnable abilities":  0,
					"default overdrives":   0,
					"learnable overdrives": 0,
					"submenus":             1,
				},
			},
			expUnique: expUnique{
				id:   19,
				name: "yojimbo",
			},
			category: "aeon",
			members: []string{
				"/aeons/7",
			},
			defaultAbilities: []string{
				"/player-abilities/94",
				"/player-abilities/95",
				"/player-abilities/96",
				"/player-abilities/97",
				"/player-abilities/98",
			},
			learnableAbilities:  []string{},
			defaultOverdrives:   []int32{},
			learnableOverdrives: []int32{},
			submenus:            []int32{12},
		},
	}

	testSingleResources(t, tests, "GetCharacterClass", testCfg.HandleCharacterClasses, compareCharacterClasses)
}

func TestRetrieveCharacterClasses(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/character-classes?limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   22,
			results: []int32{1, 5, 8, 12, 13, 17, 22},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/character-classes?category=1",
				expectedStatus: http.StatusOK,
			},
			count:   4,
			results: []int32{1, 2, 3, 4},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/character-classes?category=character",
				expectedStatus: http.StatusOK,
			},
			count:   8,
			results: []int32{5, 7, 8, 9, 10, 12},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/character-classes?category=aeon",
				expectedStatus: http.StatusOK,
			},
			count:   10,
			results: []int32{13, 15, 18, 20, 21, 22},
		},
	}

	testIdList(t, tests, testCfg.e.characterClasses.endpoint, "RetrieveCharacterClasses", testCfg.HandleCharacterClasses, compareAPIResourceLists[NamedApiResourceList])
}
