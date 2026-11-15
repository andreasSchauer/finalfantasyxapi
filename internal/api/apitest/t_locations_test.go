package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/sections/sectiontypes"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

func TestGetLocation(t *testing.T) {
	t.Parallel()
	tests := []Location{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/locations/0",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "location with provided id '0' doesn't exist. max id: 26",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/locations/27",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "location with provided id '27' doesn't exist. max id: 26",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/locations/a",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "location not found: 'a'",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/locations/15/",
				ExpectedStatus: http.StatusOK,
				ExpLengths: map[string]int{
					"connected locations": 4,
					"sublocations":        3,
					"characters":          0,
					"aeons":               1,
					"shops":               6,
					"treasures":           27,
					"monsters":            19,
					"formations":          18,
					"quests":              7,
					"bg music":            11,
					"cues music":          10,
					"fmvs music":          1,
					"boss music":          3,
					"fmvs":                2,
				},
			},
			Unique:             newExpUnique(15, "macalania"),
			ConnectedLocations: []int32{14, 16, 19, 20},
			Sublocations:       []int32{26, 27, 28},
			LocRel: LocRel{
				Characters: []int32{},
				Aeons:      []int32{4},
				Shops:      []int32{22, 25, 36},
				Treasures:  []int32{187, 193, 199, 209, 213},
				Monsters:   []int32{80, 83, 87, 93, 94, 97, 297},
				Formations: []int32{120, 125, 129, 135, 137},
				Quests:     []int32{10, 82, 83, 84, 85, 86, 87},
				Fmvs:       []int32{27, 36},
				Music: &TestLocMusic{
					BgMusic:   []int32{12, 22, 30, 43, 52, 55, 56},
					CuesMusic: []int32{4, 57, 59},
					FmvsMusic: []int32{70},
					BossMusic: []int32{16, 55, 57},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/locations/15?rel_availability=always&rel_repeatable=true",
				ExpectedStatus: http.StatusOK,
				ExpLengths: map[string]int{
					"connected locations": 4,
					"sublocations":        3,
					"characters":          0,
					"aeons":               1,
					"shops":               0,
					"treasures":           0,
					"monsters":            10,
					"formations":          10,
					"quests":              0,
					"bg music":            11,
					"cues music":          10,
					"fmvs music":          1,
					"boss music":          3,
					"fmvs":                2,
				},
			},
			Unique:             newExpUnique(15, "macalania"),
			ConnectedLocations: []int32{14, 16, 19, 20},
			Sublocations:       []int32{26, 27, 28},
			LocRel: LocRel{
				Characters: []int32{},
				Aeons:      []int32{4},
				Shops:      []int32{},
				Treasures:  []int32{},
				Monsters:   []int32{80, 82, 85, 87, 89, 90},
				Formations: []int32{120, 123, 125, 127, 129, 132},
				Quests:     []int32{},
				Fmvs:       []int32{27, 36},
				Music: &TestLocMusic{
					BgMusic:   []int32{12, 22, 30, 43, 52, 55, 56},
					CuesMusic: []int32{4, 57, 59},
					FmvsMusic: []int32{70},
					BossMusic: []int32{16, 55, 57},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/locations/15?rel_availability=always&rel_repeatable=false",
				ExpectedStatus: http.StatusOK,
				ExpLengths: map[string]int{
					"connected locations": 4,
					"sublocations":        3,
					"characters":          0,
					"aeons":               1,
					"shops":               1,
					"treasures":           17,
					"monsters":            1,
					"formations":          1,
					"quests":              5,
					"bg music":            11,
					"cues music":          10,
					"fmvs music":          1,
					"boss music":          3,
					"fmvs":                2,
				},
			},
			Unique:             newExpUnique(15, "macalania"),
			ConnectedLocations: []int32{14, 16, 19, 20},
			Sublocations:       []int32{26, 27, 28},
			LocRel: LocRel{
				Characters: []int32{},
				Aeons:      []int32{4},
				Shops:      []int32{24},
				Treasures:  []int32{187, 190, 193, 197, 203, 210},
				Monsters:   []int32{297},
				Formations: []int32{136},
				Quests:     []int32{10, 82, 83, 84, 85},
				Fmvs:       []int32{27, 36},
				Music: &TestLocMusic{
					BgMusic:   []int32{12, 22, 30, 43, 52, 55, 56},
					CuesMusic: []int32{4, 57, 59},
					FmvsMusic: []int32{70},
					BossMusic: []int32{16, 55, 57},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/locations/15?rel_availability=pre-story",
				ExpectedStatus: http.StatusOK,
				ExpLengths: map[string]int{
					"connected locations": 4,
					"sublocations":        3,
					"characters":          0,
					"aeons":               1,
					"shops":               4,
					"treasures":           7,
					"monsters":            12,
					"formations":          7,
					"quests":              0,
					"bg music":            11,
					"cues music":          10,
					"fmvs music":          1,
					"boss music":          3,
					"fmvs":                2,
				},
			},
			Unique:             newExpUnique(15, "macalania"),
			ConnectedLocations: []int32{14, 16, 19, 20},
			Sublocations:       []int32{26, 27, 28},
			LocRel: LocRel{
				Characters: []int32{},
				Aeons:      []int32{4},
				Shops:      []int32{22, 23, 25, 26},
				Treasures:  []int32{199, 200, 201, 202, 205, 206, 208},
				Monsters:   []int32{86, 87, 93, 94, 96, 97},
				Formations: []int32{126, 130, 131, 133, 134, 135, 137},
				Quests:     []int32{},
				Fmvs:       []int32{27, 36},
				Music: &TestLocMusic{
					BgMusic:   []int32{12, 22, 30, 43, 52, 55, 56},
					CuesMusic: []int32{4, 57, 59},
					FmvsMusic: []int32{70},
					BossMusic: []int32{16, 55, 57},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/locations/OmeGa_rUInS",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"connected locations": true,
					"characters":          true,
					"aeons":               true,
					"shops":               true,
					"quests":              true,
					"cues music":          true,
					"fmvs music":          true,
					"fmvs":                true,
				},
				ExpLengths: map[string]int{
					"connected locations": 0,
					"sublocations":        1,
					"characters":          0,
					"aeons":               0,
					"shops":               0,
					"treasures":           16,
					"monsters":            24,
					"formations":          22,
					"quests":              0,
					"bg music":            2,
					"cues music":          0,
					"fmvs music":          0,
					"boss music":          2,
					"fmvs":                0,
				},
			},
			Unique:       newExpUnique(26, "omega ruins"),
			Sublocations: []int32{43},
			LocRel: LocRel{
				Treasures:  []int32{327, 332, 337, 342},
				Monsters:   []int32{190, 201, 210, 239, 245, 250, 255},
				Formations: []int32{258, 262, 265, 283, 289, 296},
				Quests:     []int32{},
				Music: &TestLocMusic{
					BgMusic:   []int32{81, 82},
					BossMusic: []int32{16, 80},
				},
			},
		},
	}

	testSingleResources(t, tests, "GetLocation", testSrv.HandleLocations)
}

func TestRetrieveLocations(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/locations?limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   26,
			Results: []int32{1, 6, 18, 26},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/locations?monsters=false",
				ExpectedStatus: http.StatusOK,
			},
			Count:   2,
			Results: []int32{7, 13},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/locations?characters=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   6,
			Results: []int32{1, 4, 5, 8, 10, 12},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/locations?aeons=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   8,
			Results: []int32{2, 4, 6, 11, 15, 19, 21, 22},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/locations?item=45&methods=monster",
				ExpectedStatus: http.StatusOK,
			},
			Count:   4,
			Results: []int32{18, 20, 25, 26},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/locations?key_item=13",
				ExpectedStatus: http.StatusOK,
			},
			Count:   1,
			Results: []int32{4},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/locations?boss_fights=true&shops=true&treasures=true&sidequests=true&fmvs=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   3,
			Results: []int32{8, 15, 18},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/locations?availability=pre-airship&item=56&methods=monster",
				ExpectedStatus: http.StatusOK,
			},
			Count:   5,
			Results: []int32{14, 15, 20, 22, 24},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/locations?availability=post-game&item=56&methods=monster&pre_airship=false",
				ExpectedStatus: http.StatusOK,
			},
			Count:   5,
			Results: []int32{14, 20, 22, 24, 26},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/locations?availability=always&auto_ability=4&item=27&methods=monster",
				ExpectedStatus: http.StatusOK,
			},
			Count:   1,
			Results: []int32{14},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/locations?availability=post&monsters=true&treasures=false&item=53",
				ExpectedStatus: http.StatusOK,
			},
			Count:   2,
			Results: []int32{18, 23},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/locations?availability=pre-story&monster=31",
				ExpectedStatus: http.StatusOK,
			},
			Count:   1,
			Results: []int32{8},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/locations?availability=always&monster=87&repeatable=true&boss_fights=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   1,
			Results: []int32{15},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/locations?availability=pre-story",
				ExpectedStatus: http.StatusOK,
			},
			Count:   3,
			Results: []int32{1, 3, 17},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/locations?availability=post&sidequests=true&boss_fights=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   3,
			Results: []int32{16, 18, 20},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/locations?availability=post&monsters=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   11,
			Results: []int32{2, 16, 20, 23, 24, 26},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/locations?availability=post&monsters=false",
				ExpectedStatus: http.StatusOK,
			},
			Count:   15,
			Results: []int32{1, 7, 12, 15, 19, 22},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/locations?availability=always&monster=48",
				ExpectedStatus: http.StatusOK,
			},
			Count:   3,
			Results: []int32{10, 11, 12},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/locations?availability=pre-story&shops=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   11,
			Results: []int32{7, 12, 15, 20, 23},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/locations?availability=always&key_item=23",
				ExpectedStatus: http.StatusOK,
			},
			Count:   1,
			Results: []int32{16},
		},
	}

	testIDList(t, tests, apilookups.E.Locations.Endpoint, "RetrieveLocations", testSrv.HandleLocations, compareAPIResourceLists[apiresource.Named])
}

func TestSubsectionLocations(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/locations/5/connected/",
				ExpectedStatus: http.StatusOK,
			},
			Count:          2,
			ParentResource: h.GetStrPtr("/locations/5"),
			Results:        []int32{4, 6},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/locations/11/connected/",
				ExpectedStatus: http.StatusOK,
			},
			Count:          2,
			ParentResource: h.GetStrPtr("/locations/11"),
			Results:        []int32{10, 12},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/locations/20/connected/",
				ExpectedStatus: http.StatusOK,
			},
			Count:          4,
			ParentResource: h.GetStrPtr("/locations/20"),
			Results:        []int32{15, 21, 22, 23},
		},
	}

	testIDList(t, tests, apilookups.E.Locations.Endpoint, "SubsectionLocations", testSrv.HandleLocations, compareSimpleResourceLists[apiresource.Named, sectiontypes.LocationSimple])
}
