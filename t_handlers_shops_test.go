package main

import (
	//"net/http"
	"testing"
	//h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)



func TestGetShop(t *testing.T) {
	tests := []expShop{}

	testSingleResources(t, tests, "GetShop", testCfg.HandleShops, compareShops)
}

func TestRetrieveShops(t *testing.T) {
	tests := []expListIDs{}

	testIdList(t, tests, testCfg.e.shops.endpoint, "RetrieveShops", testCfg.HandleShops, compareAPIResourceLists[UnnamedApiResourceList])
}
