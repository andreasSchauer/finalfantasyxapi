package main


type expShop struct {
	testGeneral
	expIdOnly
	area            int32
	category		string
	preAirship		*testSubShop
	postAirship		*testSubShop
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
	items		[]testShopItem
	equipment	[]testShopEquipment
}

func compareSubShops(test test, exp testSubShop, got SubShop) {
	checkTestStructsInSlice(test, "sub shop - items", exp.items, got.Items, compareShopItems)
	checkTestStructsInSlice(test, "sub shop - equipment", exp.equipment, got.Equipment, compareShopEquipment)
}

type testShopItem struct {
	index	int
	item	int32
	price	int32
}

func (t testShopItem) GetIndex() int {
	return t.index
}

func compareShopItems(test test, exp testShopItem, got ShopItem) {
	compIdApiResource(test, "shop item - item", test.cfg.e.shops.endpoint, exp.item, got.Item)
	compare(test, "shop item - price", exp.price, got.Price)
}

type testShopEquipment struct {
	index		int
	equipment 	testFoundEquipment
	price		int32
}

func (t testShopEquipment) GetIndex() int {
	return t.index
}

func compareShopEquipment(test test, exp testShopEquipment, got ShopEquipment) {
	compareFoundEquipment(test, exp.equipment, got.Equipment)
	compare(test, "shop equipment - price", exp.price, got.Price)
}