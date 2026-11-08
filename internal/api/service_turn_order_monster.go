package api

import (
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/servicetypes"
)

func fetchFormationMons(params servicetypes.TurnOrderParams) servicetypes.TurnOrderParams {
	if params.FormationID == nil {
		return params
	}

	formation, _ := lookup.GetResourceByID(*params.FormationID, lookup.L.MonsterFormationsID)

	if formation.FormationData.IsForcedAmbush {
		params.BattleStart = string(database.BattleStartAmbush)
	}

	for _, monAmt := range formation.Monsters {
		monID := monAmt.MonsterID

		mon := servicetypes.TurnOrderMon{
			MonsterID: &monID,
		}

		params.Mons = append(params.Mons, mon)
	}

	return params
}

func assembleMonTurnOrder(monID int32, altState *int32) (apitypes.Monster, *seedtypes.AgilityTier) {
	penanceArmIDs := []int32{306, 307}
	firstTurnAglIDs := []int32{167, 168, 215, 306, 307}
	var firstTurnAglTier *seedtypes.AgilityTier

	if slices.Contains(penanceArmIDs, monID) {
		altState = h.GetInt32Ptr(2)
	}

	monster, _ := quickAssembleMon(monID, altState)

	if slices.Contains(firstTurnAglIDs, monID) {
		firstTurnAglTier = getFirstTurnAglTier(monID)
	}

	return monster, firstTurnAglTier
}

func getFirstTurnAglTier(monID int32) *seedtypes.AgilityTier {
	mon, _ := quickAssembleMon(monID, nil)
	altState := mon.AlteredStates[0]

	if len(altState.Alts) == 0 {
		return nil
	}

	baseStats := altState.Alts[0].BaseStats

	aglBS := getBaseStat("agility", baseStats)
	tier := getTierByAgility(aglBS.Value)

	return &tier
}

func getTurnOrderMonAgility(monID int32, monster apitypes.Monster) int32 {
	agility := getBaseStatVal("agility", monster.BaseStats)

	return handleNoTurnMons(monID, agility)
}

func handleNoTurnMons(monID, agility int32) int32 {
	noTurnIDs := []int32{33, 166, 194, 212, 228, 300}

	if slices.Contains(noTurnIDs, monID) {
		return 0
	}

	return agility
}

func fetchMonsterHasteStatus(mon apitypes.Monster, statusPtr *string, ignImmunities bool) (*string, error) {
	const statusSlow = string(database.HasteStatusSlow)
	const statusHaste = string(database.HasteStatusHaste)
	const statusAutoHaste = string(database.HasteStatusAutoHaste)

	if monHasAppliedStatus(mon) {
		monStatus := mon.AppliedState.AppliedStatus.StatusCondition.Name

		if statusPtr == nil && (monStatus == statusHaste || monStatus == statusSlow) {
			statusPtr = &monStatus
		}
	}

	if statusPtr == nil {
		return nil, nil
	}

	status := *statusPtr

	if status == statusAutoHaste {
		status = statusHaste
	}

	err := enforceMonImmunity(status, mon, ignImmunities)
	if err != nil {
		return nil, err
	}

	return &status, nil
}
