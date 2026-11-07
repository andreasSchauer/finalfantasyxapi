package handlerinput

import (
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
)

// CombineResources does this: s1 [1,2,3] s2 [2,3,4,5] => [1,2,3,4,5]
func CombineResources[T apiresource.HasAPIResource](s1, s2 []T) []T {
	s1Map := GetResourceMap(s1)

	for _, item := range s2 {
		key := GetAPIResourceID(item)
		_, ok := s1Map[key]
		if !ok {
			s1Map[key] = item
		}
	}

	return ResourceMapToSlice(s1Map)
}

// SeparateResources does this: items [1,2,3,4,5] changeItems [2,4] => keptItems [1,3,5] removedItems [2,4]
func SeparateResources[T apiresource.HasAPIResource, C apiresource.HasAPIResource](items []T, itemsToRemove []C) ([]T, []T) {
	removeMap := GetResourceMap(itemsToRemove)
	kept := []T{}
	removed := []T{}

	for _, item := range items {
		id := GetAPIResourceID(item)
		_, ok := removeMap[id]
		if !ok {
			kept = append(kept, item)
			continue
		}
		removed = append(removed, item)
	}

	return kept, removed
}

// RemoveResources does this: items [1,2,3,4,5] changeItems [2,4] => [1,3,5]
func RemoveResources[T apiresource.HasAPIResource, C apiresource.HasAPIResource](items []T, itemsToRemove []C) []T {
	keptItems, _ := SeparateResources(items, itemsToRemove)
	return keptItems
}

func SeparateResourcesURL[T apiresource.HasAPIResource, C apiresource.HasAPIResource](items []T, itemsToRemove []C) ([]T, []T) {
	removeMap := GetResourceURLMap(itemsToRemove)
	kept := []T{}
	removed := []T{}

	for _, item := range items {
		url := item.GetAPIResource().GetURL()
		_, ok := removeMap[url]
		if !ok {
			kept = append(kept, item)
			continue
		}
		removed = append(removed, item)
	}

	return kept, removed
}

func RemoveResourcesURL[T apiresource.HasAPIResource, C apiresource.HasAPIResource](items []T, itemsToRemove []C) []T {
	keptItems, _ := SeparateResourcesURL(items, itemsToRemove)
	return keptItems
}

// GetSharedResources does this: s1 [1,2,3,4,5] s2 [2,4,5,7,8,9] => [2,4,5]
func GetSharedResources[T apiresource.HasAPIResource](s1, s2 []T) []T {
	sharedItems := []T{}
	s2map := GetResourceMap(s2)

	for _, item := range s1 {
		id := GetAPIResourceID(item)
		_, ok := s2map[id]
		if ok {
			sharedItems = append(sharedItems, item)
		}
	}

	return sharedItems
}

func ResourcesContain[T apiresource.HasAPIResource](items []T, target T) bool {
	for _, item := range items {
		if item.GetAPIResource() == target.GetAPIResource() {
			return true
		}
	}
	return false
}

func GetAPIResourceID[T apiresource.HasAPIResource](item T) int32 {
	resource := item.GetAPIResource()
	return resource.GetID()
}

func GetResourceMap[T apiresource.HasAPIResource](items []T) map[int32]T {
	resourceMap := make(map[int32]T, len(items))

	for _, item := range items {
		id := GetAPIResourceID(item)
		resourceMap[id] = item
	}

	return resourceMap
}

func GetResourceURLMap[T apiresource.HasAPIResource](items []T) map[string]T {
	resourceMap := make(map[string]T, len(items))

	for _, item := range items {
		url := item.GetAPIResource().GetURL()
		resourceMap[url] = item
	}

	return resourceMap
}

func GetSimpleResourceURLMap[T apiresource.SimpleResource](items []T) map[string]T {
	resourceMap := make(map[string]T, len(items))

	for _, item := range items {
		url := item.GetURL()
		resourceMap[url] = item
	}

	return resourceMap
}

func ResourceMapToSlice[T apiresource.HasAPIResource](lookup map[int32]T) []T {
	s := []T{}

	for id := range lookup {
		s = append(s, lookup[id])
	}

	slices.SortStableFunc(s, SortAPIResources)

	return s
}

func SortAPIResources[T apiresource.HasAPIResource](a, b T) int {
	if GetAPIResourceID(a) < GetAPIResourceID(b) {
		return -1
	}

	if GetAPIResourceID(a) > GetAPIResourceID(b) {
		return 1
	}

	return 0
}
