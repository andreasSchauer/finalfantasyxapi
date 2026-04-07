package api

type expMasterItem struct {
	testGeneral
	expUnique
	itemType			int32
	typedItem			string
	obtainableFrom		ObtainableFrom
}

func (i expMasterItem) GetTestGeneral() testGeneral {
	return i.testGeneral
}

func compareMasterItems(test test, exp expMasterItem, got MasterItem) {
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compIdApiResource(test, "item type", test.cfg.e.itemType.endpoint, exp.itemType, got.Type)
	compPathApiResource(test, "typed item", exp.typedItem, got.TypedItem)
	compStructs(test, "obtainable from", exp.obtainableFrom, got.ObtainableFrom)
}