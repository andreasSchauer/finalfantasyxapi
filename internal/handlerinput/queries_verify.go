package handlerinput

import (
	"net/http"
	"net/url"

	"github.com/andreasSchauer/finalfantasyxapi/internal/queryparam"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

type queryVerifyer[T any] struct {
	usedNames   url.Values
	param       queryparam.QueryParam
	endpoint    ts.EndpointName
	queryLookup map[ts.QueryParamName]queryparam.QueryParam
	segment     *string
	key         *T
}

// verifyQueryParams verifies the correct usage of all query parameters of an endpoint
func verifyQueryParams[T any](r *http.Request, endpoint ts.EndpointName, queryLookup map[ts.QueryParamName]queryparam.QueryParam, key *T, segment *string) error {
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

// verifyQueryParamsAltList verifies the correct usage of all query parameters of an alternative list that is used on an endpoint, like /endpoint/sections, or /endpoint/parameters
func verifyQueryParamsAltList(r *http.Request, endpoint ts.EndpointName, queryLookup map[ts.QueryParamName]queryparam.QueryParam, listName *string) error {
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
