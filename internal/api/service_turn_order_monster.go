package api

import (
	"fmt"
	"net/http"
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func fetchFormationMons(cfg *Config, params TurnOrderParams) TurnOrderParams {
	if params.Formation == nil {
		return params
	}

	formation, _ := seeding.GetResourceByID(*params.Formation, cfg.l.MonsterFormationsID)

	if formation.FormationData.IsForcedAmbush {
		params.BattleStart = string(database.BattleStartAmbush)
	}

	for _, monAmt := range formation.Monsters {
		monID := monAmt.MonsterID

		mon := turnOrderMon{
			ID: &monID,
		}

		params.Mons = append(params.Mons, mon)
	}

	return params
}

func assembleMonTurnOrder(cfg *Config, monID int32, altState *int32) (Monster, *seeding.AgilityTier, error) {
	penanceArmIDs := []int32{306, 307}
	firstTurnAglIDs := []int32{167, 168, 215, 306, 307}
	var firstTurnAglTier *seeding.AgilityTier

	if slices.Contains(penanceArmIDs, monID) {
		altState = h.GetInt32Ptr(2)
	}

	monster, err := quickAssembleMon(cfg, monID, altState)
	if err != nil {
		return Monster{}, nil, err
	}

	if slices.Contains(firstTurnAglIDs, monID) {
		firstTurnAglTier = getAltStateAglTier(cfg, monster)
	}

	return monster, firstTurnAglTier, nil
}

func getAltStateAglTier(cfg *Config, mon Monster) *seeding.AgilityTier {
	altState := mon.AlteredStates[0]

	if len(altState.Alts) == 0 {
		return nil
	}

	baseStats := altState.Alts[0].BaseStats

	aglBS := getBaseStat(cfg, "agility", baseStats)
	tier := getAgilityTier(cfg, aglBS.Value)

	return &tier
}

func getTurnOrderMonAgility(cfg *Config, monID int32, monster Monster) int32 {
	agilityBS := getBaseStat(cfg, "agility", monster.BaseStats)
	agility := agilityBS.Value

	return handleNoTurnMons(monID, agility)
}

func handleNoTurnMons(monID, agility int32) int32 {
	noTurnIDs := []int32{33, 166, 194, 212, 228, 300}

	if slices.Contains(noTurnIDs, monID) {
		return 0
	}

	return agility
}

func fetchMonsterHasteStatus(params TurnOrderParams, mon Monster, statusPtr *string) (*string, error) {
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

	if !params.IgnImmunities {
		if monImmuneToSlow(mon) && status == statusSlow {
			return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("monster '%s' is immune to 'slow'", h.NameToString(mon.Name, mon.Version, nil)), nil)
		}

		if monImmuneToHaste(mon) && (status == statusHaste || status == statusAutoHaste) {
			return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("monster '%s' is immune to 'haste'", h.NameToString(mon.Name, mon.Version, nil)), nil)
		}
	}

	return &status, nil
}
