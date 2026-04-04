package api

import (
	"errors"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

// query uses an enum type (id or string possible) that needs to be checked for validity and then returns all resources matching that type
func enumQuery[T h.HasID, R any, A APIResource, L APIResourceList, E, N any](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], et EnumType[E, N], inputRes []A, queryName string, dbQuery DbQueryEnum[E]) ([]A, error) {
	queryParam := i.queryLookup[queryName]
	enum, err := parseEnumQuery(r, i.endpoint, queryParam, et)
	if errors.Is(err, errEmptyQuery) {
		return inputRes, nil
	}
	if err != nil {
		return nil, err
	}

	typedStr := et.convFunc(enum.Name)

	dbIDs, err := dbQuery(r.Context(), typedStr)
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resourceType, queryParam, err)
	}

	resources := idsToAPIResources(cfg, i, dbIDs)
	return resources, nil
}


// like enum query, but with more specialized logic in between (wrapperFn). For example, if types are grouped together (ctbIconType)
func enumQueryWrapper[T h.HasID, R any, A APIResource, L APIResourceList, E, N any](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], et EnumType[E, N], inputRes []A, queryName string, wrapperFn func(*Config, *http.Request, E) ([]int32, error)) ([]A, error) {
	queryParam := i.queryLookup[queryName]
	enum, err := parseEnumQuery(r, i.endpoint, queryParam, et)
	if errors.Is(err, errEmptyQuery) {
		return inputRes, nil
	}
	if err != nil {
		return nil, err
	}

	typedStr := et.convFunc(enum.Name)
	
	dbIDs, err := wrapperFn(cfg, r, typedStr)
	if err != nil {
		return nil, err
	}

	resources := idsToAPIResources(cfg, i, dbIDs)
	return resources, nil
}