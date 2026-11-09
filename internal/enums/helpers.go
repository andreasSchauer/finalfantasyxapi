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


func enumSliceToMap(enumVals []EnumVal) map[string]EnumVal {
	typeMap := make(map[string]EnumVal)

	for i, enumVal := range enumVals {
		typeMap[enumVal.Name] = EnumVal{
			ID:          int32(i + 1),
			Name:        enumVal.Name,
			Description: enumVal.Description,
		}
	}

	return typeMap
}

func getEnumValIDs(enumVals []EnumVal) []EnumVal {
	for i := range enumVals {
		enumVals[i].ID = int32(i + 1)
	}

	return enumVals
}