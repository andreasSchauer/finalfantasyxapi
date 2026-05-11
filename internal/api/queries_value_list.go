package api

import (
	"fmt"
	"net/http"
	"slices"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func filterByValues[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], query string, queryParam QueryParam, dbQueryMap map[string]DbQueryNoInput) ([]int32, error) {
	values, err := queryListSplit(cfg, query)
	if err != nil {
		return nil, err
	}
	valueMap := make(map[string]bool)

	filteredLists := []filteredIdList{}

	for _, value := range values {
		dbQuery, ok := dbQueryMap[value]
		if !ok {
			return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid value '%s' used for parameter '%s'. allowed values: %s.", query, queryParam.Name, h.FormatStringSlice(queryParam.AllowedValues)), nil)
		}

		dbIDs, err := dbQuery(r.Context())
		if err != nil {
			return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by value '%s'.", i.resourceType, value), err)
		}
		filteredList := fidl(dbIDs, err)
		filteredLists = append(filteredLists, filteredList)

		valueMap[value] = true
	}

	return filterIdSlices(filteredLists)
}


func filterByValues2[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], query string, queryParam QueryParam, dbQuery DbQueryValueList) ([]int32, error) {
	values, err := queryListSplit(cfg, query)
	if err != nil {
		return nil, err
	}

	for _, value := range values {
		if !slices.Contains(queryParam.AllowedValues, value) {
			return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid value '%s' used for parameter '%s'. allowed values: %s.", query, queryParam.Name, h.FormatStringSlice(queryParam.AllowedValues)), nil)
		}
	}

	dbIDs, err := dbQuery(r.Context(), values)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by value list '%s'.", i.resourceType, h.FormatStringSlice(values)), err)
	}

	return dbIDs, nil
}