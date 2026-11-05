package apitest

import "github.com/andreasSchauer/finalfantasyxapi/internal/api"

type expDropChanceResponse struct {
	testGeneral
	api.DropChanceResponse
}

func (e expDropChanceResponse) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareDropChanceResponses(test test, exp expDropChanceResponse, got api.DropChanceResponse) {
	test.t.Helper()
	test.compare("drop chance fin blow", exp.DropChanceFinBlow, got.DropChanceFinBlow)
	test.compare("drop chance no fin blow", exp.DropChanceNoFinBlow, got.DropChanceNoFinBlow)
	test.compStructsWithFn("wanted equipment", exp.WantedEquipment, got.WantedEquipment, compareWantedEquipments)
	test.compStructsWithFn("equipment match chances", exp.EquipmentMatchChances, got.EquipmentMatchChances, compareEquipmentMatchChances)
	test.compStructsWithFn("character chances", exp.CharacterChances, got.CharacterChances, compareCharacterChances)
	test.compStructsWithFn("monster drop chances", exp.MonsterChances, got.MonsterChances, compareMonsterDropChances)
}

func compareWantedEquipments(test test, fieldName string, exp, got api.WantedEquipment) {
	test.t.Helper()
	test.compare(fieldName+" min slots", exp.MinSlots, got.MinSlots)
	test.compare(fieldName+" max slots", exp.MaxSlots, got.MaxSlots)
	test.compareSlices(fieldName+" req abilities", exp.ReqAbilities, got.ReqAbilities)
	test.compare(fieldName+" req empty slots", exp.ReqEmptySlots, got.ReqEmptySlots)
	test.compare(fieldName+" min optional slots", exp.MinOptionalSlots, got.MinOptionalSlots)
	test.compare(fieldName+" max optional slots", exp.MaxOptionalSlots, got.MaxOptionalSlots)
	test.compare(fieldName+" optional slots are", exp.OptionalSlotsAre, got.OptionalSlotsAre)
}

func compareEquipmentMatchChances(test test, fieldName string, exp, got api.EquipmentMatchChances) {
	test.t.Helper()
	test.compare(fieldName+" fin blow", exp.FinBlow, got.FinBlow)
	test.compare(fieldName+" no fin blow", exp.NoFinBlow, got.NoFinBlow)
}

func compareCharacterChances(test test, fieldName string, exp, got api.CharacterChances) {
	test.t.Helper()
	test.compareSlices(fieldName+" eligible chars", exp.EligibleChars, got.EligibleChars)
	test.compare(fieldName+" any fin blow", exp.AnyFinBlow, got.AnyFinBlow)
	test.compare(fieldName+" any no fin blow", exp.AnyNoFinBlow, got.AnyNoFinBlow)
	test.compStructPtrsWithFn(fieldName+" specific character", exp.SpecificChar, got.SpecificChar, compareSpecificCharChances)
}

func compareSpecificCharChances(test test, fieldName string, exp, got api.SpecificCharChances) {
	test.t.Helper()
	test.compare(fieldName+" is eligible", exp.IsEligible, got.IsEligible)
	test.compare(fieldName+" fin blow", exp.FinBlow, got.FinBlow)
	test.compare(fieldName+" no fin blow", exp.NoFinBlow, got.NoFinBlow)
}

func compareMonsterDropChances(test test, fieldName string, exp, got api.MonsterDropChances) {
	test.t.Helper()
	test.compare(fieldName+" monster id", exp.MonsterID, got.MonsterID)
	test.compare(fieldName+" monster name", exp.MonsterName, got.MonsterName)
	test.compare(fieldName+" equipment drop chance", exp.EquipmentDropChance, got.EquipmentDropChance)
	test.compStructSlices(fieldName+" slots chances", exp.SlotsChances, got.SlotsChances)
	test.compStructSlices(fieldName+" shots chances", exp.ShotsChances, got.ShotsChances)
	compareDroppedAutoAbilities(test, fieldName+" auto abilities", exp.AutoAbilities, got.AutoAbilities)
}

func compareDroppedAutoAbilities(test test, fieldName string, exp, got api.DroppedAutoAbilities) {
	test.t.Helper()
	test.compStructPtrsWithFn(fieldName+" forced", exp.Forced, got.Forced, compareEqAbilitiesSimple)
	test.compStructSlicesWithFn(fieldName+" dropped", exp.Dropped, got.Dropped, compareEqAbilitiesChance)
}

func compareEqAbilitiesSimple(test test, fieldName string, exp, got api.EqAbilitySimple) {
	test.t.Helper()
	test.compare(fieldName+" id", exp.ID, got.ID)
	test.compare(fieldName+" name", exp.Name, got.Name)
	test.compareSlices(fieldName+" characters", exp.Characters, got.Characters)
}

func compareEqAbilitiesChance(test test, fieldName string, exp, got api.EqAbilityChance) {
	test.t.Helper()
	compareEqAbilitiesSimple(test, fieldName+" dropped auto ability info", exp.EqAbilitySimple, got.EqAbilitySimple)
	test.compare(fieldName+" chance", exp.Chance, got.Chance)
}
