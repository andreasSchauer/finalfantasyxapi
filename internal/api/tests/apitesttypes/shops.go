package apitesttypes

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
	PreAirship  *TestSubShop
	PostAirship *TestSubShop
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

type TestSubShop struct {
	Items     []TestShopItem
	Equipment []TestShopEquipment
}

func (e TestSubShop) Compare(test Test, fn string, got apitypes.SubShop) {
	test.T.Helper()
	test.checkTestStructsInSlice(childFn(fn, "items"), e.Items, got.Items)
	test.checkTestStructsInSlice(childFn(fn, "equipment"), e.Equipment, got.Equipment)
}

type TestShopItem struct {
	Index int
	Item  int32
	Price int32
}

func NewTestShopItem(idx int, itemID, price int32) TestShopItem {
	return TestShopItem{
		Index: idx,
		Item:  itemID,
		Price: price,
	}
}

func (e TestShopItem) GetIndex() int {
	return e.Index
}

func (e TestShopItem) Compare(test Test, fn string, got apitypes.ShopItem) {
	test.T.Helper()
	test.compIDAPIResource(childFn(fn, "item"), apilookups.E.Items.Endpoint, e.Item, got.Item)
	test.compare(childFn(fn, "price"), e.Price, got.Price)
}

type TestShopEquipment struct {
	Index     int
	Equipment TestFoundEquipment
	Price     int32
}

func (e TestShopEquipment) GetIndex() int {
	return e.Index
}

func (e TestShopEquipment) Compare(test Test, fn string, got apitypes.ShopEquipment) {
	test.T.Helper()
	e.Equipment.Compare(test, childFn(fn, "found equipment"), got.Equipment)
	test.compare(childFn(fn, "price"), e.Price, got.Price)
}
