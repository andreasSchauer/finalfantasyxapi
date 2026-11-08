package api

import (
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/queryparam"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func assignDefaultParams(endpoint ts.EndpointName, eLookup, qLookup *map[ts.QueryParamName]queryparam.QueryParam) {
	paramMap := queryparam.ParamSliceToMap(Q.DefaultParamSlice, endpoint)
	*eLookup = paramMap
	*qLookup = paramMap
}

func assignQueryParams[T seedtypes.Lookupable, R any, A apiresource.APIResource](params []queryparam.QueryParam, i *handlerinput.Main[T, R, A], queryLookup *map[ts.QueryParamName]queryparam.QueryParam) {
	paramMap := completeQueryParams(params, *i)

	i.QueryLookup = paramMap
	*queryLookup = paramMap
}

func completeQueryParams[T seedtypes.Lookupable, R any, A apiresource.APIResource](params []queryparam.QueryParam, i handlerinput.Main[T, R, A]) map[ts.QueryParamName]queryparam.QueryParam {
	params = slices.Concat(params, Q.DefaultParamSlice)

	if hasSimpleView(i.Subsections) {
		queryParamIDs := queryparam.QueryParam{
			Name:        ts.QpnIDs,
			Description: "Used to manually input the ids of resources to be batch-fetched for simple display. The original order will be preserved, but duplicates will be removed.",
			Type:        ts.QptIDList,
			IsExclusive: true,
			ParamUse:    queryparam.PuSimple,
			ForSegment:  ts.GetSnPtr(ts.SnSimple),
		}
		params = append(params, queryParamIDs)
	}

	if hasFilters(params) {
		queryParamFlip := queryparam.QueryParam{
			Name:        ts.QpnFlip,
			Description: "Flips the filtered results in a list response and returns the negative.",
			Type:        ts.QptBool,
			ParamUse:    queryparam.PuList,
		}
		params = append(params, queryParamFlip)
	}

	return queryparam.ParamSliceToMap(params, i.Endpoint)
}

func hasSimpleView(sections map[ts.SectionName]handlerinput.Subsection) bool {
	if sections == nil {
		return false
	}

	_, ok := sections[ts.SnSimple]
	return ok
}

func hasFilters(params []queryparam.QueryParam) bool {
	for _, param := range params {
		if param.ParamUse == queryparam.PuList && !hasParam(param, Q.DefaultParamSlice) {
			return true
		}
	}

	return false
}

func hasParam(target queryparam.QueryParam, params []queryparam.QueryParam) bool {
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
