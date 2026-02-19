package api

import (
	"net/http"
	"testing"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func TestGetAeon(t *testing.T) {
	tests := []expAeon{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/aeons/11",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "aeon with provided id '11' doesn't exist. max id: 10.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/aeons/3",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"character classes": 	3,
					"aeon commands": 		0,
					"default abilities":	5,
					"overdrives":			1,
					"weapon abilities":		3,
					"armor abilities":		4,
				},
			},
			expUnique: expUnique{
				id:		3,
				name: 	"ixion",
			},
			area: 				115,
			battlesToRegen: 	20,
			agility: AgilityParams{
				TickSpeed: 	15,
				MinICV: 	h.GetInt32Ptr(43),
				MaxICV: 	h.GetInt32Ptr(45),
			},
			celestialWeapon: 	h.GetInt32Ptr(5),
			characterClasses: 	[]int32{2, 3, 15},
			baseStats: map[string]int32{
				"hp": 				891,
				"mp": 				25,
				"strength": 		20,
				"defense": 			26,
				"magic": 			20,
				"magic defense": 	29,
				"agility": 			8,
				"luck": 			17,
				"evasion": 			11,
				"accuracy": 		12,
			},
			aeonCommands: 		[]int32{},
			defaultAbilities: 	[]int32{49, 50, 70, 74, 90},
			overdrives: 		[]int32{120},
			weaponAbilities: []expAeonEquipment{
				{
					autoAbility: 		54,
					celestialWeapon: 	false,
				},
				{
					autoAbility: 		2,
					celestialWeapon: 	false,
				},
				{
					autoAbility: 		51,
					celestialWeapon: 	true,
				},
			},
			armorAbilities: []expAeonEquipment{
				{
					autoAbility: 		1,
					celestialWeapon: 	false,
				},
				{
					autoAbility: 		128,
					celestialWeapon: 	false,
				},
				{
					autoAbility: 		127,
					celestialWeapon: 	false,
				},
				{
					autoAbility: 		60,
					celestialWeapon: 	false,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/aeons/9?battles=347&yuna_stats=strength-24",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{
					"weapon abilities": true,
					"armor abilities": 	true,
				},
				expLengths: map[string]int{
					"character classes": 	3,
					"aeon commands": 		5,
					"default abilities":	10,
					"overdrives":			1,
					"weapon abilities":		3,
					"armor abilities":		3,
				},
			},
			expUnique: expUnique{
				id:		9,
				name: 	"sandy",
			},
			area: 				210,
			battlesToRegen: 	30,
			agility: AgilityParams{
				TickSpeed: 	10,
				MinICV: 	h.GetInt32Ptr(26),
				MaxICV: 	h.GetInt32Ptr(30),
			},
			celestialWeapon: 	nil,
			characterClasses: 	[]int32{2, 4, 21},
			baseStats: map[string]int32{
				"hp": 				4590,
				"mp": 				71,
				"strength": 		149,
				"defense": 			53,
				"magic": 			53,
				"magic defense": 	57,
				"agility": 			22,
				"luck": 			17,
				"evasion": 			27,
				"accuracy": 		18,
			},
			aeonCommands: 		[]int32{2, 3, 4, 7, 9},
			defaultAbilities: 	[]int32{45,46, 47, 56, 60, 61, 62, 68, 100, 102},
			overdrives: 		[]int32{124},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/aeons/1?yuna_stats=hp-2000,mp-250,strength-10,defense-28,magic-72,magic_defense-50,agility-41,luck-25,evasion-30,accuracy-15",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{
					"weapon abilities": true,
					"armor abilities": 	true,
				},
				expLengths: map[string]int{
					"character classes": 	3,
					"aeon commands": 		0,
					"default abilities":	5,
					"overdrives":			2,
				},
			},
			expUnique: expUnique{
				id:		1,
				name: 	"valefor",
			},
			area: 				33,
			battlesToRegen: 	8,
			agility: AgilityParams{
				TickSpeed: 	7,
				MinICV: 	h.GetInt32Ptr(20),
				MaxICV: 	h.GetInt32Ptr(21),
			},
			celestialWeapon: 	h.GetInt32Ptr(2),
			characterClasses: 	[]int32{2, 3, 13},
			baseStats: map[string]int32{
				"hp": 				2146,
				"mp": 				68,
				"strength": 		47,
				"defense": 			72,
				"magic": 			76,
				"magic defense": 	59,
				"agility": 			35,
				"luck": 			25,
				"evasion": 			27,
				"accuracy": 		44,
			},
			aeonCommands: 		[]int32{},
			defaultAbilities: 	[]int32{69, 70, 71, 72, 88},
			overdrives: 		[]int32{117, 118},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/aeons/1?yuna_stats=hp-2000,mp-250,strength-10,defense-28,magic-72,magic_defense-50,agility-41,luck-25,evasion-30,accuracy-15&battles=600",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{
					"area": 					true,
					"battles to regenerate": 	true,
					"celestial weapon": 		true,
					"character classes": 		true,
					"aeon commands":			true,
					"default abilities":		true,
					"overdrives":				true,
					"weapon abilities": 		true,
					"armor abilities": 			true,
				},
				expLengths: map[string]int{},
			},
			expUnique: expUnique{
				id:		1,
				name: 	"valefor",
			},
			agility: AgilityParams{
				TickSpeed: 	7,
				MinICV: 	h.GetInt32Ptr(20),
				MaxICV: 	h.GetInt32Ptr(21),
			},
			baseStats: map[string]int32{
				"hp": 				2225,
				"mp": 				71,
				"strength": 		47,
				"defense": 			72,
				"magic": 			76,
				"magic defense": 	69,
				"agility": 			35,
				"luck": 			25,
				"evasion": 			42,
				"accuracy": 		44,
			},
		},
	}

	testSingleResources(t, tests, "GetAeon", testCfg.HandleAeons, compareAeons)
}

func TestRetrieveAeons(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/aeons",
				expectedStatus: http.StatusOK,
			},
			count:   10,
			results: []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/aeons?optional=true",
				expectedStatus: http.StatusOK,
			},
			count:   5,
			results: []int32{6, 7, 8, 9, 10},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/aeons?optional=false",
				expectedStatus: http.StatusOK,
			},
			count:   5,
			results: []int32{1, 2, 3, 4, 5},
		},
	}

	testIdList(t, tests, testCfg.e.aeons.endpoint, "RetrieveAeons", testCfg.HandleAeons, compareAPIResourceLists[NamedApiResourceList])
}