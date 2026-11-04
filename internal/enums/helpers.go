package enums

import (
	"cmp"
	"slices"
)

func EnumLookupToSlice(lookup map[string]EnumResponse) []EnumResponse {
	enums := make([]EnumResponse, 0, len(lookup))

	for key := range lookup {
		enums = append(enums, lookup[key])
	}

	slices.SortFunc(enums, func(a, b EnumResponse) int {
		return cmp.Compare(a.Name, b.Name)
	})
	return enums
}