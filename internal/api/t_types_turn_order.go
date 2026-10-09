package api

type expTurnOrderResponse struct {
	testGeneral
	TurnOrderResponse
	testTurnOrder		[]testBattleTurn
}

func (e expTurnOrderResponse) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareTurnOrderResponses(test test, exp expTurnOrderResponse, got TurnOrderResponse) {
	test.t.Helper()
	compare(test, "turns amt", exp.TurnsAmt, got.TurnsAmt)
	compare(test, "ign first turn", exp.IgnFirstTurn, got.IgnFirstTurn)
	compare(test, "battle start", exp.BattleStart, got.BattleStart)
	compare(test, "rng", exp.RNG, got.RNG)
	compTestStructSlices(test, "player party", exp.PlayerParty, got.PlayerParty, compareParticipants)
	compTestStructSlices(test, "opponent party", exp.OpponentParty, got.OpponentParty, compareParticipants)
	checkTestStructsInSlice(test, "turn order", exp.testTurnOrder, got.TurnOrder, compareBattleTurns)
}

func compareParticipants(test test, fieldName string, exp Participant, got Participant) {
	test.t.Helper()
	compare(test, fieldName+" name", exp.Name, got.Name)
	compare(test, fieldName+" party", string(exp.Party), string(got.Party))
	compare(test, fieldName+" agility", exp.Agility, got.Agility)
	compare(test, fieldName+" tick speed", exp.TickSpeed, got.TickSpeed)
	compare(test, fieldName+" min icv", exp.MinICV, got.MinICV)
	compare(test, fieldName+" max icv", exp.MaxICV, got.MaxICV)
	compare(test, fieldName+" first strike", exp.FirstStrike, got.FirstStrike)
	compare(test, fieldName+" status", exp.Status, got.Status)
	compare(test, fieldName+" alt state", exp.AltState, got.AltState)
	compare(test, fieldName+" turns received", exp.TurnsReceived, got.TurnsReceived)
	compare(test, fieldName+" turns percent", exp.TurnsPercent, got.TurnsPercent)
}

type testBattleTurn struct {
	index		int
	turn        int32
	currentTick int32
	name        string
	party       BattleParty
}

func (t testBattleTurn) GetIndex() int {
	return t.index
}

func compareBattleTurns(test test, fieldName string, exp testBattleTurn, got BattleTurn) {
	test.t.Helper()
	compare(test, fieldName+" turn", exp.turn, got.Turn)
	compare(test, fieldName+" current tick", exp.currentTick, got.CurrentTick)
	compare(test, fieldName+" name", exp.name, got.Name)
	compare(test, fieldName+" party", string(exp.party), string(got.Party))
}