package turnorder

import (
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/service/servicetypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)


type turnParams struct {
	Name          string
	Party         servicetypes.BattleParty
	PriorityKey   string
	Agility       int32
	TickSpeed     int32
	TicksNextTurn int32
}

func readyTurnQueue(participants []servicetypes.Participant, rng string) []turnParams {
	var turnQueue []turnParams

	for _, participant := range participants {
		if participant.TickSpeed == 0 || participant.MinICV == nil || participant.MaxICV == nil {
			continue
		}

		params := getTurnParams(participant, rng)

		turnQueue = append(turnQueue, params)
	}

	return sortTurnQueue(turnQueue)
}


func getTurnParams(p servicetypes.Participant, rng string) turnParams {
	params := turnParams{
		Name:        p.Name,
		PriorityKey: p.GetKey(),
		Party:       p.Party,
		Agility:     p.Agility,
		TickSpeed:   p.TickSpeed,
	}

	switch rng {
	case string(database.TurnOrderRngBest):
		switch p.Party {
		case servicetypes.BattlePartyPlayer:
			params.TicksNextTurn = *p.MinICV

		case servicetypes.BattlePartyOpponent:
			params.TicksNextTurn = *p.MaxICV
		}

	case string(database.TurnOrderRngWorst):
		switch p.Party {
		case servicetypes.BattlePartyPlayer:
			params.TicksNextTurn = *p.MaxICV

		case servicetypes.BattlePartyOpponent:
			params.TicksNextTurn = *p.MinICV
		}

	case string(database.TurnOrderRngMedian):
		params.TicksNextTurn = (*p.MinICV + *p.MaxICV) / 2
	}

	return params
}

func sortTurnQueue(turnQueue []turnParams) []turnParams {
	slices.SortStableFunc(turnQueue, sortTurnQueueTicks)

	return turnQueue
}

func sortTurnQueueTicks(a, b turnParams) int {
	if a.TicksNextTurn < b.TicksNextTurn {
		return -1
	}

	if a.TicksNextTurn > b.TicksNextTurn {
		return 1
	}

	return sortTurnQueueAgility(a, b)
}

func sortTurnQueueAgility(a, b turnParams) int {
	if a.Agility > b.Agility {
		return -1
	}

	if a.Agility < b.Agility {
		return 1
	}

	return sortTurnQueuePriority(a, b)
}

func sortTurnQueuePriority(a, b turnParams) int {
	priorities := getPriorityMap()

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

func getPriorityMap() map[string]int {
	prioritySlice := []string{"tidus", "yuna", "auron", "kimahri", "wakka", "lulu", "rikku", "valefor", "ifrit", "ixion", "shiva", "bahamut", "anima", "yojimbo", "cindy", "sandy", "mindy"}
	priorityMap := make(map[string]int, len(prioritySlice))

	for i, name := range prioritySlice {
		priorityMap[name] = i
	}

	return priorityMap
}
