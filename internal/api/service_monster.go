package api

import (
	"fmt"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
)

func quickAssembleMon(server *Server, monID int32, altStatePtr *int32) (Monster, error) {
	monsterLookup, _ := lookup.GetResourceByID(monID, lookup.L.MonstersID)

	monster := Monster{
		ID:               monsterLookup.ID,
		Name:             monsterLookup.Name,
		Version:          monsterLookup.Version,
		Specification:    monsterLookup.Specification,
		HasOverdrive:     monsterLookup.HasOverdrive,
		IsUnderwater:     monsterLookup.IsUnderwater,
		IsZombie:         monsterLookup.IsZombie,
		Distance:         monsterLookup.Distance,
		Properties:       namesToNamedAPIResources(server, E.Properties, monsterLookup.Properties),
		AutoAbilities:    namesToNamedAPIResources(server, E.AutoAbilities, monsterLookup.AutoAbilities),
		StealGil:         monsterLookup.StealGil,
		DoomCountdown:    monsterLookup.DoomCountdown,
		PoisonRate:       monsterLookup.PoisonRate,
		ThreatenChance:   monsterLookup.ThreatenChance,
		ZanmatoLevel:     monsterLookup.ZanmatoLevel,
		BaseStats:        E.Stats.toResAmtType(server, monsterLookup.BaseStats, newBaseStat),
		ElemResists:      getMonsterElemResists(server, monsterLookup.ElemResists),
		StatusImmunities: namesToNamedAPIResources(server, E.StatusConditions, monsterLookup.StatusImmunities),
		StatusResists:    E.StatusConditions.toResAmtType(server, monsterLookup.StatusResists, newStatusResist),
		Abilities:        convertObjSlice(server, monsterLookup.Abilities, convertMonsterAbility),
		AlteredStates:    getMonsterAlteredStates(server, nil, monsterLookup),
	}

	monster, err := applyAlteredStateFromJSON(server, monster, altStatePtr)
	if err != nil {
		return Monster{}, err
	}

	return monster, nil
}

func monsterHasFirstStrike(mon Monster) bool {
	for _, aa := range mon.AutoAbilities {
		if aa.Name == "first strike" {
			return true
		}
	}

	return false
}

func monHasAppliedStatus(mon Monster) bool {
	return mon.AppliedState != nil && mon.AppliedState.AppliedStatus != nil
}

func enforceMonImmunity(status string, mon Monster, ignImmunities bool) error {
	if !ignImmunities && monImmuneToStatus(status, mon) {
		return newHTTPError(http.StatusBadRequest, fmt.Sprintf("monster '%s' is immune to '%s'", h.NameToString(mon.Name, mon.Version, nil), status), nil)
	}

	return nil
}

func monImmuneToStatus(status string, mon Monster) bool {
	for _, condition := range mon.StatusImmunities {
		if condition.Name == status {
			return true
		}
	}

	return false
}
