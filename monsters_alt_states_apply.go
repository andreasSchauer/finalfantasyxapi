package main

import (
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type AppliedState struct {
	Condition     string           `json:"condition"`
	IsTemporary   bool             `json:"is_temporary"`
	AppliedStatus *InflictedStatus `json:"applied_status,omitempty"`
}

func (as AppliedState) IsZero() bool {
	return as.Condition == ""
}

func (cfg *Config) applyAlteredState(r *http.Request, mon Monster) (Monster, error) {
	altStateID, err := cfg.getAltStateID(r, mon)
	if errors.Is(err, errEmptyQuery) {
		return mon, nil
	}
	if err != nil {
		return Monster{}, err
	}

	altState := mon.AlteredStates[altStateID-1]
	appliedState := AppliedState{
		Condition:   altState.Condition,
		IsTemporary: altState.IsTemporary,
	}

	defaultState := AlteredState{
		URL:         cfg.createResourceURL(cfg.e.monsters.endpoint, mon.ID),
		Condition:   "default",
		IsTemporary: false,
	}

	for _, change := range altState.Changes {
		switch database.AlterationType(change.AlterationType) {
		case database.AlterationTypeChange:
			mon, appliedState, defaultState = applyAltStateChange(mon, change, appliedState, defaultState)

		case database.AlterationTypeGain:
			mon, appliedState, defaultState = applyAltStateGain(mon, change, appliedState, defaultState)

		case database.AlterationTypeLoss:
			mon, appliedState, defaultState = applyAltStateLoss(mon, change, appliedState, defaultState)
		}
	}

	mon.AppliedState = &appliedState

	if appliedState.IsZero() {
		mon.AppliedState = nil
	}

	mon.AlteredStates = replaceAltState(mon.AlteredStates, defaultState, altStateID)

	return mon, nil
}

func (cfg *Config) getAltStateID(r *http.Request, mon Monster) (int, error) {
	queryParam := cfg.q.monsters["altered_state"]
	query := r.URL.Query().Get(queryParam.Name)

	if query == "" {
		return 0, errEmptyQuery
	}

	if len(mon.AlteredStates) == 0 {
		return 0, newHTTPError(http.StatusBadRequest, fmt.Sprintf("%s has no altered states.", mon.Error()), nil)
	}

	id, err := parseQueryIdVal(query, queryParam, len(mon.AlteredStates))
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func applyAltStateChange(mon Monster, change AltStateChange, appliedState AppliedState, defaultState AlteredState) (Monster, AppliedState, AlteredState) {
	defStateChange := AltStateChange{
		AlterationType: database.AlterationTypeChange,
	}

	if change.Distance != nil {
		newDistance := mon.Distance
		defStateChange.Distance = &newDistance
		mon.Distance = *change.Distance
	}

	mon.BaseStats, defStateChange.BaseStats = modifyResourcesChange(mon.BaseStats, change.BaseStats)
	mon.ElemResists, defStateChange.ElemResists = modifyResourcesChange(mon.ElemResists, change.ElemResists)

	defaultState.Changes = append(defaultState.Changes, defStateChange)

	return mon, appliedState, defaultState
}

func applyAltStateGain(mon Monster, change AltStateChange, appliedState AppliedState, defaultState AlteredState) (Monster, AppliedState, AlteredState) {
	defStateLoss := AltStateChange{
		AlterationType: database.AlterationTypeLoss,
	}
	defStateGain := AltStateChange{
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
		defaultState.Changes = append(defaultState.Changes, defStateLoss)
	}

	if !defStateGain.IsZero() {
		defaultState.Changes = append(defaultState.Changes, defStateGain)
	}

	return mon, appliedState, defaultState
}

func applyAltStateLoss(mon Monster, change AltStateChange, appliedState AppliedState, defaultState AlteredState) (Monster, AppliedState, AlteredState) {
	defStateGain := AltStateChange{
		AlterationType: database.AlterationTypeGain,
	}

	mon.Properties, defStateGain.Properties = modifyResourcesLoss(mon.Properties, change.Properties)
	mon.AutoAbilities, defStateGain.AutoAbilities = modifyResourcesLoss(mon.AutoAbilities, change.AutoAbilities)

	if !defStateGain.IsZero() {
		defaultState.Changes = append(defaultState.Changes, defStateGain)
	}

	return mon, appliedState, defaultState
}

// put default in first and cut out the currently applied state
func replaceAltState(states []AlteredState, def AlteredState, i int) []AlteredState {
	allStates := h.Unshift(states, def)
	size := len(allStates)
	s1 := allStates[0:i]
	s2 := allStates[i+1 : size]
	return slices.Concat(s1, s2)
}
