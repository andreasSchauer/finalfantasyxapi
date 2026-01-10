package helpers

import (
    "cmp"
    "slices"
    "strings"
)


func GetMapKeyStr[T any](itemMap map[string]T) string {
	keys := []string{}

	for key := range itemMap {
		keys = append(keys, key)
	}

	slices.SortStableFunc(keys, func(a, b string) int {
		return cmp.Compare(a, b)
	})

	return strings.Join(keys, ", ")
}



func GetNameWithSpaces(name string) string {
	nameWithSpaces := strings.ReplaceAll(name, "-", " ")
	return strings.ReplaceAll(nameWithSpaces, " >", "->")
}