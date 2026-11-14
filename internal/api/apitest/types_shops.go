package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Shop struct {
	testGeneral
	IDOnly
	area        int32
	category    database.ShopCategory
	preAirship  *testSubShop
	postAirship *testSubShop
}

func (e Shop) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func (e Shop) Compare(test test, got apitypes.Shop) {
	test.t.Helper()
	test.compareExpIDOnly(e.IDOnly, got.ID)
	test.compIDAPIResource("area", apilookups.E.Areas.Endpoint, e.area, got.Area)
	test.compare("category", string(e.category), got.Category)
	test.compTestStructPtrs("pre airship", e.preAirship, got.PreAirship)
	test.compTestStructPtrs("post airship", e.postAirship, got.PostAirship)
}

type testSubShop struct {
	items     []testShopItem
	equipment []testShopEquipment
}

func (e testSubShop) Compare(test test, fieldName string, got apitypes.SubShop) {
	test.t.Helper()
	test.checkTestStructsInSlice(fieldName+" - items", e.items, got.Items)
	test.checkTestStructsInSlice(fieldName+" - equipment", e.equipment, got.Equipment)
}

type testShopItem struct {
	index int
	item  int32
	price int32
}

func newTestShopItem(idx int, itemID, price int32) testShopItem {
	return testShopItem{
		index: idx,
		item:  itemID,
		price: price,
	}
}


func (e testShopItem) GetIndex() int {
	return e.index
}

func (e testShopItem) Compare(test test, fieldName string, got apitypes.ShopItem) {
	test.t.Helper()
	test.compIDAPIResource(fieldName+" - item", apilookups.E.Items.Endpoint, e.item, got.Item)
	test.compare(fieldName+" - price", e.price, got.Price)
}

type testShopEquipment struct {
	index     int
	equipment testFoundEquipment
	price     int32
}

func (e testShopEquipment) GetIndex() int {
	return e.index
}

func (e testShopEquipment) Compare(test test, fieldName string, got apitypes.ShopEquipment) {
	test.t.Helper()
	e.equipment.Compare(test, fieldName+" - found equipment", got.Equipment)
	test.compare(fieldName+" - price", e.price, got.Price)
}
