package api

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

// used for query filters that can't really be generalized. this one simply checks, if it's empty and then calls the wrapperFn
func basicQueryWrapper[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], inputIDs []int32, queryName QueryParamName, wrapperFn QueryWrapBasic) IdFilter {
	return func (ctx context.Context) ([]int32, error) {
		queryParam := i.queryLookup[queryName]
		query, err := basicQueryChecks(r, queryParam, i.queryLookup)
		if err != nil {
			return inputIDs, nil
		}

		dbIDs, err := wrapperFn(cfg, r, ctx, query, queryParam)
		if errors.Is(err, errQueryRedirect) {
			return inputIDs, nil
		}
		if err != nil {
			return nil, err
		}

		return dbIDs, nil
	}
}

func basicQueryChecks(r *http.Request, queryParam QueryParam, queryLookup map[QueryParamName]QueryParam) (string, error) {
	query, err := checkEmptyQuery(r, queryParam)
	if err != nil {
		return "", err
	}

	if replParamsPresent(r, queryParam, queryLookup) {
		return "", errQueryRedirect
	}

	return query, nil
}

// checks, if a queryParam is empty and returns errEmptyQuery, if it is
func checkEmptyQuery(r *http.Request, queryParam QueryParam) (string, error) {
	query := r.URL.Query().Get(string(queryParam.Name))
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

func replParamsPresent(r *http.Request, queryParam QueryParam, queryLookup map[QueryParamName]QueryParam) bool {
	for _, param := range queryParam.ReplacedBy {
		p := queryLookup[param]
		_, err := checkEmptyQuery(r, p)
		if !queryIsEmpty(err) {
			return true
		}
	}
	return false
}
