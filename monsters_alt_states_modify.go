package main

import (
	"slices"
)

// remove lost items and put them into defStateGain
func modifyResourcesLoss[T HasAPIResource](items, changeItems []T) ([]T, []T) {
	if changeItems == nil {
		return items, nil
	}

	defStateGainItems := []T{}
	itemsToRemove := []T{}

	for _, item := range changeItems {
		itemsToRemove = append(itemsToRemove, item)
		defStateGainItems = append(defStateGainItems, item)
	}

	items = removeResources(items, itemsToRemove)

	slices.SortStableFunc(items, sortAPIResources)
	slices.SortStableFunc(defStateGainItems, sortAPIResources)

	return items, defStateGainItems
}

// add gained items, but also put them into defStateLoss
func modifyResourcesGain[T HasAPIResource](items, changeItems []T) ([]T, []T) {
	if changeItems == nil {
		return items, nil
	}

	defStateLossItems := []T{}

	for _, item := range changeItems {
		items = append(items, item)
		defStateLossItems = append(defStateLossItems, item)
	}

	slices.SortStableFunc(items, sortAPIResources)
	slices.SortStableFunc(defStateLossItems, sortAPIResources)

	return items, defStateLossItems
}

// replace the items with those of the changeState
func modifyResourcesChange[T HasAPIResource](items, changeItems []T) ([]T, []T) {
	if changeItems == nil {
		return items, changeItems
	}

	defStateChangeItems := []T{}
	replaceMap := getResourceMap(changeItems)

	for i, item := range items {
		key := createAPIResourceKey(item)
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
func modifyGainedImmunities(mon Monster, change, defStateGain, defStateLoss AltStateChange, appliedState AppliedState) (Monster, AltStateChange, AltStateChange, AppliedState) {
	if change.StatusImmunities == nil {
		return mon, defStateGain, defStateLoss, appliedState
	}

	mon.StatusImmunities, defStateLoss.StatusImmunities = modifyResourcesGain(mon.StatusImmunities, change.StatusImmunities)

	keptItems, removedItems := separateResources(mon.StatusResists, change.StatusImmunities)

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
