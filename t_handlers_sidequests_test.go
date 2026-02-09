package main

import (
	//"net/http"
	"testing"
	//h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)


func TestGetSidequest(t *testing.T) {
	tests := []expSidequest{}

	testSingleResources(t, tests, "GetSidequest", testCfg.HandleSidequests, compareSidequests)
}

func TestRetrieveSidequests(t *testing.T) {
	tests := []expListIDs{}

	testIdList(t, tests, testCfg.e.sidequests.endpoint, "RetrieveSidequests", testCfg.HandleSidequests, compareAPIResourceLists[NamedApiResourceList])
}
