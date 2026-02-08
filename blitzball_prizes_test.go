package main

import (
	//"net/http"
	"testing"
	//h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type expBlitzballPosition struct {
	testGeneral
	expIdOnly
	category		string
	slot			string
	items			[]testPossibleItem
}

func (e expBlitzballPosition) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareBlitzballPositions(test test, exp expBlitzballPosition, got BlitzballPrize) {
	compareExpIdOnly(test, exp.expIdOnly, got.ID)
	compare(test, "category", exp.category, got.Category)
	compare(test, "slot", exp.slot, got.Slot)
	compTestStructSlices(test, "items", exp.items, got.Items, comparePossibleItems)
}

func TestGetBlitzballPosition(t *testing.T) {
	tests := []expBlitzballPosition{}

	testSingleResources(t, tests, "GetBlitzballPrize", testCfg.HandleBlitzballPrizes, compareBlitzballPositions)
}

func TestRetrieveBlitzballPositions(t *testing.T) {
	tests := []expListIDs{}

	testIdList(t, tests, testCfg.e.blitzballPrizes.endpoint, "RetrieveBlitzballPositions", testCfg.HandleBlitzballPrizes, compareAPIResourceLists[UnnamedApiResourceList])
}
