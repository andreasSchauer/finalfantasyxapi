package api

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/queryparam"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

type QueryParameterList struct {
	apiresource.ListParams
	Results []queryparam.QueryParam `json:"results"`
}

func (l QueryParameterList) GetListParams() apiresource.ListParams {
	return l.ListParams
}

func getQueryParamList(r *http.Request, endpoint ts.EndpointName, queryLookup map[ts.QueryParamName]queryparam.QueryParam) (QueryParameterList, error) {

	queryParams := queryparam.ParamMapToSlice(queryLookup)
	queryParams = getAllowedResources(endpoint, queryParams)
	queryParams = getAllowedValuesFromTypes(queryParams)

	listParams, shownResources, err := handlerinput.CreatePaginatedList(r, queryParams)
	if err != nil {
		return QueryParameterList{}, err
	}

	list := QueryParameterList{
		ListParams: listParams,
		Results:    createQueryParamRefResURLs(shownResources),
	}

	return list, nil
}

func createQueryParamRefResURLs(params []queryparam.QueryParam) []queryparam.QueryParam {
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
				param.References[j] = ah.CreateEnumURL(ref)
				continue
			}

			ref := param.ReferencesInt[j]
			param.References[j] = ah.CreateListURL(ref)
		}
		paramsNew[i] = param
	}

	return paramsNew
}

func getAllowedResources(endpoint ts.EndpointName, params []queryparam.QueryParam) []queryparam.QueryParam {
	for idx, param := range params {
		for _, id := range param.AllowedIDs {
			allowedRes := ah.CreateResourceURL(endpoint, id)
			param.AllowedResources = append(param.AllowedResources, allowedRes)
		}
		params[idx] = param
	}

	return params
}

func getAllowedValuesFromTypes(params []queryparam.QueryParam) []queryparam.QueryParam {
	for idx, param := range params {
		for _, typeRes := range param.EnumVals {
			param.AllowedValues = append(param.AllowedValues, ts.QueryValue(typeRes.Name))
		}
		params[idx] = param
	}

	return params
}
