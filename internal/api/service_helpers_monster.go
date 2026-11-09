package api

import (
	"fmt"
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	th "github.com/andreasSchauer/finalfantasyxapi/internal/typehelpers"
)

func quickAssembleMon(monID int32, altStatePtr *int32) (apitypes.Monster, error) {
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
		ElemResists:      getMonsterElemResists(monsterLookup.ElemResists),
		StatusImmunities: handlerinput.NamesToNamedAPIResources(apilookups.E.StatusConditions, monsterLookup.StatusImmunities),
		StatusResists:    apilookups.E.StatusConditions.ToResAmtType(monsterLookup.StatusResists, th.NewStatusResist),
		Abilities:        ah.ConvertObjSlice(monsterLookup.Abilities, ConvertMonsterAbility),
		AlteredStates:    getMonsterAlteredStates(nil, monsterLookup),
	}

	monster, err := applyAlteredStateFromJSON(monster, altStatePtr)
	if err != nil {
		return apitypes.Monster{}, err
	}

	return monster, nil
}

func monsterHasFirstStrike(mon apitypes.Monster) bool {
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
