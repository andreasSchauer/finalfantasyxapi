package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api"
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api_helpers"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

// takes an expected path and checks, if it matches the URL of target struct's API Resource.
// /endpoint/23 == struct with api resource linking to host/api/endpoint/23
func (test test) compPathAPIResource[T api.HasAPIResource](fieldName, expPath string, gotRes T) {
	test.t.Helper()

	if test.dontCheck != nil && test.dontCheck[fieldName] {
		return
	}

	gotURL := gotRes.GetAPIResource().GetURL()
	gotPath := ah.URLToPath(gotURL)

	test.compare(fieldName, expPath, gotPath)
}

// takes an expected path ptr and checks, if it matches the URL of target struct ptr's API Resource.
// /endpoint/23 == struct with api resource linking to host/api/endpoint/23 (both ptrs)
func (test test) compPathAPIResourcePtrs[T api.HasAPIResource](fieldName string, expPathPtr *string, gotResPtr *T) {
	test.t.Helper()

	if test.dontCheck != nil && test.dontCheck[fieldName] {
		return
	}

	if !test.presenceResPtrs(fieldName, expPathPtr, gotResPtr) {
		return
	}

	test.compPathAPIResource(fieldName, *expPathPtr, *gotResPtr)
}

// takes an id, assembles the url, and checks, if it matches the URL of target struct's API Resource.
// 23 == struct with api resource linking to host/api/endpoint/23
func (test test) compIDAPIResource[T api.HasAPIResource](fieldName string, endpoint ts.EndpointName, expID int32, gotRes T) {
	test.t.Helper()
	expPath := ah.CompleteTestPath(endpoint, expID)
	test.compPathAPIResource(fieldName, expPath, gotRes)
}

// takes an id ptr, assembles the url, and checks, if it matches the URL of target struct ptr's API Resource.
// 23 == struct with api resource linking to host/api/endpoint/23 (both ptrs)
func (test test) compIDAPIResourcePtrs[T api.HasAPIResource](fieldName string, endpoint ts.EndpointName, expIDPtr *int32, gotResPtr *T) {
	test.t.Helper()

	var expPathPtr *string

	if expIDPtr == nil {
		expPathPtr = nil
	} else {
		expPath := ah.CompleteTestPath(endpoint, *expIDPtr)
		expPathPtr = &expPath
	}

	test.compPathAPIResourcePtrs(fieldName, expPathPtr, gotResPtr)
}
