package api

import (
	"fmt"
	"slices"
	"strings"

	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (i handlerInput[T, R, A]) assignDefaultParams() map[ts.QueryParamName]QueryParam {
	return i.completeQueryParamsInit([]QueryParam{})
}

func (i handlerInputEndpoints) assignDefaultParams() map[ts.QueryParamName]QueryParam {
	return paramSliceToMap(Q.defaultParamSlice, &i.Endpoint)
}

func (i handlerInputEnums) assignDefaultParams() map[ts.QueryParamName]QueryParam {
	return paramSliceToMap(Q.defaultParamSlice, &i.Endpoint)
}

// make a wrapper method around this for each handlerInput?
func (i handlerInput[T, R, A]) completeQueryParamsInit(params []QueryParam) map[ts.QueryParamName]QueryParam {
	params = slices.Concat(params, Q.defaultParamSlice)

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

	return paramSliceToMap(params, &i.Endpoint)
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
		if param.ParamUse == puList && !hasParam(param, Q.defaultParamSlice) {
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

func (q QueryParam) assignParamUsage() QueryParam {
	s := fmt.Sprintf("?%s=", q.Name)

	switch q.Type {
	case ts.QptBool:
		q.Usage = s + "{bool}"
		q.ExampleUses = []string{s + "true", s + "false"}

	case ts.QptEnum:
		e := q.EnumVals[0].Name
		q.Usage = s + "{value|id}"
		q.ExampleUses = []string{s + "1", s + e}

	case ts.QptEnumList:
		e1 := q.EnumVals[0].Name
		e2 := q.EnumVals[1].Name
		q.Usage = s + "{value|id},..."
		q.ExampleUses = []string{s + "1,2", s + fmt.Sprintf("%s,%s", e1, e2)}

	case ts.QptID:
		q.Usage = s + "{id}"
		q.ExampleUses = []string{s + "1"}

	case ts.QptIDNul:
		q.Usage = s + "{id|'none'}"
		q.ExampleUses = []string{s + "1", s + "none"}

	case ts.QptIDList:
		q.Usage = s + "{id},..."
		q.ExampleUses = []string{s + "1", s + "1,2"}

	case ts.QptInt:
		q.Usage = s + "{int}"
		q.ExampleUses = []string{s + "1"}

	case ts.QptIntList:
		q.Usage = s + "{int},...|{int}-{int}"
		q.ExampleUses = []string{s + "1", s + "1,2", s + "1-3", s + "1,2-4"}

	case ts.QptNameID:
		e := q.ExampleVals[0]
		q.Usage = s + "{name|id}"
		q.ExampleUses = []string{s + "1", s + e}

	case ts.QptNameIDList:
		e1 := q.ExampleVals[0]
		e2 := q.ExampleVals[1]
		q.Usage = s + "{name|id},..."
		q.ExampleUses = []string{s + "1", s + "1,2", s + fmt.Sprintf("%s,%s", e1, e2)}

	case ts.QptNameIDListNul:
		e1 := q.ExampleVals[0]
		e2 := q.ExampleVals[1]
		q.Usage = s + "{name|id},...|{'none'}"
		q.ExampleUses = []string{s + "1", s + "1,2", s + fmt.Sprintf("%s,%s", e1, e2), s + "none"}

	case ts.QptValue:
		e1 := string(q.AllowedValues[0])
		e2 := string(q.AllowedValues[1])
		q.Usage = s + "{val}"
		q.ExampleUses = []string{s + e1, s + e2}

	case ts.QptValueList:
		e1 := string(q.AllowedValues[0])
		e2 := string(q.AllowedValues[1])
		q.Usage = s + "{val}"
		q.ExampleUses = []string{s + e1, s + fmt.Sprintf("%s,%s", e1, e2)}

	case ts.QptStat:
		q.Usage = s + "{stat}:{int},..."

	default:
		return q
	}

	if q.SpecialInputs != nil {
		for _, input := range q.SpecialInputs {
			usageTrimmed := strings.TrimSuffix(q.Usage, "}")
			q.Usage = fmt.Sprintf("%s|'%s'}", usageTrimmed, input.Key)
			q.ExampleUses = append(q.ExampleUses, s+string(input.Key))
		}
	}

	return q
}

func exampleUse(queryName ts.QueryParamName, value string) string {
	return fmt.Sprintf("?%s=%s", queryName, value)
}
