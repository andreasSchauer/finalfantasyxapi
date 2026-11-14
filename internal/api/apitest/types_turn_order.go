package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/service/servicetypes"
)

type TurnOrderResponse struct {
	testGeneral
	servicetypes.TurnOrderResponse
	testTurnOrder []testBattleTurn
}

func (e TurnOrderResponse) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func (e TurnOrderResponse) Compare(test test, got servicetypes.TurnOrderResponse) {
	test.t.Helper()
	test.compare("turns amt", e.TurnsAmt, got.TurnsAmt)
	test.compare("ign first turn", e.IgnFirstTurn, got.IgnFirstTurn)
	test.compare("ign immunities", e.IgnImmunities, got.IgnImmunities)
	test.compare("battle start", e.BattleStart, got.BattleStart)
	test.compare("rng", e.RNG, got.RNG)
	test.compStructSlices("player party", e.PlayerParty, got.PlayerParty)
	test.compStructSlices("opponent party", got.OpponentParty, got.OpponentParty)
	test.checkTestStructsInSlice("turn order", e.testTurnOrder, got.TurnOrder)
}

type testBattleTurn struct {
	index       int
	turn        int32
	currentTick int32
	name        string
	party       servicetypes.BattleParty
}

func (e testBattleTurn) GetIndex() int {
	return e.index
}

func (e testBattleTurn) Compare(test test, fieldName string, got servicetypes.BattleTurn) {
	test.t.Helper()
	test.compare(fieldName+" turn", e.turn, got.Turn)
	test.compare(fieldName+" current tick", e.currentTick, got.CurrentTick)
	test.compare(fieldName+" name", e.name, got.Name)
	test.compare(fieldName+" party", string(e.party), string(got.Party))
}
