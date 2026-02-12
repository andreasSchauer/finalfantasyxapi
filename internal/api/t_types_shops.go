package api

type expShop struct {
	testGeneral
	expIdOnly
	area        int32
	category    string
	preAirship  *testSubShop
	postAirship *testSubShop
}

func (e expShop) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareShops(test test, exp expShop, got Shop) {
	compareExpIdOnly(test, exp.expIdOnly, got.ID)
	compIdApiResource(test, "area", test.cfg.e.areas.endpoint, exp.area, got.Area)
	compare(test, "category", exp.category, got.Category)
	compTestStructPtrs(test, "pre airship", exp.preAirship, got.PreAirship, compareSubShops)
	compTestStructPtrs(test, "post airship", exp.postAirship, got.PostAirship, compareSubShops)
}

type testSubShop struct {
	items     []testShopItem
	equipment []testShopEquipment
}

func compareSubShops(test test, fieldName string, exp testSubShop, got SubShop) {
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

func compareShopItems(test test, fieldName string, exp testShopItem, got ShopItem) {
	compIdApiResource(test, fieldName+" - item", test.cfg.e.items.endpoint, exp.item, got.Item)
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

func compareShopEquipment(test test, fieldName string, exp testShopEquipment, got ShopEquipment) {
	compareFoundEquipment(test, fieldName+" - found equipment", exp.equipment, got.Equipment)
	compare(test, fieldName+" - price", exp.price, got.Price)
}
