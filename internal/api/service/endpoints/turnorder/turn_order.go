// Package turnorder asdjiasd
package turnorder

import (
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/service/servicetypes"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)


func Handle(params servicetypes.TurnOrderParams) (servicetypes.TurnOrderResponse, error) {
	var err error

	response := servicetypes.TurnOrderResponse{
		TurnsAmt:      params.TurnsAmt,
		IgnFirstTurn:  params.IgnFirstTurn,
		IgnImmunities: params.IgnImmunities,
		RNG:           params.RNG,
	}

	err = getParticipants(&params, &response)
	if err != nil {
		return servicetypes.TurnOrderResponse{}, err
	}
	participants := slices.Concat(response.PlayerParty, response.OpponentParty)
	response.BattleStart = params.BattleStart // forced ambushes change params.BattleStart to "ambush"

	response.TurnOrder = calcTurnOrder(params, participants)

	completeTurnOrderResponse(&response, len(participants), params.TurnsAmt)

	return response, nil
}

func calcTurnOrder(params servicetypes.TurnOrderParams, participants []servicetypes.Participant) []servicetypes.BattleTurn {
	turnQueue := readyTurnQueue(participants, params.RNG)
	turns := []servicetypes.BattleTurn{}
	var ticksTotal int32
	var turnsTotal int32

	for len(turns) != int(params.TurnsAmt) && len(turnQueue) != 0 {
		currentTurn := &turnQueue[0]
		turnsTotal++
		ticksPassed := currentTurn.TicksNextTurn

		if ticksPassed > 0 {
			ticksTotal += ticksPassed
		}

		for i := range turnQueue {
			if ticksPassed <= 0 {
				break
			}

			turnQueue[i].TicksNextTurn -= ticksPassed
		}

		battleTurn := servicetypes.BattleTurn{
			Turn:        turnsTotal,
			CurrentTick: ticksTotal,
			Name:        currentTurn.Name,
			Party:       currentTurn.Party,
			PriorityKey: currentTurn.PriorityKey,
		}
		turns = append(turns, battleTurn)

		currentTurn.TicksNextTurn = currentTurn.TickSpeed * 3
		turnQueue = sortTurnQueue(turnQueue)
	}

	return turns
}

func completeTurnOrderResponse(response *servicetypes.TurnOrderResponse, participantsAmt int, turnsAmt int32) {
	turnCounter := make(map[string]int32, participantsAmt)

	for _, turn := range response.TurnOrder {
		turnCounter[turn.PriorityKey]++
	}

	calcParticipantsTurns(response.PlayerParty, turnCounter, turnsAmt)
	calcParticipantsTurns(response.OpponentParty, turnCounter, turnsAmt)
}

func calcParticipantsTurns(party []servicetypes.Participant, turnCounter map[string]int32, totalTurns int32) {
	for i := range party {
		participant := &party[i]

		turnsReceived := turnCounter[participant.GetKey()]
		participant.TurnsReceived = turnsReceived

		turnsPercent := float64(turnsReceived) / float64(totalTurns)
		participant.TurnsPercent = h.DecimalToPercent(turnsPercent)
	}
}
