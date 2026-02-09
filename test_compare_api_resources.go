package main


// takes an expected path and checks, if it matches the URL of target struct's API Resource.
// /endpoint/23 == struct with api resource linking to host/api/endpoint/23
func compPathApiResource[T HasAPIResource](test test, fieldName, expPath string, gotRes T) {
	test.t.Helper()

	if test.dontCheck != nil && test.dontCheck[fieldName] {
		return
	}

	expURL := completeTestURL(test.cfg, expPath)
	gotURL := gotRes.GetAPIResource().GetURL()

	compare(test, fieldName, expURL, gotURL)
}


// takes an expected path ptr and checks, if it matches the URL of target struct ptr's API Resource.
// /endpoint/23 == struct with api resource linking to host/api/endpoint/23 (both ptrs)
func compPathApiResourcePtrs[T HasAPIResource](test test, fieldName string, expPathPtr *string, gotResPtr *T) {
	test.t.Helper()

	if test.dontCheck != nil && test.dontCheck[fieldName] {
		return
	}

	if !bothResourcePtrsPresent(test, fieldName, expPathPtr, gotResPtr) {
		return
	}

	compPathApiResource(test, fieldName, *expPathPtr, *gotResPtr)
}


// takes an id, assembles the url, and checks, if it matches the URL of target struct's API Resource.
// 23 == struct with api resource linking to host/api/endpoint/23
func compIdApiResource[T HasAPIResource](test test, fieldName, endpoint string, expID int32, gotRes T) {
	expPath := completeTestPath(endpoint, expID)
	compPathApiResource(test, fieldName, expPath, gotRes)
}


// takes an id ptr, assembles the url, and checks, if it matches the URL of target struct ptr's API Resource.
// 23 == struct with api resource linking to host/api/endpoint/23 (both ptrs)
func compIdApiResourcePtrs[T HasAPIResource](test test, fieldName, endpoint string, expIDPtr *int32, gotResPtr *T) {
	test.t.Helper()

	var expPathPtr *string

	if expIDPtr == nil {
		expPathPtr = nil
	} else {
		expPath := completeTestPath(endpoint, *expIDPtr)
		expPathPtr = &expPath
	}

	compPathApiResourcePtrs(test, fieldName, expPathPtr, gotResPtr)
}
