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

// not the biggest fan of these function names
// will split into multiple files once other resource types get added

// In each function I can put the general logic into its own function.
// Then I make a wrapper that converts into the needed resource like I did with the retrieval helpers

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
