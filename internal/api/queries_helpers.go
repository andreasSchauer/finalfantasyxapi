package api

import (
	"cmp"
	"net/http"
	"net/url"
	"slices"
	"strconv"
	"strings"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func paramMapToSlice(lookup map[typestrings.QueryParamName]QueryParam) []QueryParam {
	queryParams := []QueryParam{}

	for key := range lookup {
		queryParams = append(queryParams, lookup[key])
	}

	slices.SortStableFunc(queryParams, func(a, b QueryParam) int {
		return cmp.Compare(a.ID, b.ID)
	})

	return queryParams
}

func paramSliceToMap(params []QueryParam, endpoint *typestrings.EndpointName) map[typestrings.QueryParamName]QueryParam {
	paramMap := make(map[typestrings.QueryParamName]QueryParam, len(params))

	for i, param := range params {
		param.ID = i + 1

		param = param.assignParamUsage()

		if endpoint != nil {
			param.Endpoint = *endpoint
		}

		paramMap[param.Name] = param
	}

	return paramMap
}

func queryMapToString(lookup map[typestrings.QueryParamName]QueryParam) string {
	params := paramMapToSlice(lookup)
	names := []string{}

	for _, param := range params {
		names = append(names, string(param.Name))
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
	queryParamLimit := cfg.Q.defaultParams[typestrings.QpnLimit]
	query, err := queryParamLimit.checkEmpty(r)
	if err != nil {
		return 0, err
	}

	limit, err := queryParamLimit.checkInt(query)
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

func getLimitMax(cfg *Config) string {
	queryParamLimit := cfg.Q.defaultParams[typestrings.QpnLimit]

	for _, input := range queryParamLimit.SpecialInputs {
		if input.Key == typestrings.QsvMax {
			return strconv.Itoa(input.Val)
		}
	}

	return ""
}

func setLimitMax(cfg *Config, r *http.Request, q url.Values) {
	limit := q.Get(string(typestrings.QpnLimit))
	if limit != "" {
		return
	}

	q.Set(string(typestrings.QpnLimit), getLimitMax(cfg))
	r.URL.RawQuery = q.Encode()
}
