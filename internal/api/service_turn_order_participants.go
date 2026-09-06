package api

import (
	"fmt"
	"net/http"
	"slices"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type Participant struct {
	Name    string      `json:"name"`
	Party   BattleParty `json:"-"`
	Agility int32       `json:"agility"`
	AgilityVals
	FirstStrike     bool    `json:"first_strike"`
	Status          *string `json:"status"`
	AltState        *int32  `json:"alt_state,omitempty"`
	TurnsReceived   int32   `json:"turns_received"`
	TurnsPercentage float64 `json:"turns_percentage"`
}

func (p Participant) getKey() string {
	if p.Party == battlePartyPlayer {
		return p.Name
	}

	return fmt.Sprintf("mon|%s", p.Name)
}

type BattleParty string

const (
	battlePartyPlayer   BattleParty = "player"
	battlePartyOpponent BattleParty = "opponent"
)

type participantMaps struct {
	duplicates   map[string]bool
	namesTotal   map[string]int
	namesCurrent map[string]int
}

func getParticipants(cfg *Config, params TurnOrderParams) ([]Participant, []Participant, error) {
	maps := &participantMaps{
		namesTotal:   make(map[string]int),
		namesCurrent: make(map[string]int),
		duplicates:   make(map[string]bool),
	}
	params = fetchFormationMons(cfg, params)

	playerParty, err := getParticipantsParty(cfg, params, maps)
	if err != nil {
		return nil, nil, err
	}

	mons, params, err := getParticipantsMons(cfg, params, maps)
	if err != nil {
		return nil, nil, err
	}

	monsCustom, err := getParticipantsMonsCustom(cfg, params, maps)
	if err != nil {
		return nil, nil, err
	}

	opponentParty := createDuplicateNames(slices.Concat(mons, monsCustom), maps)

	return playerParty, opponentParty, nil
}

func getParticipantsParty(cfg *Config, params TurnOrderParams, maps *participantMaps) ([]Participant, error) {
	const idFormationSpectral int32 = 253
	var playerParty []Participant

	for _, partyMember := range params.Party {
		_, isDupe := maps.duplicates[partyMember.getDuplicateKey()]
		if isDupe {
			return nil, newHTTPError(http.StatusBadRequest, "each party member can only appear once.", nil)
		}

		unit, _ := seeding.GetResourceByID(partyMember.ID, cfg.l.PlayerUnitsID)

		participant := Participant{
			Name:        unit.Name,
			Party:       battlePartyPlayer,
			Agility:     partyMember.Agl,
			FirstStrike: partyMember.FS,
			Status:      partyMember.Status,
		}
		participant.AgilityVals = extractAglTierChar(cfg, participant, params)

		if params.Formation != nil && *params.Formation == idFormationSpectral {
			*participant.MinICV = 0
			*participant.MaxICV = 0
		}

		playerParty = append(playerParty, participant)
		maps.duplicates[partyMember.getDuplicateKey()] = true
		maps.namesTotal[participant.getKey()]++
	}

	return playerParty, nil
}

func getParticipantsMons(cfg *Config, params TurnOrderParams, maps *participantMaps) ([]Participant, TurnOrderParams, error) {
	const idSpectralKeeper int32 = 193
	const idFormationSpectral int32 = 253
	var monParty []Participant

	for _, mon := range params.Mons {
		_, isDupe := maps.duplicates[mon.getDuplicateKey()]
		if isDupe {
			return nil, TurnOrderParams{}, newHTTPError(http.StatusBadRequest, "exact duplicate mons are not allowed", nil)
		}

		monster, err := getMonsterFromJson(cfg, mon)
		if err != nil {
			return nil, TurnOrderParams{}, err
		}

		agilityBS := getBaseStat(cfg, "agility", monster.BaseStats)
		agility := agilityBS.Value
		firstStrike := monHasFirstStrike(monster)

		if mon.AglOverride != nil {
			agility = *mon.AglOverride
		}

		agility = handleMonEdgeCases(mon, agility)

		participant := Participant{
			Name:        h.NameToString(monster.Name, monster.Version, nil),
			Party:       battlePartyOpponent,
			Agility:     agility,
			FirstStrike: firstStrike,
			AltState:    mon.AltState,
		}
		participant.Status, err = fetchMonsterStatus(monster, mon.Status)
		if err != nil {
			return nil, TurnOrderParams{}, err
		}

		participant.AgilityVals = extractAglTierMon(cfg, participant, params)

		if params.Formation != nil && *params.Formation == idFormationSpectral && mon.ID == idSpectralKeeper {
			*participant.MinICV = 21
			*participant.MaxICV = 21
		}

		monParty = append(monParty, participant)
		maps.duplicates[mon.getDuplicateKey()] = true
		maps.namesTotal[participant.getKey()]++
	}

	return monParty, params, nil
}

func getParticipantsMonsCustom(cfg *Config, params TurnOrderParams, maps *participantMaps) ([]Participant, error) {
	var monParty []Participant

	for _, mon := range params.MonsCustom {
		_, isDupe := maps.duplicates[mon.getDuplicateKey()]
		if isDupe {
			return nil, newHTTPError(http.StatusBadRequest, "exact duplicate mons are not allowed", nil)
		}

		participant := Participant{
			Name:        mon.Name,
			Party:       battlePartyOpponent,
			Agility:     mon.Agl,
			FirstStrike: mon.FS,
			Status:      mon.Status,
		}
		participant.AgilityVals = extractAglTierMon(cfg, participant, params)

		monParty = append(monParty, participant)
		maps.duplicates[mon.getDuplicateKey()] = true
		maps.namesTotal[participant.getKey()]++
	}

	return monParty, nil
}

func createDuplicateNames(party []Participant, maps *participantMaps) []Participant {
	for i := range party {
		mon := &party[i]
		key := mon.getKey()
		maps.namesCurrent[key]++

		if maps.namesTotal[key] <= 1 {
			continue
		}

		mon.Name = fmt.Sprintf("%s (%d)", mon.Name, maps.namesCurrent[key])
	}

	return party
}
