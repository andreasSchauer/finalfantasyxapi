package main

import (
	//"net/http"
	"testing"
	//h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)


func TestGetSubquest(t *testing.T) {
	tests := []expSubquest{}

	testSingleResources(t, tests, "GetSubquest", testCfg.HandleSubquests, compareSubquests)
}

func TestRetrieveSubquests(t *testing.T) {
	tests := []expListIDs{}

	testIdList(t, tests, testCfg.e.subquests.endpoint, "RetrieveSubquests", testCfg.HandleSubquests, compareAPIResourceLists[NamedApiResourceList])
}
