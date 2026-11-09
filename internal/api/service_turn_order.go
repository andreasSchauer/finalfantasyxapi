package api

import (
	"slices"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/servicetypes"
)

type TurnParams struct {
	Name          string
	Party         servicetypes.BattleParty
	PriorityKey   string
	Agility       int32
	TickSpeed     int32
	TicksNextTurn int32
}

func handleTurnOrder(params servicetypes.TurnOrderParams) (servicetypes.TurnOrderResponse, error) {
	var err error

	response := servicetypes.TurnOrderResponse{
		TurnsAmt:      params.TurnsAmt,
		IgnFirstTurn:  params.IgnFirstTurn,
		IgnImmunities: params.IgnImmunities,
		RNG:           params.RNG,
	}

	params, response.PlayerParty, response.OpponentParty, err = getParticipants(params)
	if err != nil {
		return servicetypes.TurnOrderResponse{}, err
	}
	participants := slices.Concat(response.PlayerParty, response.OpponentParty)
	response.BattleStart = params.BattleStart

	response.TurnOrder = calcTurnOrder(params, participants)

	response = CompleteTurnOrderResponse(response, len(participants), params.TurnsAmt)

	return response, nil
}

func calcTurnOrder(params servicetypes.TurnOrderParams, participants []servicetypes.Participant) []servicetypes.BattleTurn {
	turnQueue := readyTurnQueue(participants, params.RNG)
	turns := []servicetypes.BattleTurn{}
	var ticksTotal int32
	var turnsTotal int32

	for {
		if len(turns) == int(params.TurnsAmt) || len(turnQueue) == 0 {
			break
		}

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

func CompleteTurnOrderResponse(response servicetypes.TurnOrderResponse, participantsAmt int, turnsAmt int32) servicetypes.TurnOrderResponse {
	turnCounter := make(map[string]int32, participantsAmt)

	for _, turn := range response.TurnOrder {
		turnCounter[turn.PriorityKey]++
	}

	for i, player := range response.PlayerParty {
		response.PlayerParty[i] = calcParticipantTurns(player, turnCounter, turnsAmt)
	}

	for i, mon := range response.OpponentParty {
		response.OpponentParty[i] = calcParticipantTurns(mon, turnCounter, turnsAmt)
	}

	return response
}

func calcParticipantTurns(participant servicetypes.Participant, turnCounter map[string]int32, totalTurns int32) servicetypes.Participant {
	turnsReceived := turnCounter[participant.GetKey()]
	participant.TurnsReceived = turnsReceived

	turnsPercent := float64(turnsReceived) / float64(totalTurns)
	participant.TurnsPercent = h.DecimalToPercent(turnsPercent)

	return participant
}
