package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

func TestGetAgilityTier(t *testing.T) {
	t.Parallel()
	tests := []AgilityTier{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/agility-tiers/20",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "agility tier with provided id '20' doesn't exist. max id: 19",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/agility-tiers/15",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"char min icvs": 9,
				},
			},
			IDOnly:      newExpIDOnly(15),
			FromAgility: 35,
			ToAgility:   43,
			TickSpeed:   7,
			MonMaxICV:   h.GetInt32Ptr(23),
			MonMinICV:   h.GetInt32Ptr(21),
			CharMaxICV:  h.GetInt32Ptr(21),
			CharMinICVs: []apitypes.AgilitySubtier{
				{
					FromAgility: 35,
					ToAgility:   35,
					CharMinICV:  h.GetInt32Ptr(20),
				},
				{
					FromAgility: 36,
					ToAgility:   36,
					CharMinICV:  h.GetInt32Ptr(19),
				},
				{
					FromAgility: 37,
					ToAgility:   37,
					CharMinICV:  h.GetInt32Ptr(18),
				},
				{
					FromAgility: 38,
					ToAgility:   38,
					CharMinICV:  h.GetInt32Ptr(17),
				},
				{
					FromAgility: 39,
					ToAgility:   39,
					CharMinICV:  h.GetInt32Ptr(16),
				},
				{
					FromAgility: 40,
					ToAgility:   40,
					CharMinICV:  h.GetInt32Ptr(15),
				},
				{
					FromAgility: 41,
					ToAgility:   41,
					CharMinICV:  h.GetInt32Ptr(14),
				},
				{
					FromAgility: 42,
					ToAgility:   42,
					CharMinICV:  h.GetInt32Ptr(13),
				},
				{
					FromAgility: 43,
					ToAgility:   43,
					CharMinICV:  h.GetInt32Ptr(12),
				},
			},
		},
	}

	testSingleResources(t, tests, "GetAgilityTier", testSrv.HandleAgilityTiers)
}

func TestRetrieveAgilityTiers(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/agility-tiers",
				ExpectedStatus: http.StatusOK,
			},
			Count:   19,
			Results: []int32{1, 3, 7, 8, 12, 17, 19},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/agility-tiers?agility=95",
				ExpectedStatus: http.StatusOK,
			},
			Count:   1,
			Results: []int32{17},
		},
	}

	testIDList(t, tests, apilookups.E.AgilityTiers.Endpoint, "RetrieveAgilityTiers", testSrv.HandleAgilityTiers, compareAPIResourceLists[apiresource.Named])
}
