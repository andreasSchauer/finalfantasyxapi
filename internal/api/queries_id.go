package api

import (
	"errors"
	"fmt"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

// query uses an id of another resource type to filter resources
func idQuery[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], inputRes []A, queryName string, maxID int, dbQuery DbQueryIntMany) ([]A, error) {
	queryParam := i.queryLookup[queryName]

	id, err := parseIdQuery(r, queryParam, maxID)
	if errors.Is(err, errEmptyQuery) {
		return inputRes, nil
	}
	if err != nil {
		return nil, err
	}

	dbIDs, err := dbQuery(r.Context(), id)
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resourceType, queryParam, err)
	}

	resources := idsToAPIResources(cfg, i, dbIDs)

	return resources, nil
}

// like idOnlyQuery, but with more specialized logic in between (wrapperFn)
func idQueryWrapper[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], inputRes []A, queryName string, maxID int, wrapperFn func(*Config, *http.Request, int32) ([]A, error)) ([]A, error) {
	queryParam := i.queryLookup[queryName]

	id, err := parseIdQuery(r, queryParam, maxID)
	if errors.Is(err, errEmptyQuery) {
		return inputRes, nil
	}
	if err != nil {
		return nil, err
	}

	resources, err := wrapperFn(cfg, r, id)
	if err != nil {
		return nil, err
	}

	return resources, nil
}

func idQueryNul[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], inputRes []A, queryName string, maxID int, dbQuery DbQueryNullIntMany) ([]A, error) {
	queryParam := i.queryLookup[queryName]

	idPtr, err := parseIdQueryNul(r, queryParam, maxID)
	if errors.Is(err, errEmptyQuery) {
		return inputRes, nil
	}
	if err != nil {
		return nil, err
	}

	dbIds, err := dbQuery(r.Context(), h.GetNullInt32(idPtr))
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resourceType, queryParam, err)
	}

	resources := idsToAPIResources(cfg, i, dbIds)
	return resources, nil
}

// can be returned by id wrapperFns that use a second parameter for a set of allowed values
func filterByIdAndValues[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], id int32, queryName, pResType string, dbQueryMap map[string]DbQueryIntMany) ([]A, error) {
	queryParam := i.queryLookup[queryName]
	query, err := checkEmptyQuery(r, queryParam)
	if err != nil {
		return dbQueriesToApiResources(cfg, r, i, id, pResType, dbQueryMap)
	}

	values := querySplit(query, ",")
	valueMap := make(map[string]bool)

	filteredLists := []filteredResList[A]{}

	for _, value := range values {
		dbQuery, ok := dbQueryMap[value]
		if !ok {
			return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid value '%s' used for parameter '%s'. allowed values: %s.", query, queryParam.Name, h.FormatStringSlice(queryParam.AllowedValues)), nil)
		}

		if valueMap[value] {
			return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("duplicate use of value '%s' for parameter '%s'. each value can only be used once.", value, queryParam.Name), nil)
		}

		filteredLists = append(filteredLists, frl(getResourcesDbID(cfg, r, i, id, pResType, dbQuery)))

		valueMap[value] = true
	}

	return combineFilteredAPIResources(filteredLists)
}
