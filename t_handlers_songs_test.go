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
				requestURL: "/api/songs/a",
				expectedStatus: http.StatusBadRequest,
				expectedErr: "",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL: 	"/api/songs/0",
				expectedStatus: http.StatusOK,
				dontCheck: 		map[string]bool{},
				expLengths: 	map[string]int{},
			},
			expUnique: newExpUnique(0, ""),
			composer: h.GetStrPtr(""),
			arranger: h.GetStrPtr(""),
			durationInSeconds: 0,
			canLoop: false,
			backgroundMusic: []testBackgroundMusic{
				{
					index: 0,
					replacesEncounterMusic: false,
					areas: []int32{},
				},
			},
			cues: []testCue{
				{
					index: 0,
					triggerArea: h.GetInt32Ptr(0),
					includedAreas: []int32{},
					replacesEncounterMusic: false,
					replacesBGMusic: h.GetStrPtr(""),
				},
			},
			bossFights: []int32{},
			fmvs: []int32{},
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
			count:   0,
			results: []int32{},
		},
	}

	testIdList(t, tests, testCfg.e.songs.endpoint, "RetrieveSongs", testCfg.HandleSongs, compareAPIResourceLists[NamedApiResourceList])
}