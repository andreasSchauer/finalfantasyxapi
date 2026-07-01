package api

import (
	"net/http"
)

type QueryParameterList struct {
	ListParams
	Results []QueryParam `json:"results"`
}

func (l QueryParameterList) getListParams() ListParams {
	return l.ListParams
}

func getQueryParamList(cfg *Config, r *http.Request, endpoint EndpointName, queryLookup map[QueryParamName]QueryParam) (QueryParameterList, error) {

	queryParams := queryMapToSlice(queryLookup)
	queryParams = getAllowedResources(cfg, endpoint, queryParams)
	queryParams = getAllowedValuesFromTypes(queryParams)

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

func createQueryParamRefResURLs(cfg *Config, params []QueryParam) []QueryParam {
	paramsNew := params

	for i := range paramsNew {
		param := paramsNew[i]
		if param.ReferencesInt == nil {
			continue
		}
		param.References = make([]string, len(param.ReferencesInt))

		for j := range param.ReferencesInt {
			if param.Type == qptEnum || param.Type == qptEnumList {
				ref := param.ReferencesEnumsInt[j]
				param.References[j] = createEnumURL(cfg, ref)
				continue
			}
			
			ref := param.ReferencesInt[j]
			param.References[j] = createListURL(cfg, ref)
		}
		paramsNew[i] = param
	}

	return paramsNew
}

func getAllowedResources(cfg *Config, endpoint EndpointName, params []QueryParam) []QueryParam {
	for idx, param := range params {
		for _, id := range param.AllowedIDs {
			allowedRes := createResourceURL(cfg, endpoint, id)
			param.AllowedResources = append(param.AllowedResources, allowedRes)
		}
		params[idx] = param
	}

	return params
}

func getAllowedValuesFromTypes(params []QueryParam) []QueryParam {
	for idx, param := range params {
		if param.EnumLookup == nil {
			continue
		}

		types := createEnumValSlice(param.EnumLookup)

		for _, typeRes := range types {
			param.AllowedValues = append(param.AllowedValues, QueryValue(typeRes.Name))
		}
		params[idx] = param
	}

	return params
}
