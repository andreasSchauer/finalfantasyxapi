package main

import (
	//"net/http"
	"testing"
	//h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)



func TestGetArenaCreation(t *testing.T) {
	tests := []expArenaCreation{}

	testSingleResources(t, tests, "GetArenaCreation", testCfg.HandleArenaCreations, compareArenaCreations)
}

func TestRetrieveArenaCreations(t *testing.T) {
	tests := []expListIDs{}

	testIdList(t, tests, testCfg.e.arenaCreations.endpoint, "RetrieveArenaCreations", testCfg.HandleArenaCreations, compareAPIResourceLists[NamedApiResourceList])
}
