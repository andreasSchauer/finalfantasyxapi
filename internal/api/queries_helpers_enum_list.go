package api

import (
	"fmt"
	"net/http"
)

// checks for emptiness of enum-list-queryParam and converts its input into a slice of valid enum-strings.
func parseEnumListQuery[E, N any](r *http.Request, endpoint string, queryParam QueryType, et EnumType[E, N]) ([]E, error) {
	query, err := checkEmptyQuery(r, queryParam)
	if err != nil {
		return nil, err
	}

	return queryEnumsToSlice(query, endpoint, queryParam, et)
}

// converts a list of unique query enum values or ids into a slice of valid typed enum-strings.
func queryEnumsToSlice[E, N any](query, endpoint string, queryParam QueryType, et EnumType[E, N]) ([]E, error) {
	enumStrs := querySplit(query, ",")
	enums := []E{}

	for _, enumStr := range enumStrs {
		enum, err := checkQueryEnum(enumStr, endpoint, queryParam, et)
		if err != nil {
			return nil, err
		}
		typedStr := et.convFunc(enum.Name)
		enums = append(enums, typedStr)
	}

	err := checkDuplicateEnums(queryParam, enums)
	if err != nil {
		return nil, err
	}

	return enums, nil
}

func checkDuplicateEnums[E any](queryParam QueryType, enums []E) error {
	enumMap := make(map[any]bool)

	for _, enum := range enums {
		if enumMap[enum] {
			return newHTTPError(http.StatusBadRequest, fmt.Sprintf("duplicate use of enum '%v' for parameter '%s'. each enum can only be used once.", enum, queryParam.Name), nil)
		}
		enumMap[enum] = true
	}

	return nil
}
