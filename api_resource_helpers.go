package main

import (
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

// s1 [1,2,3] s2 [2,3,4,5] => [1,2,3,4,5]
func combineResources[T HasAPIResource](s1, s2 []T) []T {
	s1Map := getResourceMap(s1)

	for _, item := range s2 {
		key := createAPIResourceKey(item)
		_, ok := s1Map[key]
		if !ok {
			s1Map[key] = item
		}
	}

	return resourceMapToSlice(s1Map)
}

// items [1,2,3,4,5] changeItems [2,4] => keptItems [1,3,5] removedItems [2,4]
func separateResources[T HasAPIResource, C HasAPIResource](items []T, itemsToRemove []C) ([]T, []T) {
	removeMap := getResourceMap(itemsToRemove)
	kept := []T{}
	removed := []T{}

	for _, item := range items {
		key := createAPIResourceKey(item)
		_, ok := removeMap[key]
		if !ok {
			kept = append(kept, item)
			continue
		}
		removed = append(removed, item)
	}

	return kept, removed
}

// items [1,2,3,4,5] changeItems [2,4] => [1,3,5]
func removeResources[T HasAPIResource, C HasAPIResource](items []T, itemsToRemove []C) []T {
	keptItems, _ := separateResources(items, itemsToRemove)
	return keptItems
}

// s1 [1,2,3,4,5] s2 [2,4,5,7,8,9] => [2,4,5]
func getSharedResources[T HasAPIResource](s1, s2 []T) []T {
	sharedItems := []T{}
	s2map := getResourceMap(s2)

	for _, item := range s1 {
		key := createAPIResourceKey(item)
		_, ok := s2map[key]
		if ok {
			sharedItems = append(sharedItems, item)
		}
	}

	return sharedItems
}

func resourcesContain[T HasAPIResource](items []T, target T) bool {
	for _, item := range items {
		if item.getAPIResource() == target.getAPIResource() {
			return true
		}
	}
	return false
}

func createAPIResourceKey[T HasAPIResource](item T) string {
	resource := item.getAPIResource()
	key := seeding.CreateLookupKey(resource)
	return key
}

func getResourceMap[T HasAPIResource](items []T) map[string]T {
	resourceMap := make(map[string]T)

	for _, item := range items {
		key := createAPIResourceKey(item)
		resourceMap[key] = item
	}

	return resourceMap
}

func resourceMapToSlice[T HasAPIResource](lookup map[string]T) []T {
	s := []T{}

	for key := range lookup {
		s = append(s, lookup[key])
	}

	slices.SortStableFunc(s, sortAPIResources)

	return s
}

func getResourceAmountMap[T ResourceAmount](items []T) map[string]int32 {
	amountMap := make(map[string]int32)

	for _, item := range items {
		key := item.GetName()
		amountMap[key] = item.GetVal()
	}

	return amountMap
}

func sortAPIResources[T HasAPIResource](a, b T) int {
	if a.getAPIResource().getID() < b.getAPIResource().getID() {
		return -1
	}

	if a.getAPIResource().getID() > b.getAPIResource().getID() {
		return 1
	}

	return 0
}


func toHasAPIResSlice[T HasAPIResource](in []T) []HasAPIResource {
	out := make([]HasAPIResource, len(in))
	for i, v := range in {
		out[i] = v
	}
	return out
}