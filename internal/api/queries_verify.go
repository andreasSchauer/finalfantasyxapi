package api

import (
	"net/http"
	"net/url"

	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

type queryVerifyer[T any] struct {
	usedNames   url.Values
	param       QueryParam
	endpoint    ts.EndpointName
	queryLookup map[ts.QueryParamName]QueryParam
	segment     *string
	key         *T
}

// verifies the correct usage of all query parameters of an endpoint
func verifyQueryParams[T any](r *http.Request, endpoint ts.EndpointName, queryLookup map[ts.QueryParamName]QueryParam, key *T, segment *string) error {
	var err error
	q := r.URL.Query()

	for queryName := range q {
		qv := queryVerifyer[T]{
			usedNames:   q,
			endpoint:    endpoint,
			queryLookup: queryLookup,
			segment:     segment,
			key:         key,
		}

		qv.param, err = qv.getParam(queryName)
		if err != nil {
			return err
		}

		err = qv.usage()
		if err != nil {
			return err
		}
	}

	return nil
}

// verifies, if no query params are in use for an endpoint that receives a post request
func verifyQueryParamsService(r *http.Request) error {
	if len(r.URL.Query()) > 0 {
		return newHTTPError(http.StatusBadRequest, "computational endpoints don't have any query parameters", nil)
	}

	return nil
}

// verifies the correct usage of all query parameters of an alternative list that is used on an endpoint, like /endpoint/sections, or /endpoint/parameters
func verifyQueryParamsAltList(r *http.Request, endpoint ts.EndpointName, queryLookup map[ts.QueryParamName]QueryParam, listName *string) error {
	var err error
	q := r.URL.Query()

	for queryName := range q {
		qv := queryVerifyer[any]{
			usedNames:   q,
			endpoint:    endpoint,
			queryLookup: queryLookup,
			segment:     listName,
		}

		qv.param, err = qv.getParamAltList(queryName, listName)
		if err != nil {
			return err
		}

		err = qv.usage()
		if err != nil {
			return err
		}
	}

	return nil
}
