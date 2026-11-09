package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type expShop struct {
	testGeneral
	expIDOnly
	area        int32
	category    database.ShopCategory
	preAirship  *testSubShop
	postAirship *testSubShop
}

func (e expShop) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareShops(test test, exp expShop, got apitypes.Shop) {
	test.t.Helper()
	test.compareExpIDOnly(exp.expIDOnly, got.ID)
	test.compIDAPIResource("area", apilookups.E.Areas.Endpoint, exp.area, got.Area)
	test.compare("category", string(exp.category), got.Category)
	test.compTestStructPtrs("pre airship", exp.preAirship, got.PreAirship, compareSubShops)
	test.compTestStructPtrs("post airship", exp.postAirship, got.PostAirship, compareSubShops)
}

type testSubShop struct {
	items     []testShopItem
	equipment []testShopEquipment
}

func compareSubShops(test test, fieldName string, exp testSubShop, got apitypes.SubShop) {
	test.t.Helper()
	test.checkTestStructsInSlice(fieldName+" - items", exp.items, got.Items, compareShopItems)
	test.checkTestStructsInSlice(fieldName+" - equipment", exp.equipment, got.Equipment, compareShopEquipment)
}

type testShopItem struct {
	index int
	item  int32
	price int32
}

func (t testShopItem) GetIndex() int {
	return t.index
}

func newTestShopItem(idx int, itemID, price int32) testShopItem {
	return testShopItem{
		index: idx,
		item:  itemID,
		price: price,
	}
}

func compareShopItems(test test, fieldName string, exp testShopItem, got apitypes.ShopItem) {
	test.t.Helper()
	test.compIDAPIResource(fieldName+" - item", apilookups.E.Items.Endpoint, exp.item, got.Item)
	test.compare(fieldName+" - price", exp.price, got.Price)
}

type testShopEquipment struct {
	index     int
	equipment testFoundEquipment
	price     int32
}

func (t testShopEquipment) GetIndex() int {
	return t.index
}

func compareShopEquipment(test test, fieldName string, exp testShopEquipment, got apitypes.ShopEquipment) {
	test.t.Helper()
	compareFoundEquipment(test, fieldName+" - found equipment", exp.equipment, got.Equipment)
	test.compare(fieldName+" - price", exp.price, got.Price)
}
