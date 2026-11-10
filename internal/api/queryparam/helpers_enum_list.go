package queryparam

import (
	"fmt"
	"net/http"
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/enums"
)

// ParseEnumList checks for emptiness of enum-list-queryParam and converts its input into a slice of valid enum-strings.
func (q QueryParam) ParseEnumList[E, N any](r *http.Request, et enums.EnumType[E, N]) ([]E, error) {
	query, err := q.CheckEmpty(r)
	if err != nil {
		return nil, err
	}

	return q.sliceEnums(query, et)
}

// converts a list of unique query enum values or ids into a slice of valid typed enum-strings.
func (q QueryParam) sliceEnums[E, N any](query string, et enums.EnumType[E, N]) ([]E, error) {
	enumStrs, err := QueryListSplit(query)
	if err != nil {
		return nil, err
	}
	enums := []E{}

	for _, enumStr := range enumStrs {
		enum, err := q.CheckEnum(enumStr, et)
		if err != nil {
			return nil, err
		}

		if et.Aliasses != nil {
			aliasVals, ok := et.Aliasses[string(enum.Name)]
			if ok {
				enums = slices.Concat(enums, aliasVals)
				continue
			}
		}

		typedStr := et.ConvFunc(enum.Name)
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
