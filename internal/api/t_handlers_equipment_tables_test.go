package api

import (
	"net/http"
	"testing"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func TestGetEquipmentTable(t *testing.T) {
	tests := []expEquipmentTable{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/equipment-tables/164",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "equipment table with provided id '164' doesn't exist. max id: 163.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/equipment-tables/2",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"required auto-abilities":     	4,
					"selectable auto-abilities":	0,
					"equipment":            		1,
				},
			},
			expIdOnly:   newExpIdOnly(2),
			celestialWeapon: 		 h.GetInt32Ptr(2),
			specificCharacter: 		 h.GetInt32Ptr(2),
			requiredAutoAbilities: 	 []int32{51, 49, 44, 43},
			selectableAutoAbilities: []testAbilityPool{},
			emptySlotsAmt: 			 0,
			equipment: 				 []int32{2},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/equipment-tables/20",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"required auto-abilities":     	1,
					"selectable auto-abilities":	0,
					"equipment":            		7,
				},
			},
			expIdOnly:   newExpIdOnly(20),
			celestialWeapon: 		 nil,
			specificCharacter: 		 nil,
			requiredAutoAbilities: 	 []int32{45},
			selectableAutoAbilities: []testAbilityPool{},
			emptySlotsAmt: 			 0,
			equipment: 				 []int32{65, 66, 67, 68, 69, 70, 71},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/equipment-tables/145",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"required auto-abilities":     	0,
					"selectable auto-abilities":	2,
					"selectable auto-abilities 0 - auto-abilities": 4,
					"selectable auto-abilities 1 - auto-abilities": 4,
					"equipment":            		7,
				},
			},
			expIdOnly:   newExpIdOnly(145),
			celestialWeapon: 		 nil,
			specificCharacter: 		 nil,
			requiredAutoAbilities: 	 []int32{},
			selectableAutoAbilities: []testAbilityPool{
				{
					index:   		0,
					autoAbilities:	[]int32{114, 113, 112, 111},
					reqAmount:   	1,
				},
				{
					index:   		1,
					autoAbilities:	[]int32{118, 117, 116, 115},
					reqAmount:   	1,
				},
			},
			emptySlotsAmt: 			 0,
			equipment: 				 []int32{929, 930, 931, 932, 933, 934, 935},
		},
	}

	testSingleResources(t, tests, "GetEquipmentTable", testCfg.HandleEquipmentTables, compareEquipmentTables)
}

func TestRetrieveEquipmentTables(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/equipment-tables?limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   163,
			results: []int32{1, 25, 58, 63, 77, 117, 129, 151, 163},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/equipment-tables?auto_abilities=7,8",
				expectedStatus: http.StatusOK,
			},
			count:   3,
			results: []int32{13, 31, 49},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/equipment-tables?type=weapon&limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   78,
			results: []int32{1, 13, 29, 45, 67, 78},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/equipment-tables?celestial_weapon=false&limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   154,
			results: []int32{10, 28, 47, 84, 94, 113, 127, 128, 145, 154},
		},
	}

	testIdList(t, tests, testCfg.e.equipmentTables.endpoint, "RetrieveEquipmentTables", testCfg.HandleEquipmentTables, compareAPIResourceLists[UnnamedApiResourceList])
}
