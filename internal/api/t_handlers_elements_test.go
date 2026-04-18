package api

import (
	"net/http"
	"testing"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func TestGetElement(t *testing.T) {
	tests := []expElement{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/elements/6",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "element with provided id '6' doesn't exist. max id: 5.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/elements/2",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"auto-abilities": 		4,
					"player abilities": 	3,
					"overdrive abilities": 	17,
					"item abilities": 		3,
					"enemy abilities": 		7,
					"monsters weak": 		39,
					"monsters halved": 		24,
					"monsters immune": 		7,
					"monsters absorb": 		19,
				},
			},
			expUnique: 			newExpUnique(2, "lightning"),
			statusProtection: 	h.GetInt32Ptr(24),
			autoAbilities: 		[]int32{6, 58, 59, 60},
			playerAbilities: 	[]int32{70, 74, 78},
			overdriveAbilities: []int32{20, 57, 69, 135, 186},
			itemAbilities: 		[]int32{27, 28, 29},
			enemyAbilities: 	[]int32{90, 270, 401, 416},
			monstersWeak: 		[]int32{4, 18, 102, 132, 180, 227},
			monstersHalved: 	[]int32{30, 65, 140, 163, 247, 293},
			monstersImmune: 	[]int32{243, 274, 294, 297},
			monstersAbsorb: 	[]int32{27, 54, 77, 138, 205, 255, 295, 301},
		},
	}

	testSingleResources(t, tests, "GetElement", testCfg.HandleElements, compareElements)
}

func TestRetrieveElements(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/elements",
				expectedStatus: http.StatusOK,
			},
			count:   5,
			results: []int32{1, 2, 3, 4, 5},
		},
	}

	testIdList(t, tests, testCfg.e.elements.endpoint, "RetrieveElements", testCfg.HandleElements, compareAPIResourceLists[NamedApiResourceList])
}
