package api

import (
	"net/http"
	"strconv"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getAlteredStatesAlterations(states []apitypes.AlteredState) []apitypes.AlteredState {
	for i := range states {
		states[i].Alterations = getAlterations(states[i].Alts)
	}

	return states
}

func getMonsterAlteredStates(r *http.Request, mon seedtypes.Monster) []apitypes.AlteredState {
	alteredStates := []apitypes.AlteredState{}

	for i, altState := range mon.AlteredStates {
		var url string

		// if this function is triggered via a jsonParam
		if r != nil {
			q := r.URL.Query()
			q.Set(string(ts.QpnAlteredState), strconv.Itoa(i+1))
			url = ah.CreateResourceURLQuery(E.Monsters.Endpoint, mon.ID, q)
		}

		alteredState := apitypes.AlteredState{
			URL:         url,
			Condition:   altState.Condition,
			IsTemporary: altState.IsTemporary,
			Alts:        getAlts(altState),
		}

		alteredStates = append(alteredStates, alteredState)
	}

	return alteredStates
}

func getAlts(as seedtypes.AlteredState) []apitypes.Alt {
	alts := []apitypes.Alt{}

	for _, change := range as.Changes {
		altStateChange := convertAlt(change)
		altStateChange.AlterationType = database.AlterationType(change.AlterationType)
		altStateChange.Distance = change.Distance
		alts = append(alts, altStateChange)
	}

	return alts
}

func convertAlt(asc seedtypes.Alt) apitypes.Alt {
	return apitypes.Alt{
		Properties:       h.SliceOrNil(namesToNamedAPIResources(E.Properties, asc.Properties)),
		AutoAbilities:    h.SliceOrNil(namesToNamedAPIResources(E.AutoAbilities, asc.AutoAbilities)),
		BaseStats:        h.SliceOrNil(E.Stats.ToResAmtType(asc.BaseStats, newBaseStat)),
		ElemResists:      h.SliceOrNil(namesToElemResists(asc.ElemResists)),
		StatusImmunities: h.SliceOrNil(namesToNamedAPIResources(E.StatusConditions, asc.StatusImmunities)),
		AddedStatus:      ah.ConvertObjPtr(asc.AddedStatus, convertInflictedStatus),
	}
}
