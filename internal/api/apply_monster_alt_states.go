package api

import (
	"fmt"
	"net/http"
	"slices"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func applyAlteredStateFromQuery(r *http.Request, mon apitypes.Monster, queryName ts.QueryParamName) (apitypes.Monster, error) {
	altStateID, err := getAltStateID(r, mon, queryName)
	if errs.QueryIsEmpty(err) {
		return mon, nil
	}
	if err != nil {
		return apitypes.Monster{}, err
	}

	return applyAlteredState(mon, altStateID), nil
}

func ApplyAlteredStateFromJSON(mon apitypes.Monster, altStatePtr *int32) (apitypes.Monster, error) {
	if altStatePtr == nil {
		return mon, nil
	}

	altStateID := int(*altStatePtr)

	err := vpAltStateID(mon, altStateID)
	if err != nil {
		return apitypes.Monster{}, err
	}

	return applyAlteredState(mon, altStateID), nil
}

func applyAlteredState(mon apitypes.Monster, altStateID int) apitypes.Monster {
	altState := mon.AlteredStates[altStateID-1]
	appliedState := apitypes.AppliedState{
		Condition:   altState.Condition,
		IsTemporary: altState.IsTemporary,
	}

	defaultState := apitypes.AlteredState{
		URL:         ah.CreateResourceURL(apilookups.E.Monsters.Endpoint, mon.ID),
		Condition:   "default",
		IsTemporary: false,
	}

	for _, change := range altState.Alts {
		switch database.AlterationType(change.AlterationType) {
		case database.AlterationTypeChange:
			mon, appliedState, defaultState = applyAltChange(mon, change, appliedState, defaultState)

		case database.AlterationTypeGain:
			mon, appliedState, defaultState = applyAltGain(mon, change, appliedState, defaultState)

		case database.AlterationTypeLoss:
			mon, appliedState, defaultState = applyAltLoss(mon, change, appliedState, defaultState)
		}
	}

	mon.AppliedState = &appliedState
	if appliedState.IsZero() {
		mon.AppliedState = nil
	}

	mon.AlteredStates = replaceAltState(mon.AlteredStates, defaultState, altStateID)

	return mon
}

func getAltStateID(r *http.Request, mon apitypes.Monster, queryName ts.QueryParamName) (int, error) {
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

func vpAltStateID(mon apitypes.Monster, altStateID int) error {
	altStateAmt := len(mon.AlteredStates)

	if altStateAmt == 0 {
		return errs.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%s has no altered states", mon.String()), nil)
	}

	if altStateID > altStateAmt || altStateID <= 0 {
		return errs.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("provided alt state id '%d' used for %s is out of range. max id: %d", altStateID, mon.String(), altStateAmt), nil)
	}

	return nil
}

func applyAltChange(mon apitypes.Monster, change apitypes.Alt, appliedState apitypes.AppliedState, defaultState apitypes.AlteredState) (apitypes.Monster, apitypes.AppliedState, apitypes.AlteredState) {
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

	return mon, appliedState, defaultState
}

func applyAltGain(mon apitypes.Monster, change apitypes.Alt, appliedState apitypes.AppliedState, defaultState apitypes.AlteredState) (apitypes.Monster, apitypes.AppliedState, apitypes.AlteredState) {
	defStateLoss := apitypes.Alt{
		AlterationType: database.AlterationTypeLoss,
	}
	defStateGain := apitypes.Alt{
		AlterationType: database.AlterationTypeGain,
	}

	mon.Properties, defStateLoss.Properties = modifyResourcesGain(mon.Properties, change.Properties)
	mon.AutoAbilities, defStateLoss.AutoAbilities = modifyResourcesGain(mon.AutoAbilities, change.AutoAbilities)

	mon, defStateGain, defStateLoss, appliedState = modifyGainedImmunities(mon, change, defStateGain, defStateLoss, appliedState)

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

	return mon, appliedState, defaultState
}

func applyAltLoss(mon apitypes.Monster, change apitypes.Alt, appliedState apitypes.AppliedState, defaultState apitypes.AlteredState) (apitypes.Monster, apitypes.AppliedState, apitypes.AlteredState) {
	defStateGain := apitypes.Alt{
		AlterationType: database.AlterationTypeGain,
	}

	mon.Properties, defStateGain.Properties = modifyResourcesLoss(mon.Properties, change.Properties)
	mon.AutoAbilities, defStateGain.AutoAbilities = modifyResourcesLoss(mon.AutoAbilities, change.AutoAbilities)

	if !defStateGain.IsZero() {
		defaultState.Alts = append(defaultState.Alts, defStateGain)
	}

	return mon, appliedState, defaultState
}

// put default in first and cut out the currently applied state
func replaceAltState(states []apitypes.AlteredState, def apitypes.AlteredState, i int) []apitypes.AlteredState {
	allStates := h.Unshift(states, def)
	size := len(allStates)
	s1 := allStates[0:i]
	s2 := allStates[i+1 : size]
	return slices.Concat(s1, s2)
}
