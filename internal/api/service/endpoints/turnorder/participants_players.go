package turnorder

import (
	"fmt"
	"net/http"

	th "github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/typehelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/service/servicetypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
)

func getPlayerParty(params *servicetypes.TurnOrderParams, maps *participantMaps) ([]servicetypes.Participant, error) {
	playerParty := []servicetypes.Participant{}

	for i, partyMember := range params.Party {
		isDupe := maps.duplicates[partyMember.GetDuplicateKey()]
		if isDupe {
			return nil, errs.NewHTTPError(http.StatusBadRequest, errs.FieldErrorMsg(ts.PfnParty, fmt.Sprintf("duplicate unit id found at index %d. each party member can only appear once", i)), nil)
		}

		participant := getPartyParticipant(params, partyMember)

		playerParty = append(playerParty, participant)
		maps.duplicates[partyMember.GetDuplicateKey()] = true
		maps.namesTotal[participant.GetKey()]++
	}

	return playerParty, nil
}

func getPartyParticipant(params *servicetypes.TurnOrderParams, partyMember servicetypes.TurnOrderParty) servicetypes.Participant {
	unit, _ := lookup.GetResourceByID(partyMember.UnitID, lookup.L.PlayerUnitsID)

	participant := servicetypes.Participant{
		Name:        unit.Name,
		Party:       servicetypes.BattlePartyPlayer,
		Agility:     partyMember.Agility,
		FirstStrike: partyMember.FirstStrike,
		Status:      partyMember.Status,
	}
	participant.AgilityVals = getAglValsChar(participant, params)

	if isSpectralKeeperFormation(params) {
		participant.MinICV, participant.MaxICV = th.GetEqualICVs(0)
	}

	return participant
}

func isSpectralKeeperFormation(params *servicetypes.TurnOrderParams) bool {
	const idFormationSpectral int32 = 253
	return params.FormationID != nil && *params.FormationID == idFormationSpectral
}