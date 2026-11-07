package handlerinput

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/queryparam"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

// verifies the correct usage of a query parameter
func (q queryVerifyer[T]) usage() error {
	err := q.requestType()
	if err != nil {
		return err
	}

	err = q.context()
	if err != nil {
		return err
	}

	return nil
}

// used for normal endpoints like /endpoint and /endpoint/{id|key}. simply looks up the query param and returns an error, if it doesn't exist.
func (q queryVerifyer[T]) getParam(queryName string) (queryparam.QueryParam, error) {
	queryParam, ok := q.queryLookup[ts.QueryParamName(queryName)]
	if !ok {
		return queryparam.QueryParam{}, errs.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("parameter '%s' does not exist for endpoint /%s. use /api/%s/parameters for available parameters", queryName, q.endpoint, q.endpoint), nil)
	}

	return queryParam, nil
}

// used for alternative lists like /endpoint/sections and /endpoint/parameters. simply looks up the query param and returns an error, if it doesn't exist. enforces the used parameter to be a default parameter.
func (q queryVerifyer[T]) getParamAltList(queryName string, listName *string) (queryparam.QueryParam, error) {
	_, err := q.getParam(queryName)
	if err != nil {
		return queryparam.QueryParam{}, err
	}
	
	defaultParams := queryparam.GetDefaultParamMap()
	queryParam, ok := defaultParams[ts.QueryParamName(queryName)]
	if !ok {
		return queryparam.QueryParam{}, errs.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("only the following default parameters are allowed when using /api/%s/%s: %s", q.endpoint, *listName, queryparam.QueryMapToString(defaultParams)), nil)
	}

	return queryParam, nil
}
