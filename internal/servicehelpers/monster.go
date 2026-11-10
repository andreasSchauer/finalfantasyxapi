package servicehelpers

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiconv"
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/get"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	th "github.com/andreasSchauer/finalfantasyxapi/internal/typehelpers"
)

func QuickAssembleMon(monID int32, altStatePtr *int32) (apitypes.Monster, error) {
	monsterLookup, _ := lookup.GetResourceByID(monID, lookup.L.MonstersID)

	monster := apitypes.Monster{
		ID:               monsterLookup.ID,
		Name:             monsterLookup.Name,
		Version:          monsterLookup.Version,
		Specification:    monsterLookup.Specification,
		HasOverdrive:     monsterLookup.HasOverdrive,
		IsUnderwater:     monsterLookup.IsUnderwater,
		IsZombie:         monsterLookup.IsZombie,
		Distance:         monsterLookup.Distance,
		Properties:       handlerinput.NamesToNamedAPIResources(apilookups.E.Properties, monsterLookup.Properties),
		AutoAbilities:    handlerinput.NamesToNamedAPIResources(apilookups.E.AutoAbilities, monsterLookup.AutoAbilities),
		StealGil:         monsterLookup.StealGil,
		DoomCountdown:    monsterLookup.DoomCountdown,
		PoisonRate:       monsterLookup.PoisonRate,
		ThreatenChance:   monsterLookup.ThreatenChance,
		ZanmatoLevel:     monsterLookup.ZanmatoLevel,
		BaseStats:        apilookups.E.Stats.ToResAmtType(monsterLookup.BaseStats, th.NewBaseStat),
		ElemResists:      get.GetMonsterElemResists(monsterLookup.ElemResists),
		StatusImmunities: handlerinput.NamesToNamedAPIResources(apilookups.E.StatusConditions, monsterLookup.StatusImmunities),
		StatusResists:    apilookups.E.StatusConditions.ToResAmtType(monsterLookup.StatusResists, th.NewStatusResist),
		Abilities:        ah.ConvertObjSlice(monsterLookup.Abilities, apiconv.ConvertMonsterAbility),
		AlteredStates:    get.GetMonsterAlteredStates(nil, monsterLookup),
	}

	monster, err := get.ApplyAlteredStateFromJSON(monster, altStatePtr)
	if err != nil {
		return apitypes.Monster{}, err
	}

	return monster, nil
}

func MonsterHasFirstStrike(mon apitypes.Monster) bool {
	for _, aa := range mon.AutoAbilities {
		if aa.Name == "first strike" {
			return true
		}
	}

	return false
}

func monHasAppliedStatus(mon apitypes.Monster) bool {
	return mon.AppliedState != nil && mon.AppliedState.AppliedStatus != nil
}

func enforceMonImmunity(status string, mon apitypes.Monster, ignImmunities bool) error {
	if !ignImmunities && monImmuneToStatus(status, mon) {
		return errs.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("monster '%s' is immune to '%s'", h.NameToString(mon.Name, mon.Version, nil), status), nil)
	}

	return nil
}

func monImmuneToStatus(status string, mon apitypes.Monster) bool {
	for _, condition := range mon.StatusImmunities {
		if condition.Name == status {
			return true
		}
	}

	return false
}

func FetchMonsterHasteStatus(mon apitypes.Monster, statusPtr *string, ignImmunities bool) (*string, error) {
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
