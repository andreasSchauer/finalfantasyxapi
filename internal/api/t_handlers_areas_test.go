package api

import (
	"net/http"
	"testing"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func TestGetArea(t *testing.T) {
	t.Parallel()
	tests := []expArea{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/0",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "area with provided id '0' doesn't exist. max id: 237.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/238",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "area with provided id '238' doesn't exist. max id: 237.",
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
				requestURL:     "/api/areas/143/",
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
					"quests":	       4,
					"bg music":        1,
					"cues music":      0,
					"fmvs music":      0,
					"boss music":      0,
					"fmvs":            0,
				},
			},
			expNameVer:        newExpNameVer(143, "north", 1),
			displayName:       "macalania woods - north",
			parentLocation:    15,
			parentSublocation: 26,
			connectedAreas:    []int32{142, 147},
			expLocRel: expLocRel{
				treasures:  []int32{191},
				monsters:   []int32{81, 84, 85},
				formations: []int32{120, 122, 125},
				quests:     []int32{83, 85, 86, 87},
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
					"quests":	  true,
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
					"quests":	       0,
					"cues music":      4,
					"fmvs music":      0,
					"boss music":      0,
					"fmvs":            0,
				},
			},
			expNameVer:        newExpNameVer(36, "besaid village", 0),
			displayName:       "besaid village",
			parentLocation:    4,
			parentSublocation: 8,
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
				requestURL:     "/api/areas/159?rel_availability=pre-story&rel_repeatable=true",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"parent location": 	  true,
					"parent sublocation": true,
					"connected areas":	  true,
					"characters": 		  true,
					"aeons":      		  true,
					"music": 	  		  true,
					"fmvs":       		  true,
				},
				expLengths: map[string]int{
					"shops":           0,
					"treasures":       0,
					"monsters":        4,
					"formations":      3,
					"quests":	       0,
				},
			},
			expNameVer:        newExpNameVer(159, "road", 1),
			displayName:       "macalania - road",
			parentLocation: 	15,
			parentSublocation: 	28,
			expLocRel: expLocRel{
				shops: 		[]int32{},
				treasures:  []int32{},
				monsters:  	[]int32{87, 88, 89, 90},
				formations: []int32{128, 129, 132},
				quests:  	[]int32{},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/159?rel_availability=always&rel_repeatable=false",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"parent location": 	  true,
					"parent sublocation": true,
					"connected areas":	  true,
					"characters": 		  true,
					"aeons":      		  true,
					"music": 	  		  true,
					"fmvs":       		  true,
				},
				expLengths: map[string]int{
					"shops":           0,
					"treasures":       0,
					"monsters":        1,
					"formations":      1,
					"quests":	       0,
				},
			},
			expNameVer:        newExpNameVer(159, "road", 1),
			displayName:       "macalania - road",
			parentLocation: 	15,
			parentSublocation: 	28,
			expLocRel: expLocRel{
				shops: 		[]int32{},
				treasures:  []int32{},
				monsters:  	[]int32{297},
				formations: []int32{136},
				quests:  	[]int32{},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/203?rel_availability=post",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"parent location": 	  true,
					"parent sublocation": true,
					"connected areas":	  true,
					"characters": 		  true,
					"aeons":      		  true,
					"music": 	  		  true,
					"fmvs":       		  true,
				},
				expLengths: map[string]int{
					"shops":           0,
					"treasures":       0,
					"monsters":        30,
					"formations":      28,
					"quests":	       30,
				},
			},
			expNameVer:        newExpNameVer(203, "arena", 0),
			displayName:       "calm lands - arena",
			parentLocation: 	20,
			parentSublocation: 	35,
			expLocRel: expLocRel{
				shops: 		[]int32{},
				treasures:  []int32{},
				monsters:  	[]int32{256, 263, 267, 279, 288, 292},
				formations: []int32{297, 304, 308, 317, 325, 331},
				quests:  	[]int32{1, 11, 15, 18, 22, 29, 34, 40, 46},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/203?rel_availability=always&rel_repeatable=true",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"parent location": 	  true,
					"parent sublocation": true,
					"connected areas":	  true,
					"characters": 		  true,
					"aeons":      		  true,
					"music": 	  		  true,
					"fmvs":       		  true,
				},
				expLengths: map[string]int{
					"shops":           0,
					"treasures":       0,
					"monsters":        7,
					"formations":      7,
					"quests":	       0,
				},
			},
			expNameVer:        newExpNameVer(203, "arena", 0),
			displayName:       "calm lands - arena",
			parentLocation: 	20,
			parentSublocation: 	35,
			expLocRel: expLocRel{
				shops: 		[]int32{},
				treasures:  []int32{},
				monsters:  	[]int32{261, 262, 264, 265, 266, 283, 291},
				formations: []int32{302, 303, 305, 306, 307, 324, 330},
				quests:  	[]int32{},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/203?rel_availability=always&rel_repeatable=false",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"parent location": 	  true,
					"parent sublocation": true,
					"connected areas":	  true,
					"characters": 		  true,
					"aeons":      		  true,
					"music": 	  		  true,
					"fmvs":       		  true,
				},
				expLengths: map[string]int{
					"shops":           1,
					"treasures":       1,
					"monsters":        0,
					"formations":      0,
					"quests":	       7,
				},
			},
			expNameVer:        newExpNameVer(203, "arena", 0),
			displayName:       "calm lands - arena",
			parentLocation: 	20,
			parentSublocation: 	35,
			expLocRel: expLocRel{
				shops: 		[]int32{33},
				treasures:  []int32{271},
				monsters:  	[]int32{},
				formations: []int32{},
				quests:  	[]int32{16, 17, 19, 20, 21, 38, 44},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/68",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"connected areas": true,
					"characters":      true,
					"aeons":           true,
					"treasures":       true,
					"monsters":        true,
					"formations":      true,
					"quests":	       true,
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
					"quests":	       0,
					"bg music":        2,
					"cues music":      1,
					"fmvs music":      0,
					"boss music":      0,
					"fmvs":            0,
				},
			},
			expNameVer:        newExpNameVer(68, "main gate", 0),
			displayName:       "stadium - main gate",
			parentLocation:    8,
			parentSublocation: 14,
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
				requestURL:     "/api/areas/138",
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
				expLengths: map[string]int{
					"quests": 9,
				},
			},
			expNameVer:        newExpNameVer(138, "agency front", 0),
			displayName:       "thunder plains - agency front",
			parentLocation:    14,
			parentSublocation: 25,
			expLocRel: expLocRel{
				quests: []int32{88, 89, 90, 91, 92, 93, 94, 95, 96},
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
					"quests":	       true,
				},
				expLengths: map[string]int{
					"characters": 1,
					"aeons":      0,
					"shops":      0,
					"treasures":  1,
					"formations": 2,
					"monsters":   4,
					"quests":	  0,
					"bg music":   1,
					"cues music": 3,
					"fmvs music": 1,
					"boss music": 1,
					"fmvs":       6,
				},
			},
			expNameVer:        newExpNameVer(42, "deck", 0),
			displayName:       "ss liki - deck",
			parentLocation:    5,
			parentSublocation: 9,
			expLocRel: expLocRel{
				characters: []int32{5},
				treasures:  []int32{45},
				monsters:   []int32{19, 20, 21, 22},
				formations: []int32{26, 27},
				fmvs:       []int32{9, 12, 13, 14},
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
	t.Parallel()
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?comp_sphere=fa",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid boolean value 'fa' used for parameter 'comp_sphere'. usage: '?comp_sphere={bool}'.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?item=113",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "provided id '113' used for parameter 'item' is out of range. max id: 112.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?key_item=61",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "provided id '61' used for parameter 'key_item' is out of range. max id: 60.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?location=0",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "provided id '0' used for parameter 'location' is out of range. max id: 26.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/",
				expectedStatus: http.StatusOK,
			},
			count:   237,
			next:    h.GetStrPtr("/areas?limit=20&offset=20"),
			results: []int32{1, 5, 20},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   237,
			results: []int32{1, 50, 237},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?offset=50&limit=30",
				expectedStatus: http.StatusOK,
			},
			count:    237,
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
			results: []int32{87, 96, 201},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?item=7&availability=post-game&monsters=false",
				expectedStatus: http.StatusOK,
			},
			count:   5,
			results: []int32{35, 75, 127, 138, 206},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?item=3&methods=quest,monster",
				expectedStatus: http.StatusOK,
			},
			count:   17,
			results: []int32{26, 77, 138, 159, 174, 190, 207},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?characters=true",
				expectedStatus: http.StatusOK,
			},
			count:   7,
			results: []int32{1, 20, 102},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?sidequests=true",
				expectedStatus: http.StatusOK,
			},
			count:   12,
			results: []int32{74, 138, 142, 143, 145, 180, 183, 201},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?key_item=37",
				expectedStatus: http.StatusOK,
			},
			count:   2,
			results: []int32{46, 167},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?availability=post&limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   22,
			results: []int32{13, 27, 105, 208, 233, 237},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?location=13&treasures=true",
				expectedStatus: http.StatusOK,
			},
			count:   5,
			results: []int32{127, 130, 132, 134, 136},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?shops=true&airship=true",
				expectedStatus: http.StatusOK,
			},
			count:   10,
			results: []int32{47, 68, 91, 139, 199, 211},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?sublocation=27&boss_fights=true", // lake macalania
				expectedStatus: http.StatusOK,
			},
			count:   1,
			results: []int32{156},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?comp_sphere=true",
				expectedStatus: http.StatusOK,
			},
			count:   5,
			results: []int32{6, 91, 155, 176, 182},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?fmvs=true&limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   26,
			results: []int32{3, 34, 49, 125, 179, 235},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?availability=pre-airship&item=56&methods=monster",
				expectedStatus: http.StatusOK,
			},
			count:   6,
			results: []int32{137, 140, 156, 204, 209, 223},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?availability=post-game&item=56&methods=monster&pre_airship=false",
				expectedStatus: http.StatusOK,
			},
			count:   7,
			results: []int32{137, 140, 203, 209, 223, 236, 237},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?availability=always&auto_ability=4&item=27&methods=monster",
				expectedStatus: http.StatusOK,
			},
			count:   2,
			results: []int32{137, 140},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?availability=post&monsters=true&treasures=false&item=53",
				expectedStatus: http.StatusOK,
			},
			count:   7,
			results: []int32{26, 99, 100, 170, 190, 211, 228},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?availability=pre-story&monster=31",
				expectedStatus: http.StatusOK,
			},
			count:   3,
			results: []int32{76, 77, 78},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?availability=pre-story&shops=true&monster=87&repeatable=true",
				expectedStatus: http.StatusOK,
			},
			count:   1,
			results: []int32{159},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?availability=post-story",
				expectedStatus: http.StatusOK,
			},
			count:   2,
			results: []int32{234, 235},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?availability=post&sidequests=true&boss_fights=true",
				expectedStatus: http.StatusOK,
			},
			count:   2,
			results: []int32{170, 203},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?availability=post&monsters=true",
				expectedStatus: http.StatusOK,
			},
			count:   18,
			results: []int32{7, 26, 106, 170, 207, 230, 237},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?availability=post&monsters=false&limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   219,
			results: []int32{1, 18, 63, 74, 110, 136, 153, 180, 229, 235},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?availability=always&monster=48",
				expectedStatus: http.StatusOK,
			},
			count:   5,
			results: []int32{99, 100, 104, 107, 117},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?availability=pre-story&shops=true",
				expectedStatus: http.StatusOK,
			},
			count:   18,
			results: []int32{65, 96, 123, 147, 183, 200, 212},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?availability=always&key_item=23",
				expectedStatus: http.StatusOK,
			},
			count:   1,
			results: []int32{170},
		},
	}

	testIdList(t, tests, testCfg.e.areas.endpoint, "RetrieveAreas", testCfg.HandleAreas, compareAPIResourceLists[AreaApiResourceList])
}

