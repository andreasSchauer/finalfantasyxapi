package api

import (
	"fmt"
	"slices"
	"strings"
)

func (cfg *Config) assignDefaultParams(endpoint EndpointName, sections map[SectionName]Subsection) map[QueryParamName]QueryParam {
	return cfg.completeQueryParamsInit([]QueryParam{}, endpoint, sections)
}

func (cfg *Config) completeQueryParamsInit(params []QueryParam, endpoint EndpointName, sections map[SectionName]Subsection) map[QueryParamName]QueryParam {
	params = slices.Concat(params, cfg.Q.defaultParamSlice)

	if hasSimpleView(sections) {
		queryParamIDs := QueryParam{
			Name:        QpnIDs,
			Description: "Used to manually input the ids of resources to be batch-fetched for simple display. The original order will be preserved, but duplicates will be removed.",
			Type:        qptIDList,
			IsExclusive: true,
			ParamUse:    puSimple,
			ForSegment:  getSnPtr(SnSimple),
		}
		params = append(params, queryParamIDs)
	}

	if hasFilters(cfg, params) {
		queryParamFlip := QueryParam{
			Name:        QpnFlip,
			Description: "Flips the filtered results in a list response and returns the negative.",
			Type:        qptBool,
			ParamUse:    puList,
		}
		params = append(params, queryParamFlip)
	}

	return paramSliceToMap(params, &endpoint)
}

func hasSimpleView(sections map[SectionName]Subsection) bool {
	if sections == nil {
		return false
	}

	_, ok := sections[SnSimple]
	return ok
}

func hasFilters(cfg *Config, params []QueryParam) bool {
	for _, param := range params {
		if param.ParamUse == puList && !hasParam(param, cfg.Q.defaultParamSlice) {
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
	case qptBool:
		q.Usage = s + "{bool}"
		q.ExampleUses = []string{s + "true", s + "false"}

	case qptEnum:
		e := q.EnumVals[0].Name
		q.Usage = s + "{value|id}"
		q.ExampleUses = []string{s + "1", s + e}

	case qptEnumList:
		e1 := q.EnumVals[0].Name
		e2 := q.EnumVals[1].Name
		q.Usage = s + "{value|id},..."
		q.ExampleUses = []string{s + "1,2", s + fmt.Sprintf("%s,%s", e1, e2)}

	case qptID:
		q.Usage = s + "{id}"
		q.ExampleUses = []string{s + "1"}

	case qptIDNul:
		q.Usage = s + "{id|'none'}"
		q.ExampleUses = []string{s + "1", s + "none"}

	case qptIDList:
		q.Usage = s + "{id},..."
		q.ExampleUses = []string{s + "1", s + "1,2"}

	case qptInt:
		q.Usage = s + "{int}"
		q.ExampleUses = []string{s + "1"}

	case qptIntList:
		q.Usage = s + "{int},...|{int}-{int}"
		q.ExampleUses = []string{s + "1", s + "1,2", s + "1-3", s + "1,2-4"}

	case qptNameID:
		e := q.ExampleVals[0]
		q.Usage = s + "{name|id}"
		q.ExampleUses = []string{s + "1", s + e}

	case qptNameIDList:
		e1 := q.ExampleVals[0]
		e2 := q.ExampleVals[1]
		q.Usage = s + "{name|id},..."
		q.ExampleUses = []string{s + "1", s + "1,2", s + fmt.Sprintf("%s,%s", e1, e2)}

	case qptNameIDListNul:
		e1 := q.ExampleVals[0]
		e2 := q.ExampleVals[1]
		q.Usage = s + "{name|id},...|{'none'}"
		q.ExampleUses = []string{s + "1", s + "1,2", s + fmt.Sprintf("%s,%s", e1, e2), s + "none"}

	case qptValue:
		e1 := string(q.AllowedValues[0])
		e2 := string(q.AllowedValues[1])
		q.Usage = s + "{val}"
		q.ExampleUses = []string{s + e1, s + e2}

	case qptValueList:
		e1 := string(q.AllowedValues[0])
		e2 := string(q.AllowedValues[1])
		q.Usage = s + "{val}"
		q.ExampleUses = []string{s + e1, s + fmt.Sprintf("%s,%s", e1, e2)}

	case qptStat:
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

func exampleUse(queryName QueryParamName, value string) string {
	return fmt.Sprintf("?%s=%s", queryName, value)
}
