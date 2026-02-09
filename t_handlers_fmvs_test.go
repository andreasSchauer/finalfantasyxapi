package main

import (
	//"net/http"
	"testing"
	//h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)




func TestGetFMV(t *testing.T) {
	tests := []expFMV{}

	testSingleResources(t, tests, "GetFMV", testCfg.HandleFMVs, compareFMVs)
}

func TestRetrieveFMVs(t *testing.T) {
	tests := []expListIDs{}

	testIdList(t, tests, testCfg.e.fmvs.endpoint, "RetrieveFMVs", testCfg.HandleFMVs, compareAPIResourceLists[NamedApiResourceList])
}
