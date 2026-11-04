package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

// checks the basic fields of an APIResourceList (count, pagination urls) and then checks for the stated resources
func compareAPIResourceLists[A api.APIResource](test test, endpoint ts.EndpointName, expList expListIDs, gotList api.APIResourceList[A]) {
	test.t.Helper()
	test.compareListParams(expList.GetListParams(), gotList.GetListParams())
	test.checkResIDsInSlice("results", endpoint, expList.results, gotList.GetResults())
}

// checks the basic fields of an APIResourceList (count, pagination urls) and then checks for the stated resource paths
func compareAPIResourceListPaths[A api.APIResource](test test, expList expListPaths, gotList api.APIResourceList[A]) {
	test.t.Helper()
	test.compareListParams(expList.GetListParams(), gotList.GetListParams())
	test.checkResPathsInSlice("results", expList.results, gotList.GetResults())
}

// checks if both slices are present and if the provided slice of resources contains all stated resources
func (test test) checkResPathsInSlice[T api.HasAPIResource](fieldName string, expPaths []string, gotRes []T) {
	test.t.Helper()
	test.checkSliceBasic(fieldName, expPaths, gotRes)

	gotMap := api.GetResourceURLMap(gotRes)
	if len(gotMap) != len(gotRes) {
		test.t.Errorf("there appear to be duplicates in '%s'.\n\n", fieldName)
	}

	for _, expPath := range expPaths {
		expURL := api.CompleteTestURL(test.srv, expPath)
		_, ok := gotMap[expURL]
		if !ok {
			test.t.Errorf("'%s' doesn't contain all wanted resources. missing '%s'.\n\n", fieldName, expPath)
		}
	}
}

func (test test) checkResIDsInSlice[T api.HasAPIResource](fieldName string, endpoint ts.EndpointName, expIDs []int32, gotRes []T) {
	test.t.Helper()
	test.checkSliceBasic(fieldName, expIDs, gotRes)

	gotMap := api.GetResourceURLMap(gotRes)
	if len(gotMap) != len(gotRes) {
		test.t.Errorf("there appear to be duplicates in '%s'.\n\n", fieldName)
	}

	for _, expID := range expIDs {
		expURL := api.CreateResourceURL(test.srv, endpoint, expID)
		_, ok := gotMap[expURL]
		if !ok {
			test.t.Errorf("'%s' doesn't contain all wanted resources. missing '%s'.\n\n", fieldName, api.CompleteTestPath(endpoint, expID))
		}
	}
}
