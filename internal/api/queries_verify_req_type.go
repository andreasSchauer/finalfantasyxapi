package api

import (
	"fmt"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

// verifies, if the query param is used in the correct type of request.
func (q queryVerifyer[T]) requestType() error {
	err := q.single()
	if err != nil {
		return err
	}

	err = q.list()
	if err != nil {
		return err
	}

	err = q.subsection()
	if err != nil {
		return err
	}

	return nil
}

// checks, if a query param that is meant for single resource requests is used in the correct context. returns an error, if no key is provided (meaning the parameter was combined with a list request). also returns an error if the key is an id and the given id is not among the allowed ids, and if that restriction exists for the query param.
func (q queryVerifyer[T]) single() error {
	if q.param.ParamUse == puSingle {
		if q.key == nil {
			return errSingleResParam(q.param.Name)
		}

		idPtr, ok := any(q.key).(*int32)
		if ok {
			err := q.allowedIDs(*idPtr)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// checks, if the requested query param is used on an id that the query param expects.
func (q queryVerifyer[T]) allowedIDs(id int32) error {
	if q.param.AllowedIDs == nil {
		return nil
	}

	allowedIDPresent := false

	for _, reqID := range q.param.AllowedIDs {
		if id == reqID {
			allowedIDPresent = true
		}
	}
	if !allowedIDPresent {
		return newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid id '%d'. parameter '%s' can only be used with ids: %s", id, q.param.Name, h.FormatIntSlice(q.param.AllowedIDs)), nil)
	}

	return nil
}

// checks, if a query param that is meant for list requests is used in the correct context. returns an error, if a key is provided (meaning the parameter was combined with a single resource request).
func (q queryVerifyer[T]) list() error {
	if q.param.ParamUse == puList && q.key != nil {
		return errListResParam(q.param.Name)
	}

	return nil
}

// checks, if a query param that is meant for a specific segment is used in the correct context, for example /endpoint/simple. returns an error, if the given segment doesn't match with the requested segment.
func (q queryVerifyer[T]) subsection() error {
	if q.param.ForSegment != nil && !segmentsMatch(q.param.ForSegment, q.segment) {
		return newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid usage of parameter '%s'. parameter '%s' can only be used in the following format: '/api/%s/%s%s'", q.param.Name, q.param.Name, q.endpoint, *q.param.ForSegment, q.param.Usage), nil)
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
