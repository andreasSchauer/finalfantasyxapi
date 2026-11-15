package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/tests/apitesttypes"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

func TestGetEquipment(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.Equipment{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/equipment/1053",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "equipment with provided id '1053' doesn't exist. max id: 1052",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/equipment/1?table=4",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "invalid value '4' used for parameter 'table'. 'table' can range from 1 to 3 for this resource",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/equipment/1?table=1",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"required auto-abilities":   4,
					"selectable auto-abilities": 0,
					"treasures":                 1,
					"shops":                     0,
				},
			},
			Unique:                  apitesttypes.NewExpUnique(1, "caladbolg"),
			EquipmentTable:          1,
			Priority:                h.GetInt32Ptr(1),
			CelestialWeapon:         h.GetInt32Ptr(1),
			RequiredAutoAbilities:   []int32{51, 49, 38, 39},
			SelectableAutoAbilities: []apitesttypes.TestAbilityPool{},
			RequiredSlots:           nil,
			Treasures:               []int32{270},
			Shops:                   []int32{},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/equipment/1?table=2",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"required auto-abilities":   2,
					"selectable auto-abilities": 0,
					"treasures":                 1,
					"shops":                     0,
				},
			},
			Unique:                  apitesttypes.NewExpUnique(1, "caladbolg"),
			EquipmentTable:          8,
			Priority:                h.GetInt32Ptr(1),
			CelestialWeapon:         h.GetInt32Ptr(1),
			RequiredAutoAbilities:   []int32{53, 48},
			SelectableAutoAbilities: []apitesttypes.TestAbilityPool{},
			RequiredSlots:           h.GetInt32Ptr(2),
			Treasures:               []int32{270},
			Shops:                   []int32{},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/equipment/1?table=3",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"required auto-abilities":   1,
					"selectable auto-abilities": 0,
					"treasures":                 1,
					"shops":                     0,
				},
			},
			Unique:                  apitesttypes.NewExpUnique(1, "caladbolg"),
			EquipmentTable:          9,
			Priority:                h.GetInt32Ptr(1),
			CelestialWeapon:         h.GetInt32Ptr(1),
			RequiredAutoAbilities:   []int32{53},
			SelectableAutoAbilities: []apitesttypes.TestAbilityPool{},
			RequiredSlots:           h.GetInt32Ptr(3),
			Treasures:               []int32{270},
			Shops:                   []int32{},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/equipment/8",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"required auto-abilities":   4,
					"selectable auto-abilities": 0,
					"treasures":                 2,
					"shops":                     0,
				},
			},
			Unique:                  apitesttypes.NewExpUnique(8, "brotherhood"),
			EquipmentTable:          10,
			Priority:                h.GetInt32Ptr(2),
			CelestialWeapon:         nil,
			RequiredAutoAbilities:   []int32{30, 31, 7, 1},
			SelectableAutoAbilities: []apitesttypes.TestAbilityPool{},
			RequiredSlots:           nil,
			Treasures:               []int32{36, 174},
			Shops:                   []int32{},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/equipment/8?table=2",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"required auto-abilities":   1,
					"selectable auto-abilities": 0,
					"treasures":                 2,
					"shops":                     0,
				},
			},
			Unique:                  apitesttypes.NewExpUnique(8, "brotherhood"),
			EquipmentTable:          11,
			Priority:                h.GetInt32Ptr(2),
			CelestialWeapon:         nil,
			RequiredAutoAbilities:   []int32{30},
			SelectableAutoAbilities: []apitesttypes.TestAbilityPool{},
			RequiredSlots:           h.GetInt32Ptr(3),
			Treasures:               []int32{36, 174},
			Shops:                   []int32{},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/equipment/333?rel_availability=pre-story",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"required auto-abilities":   1,
					"selectable auto-abilities": 0,
					"treasures":                 0,
					"shops":                     1,
				},
			},
			Unique:                  apitesttypes.NewExpUnique(333, "scout"),
			EquipmentTable:          58,
			Priority:                h.GetInt32Ptr(49),
			CelestialWeapon:         nil,
			RequiredAutoAbilities:   []int32{1},
			SelectableAutoAbilities: []apitesttypes.TestAbilityPool{},
			RequiredSlots:           nil,
			Treasures:               []int32{},
			Shops:                   []int32{2},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/equipment/945",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"required auto-abilities":                      0,
					"selectable auto-abilities":                    2,
					"selectable auto-abilities 0 - auto-abilities": 4,
					"treasures": 0,
					"shops":     0,
				},
			},
			Unique:                apitesttypes.NewExpUnique(945, "mythril armguard"),
			EquipmentTable:        147,
			Priority:              h.GetInt32Ptr(69),
			CelestialWeapon:       nil,
			RequiredAutoAbilities: []int32{},
			SelectableAutoAbilities: []apitesttypes.TestAbilityPool{
				{
					Index:         0,
					AutoAbilities: []int32{106, 105, 104, 103},
					ReqAmount:     1,
				},
				{
					Index:         1,
					AutoAbilities: []int32{110, 109, 108, 107},
					ReqAmount:     1,
				},
			},
			RequiredSlots: nil,
			Treasures:     []int32{},
			Shops:         []int32{},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/equipment/805?rel_availability=always",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"equipment table":           true,
					"priority":                  true,
					"celestial weapon":          true,
					"required auto-abilities":   true,
					"selectable auto-abilities": true,
					"required slots":            true,
				},
				ExpLengths: map[string]int{
					"treasures": 1,
					"shops":     0,
				},
			},
			Unique:    apitesttypes.NewExpUnique(805, "lucid ring"),
			Treasures: []int32{194},
			Shops:     []int32{},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/equipment/805?rel_availability=pre-story",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"equipment table":           true,
					"priority":                  true,
					"celestial weapon":          true,
					"required auto-abilities":   true,
					"selectable auto-abilities": true,
					"required slots":            true,
				},
				ExpLengths: map[string]int{
					"treasures": 1,
					"shops":     0,
				},
			},
			Unique:    apitesttypes.NewExpUnique(805, "lucid ring"),
			Treasures: []int32{261},
			Shops:     []int32{},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/equipment/805?rel_availability=post",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"equipment table":           true,
					"priority":                  true,
					"celestial weapon":          true,
					"required auto-abilities":   true,
					"selectable auto-abilities": true,
					"required slots":            true,
				},
				ExpLengths: map[string]int{
					"treasures": 0,
					"shops":     1,
				},
			},
			Unique:    apitesttypes.NewExpUnique(805, "lucid ring"),
			Treasures: []int32{},
			Shops:     []int32{5},
		},
	}

	testSingleResources(t, tests, "GetEquipment", testSrv.HandleEquipment)
}

func TestRetrieveEquipment(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.ListIDs{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/equipment?limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   1052,
			Results: []int32{1, 123, 265, 373, 429, 555, 767, 998, 1009, 1052},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/equipment?auto_abilities=10,12,14,16&character=1",
				ExpectedStatus: http.StatusOK,
			},
			Count:   2,
			Results: []int32{100, 149},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/equipment?type=armor&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   581,
			Results: []int32{472, 677, 793, 822, 999, 1052},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/equipment?celestial_weapon=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   7,
			Results: []int32{1, 2, 3, 4, 5, 6, 7},
		},
	}

	testIDList(t, tests, apilookups.E.Equipment.Endpoint, "RetrieveEquipment", testSrv.HandleEquipment, compareAPIResourceLists[apiresource.Named])
}
