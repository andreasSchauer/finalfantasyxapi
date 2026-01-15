package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	//h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type filteredResList[T APIResource] struct {
	resources []T
	err       error
}

func frl[T APIResource](res []T, err error) filteredResList[T] {
	return filteredResList[T]{
		resources: res,
		err:       err,
	}
}

// not the biggest fan of these function names
// will split into multiple files once other resource types get added

func idQueryLocBased(cfg *Config, r *http.Request, i handlerInput[seeding.Area, Area, LocationApiResourceList], inputAreas []LocationAPIResource, queryName string, maxID int, dbQuery func(context.Context, int32) ([]int32, error)) ([]LocationAPIResource, error) {
	queryParam := i.queryLookup[queryName]

	id, err := parseIDOnlyQuery(r, queryParam, maxID)
	if errors.Is(err, errEmptyQuery) {
		return inputAreas, nil
	}
	if err != nil {
		return nil, err
	}

	dbIDs, err := dbQuery(r.Context(), id)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't retrieve %ss for parameter '%s'.", i.resourceType, queryParam.Name), err)
	}

	resources := idsToLocationAPIResources(cfg, i, dbIDs)

	return resources, nil
}

func idQueryWrapperLocBased(r *http.Request, i handlerInput[seeding.Area, Area, LocationApiResourceList], inputAreas []LocationAPIResource, queryName string, maxID int, dbQuery func(*http.Request, int32) ([]LocationAPIResource, error)) ([]LocationAPIResource, error) {
	queryParam := i.queryLookup[queryName]

	id, err := parseIDOnlyQuery(r, queryParam, maxID)
	if errors.Is(err, errEmptyQuery) {
		return inputAreas, nil
	}
	if err != nil {
		return nil, err
	}

	resources, err := dbQuery(r, id)
	if err != nil {
		return nil, err
	}

	return resources, nil
}

func boolQueryLocBased(cfg *Config, r *http.Request, i handlerInput[seeding.Area, Area, LocationApiResourceList], inputAreas []LocationAPIResource, queryName string, dbQuery func(context.Context, bool) ([]int32, error)) ([]LocationAPIResource, error) {
	queryParam := i.queryLookup[queryName]

	b, err := parseBooleanQuery(r, queryParam)
	if errors.Is(err, errEmptyQuery) {
		return inputAreas, nil
	}
	if err != nil {
		return nil, err
	}

	dbIDs, err := dbQuery(r.Context(), b)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't retrieve %ss for parameter '%s'.", i.resourceType, queryParam.Name), err)
	}

	resources := idsToLocationAPIResources(cfg, i, dbIDs)

	return resources, nil
}

func boolAccumulatorLocBased(cfg *Config, r *http.Request, i handlerInput[seeding.Area, Area, LocationApiResourceList], inputAreas []LocationAPIResource, queryName string, dbQuery func(context.Context) ([]int32, error)) ([]LocationAPIResource, error) {
	queryParam := i.queryLookup[queryName]

	b, err := parseBooleanQuery(r, queryParam)
	if errors.Is(err, errEmptyQuery) {
		return inputAreas, nil
	}
	if err != nil {
		return nil, err
	}

	dbIDs, err := dbQuery(r.Context())
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't retrieve %ss for parameter '%s'.", i.resourceType, queryParam.Name), err)
	}

	resources := idsToLocationAPIResources(cfg, i, dbIDs)

	if !b {
		resources = removeResources(inputAreas, resources)
	}

	return resources, nil
}
