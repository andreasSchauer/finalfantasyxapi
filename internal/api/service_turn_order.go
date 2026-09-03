package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
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
		participant.TickSpeed, participant.MinICV, participant.MaxICV = extractAglTierChar(cfg, partyMember, params.IgnFirstTurn)

		response.PlayerParty = append(response.PlayerParty, participant)
		duplicates[partyMember.getParticipantKey()] = true
	}

	params.Mons = fetchFormationMons(cfg, params.Formation, params.Mons)

	for _, mon := range params.Mons {
		_, isDupe := duplicates[mon.getParticipantKey()]
		if isDupe {
			return TurnOrderResponse{}, newHTTPError(http.StatusBadRequest, "exact duplicate mons are not allowed", nil)
		}

		monster, err := getMonsterFromJson(cfg, mon)
		if err != nil {
			return TurnOrderResponse{}, err
		}

		agilityBS := getBaseStat(cfg, "agility", monster.BaseStats)
		agility := agilityBS.Value
		firstStrike := monHasFirstStrike(monster)

		if mon.AglOverride != nil {
			agility = *mon.AglOverride
		}
		
		participant := Participant{
			Name: h.NameToString(monster.Name, monster.Version, nil),
			Party: battlePartyOpponent,
			Agility: agility,
			FirstStrike: firstStrike,
			AltState: mon.AltState,
			Offset: mon.Offset,
		}
		participant.TickSpeed, participant.MinICV, participant.MaxICV = extractAglTierMon(cfg, mon, agility, firstStrike, params.IgnFirstTurn)

		participant.Status, err = fetchMonsterStatus(monster, mon.Status)
		if err != nil {
			return TurnOrderResponse{}, err
		}

		response.OpponentParty = append(response.OpponentParty, participant)
		duplicates[mon.getParticipantKey()] = true
	}

	for _, mon := range params.MonsCustom {
		_, isDupe := duplicates[mon.getParticipantKey()]
		if isDupe {
			return TurnOrderResponse{}, newHTTPError(http.StatusBadRequest, "exact duplicate mons are not allowed", nil)
		}

		participant := Participant{
			Name: mon.Name,
			Party: battlePartyOpponent,
			Agility: mon.Agl,
			FirstStrike: mon.FS,
			Status: mon.Status,
			Offset: mon.Offset,
		}
		participant.TickSpeed, participant.MinICV, participant.MaxICV = extractAglTierMonCustom(cfg, mon, params.IgnFirstTurn)

		response.OpponentParty = append(response.OpponentParty, participant)
		duplicates[mon.getParticipantKey()] = true
	}

	

	return response, nil
}

func fetchFormationMons(cfg *Config, formationPtr *int32, mons []turnOrderMon) []turnOrderMon {
	if formationPtr == nil {
		return mons
	}

	formation, _ := seeding.GetResourceByID(*formationPtr, cfg.l.MonsterFormationsID)

	mons = []turnOrderMon{}

	for _, monAmt := range formation.Monsters {
		monID := monAmt.MonsterID

		mon := turnOrderMon{
			ID: monID,
		}

		mons = append(mons, mon)
	}

	return mons
}

func monHasFirstStrike(mon Monster) bool {
	for _, aa := range mon.AutoAbilities {
		if aa.Name == "first strike" {
			return true
		}
	}

	return false
}


// I feel like the conditions especially in the start, can be written a bit cleaner
func fetchMonsterStatus(mon Monster, statusPtr *string) (*string, error) {
	const statusHaste = string(database.HasteStatusHaste)
	const statusAutoHaste = string(database.HasteStatusAutoHaste)
	immuneToHaste := monImmuneToHaste(mon)
	hasAppliedStatus := monHasAppliedStatus(mon)

	if statusPtr == nil && !hasAppliedStatus {
		return nil, nil
	}

	if hasAppliedStatus {
		monStatus := mon.AppliedState.AppliedStatus.StatusCondition.Name
		
		if monStatus == statusHaste {
			return &monStatus, nil
		}
	}

	status := *statusPtr

	if immuneToHaste && (status == statusHaste || status == statusAutoHaste) {
		return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("monster '%s' is immune to 'haste'", h.NameToString(mon.Name, mon.Version, nil)), nil)
	}

	return &status, nil
}

func monHasAppliedStatus(mon Monster) bool {
	return mon.AppliedState != nil && mon.AppliedState.AppliedStatus != nil
}

func monImmuneToHaste(mon Monster) bool {
	for _, condition := range mon.StatusImmunities {
		if condition.Name == string(database.HasteStatusHaste) {
			return true
		}
	}

	return false
}

