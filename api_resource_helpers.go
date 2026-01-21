package main

import (
	"slices"
)

// s1 [1,2,3] s2 [2,3,4,5] => [1,2,3,4,5]
func combineResources[T HasAPIResource](s1, s2 []T) []T {
	s1Map := getResourceMap(s1)

	for _, item := range s2 {
		key := getAPIResourceID(item)
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
		id := getAPIResourceID(item)
		_, ok := removeMap[id]
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
		id := getAPIResourceID(item)
		_, ok := s2map[id]
		if ok {
			sharedItems = append(sharedItems, item)
		}
	}

	return sharedItems
}

func resourcesContain[T HasAPIResource](items []T, target T) bool {
	for _, item := range items {
		if item.GetAPIResource() == target.GetAPIResource() {
			return true
		}
	}
	return false
}

func getAPIResourceID[T HasAPIResource](item T) int32 {
	resource := item.GetAPIResource()
	return resource.GetID()
}

func getResourceMap[T HasAPIResource](items []T) map[int32]T {
	resourceMap := make(map[int32]T)

	for _, item := range items {
		id := getAPIResourceID(item)
		resourceMap[id] = item
	}

	return resourceMap
}

func getResourceURLMap[T HasAPIResource](items []T) map[string]T {
	resourceMap := make(map[string]T)

	for _, item := range items {
		url := item.GetAPIResource().GetURL()
		resourceMap[url] = item
	}

	return resourceMap
}

func resourceMapToSlice[T HasAPIResource](lookup map[int32]T) []T {
	s := []T{}

	for id := range lookup {
		s = append(s, lookup[id])
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
	if getAPIResourceID(a) < getAPIResourceID(b) {
		return -1
	}

	if getAPIResourceID(a) > getAPIResourceID(b) {
		return 1
	}

	return 0
}

func toHasAPIResSlice[T HasAPIResource](s []T) []HasAPIResource {
	out := make([]HasAPIResource, len(s))
	for i, v := range s {
		out[i] = v
	}
	return out
}
