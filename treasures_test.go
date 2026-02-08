package main

import (
	//"net/http"
	"testing"
	//h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type expTreasure struct {
	testGeneral
	expIdOnly
	area            int32
	isPostAirship   bool
	isAnimaTreasure bool
	treasureType    int32
	lootType        int32
	gilAmount       *int32
	items           []testItemAmount
	equipment       *testFoundEquipment
}

func (e expTreasure) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareTreasures(test test, exp expTreasure, got Treasure) {
	compareExpIdOnly(test, exp.expIdOnly, got.ID)
	compIdApiResource(test, "area", testCfg.e.areas.endpoint, exp.area, got.Area)
	compare(test, "is post airship", exp.isPostAirship, got.IsPostAirship)
	compare(test, "is anima treasure", exp.isAnimaTreasure, got.IsAnimaTreasure)
	compIdApiResource(test, "treasure type", testCfg.e.treasureType.endpoint, exp.treasureType, got.TreasureType)
	compIdApiResource(test, "loot type", testCfg.e.lootType.endpoint, exp.lootType, got.LootType)
	compare(test, "gil amount", exp.gilAmount, got.GilAmount)
	compTestStructSlices(test, "items", exp.items, got.Items, compareItemAmounts)
	compTestStructPtrs(test, "equipment", exp.equipment, got.Equipment, compareFoundEquipment)
}

type testFoundEquipment struct {
	equipmentName    int32
	abilities        []int32
	emptySlotsAmount int32
}

func compareFoundEquipment(test test, exp testFoundEquipment, got FoundEquipment) {
	enEndpoint := test.cfg.e.equipment.endpoint
	aaEndpoint := test.cfg.e.autoAbilities.endpoint

	compIdApiResource(test, "found equipment - equipment name", enEndpoint, exp.equipmentName, got.EquipmentName)
	compareResListTest(test, rltIDs("found equipment - abilities", aaEndpoint, exp.abilities, got.Abilities))
	compare(test, "found equipment - empty slots amount", exp.emptySlotsAmount, got.EmptySlotsAmount)
}

func TestGetTreasure(t *testing.T) {
	tests := []expTreasure{}

	testSingleResources(t, tests, "GetTreasure", testCfg.HandleTreasures, compareTreasures)
}

func TestRetrieveTreasures(t *testing.T) {
	tests := []expListIDs{}

	testIdList(t, tests, testCfg.e.treasures.endpoint, "RetrieveTreasures", testCfg.HandleTreasures, compareAPIResourceLists[UnnamedApiResourceList])
}
