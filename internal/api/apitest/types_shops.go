package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Shop struct {
	TestGeneral
	IDOnly
	area        int32
	category    database.ShopCategory
	preAirship  *testSubShop
	postAirship *testSubShop
}

func (e Shop) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e Shop) Compare(test Test, got apitypes.Shop) {
	test.T.Helper()
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

func (e testSubShop) Compare(test Test, fieldName string, got apitypes.SubShop) {
	test.T.Helper()
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

func (e testShopItem) Compare(test Test, fieldName string, got apitypes.ShopItem) {
	test.T.Helper()
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

func (e testShopEquipment) Compare(test Test, fieldName string, got apitypes.ShopEquipment) {
	test.T.Helper()
	e.equipment.Compare(test, fieldName+" - found equipment", got.Equipment)
	test.compare(fieldName+" - price", e.price, got.Price)
}
