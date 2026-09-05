package api

import (
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type TurnOrderResponse struct {
	URL           string        `json:"url"`
	Next          string        `json:"next"`
	IgnFirstTurn  bool          `json:"ign_first_turn"`
	BattleStart   string        `json:"battle_start"`
	RNG           string        `json:"rng"`
	PlayerParty   []Participant `json:"player_party"`
	OpponentParty []Participant `json:"opponent_party"`
	TurnOrder     []BattleTurn  `json:"turn_order"`
}

func (r TurnOrderResponse) GetURL() string {
	return r.URL
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

func handleTurnOrder(cfg *Config, params TurnOrderParams, url string) (TurnOrderResponse, error) {
	var err error

	response := TurnOrderResponse{
		URL:          url,
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
	priorities := getPriorityMap()
	turnQueue := readyTurnQueue(participants, priorities, params.RNG)
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

			turn := &turnQueue[i]
			turn.TicksNextTurn -= ticksPassed
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
		turnQueue = sortTurnQueue(turnQueue, priorities)
	}

	return turns
}

func getPriorityMap() map[string]int {
	prioritySlice := []string{"tidus", "yuna", "auron", "kimahri", "wakka", "lulu", "rikku", "valefor", "ifrit", "ixion", "shiva", "bahamut", "anima", "yojimbo", "cindy", "sandy", "mindy"}
	priorityMap := make(map[string]int, len(prioritySlice))

	for i, name := range prioritySlice {
		priorityMap[name] = i
	}

	return priorityMap
}

func readyTurnQueue(participants []Participant, priorities map[string]int, rng string) []TurnParams {
	var turnQueue []TurnParams

	for _, participant := range participants {
		if participant.TickSpeed == 0 || participant.MinICV == nil || participant.MaxICV == nil {
			continue
		}

		params := TurnParams{
			Name:        participant.Name,
			PriorityKey: participant.getKey(),
			Party:       participant.Party,
			Agility:     participant.Agility,
			TickSpeed:   participant.TickSpeed,
		}

		switch rng {
		case string(database.TurnOrderRngBest):
			switch participant.Party {
			case battlePartyPlayer:
				params.TicksNextTurn = *participant.MinICV

			case battlePartyOpponent:
				params.TicksNextTurn = *participant.MaxICV
			}

		case string(database.TurnOrderRngWorst):
			switch participant.Party {
			case battlePartyPlayer:
				params.TicksNextTurn = *participant.MaxICV

			case battlePartyOpponent:
				params.TicksNextTurn = *participant.MinICV
			}

		case string(database.TurnOrderRngMedian):
			params.TicksNextTurn = (*participant.MinICV + *participant.MaxICV) / 2
		}

		turnQueue = append(turnQueue, params)
	}

	return sortTurnQueue(turnQueue, priorities)
}

func sortTurnQueue(turnQueue []TurnParams, priorities map[string]int) []TurnParams {
	slices.SortStableFunc(turnQueue, func(a, b TurnParams) int {
		if a.TicksNextTurn < b.TicksNextTurn {
			return -1
		}

		if a.TicksNextTurn > b.TicksNextTurn {
			return 1
		}

		return sortTurnQueueAgility(a, b, priorities)
	})

	return turnQueue
}

func sortTurnQueueAgility(a, b TurnParams, priorities map[string]int) int {
	if a.Agility > b.Agility {
		return -1
	}

	if a.Agility < b.Agility {
		return 1
	}

	return sortTurnQueuePriority(a, b, priorities)
}

func sortTurnQueuePriority(a, b TurnParams, priorities map[string]int) int {
	aVal, ok := priorities[a.PriorityKey]
	if !ok {
		aVal = 99
	}
	bVal, ok := priorities[b.PriorityKey]
	if !ok {
		bVal = 100
	}

	if aVal < bVal {
		return -1
	}

	if aVal > bVal {
		return 1
	}

	return 0
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

	turnsPercentage := float64(turnsReceived) / float64(totalTurns)
	participant.TurnsPercentage = h.FloatRound(turnsPercentage, 2)

	return participant
}
