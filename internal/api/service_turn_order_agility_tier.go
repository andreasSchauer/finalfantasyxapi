package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/servicetypes"
)

func getAllAgilityTiers() []seedtypes.AgilityTier {
	tiers := make([]seedtypes.AgilityTier, len(lookup.L.AgilityTiersID))

	for _, tier := range lookup.L.AgilityTiersID {
		tiers = append(tiers, tier)
	}

	return tiers
}

func getTierByAgility(agility int32) seedtypes.AgilityTier {
	tiers := getAllAgilityTiers()
	var agilityTier seedtypes.AgilityTier

	for _, tier := range tiers {
		if agility >= tier.MinAgility && agility <= tier.MaxAgility {
			agilityTier = tier
			break
		}
	}

	return agilityTier
}

func extractAglTierChar(participant servicetypes.Participant, params servicetypes.TurnOrderParams) servicetypes.AgilityVals {
	agilityTier := getTierByAgility(participant.Agility)
	aglVals := servicetypes.AgilityVals{
		TickSpeed: agilityTier.TickSpeed,
		MaxICV:    agilityTier.CharacterMaxICV,
	}

	for _, subtier := range agilityTier.CharacterMinICVs {
		if participant.Agility >= subtier.MinAgility && participant.Agility <= subtier.MaxAgility {
			aglVals.MinICV = subtier.CharacterMinICV
			break
		}
	}

	return calcAgilityVals(aglVals, participant, params, nil)
}

func extractAglTierMon(participant servicetypes.Participant, params servicetypes.TurnOrderParams, monFirstTurnTier *seedtypes.AgilityTier) servicetypes.AgilityVals {
	agilityTier := getTierByAgility(participant.Agility)
	aglVals := servicetypes.AgilityVals{
		TickSpeed: agilityTier.TickSpeed,
		MinICV:    agilityTier.MonsterMinICV,
		MaxICV:    agilityTier.MonsterMaxICV,
	}

	return calcAgilityVals(aglVals, participant, params, monFirstTurnTier)
}

func calcAgilityVals(aglVals servicetypes.AgilityVals, participant servicetypes.Participant, params servicetypes.TurnOrderParams, monFirstTurnTier *seedtypes.AgilityTier) servicetypes.AgilityVals {
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
		return tickSpeed / 2

	case string(database.HasteStatusSlow):
		return tickSpeed * 2

	default:
		return tickSpeed
	}
}

func calcICVs(minPtr, maxPtr *int32, participant servicetypes.Participant, params servicetypes.TurnOrderParams, monFirstTurnTier *seedtypes.AgilityTier) (*int32, *int32) {
	if monFirstTurnTier != nil {
		minPtr = monFirstTurnTier.MonsterMinICV
		maxPtr = monFirstTurnTier.MonsterMaxICV
	}

	if minPtr == nil || maxPtr == nil {
		return nil, nil
	}

	if params.IgnFirstTurn {
		return getEqualICVs(0)
	}

	if participant.FirstStrike {
		switch participant.Party {
		case servicetypes.BattlePartyPlayer:
			return getEqualICVs(0)

		case servicetypes.BattlePartyOpponent:
			return getEqualICVs(-1)
		}
	}

	switch params.BattleStart {
	case string(database.BattleStartAmbush):
		switch participant.Party {
		case servicetypes.BattlePartyPlayer:
			return calcICVsStartFavorable(*minPtr, participant.Status)

		case servicetypes.BattlePartyOpponent:
			return getEqualICVs(0)
		}

	case string(database.BattleStartPreemptive):
		switch participant.Party {
		case servicetypes.BattlePartyPlayer:
			return getEqualICVs(0)

		case servicetypes.BattlePartyOpponent:
			return calcICVsStartFavorable(*minPtr, participant.Status)
		}
	}

	minICV := *minPtr
	maxICV := *maxPtr

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
