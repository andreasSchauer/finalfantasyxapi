package api

import (
	"net/http"
	"testing"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func TestGetCharacter(t *testing.T) {
	tests := []expCharacter{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/characters/9",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "character with provided id '9' doesn't exist. max id: 8.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/characters/1",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"character classes": 		2,
					"default abilities": 		0,
					"standard sg abilities": 	11,
					"expert sg abilities": 		11,
				},
			},
			expUnique: expUnique{
				id: 	1,
				name: 	"tidus",
			},
			area: 				1,
			weaponType: 		"sword",
			celestialWeapon: 	h.GetInt32Ptr(1),
			overdriveCommand: 	h.GetInt32Ptr(1),
			characterClasses: 	[]int32{1, 5},
			baseStats: map[string]int32{
				"hp": 				520,
				"mp": 				12,
				"strength": 		15,
				"defense": 			10,
				"magic": 			5,
				"magic defense":	5,
				"agility": 			10,
				"luck": 			18,
				"evasion": 			10,
				"accuracy": 		10,
			},
			defaultAbilities: 	[]int32{},
			stdSgAbilities: 	[]int32{9, 10, 13, 22, 27, 38, 56, 59},
			expSgAbilities: 	[]int32{10, 22, 25, 57, 58, 59},
			overdriveModes: map[string]int32{
				"warrior": 		150,
				"victim": 		120,
				"coward": 		600,
				"daredevil": 	170,
				"loner": 		60,
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/characters/5",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"character classes": 		2,
					"default abilities": 		1,
					"standard sg abilities": 	5,
					"expert sg abilities": 		4,
				},
			},
			expUnique: expUnique{
				id: 	5,
				name: 	"kimahri",
			},
			area:				42,
			weaponType: 		"spear",
			celestialWeapon: 	h.GetInt32Ptr(5),
			overdriveCommand: 	h.GetInt32Ptr(5),
			characterClasses: 	[]int32{1, 9},
			baseStats: map[string]int32{
				"hp": 				644,
				"mp": 				78,
				"strength": 		16,
				"defense": 			15,
				"magic": 			17,
				"magic defense": 	5,
				"agility": 			6,
				"luck": 			18,
				"evasion": 			5,
				"accuracy": 		5,
			},
			defaultAbilities: 	[]int32{33},
			stdSgAbilities: 	[]int32{12, 32, 33, 52, 87},
			expSgAbilities: 	[]int32{12, 32, 33, 52},
			overdriveModes: map[string]int32{
				"comrade": 		100,
				"healer": 		100,
				"tactician": 	60,
				"rook": 		120,
				"daredevil": 	200,
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/characters/4",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"character classes":		2,
					"default abilities": 		4,
					"standard sg abilities": 	20,
					"expert sg abilities": 		19,
				},
			},
			expUnique: expUnique{
				id: 	4,
				name: 	"lulu",
			},
			area: 				36,
			weaponType: 		"doll",
			celestialWeapon: 	h.GetInt32Ptr(4),
			overdriveCommand: 	h.GetInt32Ptr(4),
			characterClasses: 	[]int32{1, 8},
			baseStats: map[string]int32{
				"hp": 				380,
				"mp": 				92,
				"strength": 		5,
				"defense": 			8,
				"magic": 			20,
				"magic defense": 	30,
				"agility": 			5,
				"luck": 			17,
				"evasion": 			40,
				"accuracy": 		3,
			},
			defaultAbilities: 	[]int32{69, 70, 71, 72},
			stdSgAbilities: 	[]int32{21, 29, 43, 69, 74, 78, 83, 86},
			expSgAbilities: 	[]int32{29, 30, 43, 70, 74, 75, 79, 81, 87},
			overdriveModes: map[string]int32{
				"comrade": 		100,
				"victim": 		130,
				"tactician": 	75,
				"slayer": 		130,
				"hero": 		70,
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/characters/8",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"character classes": 		1,
					"default abilities": 		15,
					"standard sg abilities": 	0,
					"expert sg abilities": 		0,
				},
			},
			expUnique: expUnique{
				id: 	8,
				name: 	"seymour",
			},
			area: 				103,
			weaponType: 		"seymour-staff",
			celestialWeapon: 	nil,
			overdriveCommand: 	nil,
			characterClasses: 	[]int32{12},
			baseStats: map[string]int32{
				"hp": 				1200,
				"mp": 				999,
				"strength": 		20,
				"defense": 			25,
				"magic": 			35,
				"magic defense": 	100,
				"agility": 			20,
				"luck": 			18,
				"evasion": 			10,
				"accuracy": 		10,
			},
			defaultAbilities: 	[]int32{45, 46, 48, 69,71, 74, 75, 76},
			stdSgAbilities: 	[]int32{},
			expSgAbilities: 	[]int32{},
			overdriveModes: 	nil,
		},
	}

	testSingleResources(t, tests, "GetCharacter", testCfg.HandleCharacters, compareCharacters)
}

func TestRetrieveCharacters(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/characters",
				expectedStatus: http.StatusOK,
			},
			count:   8,
			results: []int32{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/characters?story_based=true",
				expectedStatus: http.StatusOK,
			},
			count:   1,
			results: []int32{8},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/characters?underwater=true",
				expectedStatus: http.StatusOK,
			},
			count:   3,
			results: []int32{1, 3, 7},
		},
	}

	testIdList(t, tests, testCfg.e.characters.endpoint, "RetrieveCharacters", testCfg.HandleCharacters, compareAPIResourceLists[NamedApiResourceList])
}