package api

type expEquipmentTable struct {
	testGeneral
	expIdOnly
	celestialWeapon			*int32
	specificCharacter		*int32
	requiredAutoAbilities	[]int32
	selectableAutoAbilities	[]testAbilityPool
	emptySlotsAmt			int32
	equipment				[]int32
}

func (t expEquipmentTable) GetTestGeneral() testGeneral {
	return t.testGeneral
}

func compareEquipmentTables(test test, exp expEquipmentTable, got EquipmentTable) {
	compareExpIdOnly(test, exp.expIdOnly, got.ID)
	compIdApiResourcePtrs(test, "celestial weapon", test.cfg.e.celestialWeapons.endpoint, exp.celestialWeapon, got.CelestialWeapon)
	compIdApiResourcePtrs(test, "specific character", test.cfg.e.characters.endpoint, exp.specificCharacter, got.SpecificCharacter)
	checkResIDsInSlice(test, "required auto-abilities", test.cfg.e.autoAbilities.endpoint, exp.requiredAutoAbilities, got.RequiredAutoAbilities)
	compTestStructSlices(test, "selectable auto-abilities", exp.selectableAutoAbilities, got.SelectableAutoAbilities, compareAbilityPools)
	compare(test, "empty slots amount", exp.emptySlotsAmt, got.EmptySlotsAmt)
	checkResIDsInSlice(test, "equipment", test.cfg.e.equipment.endpoint, exp.equipment, got.Equipment)
}

type testAbilityPool struct {
	index			int
	autoAbilities 	[]int32
	reqAmount		int32
}

func compareAbilityPools(test test, fieldName string, exp testAbilityPool, got AbilityPool) {
	checkResIDsInSlice(test, fieldName+" - auto abilities", test.cfg.e.autoAbilities.endpoint, exp.autoAbilities, got.AutoAbilities)
	compare(test, fieldName+" - required amount", exp.reqAmount, got.ReqAmount)
}