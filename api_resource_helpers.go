package main

import (
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func filterResourceSlice[T HasAPIResource, C HasAPIResource](items []T, changeItems []C) ([]T, []T) {
	changeItemsMap := getResourceMap(changeItems)
	var filteredItems []T
	var defaultItems []T

	for _, item := range items {
		key := createAPIResourceKey(item)
		_, ok := changeItemsMap[key]
		if !ok {
			filteredItems = append(filteredItems, item)
			continue
		}
		defaultItems = append(defaultItems, item)
	}

	return filteredItems, defaultItems
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

func sortAPIResources[T HasAPIResource](a, b T) int {
	if a.getAPIResource().getID() < b.getAPIResource().getID() {
		return -1
	}

	if a.getAPIResource().getID() > b.getAPIResource().getID() {
		return 1
	}

	return 0
}
