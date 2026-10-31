package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api"
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

func compareShops(test test, exp expShop, got api.Shop) {
	test.t.Helper()
	compareExpIDOnly(test, exp.expIDOnly, got.ID)
	compIDAPIResource(test, "area", test.cfg.E.Areas.Endpoint, exp.area, got.Area)
	compare(test, "category", string(exp.category), got.Category)
	compTestStructPtrs(test, "pre airship", exp.preAirship, got.PreAirship, compareSubShops)
	compTestStructPtrs(test, "post airship", exp.postAirship, got.PostAirship, compareSubShops)
}

type testSubShop struct {
	items     []testShopItem
	equipment []testShopEquipment
}

func compareSubShops(test test, fieldName string, exp testSubShop, got api.SubShop) {
	test.t.Helper()
	checkTestStructsInSlice(test, fieldName+" - items", exp.items, got.Items, compareShopItems)
	checkTestStructsInSlice(test, fieldName+" - equipment", exp.equipment, got.Equipment, compareShopEquipment)
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

func compareShopItems(test test, fieldName string, exp testShopItem, got api.ShopItem) {
	test.t.Helper()
	compIDAPIResource(test, fieldName+" - item", test.cfg.E.Items.Endpoint, exp.item, got.Item)
	compare(test, fieldName+" - price", exp.price, got.Price)
}

type testShopEquipment struct {
	index     int
	equipment testFoundEquipment
	price     int32
}

func (t testShopEquipment) GetIndex() int {
	return t.index
}

func compareShopEquipment(test test, fieldName string, exp testShopEquipment, got api.ShopEquipment) {
	test.t.Helper()
	compareFoundEquipment(test, fieldName+" - found equipment", exp.equipment, got.Equipment)
	compare(test, fieldName+" - price", exp.price, got.Price)
}
