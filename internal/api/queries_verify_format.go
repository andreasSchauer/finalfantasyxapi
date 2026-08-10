package api

import (
	"fmt"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)


// verifies, if the query param is used in the correct format, meaning if it's used for the correct type of request.
func vpFormat[T any](queryParam QueryParam, endpoint EndpointName, key *T, segment *string) error {
	err := vpSingle(queryParam, key)
	if err != nil {
		return err
	}
	
	err = vpList(queryParam, key)
	if err != nil {
		return err
	}
	
	err = vpSegmentOnly(queryParam, segment, endpoint)
	if err != nil {
		return err
	}

	return nil
}


// checks, if a query param that is meant for single resource requests is used in the correct context. returns an error, if no key is provided (meaning the parameter was combined with a list request). also returns an error if the key is an id and the given id is not among the allowed ids, and if that restriction exists for the query param.
func vpSingle[T any](queryParam QueryParam, keyPtr *T) error {
	if queryParam.ParamUse == puSingle {
		if keyPtr == nil {
			return errSingleResParam(queryParam.Name)
		}

		idPtr, ok := any(keyPtr).(*int32)
		if ok {
			err := verifyAllowedIDs(queryParam, *idPtr)
			if err != nil {
				return err
			}
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


// checks, if a query param that is meant for list requests is used in the correct context. returns an error, if a key is provided (meaning the parameter was combined with a single resource request).
func vpList[T any](queryParam QueryParam, key *T) error {
	if queryParam.ParamUse == puList && key != nil {
		return errListResParam(queryParam.Name)
	}
	return nil
}


// checks, if a query param that is meant for a specific segment is used in the correct context, for example /endpoint/simple. returns an error, if the given segment doesn't match with the requested segment.
func vpSegmentOnly(queryParam QueryParam, segment *string, endpoint EndpointName) error {
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
