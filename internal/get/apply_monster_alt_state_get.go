package get

import (
	"net/http"
	"strconv"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiconv"
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getAlteredStatesAlterations(states []apitypes.AlteredState) []apitypes.AlteredState {
	for i := range states {
		states[i].Alterations = getAlterations(states[i].Alts)
	}

	return states
}

func GetMonsterAlteredStates(r *http.Request, mon seedtypes.Monster) []apitypes.AlteredState {
	alteredStates := []apitypes.AlteredState{}

	for i, altState := range mon.AlteredStates {
		var url string

		// if this function is triggered via a jsonParam
		if r != nil {
			q := r.URL.Query()
			q.Set(string(ts.QpnAlteredState), strconv.Itoa(i+1))
			url = ah.CreateResourceURLQuery(apilookups.E.Monsters.Endpoint, mon.ID, q)
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
		altStateChange := apiconv.ConvertAlt(change)
		altStateChange.AlterationType = database.AlterationType(change.AlterationType)
		altStateChange.Distance = change.Distance
		alts = append(alts, altStateChange)
	}

	return alts
}

func getAlterations(alts []apitypes.Alt) apitypes.Alterations {
	alterations := apitypes.Alterations{}

	for _, alt := range alts {
		switch alt.AlterationType {
		case database.AlterationTypeChange:
			change := apitypes.AltStateChange{
				Distance:    alt.Distance,
				BaseStats:   alt.BaseStats,
				ElemResists: alt.ElemResists,
			}
			alterations.Change = &change

		case database.AlterationTypeGain:
			gain := apitypes.AltStateGain{
				Properties:       alt.Properties,
				AutoAbilities:    alt.AutoAbilities,
				StatusImmunities: alt.StatusImmunities,
				StatusResists:    alt.StatusResists,
				Status:           alt.AddedStatus,
			}
			alterations.Gain = &gain

		case database.AlterationTypeLoss:
			loss := apitypes.AltStateLoss{
				Properties:       alt.Properties,
				AutoAbilities:    alt.AutoAbilities,
				StatusImmunities: alt.StatusImmunities,
				Status:           alt.RemovedStatus,
			}
			alterations.Loss = &loss
		}
	}

	return alterations
}
