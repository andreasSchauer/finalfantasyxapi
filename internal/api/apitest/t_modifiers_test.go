package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
)

func TestGetModifier(t *testing.T) {
	t.Parallel()
	tests := []Modifier{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/modifiers/32",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "modifier with provided id '32' doesn't exist. max id: 31",
			},
		},
		{
			TestGeneral: TestGeneral{
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
			Unique:             newExpUnique(3, "buff-factor-str-based"),
			autoAbilities:      []int32{},
			playerAbilities:    []int32{27},
			overdriveAbilities: []int32{162, 163, 170},
			itemAbilities:      []int32{},
			triggerCommands:    []int32{},
			enemyAbilities:     []int32{},
			statusConditions:   []int32{},
			properties:         []int32{},
		},
		{
			TestGeneral: TestGeneral{
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
			Unique:             newExpUnique(16, "mp-cost"),
			autoAbilities:      []int32{40, 42, 43},
			playerAbilities:    []int32{},
			overdriveAbilities: []int32{174, 175},
			itemAbilities:      []int32{68, 69},
			triggerCommands:    []int32{},
			enemyAbilities:     []int32{},
			statusConditions:   []int32{},
			properties:         []int32{},
		},
		{
			TestGeneral: TestGeneral{
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
			Unique:             newExpUnique(18, "overdrive-gauge"),
			autoAbilities:      []int32{},
			playerAbilities:    []int32{},
			overdriveAbilities: []int32{},
			itemAbilities:      []int32{},
			triggerCommands:    []int32{6},
			enemyAbilities:     []int32{},
			statusConditions:   []int32{},
			properties:         []int32{},
		},
		{
			TestGeneral: TestGeneral{
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
			Unique:             newExpUnique(20, "physical-damage-taken"),
			autoAbilities:      []int32{103, 104, 105, 106},
			playerAbilities:    []int32{},
			overdriveAbilities: []int32{},
			itemAbilities:      []int32{},
			triggerCommands:    []int32{},
			enemyAbilities:     []int32{86},
			statusConditions:   []int32{9, 26, 31, 33, 43},
			properties:         []int32{1},
		},
	}

	testSingleResources(t, tests, "GetModifier", testSrv.HandleModifiers)
}

func TestRetrieveModifiers(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/modifiers?limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   31,
			Results: []int32{1, 7, 12, 17, 20, 25, 31},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/modifiers?category=percentage",
				ExpectedStatus: http.StatusOK,
			},
			Count:   9,
			Results: []int32{5, 7, 9, 25, 26, 29},
		},
	}

	testIDList(t, tests, apilookups.E.Modifiers.Endpoint, "RetrieveModifiers", testSrv.HandleModifiers, compareAPIResourceLists[apiresource.Named])
}
