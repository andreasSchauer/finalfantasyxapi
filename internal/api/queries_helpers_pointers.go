package api

import (
	"errors"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

// can make those return sql.NullTypes, if the pointers have no other use
// might modify those to return errEmptyQuery, then I can use them for filter queries, if needed

func getQueryBoolPtr(r *http.Request, queryName string, queryLookup map[string]QueryType) (*bool, error) {
	queryParam := queryLookup[queryName]

	b, err := parseBooleanQuery(r, queryParam)
	if errors.Is(err, errEmptyQuery) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &b, nil
}

func getQueryIntPtr(r *http.Request, queryName string, queryLookup map[string]QueryType) (*int32, error) {
	queryParam := queryLookup[queryName]

	integer, err := parseIntQuery(r, queryParam)
	if errors.Is(err, errEmptyQuery) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	integer32 := int32(integer)
	return &integer32, nil
}

// this is actually getQueryNameOrIdPtr
func getQueryIdPtr[T h.HasID, R any, A APIResource, L APIResourceList](r *http.Request, i handlerInput[T, R, A, L], queryName string, queryLookup map[string]QueryType) (*int32, error) {
	queryParam := queryLookup[queryName]

	id, err := parseNameOrIdQuery(r, queryParam, i.resourceType, i.objLookup)
	if errors.Is(err, errEmptyQuery) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func getQueryEnumPtr[E, N any](r *http.Request, queryName, endpoint string, et EnumType[E, N], queryLookup map[string]QueryType) (*string, error) {
	queryParam := queryLookup[queryName]

	enumRes, err := parseEnumQuery(r, endpoint, queryParam, et)
	if errors.Is(err, errEmptyQuery) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &enumRes.Name, nil
}
