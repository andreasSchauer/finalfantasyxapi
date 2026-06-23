package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

// can make those return sql.NullTypes, if the pointers have no other use

func getQueryBoolPtr(r *http.Request, queryName QueryParamName, queryLookup map[QueryParamName]QueryParam) (*bool, error) {
	queryParam := queryLookup[queryName]

	b, err := parseBooleanQuery(r, queryParam)
	if err != nil {
		return nil, err
	}

	return &b, nil
}

func getQueryIntPtr(r *http.Request, queryName QueryParamName, queryLookup map[QueryParamName]QueryParam) (*int32, error) {
	queryParam := queryLookup[queryName]

	integer, err := parseIntQuery(r, queryParam)
	if err != nil {
		return nil, err
	}

	integer32 := int32(integer)
	return &integer32, nil
}

func getQueryValuePtr(r *http.Request, queryName QueryParamName, queryLookup map[QueryParamName]QueryParam) (*string, error) {
	queryParam := queryLookup[queryName]

	value, err := parseValueQuery(r, queryParam)
	if err != nil {
		return nil, err
	}

	return &value, nil
}

func getQueryIdPtr[T seeding.Lookupable, R any, A APIResource, L APIResourceList](r *http.Request, i handlerInput[T, R, A, L], queryName QueryParamName, queryLookup map[QueryParamName]QueryParam) (*int32, error) {
	queryParam := queryLookup[queryName]

	id, err := parseIdQuery(r, queryParam, i.objLookup)
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func getQueryNameIdPtr[T seeding.Lookupable, R any, A APIResource, L APIResourceList](r *http.Request, i handlerInput[T, R, A, L], queryName QueryParamName, queryLookup map[QueryParamName]QueryParam) (*int32, error) {
	queryParam := queryLookup[queryName]

	id, err := parseNameIdQuery(r, queryParam, i.resTypeSing, i.objLookup)
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func getQueryEnumPtr[E, N any](r *http.Request, queryName QueryParamName, endpoint EndpointName, et EnumType[E, N], queryLookup map[QueryParamName]QueryParam) (*string, error) {
	queryParam := queryLookup[queryName]

	enumRes, err := parseEnumQuery(r, endpoint, queryParam, et)
	if err != nil {
		return nil, err
	}

	return &enumRes.Name, nil
}
