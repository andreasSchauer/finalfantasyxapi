package api

import (
	"cmp"
	"slices"
	"strings"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)



func queryMapToSlice(lookup map[string]QueryParam) []QueryParam {
	queryParams := []QueryParam{}

	for key := range lookup {
		queryParams = append(queryParams, lookup[key])
	}

	slices.SortStableFunc(queryParams, func(a, b QueryParam) int {
		return cmp.Compare(a.ID, b.ID)
	})

	return queryParams
}

func querySliceToMap(cfg *Config, params []QueryParam) map[string]QueryParam {
	paramMap := make(map[string]QueryParam)

	for i, param := range params {
		param.ID = i + 1

		param = cfg.assignParamUsage(param)
		paramMap[param.Name] = param
	}

	return paramMap
}

func queryMapToString(lookup map[string]QueryParam) string {
	params := queryMapToSlice(lookup)
	names := []string{}

	for _, param := range params {
		names = append(names, param.Name)
	}

	return h.FormatStringSlice(names)
}

func querySplit(query, sep string) []string {
	queryTrimmed := strings.TrimSuffix(query, sep)
	return strings.Split(queryTrimmed, sep)
}
