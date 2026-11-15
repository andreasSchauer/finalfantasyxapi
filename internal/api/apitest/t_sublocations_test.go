package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/sections/sectiontypes"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

func TestGetSublocation(t *testing.T) {
	t.Parallel()
	tests := []Sublocation{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sublocations/0",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "sublocation with provided id '0' doesn't exist. max id: 43",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sublocations/44",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "sublocation with provided id '44' doesn't exist. max id: 43",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sublocations/a",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "sublocation not found: 'a'",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sublocations/35/",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"characters": true,
					"aeons":      true,
					"cues music": true,
					"fmvs music": true,
					"fmvs":       true,
				},
				ExpLengths: map[string]int{
					"connected sublocations": 4,
					"areas":                  6,
					"characters":             0,
					"aeons":                  0,
					"shops":                  3,
					"treasures":              9,
					"monsters":               51,
					"formations":             50,
					"quests":                 46,
					"bg music":               2,
					"cues music":             1,
					"fmvs music":             0,
					"boss music":             1,
					"fmvs":                   0,
				},
			},
			Unique:                newExpUnique(35, "calm lands"),
			ParentLocation:        20,
			ConnectedSublocations: []int32{26, 36, 38, 39},
			Areas:                 []int32{200, 201, 202, 203, 204, 205},
			LocRel: LocRel{
				Shops:      []int32{31, 32, 33},
				Treasures:  []int32{265, 268, 270, 272},
				Monsters:   []int32{138, 142, 149, 259, 270, 282, 292},
				Formations: []int32{193, 198, 205, 306, 312, 320, 331},
				Quests:     []int32{1, 11, 19, 34, 42, 45, 46, 69, 74, 77},
				Music: &testLocMusic{
					BgMusic:   []int32{71, 73},
					BossMusic: []int32{16},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sublocations/26?rel_availability=pre-story",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"music": true,
					"fmvs":  true,
				},
				ExpLengths: map[string]int{
					"connected sublocations": 4,
					"areas":                  13,
					"characters":             0,
					"aeons":                  0,
					"shops":                  1,
					"treasures":              0,
					"monsters":               1,
					"formations":             1,
					"quests":                 0,
				},
			},
			Unique:                newExpUnique(26, "macalania woods"),
			ParentLocation:        15,
			ConnectedSublocations: []int32{25, 27, 34, 35},
			Areas:                 []int32{141, 143, 146, 149, 153},
			LocRel: LocRel{
				Characters: []int32{},
				Aeons:      []int32{},
				Shops:      []int32{22},
				Treasures:  []int32{},
				Monsters:   []int32{86},
				Formations: []int32{126},
				Quests:     []int32{},
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sublocations/26?rel_availability=always&rel_repeatable=true",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"music": true,
					"fmvs":  true,
				},
				ExpLengths: map[string]int{
					"connected sublocations": 4,
					"areas":                  13,
					"characters":             0,
					"aeons":                  0,
					"shops":                  0,
					"treasures":              0,
					"monsters":               6,
					"formations":             6,
					"quests":                 0,
				},
			},
			Unique:                newExpUnique(26, "macalania woods"),
			ParentLocation:        15,
			ConnectedSublocations: []int32{25, 27, 34, 35},
			Areas:                 []int32{141, 143, 146, 149, 153},
			LocRel: LocRel{
				Characters: []int32{},
				Aeons:      []int32{},
				Shops:      []int32{},
				Treasures:  []int32{},
				Monsters:   []int32{80, 81, 82, 83, 84, 85},
				Formations: []int32{120, 121, 122, 123, 124, 125},
				Quests:     []int32{},
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sublocations/26?rel_availability=always&rel_repeatable=false",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"music": true,
					"fmvs":  true,
				},
				ExpLengths: map[string]int{
					"connected sublocations": 4,
					"areas":                  13,
					"characters":             0,
					"aeons":                  0,
					"shops":                  0,
					"treasures":              8,
					"monsters":               0,
					"formations":             0,
					"quests":                 5,
				},
			},
			Unique:                newExpUnique(26, "macalania woods"),
			ParentLocation:        15,
			ConnectedSublocations: []int32{25, 27, 34, 35},
			Areas:                 []int32{141, 143, 146, 149, 153},
			LocRel: LocRel{
				Characters: []int32{},
				Aeons:      []int32{},
				Shops:      []int32{},
				Treasures:  []int32{187, 188, 189, 190, 191, 192, 193, 194},
				Monsters:   []int32{},
				Formations: []int32{},
				Quests:     []int32{10, 82, 83, 84, 85},
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sublocations/14/",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"connected sublocations": 1,
					"areas":                  7,
					"characters":             1,
					"aeons":                  0,
					"shops":                  1,
					"treasures":              5,
					"monsters":               3,
					"formations":             3,
					"quests":                 2,
					"bg music":               3,
					"cues music":             3,
					"fmvs music":             2,
					"boss music":             0,
					"fmvs":                   2,
				},
			},
			Unique:                newExpUnique(14, "stadium"),
			ParentLocation:        8,
			ConnectedSublocations: []int32{15},
			Areas:                 []int32{68, 69, 70, 71, 72, 73, 74},
			LocRel: LocRel{
				Characters: []int32{6},
				Aeons:      []int32{},
				Shops:      []int32{5},
				Treasures:  []int32{70, 71, 72, 73, 74},
				Monsters:   []int32{34, 35, 36},
				Formations: []int32{38, 39, 40},
				Quests:     []int32{97, 98},
				Music: &testLocMusic{
					BgMusic:   []int32{17, 32, 34},
					CuesMusic: []int32{16, 35, 36},
					FmvsMusic: []int32{16, 20},
					BossMusic: []int32{},
				},
				Fmvs: []int32{18, 19},
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sublocations/beSaiD",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"shops":  true,
					"quests": true,
				},
				ExpLengths: map[string]int{
					"connected sublocations": 2,
					"areas":                  16,
					"characters":             1,
					"aeons":                  1,
					"shops":                  0,
					"treasures":              23,
					"monsters":               12,
					"formations":             14,
					"quests":                 0,
					"bg music":               5,
					"cues music":             3,
					"fmvs music":             3,
					"boss music":             2,
					"fmvs":                   3,
				},
			},
			Unique:                newExpUnique(7, "besaid"),
			ParentLocation:        4,
			ConnectedSublocations: []int32{8, 9},
			Areas:                 []int32{20, 24, 29, 31, 35},
			LocRel: LocRel{
				Characters: []int32{3},
				Aeons:      []int32{1},
				Treasures:  []int32{15, 18, 24, 40, 44},
				Monsters:   []int32{7, 11, 14, 18, 293},
				Formations: []int32{9, 13, 17, 22, 25},
				Quests:     []int32{},
				Fmvs:       []int32{5, 6, 8},
				Music: &testLocMusic{
					BgMusic:   []int32{18, 20, 22, 23, 27},
					CuesMusic: []int32{17},
					FmvsMusic: []int32{17, 23, 34},
					BossMusic: []int32{9, 34},
				},
			},
		},
	}

	testSingleResources(t, tests, "GetSublocation", testSrv.HandleSublocations)
}

