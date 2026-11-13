package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Shop struct {
	TestGeneral
	IDOnly
	area        int32
	category    database.ShopCategory
	preAirship  *TestSubShop
	postAirship *TestSubShop
}

func (e Shop) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

type TestSubShop struct {
	items     []TestShopItem
	equipment []TestShopEquipment
}

type TestShopItem struct {
	index int
	item  int32
	price int32
}

func (t TestShopItem) GetIndex() int {
	return t.index
}

func newTestShopItem(idx int, itemID, price int32) TestShopItem {
	return TestShopItem{
		index: idx,
		item:  itemID,
		price: price,
	}
}

type TestShopEquipment struct {
	index     int
	equipment TestFoundEquipment
	price     int32
}

func (t TestShopEquipment) GetIndex() int {
	return t.index
}
