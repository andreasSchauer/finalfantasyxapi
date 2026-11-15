package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/tests/apitesttypes"
)

func TestGetModifier(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.Modifier{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/modifiers/32",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "modifier with provided id '32' doesn't exist. max id: 31",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/modifiers/3",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"auto-abilities":      0,
					"player abilities":    1,
					"overdrive abilities": 3,
					"item abilities":      0,
					"trigger commands":    0,
					"enemy abilities":     0,
					"status conditions":   0,
					"properties":          0,
				},
			},
			Unique:             apitesttypes.NewExpUnique(3, "buff-factor-str-based"),
			AutoAbilities:      []int32{},
			PlayerAbilities:    []int32{27},
			OverdriveAbilities: []int32{162, 163, 170},
			ItemAbilities:      []int32{},
			TriggerCommands:    []int32{},
			EnemyAbilities:     []int32{},
			StatusConditions:   []int32{},
			Properties:         []int32{},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/modifiers/16",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"auto-abilities":      3,
					"player abilities":    0,
					"overdrive abilities": 2,
					"item abilities":      2,
					"trigger commands":    0,
					"enemy abilities":     0,
					"status conditions":   0,
					"properties":          0,
				},
			},
			Unique:             apitesttypes.NewExpUnique(16, "mp-cost"),
			AutoAbilities:      []int32{40, 42, 43},
			PlayerAbilities:    []int32{},
			OverdriveAbilities: []int32{174, 175},
			ItemAbilities:      []int32{68, 69},
			TriggerCommands:    []int32{},
			EnemyAbilities:     []int32{},
			StatusConditions:   []int32{},
			Properties:         []int32{},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/modifiers/18",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"auto-abilities":      0,
					"player abilities":    0,
					"overdrive abilities": 0,
					"item abilities":      0,
					"trigger commands":    1,
					"enemy abilities":     0,
					"status conditions":   0,
					"properties":          0,
				},
			},
			Unique:             apitesttypes.NewExpUnique(18, "overdrive-gauge"),
			AutoAbilities:      []int32{},
			PlayerAbilities:    []int32{},
			OverdriveAbilities: []int32{},
			ItemAbilities:      []int32{},
			TriggerCommands:    []int32{6},
			EnemyAbilities:     []int32{},
			StatusConditions:   []int32{},
			Properties:         []int32{},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/modifiers/20",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"auto-abilities":      4,
					"player abilities":    0,
					"overdrive abilities": 0,
					"item abilities":      0,
					"trigger commands":    0,
					"enemy abilities":     1,
					"status conditions":   6,
					"properties":          1,
				},
			},
			Unique:             apitesttypes.NewExpUnique(20, "physical-damage-taken"),
			AutoAbilities:      []int32{103, 104, 105, 106},
			PlayerAbilities:    []int32{},
			OverdriveAbilities: []int32{},
			ItemAbilities:      []int32{},
			TriggerCommands:    []int32{},
			EnemyAbilities:     []int32{86},
			StatusConditions:   []int32{9, 26, 31, 33, 43},
			Properties:         []int32{1},
		},
	}

	testSingleResources(t, tests, "GetModifier", testSrv.HandleModifiers)
}

func TestRetrieveModifiers(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.ListIDs{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/modifiers?limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   31,
			Results: []int32{1, 7, 12, 17, 20, 25, 31},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/modifiers?category=percentage",
				ExpectedStatus: http.StatusOK,
			},
			Count:   9,
			Results: []int32{5, 7, 9, 25, 26, 29},
		},
	}

	testIDList(t, tests, apilookups.E.Modifiers.Endpoint, "RetrieveModifiers", testSrv.HandleModifiers, compareAPIResourceLists[apiresource.Named])
}
