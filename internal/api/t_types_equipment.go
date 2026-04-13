package api

type expEquipment struct {
	testGeneral
	expUnique
	equipmentTable			int32
	priority				int32
	celestialWeapon			*int32
	requiredAutoAbilities	[]int32
	selectableAutoAbilities	[]testAbilityPool
	emptySlotsAmt			int32
	treasures				[]int32
	shops					[]int32
}

func (e expEquipment) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareEquipment(test test, exp expEquipment, got EquipmentName) {
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compIdApiResource(test, "equipment table", test.cfg.e.equipmentTables.endpoint, exp.equipmentTable, got.EquipmentTable)
	compare(test, "priority", exp.priority, got.Priority)
	compIdApiResourcePtrs(test, "celestial weapon", test.cfg.e.celestialWeapons.endpoint, exp.celestialWeapon, got.CelestialWeapon)
	checkResIDsInSlice(test, "required auto-abilities", test.cfg.e.autoAbilities.endpoint, exp.requiredAutoAbilities, got.RequiredAutoAbilities)
	compTestStructSlices(test, "selectable auto-abilities", exp.selectableAutoAbilities, got.SelectableAutoAbilities, compareAbilityPools)
	compare(test, "empty slots amount", exp.emptySlotsAmt, got.EmptySlotsAmt)
	checkResIDsInSlice(test, "treasures", test.cfg.e.treasures.endpoint, exp.treasures, got.Treasures)
	checkResIDsInSlice(test, "shops", test.cfg.e.shops.endpoint, exp.shops, got.Shops)
}