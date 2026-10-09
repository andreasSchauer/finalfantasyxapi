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
	compareTotalDropChances(test, "total drop chances", exp.TotalDropChances, got.TotalDropChances)
	compareEquipmentChances(test, "equipment chances", exp.EquipmentChances, got.EquipmentChances)
	compareCharacterChances(test, "character chances", exp.CharacterChances, got.CharacterChances)
}

func compareTotalDropChances(test test, fieldName string, exp TotalDropChances, got TotalDropChances) {
	test.t.Helper()
	compare(test, fieldName+" with fin blow", exp.WithFinBlow, got.WithFinBlow)
	compare(test, fieldName+" no fin blow", exp.NoFinBlow, got.NoFinBlow)
}

func compareEquipmentChances(test test, fieldName string, exp EquipmentChances, got EquipmentChances) {
	test.t.Helper()
	compare(test, fieldName+" monster drop", exp.MonsterDrop, got.MonsterDrop)
	compare(test, fieldName+" match fin blow", exp.MatchFinBlow, got.MatchFinBlow)
	compare(test, fieldName+" match no fin blow", exp.MatchNoFinBlow, got.MatchNoFinBlow)
}

func compareCharacterChances(test test, fieldName string, exp CharacterChances, got CharacterChances) {
	test.t.Helper()
	compare(test, fieldName+" eligible chars", exp.EligibleChars, got.EligibleChars)
	compare(test, fieldName+" any char fin blow", exp.AnyCharFinBlow, got.AnyCharFinBlow)
	compare(test, fieldName+" any char no fin blow", exp.AnyCharNoFinBlow, got.AnyCharNoFinBlow)
	compare(test, fieldName+" char fin blow", exp.CharFinBlow, got.CharFinBlow)
	compare(test, fieldName+" char no fin blow", exp.CharNoFinBlow, got.CharNoFinBlow)
}