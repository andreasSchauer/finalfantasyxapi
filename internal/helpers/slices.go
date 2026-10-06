package helpers

import (
	"slices"
)

func Unshift[T any](items []T, added T) []T {
	newSlice := []T{added}
	newSlice = slices.Concat(newSlice, items)

	return newSlice
}

func Filter[T any](s []T, fn func(T) bool) []T {
	if s == nil {
		return nil
	}

	newSlice := []T{}

	for _, item := range s {
		if fn(item) {
			newSlice = append(newSlice, item)
		}
	}

	return newSlice
}


func GetNumSlice(min, max int32) []int32 {
	var s []int32

	for i := min; i <= max; i++ {
		s = append(s, i)
	}

	return s
}