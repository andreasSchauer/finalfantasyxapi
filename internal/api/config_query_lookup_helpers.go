package api

import (
	"fmt"
	"slices"
	"strings"
)


func (cfg *Config) assignDefaultParams() map[QueryParamName]QueryParam {
	return cfg.completeQueryParamsInit([]QueryParam{}, false)
}

func (cfg *Config) completeQueryParamsInit(params []QueryParam, hasSimpleView bool) map[QueryParamName]QueryParam {
	params = slices.Concat(params, cfg.q.defaultParamSlice)

	if hasSimpleView {
		queryParamIDs := QueryParam{
			Name:        qpnIDs,
			Description: "Used to manually input the ids of resources to be batch-fetched for simple display. The original order will be preserved, but duplicates will be removed.",
			Type:        qptIdList,
			IsExclusive: true,
			ForList:     false,
			ForSingle:   false,
			ForSegment:  getSnPtr(snSimple),
		}
		params = append(params, queryParamIDs)
	}

	if hasFilters(cfg, params) {
		queryParamFlip := QueryParam{
			Name:        qpnFlip,
			Description: "Flips the filtered results in a list response and returns the negative.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		}
		params = append(params, queryParamFlip)
	}

	return querySliceToMap(cfg, params)
}

func hasSimpleView(sections map[SectionName]Subsection) bool {
	if sections == nil {
		return false
	}

	_, ok := sections[snSimple]
	return ok
}

func hasFilters(cfg *Config, params []QueryParam) bool {
	for _, param := range params {
		if param.ForList && !hasParam(param, cfg.q.defaultParamSlice) {
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


func (cfg *Config) assignParamUsage(p QueryParam) QueryParam {
	s := fmt.Sprintf("?%s=", p.Name)

	switch p.Type {
	case qptBool:
		p.Usage = s + "{bool}"
		p.ExampleUses = []string{s + "true", s + "false"}

	case qptEnum:
		enums := createEnumValSlice(p.EnumLookup)
		e := enums[0].Name
		p.Usage = s + "{value|id}"
		p.ExampleUses = []string{s + "1", s + e}

	case qptEnumList:
		enums := createEnumValSlice(p.EnumLookup)
		e1 := enums[0].Name
		e2 := enums[1].Name
		p.Usage = s + "{value|id},..."
		p.ExampleUses = []string{s + "1,2", s + fmt.Sprintf("%s,%s", e1, e2)}

	case qptId:
		p.Usage = s + "{id}"
		p.ExampleUses = []string{s + "1"}

	case qptIdNul:
		p.Usage = s + "{id|'none'}"
		p.ExampleUses = []string{s + "1", s + "none"}

	case qptIdList:
		p.Usage = s + "{id},..."
		p.ExampleUses = []string{s + "1", s + "1,2"}

	case qptInt:
		p.Usage = s + "{int}"
		p.ExampleUses = []string{s + "1"}

	case qptIntList:
		p.Usage = s + "{int},...|{int}-{int}"
		p.ExampleUses = []string{s + "1", s + "1,2", s + "1-3", s + "1,2-4"}

	case qptNameId:
		e := p.ExampleVals[0]
		p.Usage = s + "{name|id}"
		p.ExampleUses = []string{s + "1", s + e}

	case qptNameIdList:
		e1 := p.ExampleVals[0]
		e2 := p.ExampleVals[1]
		p.Usage = s + "{name|id},..."
		p.ExampleUses = []string{s + "1", s + "1,2", s + fmt.Sprintf("%s,%s", e1, e2)}

	case qptNameIdListNul:
		e1 := p.ExampleVals[0]
		e2 := p.ExampleVals[1]
		p.Usage = s + "{name|id},...|{'none'}"
		p.ExampleUses = []string{s + "1", s + "1,2", s + fmt.Sprintf("%s,%s", e1, e2), s + "none"}

	case qptValue:
		e1 := string(p.AllowedValues[0])
		e2 := string(p.AllowedValues[1])
		p.Usage = s + "{val}"
		p.ExampleUses = []string{s + e1, s + e2}

	case qptValueList:
		e1 := string(p.AllowedValues[0])
		e2 := string(p.AllowedValues[1])
		p.Usage = s + "{val}"
		p.ExampleUses = []string{s + e1, s + fmt.Sprintf("%s,%s", e1, e2)}

	case qptStat:
		p.Usage = s + "{stat}={int},..."

	default:
		return p
	}

	if p.SpecialInputs != nil {
		for _, input := range p.SpecialInputs {
			usageTrimmed := strings.TrimSuffix(p.Usage, "}")
			p.Usage = fmt.Sprintf("%s|'%s'}", usageTrimmed, input.Key)
			p.ExampleUses = append(p.ExampleUses, s+string(input.Key))
		}
	}

	return p
}