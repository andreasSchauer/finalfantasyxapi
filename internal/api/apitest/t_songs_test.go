package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/sections/sectiontypes"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

func TestGetSong(t *testing.T) {
	t.Parallel()
	tests := []Song{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/songs/a",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "song not found: 'a'",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/songs/a/2",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "wrong format. usage: '/api/songs', '/api/songs/{id}', '/api/songs/{name}'",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/songs/2/a",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "endpoint /songs doesn't have any subsections",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/songs/44",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"background music": 1,
					"cues":             3,
					"boss fights":      0,
					"fmvs":             0,
				},
			},
			Unique:            newExpUnique(44, "twilight"),
			Composer:          h.GetStrPtr("junya nakano"),
			Arranger:          h.GetStrPtr("junya nakano"),
			DurationInSeconds: 280,
			CanLoop:           true,
			BackgroundMusic: []testBackgroundMusic{
				{
					Index:                  0,
					ReplacesEncounterMusic: false,
					Areas:                  []int32{232, 233, 234},
				},
			},
			Cues: []testCue{
				{
					Index:                  0,
					TriggerArea:            h.GetInt32Ptr(103),
					IncludedAreas:          []int32{},
					ReplacesEncounterMusic: false,
					ReplacesBGMusic:        nil,
				},
				{
					Index:                  1,
					TriggerArea:            h.GetInt32Ptr(158),
					IncludedAreas:          []int32{},
					ReplacesEncounterMusic: false,
					ReplacesBGMusic:        nil,
				},
				{
					Index:                  2,
					TriggerArea:            h.GetInt32Ptr(214),
					IncludedAreas:          []int32{214},
					ReplacesEncounterMusic: false,
					ReplacesBGMusic:        h.GetStrPtr("until-trigger"),
				},
			},
			BossFights: []int32{},
			Fmvs:       []int32{},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/songs/a_fLeeting_dream",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"background music": 1,
					"cues":             6,
					"boss fights":      0,
					"fmvs":             1,
				},
			},
			Unique:            newExpUnique(78, "a fleeting dream"),
			Composer:          h.GetStrPtr("nobuo uematsu"),
			Arranger:          h.GetStrPtr("masashi hamauzu"),
			DurationInSeconds: 264,
			CanLoop:           true,
			BackgroundMusic: []testBackgroundMusic{
				{
					Index:                  0,
					ReplacesEncounterMusic: true,
					Areas:                  []int32{219, 220},
				},
			},
			Cues: []testCue{
				{
					Index:                  0,
					TriggerArea:            h.GetInt32Ptr(42),
					IncludedAreas:          []int32{},
					ReplacesEncounterMusic: false,
					ReplacesBGMusic:        nil,
				},
				{
					Index:                  2,
					TriggerArea:            h.GetInt32Ptr(187),
					IncludedAreas:          []int32{179, 180, 181, 182, 183, 184, 185, 186, 187, 188, 189, 190},
					ReplacesEncounterMusic: false,
					ReplacesBGMusic:        h.GetStrPtr("until-trigger"),
				},
				{
					Index:                  5,
					TriggerArea:            nil,
					IncludedAreas:          []int32{},
					ReplacesEncounterMusic: false,
					ReplacesBGMusic:        nil,
				},
			},
			BossFights: []int32{},
			Fmvs:       []int32{51},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/songs/80",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"background music": true,
					"cues":             true,
				},
				ExpLengths: map[string]int{
					"background music": 0,
					"cues":             0,
					"boss fights":      5,
					"fmvs":             1,
				},
			},
			Unique:            newExpUnique(80, "challenge"),
			Composer:          h.GetStrPtr("masashi hamauzu"),
			Arranger:          h.GetStrPtr("masashi hamauzu"),
			DurationInSeconds: 258,
			CanLoop:           true,
			BossFights:        []int32{85, 235, 254, 255, 296},
			Fmvs:              []int32{22},
		},
	}

	testSingleResources(t, tests, "GetSongs", testSrv.HandleSongs)
}

func TestRetrieveSongs(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/songs?limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   95,
			Results: []int32{1, 15, 18, 33, 47, 56, 71, 80, 93, 95},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/songs?location=13",
				ExpectedStatus: http.StatusOK,
			},
			Count:   4,
			Results: []int32{20, 43, 49, 77},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/songs?sublocation=15", // luca
				ExpectedStatus: http.StatusOK,
			},
			Count:   12,
			Results: []int32{10, 16, 17, 27, 32, 34, 75, 77},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/songs?area=211", // gagazet - mountain gate
				ExpectedStatus: http.StatusOK,
			},
			Count:   3,
			Results: []int32{9, 27, 75},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/songs?composer=nobuo-uematsu&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   53,
			Results: []int32{2, 10, 19, 40, 42, 67, 78, 91, 95},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/songs?arranger=3&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   34,
			Results: []int32{18, 41, 46, 61, 64, 76, 88},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/songs?special_use=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   8,
			Results: []int32{9, 10, 11, 16, 24, 34, 37, 40},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/songs?fmvs=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   20,
			Results: []int32{5, 6, 17, 29, 61, 63, 64, 70, 89},
		},
	}

	testIDList(t, tests, apilookups.E.Songs.Endpoint, "RetrieveSongs", testSrv.HandleSongs, compareAPIResourceLists[apiresource.Named])
}

func TestSubsectionSongs(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/locations/25/songs",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleLocations,
			},
			Count:          11,
			ParentResource: h.GetStrPtr("/locations/25"),
			Results:        []int32{5, 44, 57, 60, 78, 79, 84, 85, 87, 88, 89},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sublocations/31/songs",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleSublocations,
			},
			Count:          4,
			ParentResource: h.GetStrPtr("/sublocations/31"),
			Results:        []int32{2, 27, 59, 60},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas/148/songs",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleAreas,
			},
			Count:          4,
			ParentResource: h.GetStrPtr("/areas/148"),
			Results:        []int32{4, 16, 30, 51},
		},
	}

	testIDList(t, tests, apilookups.E.Songs.Endpoint, "SubsectionSongs", nil, compareSimpleResourceLists[apiresource.Named, sectiontypes.SongSimple])
}
