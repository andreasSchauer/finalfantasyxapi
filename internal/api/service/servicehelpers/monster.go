package servicehelpers

import (
	"fmt"
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apiconvert"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/get"
	th "github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/typehelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
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
		ElemResists:      get.MonsterElemResists(monsterLookup.ElemResists),
		StatusImmunities: handlerinput.NamesToNamedAPIResources(apilookups.E.StatusConditions, monsterLookup.StatusImmunities),
		StatusResists:    apilookups.E.StatusConditions.ToResAmtType(monsterLookup.StatusResists, th.NewStatusResist),
		Abilities:        ah.ConvertObjSlice(monsterLookup.Abilities, apiconvert.MonsterAbility),
		AlteredStates:    get.MonsterAlteredStates(nil, monsterLookup),
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
