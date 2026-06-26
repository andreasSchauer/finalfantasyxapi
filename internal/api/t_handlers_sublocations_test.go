package api

import (
	"net/http"
	"testing"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func TestGetSublocation(t *testing.T) {
	t.Parallel()
	tests := []expSublocation{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations/0",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "sublocation with provided id '0' doesn't exist. max id: 41.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations/42",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "sublocation with provided id '42' doesn't exist. max id: 41.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations/a",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "sublocation not found: 'a'.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations/34/",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"characters": true,
					"aeons":      true,
					"cues music": true,
					"fmvs music": true,
					"fmvs":       true,
				},
				expLengths: map[string]int{
					"connected sublocations": 4,
					"areas":                  6,
					"characters":             0,
					"aeons":                  0,
					"shops":                  3,
					"treasures":              9,
					"monsters":               51,
					"formations":             50,
					"quests":	              46,
					"bg music":               2,
					"cues music":             1,
					"fmvs music":             0,
					"boss music":             1,
					"fmvs":                   0,
				},
			},
			expUnique:             newExpUnique(34, "calm lands"),
			parentLocation:        20,
			connectedSublocations: []int32{25, 35, 36, 37},
			areas:                 []int32{202, 203, 204, 205, 206, 207},
			expLocRel: expLocRel{
				shops:      []int32{31, 32, 33},
				treasures:  []int32{265, 268, 270, 272},
				monsters:   []int32{138, 142, 149, 259, 270, 282, 292},
				formations: []int32{193, 198, 205, 306, 312, 320, 331},
				quests:     []int32{1, 11, 19, 34, 42, 45, 46, 69, 74, 77},
				music: &testLocMusic{
					bgMusic:   []int32{71, 73},
					bossMusic: []int32{16},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations/25?rel_availability=pre-story",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{
					"music": true,
					"fmvs": true,
				},
				expLengths: map[string]int{
					"connected sublocations": 4,
					"areas":                  13,
					"characters":             0,
					"aeons":                  0,
					"shops":                  1,
					"treasures":              0,
					"monsters":               1,
					"formations":             1,
					"quests":	              0,
				},
			},
			expUnique:             newExpUnique(25, "macalania woods"),
			parentLocation:        15,
			connectedSublocations: []int32{24, 26, 33, 34},
			areas:                 []int32{143, 145, 148, 151, 155},
			expLocRel: expLocRel{
				characters: []int32{},
				aeons:      []int32{},
				shops:      []int32{22},
				treasures:  []int32{},
				monsters:   []int32{86},
				formations: []int32{126},
				quests:     []int32{},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations/25?rel_availability=always&rel_repeatable=true",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{
					"music": true,
					"fmvs": true,
				},
				expLengths: map[string]int{
					"connected sublocations": 4,
					"areas":                  13,
					"characters":             0,
					"aeons":                  0,
					"shops":                  0,
					"treasures":              0,
					"monsters":               6,
					"formations":             6,
					"quests":	              0,
				},
			},
			expUnique:             newExpUnique(25, "macalania woods"),
			parentLocation:        15,
			connectedSublocations: []int32{24, 26, 33, 34},
			areas:                 []int32{143, 145, 148, 151, 155},
			expLocRel: expLocRel{
				characters: []int32{},
				aeons:      []int32{},
				shops:      []int32{},
				treasures:  []int32{},
				monsters:   []int32{80, 81, 82, 83, 84, 85},
				formations: []int32{120, 121, 122, 123, 124, 125},
				quests:     []int32{},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations/25?rel_availability=always&rel_repeatable=false",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{
					"music": true,
					"fmvs": true,
				},
				expLengths: map[string]int{
					"connected sublocations": 4,
					"areas":                  13,
					"characters":             0,
					"aeons":                  0,
					"shops":                  0,
					"treasures":              8,
					"monsters":               0,
					"formations":             0,
					"quests":	              5,
				},
			},
			expUnique:             newExpUnique(25, "macalania woods"),
			parentLocation:        15,
			connectedSublocations: []int32{24, 26, 33, 34},
			areas:                 []int32{143, 145, 148, 151, 155},
			expLocRel: expLocRel{
				characters: []int32{},
				aeons:      []int32{},
				shops:      []int32{},
				treasures:  []int32{187, 188, 189, 190, 191, 192, 193, 194},
				monsters:   []int32{},
				formations: []int32{},
				quests:     []int32{10, 82, 83, 84, 85},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations/13/",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"connected sublocations": 1,
					"areas":                  7,
					"characters":             1,
					"aeons":                  0,
					"shops":                  1,
					"treasures":              5,
					"monsters":               3,
					"formations":             3,
					"quests":	              2,
					"bg music":               3,
					"cues music":             3,
					"fmvs music":             2,
					"boss music":             0,
					"fmvs":                   2,
				},
			},
			expUnique:             newExpUnique(13, "stadium"),
			parentLocation:        8,
			connectedSublocations: []int32{14},
			areas:                 []int32{69, 70, 71, 72, 73, 74, 75},
			expLocRel: expLocRel{
				characters: []int32{6},
				aeons:      []int32{},
				shops:      []int32{5},
				treasures:  []int32{70, 71, 72, 73, 74},
				monsters:   []int32{34, 35, 36},
				formations: []int32{38, 39, 40},
				quests:     []int32{97, 98},
				music: &testLocMusic{
					bgMusic:   []int32{17, 32, 34},
					cuesMusic: []int32{16, 35, 36},
					fmvsMusic: []int32{16, 20},
					bossMusic: []int32{},
				},
				fmvs: []int32{18, 19},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations/beSaiD",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"shops":      true,
					"quests":	  true,
				},
				expLengths: map[string]int{
					"connected sublocations": 2,
					"areas":                  16,
					"characters":             1,
					"aeons":                  1,
					"shops":                  0,
					"treasures":              23,
					"monsters":               12,
					"formations":             14,
					"quests":	              0,
					"bg music":               5,
					"cues music":             3,
					"fmvs music":             3,
					"boss music":             2,
					"fmvs":                   3,
				},
			},
			expUnique:             newExpUnique(6, "besaid"),
			parentLocation:        4,
			connectedSublocations: []int32{7, 8},
			areas:                 []int32{20, 24, 29, 31, 35},
			expLocRel: expLocRel{
				characters: []int32{3},
				aeons:      []int32{1},
				treasures:  []int32{15, 18, 24, 40, 44},
				monsters:   []int32{7, 11, 14, 18, 293},
				formations: []int32{9, 13, 17, 22, 25},
				quests: 	[]int32{},
				fmvs:       []int32{5, 6, 8},
				music: &testLocMusic{
					bgMusic:   []int32{18, 20, 22, 23, 27},
					cuesMusic: []int32{17},
					fmvsMusic: []int32{17, 23, 34},
					bossMusic: []int32{9, 34},
				},
			},
		},
	}

	testSingleResources(t, tests, "GetSublocation", testCfg.HandleSublocations, compareSublocations)
}

