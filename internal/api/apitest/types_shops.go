package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Shop struct {
	TestGeneral
	IDOnly
	Area        int32
	Category    database.ShopCategory
	PreAirship  *testSubShop
	PostAirship *testSubShop
}

func (e Shop) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e Shop) Compare(test Test, got apitypes.Shop) {
	test.T.Helper()
	test.compareExpIDOnly(e.IDOnly, got.ID)
	test.compIDAPIResource("area", apilookups.E.Areas.Endpoint, e.Area, got.Area)
	test.compare("category", string(e.Category), got.Category)
	test.compTestStructPtrs("pre airship", e.PreAirship, got.PreAirship)
	test.compTestStructPtrs("post airship", e.PostAirship, got.PostAirship)
}

type testSubShop struct {
	Items     []testShopItem
	Equipment []testShopEquipment
}

func (e testSubShop) Compare(test Test, fieldName string, got apitypes.SubShop) {
	test.T.Helper()
	test.checkTestStructsInSlice(fieldName+" - items", e.Items, got.Items)
	test.checkTestStructsInSlice(fieldName+" - equipment", e.Equipment, got.Equipment)
}

type testShopItem struct {
	Index int
	Item  int32
	Price int32
}

func newTestShopItem(idx int, itemID, price int32) testShopItem {
	return testShopItem{
		Index: idx,
		Item:  itemID,
		Price: price,
	}
}

func (e testShopItem) GetIndex() int {
	return e.Index
}

func (e testShopItem) Compare(test Test, fieldName string, got apitypes.ShopItem) {
	test.T.Helper()
	test.compIDAPIResource(fieldName+" - item", apilookups.E.Items.Endpoint, e.Item, got.Item)
	test.compare(fieldName+" - price", e.Price, got.Price)
}

type testShopEquipment struct {
	Index     int
	Equipment testFoundEquipment
	Price     int32
}

func (e testShopEquipment) GetIndex() int {
	return e.Index
}

func (e testShopEquipment) Compare(test Test, fieldName string, got apitypes.ShopEquipment) {
	test.T.Helper()
	e.Equipment.Compare(test, fieldName+" - found equipment", got.Equipment)
	test.compare(fieldName+" - price", e.Price, got.Price)
}
