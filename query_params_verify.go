package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func verifyQueryParams(r *http.Request, endpoint string, id *int32, lookup map[string]QueryType) error {
	q := r.URL.Query()

	for param := range q {
		queryType, ok := lookup[param]
		if !ok {
			return newHTTPError(http.StatusBadRequest, fmt.Sprintf("Parameter %s does not exist for endpoint %s.", param, endpoint), nil)
		}

		err := verifyRequiredParams(q, queryType, param)
		if err != nil {
			return err
		}

		err = verifyQueryUsage(queryType, param, id)
		if err != nil {
			return err
		}
	}

	return nil
}

func verifyRequiredParams(q url.Values, queryType QueryType, param string) error {
	if queryType.RequiredWith == nil {
		return nil
	}

	for _, reqParam := range queryType.RequiredWith {
		reqParamVal := q.Get(reqParam)
		if reqParamVal == "" {
			return newHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid usage of parameter %s. Parameter %s can only be used in combination with parameter(s): %s.", param, param, strings.Join(queryType.RequiredWith, ", ")), nil)
		}
	}

	return nil
}

func verifyQueryUsage(queryType QueryType, param string, id *int32) error {
	if queryType.ForSingle {
		if id == nil {
			return newHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid usage of parameter %s. Parameter %s can only be used with single-resource-endpoints.", param, param), nil)
		}

		err := verifyAllowedIDs(queryType, param, *id)
		if err != nil {
			return err
		}
	}

	if queryType.ForList && id != nil {
		return newHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid usage of parameter %s. Parameter %s can only be used with list-endpoints.", param, param), nil)
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
		idsString := getIDsString(queryType.AllowedIDs)
		return newHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid id %d. Parameter %s can only be used with ids %s.", id, param, idsString), nil)
	}

	return nil
}

func getIDsString(IDs []int32) string {
	return strings.Trim(strings.Join(strings.Split(fmt.Sprint(IDs), " "), ", "), "[]")
}