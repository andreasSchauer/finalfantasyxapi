package main

import (
	//"net/http"
	"testing"
	//h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)


func TestGetTreasure(t *testing.T) {
	tests := []expTreasure{}

	testSingleResources(t, tests, "GetTreasure", testCfg.HandleTreasures, compareTreasures)
}

func TestRetrieveTreasures(t *testing.T) {
	tests := []expListIDs{}

	testIdList(t, tests, testCfg.e.treasures.endpoint, "RetrieveTreasures", testCfg.HandleTreasures, compareAPIResourceLists[UnnamedApiResourceList])
}
