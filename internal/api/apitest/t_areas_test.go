package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/sections/sectiontypes"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

func TestGetArea(t *testing.T) {
	t.Parallel()
	tests := []Area{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas/0",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "area with provided id '0' doesn't exist. max id: 237",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas/238",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "area with provided id '238' doesn't exist. max id: 237",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas/a",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "invalid id 'a'",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas/143/",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"characters": true,
					"aeons":      true,
					"shops":      true,
					"cues music": true,
					"fmvs music": true,
					"boss music": true,
					"fmvs":       true,
				},
				ExpLengths: map[string]int{
					"connected areas": 2,
					"characters":      0,
					"aeons":           0,
					"shops":           0,
					"treasures":       1,
					"monsters":        6,
					"formations":      6,
					"quests":          4,
					"bg music":        1,
					"cues music":      0,
					"fmvs music":      0,
					"boss music":      0,
					"fmvs":            0,
				},
			},
			NameVer:           newExpNameVer(143, "north", 1),
			DisplayName:       "macalania woods - north",
			ParentLocation:    15,
			ParentSublocation: 26,
			ConnectedAreas:    []int32{142, 147},
			LocRel: LocRel{
				Treasures:  []int32{191},
				Monsters:   []int32{81, 84, 85},
				Formations: []int32{120, 122, 125},
				Quests:     []int32{83, 85, 86, 87},
				Music: &testLocMusic{
					BgMusic: []int32{30},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas/36",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"aeons":      true,
					"shops":      true,
					"monsters":   true,
					"formations": true,
					"quests":     true,
					"cues music": true,
					"fmvs music": true,
					"boss music": true,
					"fmvs":       true,
				},
				ExpLengths: map[string]int{
					"connected areas": 7,
					"characters":      2,
					"aeons":           0,
					"shops":           0,
					"treasures":       6,
					"monsters":        0,
					"formations":      0,
					"quests":          0,
					"cues music":      4,
					"fmvs music":      0,
					"boss music":      0,
					"fmvs":            0,
				},
			},
			NameVer:           newExpNameVer(36, "besaid village", 0),
			DisplayName:       "besaid village",
			ParentLocation:    4,
			ParentSublocation: 8,
			ConnectedAreas:    []int32{26, 37, 41},
			LocRel: LocRel{
				Characters: []int32{2, 4},
				Treasures:  []int32{33, 37},
				Music: &testLocMusic{
					BgMusic: []int32{19},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas/159?rel_availability=pre-story&rel_repeatable=true",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"parent location":    true,
					"parent sublocation": true,
					"connected areas":    true,
					"characters":         true,
					"aeons":              true,
					"music":              true,
					"fmvs":               true,
				},
				ExpLengths: map[string]int{
					"shops":      0,
					"treasures":  0,
					"monsters":   4,
					"formations": 3,
					"quests":     0,
				},
			},
			NameVer:           newExpNameVer(159, "road", 1),
			DisplayName:       "macalania - road",
			ParentLocation:    15,
			ParentSublocation: 28,
			LocRel: LocRel{
				Shops:      []int32{},
				Treasures:  []int32{},
				Monsters:   []int32{87, 88, 89, 90},
				Formations: []int32{128, 129, 132},
				Quests:     []int32{},
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas/159?rel_availability=always&rel_repeatable=false",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"parent location":    true,
					"parent sublocation": true,
					"connected areas":    true,
					"characters":         true,
					"aeons":              true,
					"music":              true,
					"fmvs":               true,
				},
				ExpLengths: map[string]int{
					"shops":      0,
					"treasures":  0,
					"monsters":   1,
					"formations": 1,
					"quests":     0,
				},
			},
			NameVer:           newExpNameVer(159, "road", 1),
			DisplayName:       "macalania - road",
			ParentLocation:    15,
			ParentSublocation: 28,
			LocRel: LocRel{
				Shops:      []int32{},
				Treasures:  []int32{},
				Monsters:   []int32{297},
				Formations: []int32{136},
				Quests:     []int32{},
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas/203?rel_availability=post",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"parent location":    true,
					"parent sublocation": true,
					"connected areas":    true,
					"characters":         true,
					"aeons":              true,
					"music":              true,
					"fmvs":               true,
				},
				ExpLengths: map[string]int{
					"shops":      0,
					"treasures":  0,
					"monsters":   30,
					"formations": 28,
					"quests":     30,
				},
			},
			NameVer:           newExpNameVer(203, "arena", 0),
			DisplayName:       "calm lands - arena",
			ParentLocation:    20,
			ParentSublocation: 35,
			LocRel: LocRel{
				Shops:      []int32{},
				Treasures:  []int32{},
				Monsters:   []int32{256, 263, 267, 279, 288, 292},
				Formations: []int32{297, 304, 308, 317, 325, 331},
				Quests:     []int32{1, 11, 15, 18, 22, 29, 34, 40, 46},
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas/203?rel_availability=always&rel_repeatable=true",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"parent location":    true,
					"parent sublocation": true,
					"connected areas":    true,
					"characters":         true,
					"aeons":              true,
					"music":              true,
					"fmvs":               true,
				},
				ExpLengths: map[string]int{
					"shops":      0,
					"treasures":  0,
					"monsters":   7,
					"formations": 7,
					"quests":     0,
				},
			},
			NameVer:           newExpNameVer(203, "arena", 0),
			DisplayName:       "calm lands - arena",
			ParentLocation:    20,
			ParentSublocation: 35,
			LocRel: LocRel{
				Shops:      []int32{},
				Treasures:  []int32{},
				Monsters:   []int32{261, 262, 264, 265, 266, 283, 291},
				Formations: []int32{302, 303, 305, 306, 307, 324, 330},
				Quests:     []int32{},
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas/203?rel_availability=always&rel_repeatable=false",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"parent location":    true,
					"parent sublocation": true,
					"connected areas":    true,
					"characters":         true,
					"aeons":              true,
					"music":              true,
					"fmvs":               true,
				},
				ExpLengths: map[string]int{
					"shops":      1,
					"treasures":  1,
					"monsters":   0,
					"formations": 0,
					"quests":     7,
				},
			},
			NameVer:           newExpNameVer(203, "arena", 0),
			DisplayName:       "calm lands - arena",
			ParentLocation:    20,
			ParentSublocation: 35,
			LocRel: LocRel{
				Shops:      []int32{33},
				Treasures:  []int32{271},
				Monsters:   []int32{},
				Formations: []int32{},
				Quests:     []int32{16, 17, 19, 20, 21, 38, 44},
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas/68",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"connected areas": true,
					"characters":      true,
					"aeons":           true,
					"treasures":       true,
					"monsters":        true,
					"formations":      true,
					"quests":          true,
					"fmvs music":      true,
					"boss music":      true,
					"fmvs":            true,
				},
				ExpLengths: map[string]int{
					"connected areas": 6,
					"characters":      0,
					"aeons":           0,
					"shops":           1,
					"treasures":       0,
					"monsters":        0,
					"formations":      0,
					"quests":          0,
					"bg music":        2,
					"cues music":      1,
					"fmvs music":      0,
					"boss music":      0,
					"fmvs":            0,
				},
			},
			NameVer:           newExpNameVer(68, "main gate", 0),
			DisplayName:       "stadium - main gate",
			ParentLocation:    8,
			ParentSublocation: 14,
			LocRel: LocRel{
				Shops: []int32{5},
				Music: &testLocMusic{
					CuesMusic: []int32{35},
					BgMusic:   []int32{32, 34},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas/138",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
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
				ExpLengths: map[string]int{
					"quests": 9,
				},
			},
			NameVer:           newExpNameVer(138, "agency front", 0),
			DisplayName:       "thunder plains - agency front",
			ParentLocation:    14,
			ParentSublocation: 25,
			LocRel: LocRel{
				Quests: []int32{88, 89, 90, 91, 92, 93, 94, 95, 96},
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas/42",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"connected areas": true,
					"aeons":           true,
					"shops":           true,
					"quests":          true,
				},
				ExpLengths: map[string]int{
					"characters": 1,
					"aeons":      0,
					"shops":      0,
					"treasures":  1,
					"formations": 2,
					"monsters":   4,
					"quests":     0,
					"bg music":   1,
					"cues music": 3,
					"fmvs music": 1,
					"boss music": 1,
					"fmvs":       6,
				},
			},
			NameVer:           newExpNameVer(42, "deck", 0),
			DisplayName:       "ss liki - deck",
			ParentLocation:    5,
			ParentSublocation: 9,
			LocRel: LocRel{
				Characters: []int32{5},
				Treasures:  []int32{45},
				Monsters:   []int32{19, 20, 21, 22},
				Formations: []int32{26, 27},
				Fmvs:       []int32{9, 12, 13, 14},
				Music: &testLocMusic{
					BgMusic:   []int32{28},
					CuesMusic: []int32{},
					FmvsMusic: []int32{16},
					BossMusic: []int32{16},
				},
			},
		},
	}

	testSingleResources(t, tests, "GetArea", testSrv.HandleAreas)
}

