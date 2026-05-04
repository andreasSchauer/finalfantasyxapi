package api

import (
	"slices"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type filteredIdList struct {
	IDs []int32
	err error
}

func fidl(ids []int32, err error) filteredIdList {
	return filteredIdList{
		IDs: ids,
		err: err,
	}
}

func combineIdSlices(idSlices ...[]int32) []int32 {
	ids := []int32{}

	for _, slice := range idSlices {
		ids = addUniqueIDsToSlice(ids, slice)
	}

	return ids
}

func filterIdSlices(filteredLists []filteredIdList) ([]int32, error) {
	switch len(filteredLists) {
	case 0:
		return []int32{}, nil

	case 1:
		return filteredLists[0].IDs, nil

	default:
		ids := filteredLists[0].IDs

		for _, filtered := range filteredLists {
			if filtered.err != nil {
				return nil, filtered.err
			}
			ids = getSharedIDs(ids, filtered.IDs)
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

func sortNamesByID[T h.HasID](s []string, lookup map[string]T) []string {
	slices.SortStableFunc(s, func(a, b string) int {
		A, _ := seeding.GetResource(a, lookup)
		B, _ := seeding.GetResource(b, lookup)

		if A.GetID() < B.GetID() {
			return -1
		}

		if A.GetID() > B.GetID() {
			return 1
		}

		return 0
	})

	return s
}
