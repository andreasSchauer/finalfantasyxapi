package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/service/servicetypes"
)

type TurnOrderResponse struct {
	TestGeneral
	servicetypes.TurnOrderResponse
	testTurnOrder []TestBattleTurn
}

func (e TurnOrderResponse) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

type TestBattleTurn struct {
	index       int
	turn        int32
	currentTick int32
	name        string
	party       servicetypes.BattleParty
}

func (t TestBattleTurn) GetIndex() int {
	return t.index
}
