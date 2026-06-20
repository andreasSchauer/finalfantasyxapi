package api

import (
	"net/http"
	"testing"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func TestGetLocation(t *testing.T) {
	t.Parallel()
	tests := []expLocation{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations/0",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "location with provided id '0' doesn't exist. max id: 26.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations/27",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "location with provided id '27' doesn't exist. max id: 26.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations/a",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "location not found: 'a'.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations/15/",
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"connected locations": 4,
					"sublocations":        3,
					"characters":          0,
					"aeons":               1,
					"shops":               6,
					"treasures":           27,
					"monsters":            19,
					"formations":          18,
					"quests":          	   7,
					"bg music":            11,
					"cues music":          10,
					"fmvs music":          1,
					"boss music":          3,
					"fmvs":                2,
				},
			},
			expUnique:          newExpUnique(15, "macalania"),
			connectedLocations: []int32{14, 16, 19, 20},
			sublocations:       []int32{25, 26, 27},
			expLocRel: expLocRel{
				characters: []int32{},
				aeons:      []int32{4},
				shops:      []int32{22, 25, 36},
				treasures:  []int32{187, 193, 199, 209, 213},
				monsters:   []int32{80, 83, 87, 93, 94, 97, 297},
				formations: []int32{120, 125, 129, 135, 137},
				quests:     []int32{10, 82, 83, 84, 85, 86, 87},
				fmvs:       []int32{27, 36},
				music: &testLocMusic{
					bgMusic:   []int32{12, 22, 30, 43, 52, 55, 56},
					cuesMusic: []int32{4, 57, 59},
					fmvsMusic: []int32{70},
					bossMusic: []int32{16, 55, 57},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations/15?rel_availability=always&rel_repeatable=true",
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"connected locations": 4,
					"sublocations":        3,
					"characters":          0,
					"aeons":               1,
					"shops":               0,
					"treasures":           0,
					"monsters":            10,
					"formations":          10,
					"quests":          	   0,
					"bg music":            11,
					"cues music":          10,
					"fmvs music":          1,
					"boss music":          3,
					"fmvs":                2,
				},
			},
			expUnique:          newExpUnique(15, "macalania"),
			connectedLocations: []int32{14, 16, 19, 20},
			sublocations:       []int32{25, 26, 27},
			expLocRel: expLocRel{
				characters: []int32{},
				aeons:      []int32{4},
				shops:      []int32{},
				treasures:  []int32{},
				monsters:   []int32{80, 82, 85, 87, 89, 90},
				formations: []int32{120, 123, 125, 127, 129, 132},
				quests:     []int32{},
				fmvs:       []int32{27, 36},
				music: &testLocMusic{
					bgMusic:   []int32{12, 22, 30, 43, 52, 55, 56},
					cuesMusic: []int32{4, 57, 59},
					fmvsMusic: []int32{70},
					bossMusic: []int32{16, 55, 57},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations/15?rel_availability=always&rel_repeatable=false",
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"connected locations": 4,
					"sublocations":        3,
					"characters":          0,
					"aeons":               1,
					"shops":               1,
					"treasures":           17,
					"monsters":            1,
					"formations":          1,
					"quests":          	   5,
					"bg music":            11,
					"cues music":          10,
					"fmvs music":          1,
					"boss music":          3,
					"fmvs":                2,
				},
			},
			expUnique:          newExpUnique(15, "macalania"),
			connectedLocations: []int32{14, 16, 19, 20},
			sublocations:       []int32{25, 26, 27},
			expLocRel: expLocRel{
				characters: []int32{},
				aeons:      []int32{4},
				shops:      []int32{24},
				treasures:  []int32{187, 190, 193, 197, 203, 210},
				monsters:   []int32{297},
				formations: []int32{136},
				quests:     []int32{10, 82, 83, 84, 85},
				fmvs:       []int32{27, 36},
				music: &testLocMusic{
					bgMusic:   []int32{12, 22, 30, 43, 52, 55, 56},
					cuesMusic: []int32{4, 57, 59},
					fmvsMusic: []int32{70},
					bossMusic: []int32{16, 55, 57},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations/15?rel_availability=pre-story",
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"connected locations": 4,
					"sublocations":        3,
					"characters":          0,
					"aeons":               1,
					"shops":               4,
					"treasures":           7,
					"monsters":            12,
					"formations":          7,
					"quests":          	   0,
					"bg music":            11,
					"cues music":          10,
					"fmvs music":          1,
					"boss music":          3,
					"fmvs":                2,
				},
			},
			expUnique:          newExpUnique(15, "macalania"),
			connectedLocations: []int32{14, 16, 19, 20},
			sublocations:       []int32{25, 26, 27},
			expLocRel: expLocRel{
				characters: []int32{},
				aeons:      []int32{4},
				shops:      []int32{22, 23, 25, 26},
				treasures:  []int32{199, 200, 201, 202, 205, 206, 208},
				monsters:   []int32{86, 87, 93, 94, 96, 97},
				formations: []int32{126, 130, 131, 133, 134, 135, 137},
				quests:     []int32{},
				fmvs:       []int32{27, 36},
				music: &testLocMusic{
					bgMusic:   []int32{12, 22, 30, 43, 52, 55, 56},
					cuesMusic: []int32{4, 57, 59},
					fmvsMusic: []int32{70},
					bossMusic: []int32{16, 55, 57},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations/OmeGa_rUInS",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"connected locations": true,
					"characters":          true,
					"aeons":               true,
					"shops":               true,
					"quests":	           true,
					"cues music":          true,
					"fmvs music":          true,
					"fmvs":                true,
				},
				expLengths: map[string]int{
					"connected locations": 0,
					"sublocations":        1,
					"characters":          0,
					"aeons":               0,
					"shops":               0,
					"treasures":           16,
					"monsters":            24,
					"formations":          22,
					"quests":	           0,
					"bg music":            2,
					"cues music":          0,
					"fmvs music":          0,
					"boss music":          2,
					"fmvs":                0,
				},
			},
			expUnique:    newExpUnique(26, "omega ruins"),
			sublocations: []int32{41},
			expLocRel: expLocRel{
				treasures:  []int32{327, 332, 337, 342},
				monsters:   []int32{190, 201, 210, 239, 245, 250, 255},
				formations: []int32{258, 262, 265, 283, 289, 296},
				quests:     []int32{},
				music: &testLocMusic{
					bgMusic:   []int32{81, 82},
					bossMusic: []int32{16, 80},
				},
			},
		},
	}

	testSingleResources(t, tests, "GetLocation", testCfg.HandleLocations, compareLocations)
}

