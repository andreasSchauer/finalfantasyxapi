package api

import (
	"fmt"
	"net/http"
	"slices"
)

// checks for emptiness of enum-list-queryParam and converts its input into a slice of valid enum-strings.
func parseEnumListQuery[E, N any](cfg *Config, r *http.Request, endpoint string, queryParam QueryParam, et EnumType[E, N]) ([]E, error) {
	query, err := checkEmptyQuery(r, queryParam)
	if err != nil {
		return nil, err
	}

	return queryEnumsToSlice(cfg, query, endpoint, queryParam, et)
}

// converts a list of unique query enum values or ids into a slice of valid typed enum-strings.
func queryEnumsToSlice[E, N any](cfg *Config, query, endpoint string, queryParam QueryParam, et EnumType[E, N]) ([]E, error) {
	enumStrs, err := queryListSplit(cfg, query)
	if err != nil {
		return nil, err
	}
	enums := []E{}

	for _, enumStr := range enumStrs {
		enum, err := checkQueryEnum(enumStr, endpoint, queryParam, et)
		if err != nil {
			return nil, err
		}

		if et.aliasses != nil {
			aliasVals, ok := et.aliasses[string(enum.Name)]
			if ok {
				enums = slices.Concat(enums, aliasVals)
			}
		}

		typedStr := et.convFunc(enum.Name)
		enums = append(enums, typedStr)
	}

	enums = removeDuplicateEnums(enums)

	return enums, nil
}

func removeDuplicateEnums[E any](enums []E) []E {
	enumMap := make(map[any]bool)
	newEnums := []E{}

	for _, enum := range enums {
		if enumMap[enum] {
			continue
		}
		enumMap[enum] = true
		newEnums = append(newEnums, enum)
	}

	return sortEnums(newEnums)
}


func sortEnums[E any](enums []E) []E {
	slices.SortStableFunc(enums, func(a, b E) int {
		A := fmt.Sprint(a)
		B := fmt.Sprint(b)

		if A < B {
			return -1
		}

		if B > A {
			return 1
		}

		return 0
	})

	return enums
}