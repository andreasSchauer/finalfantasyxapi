package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/servicetypes"
)

type expTurnOrderResponse struct {
	testGeneral
	servicetypes.TurnOrderResponse
	testTurnOrder []testBattleTurn
}

func (e expTurnOrderResponse) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareTurnOrderResponses(test test, exp expTurnOrderResponse, got servicetypes.TurnOrderResponse) {
	test.t.Helper()
	test.compare("turns amt", exp.TurnsAmt, got.TurnsAmt)
	test.compare("ign first turn", exp.IgnFirstTurn, got.IgnFirstTurn)
	test.compare("ign immunities", exp.IgnImmunities, got.IgnImmunities)
	test.compare("battle start", exp.BattleStart, got.BattleStart)
	test.compare("rng", exp.RNG, got.RNG)
	test.compTestStructSlices("player party", exp.PlayerParty, got.PlayerParty, compareParticipants)
	test.compTestStructSlices("opponent party", exp.OpponentParty, got.OpponentParty, compareParticipants)
	test.checkTestStructsInSlice("turn order", exp.testTurnOrder, got.TurnOrder, compareBattleTurns)
}

func compareParticipants(test test, fieldName string, exp, got servicetypes.Participant) {
	test.t.Helper()
	test.compare(fieldName+" name", exp.Name, got.Name)
	test.compare(fieldName+" party", string(exp.Party), string(got.Party))
	test.compare(fieldName+" agility", exp.Agility, got.Agility)
	test.compare(fieldName+" tick speed", exp.TickSpeed, got.TickSpeed)
	test.compare(fieldName+" min icv", exp.MinICV, got.MinICV)
	test.compare(fieldName+" max icv", exp.MaxICV, got.MaxICV)
	test.compare(fieldName+" first strike", exp.FirstStrike, got.FirstStrike)
	test.compare(fieldName+" status", exp.Status, got.Status)
	test.compare(fieldName+" alt state", exp.AltState, got.AltState)
	test.compare(fieldName+" turns received", exp.TurnsReceived, got.TurnsReceived)
	test.compare(fieldName+" turns percent", exp.TurnsPercent, got.TurnsPercent)
}

type testBattleTurn struct {
	index       int
	turn        int32
	currentTick int32
	name        string
	party       servicetypes.BattleParty
}

func (t testBattleTurn) GetIndex() int {
	return t.index
}

func compareBattleTurns(test test, fieldName string, exp testBattleTurn, got servicetypes.BattleTurn) {
	test.t.Helper()
	test.compare(fieldName+" turn", exp.turn, got.Turn)
	test.compare(fieldName+" current tick", exp.currentTick, got.CurrentTick)
	test.compare(fieldName+" name", exp.name, got.Name)
	test.compare(fieldName+" party", string(exp.party), string(got.Party))
}
