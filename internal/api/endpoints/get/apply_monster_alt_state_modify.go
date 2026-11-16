package get

import (
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

// remove lost items and put them into defStateGain
func modifyResourcesLoss[T apiresource.HasAPIResource](items, changeItems []T) ([]T, []T) {
	if changeItems == nil {
		return items, nil
	}

	defStateGainItems := []T{}
	itemsToRemove := []T{}

	for _, item := range changeItems {
		itemsToRemove = append(itemsToRemove, item)
		defStateGainItems = append(defStateGainItems, item)
	}

	items = apiresource.RemoveResources(items, itemsToRemove)

	slices.SortStableFunc(items, apiresource.SortAPIResources)
	slices.SortStableFunc(defStateGainItems, apiresource.SortAPIResources)

	return items, defStateGainItems
}

// add gained items, but also put them into defStateLoss
func modifyResourcesGain[T apiresource.HasAPIResource](items, changeItems []T) ([]T, []T) {
	if changeItems == nil {
		return items, nil
	}

	defStateLossItems := []T{}

	for _, item := range changeItems {
		items = append(items, item)
		defStateLossItems = append(defStateLossItems, item)
	}

	slices.SortStableFunc(items, apiresource.SortAPIResources)
	slices.SortStableFunc(defStateLossItems, apiresource.SortAPIResources)

	return items, defStateLossItems
}

// replace the items with those of the changeState
func modifyResourcesChange[T apiresource.HasAPIResource](items, changeItems []T) ([]T, []T) {
	if changeItems == nil {
		return items, changeItems
	}

	defStateChangeItems := []T{}
	replaceMap := apiresource.GetResourceMap(changeItems)

	for i, item := range items {
		key := apiresource.GetAPIResourceID(item)
		_, ok := replaceMap[key]
		if ok {
			defStateChangeItems = append(defStateChangeItems, item)
			items[i] = replaceMap[key]
		}
	}

	return items, defStateChangeItems
}

// if a resistance becomes an immunity, the resistance needs to be added to defStateGain
// while the immunity needs to be added to defStateLoss
func modifyGainedImmunities(mon apitypes.Monster, change, defStateGain, defStateLoss apitypes.Alt, appliedState apitypes.AppliedState) (apitypes.Monster, apitypes.Alt, apitypes.Alt, apitypes.AppliedState) {
	if change.StatusImmunities == nil {
		return mon, defStateGain, defStateLoss, appliedState
	}

	mon.StatusImmunities, defStateLoss.StatusImmunities = modifyResourcesGain(mon.StatusImmunities, change.StatusImmunities)

	keptItems, removedItems := apiresource.SeparateResources(mon.StatusResists, change.StatusImmunities)

	if len(removedItems) == 0 {
		removedItems = nil
	}

	defStateGain.StatusResists = removedItems
	mon.StatusResists = keptItems

	if change.AddedStatus != nil {
		status := change.AddedStatus
		appliedState.AppliedStatus = status
		defStateLoss.RemovedStatus = &status.StatusCondition
	}

	return mon, defStateGain, defStateLoss, appliedState
}
