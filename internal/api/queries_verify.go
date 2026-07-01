package api

import (
	"net/http"
	"net/url"
)

// verifies the correct usage of all query parameters of an alternative list that is used on an endpoint that expects ids as its primary key for single resources, like /endpoint/sections, or /endpoint/parameters
func verifyQueryParamsAltListID(cfg *Config, r *http.Request, endpoint EndpointName, listName *string) error {
	q := r.URL.Query()

	for query := range q {
		queryParam, err := getParamAltList(cfg, endpoint, query, listName)
		if err != nil {
			return err
		}

		err = verifyQueryUsageID(q, queryParam, endpoint, cfg.q.defaultParams, nil, listName)
		if err != nil {
			return err
		}
	}

	return nil
}

// verifies the correct usage of all query parameters of an endpoint that uses ids as its primary key for single resources, like /endpoint/{id}.
func verifyQueryParamsID(r *http.Request, endpoint EndpointName, queryLookup map[QueryParamName]QueryParam, id *int32, segment *string) error {
	q := r.URL.Query()

	for query := range q {
		queryParam, err := getParamEndpoint(endpoint, queryLookup, query)
		if err != nil {
			return err
		}

		err = verifyQueryUsageID(q, queryParam, endpoint, queryLookup, id, segment)
		if err != nil {
			return err
		}
	}

	return nil
}

// checks the usage of an endpoint that uses ids as its primary key for single resources, like /locations/{id}.
func verifyQueryUsageID(q url.Values, queryParam QueryParam, endpoint EndpointName, queryLookup map[QueryParamName]QueryParam, id *int32, segment *string) error {
	err := verifyExclusiveParam(q, queryParam, queryLookup)
	if err != nil {
		return err
	}

	err = verifySegmentOnlyParam(queryParam, segment, endpoint)
	if err != nil {
		return err
	}

	err = verifySingleResourceParamID(queryParam, id)
	if err != nil {
		return err
	}

	err = verifyListResourceParamID(queryParam, id)
	if err != nil {
		return err
	}

	err = verifyRequiredParams(q, queryParam)
	if err != nil {
		return err
	}

	err = verifyForbiddenParams(q, queryParam)
	if err != nil {
		return err
	}

	err = verifyUsableWith(q, queryParam)
	if err != nil {
		return err
	}

	return nil
}

// verifies the correct usage of all query parameters of an alternative list that is used on an endpoint that expects keys as its primary key for single resources, like /enums/parameters
func verifyQueryParamsAltListKey(cfg *Config, r *http.Request, endpoint EndpointName, listName *string) error {
	q := r.URL.Query()

	for query := range q {
		queryParam, err := getParamAltList(cfg, endpoint, query, listName)
		if err != nil {
			return err
		}

		err = verifyQueryUsageKey(q, queryParam, endpoint, cfg.q.defaultParams, nil, listName)
		if err != nil {
			return err
		}
	}

	return nil
}

// verifies the correct usage of all query parameters of an endpoint that uses keys as its primary key for single resources, like /enums/{EnumName}.
func verifyQueryParamsKey(r *http.Request, endpoint EndpointName, queryLookup map[QueryParamName]QueryParam, key *string) error {
	q := r.URL.Query()

	for query := range q {
		queryParam, err := getParamEndpoint(endpoint, queryLookup, query)
		if err != nil {
			return err
		}

		err = verifyQueryUsageKey(q, queryParam, endpoint, queryLookup, key, nil)
		if err != nil {
			return err
		}
	}

	return nil
}

// checks the usage of an endpoint that uses keys as its primary key for single resources, like /enums/{EnumName}.
func verifyQueryUsageKey(q url.Values, queryParam QueryParam, endpoint EndpointName, queryLookup map[QueryParamName]QueryParam, key, segment *string) error {
	err := verifyExclusiveParam(q, queryParam, queryLookup)
	if err != nil {
		return err
	}

	err = verifySegmentOnlyParam(queryParam, segment, endpoint)
	if err != nil {
		return err
	}

	err = verifySingleResourceParamKey(queryParam, key)
	if err != nil {
		return err
	}

	err = verifyListResourceParamKey(queryParam, key)
	if err != nil {
		return err
	}

	err = verifyRequiredParams(q, queryParam)
	if err != nil {
		return err
	}

	err = verifyForbiddenParams(q, queryParam)
	if err != nil {
		return err
	}

	err = verifyUsableWith(q, queryParam)
	if err != nil {
		return err
	}

	return nil
}
