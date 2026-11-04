package api

import (
	"fmt"
	"net/http"

	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

// verifies, if the query param is used in the correct context (the relationships with the parameters it is combined with are correct (exclusive, required, forbidden, usable with))
func (q queryVerifyer[T]) context() error {
	err := q.exclusive()
	if err != nil {
		return err
	}

	err = q.requiresAll()
	if err != nil {
		return err
	}

	err = q.conflictsWith()
	if err != nil {
		return err
	}

	err = q.requiresOne()
	if err != nil {
		return err
	}

	return nil
}

// verifies the use of an exclusive query param
func (q queryVerifyer[T]) exclusive() error {
	if q.param.IsExclusive && !q.canUseExclusiveParam() {
		return newHTTPError(http.StatusBadRequest, fmt.Sprintf("parameter '%s' can't be combined with other parameters", q.param.Name), nil)
	}

	return nil
}

// checks, if only one exclusive query param is used. returns false, if the query param doesn't exist, or if an exclusive query param is combined with any other query param (len > 1).
func (q queryVerifyer[T]) canUseExclusiveParam() bool {
	for query := range q.usedNames {
		queryParam, ok := q.queryLookup[ts.QueryParamName(query)]
		if !ok {
			return false
		}

		if queryParam.IsExclusive && len(q.usedNames) > 1 {
			return false
		}
	}

	return true
}

// checks, if all requiresAll query parameters of the requested query param are present. returns an error, if at least one requiresAll param is missing.
func (q queryVerifyer[T]) requiresAll() error {
	if q.param.RequiresAll == nil {
		return nil
	}

	for _, reqParam := range q.param.RequiresAll {
		reqParamVal := q.usedNames.Get(string(reqParam))

		if reqParamVal == "" {
			return newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid usage of parameter '%s'. when using parameter '%s', the following parameter(s) must be present: %s", q.param.Name, q.param.Name, ts.FormatQpnSlice(q.param.RequiresAll)), nil)
		}
	}

	return nil
}

// checks, if a conflicting query parameter of the requested query param is present. returns an error, if at least one conflicting param is present.
func (q queryVerifyer[T]) conflictsWith() error {
	if q.param.ConflictsWith == nil {
		return nil
	}

	for _, conflictParam := range q.param.ConflictsWith {
		fbParamVal := q.usedNames.Get(string(conflictParam))

		if fbParamVal != "" {
			return newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid usage of parameter '%s'. parameter '%s' can't be used in combination with the following parameter(s): %s", q.param.Name, q.param.Name, ts.FormatQpnSlice(q.param.ConflictsWith)), nil)
		}
	}

	return nil
}

// checks, if at least one of the query params, that the requested query param must be combined with, is present. returns an error, if none of them are present.
func (q queryVerifyer[T]) requiresOne() error {
	if q.param.RequiresOne == nil {
		return nil
	}

	for _, reqParam := range q.param.RequiresOne {
		reqParamVal := q.usedNames.Get(string(reqParam))

		if reqParamVal != "" {
			return nil
		}
	}

	return newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid usage of parameter '%s'. parameter '%s' can only be used in combination with at least one of the following parameters: %s", q.param.Name, q.param.Name, ts.FormatQpnSlice(q.param.RequiresOne)), nil)
}

// checks, if the query param is a default param.
func isDefaultParam(cfg *Config, queryName ts.QueryParamName) bool {
	_, ok := cfg.Q.defaultParams[queryName]
	return ok
}
