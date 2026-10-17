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
	compare(test, "drop chance fin blow", exp.DropChanceFinBlow, got.DropChanceFinBlow)
	compare(test, "drop chance no fin blow", exp.DropChanceNoFinBlow, got.DropChanceNoFinBlow)
	compStructsWithFn(test, "equipment match chances", exp.EquipmentMatchChances, got.EquipmentMatchChances, compareEquipmentMatchChances)
	compStructsWithFn(test, "character chances", exp.CharacterChances, got.CharacterChances, compareCharacterChances)
	compStructsWithFn(test, "monster drop chances", exp.MonsterChances, got.MonsterChances, compareMonsterDropChances)
}

func compareEquipmentMatchChances(test test, fieldName string, exp, got EquipmentMatchChances) {
	test.t.Helper()
	compare(test, fieldName+" fin blow", exp.FinBlow, got.FinBlow)
	compare(test, fieldName+" no fin blow", exp.NoFinBlow, got.NoFinBlow)
}

func compareCharacterChances(test test, fieldName string, exp, got CharacterChances) {
	test.t.Helper()
	compare(test, fieldName+" eligible chars", exp.EligibleChars, got.EligibleChars)
	compare(test, fieldName+" any fin blow", exp.AnyFinBlow, got.AnyFinBlow)
	compare(test, fieldName+" any no fin blow", exp.AnyNoFinBlow, got.AnyNoFinBlow)
	compStructPtrsWithFn(test, fieldName+" specific character", exp.SpecificChar, got.SpecificChar, compareSpecificCharChances)
}

func compareSpecificCharChances(test test, fieldName string, exp, got SpecificCharChances) {
	test.t.Helper()
	compare(test, fieldName+" is eligible", exp.IsEligible, got.IsEligible)
	compare(test, fieldName+" fin blow", exp.FinBlow, got.FinBlow)
	compare(test, fieldName+" no fin blow", exp.NoFinBlow, got.NoFinBlow)
}

func compareMonsterDropChances(test test, fieldName string, exp, got MonsterDropChances) {
	test.t.Helper()
	compare(test, fieldName+" monster id", exp.MonsterID, got.MonsterID)
	compare(test, fieldName+" monster name", exp.MonsterName, got.MonsterName)
	compare(test, fieldName+" equipment drop chance", exp.EquipmentDropChance, got.EquipmentDropChance)
	compStructSlices(test, fieldName+" slots chances", exp.SlotsChances, got.SlotsChances)
	compStructSlices(test, fieldName+" shots chances", exp.ShotsChances, got.ShotsChances)
	compareDroppedAutoAbilities(test, fieldName+" auto abilities", exp.AutoAbilities, got.AutoAbilities)
}

func compareDroppedAutoAbilities(test test, fieldName string, exp, got DroppedAutoAbilities) {
	test.t.Helper()
	compStructPtrsWithFn(test, fieldName+" forced", exp.Forced, got.Forced, compareEqAbilitiesSimple)
	compStructSlicesWithFn(test, fieldName+" dropped", exp.Dropped, got.Dropped, compareEqAbilitiesChance)
}

func compareEqAbilitiesSimple(test test, fieldName string, exp, got EqAbilitySimple) {
	test.t.Helper()
	compare(test, fieldName+" id", exp.ID, got.ID)
	compare(test, fieldName+" name", exp.Name, got.Name)
	compareSlices(test, fieldName+" characters", exp.Characters, got.Characters)
}

func compareEqAbilitiesChance(test test, fieldName string, exp, got EqAbilityChance) {
	test.t.Helper()
	compareEqAbilitiesSimple(test, fieldName+" dropped auto ability info", exp.EqAbilitySimple, got.EqAbilitySimple)
	compare(test, fieldName+" chance", exp.Chance, got.Chance)
}
