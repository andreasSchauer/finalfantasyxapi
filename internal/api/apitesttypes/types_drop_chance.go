package apitesttypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/service/servicetypes"
)

type DropChanceResponse struct {
	TestGeneral
	servicetypes.DropChanceResponse
}

func (e DropChanceResponse) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e DropChanceResponse) Compare(test Test, got servicetypes.DropChanceResponse) {
	test.T.Helper()
	test.compare("drop chance fin blow", e.DropChanceFinBlow, got.DropChanceFinBlow)
	test.compare("drop chance no fin blow", e.DropChanceNoFinBlow, got.DropChanceNoFinBlow)
	test.compStructs("wanted equipment", e.WantedEquipment, got.WantedEquipment)
	test.compStructs("equipment match chances", e.EquipmentMatchChances, got.EquipmentMatchChances)
	test.compStructs("character chances", e.CharacterChances, got.CharacterChances)
	test.compStructs("monster drop chances", e.MonsterChances, got.MonsterChances)
}
