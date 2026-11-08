package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/servicetypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

type participantMaps struct {
	duplicates   map[string]bool
	namesTotal   map[string]int
	namesCurrent map[string]int
}

func getParticipants(params servicetypes.TurnOrderParams) (servicetypes.TurnOrderParams, []servicetypes.Participant, []servicetypes.Participant, error) {
	maps := &participantMaps{
		namesTotal:   make(map[string]int),
		namesCurrent: make(map[string]int),
		duplicates:   make(map[string]bool),
	}
	params = fetchFormationMons(params)

	playerParty, err := getParticipantsParty(params, maps)
	if err != nil {
		return servicetypes.TurnOrderParams{}, nil, nil, err
	}

	mons, err := getParticipantsMons(params, maps)
	if err != nil {
		return servicetypes.TurnOrderParams{}, nil, nil, err
	}

	opponentParty := createDuplicateNames(mons, maps)

	return params, playerParty, opponentParty, nil
}

func getParticipantsParty(params servicetypes.TurnOrderParams, maps *participantMaps) ([]servicetypes.Participant, error) {
	const idFormationSpectral int32 = 253
	playerParty := []servicetypes.Participant{}

	for i, partyMember := range params.Party {
		isDupe := maps.duplicates[partyMember.GetDuplicateKey()]
		if isDupe {
			return nil, errs.NewHTTPError(http.StatusBadRequest, errs.FieldErrorMsg(ts.PfnParty, fmt.Sprintf("duplicate unit id found at index %d. each party member can only appear once", i)), nil)
		}

		unit, _ := lookup.GetResourceByID(partyMember.UnitID, lookup.L.PlayerUnitsID)

		participant := servicetypes.Participant{
			Name:        unit.Name,
			Party:       servicetypes.BattlePartyPlayer,
			Agility:     partyMember.Agility,
			FirstStrike: partyMember.FirstStrike,
			Status:      partyMember.Status,
		}
		participant.AgilityVals = extractAglTierChar(participant, params)

		if params.FormationID != nil && *params.FormationID == idFormationSpectral {
			participant.MinICV, participant.MaxICV = getEqualICVs(0)
		}

		playerParty = append(playerParty, participant)
		maps.duplicates[partyMember.GetDuplicateKey()] = true
		maps.namesTotal[participant.GetKey()]++
	}

	return playerParty, nil
}

func getParticipantsMons(params servicetypes.TurnOrderParams, maps *participantMaps) ([]servicetypes.Participant, error) {
	monParty := []servicetypes.Participant{}

	for i, mon := range params.Mons {
		isDupe := maps.duplicates[mon.GetDuplicateKey()]
		if isDupe {
			return nil, errs.NewHTTPError(http.StatusBadRequest, errs.FieldErrorMsg(ts.PfnMons, fmt.Sprintf("duplicate monster found at index %d. exact duplicate monsters are not allowed", i)), nil)
		}

		participant, firstTurnAglTier, err := fetchTurnOrderMonster(params, mon)
		if err != nil {
			return nil, err
		}

		participant = populateParticipantMon(params, participant, mon, firstTurnAglTier)

		monParty = append(monParty, participant)
		maps.duplicates[mon.GetDuplicateKey()] = true
		maps.namesTotal[participant.GetKey()]++
	}

	return monParty, nil
}

func fetchTurnOrderMonster(params servicetypes.TurnOrderParams, mon servicetypes.TurnOrderMon) (servicetypes.Participant, *seedtypes.AgilityTier, error) {
	var err error

	if mon.MonsterID == nil {
		return servicetypes.Participant{}, nil, nil
	}

	monster, firstTurnAglTier := assembleMonTurnOrder(*mon.MonsterID, mon.AltState)
	participant := servicetypes.Participant{
		Name:        h.NameToString(monster.Name, monster.Version, nil),
		Agility:     getTurnOrderMonAgility(*mon.MonsterID, monster),
		FirstStrike: monsterHasFirstStrike(monster),
		AltState:    mon.AltState,
	}

	participant.Status, err = fetchMonsterHasteStatus(monster, mon.Status, params.IgnImmunities)
	if err != nil {
		return servicetypes.Participant{}, nil, err
	}

	return participant, firstTurnAglTier, nil
}

func populateParticipantMon(params servicetypes.TurnOrderParams, participant servicetypes.Participant, mon servicetypes.TurnOrderMon, firstTurnAglTier *seedtypes.AgilityTier) servicetypes.Participant {
	if mon.Name != nil {
		participant.Name = *mon.Name
	}

	if mon.Agility != nil {
		participant.Agility = *mon.Agility
	}

	if mon.FirstStrike != nil {
		participant.FirstStrike = *mon.FirstStrike
	}

	if mon.Status != nil {
		participant.Status = mon.Status
	}

	participant.Party = servicetypes.BattlePartyOpponent
	participant.AgilityVals = extractAglTierMon(participant, params, firstTurnAglTier)

	if mon.MonsterID == nil {
		return participant
	}

	if isSpectralKeeperInStoryFight(*mon.MonsterID, params.FormationID) {
		participant.MinICV, participant.MaxICV = getEqualICVs(21)
	}

	return participant
}

func isSpectralKeeperInStoryFight(monID int32, formationPtr *int32) bool {
	const idSpectralKeeper int32 = 193
	const idFormationSpectral int32 = 253
	return formationPtr != nil && *formationPtr == idFormationSpectral && monID == idSpectralKeeper
}

func createDuplicateNames(party []servicetypes.Participant, maps *participantMaps) []servicetypes.Participant {
	for i := range party {
		mon := &party[i]
		key := mon.GetKey()
		maps.namesCurrent[key]++

		if maps.namesTotal[key] <= 1 {
			continue
		}

		mon.Name = fmt.Sprintf("%s (%d)", mon.Name, maps.namesCurrent[key])
	}

	return party
}
