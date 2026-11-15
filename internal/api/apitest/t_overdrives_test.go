package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apitesttypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/sections/sectiontypes"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

func TestGetOverdrive(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.Overdrive{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/overdrives/125",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "overdrive with provided id '125' doesn't exist. max id: 124",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/overdrives/15",
				ExpectedStatus: http.StatusOK,
				ExpLengths: map[string]int{
					"overdrive abilities": 7,
				},
			},
			Unique:             apitesttypes.NewExpUnique(15, "status reels"),
			Rank:               h.GetInt32Ptr(4),
			CountdownInSec:     h.GetInt32Ptr(20),
			User:               7,
			OverdriveCommand:   h.GetInt32Ptr(3),
			OverdriveAbilities: []int32{17, 32, 33, 34, 35, 36, 37},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/overdrives/50",
				ExpectedStatus: http.StatusOK,
				ExpLengths: map[string]int{
					"overdrive abilities": 3,
				},
			},
			Unique:             apitesttypes.NewExpUnique(50, "banishing blade"),
			Rank:               h.GetInt32Ptr(6),
			CountdownInSec:     h.GetInt32Ptr(4),
			User:               10,
			OverdriveCommand:   h.GetInt32Ptr(6),
			OverdriveAbilities: []int32{113, 114, 115},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/overdrives/64",
				ExpectedStatus: http.StatusOK,
				ExpLengths: map[string]int{
					"overdrive abilities": 1,
				},
			},
			Unique:             apitesttypes.NewExpUnique(64, "burning soul"),
			Rank:               h.GetInt32Ptr(6),
			CountdownInSec:     nil,
			User:               11,
			OverdriveCommand:   h.GetInt32Ptr(7),
			OverdriveAbilities: []int32{130},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/overdrives/124",
				ExpectedStatus: http.StatusOK,
				ExpLengths: map[string]int{
					"overdrive abilities": 1,
				},
			},
			Unique:             apitesttypes.NewExpUnique(124, "delta attack"),
			Rank:               h.GetInt32Ptr(6),
			CountdownInSec:     nil,
			User:               4,
			OverdriveCommand:   nil,
			OverdriveAbilities: []int32{190},
		},
	}

	testSingleResources(t, tests, "GetOverdrive", testSrv.HandleOverdrives)
}

func TestRetrieveOverdrives(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.ListIDs{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/overdrives?limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   124,
			Results: []int32{1, 23, 45, 68, 77, 83, 84, 112, 124},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/overdrives?user=lulu&rank=5",
				ExpectedStatus: http.StatusOK,
			},
			Count:   14,
			Results: []int32{17, 18, 21, 24, 26, 29, 31},
		},
	}

	testIDList(t, tests, apilookups.E.Overdrives.Endpoint, "RetrieveOverdrives", testSrv.HandleOverdrives, compareAPIResourceLists[apiresource.Named])
}

func TestSubsectionOverdrives(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.ListIDs{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/overdrive-commands/1/overdrives",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleOverdriveCommands,
			},
			Count:          4,
			ParentResource: h.GetStrPtr("/overdrive-commands/1"),
			Results:        []int32{1, 2, 3, 4},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/character-classes/8/default-overdrives",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleCharacterClasses,
			},
			Count:          4,
			ParentResource: h.GetStrPtr("/character-classes/8"),
			Results:        []int32{17, 18, 19, 20},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/character-classes/7/learnable-overdrives",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleCharacterClasses,
			},
			Count:          3,
			ParentResource: h.GetStrPtr("/character-classes/7"),
			Results:        []int32{14, 15, 16},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/characters/6/overdrives",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleCharacters,
			},
			Count:          4,
			ParentResource: h.GetStrPtr("/characters/6"),
			Results:        []int32{48, 49, 50, 51},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/aeons/6/overdrives",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleAeons,
			},
			Count:          1,
			ParentResource: h.GetStrPtr("/aeons/6"),
			Results:        []int32{123},
		},
	}

	testIDList(t, tests, apilookups.E.Overdrives.Endpoint, "SubsectionOverdrives", nil, compareSimpleResourceLists[apiresource.Named, sectiontypes.OverdriveSimple])
}