func TestRetrieveAreas(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas?comp_sphere=fa",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "invalid boolean value 'fa' used for parameter 'comp_sphere'. usage: '?comp_sphere={bool}'",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas?item=113",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "provided id '113' used for parameter 'item' is out of range. max id: 112",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas?key_item=61",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "provided id '61' used for parameter 'key_item' is out of range. max id: 60",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas?location=0",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "provided id '0' used for parameter 'location' is out of range. max id: 26",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas/",
				ExpectedStatus: http.StatusOK,
			},
			Count:   237,
			Next:    h.GetStrPtr("/areas?limit=20&offset=20"),
			Results: []int32{1, 5, 20},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas?limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   237,
			Results: []int32{1, 50, 237},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas?offset=50&limit=30",
				ExpectedStatus: http.StatusOK,
			},
			Count:    237,
			Previous: h.GetStrPtr("/areas?limit=30&offset=20"),
			Next:     h.GetStrPtr("/areas?limit=30&offset=80"),
			Results:  []int32{51, 80},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas?monsters=true&chocobo=true&save_sphere=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   3,
			Results: []int32{87, 96, 201},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas?item=7&availability=post-game&monsters=false",
				ExpectedStatus: http.StatusOK,
			},
			Count:   5,
			Results: []int32{35, 75, 127, 138, 206},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas?item=3&methods=quest,monster",
				ExpectedStatus: http.StatusOK,
			},
			Count:   17,
			Results: []int32{26, 77, 138, 159, 174, 190, 207},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas?characters=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   7,
			Results: []int32{1, 20, 102},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas?sidequests=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   12,
			Results: []int32{74, 138, 142, 143, 145, 180, 183, 201},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas?key_item=37",
				ExpectedStatus: http.StatusOK,
			},
			Count:   2,
			Results: []int32{46, 167},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas?availability=post&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   22,
			Results: []int32{13, 27, 105, 208, 233, 237},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas?location=13&treasures=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   5,
			Results: []int32{127, 130, 132, 134, 136},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas?shops=true&airship=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   10,
			Results: []int32{47, 68, 91, 139, 199, 211},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas?sublocation=27&boss_fights=true", // lake macalania
				ExpectedStatus: http.StatusOK,
			},
			Count:   1,
			Results: []int32{156},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas?comp_sphere=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   5,
			Results: []int32{6, 91, 155, 176, 182},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas?fmvs=true&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   26,
			Results: []int32{3, 34, 49, 125, 179, 235},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas?availability=pre-airship&item=56&methods=monster",
				ExpectedStatus: http.StatusOK,
			},
			Count:   6,
			Results: []int32{137, 140, 156, 204, 209, 223},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas?availability=post-game&item=56&methods=monster&pre_airship=false",
				ExpectedStatus: http.StatusOK,
			},
			Count:   7,
			Results: []int32{137, 140, 203, 209, 223, 236, 237},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas?availability=always&auto_ability=4&item=27&methods=monster",
				ExpectedStatus: http.StatusOK,
			},
			Count:   2,
			Results: []int32{137, 140},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas?availability=post&monsters=true&treasures=false&item=53",
				ExpectedStatus: http.StatusOK,
			},
			Count:   7,
			Results: []int32{26, 99, 100, 170, 190, 211, 228},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas?availability=pre-story&monster=31",
				ExpectedStatus: http.StatusOK,
			},
			Count:   3,
			Results: []int32{76, 77, 78},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas?availability=pre-story&shops=true&monster=87&repeatable=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   1,
			Results: []int32{159},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas?availability=post-story",
				ExpectedStatus: http.StatusOK,
			},
			Count:   2,
			Results: []int32{234, 235},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas?availability=post&sidequests=true&boss_fights=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   2,
			Results: []int32{170, 203},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas?availability=post&monsters=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   18,
			Results: []int32{7, 26, 106, 170, 207, 230, 237},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas?availability=post&monsters=false&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   219,
			Results: []int32{1, 18, 63, 74, 110, 136, 153, 180, 229, 235},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas?availability=always&monster=48",
				ExpectedStatus: http.StatusOK,
			},
			Count:   5,
			Results: []int32{99, 100, 104, 107, 117},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas?availability=pre-story&shops=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   18,
			Results: []int32{65, 96, 123, 147, 183, 200, 212},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas?availability=always&key_item=23",
				ExpectedStatus: http.StatusOK,
			},
			Count:   1,
			Results: []int32{170},
		},
	}

	testIDList(t, tests, apilookups.E.Areas.Endpoint, "RetrieveAreas", testSrv.HandleAreas, compareAPIResourceLists[apiresource.Area])
}

