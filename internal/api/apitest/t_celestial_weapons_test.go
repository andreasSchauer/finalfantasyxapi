package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
)

func TestGetCelestialWeapon(t *testing.T) {
	t.Parallel()
	tests := []CelestialWeapon{
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/celestial-weapons/8",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "celestial weapon with provided id '8' doesn't exist. max id: 7",
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/celestial-weapons/3",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"auto-abilities": 4,
				},
			},
			Unique:     newExpUnique(3, "world champion"),
			equipment:     3,
			autoAbilities: []int32{51, 49, 44, 38},
			crest:         19,
			sigil:         20,
			wpnTreasure:   81,
			crestTreasure: 71,
			sigilQuest:    98,
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/celestial-weapons/6",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"auto-abilities": 4,
				},
			},
			Unique:     newExpUnique(6, "masamune"),
			equipment:     6,
			autoAbilities: []int32{51, 49, 3, 37},
			crest:         15,
			sigil:         16,
			wpnTreasure:   115,
			crestTreasure: 107,
			sigilQuest:    46,
		},
	}

	testSingleResources(t, tests, "GetCelestialWeapon", testSrv.HandleCelestialWeapons, compareCelestialWeapons)
}

func TestRetrieveCelestialWeapons(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/celestial-weapons",
				expectedStatus: http.StatusOK,
			},
			count:   7,
			results: []int32{1, 2, 3, 4, 5, 6, 7},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/celestial-weapons?formula=hp-high",
				expectedStatus: http.StatusOK,
			},
			count:   4,
			results: []int32{1, 3, 5, 7},
		},
	}

	testIDList(t, tests, apilookups.E.CelestialWeapons.Endpoint, "RetrieveCelestialWeapons", testSrv.HandleCelestialWeapons, compareAPIResourceLists[apiresource.Named])
}
