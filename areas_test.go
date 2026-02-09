package main

import (
	"net/http"
	"testing"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type expArea struct {
	testGeneral
	expNameVer
	parentLocation    int32
	parentSublocation int32
	connectedAreas    []int32
	expLocRel
}

func (e expArea) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareAreas(test test, exp expArea, got Area) {
	compareExpNameVer(test, exp.expNameVer, got.ID, got.Name, got.Version)
	compIdApiResource(test, "location", testCfg.e.locations.endpoint, exp.parentLocation, got.ParentLocation)
	compIdApiResource(test, "sublocation", testCfg.e.sublocations.endpoint, exp.parentSublocation, got.ParentSublocation)
	compareResListTest(test, rltIDs("connected areas", testCfg.e.areas.endpoint, exp.connectedAreas, got.ConnectedAreas))
	compareLocRel(test, exp.expLocRel, got.LocRel)
}

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
					"characters": 	true,
					"aeons": 		true,
					"shops": 		true,
					"cues music": 	true,
					"fmvs music": 	true,
					"boss music": 	true,
					"fmvs": 		true,
				},
				expLengths: map[string]int{
					"connected areas": 	2,
					"characters":		0,
					"aeons":			0,
					"shops":			0,
					"treasures":		1,
					"monsters":        	6,
					"formations":      	6,
					"sidequests":		1,
					"bg music": 		1,
					"cues music": 		0,
					"fmvs music": 		0,
					"boss music": 		0,
					"fmvs": 			0,
				},
			},
			expNameVer: expNameVer{
				id:      145,
				name:    "north",
				version: h.GetInt32Ptr(1),
			},
			parentLocation:    15,
			parentSublocation: 25,
			connectedAreas:    []int32{144, 149},
			expLocRel: expLocRel{
				treasures:	[]int32{191},
				monsters:   []int32{81, 84, 85},
				formations: []int32{120, 122, 125},
				sidequests: []int32{6},
				music: h.GetStructPtr(testLocMusic{
					bgMusic: []int32{30},
				}),
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/36",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"aeons": 		true,
					"shops": 		true,
					"monsters":		true,
					"formations": 	true,
					"sidequests": 	true,
					"cues music": 	true,
					"fmvs music": 	true,
					"boss music": 	true,
					"fmvs": 		true,
				},
				expLengths: map[string]int{
					"connected areas": 	7,
					"characters":      	2,
					"aeons":			0,
					"shops": 			0,
					"treasures":       	6,
					"monsters":        	0,
					"formations": 		0,
					"sidequests": 		0,
					"cues music": 		0,
					"fmvs music": 		0,
					"boss music": 		0,
					"fmvs": 			0,
				},
			},
			expNameVer: expNameVer{
				id:      36,
				name:    "besaid village",
				version: nil,
			},
			parentLocation:    4,
			parentSublocation: 7,
			connectedAreas:    []int32{26, 37, 41},
			expLocRel: expLocRel{
				characters: []int32{2, 4},
				treasures:  []int32{33, 37},
				music: h.GetStructPtr(testLocMusic{
					bgMusic:    []int32{19},
				}),
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/69",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"connected areas": 	true,
					"characters": 		true,
					"aeons": 			true,
					"treasures": 		true,
					"monsters": 		true,
					"formations": 		true,
					"sidequests": 		true,
					"fmvs music": 		true,
					"boss music": 		true,
					"fmvs": 			true,
				},
				expLengths: map[string]int{
					"connected areas": 	6,
					"characters":      	0,
					"aeons":			0,
					"shops":           	1,
					"treasures":       	0,
					"monsters":        	0,
					"formations": 		0,
					"sidequests": 		0,
					"bg music":        	2,
					"cues music":     	1,
					"fmvs music": 		0,
					"boss music": 		0,
					"fmvs": 			0,
				},
			},
			expNameVer: expNameVer{
				id:      69,
				name:    "main gate",
				version: nil,
			},
			parentLocation:    8,
			parentSublocation: 13,
			expLocRel: expLocRel{
				shops:     []int32{5},
				music: h.GetStructPtr(testLocMusic{
					cuesMusic: []int32{35},
					bgMusic:   []int32{32, 34},
				}),
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/140",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"connected areas": 	true,
					"characters": 		true,
					"aeons": 			true,
					"shops": 			true,
					"treasures": 		true,
					"monsters": 		true,
					"formations": 		true,
					"music": 			true,
					"fmvs": 			true,
				},
			},
			expNameVer: expNameVer{
				id:      140,
				name:    "agency front",
				version: nil,
			},
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
					"connected areas": 	true,
					"aeons": 			true,
					"shops": 			true,
					"sidequests":		true,
				},
				expLengths: map[string]int{
					"characters": 	1,
					"aeons": 		0,
					"shops": 		0,
					"treasures": 	1,
					"formations": 	1,
					"monsters":   	2,
					"sidequests":	0,
					"bg music": 	1,
					"cues music":	0,
					"fmvs music": 	1,
					"boss music": 	1,
					"fmvs":       	5,
				},
			},
			expNameVer: expNameVer{
				id:      42,
				name:    "deck",
				version: nil,
			},
			parentLocation:    5,
			parentSublocation: 8,
			expLocRel: expLocRel{
				characters: []int32{5},
				treasures: 	[]int32{45},
				monsters:   []int32{19, 20},
				formations: []int32{26},
				fmvs:       []int32{9, 12, 13},
				music: h.GetStructPtr(testLocMusic{
					bgMusic: 	[]int32{28},
					cuesMusic: 	[]int32{},
					fmvsMusic:  []int32{16},
					bossMusic:  []int32{16},
				}),
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
				expectedErr:    "invalid boolean value 'fa'. usage: '?comp_sphere={boolean}'.",
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

func TestAreasConnected(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/36/connected/",
				expectedStatus: http.StatusOK,
			},
			count:          7,
			parentResource: h.GetStrPtr("/areas/36"),
			results:        []int32{26, 30, 37, 38, 39, 40, 41},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/211/connected/",
				expectedStatus: http.StatusOK,
			},
			count:          2,
			parentResource: h.GetStrPtr("/areas/211"),
			results:        []int32{207, 212},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/9/connected/",
				expectedStatus: http.StatusOK,
			},
			count:          4,
			parentResource: h.GetStrPtr("/areas/9"),
			results:        []int32{8, 10, 11, 15},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/238/connected/",
				expectedStatus: http.StatusOK,
			},
			count:          0,
			parentResource: h.GetStrPtr("/areas/238"),
			results:        []int32{},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/151/connected/",
				expectedStatus: http.StatusOK,
			},
			count:          3,
			parentResource: h.GetStrPtr("/areas/151"),
			results:        []int32{143, 152, 201},
		},
	}

	testIdList(t, tests, testCfg.e.areas.endpoint, "SubsectionAreasConnected", testCfg.HandleAreas, compareSubResourceLists[AreaAPIResource, AreaSub])
}

