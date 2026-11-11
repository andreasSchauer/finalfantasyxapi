package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/sections/sectiontypes"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

func TestGetOverdrive(t *testing.T) {
	t.Parallel()
	tests := []Overdrive{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrives/125",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "overdrive with provided id '125' doesn't exist. max id: 124",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrives/15",
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"overdrive abilities": 7,
				},
			},
			Unique:          newExpUnique(15, "status reels"),
			rank:               h.GetInt32Ptr(4),
			countdownInSec:     h.GetInt32Ptr(20),
			user:               7,
			overdriveCommand:   h.GetInt32Ptr(3),
			overdriveAbilities: []int32{17, 32, 33, 34, 35, 36, 37},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrives/50",
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"overdrive abilities": 3,
				},
			},
			Unique:          newExpUnique(50, "banishing blade"),
			rank:               h.GetInt32Ptr(6),
			countdownInSec:     h.GetInt32Ptr(4),
			user:               10,
			overdriveCommand:   h.GetInt32Ptr(6),
			overdriveAbilities: []int32{113, 114, 115},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrives/64",
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"overdrive abilities": 1,
				},
			},
			Unique:          newExpUnique(64, "burning soul"),
			rank:               h.GetInt32Ptr(6),
			countdownInSec:     nil,
			user:               11,
			overdriveCommand:   h.GetInt32Ptr(7),
			overdriveAbilities: []int32{130},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrives/124",
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"overdrive abilities": 1,
				},
			},
			Unique:          newExpUnique(124, "delta attack"),
			rank:               h.GetInt32Ptr(6),
			countdownInSec:     nil,
			user:               4,
			overdriveCommand:   nil,
			overdriveAbilities: []int32{190},
		},
	}

	testSingleResources(t, tests, "GetOverdrive", testSrv.HandleOverdrives, compareOverdrives)
}

func TestRetrieveOverdrives(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrives?limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   124,
			results: []int32{1, 23, 45, 68, 77, 83, 84, 112, 124},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrives?user=lulu&rank=5",
				expectedStatus: http.StatusOK,
			},
			count:   14,
			results: []int32{17, 18, 21, 24, 26, 29, 31},
		},
	}

	testIDList(t, tests, apilookups.E.Overdrives.Endpoint, "RetrieveOverdrives", testSrv.HandleOverdrives, compareAPIResourceLists[apiresource.Named])
}

func TestSubsectionOverdrives(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-commands/1/overdrives",
				expectedStatus: http.StatusOK,
				handler:        testSrv.HandleOverdriveCommands,
			},
			count:          4,
			parentResource: h.GetStrPtr("/overdrive-commands/1"),
			results:        []int32{1, 2, 3, 4},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/character-classes/8/default-overdrives",
				expectedStatus: http.StatusOK,
				handler:        testSrv.HandleCharacterClasses,
			},
			count:          4,
			parentResource: h.GetStrPtr("/character-classes/8"),
			results:        []int32{17, 18, 19, 20},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/character-classes/7/learnable-overdrives",
				expectedStatus: http.StatusOK,
				handler:        testSrv.HandleCharacterClasses,
			},
			count:          3,
			parentResource: h.GetStrPtr("/character-classes/7"),
			results:        []int32{14, 15, 16},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/characters/6/overdrives",
				expectedStatus: http.StatusOK,
				handler:        testSrv.HandleCharacters,
			},
			count:          4,
			parentResource: h.GetStrPtr("/characters/6"),
			results:        []int32{48, 49, 50, 51},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/aeons/6/overdrives",
				expectedStatus: http.StatusOK,
				handler:        testSrv.HandleAeons,
			},
			count:          1,
			parentResource: h.GetStrPtr("/aeons/6"),
			results:        []int32{123},
		},
	}

	testIDList(t, tests, apilookups.E.Overdrives.Endpoint, "SubsectionOverdrives", nil, compareSimpleResourceLists[apiresource.Named, sectiontypes.OverdriveSimple])
}
