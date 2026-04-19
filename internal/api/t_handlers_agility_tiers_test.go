package api

import (
	"net/http"
	"testing"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func TestGetAgilityTier(t *testing.T) {
	tests := []expAgilityTier{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/agility-tiers/20",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "agility tier with provided id '20' doesn't exist. max id: 19.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/agility-tiers/15",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"char min icvs": 	9,
				},
			},
			expIdOnly: 		newExpIdOnly(15),
			fromAgility: 	35,
			toAgility: 		43,
			tickSpeed: 		7,
			monMaxICV: 		h.GetInt32Ptr(23),
			monMinICV: 		h.GetInt32Ptr(21),
			charMaxICV: 	h.GetInt32Ptr(21),
			charMinICVs: 	[]AgilitySubtier{
				{
					FromAgility: 	35,
					ToAgility: 		35,
					CharMinICV: 	h.GetInt32Ptr(20),
				},
				{
					FromAgility: 	36,
					ToAgility: 		36,
					CharMinICV: 	h.GetInt32Ptr(19),
				},
				{
					FromAgility: 	37,
					ToAgility: 		37,
					CharMinICV: 	h.GetInt32Ptr(18),
				},
				{
					FromAgility: 	38,
					ToAgility: 		38,
					CharMinICV: 	h.GetInt32Ptr(17),
				},
				{
					FromAgility: 	39,
					ToAgility: 		39,
					CharMinICV: 	h.GetInt32Ptr(16),
				},
				{
					FromAgility: 	40,
					ToAgility: 		40,
					CharMinICV: 	h.GetInt32Ptr(15),
				},
				{
					FromAgility: 	41,
					ToAgility: 		41,
					CharMinICV: 	h.GetInt32Ptr(14),
				},
				{
					FromAgility: 	42,
					ToAgility: 		42,
					CharMinICV: 	h.GetInt32Ptr(13),
				},
				{
					FromAgility: 	43,
					ToAgility: 		43,
					CharMinICV: 	h.GetInt32Ptr(12),
				},
			},
		},
	}

	testSingleResources(t, tests, "GetAgilityTier", testCfg.HandleAgilityTiers, compareAgilityTiers)
}

func TestRetrieveAgilityTiers(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/agility-tiers",
				expectedStatus: http.StatusOK,
			},
			count:   19,
			results: []int32{1, 3, 7, 8, 12, 17, 19},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/agility-tiers?agility=95",
				expectedStatus: http.StatusOK,
			},
			count:   1,
			results: []int32{17},
		},
	}

	testIdList(t, tests, testCfg.e.agilityTiers.endpoint, "RetrieveAgilityTiers", testCfg.HandleAgilityTiers, compareAPIResourceLists[UnnamedApiResourceList])
}
