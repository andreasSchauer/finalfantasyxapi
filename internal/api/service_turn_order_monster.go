package api

import (
	"fmt"
	"net/http"
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func handleMonEdgeCases(mon turnOrderMon, agility int32) int32 {
	noTurnIDs := []int32{33, 166, 194, 212, 228, 300}

	if slices.Contains(noTurnIDs, mon.ID) {
		return 0
	}

	return agility
}

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
			ID: monID,
		}

		params.Mons = append(params.Mons, mon)
	}

	return params
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
