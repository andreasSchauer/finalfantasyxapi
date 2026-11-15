package apitesttypes

import (
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
)

// takes an expected path and checks, if it matches the URL of target struct's API Resource.
// /endpoint/23 == struct with api resource linking to host/api/endpoint/23
func (test Test) compPathAPIResource[T apiresource.HasAPIResource](fn, expPath string, gotRes T) {
	test.T.Helper()

	if test.DontCheck != nil && test.DontCheck[fn] {
		return
	}

	gotURL := gotRes.GetAPIResource().GetURL()
	gotPath := ah.URLToPath(gotURL)

	test.compare(fn, expPath, gotPath)
}

// takes an expected path ptr and checks, if it matches the URL of target struct ptr's API Resource.
// /endpoint/23 == struct with api resource linking to host/api/endpoint/23 (both ptrs)
func (test Test) compPathAPIResourcePtrs[T apiresource.HasAPIResource](fn string, expPathPtr *string, gotResPtr *T) {
	test.T.Helper()

	if test.DontCheck != nil && test.DontCheck[fn] {
		return
	}

	if !test.presenceResPtrs(fn, expPathPtr, gotResPtr) {
		return
	}

	test.compPathAPIResource(fn, *expPathPtr, *gotResPtr)
}

// takes an id, assembles the url, and checks, if it matches the URL of target struct's API Resource.
// 23 == struct with api resource linking to host/api/endpoint/23
func (test Test) compIDAPIResource[T apiresource.HasAPIResource](fn string, endpoint ts.EndpointName, expID int32, gotRes T) {
	test.T.Helper()
	expPath := ah.CompleteTestPath(endpoint, expID)
	test.compPathAPIResource(fn, expPath, gotRes)
}

// compIDAPIResourcePtrs takes an id ptr, assembles the url, and checks, if it matches the URL of target struct ptr's API Resource.
// 23 == struct with api resource linking to host/api/endpoint/23 (both ptrs)
func (test Test) compIDAPIResourcePtrs[T apiresource.HasAPIResource](fn string, endpoint ts.EndpointName, expIDPtr *int32, gotResPtr *T) {
	test.T.Helper()

	var expPathPtr *string

	if expIDPtr == nil {
		expPathPtr = nil
	} else {
		expPath := ah.CompleteTestPath(endpoint, *expIDPtr)
		expPathPtr = &expPath
	}

	test.compPathAPIResourcePtrs(fn, expPathPtr, gotResPtr)
}
