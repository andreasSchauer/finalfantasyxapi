package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Treasure struct {
	TestGeneral
	IDOnly
	area            int32
	availability    database.AvailabilityType
	isAnimaTreasure bool
	treasureType    database.TreasureType
	lootType        database.LootType
	gilAmount       *int32
	items           []TestResAmount[apiresource.Typed]
	equipment       *TestFoundEquipment
}

func (e Treasure) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

type TestFoundEquipment struct {
	equipmentName    int32
	abilities        []int32
	emptySlotsAmount int32
}
