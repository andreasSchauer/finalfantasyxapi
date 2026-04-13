package api

import (
	"cmp"
	"net/http"
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

func queryListSplit(cfg *Config, query string) ([]string, error) {
	segments := querySplit(query, ",")

	if len(segments) > cfg.fetchLimit {
		return nil, newHTTPErrorFetchLimit(cfg.fetchLimit)
	}

	return segments, nil
}

func queryIntMapToSlice(m map[string]int32) []int32 {
	items := []int32{}

	for _, item := range m {
		items = append(items, item)
	}

	slices.SortStableFunc(items, func(a, b int32) int {
		return cmp.Compare(a, b)
	})

	return items
}

func getRawQueryLimit(cfg *Config, r *http.Request) (int, error) {
	queryParamLimit := cfg.q.defaultParams["limit"]
	query, err := checkEmptyQuery(r, queryParamLimit)
	if err != nil {
		return 0, err
	}

	limit, err := checkQueryInt(queryParamLimit, query)
	if err != nil {
		return 0, err
	}
	if limit == 0 {
		limit = *queryParamLimit.DefaultVal
	}
	if limit > cfg.fetchLimit {
		return cfg.fetchLimit, nil
	}

	return limit, nil
}
