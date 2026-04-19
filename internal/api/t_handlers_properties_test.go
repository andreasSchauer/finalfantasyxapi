package api

import (
	"net/http"
	"testing"
)

func TestGetProperty(t *testing.T) {
	tests := []expProperty{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/properties/13",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "property with provided id '13' doesn't exist. max id: 12.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/properties/1",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"auto-abilities": 	0,
					"monsters": 		53,
				},
			},
			expUnique: 		newExpUnique(1, "armored"),
			autoAbilities: 	[]int32{},
			monsters: 		[]int32{5, 29, 79, 107, 197, 277, 292},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/properties/4",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"auto-abilities": 	1,
					"monsters": 		0,
				},
			},
			expUnique: 		newExpUnique(4, "piercing"),
			autoAbilities: 	[]int32{2},
			monsters: 		[]int32{},
		},
	}

	testSingleResources(t, tests, "GetProperty", testCfg.HandleProperties, compareProperties)
}

func TestRetrieveProperties(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/properties",
				expectedStatus: http.StatusOK,
			},
			count:   12,
			results: []int32{1, 4, 5, 8, 9, 10, 12},
		},
	}

	testIdList(t, tests, testCfg.e.properties.endpoint, "RetrieveProperties", testCfg.HandleProperties, compareAPIResourceLists[NamedApiResourceList])
}
