package apitesttypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/service/servicetypes"
)

type TurnOrderResponse struct {
	TestGeneral
	servicetypes.TurnOrderResponse
	TestTurnOrder []TestBattleTurn
}

func (e TurnOrderResponse) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e TurnOrderResponse) Compare(test Test, got servicetypes.TurnOrderResponse) {
	test.T.Helper()
	test.compare("turns amt", e.TurnsAmt, got.TurnsAmt)
	test.compare("ign first turn", e.IgnFirstTurn, got.IgnFirstTurn)
	test.compare("ign immunities", e.IgnImmunities, got.IgnImmunities)
	test.compare("battle start", e.BattleStart, got.BattleStart)
	test.compare("rng", e.RNG, got.RNG)
	test.compStructSlices("player party", e.PlayerParty, got.PlayerParty)
	test.compStructSlices("opponent party", got.OpponentParty, got.OpponentParty)
	test.checkTestStructsInSlice("turn order", e.TestTurnOrder, got.TurnOrder)
}

type TestBattleTurn struct {
	Index       int
	Turn        int32
	CurrentTick int32
	Name        string
	Party       servicetypes.BattleParty
}

func (e TestBattleTurn) GetIndex() int {
	return e.Index
}

func (e TestBattleTurn) Compare(test Test, fn string, got servicetypes.BattleTurn) {
	test.T.Helper()
	test.compare(childFn(fn, "turn"), e.Turn, got.Turn)
	test.compare(childFn(fn, "current tick"), e.CurrentTick, got.CurrentTick)
	test.compare(childFn(fn, "name"), e.Name, got.Name)
	test.compare(childFn(fn, "party"), string(e.Party), string(got.Party))
}
