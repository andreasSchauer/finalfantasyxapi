package api

import (
	"fmt"
	"net/http"
	"net/url"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func verifyDefaultParamsOnly[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], segment *string) error {
	q := r.URL.Query()

	for query := range q {
		queryParam, ok := cfg.q.defaultParams[query]
		if !ok {
			return newHTTPError(http.StatusBadRequest, fmt.Sprintf("only default parameters are allowed when using /api/%s/%s. available default parameters: %s.", i.endpoint, *segment, queryMapToString(cfg.q.defaultParams)), nil)
		}

		err := verifyQueryUsage(q, queryParam, i.endpoint, nil, segment)
		if err != nil {
			return err
		}
	}

	return nil
}

func verifyQueryParams[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], id *int32, segment *string) error {
	q := r.URL.Query()
	canUseDefaultOnlyParam := verifyDefaultOnlyParam(cfg, q, i.queryLookup)

	for query := range q {
		queryParam, ok := i.queryLookup[query]
		if !ok {
			return newHTTPError(http.StatusBadRequest, fmt.Sprintf("parameter '%s' does not exist for endpoint /%s. use /api/%s/parameters for available parameters.", query, i.endpoint, i.endpoint), nil)
		}

		if queryParam.DefaultOnly && !canUseDefaultOnlyParam {
			return newHTTPError(http.StatusBadRequest, fmt.Sprintf("parameter '%s' can only be used with default parameters. available default parameters: %s.", queryParam.Name, queryMapToString(cfg.q.defaultParams)), nil)
		}

		err := verifyQueryUsage(q, queryParam, i.endpoint, id, segment)
		if err != nil {
			return err
		}
	}

	return nil
}

func verifyDefaultOnlyParam(cfg *Config, q url.Values, lookup map[string]QueryType) bool {
	defaultOnlyCount := 0

	for query := range q {
		queryParam, ok := lookup[query]
		if !ok {
			return false
		}

		if queryParam.DefaultOnly {
			defaultOnlyCount++
			if defaultOnlyCount > 1 {
				return false
			}

			continue
		}

		if !isDefaultParam(cfg, query) {
			return false
		}
	}

	return true
}

func isDefaultParam(cfg *Config, queryName string) bool {
	_, ok := cfg.q.defaultParams[queryName]
	return ok
}

func verifyQueryUsage(q url.Values, queryParam QueryType, endpoint string, id *int32, segment *string) error {
	if queryParam.ForSegment != nil && !segmentsMatch(queryParam.ForSegment, segment) {
		return newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid usage of parameter '%s'. parameter '%s' can only be used in the following format: '/api/%s/%s%s'.", queryParam.Name, queryParam.Name, endpoint, *queryParam.ForSegment, queryParam.Usage), nil)
	}

	if queryParam.ForSingle {
		if id == nil {
			return newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid usage of parameter '%s'. parameter '%s' can only be used with single-resource-endpoints.", queryParam.Name, queryParam.Name), nil)
		}

		err := verifyAllowedIDs(queryParam, *id)
		if err != nil {
			return err
		}
	}

	if queryParam.ForList && id != nil {
		return newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid usage of parameter '%s'. parameter '%s' can only be used with list-endpoints.", queryParam.Name, queryParam.Name), nil)
	}

	err := verifyRequiredParams(q, queryParam)
	if err != nil {
		return err
	}

	return nil
}

func segmentsMatch(sParam, sReq *string) bool {
	switch {
	case sParam == nil && sReq == nil:
		return true

	case sParam != nil && sReq != nil:
		return *sParam == *sReq

	default:
		return false
	}
}

func verifyAllowedIDs(queryParam QueryType, id int32) error {
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

func verifyRequiredParams(q url.Values, queryParam QueryType) error {
	if queryParam.RequiredParams == nil {
		return nil
	}

	for _, reqParam := range queryParam.RequiredParams {
		reqParamVal := q.Get(reqParam)

		if reqParamVal != "" {
			return nil
		}
	}

	return newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid usage of parameter '%s'. parameter '%s' can only be used in combination with parameter(s): %s.", queryParam.Name, queryParam.Name, h.FormatStringSlice(queryParam.RequiredParams)), nil)
}
