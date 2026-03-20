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
			members: []int32{1, 2, 3, 4, 5, 6, 7},
			defaultAbilities: []int32{374, 376, 377, 378, 379},
			learnableAbilities: []int32{1, 12, 27, 29, 35, 51, 69, 87},
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
			members: []int32{9, 10, 11, 12, 13, 14},
			defaultAbilities: []int32{375, 380, 381},
			learnableAbilities: []int32{3, 14, 19, 29, 50, 57, 64, 85},
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
			members: []int32{16, 17, 18},
			defaultAbilities: []int32{102, 375},
			learnableAbilities:  []int32{},
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
			members: []int32{2},
			defaultAbilities: []int32{45, 53, 383, 385, 386, 387, 388, 389, 390},
			learnableAbilities: []int32{391, 392, 393, 394, 395, 396, 397},
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
			members: []int32{4},
			defaultAbilities: []int32{69, 70, 71, 72, 383, 384, 385, 387, 388, 389},
			learnableAbilities:  []int32{},
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
			members: []int32{15},
			defaultAbilities: []int32{94, 95, 96, 97, 98},
			learnableAbilities:  []int32{},
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
