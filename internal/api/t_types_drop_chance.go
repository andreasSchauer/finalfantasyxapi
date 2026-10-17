package api

type expDropChanceResponse struct {
	testGeneral
	DropChanceResponse
}

func (e expDropChanceResponse) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareDropChanceResponses(test test, exp expDropChanceResponse, got DropChanceResponse) {
	test.t.Helper()
	compStructsWithFn(test, "total drop chances", exp.TotalDropChances, got.TotalDropChances, compareTotalDropChances)
	compStructsWithFn(test, "equipment chances", exp.EquipmentChances, got.EquipmentChances, compareEquipmentChances)
	compStructsWithFn(test, "character chances", exp.CharacterChances, got.CharacterChances, compareCharacterChances)
	compStructsWithFn(test, "monster drop chances", exp.MonsterDropChances, got.MonsterDropChances, compareMonsterDropChances)
}

func compareTotalDropChances(test test, fieldName string, exp, got TotalDropChances) {
	test.t.Helper()
	compare(test, fieldName+" with fin blow", exp.WithFinBlow, got.WithFinBlow)
	compare(test, fieldName+" no fin blow", exp.NoFinBlow, got.NoFinBlow)
}

func compareEquipmentChances(test test, fieldName string, exp, got EquipmentChances) {
	test.t.Helper()
	compare(test, fieldName+" match fin blow", exp.MatchFinBlow, got.MatchFinBlow)
	compare(test, fieldName+" match no fin blow", exp.MatchNoFinBlow, got.MatchNoFinBlow)
}

func compareCharacterChances(test test, fieldName string, exp, got CharacterChances) {
	test.t.Helper()
	compare(test, fieldName+" eligible chars", exp.EligibleChars, got.EligibleChars)
	compare(test, fieldName+" any char fin blow", exp.AnyCharFinBlow, got.AnyCharFinBlow)
	compare(test, fieldName+" any char no fin blow", exp.AnyCharNoFinBlow, got.AnyCharNoFinBlow)
	compStructPtrsWithFn(test, fieldName+" specific character", exp.SpecificChar, got.SpecificChar, compareSpecificCharChances)
}

func compareSpecificCharChances(test test, fieldName string, exp, got SpecificCharChances) {
	test.t.Helper()
	compare(test, fieldName+" char is eligible", exp.CharIsEligible, got.CharIsEligible)
	compare(test, fieldName+" char fin blow", exp.CharFinBlow, got.CharFinBlow)
	compare(test, fieldName+" char no fin blow", exp.CharNoFinBlow, got.CharNoFinBlow)
}

func compareMonsterDropChances(test test, fieldName string, exp, got MonsterDropChances) {
	test.t.Helper()
	compare(test, fieldName+" monster id", exp.MonsterID, got.MonsterID)
	compare(test, fieldName+" monster name", exp.MonsterName, got.MonsterName)
	compare(test, fieldName+" equipment drop chance", exp.EquipmentDropChance, got.EquipmentDropChance)
	compStructSlices(test, fieldName+" slots chances", exp.SlotsChances, got.SlotsChances)
	compStructSlices(test, fieldName+" shots chances", exp.ShotsChances, got.ShotsChances)
	compareDroppedAutoAbilities(test, fieldName+" dropped auto abilities", exp.DroppedAutoAbilities, got.DroppedAutoAbilities)
}

func compareDroppedAutoAbilities(test test, fieldName string, exp, got DroppedAutoAbilities) {
	test.t.Helper()
	compStructPtrsWithFn(test, fieldName+" forced ability", exp.ForcedAbility, got.ForcedAbility, compareDroppedAutoAbilityInfo)
	compStructSlicesWithFn(test, fieldName+" dropped abilities", exp.DroppedAbilities, got.DroppedAbilities, compareDroppedAutoAbility)
}

func compareDroppedAutoAbilityInfo(test test, fieldName string, exp, got DroppedAutoAbilityInfo) {
	test.t.Helper()
	compare(test, fieldName+" id", exp.ID, got.ID)
	compare(test, fieldName+" name", exp.Name, got.Name)
	compareSlices(test, fieldName+" characters", exp.Characters, got.Characters)
}

func compareDroppedAutoAbility(test test, fieldName string, exp, got DroppedAutoAbility) {
	test.t.Helper()
	compareDroppedAutoAbilityInfo(test, fieldName+" dropped auto ability info", exp.DroppedAutoAbilityInfo, got.DroppedAutoAbilityInfo)
	compare(test, fieldName+" chance", exp.Chance, got.Chance)
}