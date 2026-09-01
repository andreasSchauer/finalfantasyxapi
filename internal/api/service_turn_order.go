package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)


type TurnOrderResponse struct {
	URL            	string 			`json:"url"`
	Next			string			`json:"next"`
	IgnFirstTurn	bool			`json:"ign_first_turn"`
	BattleStart		string			`json:"battle_start"`
	PlayerParty		[]Participant	`json:"player_party"`
	OpponentParty	[]Participant	`json:"opponent_party"`
	TurnOrder		[]BattleTurn	`json:"turn_order"`
}

func (r TurnOrderResponse) GetURL() string {
	return r.URL
}

type Participant struct {
	Name			string		`json:"name"`
	Party			BattleParty	`json:"-"`
	Agility			int32		`json:"agility"`
	TickSpeed		int32		`json:"tick_speed"`
	FirstStrike		bool		`json:"first_strike"`
	Status			*string		`json:"status"`
	AltState		*int32		`json:"alt_state,omitempty"`
	TurnsReceived	int32		`json:"turns_received"`
	TurnsPercentage	float64		`json:"turns_percentage"`
	Offset			int32		`json:"starting_tick"`
	MinICV			*int32		`json:"min_icv"`
	MaxICV			*int32		`json:"max_icv"`
}

func (p Participant) getPriorityKey() string {
	if p.Party == battlePartyPlayer {
		return p.Name
	}

	return fmt.Sprintf("mon|%s", p.Name)
}

type BattleParty string

const (
	battlePartyPlayer		BattleParty = "player"
	battlePartyOpponent		BattleParty = "opponent"
)

type BattleTurn struct {
	Name				string		`json:"name"`
	Party				BattleParty	`json:"party"`
	TicksNextTurn		int			`json:"ticks_next_turn"`
	TotalTicksPassed	int			`json:"total_ticks_passed"`
}

func calcTurnOrder(cfg *Config, params TurnOrderParams, url string) (TurnOrderResponse, error) {
	response := TurnOrderResponse{
		URL: url,
		IgnFirstTurn: params.IgnFirstTurn,
		BattleStart: params.BattleStart,
	}

	//priorities := getPriorityMap()
	duplicates := make(map[string]bool)

	// if formation != nil; fetch monsters from monster formation and add to params.Mons

	for _, partyMember := range params.Party {
		_, isDupe := duplicates[partyMember.getParticipantKey()]
		if isDupe {
			return TurnOrderResponse{}, newHTTPError(http.StatusBadRequest, "each party member can only appear once.", nil)
		}
		
		unit, _ := seeding.GetResourceByID(partyMember.ID, cfg.l.PlayerUnitsID)

		participant := Participant{
			Name: unit.Name,
			Party: battlePartyPlayer,
			Agility: partyMember.Agl,
			FirstStrike: partyMember.FS,
			Status: partyMember.Status,
			Offset: partyMember.Offset,
		}
		participant.TickSpeed, participant.MinICV, participant.MaxICV = extractAglTierChar(cfg, partyMember.Agl)

		response.PlayerParty = append(response.PlayerParty, participant)
		duplicates[partyMember.getParticipantKey()] = true
	}

	return response, nil
}

func getPriorityMap() map[string]int {
	prioritySlice := []string{"tidus", "yuna", "auron", "kimahri", "wakka", "lulu", "rikku", "valefor", "ifrit", "ixion", "shiva", "bahamut", "anima", "yojimbo", "cindy", "sandy", "mindy"}
	priorityMap := make(map[string]int, len(prioritySlice))

	for i, name := range prioritySlice {
		priorityMap[name] = i
	}

	return priorityMap
}


func getAllAgilityTiers(cfg *Config) []seeding.AgilityTier {
	tiers := make([]seeding.AgilityTier, len(cfg.l.AgilityTiersID))

	for _, tier := range cfg.l.AgilityTiersID {
		tiers = append(tiers, tier)
	}

	return tiers
}

func getAgilityTier(cfg *Config, agility int32) seeding.AgilityTier {
	tiers := getAllAgilityTiers(cfg)
	var agilityTier seeding.AgilityTier

	for _, tier := range tiers {
		if agility >= tier.MinAgility && agility <= tier.MaxAgility {
			agilityTier = tier
			break
		}
	}

	return agilityTier
}

func extractAglTierChar(cfg *Config, agility int32) (int32, *int32, *int32) {
	agilityTier := getAgilityTier(cfg, agility)
	var minICV *int32

	for _, subtier := range agilityTier.CharacterMinICVs {
		if agility >= subtier.MinAgility && agility <= subtier.MaxAgility {
			minICV = subtier.CharacterMinICV
		}
	}

	return agilityTier.TickSpeed, minICV, agilityTier.CharacterMaxICV
}

func extractAglTierMon(cfg *Config, agility int32) (int32, *int32, *int32) {
	agilityTier := getAgilityTier(cfg, agility)

	return agilityTier.TickSpeed, agilityTier.MonsterMinICV, agilityTier.MonsterMaxICV
}