func TestRetrieveLocations(t *testing.T) {
	t.Parallel()
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations?limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   26,
			results: []int32{1, 6, 18, 26},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations?monsters=false",
				expectedStatus: http.StatusOK,
			},
			count:   2,
			results: []int32{7, 13},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations?characters=true",
				expectedStatus: http.StatusOK,
			},
			count:   6,
			results: []int32{1, 4, 5, 8, 10, 12},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations?aeons=true",
				expectedStatus: http.StatusOK,
			},
			count:   8,
			results: []int32{2, 4, 6, 11, 15, 19, 21, 22},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations?item=45&methods=monster",
				expectedStatus: http.StatusOK,
			},
			count:   4,
			results: []int32{18, 20, 25, 26},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations?key_item=13",
				expectedStatus: http.StatusOK,
			},
			count:   1,
			results: []int32{4},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations?boss_fights=true&shops=true&treasures=true&sidequests=true&fmvs=true",
				expectedStatus: http.StatusOK,
			},
			count:   3,
			results: []int32{8, 15, 18},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations?availability=pre-airship&item=56&methods=monster&pre_airship=true",
				expectedStatus: http.StatusOK,
			},
			count:   5,
			results: []int32{14, 15, 20, 22, 24},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations?availability=post-game&item=56&methods=monster&pre_airship=false",
				expectedStatus: http.StatusOK,
			},
			count:   5,
			results: []int32{14, 20, 22, 24, 26},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations?availability=always&auto_ability=4&item=27&methods=monster",
				expectedStatus: http.StatusOK,
			},
			count:   1,
			results: []int32{14},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations?availability=post&monsters=true&treasures=false&item=53",
				expectedStatus: http.StatusOK,
			},
			count:   2,
			results: []int32{18, 23},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations?availability=pre-story&monster=31",
				expectedStatus: http.StatusOK,
			},
			count:   1,
			results: []int32{8},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations?availability=always&monster=87&repeatable=true&boss_fights=true",
				expectedStatus: http.StatusOK,
			},
			count:   1,
			results: []int32{15},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations?availability=pre-story",
				expectedStatus: http.StatusOK,
			},
			count:   3,
			results: []int32{1, 3, 17},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations?availability=post&sidequests=true&boss_fights=true",
				expectedStatus: http.StatusOK,
			},
			count:   3,
			results: []int32{16, 18, 20},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations?availability=post&monsters=true",
				expectedStatus: http.StatusOK,
			},
			count:   11,
			results: []int32{2, 16, 20, 23, 24, 26},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations?availability=post&monsters=false",
				expectedStatus: http.StatusOK,
			},
			count:   15,
			results: []int32{1, 7, 12, 15, 19, 22},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations?availability=always&monster=48",
				expectedStatus: http.StatusOK,
			},
			count:   3,
			results: []int32{10, 11, 12},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations?availability=pre-story&shops=true",
				expectedStatus: http.StatusOK,
			},
			count:   11,
			results: []int32{7, 12, 15, 20, 23},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations?availability=always&key_item=23",
				expectedStatus: http.StatusOK,
			},
			count:   1,
			results: []int32{16},
		},
	}

	testIdList(t, tests, testCfg.e.locations.endpoint, "RetrieveLocations", testCfg.HandleLocations, compareAPIResourceLists[NamedApiResourceList])
}

func TestSubsectionLocations(t *testing.T) {
	t.Parallel()
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations/5/connected/",
				expectedStatus: http.StatusOK,
			},
			count:          2,
			parentResource: h.GetStrPtr("/locations/5"),
			results:        []int32{4, 6},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations/11/connected/",
				expectedStatus: http.StatusOK,
			},
			count:          2,
			parentResource: h.GetStrPtr("/locations/11"),
			results:        []int32{10, 12},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations/20/connected/",
				expectedStatus: http.StatusOK,
			},
			count:          4,
			parentResource: h.GetStrPtr("/locations/20"),
			results:        []int32{15, 21, 22, 23},
		},
	}

	testIdList(t, tests, testCfg.e.locations.endpoint, "SubsectionLocations", testCfg.HandleLocations, compareSimpleResourceLists[NamedAPIResource, LocationSimple])
}
