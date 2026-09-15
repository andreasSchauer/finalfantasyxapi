package api

import (
	"slices"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type TurnOrderResponse struct {
	TurnsAmt      int32         `json:"turns_amt"`
	IgnFirstTurn  bool          `json:"ign_first_turn"`
	BattleStart   string        `json:"battle_start"`
	RNG           string        `json:"rng"`
	PlayerParty   []Participant `json:"player_party"`
	OpponentParty []Participant `json:"opponent_party"`
	TurnOrder     []BattleTurn  `json:"turn_order"`
}

type TurnParams struct {
	Name          string
	Party         BattleParty
	PriorityKey   string
	Agility       int32
	TickSpeed     int32
	TicksNextTurn int32
}

type BattleTurn struct {
	Turn        int32       `json:"turn"`
	CurrentTick int32       `json:"current_tick"`
	Name        string      `json:"name"`
	Party       BattleParty `json:"party"`
	PriorityKey string      `json:"-"`
}

func handleTurnOrder(cfg *Config, params TurnOrderParams) (TurnOrderResponse, error) {
	var err error

	response := TurnOrderResponse{
		TurnsAmt:     params.TurnsAmt,
		IgnFirstTurn: params.IgnFirstTurn,
		BattleStart:  params.BattleStart,
		RNG:          params.RNG,
	}

	response.PlayerParty, response.OpponentParty, err = getParticipants(cfg, params)
	if err != nil {
		return TurnOrderResponse{}, err
	}
	participants := slices.Concat(response.PlayerParty, response.OpponentParty)

	response.TurnOrder = calcTurnOrder(params, participants)

	response = completeTurnOrderResponse(response, len(participants), params.TurnsAmt)

	return response, nil
}

func calcTurnOrder(params TurnOrderParams, participants []Participant) []BattleTurn {
	turnQueue := readyTurnQueue(participants, params.RNG)
	var turns []BattleTurn
	var ticksTotal int32
	var turnsTotal int32

	for {
		if len(turns) == int(params.TurnsAmt) {
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

		battleTurn := BattleTurn{
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


func completeTurnOrderResponse(response TurnOrderResponse, participantsAmt int, turnsAmt int32) TurnOrderResponse {
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

func calcParticipantTurns(participant Participant, turnCounter map[string]int32, totalTurns int32) Participant {
	turnsReceived := turnCounter[participant.getKey()]
	participant.TurnsReceived = turnsReceived

	turnsPercent := float64(turnsReceived) / float64(totalTurns)
	participant.TurnsPercent = h.DecimalToPercent(turnsPercent)

	return participant
}
