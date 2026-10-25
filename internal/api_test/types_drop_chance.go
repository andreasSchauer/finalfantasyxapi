package apiTest

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
	compare(test, "drop chance fin blow", exp.DropChanceFinBlow, got.DropChanceFinBlow)
	compare(test, "drop chance no fin blow", exp.DropChanceNoFinBlow, got.DropChanceNoFinBlow)
	compStructsWithFn(test, "wanted equipment", exp.WantedEquipment, got.WantedEquipment, compareWantedEquipments)
	compStructsWithFn(test, "equipment match chances", exp.EquipmentMatchChances, got.EquipmentMatchChances, compareEquipmentMatchChances)
	compStructsWithFn(test, "character chances", exp.CharacterChances, got.CharacterChances, compareCharacterChances)
	compStructsWithFn(test, "monster drop chances", exp.MonsterChances, got.MonsterChances, compareMonsterDropChances)
}

func compareWantedEquipments(test test, fieldName string, exp, got api.WantedEquipment) {
	test.t.Helper()
	compare(test, fieldName+" min slots", exp.MinSlots, got.MinSlots)
	compare(test, fieldName+" max slots", exp.MaxSlots, got.MaxSlots)
	compareSlices(test, fieldName+" req abilities", exp.ReqAbilities, got.ReqAbilities)
	compare(test, fieldName+" req empty slots", exp.ReqEmptySlots, got.ReqEmptySlots)
	compare(test, fieldName+" min optional slots", exp.MinOptionalSlots, got.MinOptionalSlots)
	compare(test, fieldName+" max optional slots", exp.MaxOptionalSlots, got.MaxOptionalSlots)
	compare(test, fieldName+" optional slots are", exp.OptionalSlotsAre, got.OptionalSlotsAre)
}

func compareEquipmentMatchChances(test test, fieldName string, exp, got api.EquipmentMatchChances) {
	test.t.Helper()
	compare(test, fieldName+" fin blow", exp.FinBlow, got.FinBlow)
	compare(test, fieldName+" no fin blow", exp.NoFinBlow, got.NoFinBlow)
}

func compareCharacterChances(test test, fieldName string, exp, got api.CharacterChances) {
	test.t.Helper()
	compareSlices(test, fieldName+" eligible chars", exp.EligibleChars, got.EligibleChars)
	compare(test, fieldName+" any fin blow", exp.AnyFinBlow, got.AnyFinBlow)
	compare(test, fieldName+" any no fin blow", exp.AnyNoFinBlow, got.AnyNoFinBlow)
	compStructPtrsWithFn(test, fieldName+" specific character", exp.SpecificChar, got.SpecificChar, compareSpecificCharChances)
}

func compareSpecificCharChances(test test, fieldName string, exp, got api.SpecificCharChances) {
	test.t.Helper()
	compare(test, fieldName+" is eligible", exp.IsEligible, got.IsEligible)
	compare(test, fieldName+" fin blow", exp.FinBlow, got.FinBlow)
	compare(test, fieldName+" no fin blow", exp.NoFinBlow, got.NoFinBlow)
}

func compareMonsterDropChances(test test, fieldName string, exp, got api.MonsterDropChances) {
	test.t.Helper()
	compare(test, fieldName+" monster id", exp.MonsterID, got.MonsterID)
	compare(test, fieldName+" monster name", exp.MonsterName, got.MonsterName)
	compare(test, fieldName+" equipment drop chance", exp.EquipmentDropChance, got.EquipmentDropChance)
	compStructSlices(test, fieldName+" slots chances", exp.SlotsChances, got.SlotsChances)
	compStructSlices(test, fieldName+" shots chances", exp.ShotsChances, got.ShotsChances)
	compareDroppedAutoAbilities(test, fieldName+" auto abilities", exp.AutoAbilities, got.AutoAbilities)
}

func compareDroppedAutoAbilities(test test, fieldName string, exp, got api.DroppedAutoAbilities) {
	test.t.Helper()
	compStructPtrsWithFn(test, fieldName+" forced", exp.Forced, got.Forced, compareEqAbilitiesSimple)
	compStructSlicesWithFn(test, fieldName+" dropped", exp.Dropped, got.Dropped, compareEqAbilitiesChance)
}

func compareEqAbilitiesSimple(test test, fieldName string, exp, got api.EqAbilitySimple) {
	test.t.Helper()
	compare(test, fieldName+" id", exp.ID, got.ID)
	compare(test, fieldName+" name", exp.Name, got.Name)
	compareSlices(test, fieldName+" characters", exp.Characters, got.Characters)
}

func compareEqAbilitiesChance(test test, fieldName string, exp, got api.EqAbilityChance) {
	test.t.Helper()
	compareEqAbilitiesSimple(test, fieldName+" dropped auto ability info", exp.EqAbilitySimple, got.EqAbilitySimple)
	compare(test, fieldName+" chance", exp.Chance, got.Chance)
}