func TestSubsectionAreas(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas/36/connected/",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleAreas,
			},
			Count:          7,
			ParentResource: h.GetStrPtr("/areas/36"),
			Results:        []int32{26, 30, 37, 38, 39, 40, 41},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas/209/connected/",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleAreas,
			},
			Count:          2,
			ParentResource: h.GetStrPtr("/areas/209"),
			Results:        []int32{205, 210},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas/9/connected/",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleAreas,
			},
			Count:          4,
			ParentResource: h.GetStrPtr("/areas/9"),
			Results:        []int32{8, 10, 11, 15},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas/235/connected/",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleAreas,
			},
			Count:          0,
			ParentResource: h.GetStrPtr("/areas/235"),
			Results:        []int32{},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas/149/connected/",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleAreas,
			},
			Count:          3,
			ParentResource: h.GetStrPtr("/areas/149"),
			Results:        []int32{141, 150, 199},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/monsters/45/areas/",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleMonsters,
			},
			Count:          5,
			ParentResource: h.GetStrPtr("/monsters/45"),
			Results:        []int32{87, 88, 89, 92, 93},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/monsters/140/areas/",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleMonsters,
			},
			Count:          4,
			ParentResource: h.GetStrPtr("/monsters/140"),
			Results:        []int32{200, 201, 202, 209},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/monsters/66/areas/",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleMonsters,
			},
			Count:          1,
			ParentResource: h.GetStrPtr("/monsters/66"),
			Results:        []int32{125},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sublocations/42/areas/",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleSublocations,
			},
			Count:          7,
			ParentResource: h.GetStrPtr("/sublocations/42"),
			Results:        []int32{229, 230, 231, 232, 233, 234, 235},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/locations/10/areas/",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleLocations,
			},
			Count:          9,
			ParentResource: h.GetStrPtr("/locations/10"),
			Results:        []int32{98, 99, 100, 103, 106},
		},
	}

	testIDList(t, tests, apilookups.E.Areas.Endpoint, "SubsectionAreas", nil, compareSimpleResourceLists[apiresource.Area, sectiontypes.AreaSimple])
}
