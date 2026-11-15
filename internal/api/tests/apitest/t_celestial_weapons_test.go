package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/tests/apitesttypes"
)

func TestGetCelestialWeapon(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.CelestialWeapon{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/celestial-weapons/8",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "celestial weapon with provided id '8' doesn't exist. max id: 7",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/celestial-weapons/3",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"auto-abilities": 4,
				},
			},
			Unique:        apitesttypes.NewExpUnique(3, "world champion"),
			Equipment:     3,
			AutoAbilities: []int32{51, 49, 44, 38},
			Crest:         19,
			Sigil:         20,
			WpnTreasure:   81,
			CrestTreasure: 71,
			SigilQuest:    98,
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/celestial-weapons/6",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"auto-abilities": 4,
				},
			},
			Unique:        apitesttypes.NewExpUnique(6, "masamune"),
			Equipment:     6,
			AutoAbilities: []int32{51, 49, 3, 37},
			Crest:         15,
			Sigil:         16,
			WpnTreasure:   115,
			CrestTreasure: 107,
			SigilQuest:    46,
		},
	}

	testSingleResources(t, tests, "GetCelestialWeapon", testSrv.HandleCelestialWeapons)
}

func TestRetrieveCelestialWeapons(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.ListIDs{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/celestial-weapons",
				ExpectedStatus: http.StatusOK,
			},
			Count:   7,
			Results: []int32{1, 2, 3, 4, 5, 6, 7},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/celestial-weapons?formula=hp-high",
				ExpectedStatus: http.StatusOK,
			},
			Count:   4,
			Results: []int32{1, 3, 5, 7},
		},
	}

	testIDList(t, tests, apilookups.E.CelestialWeapons.Endpoint, "RetrieveCelestialWeapons", testSrv.HandleCelestialWeapons, compareAPIResourceLists[apiresource.Named])
}
