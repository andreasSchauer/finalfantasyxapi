package helpers

import (
    "cmp"
    "slices"
    "strings"
)

func Capitalize(s string) string {
    if s == "" {
        return s
    }
    return strings.ToUpper(s[:1]) + s[1:]
}


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