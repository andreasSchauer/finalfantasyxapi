package api

import (
	"fmt"
	"net/http"
)

// verifies the correct usage of all query parameters of an endpoint
func verifyQueryParams[T any](r *http.Request, endpoint EndpointName, queryLookup map[QueryParamName]QueryParam, key *T, segment *string) error {
	q := r.URL.Query()

	for query := range q {
		queryParam, err := getParamEndpoint(endpoint, queryLookup, query)
		if err != nil {
			return err
		}

		err = verifyQueryUsage(q, queryParam, endpoint, queryLookup, key, segment)
		if err != nil {
			return err
		}
	}

	return nil
}

func verifyQueryParamsServiceGet(r *http.Request, endpoint EndpointName, queryLookup map[QueryParamName]QueryParam) error {
	_, err := checkEmptyQuery(r, queryLookup[qpnState])
	if errExceptEmptyQuery(err) {
		return err
	}
	if queryIsEmpty(err) {
		return newHTTPError(http.StatusBadRequest, fmt.Sprintf("parameter '%s' can't be empty", qpnState), nil)
	}

	err = verifyQueryParams[any](r, endpoint, queryLookup, nil, nil)
	if err != nil {
		return err
	}

	return nil
}

// verifies, if no query params are in use for an endpoint that receives a post request
func verifyQueryParamsServicePost(r *http.Request) error {
	if len(r.URL.Query()) > 0 {
		return newHTTPError(http.StatusBadRequest, "POST requests don't have any query parameters", nil)
	}

	return nil
}

// verifies the correct usage of all query parameters of an alternative list that is used on an endpoint, like /endpoint/sections, or /endpoint/parameters
func verifyQueryParamsAltList(cfg *Config, r *http.Request, endpoint EndpointName, queryLookup map[QueryParamName]QueryParam, listName *string) error {
	q := r.URL.Query()

	for query := range q {
		queryParam, err := getParamAltList(cfg, endpoint, queryLookup, query, listName)
		if err != nil {
			return err
		}

		err = verifyQueryUsage[any](q, queryParam, endpoint, cfg.q.defaultParams, nil, listName)
		if err != nil {
			return err
		}
	}

	return nil
}
