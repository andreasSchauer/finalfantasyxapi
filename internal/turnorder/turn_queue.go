package turnorder

import (
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/servicetypes"
)

func readyTurnQueue(participants []servicetypes.Participant, rng string) []turnParams {
	var turnQueue []turnParams

	for _, participant := range participants {
		if participant.TickSpeed == 0 || participant.MinICV == nil || participant.MaxICV == nil {
			continue
		}

		params := turnParams{
			Name:        participant.Name,
			PriorityKey: participant.GetKey(),
			Party:       participant.Party,
			Agility:     participant.Agility,
			TickSpeed:   participant.TickSpeed,
		}

		switch rng {
		case string(database.TurnOrderRngBest):
			switch participant.Party {
			case servicetypes.BattlePartyPlayer:
				params.TicksNextTurn = *participant.MinICV

			case servicetypes.BattlePartyOpponent:
				params.TicksNextTurn = *participant.MaxICV
			}

		case string(database.TurnOrderRngWorst):
			switch participant.Party {
			case servicetypes.BattlePartyPlayer:
				params.TicksNextTurn = *participant.MaxICV

			case servicetypes.BattlePartyOpponent:
				params.TicksNextTurn = *participant.MinICV
			}

		case string(database.TurnOrderRngMedian):
			params.TicksNextTurn = (*participant.MinICV + *participant.MaxICV) / 2
		}

		turnQueue = append(turnQueue, params)
	}

	return sortTurnQueue(turnQueue)
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
