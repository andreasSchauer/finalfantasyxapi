package api

import (
	"slices"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

// CombineIDSlices does this: s1 [1,2,3] s2 [2,3,4,5] => [1,2,3,4,5]
func CombineIDSlices(idSlices ...[]int32) []int32 {
	ids := []int32{}

	for _, slice := range idSlices {
		ids = AddUniqueIDsToSlice(ids, slice)
	}

	return ids
}

func FilterIDSlices(slices [][]int32) ([]int32, error) {
	switch len(slices) {
	case 0:
		return []int32{}, nil

	case 1:
		return slices[0], nil

	default:
		ids := slices[0]

		for i := 1; i < len(slices); i++ {
			ids = GetSharedIDs(ids, slices[i])
		}

		return ids, nil
	}
}

func AddUniqueIDsToSlice(s1, s2 []int32) []int32 {
	idMap := GetIDMap(s1)

	for _, id := range s2 {
		_, ok := idMap[id]
		if !ok {
			s1 = append(s1, id)
		}
	}

	slices.Sort(s1)
	return s1
}

// GetSharedIDs does this: s1 [1,2,3,4,5] s2 [2,4,5,7,8,9] => [2,4,5]
func GetSharedIDs(s1, s2 []int32) []int32 {
	sharedIDs := []int32{}
	s2Map := GetIDMap(s2)

	for _, id := range s1 {
		_, ok := s2Map[id]
		if ok {
			sharedIDs = append(sharedIDs, id)
		}
	}

	slices.Sort(sharedIDs)
	return sharedIDs
}

func GetIDMap(s []int32) map[int32]bool {
	idMap := make(map[int32]bool)

	for _, id := range s {
		idMap[id] = true
	}

	return idMap
}

func SortNamesByID[T seedtypes.Lookupable](s []string, l map[string]T) []string {
	slices.SortStableFunc(s, func(a, b string) int {
		A, _ := lookup.GetResource(a, l)
		B, _ := lookup.GetResource(b, l)

		return h.SortOnID(A, B)
	})

	return s
}

// SeparateIDs does this: ids [1,2,3,4,5] idsToRemove [2,4] => keptIDs [1,3,5] removedIDs [2,4]
func SeparateIDs(ids []int32, idsToRemove []int32) ([]int32, []int32) {
	removeMap := GetIDMap(idsToRemove)
	kept := []int32{}
	removed := []int32{}

	for _, id := range ids {
		_, ok := removeMap[id]
		if !ok {
			kept = append(kept, id)
			continue
		}
		removed = append(removed, id)
	}

	return kept, removed
}

// RemoveIDs does this: items [1,2,3,4,5] changeItems [2,4] => [1,3,5]
func RemoveIDs(items []int32, itemsToRemove []int32) []int32 {
	keptIDs, _ := SeparateIDs(items, itemsToRemove)
	return keptIDs
}
