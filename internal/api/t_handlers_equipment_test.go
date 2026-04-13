package api

import (
	"net/http"
	"testing"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func TestGetEquipment(t *testing.T) {
	tests := []expEquipment{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/equipment/1053",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "equipment with provided id '1053' doesn't exist. max id: 1052.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/equipment/1?table=4",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid value '4' used for parameter 'table'. 'table' can range from 1 to 3 for this resource.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/equipment/1?table=1",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"required auto-abilities":   4,
					"selectable auto-abilities": 0,
					"treasures":                 1,
					"shops":                     0,
				},
			},
			expUnique:               newExpUnique(1, "caladbolg"),
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
			testGeneral: testGeneral{
				requestURL:     "/api/equipment/1?table=2",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"required auto-abilities":   2,
					"selectable auto-abilities": 0,
					"treasures":                 1,
					"shops":                     0,
				},
			},
			expUnique:               newExpUnique(1, "caladbolg"),
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
			testGeneral: testGeneral{
				requestURL:     "/api/equipment/1?table=3",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"required auto-abilities":   1,
					"selectable auto-abilities": 0,
					"treasures":                 1,
					"shops":                     0,
				},
			},
			expUnique:               newExpUnique(1, "caladbolg"),
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
			testGeneral: testGeneral{
				requestURL:     "/api/equipment/8",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"required auto-abilities":   4,
					"selectable auto-abilities": 0,
					"treasures":                 2,
					"shops":                     0,
				},
			},
			expUnique:               newExpUnique(8, "brotherhood"),
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
			testGeneral: testGeneral{
				requestURL:     "/api/equipment/8?table=2",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"required auto-abilities":   1,
					"selectable auto-abilities": 0,
					"treasures":                 2,
					"shops":                     0,
				},
			},
			expUnique:               newExpUnique(8, "brotherhood"),
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
			testGeneral: testGeneral{
				requestURL:     "/api/equipment/333?rel_availability=pre-story",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"required auto-abilities":   1,
					"selectable auto-abilities": 0,
					"treasures":                 0,
					"shops":                     1,
				},
			},
			expUnique:               newExpUnique(333, "scout"),
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
			testGeneral: testGeneral{
				requestURL:     "/api/equipment/945",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"required auto-abilities":                      0,
					"selectable auto-abilities":                    2,
					"selectable auto-abilities 0 - auto-abilities": 4,
					"treasures": 0,
					"shops":     0,
				},
			},
			expUnique:             newExpUnique(945, "mythril armguard"),
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
	}

	testSingleResources(t, tests, "GetEquipment", testCfg.HandleEquipment, compareEquipment)
}

func TestRetrieveEquipment(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/equipment?limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   1052,
			results: []int32{1, 123, 265, 373, 429, 555, 767, 998, 1009, 1052},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/equipment?auto_abilities=10,12,14,16&character=1",
				expectedStatus: http.StatusOK,
			},
			count:   2,
			results: []int32{100, 149},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/equipment?type=armor&limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   581,
			results: []int32{472, 677, 793, 822, 999, 1052},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/equipment?celestial_weapon=true",
				expectedStatus: http.StatusOK,
			},
			count:   7,
			results: []int32{1, 2, 3, 4, 5, 6, 7},
		},
	}

	testIdList(t, tests, testCfg.e.equipment.endpoint, "RetrieveEquipment", testCfg.HandleEquipment, compareAPIResourceLists[NamedApiResourceList])
}
