package helpers

import (
	"slices"
)

func Unshift[T any] (items[]T, added T) []T {
	newSlice := []T{added}
	newSlice = slices.Concat(newSlice, items)

	return newSlice
}


func CombineIdSlices(ids []int32, idSlices ...[]int32) []int32 {
	for _, slice := range idSlices {
		ids = addUniqueIDsToSlice(ids, slice)
	}

	return ids
}

func addUniqueIDsToSlice(s1, s2 []int32) []int32 {
	idMap := make(map[int32]bool)

	for _, id := range s1 {
		idMap[id] = true
	}

	for _, id := range s2 {
		_, ok := idMap[id]
		if !ok {
			s1 = append(s1, id)
		}
	}

	slices.Sort(s1)
	return s1
}