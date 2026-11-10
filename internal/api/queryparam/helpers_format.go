package queryparam

import (
	"cmp"
	"slices"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
)

func ParamMapToSlice(lookup map[ts.QueryParamName]QueryParam) []QueryParam {
	queryParams := []QueryParam{}

	for key := range lookup {
		queryParams = append(queryParams, lookup[key])
	}

	slices.SortStableFunc(queryParams, func(a, b QueryParam) int {
		return cmp.Compare(a.ID, b.ID)
	})

	return queryParams
}

func ParamSliceToMap(params []QueryParam, endpoint ts.EndpointName) map[ts.QueryParamName]QueryParam {
	paramMap := make(map[ts.QueryParamName]QueryParam, len(params))

	for i, param := range params {
		param.ID = i + 1
		param.AssignUsage()
		param.Endpoint = endpoint
		paramMap[param.Name] = param
	}

	return paramMap
}

func GetDefaultParamMap() map[ts.QueryParamName]QueryParam {
	return DefaultParamSliceToMap(GetDefaultParams())
}

func DefaultParamSliceToMap(params []QueryParam) map[ts.QueryParamName]QueryParam {
	paramMap := make(map[ts.QueryParamName]QueryParam, len(params))

	for i, param := range params {
		param.ID = i + 1
		param.AssignUsage()
		paramMap[param.Name] = param
	}

	return paramMap
}

func DefaultParamSliceToString(params []QueryParam) string {
	paramMap := DefaultParamSliceToMap(params)
	return QueryMapToString(paramMap)
}

func QueryMapToString(lookup map[ts.QueryParamName]QueryParam) string {
	params := ParamMapToSlice(lookup)
	names := []string{}

	for _, param := range params {
		names = append(names, string(param.Name))
	}

	return h.FormatStringSlice(names)
}
