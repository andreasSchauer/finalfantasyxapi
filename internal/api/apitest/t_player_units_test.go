package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

func TestGetPlayerUnit(t *testing.T) {
	t.Parallel()
	tests := []PlayerUnit{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/player-units/19",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "player unit with provided id '19' doesn't exist. max id: 18",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/player-units/1",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"character classes": 2,
				},
			},
			Unique: Unique{
				ID:   1,
				Name: "tidus",
			},
			unitType:         database.UnitTypeCharacter,
			typedUnit:        "/characters/1",
			area:             1,
			celestialWeapon:  h.GetInt32Ptr(1),
			characterClasses: []int32{1, 5},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/player-units/8",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"character classes": 1,
				},
			},
			Unique: Unique{
				ID:   8,
				Name: "seymour",
			},
			unitType:         database.UnitTypeCharacter,
			typedUnit:        "/characters/8",
			area:             102,
			celestialWeapon:  nil,
			characterClasses: []int32{12},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/player-units/10",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"character classes": 3,
				},
			},
			Unique: Unique{
				ID:   10,
				Name: "ifrit",
			},
			unitType:         database.UnitTypeAeon,
			typedUnit:        "/aeons/2",
			area:             60,
			celestialWeapon:  h.GetInt32Ptr(3),
			characterClasses: []int32{2, 3, 14},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/player-units/15",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"character classes": 2,
				},
			},
			Unique: Unique{
				ID:   15,
				Name: "yojimbo",
			},
			unitType:         database.UnitTypeAeon,
			typedUnit:        "/aeons/7",
			area:             210,
			celestialWeapon:  h.GetInt32Ptr(6),
			characterClasses: []int32{2, 19},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/player-units/16",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"character classes": 3,
				},
			},
			Unique: Unique{
				ID:   16,
				Name: "cindy",
			},
			unitType:         database.UnitTypeAeon,
			typedUnit:        "/aeons/8",
			area:             208,
			celestialWeapon:  nil,
			characterClasses: []int32{2, 4, 20},
		},
	}

	testSingleResources(t, tests, "GetPlayerUnits", testSrv.HandlePlayerUnits)
}

func TestRetrievePlayerUnits(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/player-units",
				ExpectedStatus: http.StatusOK,
			},
			Count:   18,
			Results: []int32{1, 3, 4, 8, 12, 17, 18},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/player-units?type=aeon",
				ExpectedStatus: http.StatusOK,
			},
			Count:   10,
			Results: []int32{9, 11, 13, 14, 15, 16, 18},
		},
	}

	testIDList(t, tests, apilookups.E.PlayerUnits.Endpoint, "RetrievePlayerUnits", testSrv.HandlePlayerUnits, compareAPIResourceLists[apiresource.Typed])
}
