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

func getQueryParamList(server *Server, r *http.Request, endpoint ts.EndpointName, queryLookup map[ts.QueryParamName]QueryParam) (QueryParameterList, error) {

	queryParams := paramMapToSlice(queryLookup)
	queryParams = getAllowedResources(server, endpoint, queryParams)
	queryParams = getAllowedValuesFromTypes(queryParams)

	listParams, shownResources, err := createPaginatedList(server, r, queryParams)
	if err != nil {
		return QueryParameterList{}, err
	}

	list := QueryParameterList{
		ListParams: listParams,
		Results:    createQueryParamRefResURLs(server, shownResources),
	}

	return list, nil
}

func createQueryParamRefResURLs(server *Server, params []QueryParam) []QueryParam {
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
				param.References[j] = createEnumURL(server, ref)
				continue
			}

			ref := param.ReferencesInt[j]
			param.References[j] = createListURL(server, ref)
		}
		paramsNew[i] = param
	}

	return paramsNew
}

func getAllowedResources(server *Server, endpoint ts.EndpointName, params []QueryParam) []QueryParam {
	for idx, param := range params {
		for _, id := range param.AllowedIDs {
			allowedRes := CreateResourceURL(server, endpoint, id)
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
