package main

import (
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type HasAPIResource interface {
	getAPIResource() IsAPIResource
}

type IsAPIResource interface {
	getID() int32
	getURL() string
	seeding.Lookupable
}



// input: items [1,2,3,4,5] changeItems [2,4]
// output: keptItems [1,3,5] removedItems [2,4]
// due to the interfaces,
// everything with the same namedAPIResources in it
// can filter each other
func removeResources[T HasAPIResource, C HasAPIResource](items []T, itemsToRemove []C) ([]T, []T) {
	removeMap := getResourceMap(itemsToRemove)
	keptItems := []T{}
	removedItems := []T{}

	for _, item := range items {
		key := createAPIResourceKey(item)
		_, ok := removeMap[key]
		if !ok {
			keptItems = append(keptItems, item)
			continue
		}
		removedItems = append(removedItems, item)
	}

	return keptItems, removedItems
}

// s1 [1,2,3,4,5] s2 [2,4,5,7,8,9] => [2,5]
func getSharedResources[T HasAPIResource](s1 []T, s2 []T) []T {
	newSlice := []T{}
	s2map := getResourceMap(s2)

	for _, item := range s1 {
		key := createAPIResourceKey(item)
		_, ok := s2map[key]
		if ok {
			newSlice = append(newSlice, item)
		}
	}

	return newSlice
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

func sortAPIResources[T HasAPIResource](a, b T) int {
	if a.getAPIResource().getID() < b.getAPIResource().getID() {
		return -1
	}

	if a.getAPIResource().getID() > b.getAPIResource().getID() {
		return 1
	}

	return 0
}
