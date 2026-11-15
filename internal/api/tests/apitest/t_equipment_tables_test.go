package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/tests/apitesttypes"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

func TestGetEquipmentTable(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.EquipmentTable{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/equipment-tables/164",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "equipment table with provided id '164' doesn't exist. max id: 163",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/equipment-tables/2",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"required auto-abilities":   4,
					"selectable auto-abilities": 0,
					"equipment":                 1,
				},
			},
			IDOnly:                  apitesttypes.NewExpIDOnly(2),
			CelestialWeapon:         h.GetInt32Ptr(2),
			SpecificCharacter:       h.GetInt32Ptr(2),
			RequiredAutoAbilities:   []int32{51, 49, 44, 43},
			SelectableAutoAbilities: []apitesttypes.TestAbilityPool{},
			RequiredSlots:           nil,
			Equipment:               []int32{2},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/equipment-tables/20",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"required auto-abilities":   1,
					"selectable auto-abilities": 0,
					"equipment":                 7,
				},
			},
			IDOnly:                  apitesttypes.NewExpIDOnly(20),
			CelestialWeapon:         nil,
			SpecificCharacter:       nil,
			RequiredAutoAbilities:   []int32{45},
			SelectableAutoAbilities: []apitesttypes.TestAbilityPool{},
			RequiredSlots:           nil,
			Equipment:               []int32{65, 66, 67, 68, 69, 70, 71},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
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
			IDOnly:                apitesttypes.NewExpIDOnly(145),
			CelestialWeapon:       nil,
			SpecificCharacter:     nil,
			RequiredAutoAbilities: []int32{},
			SelectableAutoAbilities: []apitesttypes.TestAbilityPool{
				{
					Index:         0,
					AutoAbilities: []int32{114, 113, 112, 111},
					ReqAmount:     1,
				},
				{
					Index:         1,
					AutoAbilities: []int32{118, 117, 116, 115},
					ReqAmount:     1,
				},
			},
			RequiredSlots: nil,
			Equipment:     []int32{929, 930, 931, 932, 933, 934, 935},
		},
	}

	testSingleResources(t, tests, "GetEquipmentTable", testSrv.HandleEquipmentTables)
}

func TestRetrieveEquipmentTables(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.ListIDs{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/equipment-tables?limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   163,
			Results: []int32{1, 25, 58, 63, 77, 117, 129, 151, 163},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/equipment-tables?auto_abilities=7,8",
				ExpectedStatus: http.StatusOK,
			},
			Count:   3,
			Results: []int32{13, 31, 49},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/equipment-tables?type=weapon&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   78,
			Results: []int32{1, 13, 29, 45, 67, 78},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/equipment-tables?celestial_weapon=false&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   154,
			Results: []int32{10, 28, 47, 84, 94, 113, 127, 128, 145, 154},
		},
	}

	testIDList(t, tests, apilookups.E.EquipmentTables.Endpoint, "RetrieveEquipmentTables", testSrv.HandleEquipmentTables, compareAPIResourceLists[apiresource.Unnamed])
}
