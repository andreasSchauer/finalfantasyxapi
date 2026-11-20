package servicehelpers

import (
	"fmt"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apiconvert"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/get"
	th "github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/typehelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
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

	err := get.ApplyAlteredStateFromJSON(&monster, altStatePtr)
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

func MonHasAppliedStatus(mon apitypes.Monster) bool {
	return mon.AppliedState != nil && mon.AppliedState.AppliedStatus != nil
}


func MonImmuneToStatus(status string, mon apitypes.Monster) bool {
	for _, condition := range mon.StatusImmunities {
		if condition.Name == status {
			return true
		}
	}

	return false
}


func EnforceMonImmunity(status string, mon apitypes.Monster, ignImmunities bool) error {
	if !ignImmunities && MonImmuneToStatus(status, mon) {
		return errs.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("monster '%s' is immune to '%s'", h.NameToString(mon.Name, mon.Version, nil), status), nil)
	}

	return nil
}
