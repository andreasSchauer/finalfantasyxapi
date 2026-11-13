package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type Sphere struct {
	TestGeneral
	Unique
	item            int32
	createdNode     *apitypes.CreatedNode
	monsters        []TestMonItemAmts
	treasures       map[int32]int32
	shops           []int32
	quests          map[int32]int32
	blitzballPrizes map[int32]int32
}

func (e Sphere) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}
