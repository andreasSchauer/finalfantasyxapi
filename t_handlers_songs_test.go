package main

import (
	"net/http"
	"testing"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)


func TestGetSong(t *testing.T) {
	tests := []expSong{
		{
			testGeneral: testGeneral{
				requestURL: 	"/api/songs/a",
				expectedStatus: http.StatusNotFound,
				expectedErr: 	"song not found: 'a'.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL: 	"/api/songs/a/2",
				expectedStatus: http.StatusBadRequest,
				expectedErr: 	"wrong format. usage: '/api/songs', '/api/songs/{id}', '/api/songs/{name}'",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL: 	"/api/songs/2/a",
				expectedStatus: http.StatusBadRequest,
				expectedErr: 	"endpoint /songs doesn't have any subsections.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL: 	"/api/songs/44",
				expectedStatus: http.StatusOK,
				dontCheck: 		map[string]bool{},
				expLengths: 	map[string]int{
					"background music": 1,
					"cues": 			3,
					"boss fights": 		0,
					"fmvs": 			0,
				},
			},
			expUnique: 			newExpUnique(44, "twilight"),
			composer: 			h.GetStrPtr("junya nakano"),
			arranger: 			h.GetStrPtr("junya nakano"),
			durationInSeconds: 	280,
			canLoop: 			true,
			backgroundMusic: []testBackgroundMusic{
				{
					index: 					0,
					replacesEncounterMusic: false,
					areas: 					[]int32{234, 235, 236},
				},
			},
			cues: []testCue{
				{
					index: 0,
					triggerArea: 			h.GetInt32Ptr(104),
					includedAreas: 			[]int32{},
					replacesEncounterMusic: false,
					replacesBGMusic: 		nil,
				},
				{
					index: 1,
					triggerArea: 			h.GetInt32Ptr(160),
					includedAreas: 			[]int32{},
					replacesEncounterMusic: false,
					replacesBGMusic: 		nil,
				},
				{
					index: 2,
					triggerArea: 			h.GetInt32Ptr(216),
					includedAreas: 			[]int32{216},
					replacesEncounterMusic: false,
					replacesBGMusic: 		h.GetStrPtr("until-trigger"),
				},
			},
			bossFights: []int32{},
			fmvs: []int32{},
		},
		{
			testGeneral: testGeneral{
				requestURL: 	"/api/songs/a_fLeeting_dream",
				expectedStatus: http.StatusOK,
				dontCheck: 		map[string]bool{},
				expLengths: 	map[string]int{
					"background music": 1,
					"cues": 			6,
					"boss fights": 		0,
					"fmvs": 			1,
				},
			},
			expUnique: 			newExpUnique(78, "a fleeting dream"),
			composer: 			h.GetStrPtr("nobuo uematsu"),
			arranger: 			h.GetStrPtr("masashi hamauzu"),
			durationInSeconds: 	264,
			canLoop: 			true,
			backgroundMusic: []testBackgroundMusic{
				{
					index: 					0,
					replacesEncounterMusic: true,
					areas: 					[]int32{221, 222},
				},
			},
			cues: []testCue{
				{
					index: 0,
					triggerArea: 			h.GetInt32Ptr(42),
					includedAreas: 			[]int32{},
					replacesEncounterMusic: false,
					replacesBGMusic: 		nil,
				},
				{
					index: 2,
					triggerArea: 			h.GetInt32Ptr(189),
					includedAreas: 			[]int32{181, 182, 183, 184, 185, 186, 187, 188, 189, 190, 191, 192},
					replacesEncounterMusic: false,
					replacesBGMusic: 		h.GetStrPtr("until-trigger"),
				},
				{
					index: 5,
					triggerArea: 			nil,
					includedAreas: 			[]int32{},
					replacesEncounterMusic: false,
					replacesBGMusic: 		nil,
				},
			},
			bossFights: []int32{},
			fmvs: []int32{51},
		},
		{
			testGeneral: testGeneral{
				requestURL: 	"/api/songs/80",
				expectedStatus: http.StatusOK,
				dontCheck: 		map[string]bool{
					"background music": true,
					"cues": 			true,
				},
				expLengths: 	map[string]int{
					"background music": 0,
					"cues": 			0,
					"boss fights": 		5,
					"fmvs": 			1,
				},
			},
			expUnique: 			newExpUnique(80, "challenge"),
			composer: 			h.GetStrPtr("masashi hamauzu"),
			arranger: 			h.GetStrPtr("masashi hamauzu"),
			durationInSeconds: 	258,
			canLoop: 			true,
			bossFights: []int32{85, 235, 254, 255, 296},
			fmvs: []int32{22},
		},
	}

	testSingleResources(t, tests, "GetSongs", testCfg.HandleSongs, compareSongs)
}

func TestRetrieveSongs(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/songs?limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   95,
			results: []int32{1, 15, 18, 33, 47, 56, 71, 80, 93, 95},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/songs?location=13",
				expectedStatus: http.StatusOK,
			},
			count:   4,
			results: []int32{20, 43, 49, 77},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/songs?sublocation=14",
				expectedStatus: http.StatusOK,
			},
			count:   12,
			results: []int32{10, 16, 17, 27, 32, 34, 75, 77},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/songs?area=213",
				expectedStatus: http.StatusOK,
			},
			count:   3,
			results: []int32{9, 27, 75},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/songs?composer=nobuo-uematsu&limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   53,
			results: []int32{2, 10, 19, 40, 42, 67, 78, 91, 95},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/songs?arranger=3&limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   34,
			results: []int32{18, 41, 46, 61, 64, 76, 88},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/songs?special_use=true",
				expectedStatus: http.StatusOK,
			},
			count:   8,
			results: []int32{9, 10, 11, 16, 24, 34, 37, 40},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/songs?fmvs=true",
				expectedStatus: http.StatusOK,
			},
			count:   20,
			results: []int32{5, 6, 17, 29, 61, 63, 64, 70, 89},
		},
	}

	testIdList(t, tests, testCfg.e.songs.endpoint, "RetrieveSongs", testCfg.HandleSongs, compareAPIResourceLists[NamedApiResourceList])
}


func TestSubsectionSongs(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations/25/songs",
				expectedStatus: http.StatusOK,
				handler: 		testCfg.HandleLocations,
			},
			count:          11,
			parentResource: h.GetStrPtr("/locations/25"),
			results:        []int32{5, 44, 57, 60, 78, 79, 84, 85, 87, 88, 89},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations/30/songs",
				expectedStatus: http.StatusOK,
				handler: 		testCfg.HandleSublocations,
			},
			count:          4,
			parentResource: h.GetStrPtr("/sublocations/30"),
			results:        []int32{2, 27, 59, 60},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/150/songs",
				expectedStatus: http.StatusOK,
				handler: 		testCfg.HandleAreas,
			},
			count:          4,
			parentResource: h.GetStrPtr("/areas/150"),
			results:        []int32{4, 16, 30, 51},
		},
	}

	testIdList(t, tests, testCfg.e.songs.endpoint, "SubsectionSongs", nil, compareSubResourceLists[NamedAPIResource, SongSub])
}