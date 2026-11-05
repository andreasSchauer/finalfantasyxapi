package api

import (
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func AssignDefaultParams(endpoint ts.EndpointName, eLookup, qLookup *map[ts.QueryParamName]QueryParam) {
	params := paramSliceToMap(Q.DefaultParamSlice, endpoint)
	*eLookup = params
	*qLookup = params
}

func AssignQueryParams[T seedtypes.Lookupable, R any, A apiresource.APIResource](params []QueryParam, i *handlerInput[T, R, A], queryLookup *map[ts.QueryParamName]QueryParam) {
	paramMap := completeQueryParams(params, *i)

	i.queryLookup = paramMap
	*queryLookup = paramMap
}

func completeQueryParams[T seedtypes.Lookupable, R any, A apiresource.APIResource](params []QueryParam, i handlerInput[T, R, A]) map[ts.QueryParamName]QueryParam {
	params = slices.Concat(params, Q.DefaultParamSlice)

	if hasSimpleView(i.subsections) {
		queryParamIDs := QueryParam{
			Name:        ts.QpnIDs,
			Description: "Used to manually input the ids of resources to be batch-fetched for simple display. The original order will be preserved, but duplicates will be removed.",
			Type:        ts.QptIDList,
			IsExclusive: true,
			ParamUse:    puSimple,
			ForSegment:  ts.GetSnPtr(ts.SnSimple),
		}
		params = append(params, queryParamIDs)
	}

	if hasFilters(params) {
		queryParamFlip := QueryParam{
			Name:        ts.QpnFlip,
			Description: "Flips the filtered results in a list response and returns the negative.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		}
		params = append(params, queryParamFlip)
	}

	return paramSliceToMap(params, i.Endpoint)
}

func hasSimpleView(sections map[ts.SectionName]Subsection) bool {
	if sections == nil {
		return false
	}

	_, ok := sections[ts.SnSimple]
	return ok
}

func hasFilters(params []QueryParam) bool {
	for _, param := range params {
		if param.ParamUse == puList && !hasParam(param, Q.DefaultParamSlice) {
			return true
		}
	}

	return false
}

func hasParam(target QueryParam, params []QueryParam) bool {
	if params == nil {
		return false
	}

	for _, param := range params {
		if param.Name == target.Name {
			return true
		}
	}

	return false
}
