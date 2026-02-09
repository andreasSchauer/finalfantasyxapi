package main

import (
	//"net/http"
	"testing"
	//h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type expFMV struct {
	testGeneral
	expUnique
	area		int32
	song		*int32
}

func (e expFMV) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareFMVs(test test, exp expFMV, got FMV) {
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compIdApiResource(test, "area", test.cfg.e.areas.endpoint, exp.area, got.Area)
	compIdApiResourcePtrs(test, "song", test.cfg.e.songs.endpoint, exp.song, got.Song)
}



func TestGetFMV(t *testing.T) {
	tests := []expFMV{}

	testSingleResources(t, tests, "GetFMV", testCfg.HandleFMVs, compareFMVs)
}

func TestRetrieveFMVs(t *testing.T) {
	tests := []expListIDs{}

	testIdList(t, tests, testCfg.e.fmvs.endpoint, "RetrieveFMVs", testCfg.HandleFMVs, compareAPIResourceLists[NamedApiResourceList])
}
