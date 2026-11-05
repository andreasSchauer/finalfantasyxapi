package api

import (
	"net/http"

	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

type QueryParameterList struct {
	ListParams
	Results []QueryParam `json:"results"`
}

func (l QueryParameterList) GetListParams() ListParams {
	return l.ListParams
}

func getQueryParamList(r *http.Request, endpoint ts.EndpointName, queryLookup map[ts.QueryParamName]QueryParam) (QueryParameterList, error) {

	queryParams := paramMapToSlice(queryLookup)
	queryParams = getAllowedResources(endpoint, queryParams)
	queryParams = getAllowedValuesFromTypes(queryParams)

	listParams, shownResources, err := createPaginatedList(r, queryParams)
	if err != nil {
		return QueryParameterList{}, err
	}

	list := QueryParameterList{
		ListParams: listParams,
		Results:    createQueryParamRefResURLs(shownResources),
	}

	return list, nil
}

func createQueryParamRefResURLs(params []QueryParam) []QueryParam {
	paramsNew := params

	for i := range paramsNew {
		param := paramsNew[i]
		if param.ReferencesInt == nil {
			continue
		}
		param.References = make([]string, len(param.ReferencesInt))

		for j := range param.ReferencesInt {
			if param.Type == ts.QptEnum || param.Type == ts.QptEnumList {
				ref := param.ReferencesEnumsInt[j]
				param.References[j] = CreateEnumURL(ref)
				continue
			}

			ref := param.ReferencesInt[j]
			param.References[j] = CreateListURL(ref)
		}
		paramsNew[i] = param
	}

	return paramsNew
}

func getAllowedResources(endpoint ts.EndpointName, params []QueryParam) []QueryParam {
	for idx, param := range params {
		for _, id := range param.AllowedIDs {
			allowedRes := CreateResourceURL(endpoint, id)
			param.AllowedResources = append(param.AllowedResources, allowedRes)
		}
		params[idx] = param
	}

	return params
}

func getAllowedValuesFromTypes(params []QueryParam) []QueryParam {
	for idx, param := range params {
		for _, typeRes := range param.EnumVals {
			param.AllowedValues = append(param.AllowedValues, ts.QueryValue(typeRes.Name))
		}
		params[idx] = param
	}

	return params
}