func getMonsterFromJson(cfg *Config, mon turnOrderMon) (Monster, error) {
	monsterLookup, _ := seeding.GetResourceByID(mon.ID, cfg.l.MonstersID)

	monster := Monster{
		ID:                   monsterLookup.ID,
		Name:                 monsterLookup.Name,
		Version:              monsterLookup.Version,
		Specification:        monsterLookup.Specification,
		HasOverdrive:         monsterLookup.HasOverdrive,
		IsUnderwater:         monsterLookup.IsUnderwater,
		IsZombie:             monsterLookup.IsZombie,
		Distance:             monsterLookup.Distance,
		Properties:           namesToNamedAPIResources(cfg, cfg.e.properties, monsterLookup.Properties),
		AutoAbilities:        namesToNamedAPIResources(cfg, cfg.e.autoAbilities, monsterLookup.AutoAbilities),
		StealGil:             monsterLookup.StealGil,
		DoomCountdown:        monsterLookup.DoomCountdown,
		PoisonRate:           monsterLookup.PoisonRate,
		ThreatenChance:       monsterLookup.ThreatenChance,
		ZanmatoLevel:         monsterLookup.ZanmatoLevel,
		BaseStats:            toResAmtType(cfg, cfg.e.stats, monsterLookup.BaseStats, newBaseStat),
		ElemResists:          getMonsterElemResists(cfg, monsterLookup.ElemResists),
		StatusImmunities:     namesToNamedAPIResources(cfg, cfg.e.statusConditions, monsterLookup.StatusImmunities),
		StatusResists:        toResAmtType(cfg, cfg.e.statusConditions, monsterLookup.StatusResists, newStatusResist),
		Abilities:            convertObjSlice(cfg, monsterLookup.Abilities, convertMonsterAbility),
		AlteredStates: 		  getMonsterAlteredStates(cfg, nil, monsterLookup),
	}

	return applyAlteredStateFromJson(cfg, monster, mon.AltState)
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

// both these functions need to account for first strike and status
// they should probably also take the respective type as input, instead of just the agility
// ignFirstTurn also plays a role. it sets both ICVs to the offset value
func extractAglTierChar(cfg *Config, params turnOrderParty, ignFirstTurn bool) (int32, *int32, *int32) {
	agilityTier := getAgilityTier(cfg, params.Agl)
	
	tickSpeed := agilityTier.TickSpeed
	var minICV *int32
	maxICV := agilityTier.CharacterMaxICV

	for _, subtier := range agilityTier.CharacterMinICVs {
		if params.Agl >= subtier.MinAgility && params.Agl <= subtier.MaxAgility {
			minICV = subtier.CharacterMinICV
			break
		}
	}

	return calcAgilityVals(minICV, maxICV, tickSpeed, params.Offset, params.FS, ignFirstTurn, params.Status, battlePartyPlayer)
}


func extractAglTierMon(cfg *Config, params turnOrderMon, agility int32, firstStrike, ignFirstTurn bool) (int32, *int32, *int32) {
	agilityTier := getAgilityTier(cfg, agility)
	
	tickSpeed := agilityTier.TickSpeed
	minICV := agilityTier.MonsterMinICV
	maxICV := agilityTier.MonsterMaxICV

	return calcAgilityVals(minICV, maxICV, tickSpeed, params.Offset, firstStrike, ignFirstTurn, params.Status, battlePartyOpponent)
}

func extractAglTierMonCustom(cfg *Config, params turnOrderMonCustom,ignFirstTurn bool) (int32, *int32, *int32) {
	agilityTier := getAgilityTier(cfg, params.Agl)
	
	tickSpeed := agilityTier.TickSpeed
	minICV := agilityTier.MonsterMinICV
	maxICV := agilityTier.MonsterMaxICV

	return calcAgilityVals(minICV, maxICV, tickSpeed, params.Offset, params.FS, ignFirstTurn, params.Status, battlePartyOpponent)
}

func calcAgilityVals(minICV, maxICV *int32, tickSpeed, offset int32, firstStrike, ignFirstTurn bool, status *string, partyType BattleParty) (int32, *int32, *int32) {
	tickSpeed = calcTickSpeed(tickSpeed, status)
	minICV, maxICV = calcIcvVals(minICV, maxICV, firstStrike, ignFirstTurn, status, offset, partyType)

	return tickSpeed, minICV, maxICV
}


func calcTickSpeed(tickSpeed int32, statusPtr *string) int32 {
	if statusPtr == nil {
		return tickSpeed
	}
	status := *statusPtr

	switch status {
	case string(database.HasteStatusAutoHaste), string(database.HasteStatusHaste):
		return tickSpeed /2

	case string(database.HasteStatusSlow):
		return tickSpeed * 2

	default:
		return tickSpeed
	}
}

func calcIcvVals(minPtr, maxPtr *int32, firstStrike, ignFirstTurn bool, status *string, offset int32, partyType BattleParty) (*int32, *int32) {
	if minPtr == nil && maxPtr == nil {
		return nil, nil
	}

	var minICV int32
	var maxICV int32
	
	if ignFirstTurn {
		minICV = offset
		maxICV = offset
		return &minICV, &maxICV
	}

	if firstStrike {
		switch partyType {
		case battlePartyPlayer:
			minICV = 0
			maxICV = 0
			return &minICV, &maxICV

		case battlePartyOpponent:
			minICV = -1
			maxICV = -1
			return &minICV, &maxICV
		} 
	}

	minICV = *minPtr
	maxICV = *maxPtr
	
	if status == nil {
		return &minICV, &maxICV
	}

	if *status == string(database.HasteStatusAutoHaste) {
		minICV /= 2
		maxICV /= 2
	}

	return &minICV, &maxICV
}