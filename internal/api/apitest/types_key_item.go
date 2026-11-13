package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type KeyItem struct {
	TestGeneral
	Unique
	untypedItem     int32
	category        database.KeyItemCategory
	celestialWeapon *int32
	primer          *int32
	areas           []int32
	treasures       []int32
	quests          []int32
}

func (e KeyItem) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}
