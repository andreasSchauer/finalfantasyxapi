package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

func compareShops(test Test, exp Shop, got apitypes.Shop) {
	test.t.Helper()
	test.compareExpIDOnly(exp.IDOnly, got.ID)
	test.compIDAPIResource("area", apilookups.E.Areas.Endpoint, exp.area, got.Area)
	test.compare("category", string(exp.category), got.Category)
	test.compTestStructPtrs("pre airship", exp.preAirship, got.PreAirship, compareSubShops)
	test.compTestStructPtrs("post airship", exp.postAirship, got.PostAirship, compareSubShops)
}

func compareSubShops(test Test, fieldName string, exp TestSubShop, got apitypes.SubShop) {
	test.t.Helper()
	test.checkTestStructsInSlice(fieldName+" - items", exp.items, got.Items, compareShopItems)
	test.checkTestStructsInSlice(fieldName+" - equipment", exp.equipment, got.Equipment, compareShopEquipment)
}

func compareShopItems(test Test, fieldName string, exp TestShopItem, got apitypes.ShopItem) {
	test.t.Helper()
	test.compIDAPIResource(fieldName+" - item", apilookups.E.Items.Endpoint, exp.item, got.Item)
	test.compare(fieldName+" - price", exp.price, got.Price)
}

func compareShopEquipment(test Test, fieldName string, exp TestShopEquipment, got apitypes.ShopEquipment) {
	test.t.Helper()
	compareFoundEquipment(test, fieldName+" - found equipment", exp.equipment, got.Equipment)
	test.compare(fieldName+" - price", exp.price, got.Price)
}
