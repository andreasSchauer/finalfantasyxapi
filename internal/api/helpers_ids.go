package api

import (
	"slices"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

// s1 [1,2,3] s2 [2,3,4,5] => [1,2,3,4,5]
func combineIdSlices(idSlices ...[]int32) []int32 {
	ids := []int32{}

	for _, slice := range idSlices {
		ids = addUniqueIDsToSlice(ids, slice)
	}

	return ids
}

func filterIdSlices(slices [][]int32) ([]int32, error) {
	switch len(slices) {
	case 0:
		return []int32{}, nil

	case 1:
		return slices[0], nil

	default:
		ids := slices[0]

		for i := 1; i < len(slices); i++ {
			ids = getSharedIDs(ids, slices[i])
		}

		return ids, nil
	}
}

func addUniqueIDsToSlice(s1, s2 []int32) []int32 {
	idMap := getIdMap(s1)

	for _, id := range s2 {
		_, ok := idMap[id]
		if !ok {
			s1 = append(s1, id)
		}
	}

	slices.Sort(s1)
	return s1
}

// s1 [1,2,3,4,5] s2 [2,4,5,7,8,9] => [2,4,5]
func getSharedIDs(s1, s2 []int32) []int32 {
	sharedIDs := []int32{}
	s2Map := getIdMap(s2)

	for _, id := range s1 {
		_, ok := s2Map[id]
		if ok {
			sharedIDs = append(sharedIDs, id)
		}
	}

	slices.Sort(sharedIDs)
	return sharedIDs
}

func getIdMap(s []int32) map[int32]bool {
	idMap := make(map[int32]bool)

	for _, id := range s {
		idMap[id] = true
	}

	return idMap
}

func sortNamesByID[T seeding.Lookupable](s []string, lookup map[string]T) []string {
	slices.SortStableFunc(s, func(a, b string) int {
		A, _ := seeding.GetResource(a, lookup)
		B, _ := seeding.GetResource(b, lookup)

		return h.SortOnId(A, B)
	})

	return s
}

// ids [1,2,3,4,5] idsToRemove [2,4] => keptIDs [1,3,5] removedIDs [2,4]
func separateIDs(ids []int32, idsToRemove []int32) ([]int32, []int32) {
	removeMap := getIdMap(idsToRemove)
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

// items [1,2,3,4,5] changeItems [2,4] => [1,3,5]
func removeIDs(items []int32, itemsToRemove []int32) []int32 {
	keptIDs, _ := separateIDs(items, itemsToRemove)
	return keptIDs
}
