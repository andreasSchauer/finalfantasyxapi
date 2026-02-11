package main

import (
	"net/http"
	"testing"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func TestGetArea(t *testing.T) {
	tests := []expArea{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/0",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "area with provided id '0' doesn't exist. max id: 240.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/241",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "area with provided id '241' doesn't exist. max id: 240.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/a",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid id 'a'.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/145/",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"characters": true,
					"aeons":      true,
					"shops":      true,
					"cues music": true,
					"fmvs music": true,
					"boss music": true,
					"fmvs":       true,
				},
				expLengths: map[string]int{
					"connected areas": 2,
					"characters":      0,
					"aeons":           0,
					"shops":           0,
					"treasures":       1,
					"monsters":        6,
					"formations":      6,
					"sidequests":      1,
					"bg music":        1,
					"cues music":      0,
					"fmvs music":      0,
					"boss music":      0,
					"fmvs":            0,
				},
			},
			expNameVer:        newExpNameVer(145, "north", 1),
			displayName:       "macalania woods - north",
			parentLocation:    15,
			parentSublocation: 25,
			connectedAreas:    []int32{144, 149},
			expLocRel: expLocRel{
				treasures:  []int32{191},
				monsters:   []int32{81, 84, 85},
				formations: []int32{120, 122, 125},
				sidequests: []int32{6},
				music: &testLocMusic{
					bgMusic: []int32{30},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/36",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"aeons":      true,
					"shops":      true,
					"monsters":   true,
					"formations": true,
					"sidequests": true,
					"cues music": true,
					"fmvs music": true,
					"boss music": true,
					"fmvs":       true,
				},
				expLengths: map[string]int{
					"connected areas": 7,
					"characters":      2,
					"aeons":           0,
					"shops":           0,
					"treasures":       6,
					"monsters":        0,
					"formations":      0,
					"sidequests":      0,
					"cues music":      4,
					"fmvs music":      0,
					"boss music":      0,
					"fmvs":            0,
				},
			},
			expNameVer:        newExpNameVer(36, "besaid village", 0),
			displayName:       "besaid village",
			parentLocation:    4,
			parentSublocation: 7,
			connectedAreas:    []int32{26, 37, 41},
			expLocRel: expLocRel{
				characters: []int32{2, 4},
				treasures:  []int32{33, 37},
				music: &testLocMusic{
					bgMusic: []int32{19},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/69",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"connected areas": true,
					"characters":      true,
					"aeons":           true,
					"treasures":       true,
					"monsters":        true,
					"formations":      true,
					"sidequests":      true,
					"fmvs music":      true,
					"boss music":      true,
					"fmvs":            true,
				},
				expLengths: map[string]int{
					"connected areas": 6,
					"characters":      0,
					"aeons":           0,
					"shops":           1,
					"treasures":       0,
					"monsters":        0,
					"formations":      0,
					"sidequests":      0,
					"bg music":        2,
					"cues music":      1,
					"fmvs music":      0,
					"boss music":      0,
					"fmvs":            0,
				},
			},
			expNameVer:        newExpNameVer(69, "main gate", 0),
			displayName:       "stadium - main gate",
			parentLocation:    8,
			parentSublocation: 13,
			expLocRel: expLocRel{
				shops: []int32{5},
				music: &testLocMusic{
					cuesMusic: []int32{35},
					bgMusic:   []int32{32, 34},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/140",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"connected areas": true,
					"characters":      true,
					"aeons":           true,
					"shops":           true,
					"treasures":       true,
					"monsters":        true,
					"formations":      true,
					"music":           true,
					"fmvs":            true,
				},
			},
			expNameVer:        newExpNameVer(140, "agency front", 0),
			displayName:       "thunder plains - agency front",
			parentLocation:    14,
			parentSublocation: 24,
			expLocRel: expLocRel{
				sidequests: []int32{7},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/42",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"connected areas": true,
					"aeons":           true,
					"shops":           true,
					"sidequests":      true,
				},
				expLengths: map[string]int{
					"characters": 1,
					"aeons":      0,
					"shops":      0,
					"treasures":  1,
					"formations": 1,
					"monsters":   2,
					"sidequests": 0,
					"bg music":   1,
					"cues music": 3,
					"fmvs music": 1,
					"boss music": 1,
					"fmvs":       5,
				},
			},
			expNameVer:        newExpNameVer(42, "deck", 0),
			displayName:       "ss liki - deck",
			parentLocation:    5,
			parentSublocation: 8,
			expLocRel: expLocRel{
				characters: []int32{5},
				treasures:  []int32{45},
				monsters:   []int32{19, 20},
				formations: []int32{26},
				fmvs:       []int32{9, 12, 13},
				music: &testLocMusic{
					bgMusic:   []int32{28},
					cuesMusic: []int32{},
					fmvsMusic: []int32{16},
					bossMusic: []int32{16},
				},
			},
		},
	}

	testSingleResources(t, tests, "GetArea", testCfg.HandleAreas, compareAreas)
}

