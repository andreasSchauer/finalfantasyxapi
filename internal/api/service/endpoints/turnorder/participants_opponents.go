package turnorder

import (
	"fmt"
	"net/http"
	
	th "github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/typehelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/errs"
	sh "github.com/andreasSchauer/finalfantasyxapi/internal/api/service/servicehelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/service/servicetypes"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func getOpponentParty(params *servicetypes.TurnOrderParams, maps *participantMaps) ([]servicetypes.Participant, error) {
	monParty := []servicetypes.Participant{}

	for i, mon := range params.Mons {
		isDupe := maps.duplicates[mon.GetDuplicateKey()]
		if isDupe {
			return nil, errs.NewHTTPError(http.StatusBadRequest, errs.FieldErrorMsg(ts.PfnMons, fmt.Sprintf("duplicate monster found at index %d. exact duplicate monsters are not allowed", i)), nil)
		}

		participant, err := getMonParticipant(params, mon)
		if err != nil {
			return nil, err
		}

		monParty = append(monParty, participant)
		maps.duplicates[mon.GetDuplicateKey()] = true
		maps.namesTotal[participant.GetKey()]++
	}

	return monParty, nil
}

func getMonParticipant(params *servicetypes.TurnOrderParams, mon servicetypes.TurnOrderMon) (servicetypes.Participant, error) {
	var participant servicetypes.Participant

	if mon.Custom != nil {
		participant = servicetypes.Participant{
			Name:        mon.Custom.Name,
			Agility:     mon.Custom.Agility,
			FirstStrike: mon.Custom.FirstStrike,
			Status:      mon.Custom.Status,
			Party:       servicetypes.BattlePartyOpponent,
		}
		participant.AgilityVals = getAglValsMon(participant, params, nil)
	}

	if mon.Lookup != nil {
		var firstTurnAglTier *seedtypes.AgilityTier
		var err error

		participant, firstTurnAglTier, err = fetchTurnOrderMonster(params, *mon.Lookup)
		if err != nil {
			return servicetypes.Participant{}, err
		}
		participant = populateParticipantMonLookup(params, participant, *mon.Lookup, firstTurnAglTier)
	}

	return participant, nil
}

func fetchTurnOrderMonster(params *servicetypes.TurnOrderParams, mon servicetypes.TurnOrderMonLookup) (servicetypes.Participant, *seedtypes.AgilityTier, error) {
	var err error
	var statusPtr *string

	monster, firstTurnAglTier := assembleMonTurnOrder(mon.MonsterID, mon.AltState)

	participant := servicetypes.Participant{
		Name:        h.NameToString(monster.Name, monster.Version, nil),
		Agility:     getTurnOrderMonAgility(mon.MonsterID, monster),
		FirstStrike: sh.MonsterHasFirstStrike(monster),
		AltState:    mon.AltState,
	}

	if mon.Overrides != nil {
		statusPtr = mon.Overrides.Status
	}

	participant.Status, err = fetchMonsterHasteStatus(monster, statusPtr, params.IgnImmunities)
	if err != nil {
		return servicetypes.Participant{}, nil, err
	}

	return participant, firstTurnAglTier, nil
}


func populateParticipantMonLookup(params *servicetypes.TurnOrderParams, participant servicetypes.Participant, mon servicetypes.TurnOrderMonLookup, firstTurnAglTier *seedtypes.AgilityTier) servicetypes.Participant {
	if mon.Overrides != nil {
		participant = applyParticipantOverrides(participant, *mon.Overrides)
	}

	participant.Party = servicetypes.BattlePartyOpponent
	participant.AgilityVals = getAglValsMon(participant, params, firstTurnAglTier)

	if isSpectralKeeperInStoryFight(mon.MonsterID, params.FormationID) {
		participant.MinICV, participant.MaxICV = th.GetEqualICVs(21)
	}

	return participant
}

func applyParticipantOverrides(participant servicetypes.Participant, mon servicetypes.TurnOrderMonOverrides) servicetypes.Participant {
	if mon.Agility != nil {
		participant.Agility = *mon.Agility
	}

	if mon.FirstStrike != nil {
		participant.FirstStrike = *mon.FirstStrike
	}

	if mon.Status != nil {
		participant.Status = sh.OverrideNonePtr(mon.Status)
	}

	return participant
}

func isSpectralKeeperInStoryFight(monID int32, formationPtr *int32) bool {
	const idSpectralKeeper int32 = 193
	const idFormationSpectral int32 = 253
	return formationPtr != nil && *formationPtr == idFormationSpectral && monID == idSpectralKeeper
}