func TestRetrieveSublocations(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sublocations?limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   43,
			Results: []int32{1, 25, 43},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sublocations?monsters=false",
				ExpectedStatus: http.StatusOK,
			},
			Count:   11,
			Results: []int32{2, 3, 8, 19, 23, 24, 36},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sublocations?characters=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   7,
			Results: []int32{1, 7, 8, 9, 14, 17, 20},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sublocations?aeons=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   8,
			Results: []int32{4, 7, 12, 19, 28, 33, 37, 38},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sublocations?key_item=36",
				ExpectedStatus: http.StatusOK,
			},
			Count:   1,
			Results: []int32{8},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sublocations?location=8",
				ExpectedStatus: http.StatusOK,
			},
			Count:   2,
			Results: []int32{14, 15},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sublocations?item=1&methods=treasure&boss_fights=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   4,
			Results: []int32{1, 7, 9, 17},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sublocations?shops=true&treasures=true&sidequests=true&fmvs=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   3,
			Results: []int32{14, 26, 32},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sublocations?availability=pre-airship&item=56&methods=monster",
				ExpectedStatus: http.StatusOK,
			},
			Count:   5,
			Results: []int32{25, 27, 35, 38, 41},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sublocations?availability=post-game&item=56&methods=monster&pre_airship=false",
				ExpectedStatus: http.StatusOK,
			},
			Count:   5,
			Results: []int32{25, 35, 38, 41, 43},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sublocations?availability=always&auto_ability=4&item=27&methods=monster",
				ExpectedStatus: http.StatusOK,
			},
			Count:   1,
			Results: []int32{25},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sublocations?availability=post&monsters=true&treasures=false&item=53",
				ExpectedStatus: http.StatusOK,
			},
			Count:   2,
			Results: []int32{32, 39},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sublocations?availability=pre-story&monster=31",
				ExpectedStatus: http.StatusOK,
			},
			Count:   1,
			Results: []int32{15},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sublocations?availability=always&monster=87&repeatable=true&boss_fights=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   1,
			Results: []int32{28},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sublocations?availability=pre-story",
				ExpectedStatus: http.StatusOK,
			},
			Count:   6,
			Results: []int32{1, 2, 5, 6, 31, 33},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sublocations?availability=post&sidequests=true&boss_fights=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   3,
			Results: []int32{30, 32, 35},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sublocations?availability=post&monsters=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   11,
			Results: []int32{4, 17, 30, 37, 41, 43},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sublocations?availability=post&monsters=false&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   32,
			Results: []int32{1, 2, 5, 14, 21, 24, 28, 31, 40},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sublocations?availability=always&monster=48",
				ExpectedStatus: http.StatusOK,
			},
			Count:   3,
			Results: []int32{17, 18, 20},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sublocations?availability=pre-story&shops=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   14,
			Results: []int32{13, 21, 28, 32, 34, 39},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sublocations?availability=always&key_item=23",
				ExpectedStatus: http.StatusOK,
			},
			Count:   1,
			Results: []int32{30},
		},
	}

	testIDList(t, tests, apilookups.E.Sublocations.Endpoint, "RetrieveSublocations", testSrv.HandleSublocations, compareAPIResourceLists[apiresource.Named])
}

