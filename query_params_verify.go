package main

import (
	"fmt"
	"net/http"
	"net/url"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func verifyQueryParams(r *http.Request, endpoint string, id *int32, lookup map[string]QueryType) error {
	q := r.URL.Query()

	for query := range q {
		queryParam, ok := lookup[query]
		if !ok {
			return newHTTPError(http.StatusBadRequest, fmt.Sprintf("parameter '%s' does not exist for endpoint /%s.", query, endpoint), nil)
		}

		err := verifyRequiredParams(q, queryParam, query)
		if err != nil {
			return err
		}

		err = verifyQueryUsage(queryParam, query, id)
		if err != nil {
			return err
		}
	}

	return nil
}

func verifyRequiredParams(q url.Values, queryParam QueryType, query string) error {
	if queryParam.RequiredParams == nil {
		return nil
	}

	for _, reqParam := range queryParam.RequiredParams {
		reqParamVal := q.Get(reqParam)

		if reqParamVal == "" {
			return newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid usage of parameter '%s'. parameter '%s' can only be used in combination with parameter(s): %s.", query, query, h.FormatStringSlice(queryParam.RequiredParams)), nil)
		}
	}

	return nil
}

func verifyQueryUsage(queryParam QueryType, query string, id *int32) error {
	if queryParam.ForSingle {
		if id == nil {
			return newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid usage of parameter '%s'. parameter '%s' can only be used with single-resource-endpoints.", query, query), nil)
		}

		err := verifyAllowedIDs(queryParam, query, *id)
		if err != nil {
			return err
		}
	}

	if queryParam.ForList && id != nil {
		return newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid usage of parameter '%s'. parameter '%s' can only be used with list-endpoints.", query, query), nil)
	}

	return nil
}

func verifyAllowedIDs(queryType QueryType, param string, id int32) error {
	if queryType.AllowedIDs == nil {
		return nil
	}

	allowedIDPresent := false

	for _, reqID := range queryType.AllowedIDs {
		if id == reqID {
			allowedIDPresent = true
		}
	}
	if !allowedIDPresent {
		return newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid id '%d'. parameter '%s' can only be used with ids: %s.", id, param, h.FormatIntSlice(queryType.AllowedIDs)), nil)
	}

	return nil
}
