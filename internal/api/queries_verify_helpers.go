package api

import (
	"fmt"
	"net/http"
	"net/url"
)


// verifies that either 'state' param is used alone, or that required params are not empty
func verifyRequiredParams(r *http.Request, q url.Values, queryLookup map[QueryParamName]QueryParam) error {
	var hasState bool
	paramsAmt := len(q)

	_, err := checkEmptyQuery(r, queryLookup[qpnState])
	if errExceptEmptyQuery(err) {
		return err
	}
	if !queryIsEmpty(err) {
		hasState = true
	}

	if paramsAmt == 1 && hasState {
		return nil
	}
	if paramsAmt > 1 && hasState {
		return newHTTPError(http.StatusBadRequest, fmt.Sprintf("parameter '%s' can't be coupled with other parameters", string(qpnState)), nil)
	}
	
	for _, param := range queryLookup {
		if !param.IsRequired {
			continue
		}
		
		_, err := checkEmptyQuery(r, param)
		if errExceptEmptyQuery(err) {
			return err
		}
		if queryIsEmpty(err) && !hasState {
			return newHTTPError(http.StatusBadRequest, fmt.Sprintf("parameter '%s' can't be empty.", param.Name), nil)
		}
	}

	return nil
}


// verifies the correct usage of a query parameter
func verifyQueryUsage[T any](q url.Values, queryParam QueryParam, endpoint EndpointName, queryLookup map[QueryParamName]QueryParam, key *T, segment *string) error {
	err := vpFormat(queryParam, endpoint, key, segment)
	if err != nil {
		return err
	}

	err = vpContext(q, queryParam, queryLookup)
	if err != nil {
		return err
	}

	return nil
}

// used for normal endpoints like /endpoint and /endpoint/{id|key}. simply looks up the query param and returns an error, if it doesn't exist.
func getParamEndpoint(endpoint EndpointName, queryLookup map[QueryParamName]QueryParam, query string) (QueryParam, error) {
	queryParam, ok := queryLookup[QueryParamName(query)]
	if !ok {
		return QueryParam{}, newHTTPError(http.StatusBadRequest, fmt.Sprintf("parameter '%s' does not exist for endpoint /%s. use /api/%s/parameters for available parameters.", query, endpoint, endpoint), nil)
	}

	return queryParam, nil
}


// used for alternative lists like /endpoint/sections and /endpoint/parameters. simply looks up the query param and returns an error, if it doesn't exist. enforces the used parameter to be a default parameter.
func getParamAltList(cfg *Config, endpoint EndpointName, queryLookup map[QueryParamName]QueryParam, query string, listName *string) (QueryParam, error) {
	_, err := getParamEndpoint(endpoint, queryLookup, query)
	if err != nil {
		return QueryParam{}, err
	}

	queryParam, ok := cfg.q.defaultParams[QueryParamName(query)]
	if !ok {
		return QueryParam{}, newHTTPError(http.StatusBadRequest, fmt.Sprintf("only the following default parameters are allowed when using /api/%s/%s: %s.", endpoint, *listName, queryMapToString(cfg.q.defaultParams)), nil)
	}

	return queryParam, nil
}