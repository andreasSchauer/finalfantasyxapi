package api

import (
	"fmt"
	"net/http"
	"net/url"
)

// verifies, if the query param is used in the correct context (the relationships with the parameters it is combined with are correct (exclusive, required, forbidden, usable with))
func vpContext(q url.Values, queryParam QueryParam, queryLookup map[QueryParamName]QueryParam) error {
	err := vpExclusive(q, queryParam, queryLookup)
	if err != nil {
		return err
	}

	err = vpRequired(q, queryParam)
	if err != nil {
		return err
	}

	err = vpConflictsWith(q, queryParam)
	if err != nil {
		return err
	}

	err = vpUsableWith(q, queryParam)
	if err != nil {
		return err
	}

	return nil
}

// checks, if the query param is a default param.
func isDefaultParam(cfg *Config, queryName QueryParamName) bool {
	_, ok := cfg.q.defaultParams[queryName]
	return ok
}

// verifies the use of an exclusive query param
func vpExclusive(q url.Values, queryParam QueryParam, queryLookup map[QueryParamName]QueryParam) error {
	if queryParam.IsExclusive && !canUseExclusiveParam(q, queryLookup) {
		return newHTTPError(http.StatusBadRequest, fmt.Sprintf("parameter '%s' can't be combined with other parameters.", queryParam.Name), nil)
	}

	return nil
}

// checks, if only one exclusive query param is used. returns false, if the query param doesn't exist, or if an exclusive query param is combined with any other query param (len > 1).
func canUseExclusiveParam(q url.Values, lookup map[QueryParamName]QueryParam) bool {
	for query := range q {
		queryParam, ok := lookup[QueryParamName(query)]
		if !ok {
			return false
		}

		if queryParam.IsExclusive && len(q) > 1 {
			return false
		}
	}

	return true
}

// checks, if all required query parameters of the requested query param are present. returns an error, if at least one required param is missing.
func vpRequired(q url.Values, queryParam QueryParam) error {
	if queryParam.RequiredParams == nil {
		return nil
	}

	for _, reqParam := range queryParam.RequiredParams {
		reqParamVal := q.Get(string(reqParam))

		if reqParamVal == "" {
			return newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid usage of parameter '%s'. when using parameter '%s', the following parameter(s) must be present: %s.", queryParam.Name, queryParam.Name, formatQpnSlice(queryParam.RequiredParams)), nil)
		}
	}

	return nil
}

// checks, if a conflicting query parameter of the requested query param is present. returns an error, if at least one conflicting param is present.
func vpConflictsWith(q url.Values, queryParam QueryParam) error {
	if queryParam.ConflictsWith == nil {
		return nil
	}

	for _, conflictParam := range queryParam.ConflictsWith {
		fbParamVal := q.Get(string(conflictParam))

		if fbParamVal != "" {
			return newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid usage of parameter '%s'. parameter '%s' can't be used in combination with the following parameter(s): %s.", queryParam.Name, queryParam.Name, formatQpnSlice(queryParam.ConflictsWith)), nil)
		}
	}

	return nil
}

// checks, if at least one of the query params, that the requested query param must be combined with, is present. returns an error, if none of them are present.
func vpUsableWith(q url.Values, queryParam QueryParam) error {
	if queryParam.UsableWith == nil {
		return nil
	}

	for _, reqParam := range queryParam.UsableWith {
		reqParamVal := q.Get(string(reqParam))

		if reqParamVal != "" {
			return nil
		}
	}

	return newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid usage of parameter '%s'. parameter '%s' can only be used in combination with at least one of the following parameters: %s.", queryParam.Name, queryParam.Name, formatQpnSlice(queryParam.UsableWith)), nil)
}
