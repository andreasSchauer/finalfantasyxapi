package main

import (
	"net/http"
	"slices"
	"strconv"

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
	altStateID, err := getAltStateID(r, mon)
	if err != nil {
		return Monster{}, err
	}
	if altStateID == 0 {
		return mon, nil
	}

	altState := mon.AlteredStates[altStateID-1]
	appliedState := AppliedState{
		Condition:   altState.Condition,
		IsTemporary: altState.IsTemporary,
	}

	defaultState := AlteredState{
		URL:         cfg.createURL("monsters", mon.ID),
		Condition:   "default",
		IsTemporary: false,
	}

	for _, change := range altState.Changes {
		switch database.AlterationType(change.AlterationType) {
		case database.AlterationTypeChange:
			mon, appliedState, defaultState = applyAltStateTypeChange(mon, change, appliedState, defaultState)
		case database.AlterationTypeGain:
			mon, appliedState, defaultState = applyAltStateTypeGain(mon, change, appliedState, defaultState)
		case database.AlterationTypeLoss:
			mon, appliedState, defaultState = applyAltStateTypeLoss(mon, change, appliedState, defaultState)
		}
	}

	if appliedState.IsZero() {
		mon.AppliedState = nil
	} else {
		mon.AppliedState = &appliedState
	}

	mon.AlteredStates = replaceAltState(mon.AlteredStates, defaultState, altStateID)

	return mon, nil
}

func applyAltStateTypeLoss(mon Monster, change AltStateChange, appliedState AppliedState, defaultState AlteredState) (Monster, AppliedState, AlteredState) {
	defStateChangeGain := AltStateChange{
		AlterationType: database.AlterationTypeGain,
	}

	gainChangesExist := false

	if change.Properties != nil {
		propertiesToRemove := []NamedAPIResource{}

		for _, property := range change.Properties {
			propertiesToRemove = append(propertiesToRemove, property)
			defStateChangeGain.Properties = append(defStateChangeGain.Properties, property)

			gainChangesExist = true
		}

		mon.Properties = removeResources(mon.Properties, propertiesToRemove)

		slices.SortStableFunc(mon.Properties, sortAPIResources)
		slices.SortStableFunc(defStateChangeGain.Properties, sortAPIResources)
	}

	if change.AutoAbilities != nil {
		abilitiesToRemove := []NamedAPIResource{}

		for _, ability := range change.AutoAbilities {
			abilitiesToRemove = append(abilitiesToRemove, ability)
			defStateChangeGain.AutoAbilities = append(defStateChangeGain.AutoAbilities, ability)

			gainChangesExist = true
		}

		mon.AutoAbilities = removeResources(mon.AutoAbilities, abilitiesToRemove)

		slices.SortStableFunc(mon.AutoAbilities, sortAPIResources)
		slices.SortStableFunc(defStateChangeGain.AutoAbilities, sortAPIResources)
	}

	if gainChangesExist {
		defaultState.Changes = append(defaultState.Changes, defStateChangeGain)
	}

	return mon, appliedState, defaultState
}

func applyAltStateTypeGain(mon Monster, change AltStateChange, appliedState AppliedState, defaultState AlteredState) (Monster, AppliedState, AlteredState) {
	defStateChangeLoss := AltStateChange{
		AlterationType: database.AlterationTypeLoss,
	}

	defStateChangeGain := AltStateChange{
		AlterationType: database.AlterationTypeGain,
	}

	lossChangesExist := false
	gainChangesExist := false

	if change.Properties != nil {
		for _, property := range change.Properties {
			mon.Properties = append(mon.Properties, property)
			defStateChangeLoss.Properties = append(defStateChangeLoss.Properties, property)

			lossChangesExist = true
		}

		slices.SortStableFunc(mon.Properties, sortAPIResources)
		slices.SortStableFunc(defStateChangeLoss.Properties, sortAPIResources)
	}

	if change.AutoAbilities != nil {
		for _, autoAbility := range change.AutoAbilities {
			mon.AutoAbilities = append(mon.AutoAbilities, autoAbility)
			defStateChangeLoss.AutoAbilities = append(defStateChangeLoss.AutoAbilities, autoAbility)

			lossChangesExist = true
		}

		slices.SortStableFunc(mon.AutoAbilities, sortAPIResources)
		slices.SortStableFunc(defStateChangeLoss.AutoAbilities, sortAPIResources)
	}

	// if a resistance becomes an immunity, the resistance needs to be gained back in the default
	// while the immunity needs to be lost in the default
	if change.StatusImmunities != nil {
		for _, immunity := range change.StatusImmunities {
			mon.StatusImmunities = append(mon.StatusImmunities, immunity)
			defStateChangeLoss.StatusImmunities = append(defStateChangeLoss.StatusImmunities, immunity)

			lossChangesExist = true
		}

		slices.SortStableFunc(mon.StatusImmunities, sortAPIResources)
		slices.SortStableFunc(defStateChangeLoss.StatusImmunities, sortAPIResources)

		keptItems, removedItems := separateResources(mon.StatusResists, change.StatusImmunities)

		defStateChangeGain.StatusResistances = removedItems

		if len(defStateChangeGain.StatusResistances) == 0 {
			removedItems = nil
		}

		mon.StatusResists = keptItems

		gainChangesExist = true
	}

	if change.AddedStatus != nil {
		status := change.AddedStatus
		appliedState.AppliedStatus = status
		defStateChangeLoss.RemovedStatus = &status.StatusCondition
	}

	if lossChangesExist {
		defaultState.Changes = append(defaultState.Changes, defStateChangeLoss)
	}

	if gainChangesExist {
		defaultState.Changes = append(defaultState.Changes, defStateChangeGain)
	}

	return mon, appliedState, defaultState
}

// swap old thing with new thing
func applyAltStateTypeChange(mon Monster, change AltStateChange, appliedState AppliedState, defaultState AlteredState) (Monster, AppliedState, AlteredState) {
	defStateChange := AltStateChange{
		AlterationType: database.AlterationTypeChange,
	}

	if change.Distance != nil {
		mon.Distance = *change.Distance
		defStateChange.Distance = &mon.Distance
	}

	if change.BaseStats != nil {
		defStateChange.BaseStats = []BaseStat{}

		for _, newStat := range change.BaseStats {
			for i, oldStat := range mon.BaseStats {
				if newStat.Stat.Name == oldStat.Stat.Name {
					defStateChange.BaseStats = append(defStateChange.BaseStats, oldStat)
					mon.BaseStats[i] = newStat
				}
			}
		}
	}

	if change.ElemResists != nil {
		defStateChange.ElemResists = []ElementalResist{}

		for _, newResist := range change.ElemResists {
			for i, oldResist := range mon.ElemResists {
				if newResist.Element.Name == oldResist.Element.Name {
					defStateChange.ElemResists = append(defStateChange.ElemResists, oldResist)
					mon.ElemResists[i] = newResist
				}
			}
		}
	}
	defaultState.Changes = append(defaultState.Changes, defStateChange)

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

func getAltStateID(r *http.Request, mon Monster) (int, error) {
	altStateIDStr := r.URL.Query().Get("altered-state")
	if altStateIDStr == "" {
		return 0, nil
	}

	altStateID, err := strconv.Atoi(altStateIDStr)
	if err != nil {
		return 0, newHTTPError(http.StatusBadRequest, "invalid alt state id", err)
	}

	if altStateID > len(mon.AlteredStates) || altStateID <= 0 {
		return 0, newHTTPError(http.StatusBadRequest, "provided alt state id does not exist", err)
	}

	return altStateID, nil
}