func TestAreasMonsters(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/90/monsters/",
				expectedStatus: http.StatusOK,
			},
			count:          6,
			parentResource: h.GetStrPtr("/areas/90"),
			results:        []int32{38, 39, 40, 42, 43, 45},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/23/monsters/",
				expectedStatus: http.StatusOK,
			},
			count:          4,
			parentResource: h.GetStrPtr("/areas/23"),
			results:        []int32{15, 16, 17, 18},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/239/monsters/",
				expectedStatus: http.StatusOK,
			},
			count:          21,
			next:           h.GetStrPtr("/areas/239/monsters?limit=20&offset=20"),
			parentResource: h.GetStrPtr("/areas/239"),
			results:        []int32{190, 201, 210, 239, 245, 249, 253},
		},
	}

	testIdList(t, tests, testCfg.e.monsters.endpoint, "SubsectionAreasMonsters", testCfg.HandleAreas, compareSubResourceLists[NamedAPIResource, MonsterSub])
}

func TestAreasParameters(t *testing.T) {
	tests := []expListNames{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/parameters?limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   20,
			results: []string{"limit", "offset", "item", "save_sphere", "sublocation"},
		},
	}

	testNameList(t, tests, "", "AreasParameters", testCfg.HandleAreas, compareParameterLists)
}

func TestAreasSections(t *testing.T) {
	tests := []expListNames{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/sections",
				expectedStatus: http.StatusOK,
			},
			count: 5,
			results: []string{
				"connected",
				"monsters",
				"monster-formations",
				"songs",
				"treasures",
			},
		},
	}

	testNameList(t, tests, testCfg.e.areas.endpoint, "AreasSections", testCfg.HandleAreas, compareSectionLists)
}
