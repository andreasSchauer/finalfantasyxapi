package main

import (
	//"net/http"
	"testing"
	//h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)


func TestGetBlitzballPrize(t *testing.T) {
	tests := []expBlitzballPrize{}

	testSingleResources(t, tests, "GetBlitzballPrize", testCfg.HandleBlitzballPrizes, compareBlitzballPrizes)
}

func TestRetrieveBlitzballPrizes(t *testing.T) {
	tests := []expListIDs{}

	testIdList(t, tests, testCfg.e.blitzballPrizes.endpoint, "RetrieveBlitzballPositions", testCfg.HandleBlitzballPrizes, compareAPIResourceLists[UnnamedApiResourceList])
}
