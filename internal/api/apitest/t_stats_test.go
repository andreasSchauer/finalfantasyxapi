package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apitesttypes"
)

func TestGetStat(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.Stat{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/stats/11",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "stat with provided id '11' doesn't exist. max id: 10",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/stats/2?changes_only=true",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"spheres":             5,
					"auto-abilities":      4,
					"player abilities":    0,
					"overdrive abilities": 3,
					"item abilities":      2,
					"trigger commands":    0,
					"status conditions":   0,
				},
			},
			Unique:             apitesttypes.NewExpUnique(2, "mp"),
			Spheres:            []int32{2, 6, 11, 17, 26},
			AutoAbilities:      []int32{115, 116, 117, 118},
			PlayerAbilities:    []int32{},
			OverdriveAbilities: []int32{171, 172, 173},
			ItemAbilities:      []int32{65, 67},
			TriggerCommands:    []int32{},
			StatusConditions:   []int32{},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/stats/3",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"spheres":             5,
					"auto-abilities":      5,
					"player abilities":    2,
					"overdrive abilities": 3,
					"item abilities":      0,
					"trigger commands":    1,
					"status conditions":   3,
					"properties":          0,
				},
			},
			Unique:             apitesttypes.NewExpUnique(3, "strength"),
			Spheres:            []int32{1, 6, 11, 18, 26},
			AutoAbilities:      []int32{29, 30, 31, 32, 51},
			PlayerAbilities:    []int32{11, 27},
			OverdriveAbilities: []int32{162, 163, 170},
			ItemAbilities:      []int32{},
			TriggerCommands:    []int32{4},
			StatusConditions:   []int32{1, 17, 34},
			Properties:         []int32{},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/stats/4",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"spheres":             true,
					"auto-abilities":      true,
					"player abilities":    true,
					"overdrive abilities": true,
					"item abilities":      true,
					"trigger commands":    true,
					"status conditions":   true,
				},
				ExpLengths: map[string]int{
					"properties": 2,
				},
			},
			Unique:     apitesttypes.NewExpUnique(4, "defense"),
			Properties: []int32{1, 4},
		},
	}

	testSingleResources(t, tests, "GetStat", testSrv.HandleStats)
}

func TestRetrieveStats(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.ListIDs{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/stats",
				ExpectedStatus: http.StatusOK,
			},
			Count:   10,
			Results: []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	testIDList(t, tests, apilookups.E.Stats.Endpoint, "RetrieveStats", testSrv.HandleStats, compareAPIResourceLists[apiresource.Named])
}
