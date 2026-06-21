package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type QueryParameterList struct {
	ListParams
	Results []QueryParam `json:"results"`
}

func (l QueryParameterList) getListParams() ListParams {
	return l.ListParams
}

func getQueryParamList[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L]) (QueryParameterList, error) {

	queryParams := queryMapToSlice(i.queryLookup)
	queryParams = getAllowedResources(cfg, i, queryParams)
	queryParams = getAllowedValuesFromTypes(cfg, queryParams)

	listParams, shownResources, err := createPaginatedList(cfg, r, queryParams)
	if err != nil {
		return QueryParameterList{}, err
	}

	list := QueryParameterList{
		ListParams: listParams,
		Results:    createQueryParamRefResURLs(cfg, shownResources),
	}

	return list, nil
}

func createQueryParamRefResURLs(cfg *Config, params []QueryParam) ([]QueryParam) {
	paramsNew := params

	for i := range paramsNew {
		param := paramsNew[i]
		if param.References == nil {
			continue
		}

		for j := range param.References {
			ref := param.References[j]
			param.References[j] = createListURL(cfg, ref)
		}
		paramsNew[i] = param
	}

	return paramsNew
}

func getAllowedResources[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, i handlerInput[T, R, A, L], params []QueryParam) []QueryParam {
	for idx, param := range params {
		for _, id := range param.AllowedIDs {
			allowedRes := createResourceURL(cfg, i.endpoint, id)
			param.AllowedResources = append(param.AllowedResources, allowedRes)
		}
		params[idx] = param
	}

	return params
}

func getAllowedValuesFromTypes(cfg *Config, params []QueryParam) []QueryParam {
	for idx, param := range params {
		if param.TypeLookup == nil {
			continue
		}

		types := createEnumResourceSlice(cfg, "", param.TypeLookup)

		for _, typeRes := range types {
			param.AllowedValues = append(param.AllowedValues, typeRes.Name)
		}
		params[idx] = param
	}

	return params
}
