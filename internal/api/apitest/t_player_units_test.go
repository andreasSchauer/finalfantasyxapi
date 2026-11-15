package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apitesttypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

func TestGetPlayerUnit(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.PlayerUnit{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/player-units/19",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "player unit with provided id '19' doesn't exist. max id: 18",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/player-units/1",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"character classes": 2,
				},
			},
			Unique: apitesttypes.Unique{
				ID:   1,
				Name: "tidus",
			},
			UnitType:         database.UnitTypeCharacter,
			TypedUnit:        "/characters/1",
			Area:             1,
			CelestialWeapon:  h.GetInt32Ptr(1),
			CharacterClasses: []int32{1, 5},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/player-units/8",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"character classes": 1,
				},
			},
			Unique: apitesttypes.Unique{
				ID:   8,
				Name: "seymour",
			},
			UnitType:         database.UnitTypeCharacter,
			TypedUnit:        "/characters/8",
			Area:             102,
			CelestialWeapon:  nil,
			CharacterClasses: []int32{12},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/player-units/10",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"character classes": 3,
				},
			},
			Unique: apitesttypes.Unique{
				ID:   10,
				Name: "ifrit",
			},
			UnitType:         database.UnitTypeAeon,
			TypedUnit:        "/aeons/2",
			Area:             60,
			CelestialWeapon:  h.GetInt32Ptr(3),
			CharacterClasses: []int32{2, 3, 14},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/player-units/15",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"character classes": 2,
				},
			},
			Unique: apitesttypes.Unique{
				ID:   15,
				Name: "yojimbo",
			},
			UnitType:         database.UnitTypeAeon,
			TypedUnit:        "/aeons/7",
			Area:             210,
			CelestialWeapon:  h.GetInt32Ptr(6),
			CharacterClasses: []int32{2, 19},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/player-units/16",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"character classes": 3,
				},
			},
			Unique: apitesttypes.Unique{
				ID:   16,
				Name: "cindy",
			},
			UnitType:         database.UnitTypeAeon,
			TypedUnit:        "/aeons/8",
			Area:             208,
			CelestialWeapon:  nil,
			CharacterClasses: []int32{2, 4, 20},
		},
	}

	testSingleResources(t, tests, "GetPlayerUnits", testSrv.HandlePlayerUnits)
}

func TestRetrievePlayerUnits(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.ListIDs{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/player-units",
				ExpectedStatus: http.StatusOK,
			},
			Count:   18,
			Results: []int32{1, 3, 4, 8, 12, 17, 18},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/player-units?type=aeon",
				ExpectedStatus: http.StatusOK,
			},
			Count:   10,
			Results: []int32{9, 11, 13, 14, 15, 16, 18},
		},
	}

	testIDList(t, tests, apilookups.E.PlayerUnits.Endpoint, "RetrievePlayerUnits", testSrv.HandlePlayerUnits, compareAPIResourceLists[apiresource.Typed])
}
