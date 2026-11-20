package get

import (
	"fmt"
	"net/http"
	"slices"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
)

func applyAlteredStateFromQuery(r *http.Request, mon *apitypes.Monster, queryName ts.QueryParamName) error {
	altStateID, err := getAltStateID(r, mon, queryName)
	if errs.QueryIsEmpty(err) {
		return nil
	}
	if err != nil {
		return err
	}

	applyAlteredState(mon, altStateID)
	return nil
}

func ApplyAlteredStateFromJSON(mon *apitypes.Monster, altStatePtr *int32) error {
	if altStatePtr == nil {
		return nil
	}

	altStateID := int(*altStatePtr)

	err := vpAltStateID(mon, altStateID)
	if err != nil {
		return err
	}

	applyAlteredState(mon, altStateID)
	return nil
}

func applyAlteredState(mon *apitypes.Monster, altStateID int) {
	altState := mon.AlteredStates[altStateID-1]
	appliedState := &apitypes.AppliedState{
		Condition:   altState.Condition,
		IsTemporary: altState.IsTemporary,
	}

	defaultState := &apitypes.AlteredState{
		URL:         ah.CreateResourceURL(apilookups.E.Monsters.Endpoint, mon.ID),
		Condition:   "default",
		IsTemporary: false,
	}

	for _, change := range altState.Alts {
		switch database.AlterationType(change.AlterationType) {
		case database.AlterationTypeChange:
			applyAltChange(mon, change, defaultState)

		case database.AlterationTypeGain:
			applyAltGain(mon, change, appliedState, defaultState)

		case database.AlterationTypeLoss:
			applyAltLoss(mon, change, defaultState)
		}
	}

	mon.AppliedState = appliedState
	if appliedState.IsZero() {
		mon.AppliedState = nil
	}

	replaceAltState(mon, *defaultState, altStateID)
}

func getAltStateID(r *http.Request, mon *apitypes.Monster, queryName ts.QueryParamName) (int, error) {
	queryParam := apilookups.Q.Monsters[queryName]
	query, err := queryParam.CheckEmpty(r)
	if err != nil {
		return 0, err
	}

	if len(mon.AlteredStates) == 0 {
		return 0, errs.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%s has no altered states", mon.String()), nil)
	}

	id, err := queryParam.ParseIDVal(query, len(mon.AlteredStates))
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func vpAltStateID(mon *apitypes.Monster, altStateID int) error {
	altStateAmt := len(mon.AlteredStates)

	if altStateAmt == 0 {
		return errs.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%s has no altered states", mon.String()), nil)
	}

	if altStateID > altStateAmt || altStateID <= 0 {
		return errs.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("provided alt state id '%d' used for %s is out of range. max id: %d", altStateID, mon.String(), altStateAmt), nil)
	}

	return nil
}

func applyAltChange(mon *apitypes.Monster, change apitypes.Alt, defaultState *apitypes.AlteredState) {
	defStateChange := apitypes.Alt{
		AlterationType: database.AlterationTypeChange,
	}

	if change.Distance != nil {
		newDistance := mon.Distance
		defStateChange.Distance = &newDistance
		mon.Distance = *change.Distance
	}

	mon.BaseStats, defStateChange.BaseStats = modifyResourcesChange(mon.BaseStats, change.BaseStats)
	mon.ElemResists, defStateChange.ElemResists = modifyResourcesChange(mon.ElemResists, change.ElemResists)

	defaultState.Alts = append(defaultState.Alts, defStateChange)
}

func applyAltGain(mon *apitypes.Monster, change apitypes.Alt, appliedState *apitypes.AppliedState, defaultState *apitypes.AlteredState) {
	defStateLoss := apitypes.Alt{
		AlterationType: database.AlterationTypeLoss,
	}
	defStateGain := apitypes.Alt{
		AlterationType: database.AlterationTypeGain,
	}

	mon.Properties, defStateLoss.Properties = modifyResourcesGain(mon.Properties, change.Properties)
	mon.AutoAbilities, defStateLoss.AutoAbilities = modifyResourcesGain(mon.AutoAbilities, change.AutoAbilities)

	modifyGainedImmunities(mon, &change, &defStateGain, &defStateLoss, appliedState)

	if change.AddedStatus != nil {
		appliedState.AppliedStatus = change.AddedStatus
		defStateLoss.RemovedStatus = &change.AddedStatus.StatusCondition
	}

	if !defStateLoss.IsZero() {
		defaultState.Alts = append(defaultState.Alts, defStateLoss)
	}

	if !defStateGain.IsZero() {
		defaultState.Alts = append(defaultState.Alts, defStateGain)
	}
}

func applyAltLoss(mon *apitypes.Monster, change apitypes.Alt, defaultState *apitypes.AlteredState) {
	defStateGain := apitypes.Alt{
		AlterationType: database.AlterationTypeGain,
	}

	mon.Properties, defStateGain.Properties = modifyResourcesLoss(mon.Properties, change.Properties)
	mon.AutoAbilities, defStateGain.AutoAbilities = modifyResourcesLoss(mon.AutoAbilities, change.AutoAbilities)

	if !defStateGain.IsZero() {
		defaultState.Alts = append(defaultState.Alts, defStateGain)
	}
}

// put default in first and cut out the currently applied state (at index i, determined by altStateID)
func replaceAltState(mon *apitypes.Monster, def apitypes.AlteredState, i int) {
	allStates := h.Unshift(mon.AlteredStates, def)
	size := len(allStates)
	s1 := allStates[0:i]
	s2 := allStates[i+1:size]
	mon.AlteredStates = slices.Concat(s1, s2)
}
