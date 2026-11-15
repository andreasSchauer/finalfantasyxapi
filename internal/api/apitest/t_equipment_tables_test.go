package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

func TestGetEquipmentTable(t *testing.T) {
	t.Parallel()
	tests := []EquipmentTable{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/equipment-tables/164",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "equipment table with provided id '164' doesn't exist. max id: 163",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/equipment-tables/2",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"required auto-abilities":   4,
					"selectable auto-abilities": 0,
					"equipment":                 1,
				},
			},
			IDOnly:                  newExpIDOnly(2),
			celestialWeapon:         h.GetInt32Ptr(2),
			specificCharacter:       h.GetInt32Ptr(2),
			requiredAutoAbilities:   []int32{51, 49, 44, 43},
			selectableAutoAbilities: []testAbilityPool{},
			requiredSlots:           nil,
			equipment:               []int32{2},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/equipment-tables/20",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"required auto-abilities":   1,
					"selectable auto-abilities": 0,
					"equipment":                 7,
				},
			},
			IDOnly:                  newExpIDOnly(20),
			celestialWeapon:         nil,
			specificCharacter:       nil,
			requiredAutoAbilities:   []int32{45},
			selectableAutoAbilities: []testAbilityPool{},
			requiredSlots:           nil,
			equipment:               []int32{65, 66, 67, 68, 69, 70, 71},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/equipment-tables/145",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"required auto-abilities":                      0,
					"selectable auto-abilities":                    2,
					"selectable auto-abilities 0 - auto-abilities": 4,
					"selectable auto-abilities 1 - auto-abilities": 4,
					"equipment": 7,
				},
			},
			IDOnly:                newExpIDOnly(145),
			celestialWeapon:       nil,
			specificCharacter:     nil,
			requiredAutoAbilities: []int32{},
			selectableAutoAbilities: []testAbilityPool{
				{
					index:         0,
					autoAbilities: []int32{114, 113, 112, 111},
					reqAmount:     1,
				},
				{
					index:         1,
					autoAbilities: []int32{118, 117, 116, 115},
					reqAmount:     1,
				},
			},
			requiredSlots: nil,
			equipment:     []int32{929, 930, 931, 932, 933, 934, 935},
		},
	}

	testSingleResources(t, tests, "GetEquipmentTable", testSrv.HandleEquipmentTables)
}

func TestRetrieveEquipmentTables(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/equipment-tables?limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   163,
			Results: []int32{1, 25, 58, 63, 77, 117, 129, 151, 163},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/equipment-tables?auto_abilities=7,8",
				ExpectedStatus: http.StatusOK,
			},
			Count:   3,
			Results: []int32{13, 31, 49},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/equipment-tables?type=weapon&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   78,
			Results: []int32{1, 13, 29, 45, 67, 78},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/equipment-tables?celestial_weapon=false&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   154,
			Results: []int32{10, 28, 47, 84, 94, 113, 127, 128, 145, 154},
		},
	}

	testIDList(t, tests, apilookups.E.EquipmentTables.Endpoint, "RetrieveEquipmentTables", testSrv.HandleEquipmentTables, compareAPIResourceLists[apiresource.Unnamed])
}
