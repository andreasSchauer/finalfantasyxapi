package api

import (
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

func applyAlteredState(cfg *Config, r *http.Request, mon Monster, queryName string) (Monster, error) {
	altStateID, err := getAltStateID(cfg, r, mon, queryName)
	if queryIsEmpty(err) {
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
		URL:         createResourceURL(cfg, cfg.e.monsters.endpoint, mon.ID),
		Condition:   "default",
		IsTemporary: false,
	}

	for _, change := range altState.Alts {
		switch database.AlterationType(change.AlterationType) {
		case database.AlterationTypeChange:
			mon, appliedState, defaultState = applyAlt(mon, change, appliedState, defaultState)

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

func getAltStateID(cfg *Config, r *http.Request, mon Monster, queryName string) (int, error) {
	queryParam := cfg.q.monsters[queryName]
	query := r.URL.Query().Get(queryParam.Name)

	if query == "" {
		return 0, errEmptyQuery
	}

	if len(mon.AlteredStates) == 0 {
		return 0, newHTTPError(http.StatusBadRequest, fmt.Sprintf("%s has no altered states.", mon.Error()), nil)
	}

	id, err := parseQueryID(query, queryParam, len(mon.AlteredStates))
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func applyAlt(mon Monster, change Alt, appliedState AppliedState, defaultState AlteredState) (Monster, AppliedState, AlteredState) {
	defStateChange := Alt{
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

func applyAltStateGain(mon Monster, change Alt, appliedState AppliedState, defaultState AlteredState) (Monster, AppliedState, AlteredState) {
	defStateLoss := Alt{
		AlterationType: database.AlterationTypeLoss,
	}
	defStateGain := Alt{
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

func applyAltStateLoss(mon Monster, change Alt, appliedState AppliedState, defaultState AlteredState) (Monster, AppliedState, AlteredState) {
	defStateGain := Alt{
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
func replaceAltState(states []AlteredState, def AlteredState, i int) []AlteredState {
	allStates := h.Unshift(states, def)
	size := len(allStates)
	s1 := allStates[0:i]
	s2 := allStates[i+1 : size]
	return slices.Concat(s1, s2)
}
