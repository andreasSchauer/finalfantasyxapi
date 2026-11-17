package turnorder

import (
	th "github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/typehelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/service/servicetypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func getAglValsChar(participant servicetypes.Participant, params *servicetypes.TurnOrderParams) servicetypes.AgilityVals {
	agilityTier := getTierByAgility(participant.Agility)
	aglVals := servicetypes.AgilityVals{
		TickSpeed: agilityTier.TickSpeed,
		MaxICV:    h.CopyPtr(agilityTier.CharacterMaxICV),
	}

	for _, subtier := range agilityTier.CharacterMinICVs {
		if participant.Agility >= subtier.MinAgility && participant.Agility <= subtier.MaxAgility {
			aglVals.MinICV = h.CopyPtr(subtier.CharacterMinICV)
			break
		}
	}

	calcAgilityVals(&aglVals, participant, params, nil)
	return aglVals
}

func getAglValsMon(participant servicetypes.Participant, params *servicetypes.TurnOrderParams, monFirstTurnTier *seedtypes.AgilityTier) servicetypes.AgilityVals {
	agilityTier := getTierByAgility(participant.Agility)
	aglVals := servicetypes.AgilityVals{
		TickSpeed: agilityTier.TickSpeed,
		MinICV:    h.CopyPtr(agilityTier.MonsterMinICV),
		MaxICV:    h.CopyPtr(agilityTier.MonsterMaxICV),
	}

	calcAgilityVals(&aglVals, participant, params, monFirstTurnTier)
	return aglVals
}

func calcAgilityVals(aglVals *servicetypes.AgilityVals, participant servicetypes.Participant, params *servicetypes.TurnOrderParams, monFirstTurnTier *seedtypes.AgilityTier) {
	aglVals.TickSpeed = calcTickSpeed(aglVals.TickSpeed, participant.Status)
	calcAglValICVs(aglVals, participant, params, monFirstTurnTier)
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

func calcAglValICVs(aglVals *servicetypes.AgilityVals, participant servicetypes.Participant, params *servicetypes.TurnOrderParams, monFirstTurnTier *seedtypes.AgilityTier) {
	if monFirstTurnTier != nil {
		aglVals.MinICV = h.CopyPtr(monFirstTurnTier.MonsterMinICV)
		aglVals.MaxICV = h.CopyPtr(monFirstTurnTier.MonsterMaxICV)
	}

	if aglVals.MinICV == nil || aglVals.MaxICV == nil {
		return
	}

	if params.IgnFirstTurn {
		aglValsEqualICVs(aglVals, 0)
		return
	}

	if participant.FirstStrike {
		switch participant.Party {
		case servicetypes.BattlePartyPlayer:
			aglValsEqualICVs(aglVals, 0)
			return

		case servicetypes.BattlePartyOpponent:
			aglValsEqualICVs(aglVals, -1)
			return
		}
	}

	switch params.BattleStart {
	case string(database.BattleStartAmbush):
		switch participant.Party {
		case servicetypes.BattlePartyPlayer:
			calcICVsStartFavorable(aglVals, participant.Status)
			return

		case servicetypes.BattlePartyOpponent:
			aglValsEqualICVs(aglVals, 0)
			return
		}

	case string(database.BattleStartPreemptive):
		switch participant.Party {
		case servicetypes.BattlePartyPlayer:
			aglValsEqualICVs(aglVals, 0)
			return

		case servicetypes.BattlePartyOpponent:
			calcICVsStartFavorable(aglVals, participant.Status)
			return
		}
	}

	if participant.Status != nil && *participant.Status == string(database.HasteStatusAutoHaste) {
		*aglVals.MinICV /= 2
		*aglVals.MaxICV /= 2
	}
}

func aglValsEqualICVs(aglVals *servicetypes.AgilityVals, icv int32) {
	aglVals.MinICV, aglVals.MaxICV = th.GetEqualICVs(icv)
}

func calcICVsStartFavorable(aglVals *servicetypes.AgilityVals, status *string) {
	aglValsEqualICVs(aglVals, *aglVals.MinICV*3)

	if status == nil || *status != string(database.HasteStatusAutoHaste) {
		return
	}

	*aglVals.MinICV /= 2
	*aglVals.MaxICV /= 2
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

func getAllAgilityTiers() []seedtypes.AgilityTier {
	tiers := make([]seedtypes.AgilityTier, len(lookup.L.AgilityTiersID))

	for _, tier := range lookup.L.AgilityTiersID {
		tiers = append(tiers, tier)
	}

	return tiers
}