func TestRetrieveAreas(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?comp_sphere=fa",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid boolean value 'fa'. usage: '?comp_sphere={bool}'.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?item=113",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "provided id '113' in 'item' is out of range. max id: 112.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?key_item=61",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "provided id '61' in 'key_item' is out of range. max id: 60.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?location=0",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "provided id '0' in 'location' is out of range. max id: 26.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/",
				expectedStatus: http.StatusOK,
			},
			count:   240,
			next:    h.GetStrPtr("/areas?limit=20&offset=20"),
			results: []int32{1, 5, 20},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   240,
			results: []int32{1, 50, 240},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?offset=50&limit=30",
				expectedStatus: http.StatusOK,
			},
			count:    240,
			previous: h.GetStrPtr("/areas?limit=30&offset=20"),
			next:     h.GetStrPtr("/areas?limit=30&offset=80"),
			results:  []int32{51, 80},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?monsters=true&chocobo=true&save_sphere=true",
				expectedStatus: http.StatusOK,
			},
			count:   3,
			results: []int32{88, 97, 203},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?item=7&story_based=false&monsters=false",
				expectedStatus: http.StatusOK,
			},
			count:   5,
			results: []int32{35, 129, 140, 163, 208},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?characters=true",
				expectedStatus: http.StatusOK,
			},
			count:   7,
			results: []int32{1, 20, 103},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?sidequests=true",
				expectedStatus: http.StatusOK,
			},
			count:   11,
			results: []int32{75, 140, 144, 145, 182, 185, 203},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?key_item=37",
				expectedStatus: http.StatusOK,
			},
			count:   2,
			results: []int32{46, 169},
		},
	}

	testIdList(t, tests, testCfg.e.areas.endpoint, "RetrieveAreas", testCfg.HandleAreas, compareAPIResourceLists[AreaApiResourceList])
}

func TestSubsectionAreas(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/36/connected/",
				expectedStatus: http.StatusOK,
				handler:        testCfg.HandleAreas,
			},
			count:          7,
			parentResource: h.GetStrPtr("/areas/36"),
			results:        []int32{26, 30, 37, 38, 39, 40, 41},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/211/connected/",
				expectedStatus: http.StatusOK,
				handler:        testCfg.HandleAreas,
			},
			count:          2,
			parentResource: h.GetStrPtr("/areas/211"),
			results:        []int32{207, 212},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/9/connected/",
				expectedStatus: http.StatusOK,
				handler:        testCfg.HandleAreas,
			},
			count:          4,
			parentResource: h.GetStrPtr("/areas/9"),
			results:        []int32{8, 10, 11, 15},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/238/connected/",
				expectedStatus: http.StatusOK,
				handler:        testCfg.HandleAreas,
			},
			count:          0,
			parentResource: h.GetStrPtr("/areas/238"),
			results:        []int32{},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/151/connected/",
				expectedStatus: http.StatusOK,
				handler:        testCfg.HandleAreas,
			},
			count:          3,
			parentResource: h.GetStrPtr("/areas/151"),
			results:        []int32{143, 152, 201},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/45/areas/",
				expectedStatus: http.StatusOK,
				handler:        testCfg.HandleMonsters,
			},
			count:          5,
			parentResource: h.GetStrPtr("/monsters/45"),
			results:        []int32{88, 89, 90, 93, 94},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/140/areas/",
				expectedStatus: http.StatusOK,
				handler:        testCfg.HandleMonsters,
			},
			count:          4,
			parentResource: h.GetStrPtr("/monsters/140"),
			results:        []int32{202, 203, 204, 211},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/66/areas/",
				expectedStatus: http.StatusOK,
				handler:        testCfg.HandleMonsters,
			},
			count:          1,
			parentResource: h.GetStrPtr("/monsters/66"),
			results:        []int32{127},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations/40/areas/",
				expectedStatus: http.StatusOK,
				handler:        testCfg.HandleSublocations,
			},
			count:          8,
			parentResource: h.GetStrPtr("/sublocations/40"),
			results:        []int32{231, 232, 233, 234, 235, 236, 237, 238},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations/10/areas/",
				expectedStatus: http.StatusOK,
				handler:        testCfg.HandleLocations,
			},
			count:          9,
			parentResource: h.GetStrPtr("/locations/10"),
			results:        []int32{99, 100, 101, 104, 107},
		},
	}

	testIdList(t, tests, testCfg.e.areas.endpoint, "SubsectionAreas", nil, compareSubResourceLists[AreaAPIResource, AreaSub])
}