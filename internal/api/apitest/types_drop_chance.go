package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/service/servicetypes"
)

type DropChanceResponse struct {
	testGeneral
	servicetypes.DropChanceResponse
}

func (e DropChanceResponse) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareDropChanceResponses(test test, exp DropChanceResponse, got servicetypes.DropChanceResponse) {
	test.t.Helper()
	test.compare("drop chance fin blow", exp.DropChanceFinBlow, got.DropChanceFinBlow)
	test.compare("drop chance no fin blow", exp.DropChanceNoFinBlow, got.DropChanceNoFinBlow)
	test.compStructs("wanted equipment", exp.WantedEquipment, got.WantedEquipment)
	test.compStructs("equipment match chances", exp.EquipmentMatchChances, got.EquipmentMatchChances)
	test.compStructs("character chances", exp.CharacterChances, got.CharacterChances)
	test.compStructs("monster drop chances", exp.MonsterChances, got.MonsterChances)
}