func TestSubsectionAreas(t *testing.T) {
	t.Parallel()
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
				requestURL:     "/api/areas/209/connected/",
				expectedStatus: http.StatusOK,
				handler:        testCfg.HandleAreas,
			},
			count:          2,
			parentResource: h.GetStrPtr("/areas/209"),
			results:        []int32{205, 210},
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
				requestURL:     "/api/areas/235/connected/",
				expectedStatus: http.StatusOK,
				handler:        testCfg.HandleAreas,
			},
			count:          0,
			parentResource: h.GetStrPtr("/areas/235"),
			results:        []int32{},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/149/connected/",
				expectedStatus: http.StatusOK,
				handler:        testCfg.HandleAreas,
			},
			count:          3,
			parentResource: h.GetStrPtr("/areas/149"),
			results:        []int32{141, 150, 199},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/45/areas/",
				expectedStatus: http.StatusOK,
				handler:        testCfg.HandleMonsters,
			},
			count:          5,
			parentResource: h.GetStrPtr("/monsters/45"),
			results:        []int32{87, 88, 89, 92, 93},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/140/areas/",
				expectedStatus: http.StatusOK,
				handler:        testCfg.HandleMonsters,
			},
			count:          4,
			parentResource: h.GetStrPtr("/monsters/140"),
			results:        []int32{200, 201, 202, 209},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/66/areas/",
				expectedStatus: http.StatusOK,
				handler:        testCfg.HandleMonsters,
			},
			count:          1,
			parentResource: h.GetStrPtr("/monsters/66"),
			results:        []int32{125},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations/42/areas/",
				expectedStatus: http.StatusOK,
				handler:        testCfg.HandleSublocations,
			},
			count:          7,
			parentResource: h.GetStrPtr("/sublocations/42"),
			results:        []int32{229, 230, 231, 232, 233, 234, 235},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations/10/areas/",
				expectedStatus: http.StatusOK,
				handler:        testCfg.HandleLocations,
			},
			count:          9,
			parentResource: h.GetStrPtr("/locations/10"),
			results:        []int32{98, 99, 100, 103, 106},
		},
	}

	testIdList(t, tests, testCfg.e.areas.endpoint, "SubsectionAreas", nil, compareSimpleResourceLists[AreaAPIResource, AreaSimple])
}
