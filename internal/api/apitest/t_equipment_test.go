package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

func TestGetEquipment(t *testing.T) {
	t.Parallel()
	tests := []Equipment{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/equipment/1053",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "equipment with provided id '1053' doesn't exist. max id: 1052",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/equipment/1?table=4",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "invalid value '4' used for parameter 'table'. 'table' can range from 1 to 3 for this resource",
			},
		},
		{
			TestGeneral: TestGeneral{
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
			Unique:                  newExpUnique(1, "caladbolg"),
			equipmentTable:          1,
			priority:                h.GetInt32Ptr(1),
			celestialWeapon:         h.GetInt32Ptr(1),
			requiredAutoAbilities:   []int32{51, 49, 38, 39},
			selectableAutoAbilities: []testAbilityPool{},
			requiredSlots:           nil,
			treasures:               []int32{270},
			shops:                   []int32{},
		},
		{
			TestGeneral: TestGeneral{
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
			Unique:                  newExpUnique(1, "caladbolg"),
			equipmentTable:          8,
			priority:                h.GetInt32Ptr(1),
			celestialWeapon:         h.GetInt32Ptr(1),
			requiredAutoAbilities:   []int32{53, 48},
			selectableAutoAbilities: []testAbilityPool{},
			requiredSlots:           h.GetInt32Ptr(2),
			treasures:               []int32{270},
			shops:                   []int32{},
		},
		{
			TestGeneral: TestGeneral{
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
			Unique:                  newExpUnique(1, "caladbolg"),
			equipmentTable:          9,
			priority:                h.GetInt32Ptr(1),
			celestialWeapon:         h.GetInt32Ptr(1),
			requiredAutoAbilities:   []int32{53},
			selectableAutoAbilities: []testAbilityPool{},
			requiredSlots:           h.GetInt32Ptr(3),
			treasures:               []int32{270},
			shops:                   []int32{},
		},
		{
			TestGeneral: TestGeneral{
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
			Unique:                  newExpUnique(8, "brotherhood"),
			equipmentTable:          10,
			priority:                h.GetInt32Ptr(2),
			celestialWeapon:         nil,
			requiredAutoAbilities:   []int32{30, 31, 7, 1},
			selectableAutoAbilities: []testAbilityPool{},
			requiredSlots:           nil,
			treasures:               []int32{36, 174},
			shops:                   []int32{},
		},
		{
			TestGeneral: TestGeneral{
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
			Unique:                  newExpUnique(8, "brotherhood"),
			equipmentTable:          11,
			priority:                h.GetInt32Ptr(2),
			celestialWeapon:         nil,
			requiredAutoAbilities:   []int32{30},
			selectableAutoAbilities: []testAbilityPool{},
			requiredSlots:           h.GetInt32Ptr(3),
			treasures:               []int32{36, 174},
			shops:                   []int32{},
		},
		{
			TestGeneral: TestGeneral{
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
			Unique:                  newExpUnique(333, "scout"),
			equipmentTable:          58,
			priority:                h.GetInt32Ptr(49),
			celestialWeapon:         nil,
			requiredAutoAbilities:   []int32{1},
			selectableAutoAbilities: []testAbilityPool{},
			requiredSlots:           nil,
			treasures:               []int32{},
			shops:                   []int32{2},
		},
		{
			TestGeneral: TestGeneral{
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
			Unique:                newExpUnique(945, "mythril armguard"),
			equipmentTable:        147,
			priority:              h.GetInt32Ptr(69),
			celestialWeapon:       nil,
			requiredAutoAbilities: []int32{},
			selectableAutoAbilities: []testAbilityPool{
				{
					index:         0,
					autoAbilities: []int32{106, 105, 104, 103},
					reqAmount:     1,
				},
				{
					index:         1,
					autoAbilities: []int32{110, 109, 108, 107},
					reqAmount:     1,
				},
			},
			requiredSlots: nil,
			treasures:     []int32{},
			shops:         []int32{},
		},
		{
			TestGeneral: TestGeneral{
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
			Unique:    newExpUnique(805, "lucid ring"),
			treasures: []int32{194},
			shops:     []int32{},
		},
		{
			TestGeneral: TestGeneral{
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
			Unique:    newExpUnique(805, "lucid ring"),
			treasures: []int32{261},
			shops:     []int32{},
		},
		{
			TestGeneral: TestGeneral{
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
			Unique:    newExpUnique(805, "lucid ring"),
			treasures: []int32{},
			shops:     []int32{5},
		},
	}

	testSingleResources(t, tests, "GetEquipment", testSrv.HandleEquipment)
}

func TestRetrieveEquipment(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/equipment?limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   1052,
			Results: []int32{1, 123, 265, 373, 429, 555, 767, 998, 1009, 1052},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/equipment?auto_abilities=10,12,14,16&character=1",
				ExpectedStatus: http.StatusOK,
			},
			Count:   2,
			Results: []int32{100, 149},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/equipment?type=armor&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   581,
			Results: []int32{472, 677, 793, 822, 999, 1052},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/equipment?celestial_weapon=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   7,
			Results: []int32{1, 2, 3, 4, 5, 6, 7},
		},
	}

	testIDList(t, tests, apilookups.E.Equipment.Endpoint, "RetrieveEquipment", testSrv.HandleEquipment, compareAPIResourceLists[apiresource.Named])
}
