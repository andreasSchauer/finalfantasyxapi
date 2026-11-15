package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

func TestGetAeonCommand(t *testing.T) {
	t.Parallel()
	tests := []AeonCommand{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/aeon-commands/10",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "aeon command with provided id '10' doesn't exist. max id: 9",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/aeon-commands/1",
				ExpectedStatus: http.StatusOK,
				ExpLengths: map[string]int{
					"possible abilities":               1,
					"possible abilities 0 - abilities": 6,
				},
			},
			Unique:      NewExpUnique(1, "pay"),
			User:        19,
			Topmenu:     h.GetInt32Ptr(1),
			OpenSubmenu: h.GetInt32Ptr(12),
			PossibleAbilities: []PossibleAbilityList{
				{
					User:      19,
					Abilities: []int32{94, 95, 96, 97, 98, 382},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/aeon-commands/2",
				ExpectedStatus: http.StatusOK,
				ExpLengths: map[string]int{
					"possible abilities":               3,
					"possible abilities 0 - abilities": 10,
					"possible abilities 1 - abilities": 3,
					"possible abilities 2 - abilities": 17,
				},
			},
			Unique:      NewExpUnique(2, "do as you will."),
			User:        4,
			Topmenu:     h.GetInt32Ptr(1),
			OpenSubmenu: nil,
			PossibleAbilities: []PossibleAbilityList{
				{
					User:      20,
					Abilities: []int32{375, 54, 55, 66, 80, 102},
				},
				{
					User:      21,
					Abilities: []int32{375, 100, 62},
				},
				{
					User:      22,
					Abilities: []int32{375, 101, 78, 84, 81, 102},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/aeon-commands/7",
				ExpectedStatus: http.StatusOK,
				ExpLengths: map[string]int{
					"possible abilities":               1,
					"possible abilities 0 - abilities": 10,
				},
			},
			Unique:      NewExpUnique(7, "defense!"),
			User:        21,
			Topmenu:     h.GetInt32Ptr(1),
			OpenSubmenu: nil,
			PossibleAbilities: []PossibleAbilityList{
				{
					User:      21,
					Abilities: []int32{375, 100, 45, 46, 61, 68, 102},
				},
			},
		},
	}

	testSingleResources(t, tests, "GetAeonCommand", testSrv.HandleAeonCommands)
}

func TestRetrieveAeonCommands(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/aeon-commands",
				ExpectedStatus: http.StatusOK,
			},
			Count:   9,
			Results: []int32{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	testIDList(t, tests, apilookups.E.AeonCommands.Endpoint, "RetrieveAeonCommands", testSrv.HandleAeonCommands, compareAPIResourceLists[apiresource.Named])
}
