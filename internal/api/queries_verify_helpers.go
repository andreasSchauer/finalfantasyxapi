package api

import (
	"fmt"
	"net/http"
	"net/url"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

// used for alternative lists like /endpoint/sections and /endpoint/parameters. simply looks up the query param and returns an error, if it doesn't exist.
func getParamAltList(cfg *Config, endpoint EndpointName, query string, listName *string) (QueryParam, error) {
	queryParam, ok := cfg.q.defaultParams[QueryParamName(query)]
	if !ok {
		return QueryParam{}, newHTTPError(http.StatusBadRequest, fmt.Sprintf("only the following default parameters are allowed when using /api/%s/%s: %s.", endpoint, *listName, queryMapToString(cfg.q.defaultParams)), nil)
	}

	return queryParam, nil
}

// used for normal endpoints like /endpoint and /endpint/{id|key}. simply looks up the query param and returns an error, if it doesn't exist.
func getParamEndpoint(endpoint EndpointName, queryLookup map[QueryParamName]QueryParam, query string) (QueryParam, error) {
	queryParam, ok := queryLookup[QueryParamName(query)]
	if !ok {
		return QueryParam{}, newHTTPError(http.StatusBadRequest, fmt.Sprintf("parameter '%s' does not exist for endpoint /%s. use /api/%s/parameters for available parameters.", query, endpoint, endpoint), nil)
	}

	return queryParam, nil
}

// verifies the use of an exclusive query param
func verifyExclusiveParam(q url.Values, queryParam QueryParam, queryLookup map[QueryParamName]QueryParam) error {
	if queryParam.IsExclusive && !canUseExclusiveParam(q, queryLookup) {
		return newHTTPError(http.StatusBadRequest, fmt.Sprintf("parameter '%s' can't be combined with other parameters.", queryParam.Name), nil)
	}

	return nil
}

// checks, if only one exclusive query param is used. returns false, if the query param doesn't exist, or if more than one exclusive query param is in use.
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

// checks, if the query param is a default param.
func isDefaultParam(cfg *Config, queryName QueryParamName) bool {
	_, ok := cfg.q.defaultParams[queryName]
	return ok
}

// checks, if a query param that is meant for single resource requests is used in the correct context. returns an error, if no id is provided (meaning the parameter was combined with a list request), or if the given id is not among the allowed ids, if that restriction exists for the query param. meant to be used for resource endpoints like /locations/{id}.
func verifySingleResourceParamID(queryParam QueryParam, id *int32) error {
	if queryParam.ForSingle {
		if id == nil {
			return errSingleResParam(queryParam.Name)
		}

		err := verifyAllowedIDs(queryParam, *id)
		if err != nil {
			return err
		}
	}

	return nil
}

// checks, if the requested query param is used on an id that the query param expects.
func verifyAllowedIDs(queryParam QueryParam, id int32) error {
	if queryParam.AllowedIDs == nil {
		return nil
	}

	allowedIDPresent := false

	for _, reqID := range queryParam.AllowedIDs {
		if id == reqID {
			allowedIDPresent = true
		}
	}
	if !allowedIDPresent {
		return newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid id '%d'. parameter '%s' can only be used with ids: %s.", id, queryParam.Name, h.FormatIntSlice(queryParam.AllowedIDs)), nil)
	}

	return nil
}

// checks, if a query param that is meant for list requests is used in the correct context. returns an error, if an id is provided (meaning the parameter was combined with a single resource request). meant to be used for resource endpoints like /locations.
func verifyListResourceParamID(queryParam QueryParam, id *int32) error {
	if queryParam.ForList && id != nil {
		return errListResParam(queryParam.Name)
	}
	return nil
}

// checks, if a query param that is meant for single resource requests is used in the correct context. this function is used for endpoints that don't expect ids, but a different key, like EnumNames. returns an error, if no key is provided (meaning the parameter was combined with a list request). meant to be used for special endpoints like /enums/{enum_name}.
func verifySingleResourceParamKey(queryParam QueryParam, key *string) error {
	if queryParam.ForSingle {
		if key == nil {
			return errSingleResParam(queryParam.Name)
		}
	}
	return nil
}

// checks, if a query param that is meant for list requests is used in the correct context. this function is used for endpoints that don't expect ids, but a different key, like EnumNames. returns an error, if a key is provided (meaning the parameter was combined with a list request). meant to be used for special endpoints like /enums.
func verifyListResourceParamKey(queryParam QueryParam, key *string) error {
	if queryParam.ForList && key != nil {
		return errListResParam(queryParam.Name)
	}
	return nil
}

// checks, if a query param that is meant for a specific segment is used in the correct context, for example /endpoint/simple. returns an error, if the given segment doesn't match with the requested segment.
func verifySegmentOnlyParam(queryParam QueryParam, segment *string, endpoint EndpointName) error {
	if queryParam.ForSegment != nil && !segmentsMatch(queryParam.ForSegment, segment) {
		return newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid usage of parameter '%s'. parameter '%s' can only be used in the following format: '/api/%s/%s%s'.", queryParam.Name, queryParam.Name, endpoint, *queryParam.ForSegment, queryParam.Usage), nil)
	}

	return nil
}

// checks if a requested segment matches with the intended segment. returns false, if the segments don't match.
func segmentsMatch(sParam *SectionName, sRequest *string) bool {
	switch {
	case sParam == nil && sRequest == nil:
		return true

	case sParam != nil && sRequest != nil:
		segment := *sParam
		return string(segment) == *sRequest

	default:
		return false
	}
}

// checks, if all required query parameters of the requested query param are present. returns an error, if at least one required param is missing.
func verifyRequiredParams(q url.Values, queryParam QueryParam) error {
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

// checks, if a forbidden query parameter of the requested query param is present. returns an error, if at least one forbidden param is present.
func verifyForbiddenParams(q url.Values, queryParam QueryParam) error {
	if queryParam.ForbiddenParams == nil {
		return nil
	}

	for _, frbParam := range queryParam.ForbiddenParams {
		frbParamVal := q.Get(string(frbParam))

		if frbParamVal != "" {
			return newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid usage of parameter '%s'. parameter '%s' can't be used in combination with the following parameter(s): %s.", queryParam.Name, queryParam.Name, formatQpnSlice(queryParam.ForbiddenParams)), nil)
		}
	}

	return nil
}

// checks, if at least one of the query params that the requested query param must be combined with is present. returns an error, if none of them are present.
func verifyUsableWith(q url.Values, queryParam QueryParam) error {
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
