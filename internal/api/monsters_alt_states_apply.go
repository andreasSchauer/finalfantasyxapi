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

func applyAlteredStateFromQuery(cfg *Config, r *http.Request, mon Monster, queryName QueryParamName) (Monster, error) {
	altStateID, err := getAltStateID(cfg, r, mon, queryName)
	if queryIsEmpty(err) {
		return mon, nil
	}
	if err != nil {
		return Monster{}, err
	}

	return applyAlteredState(cfg, mon, altStateID), nil
}

func applyAlteredStateFromJson(cfg *Config, mon Monster, altStatePtr *int32) (Monster, error) {
	if altStatePtr == nil {
		return mon, nil
	}

	altStateID := int(*altStatePtr)

	err := vpAltStateID(mon, altStateID)
	if err != nil {
		return Monster{}, err
	}

	return applyAlteredState(cfg, mon, altStateID), nil
}

func applyAlteredState(cfg *Config, mon Monster, altStateID int) Monster {
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

func getAltStateID(cfg *Config, r *http.Request, mon Monster, queryName QueryParamName) (int, error) {
	queryParam := cfg.q.monsters[queryName]
	query, err := checkEmptyQuery(r, queryParam)
	if err != nil {
		return 0, err
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


func vpAltStateID(mon Monster, altStateID int) error {
	altStateAmt := len(mon.AlteredStates)

	if altStateAmt == 0 {
		return newHTTPError(http.StatusBadRequest, fmt.Sprintf("%s has no altered states.", mon.Error()), nil)
	}

	if altStateID > altStateAmt || altStateID <= 0 {
		return newHTTPError(http.StatusBadRequest, fmt.Sprintf("provided alt state id '%d' used for %s is out of range. max id: %d.", altStateID, mon.Error(), altStateAmt), nil)
	}

	return nil
}

func applyAltChange(mon Monster, change Alt, appliedState AppliedState, defaultState AlteredState) (Monster, AppliedState, AlteredState) {
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

func applyAltGain(mon Monster, change Alt, appliedState AppliedState, defaultState AlteredState) (Monster, AppliedState, AlteredState) {
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

func applyAltLoss(mon Monster, change Alt, appliedState AppliedState, defaultState AlteredState) (Monster, AppliedState, AlteredState) {
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
