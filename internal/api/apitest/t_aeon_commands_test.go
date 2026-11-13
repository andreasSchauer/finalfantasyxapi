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
				requestURL:     "/api/aeon-commands/10",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "aeon command with provided id '10' doesn't exist. max id: 9",
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/aeon-commands/1",
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"possible abilities":               1,
					"possible abilities 0 - abilities": 6,
				},
			},
			Unique:   newExpUnique(1, "pay"),
			user:        19,
			topmenu:     h.GetInt32Ptr(1),
			openSubmenu: h.GetInt32Ptr(12),
			possibleAbilities: []PossibleAbilityList{
				{
					user:      19,
					abilities: []int32{94, 95, 96, 97, 98, 382},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/aeon-commands/2",
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"possible abilities":               3,
					"possible abilities 0 - abilities": 10,
					"possible abilities 1 - abilities": 3,
					"possible abilities 2 - abilities": 17,
				},
			},
			Unique:   newExpUnique(2, "do as you will."),
			user:        4,
			topmenu:     h.GetInt32Ptr(1),
			openSubmenu: nil,
			possibleAbilities: []PossibleAbilityList{
				{
					user:      20,
					abilities: []int32{375, 54, 55, 66, 80, 102},
				},
				{
					user:      21,
					abilities: []int32{375, 100, 62},
				},
				{
					user:      22,
					abilities: []int32{375, 101, 78, 84, 81, 102},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/aeon-commands/7",
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"possible abilities":               1,
					"possible abilities 0 - abilities": 10,
				},
			},
			Unique:   newExpUnique(7, "defense!"),
			user:        21,
			topmenu:     h.GetInt32Ptr(1),
			openSubmenu: nil,
			possibleAbilities: []PossibleAbilityList{
				{
					user:      21,
					abilities: []int32{375, 100, 45, 46, 61, 68, 102},
				},
			},
		},
	}

	testSingleResources(t, tests, "GetAeonCommand", testSrv.HandleAeonCommands, compareAeonCommands)
}

func TestRetrieveAeonCommands(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/aeon-commands",
				expectedStatus: http.StatusOK,
			},
			count:   9,
			results: []int32{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	testIDList(t, tests, apilookups.E.AeonCommands.Endpoint, "RetrieveAeonCommands", testSrv.HandleAeonCommands, compareAPIResourceLists[apiresource.Named])
}