func TestRetrieveSublocations(t *testing.T) {
	t.Parallel()
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations?limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   41,
			results: []int32{1, 25, 41},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations?monsters=false",
				expectedStatus: http.StatusOK,
			},
			count:   9,
			results: []int32{2, 9, 18, 21, 23},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations?characters=true",
				expectedStatus: http.StatusOK,
			},
			count:   7,
			results: []int32{1, 6, 7, 8, 13, 16, 19},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations?aeons=true",
				expectedStatus: http.StatusOK,
			},
			count:   8,
			results: []int32{3, 6, 11, 18, 27, 32, 35, 36},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations?key_item=36",
				expectedStatus: http.StatusOK,
			},
			count:   1,
			results: []int32{7},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations?location=8",
				expectedStatus: http.StatusOK,
			},
			count:   2,
			results: []int32{13, 14},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations?item=1&methods=treasure&boss_fights=true",
				expectedStatus: http.StatusOK,
			},
			count:   4,
			results: []int32{1, 6, 8, 16},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations?shops=true&treasures=true&sidequests=true&fmvs=true",
				expectedStatus: http.StatusOK,
			},
			count:   3,
			results: []int32{13, 25, 31},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations?availability=pre-airship&item=56&methods=monster&pre_airship=true",
				expectedStatus: http.StatusOK,
			},
			count:   5,
			results: []int32{24, 26, 34, 36, 39},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations?availability=post-game&item=56&methods=monster&pre_airship=false",
				expectedStatus: http.StatusOK,
			},
			count:   5,
			results: []int32{24, 34, 36, 39, 41},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations?availability=always&auto_ability=4&item=27&methods=monster",
				expectedStatus: http.StatusOK,
			},
			count:   1,
			results: []int32{24},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations?availability=post&monsters=true&treasures=false&item=53",
				expectedStatus: http.StatusOK,
			},
			count:   2,
			results: []int32{31, 37},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations?availability=pre-story&monster=31",
				expectedStatus: http.StatusOK,
			},
			count:   1,
			results: []int32{14},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations?availability=always&monster=87&repeatable=true&boss_fights=true",
				expectedStatus: http.StatusOK,
			},
			count:   1,
			results: []int32{27},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations?availability=pre-story",
				expectedStatus: http.StatusOK,
			},
			count:   5,
			results: []int32{1, 4, 5, 30, 32},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations?availability=post&sidequests=true&boss_fights=true",
				expectedStatus: http.StatusOK,
			},
			count:   3,
			results: []int32{29, 31, 34},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations?availability=post&monsters=true",
				expectedStatus: http.StatusOK,
			},
			count:   11,
			results: []int32{3, 16, 29, 35, 39, 41},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations?availability=post&monsters=false&limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   30,
			results: []int32{1, 5, 14, 21, 24, 28, 30, 38},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations?availability=always&monster=48",
				expectedStatus: http.StatusOK,
			},
			count:   3,
			results: []int32{16, 17, 19},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations?availability=pre-story&shops=true",
				expectedStatus: http.StatusOK,
			},
			count:   14,
			results: []int32{12, 20, 27, 31, 33, 37},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations?availability=always&key_item=23",
				expectedStatus: http.StatusOK,
			},
			count:   1,
			results: []int32{29},
		},
	}

	testIdList(t, tests, testCfg.e.sublocations.endpoint, "RetrieveSublocations", testCfg.HandleSublocations, compareAPIResourceLists[NamedApiResourceList])
}

