package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type AgilityVals struct {
	TickSpeed		int32		`json:"tick_speed"`
	MinICV			*int32		`json:"min_icv"`
	MaxICV			*int32		`json:"max_icv"`
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

func extractAglTierChar(cfg *Config, participant Participant, params TurnOrderParams) AgilityVals {
	agilityTier := getAgilityTier(cfg, participant.Agility)
	aglVals := AgilityVals{
		TickSpeed: 	agilityTier.TickSpeed,
		MaxICV: 	agilityTier.CharacterMaxICV,
	}

	for _, subtier := range agilityTier.CharacterMinICVs {
		if participant.Agility >= subtier.MinAgility && participant.Agility <= subtier.MaxAgility {
			aglVals.MinICV = subtier.CharacterMinICV
			break
		}
	}

	return calcAgilityVals(aglVals, participant, params, nil)
}


func extractAglTierMon(cfg *Config, participant Participant, params TurnOrderParams, monFirstTurnTier *seeding.AgilityTier) AgilityVals {
	agilityTier := getAgilityTier(cfg, participant.Agility)
	aglVals := AgilityVals{
		TickSpeed: 	agilityTier.TickSpeed,
		MinICV: 	agilityTier.MonsterMinICV,
		MaxICV: 	agilityTier.MonsterMaxICV,
	}

	return calcAgilityVals(aglVals, participant, params, monFirstTurnTier)
}


func calcAgilityVals(aglVals AgilityVals, participant Participant, params TurnOrderParams, monFirstTurnTier *seeding.AgilityTier) AgilityVals {
	aglVals.TickSpeed = calcTickSpeed(aglVals.TickSpeed, participant.Status)
	aglVals.MinICV, aglVals.MaxICV = calcICVs(aglVals.MinICV, aglVals.MaxICV, participant, params, monFirstTurnTier)

	return aglVals
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

func calcICVs(minPtr, maxPtr *int32, participant Participant, params TurnOrderParams, monFirstTurnTier *seeding.AgilityTier) (*int32, *int32) {
	if monFirstTurnTier != nil {
		minPtr = monFirstTurnTier.MonsterMinICV
		maxPtr = monFirstTurnTier.MonsterMaxICV
	}

	if minPtr == nil || maxPtr == nil {
		return nil, nil
	}

	var minICV int32
	var maxICV int32
	
	if params.IgnFirstTurn {
		return getEqualICVs(0)
	}

	if participant.FirstStrike {
		switch participant.Party {
		case battlePartyPlayer:
			return getEqualICVs(0)

		case battlePartyOpponent:
			return getEqualICVs(-1)
		} 
	}

	switch params.BattleStart {
	case string(database.BattleStartAmbush):
		switch participant.Party {
		case battlePartyPlayer:
			return calcICVsStartFavorable(*minPtr, participant.Status)

		case battlePartyOpponent:
			return getEqualICVs(0)
		}

	case string(database.BattleStartPreemptive):
		switch participant.Party {
		case battlePartyPlayer:
			return getEqualICVs(0)
			
		case battlePartyOpponent:
			return calcICVsStartFavorable(*minPtr, participant.Status)
		}
	}

	minICV = *minPtr
	maxICV = *maxPtr

	
	if participant.Status == nil {
		return &minICV, &maxICV
	}

	if *participant.Status == string(database.HasteStatusAutoHaste) {
		minICV /= 2
		maxICV /= 2
	}

	return &minICV, &maxICV
}

func getEqualICVs(val int32) (*int32, *int32) {
	minICV := val
	maxICV := val

	return &minICV, &maxICV
}

func calcICVsStartFavorable(minVal int32, status *string) (*int32, *int32) {
	minPtr, maxPtr := getEqualICVs(minVal * 3)

	if status == nil || *status != string(database.HasteStatusAutoHaste) {
		return minPtr, maxPtr
	}
	
	*minPtr /= 2
	*maxPtr /= 2

	return minPtr, maxPtr
}