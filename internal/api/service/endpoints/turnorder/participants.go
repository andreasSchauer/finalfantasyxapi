package turnorder


import (
	"fmt"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/service/servicetypes"
)


type participantMaps struct {
	duplicates   map[string]bool
	namesTotal   map[string]int
	namesCurrent map[string]int
}

func getParticipants(params *servicetypes.TurnOrderParams, response *servicetypes.TurnOrderResponse) error {
	var err error

	maps := &participantMaps{
		namesTotal:   make(map[string]int),
		namesCurrent: make(map[string]int),
		duplicates:   make(map[string]bool),
	}

	fetchFormationMons(params)
	response.BattleStart = params.BattleStart

	response.PlayerParty, err = getPlayerParty(params, maps)
	if err != nil {
		return err
	}

	response.OpponentParty, err = getOpponentParty(params, maps)
	if err != nil {
		return err
	}

	createDuplicateNames(response.OpponentParty, maps)

	return nil
}

func createDuplicateNames(party []servicetypes.Participant, maps *participantMaps) {
	for i := range party {
		mon := &party[i]
		key := mon.GetKey()
		maps.namesCurrent[key]++

		if maps.namesTotal[key] <= 1 {
			continue
		}

		mon.Name = fmt.Sprintf("%s (%d)", mon.Name, maps.namesCurrent[key])
	}
}