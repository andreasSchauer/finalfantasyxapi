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

func getQueryParamList(srv *Server, r *http.Request, endpoint ts.EndpointName, queryLookup map[ts.QueryParamName]QueryParam) (QueryParameterList, error) {

	queryParams := paramMapToSlice(queryLookup)
	queryParams = getAllowedResources(srv, endpoint, queryParams)
	queryParams = getAllowedValuesFromTypes(queryParams)

	listParams, shownResources, err := createPaginatedList(srv, r, queryParams)
	if err != nil {
		return QueryParameterList{}, err
	}

	list := QueryParameterList{
		ListParams: listParams,
		Results:    createQueryParamRefResURLs(srv, shownResources),
	}

	return list, nil
}

func createQueryParamRefResURLs(srv *Server, params []QueryParam) []QueryParam {
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
				param.References[j] = createEnumURL(srv, ref)
				continue
			}

			ref := param.ReferencesInt[j]
			param.References[j] = createListURL(srv, ref)
		}
		paramsNew[i] = param
	}

	return paramsNew
}

func getAllowedResources(srv *Server, endpoint ts.EndpointName, params []QueryParam) []QueryParam {
	for idx, param := range params {
		for _, id := range param.AllowedIDs {
			allowedRes := CreateResourceURL(srv, endpoint, id)
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
