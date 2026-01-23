package helpers

import (
	"slices"
)

func Unshift[T any] (items[]T, added T) []T {
	newSlice := []T{added}
	newSlice = slices.Concat(newSlice, items)

	return newSlice
}
