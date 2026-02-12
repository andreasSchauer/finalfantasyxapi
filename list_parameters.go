package main

import (
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type QueryParameterList struct {
	ListParams
	Results []QueryType `json:"results"`
}

func (l QueryParameterList) getListParams() ListParams {
	return l.ListParams
}

func getQueryParamList[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L]) (QueryParameterList, error) {
	queryParams := queryMapToSlice(i.queryLookup)
	queryParams = getAllowedResources(cfg, i, queryParams)
	queryParams = getAllowedValuesFromTypes(cfg, queryParams)

	listParams, shownResources, err := createPaginatedList(cfg, r, queryParams)
	if err != nil {
		return QueryParameterList{}, err
	}

	list := QueryParameterList{
		ListParams: listParams,
		Results:    shownResources,
	}

	return list, nil
}

func getAllowedResources[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, i handlerInput[T, R, A, L], params []QueryType) []QueryType {
	for idx, param := range params {
		for _, id := range param.AllowedIDs {
			allowedRes := createResourceURL(cfg, i.endpoint, id)
			param.AllowedResources = append(param.AllowedResources, allowedRes)
		}
		params[idx] = param
	}

	return params
}

func getAllowedValuesFromTypes(cfg *Config, params []QueryType) []QueryType {
	for idx, param := range params {
		if param.TypeLookup == nil {
			continue
		}

		types := createTypeResourceSlice(cfg, "", param.TypeLookup)

		for _, typeRes := range types {
			param.AllowedValues = append(param.AllowedValues, typeRes.Name)
		}
		params[idx] = param
	}

	return params
}