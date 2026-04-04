package api

import (
	"cmp"
	"net/http"
	"slices"
	"strings"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

// checks, if a queryParam is empty and returns errEmptyQuery, if it is
func checkEmptyQuery(r *http.Request, queryParam QueryParam) (string, error) {
	query := r.URL.Query().Get(queryParam.Name)
	if query == "" {
		return "", errEmptyQuery
	}

	return strings.ToLower(query), nil
}

// checks, if "none" was used as input and returns errQueryNone, if it was
func checkNoneQuery(query string) error {
	if query == "none" {
		return errQueryNone
	}

	return nil
}

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
