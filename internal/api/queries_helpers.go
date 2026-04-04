package api

import (
	"cmp"
	"net/http"
	"slices"
	"strings"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)


func checkEmptyQuery(r *http.Request, queryParam QueryType) (string, error) {
	query := r.URL.Query().Get(queryParam.Name)
	if query == "" {
		return "", errEmptyQuery
	}

	return strings.ToLower(query), nil
}


func checkNoneQuery(query string) error {
	if query == "none" {
		return errQueryNone
	}

	return nil
}


func queryMapToSlice(lookup map[string]QueryType) []QueryType {
	queryParams := []QueryType{}

	for key := range lookup {
		queryParams = append(queryParams, lookup[key])
	}

	slices.SortStableFunc(queryParams, func(a, b QueryType) int {
		return cmp.Compare(a.ID, b.ID)
	})

	return queryParams
}


func queryMapToString(lookup map[string]QueryType) string {
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