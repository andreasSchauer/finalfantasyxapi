package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type filteredResList[T HasAPIResource] struct {
	resources []T
	err       error
}

func frl[T HasAPIResource](res []T, err error) filteredResList[T] {
	return filteredResList[T]{
		resources: res,
		err:       err,
	}
}

// query uses an id of another resource type to filter resources
func idOnlyQuery[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], inputRes []A, queryName string, maxID int, dbQuery func(context.Context, int32) ([]int32, error)) ([]A, error) {
	queryParam := i.queryLookup[queryName]

	id, err := parseIDOnlyQuery(r, queryParam, maxID)
	if errors.Is(err, errEmptyQuery) {
		return inputRes, nil
	}
	if err != nil {
		return nil, err
	}

	dbIDs, err := dbQuery(r.Context(), id)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't retrieve %ss for parameter '%s'.", i.resourceType, queryParam.Name), err)
	}

	resources := idsToAPIResources(cfg, i, dbIDs)

	return resources, nil
}

// query uses an id of another resource type to filter resources but uses more specialized logic in between (found in wrapperFn)
func idOnlyQueryWrapper[T h.HasID, R any, A APIResource, L APIResourceList](r *http.Request, i handlerInput[T, R, A, L], inputRes []A, queryName string, maxID int, wrapperFn func(*http.Request, int32) ([]A, error)) ([]A, error) {
	queryParam := i.queryLookup[queryName]

	id, err := parseIDOnlyQuery(r, queryParam, maxID)
	if errors.Is(err, errEmptyQuery) {
		return inputRes, nil
	}
	if err != nil {
		return nil, err
	}

	resources, err := wrapperFn(r, id)
	if err != nil {
		return nil, err
	}

	return resources, nil
}

// db query searches for resources with matching boolean db column value
func boolQuery[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], inputRes []A, queryName string, dbQuery func(context.Context, bool) ([]int32, error)) ([]A, error) {
	queryParam := i.queryLookup[queryName]

	b, err := parseBooleanQuery(r, queryParam)
	if errors.Is(err, errEmptyQuery) {
		return inputRes, nil
	}
	if err != nil {
		return nil, err
	}

	dbIDs, err := dbQuery(r.Context(), b)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't retrieve %ss for parameter '%s'.", i.resourceType, queryParam.Name), err)
	}

	resources := idsToAPIResources(cfg, i, dbIDs)

	return resources, nil
}

// db query accumulates all resources that fulfill a certain condition (mostly if it has resources of a specific type). a false boolean flips these results
func boolQuery2[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], inputRes []A, queryName string, dbQuery func(context.Context) ([]int32, error)) ([]A, error) {
	queryParam := i.queryLookup[queryName]

	b, err := parseBooleanQuery(r, queryParam)
	if errors.Is(err, errEmptyQuery) {
		return inputRes, nil
	}
	if err != nil {
		return nil, err
	}

	dbIDs, err := dbQuery(r.Context())
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't retrieve %ss for parameter '%s'.", i.resourceType, queryParam.Name), err)
	}

	resources := idsToAPIResources(cfg, i, dbIDs)

	if !b {
		resources = removeResources(inputRes, resources)
	}

	return resources, nil
}

// query uses an enum type (id or string possible) that needs to be checked for validity and then returns all resources matching that type
func typeQuery[T h.HasID, R any, A APIResource, L APIResourceList, ET any](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], et EnumType[ET], inputRes []A, queryName string, dbQuery func(context.Context, ET) ([]int32, error)) ([]A, error) {
	queryParam := i.queryLookup[queryName]
	enum, err := parseTypeQuery(r, queryParam, et)
	if errors.Is(err, errEmptyQuery) {
		return inputRes, nil
	}
	if err != nil {
		return nil, err
	}

	modeType := et.convFunc(enum.Name)

	dbIDs, err := dbQuery(r.Context(), modeType)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't retrieve %ss for parameter '%s'.", i.resourceType, queryParam.Name), err)
	}

	resources := idsToAPIResources(cfg, i, dbIDs)
	return resources, nil
}

// used for method queries for example as a combination of all of them (see areas 'item' parameter)
func combineFilteredAPIResources[A APIResource](filteredLists []filteredResList[A]) ([]A, error) {
	resources := []A{}

	for _, filtered := range filteredLists {
		if filtered.err != nil {
			return nil, filtered.err
		}
		
		resources = combineResources(resources, filtered.resources)
	}

	return resources, nil
}
