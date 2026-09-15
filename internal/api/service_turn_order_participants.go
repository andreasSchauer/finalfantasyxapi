package api

import (
	"fmt"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type Participant struct {
	Name    string      `json:"name"`
	Party   BattleParty `json:"-"`
	Agility int32       `json:"agility"`
	AgilityVals
	FirstStrike   bool    `json:"first_strike"`
	Status        *string `json:"status"`
	AltState      *int32  `json:"alt_state,omitempty"`
	TurnsReceived int32   `json:"turns_received"`
	TurnsPercent  float64 `json:"turns_percent"`
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

	mons, err := getParticipantsMons(cfg, params, maps)
	if err != nil {
		return nil, nil, err
	}

	opponentParty := createDuplicateNames(mons, maps)

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
			Agility:     partyMember.Agility,
			FirstStrike: partyMember.FS,
			Status:      partyMember.Status,
		}
		participant.AgilityVals = extractAglTierChar(cfg, participant, params)

		if params.Formation != nil && *params.Formation == idFormationSpectral {
			participant.MinICV, participant.MaxICV = getEqualICVs(0)
		}

		playerParty = append(playerParty, participant)
		maps.duplicates[partyMember.getDuplicateKey()] = true
		maps.namesTotal[participant.getKey()]++
	}

	return playerParty, nil
}

func getParticipantsMons(cfg *Config, params TurnOrderParams, maps *participantMaps) ([]Participant, error) {
	var monParty []Participant
	var firstTurnAglTier *seeding.AgilityTier

	for _, mon := range params.Mons {
		_, isDupe := maps.duplicates[mon.getDuplicateKey()]
		if isDupe {
			return nil, newHTTPError(http.StatusBadRequest, "exact duplicate mons are not allowed", nil)
		}

		var participant Participant
		var err error

		participant, firstTurnAglTier, err = fetchTurnOrderMonster(cfg, params, mon)
		if err != nil {
			return nil, err
		}

		participant = populateParticipantMon(cfg, params, participant, mon, firstTurnAglTier)

		monParty = append(monParty, participant)
		maps.duplicates[mon.getDuplicateKey()] = true
		maps.namesTotal[participant.getKey()]++
	}

	return monParty, nil
}

func fetchTurnOrderMonster(cfg *Config, params TurnOrderParams, mon turnOrderMon) (Participant, *seeding.AgilityTier, error) {
	if mon.ID == nil {
		return Participant{}, nil, nil
	}

	monster, firstTurnAglTier, err := getTurnOrderMonFromJson(cfg, *mon.ID, mon.AltState)
	if err != nil {
		return Participant{}, nil, err
	}

	participant := Participant{
		Name:        h.NameToString(monster.Name, monster.Version, nil),
		Agility:     getTurnOrderMonAgilityNew(cfg, *mon.ID, monster),
		FirstStrike: monsterHasFirstStrike(monster),
		AltState:    mon.AltState,
	}
	participant.Status, err = fetchMonsterStatus(params, monster, mon.Status)
	if err != nil {
		return Participant{}, nil, err
	}

	return participant, firstTurnAglTier, nil
}

func populateParticipantMon(cfg *Config, params TurnOrderParams, participant Participant, mon turnOrderMon, firstTurnAglTier *seeding.AgilityTier) Participant {
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

	participant.Party = battlePartyOpponent
	participant.AgilityVals = extractAglTierMon(cfg, participant, params, firstTurnAglTier)

	if isSpectralKeeperInStoryFight(*mon.ID, params.Formation) {
		participant.MinICV, participant.MaxICV = getEqualICVs(21)
	}

	return participant
}

func isSpectralKeeperInStoryFight(monID int32, formationPtr *int32) bool {
	const idSpectralKeeper int32 = 193
	const idFormationSpectral int32 = 253
	return formationPtr != nil && *formationPtr == idFormationSpectral && monID == idSpectralKeeper
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
