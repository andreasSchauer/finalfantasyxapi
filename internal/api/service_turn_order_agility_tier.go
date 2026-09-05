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

	return calcAgilityVals(aglVals, participant, params)
}


func extractAglTierMon(cfg *Config, participant Participant, params TurnOrderParams) AgilityVals {
	agilityTier := getAgilityTier(cfg, participant.Agility)
	aglVals := AgilityVals{
		TickSpeed: 	agilityTier.TickSpeed,
		MinICV: 	agilityTier.MonsterMinICV,
		MaxICV: 	agilityTier.MonsterMaxICV,
	}

	return calcAgilityVals(aglVals, participant, params)
}


func calcAgilityVals(aglVals AgilityVals, participant Participant, params TurnOrderParams) AgilityVals {
	aglVals.TickSpeed = calcTickSpeed(aglVals.TickSpeed, participant.Status)
	aglVals.MinICV, aglVals.MaxICV = calcICVs(aglVals.MinICV, aglVals.MaxICV, participant, params)

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

func calcICVs(minPtr, maxPtr *int32, participant Participant, params TurnOrderParams) (*int32, *int32) {
	if minPtr == nil || maxPtr == nil {
		return nil, nil
	}

	var minICV int32
	var maxICV int32
	
	if params.IgnFirstTurn {
		minICV = participant.Offset
		maxICV = participant.Offset
		return &minICV, &maxICV
	}

	if participant.FirstStrike {
		switch participant.Party {
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

	switch params.BattleStart {
	case string(database.BattleStartAmbush):
		switch participant.Party {
		case battlePartyPlayer:
			return calcICVsStartFavorable(*minPtr, participant.Status)

		case battlePartyOpponent:
			minICV = 0
			maxICV = 0
			return &minICV, &maxICV
		}

	case string(database.BattleStartPreemptive):
		switch participant.Party {
		case battlePartyPlayer:
			minICV = 0
			maxICV = 0
			return &minICV, &maxICV
			
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

func calcICVsStartFavorable(minVal int32, status *string) (*int32, *int32) {
	val := minVal * 3
	minICV := val
	maxICV := val

	if status == nil || *status != string(database.HasteStatusAutoHaste) {
		return &minICV, &maxICV
	}
	
	minICV /= 2
	maxICV /= 2

	return &minICV, &maxICV
}