func TestSubsectionSublocations(t *testing.T) {
	t.Parallel()
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations/4/connected/",
				expectedStatus: http.StatusOK,
				handler:        testCfg.HandleSublocations,
			},
			count:          2,
			parentResource: h.GetStrPtr("/sublocations/4"),
			results:        []int32{5, 6},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations/11/connected/",
				expectedStatus: http.StatusOK,
				handler:        testCfg.HandleSublocations,
			},
			count:          1,
			parentResource: h.GetStrPtr("/sublocations/11"),
			results:        []int32{10},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations/26/connected/",
				expectedStatus: http.StatusOK,
				handler:        testCfg.HandleSublocations,
			},
			count:          3,
			parentResource: h.GetStrPtr("/sublocations/26"),
			results:        []int32{25, 27, 28},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations/34/connected/",
				expectedStatus: http.StatusOK,
				handler:        testCfg.HandleSublocations,
			},
			count:          4,
			parentResource: h.GetStrPtr("/sublocations/34"),
			results:        []int32{25, 35, 36, 37},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations/simple?availability=pre-story",
				expectedStatus: http.StatusOK,
				handler:        testCfg.HandleSublocations,
				dontCheck: map[string]bool{
					"parent resource": true,
				},
			},
			count:   5,
			results: []int32{1, 4, 5, 30, 32},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations/simple?availability=post&sidequests=true&boss_fights=true",
				expectedStatus: http.StatusOK,
				handler:        testCfg.HandleSublocations,
				dontCheck: map[string]bool{
					"parent resource": true,
				},
			},
			count:   3,
			results: []int32{29, 31, 34},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations/15/sublocations/",
				expectedStatus: http.StatusOK,
				handler:        testCfg.HandleLocations,
			},
			count:          3,
			parentResource: h.GetStrPtr("/locations/15"),
			results:        []int32{25, 26, 27},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations/8/sublocations/",
				expectedStatus: http.StatusOK,
				handler:        testCfg.HandleLocations,
			},
			count:          2,
			parentResource: h.GetStrPtr("/locations/8"),
			results:        []int32{13, 14},
		},
	}

	testIdList(t, tests, testCfg.e.sublocations.endpoint, "SubsectionSublocations", nil, compareSimpleResourceLists[NamedAPIResource, SublocationSimple])
}