func TestSubsectionSublocations(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sublocations/5/connected/",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleSublocations,
			},
			Count:          2,
			ParentResource: h.GetStrPtr("/sublocations/5"),
			Results:        []int32{6, 7},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sublocations/12/connected/",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleSublocations,
			},
			Count:          1,
			ParentResource: h.GetStrPtr("/sublocations/12"),
			Results:        []int32{11},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sublocations/27/connected/",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleSublocations,
			},
			Count:          3,
			ParentResource: h.GetStrPtr("/sublocations/27"),
			Results:        []int32{26, 28, 29},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sublocations/35/connected/",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleSublocations,
			},
			Count:          4,
			ParentResource: h.GetStrPtr("/sublocations/35"),
			Results:        []int32{26, 36, 38, 39},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sublocations/simple?availability=pre-story",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleSublocations,
				DontCheck: map[string]bool{
					"parent resource": true,
				},
			},
			Count:   6,
			Results: []int32{1, 2, 5, 6, 31, 33},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sublocations/simple?availability=post&sidequests=true&boss_fights=true",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleSublocations,
				DontCheck: map[string]bool{
					"parent resource": true,
				},
			},
			Count:   3,
			Results: []int32{30, 32, 35},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/locations/15/sublocations/",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleLocations,
			},
			Count:          3,
			ParentResource: h.GetStrPtr("/locations/15"),
			Results:        []int32{26, 27, 28},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/locations/8/sublocations/",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleLocations,
			},
			Count:          2,
			ParentResource: h.GetStrPtr("/locations/8"),
			Results:        []int32{14, 15},
		},
	}

	testIDList(t, tests, apilookups.E.Sublocations.Endpoint, "SubsectionSublocations", nil, compareSimpleResourceLists[apiresource.Named, sectiontypes.SublocationSimple])
}
