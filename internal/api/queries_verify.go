package api

import "net/http"

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