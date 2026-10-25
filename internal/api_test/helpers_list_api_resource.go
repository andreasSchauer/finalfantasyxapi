package apiTest

import "github.com/andreasSchauer/finalfantasyxapi/internal/api"

// checks the basic fields of an APIResourceList (count, pagination urls) and then checks for the stated resources
func compareAPIResourceLists[T api.APIResourceList](test test, endpoint api.EndpointName, expList expListIDs, gotList T) {
	test.t.Helper()
	compareListParams(test, expList.GetListParams(), gotList.GetListParams())
	checkResIDsInSlice(test, "results", endpoint, expList.results, gotList.GetResults())
}

// checks the basic fields of an APIResourceList (count, pagination urls) and then checks for the stated resource paths
func compareAPIResourceListPaths[T api.APIResourceList](test test, expList expListPaths, gotList T) {
	test.t.Helper()
	compareListParams(test, expList.GetListParams(), gotList.GetListParams())
	checkResPathsInSlice(test, "results", expList.results, gotList.GetResults())
}

// checks if both slices are present and if the provided slice of resources contains all stated resources
func checkResPathsInSlice[T api.HasAPIResource](test test, fieldName string, expPaths []string, gotRes []T) {
	test.t.Helper()
	sliceBasicChecks(test, fieldName, expPaths, gotRes)

	gotMap := api.GetResourceURLMap(gotRes)
	if len(gotMap) != len(gotRes) {
		test.t.Errorf("there appear to be duplicates in '%s'.\n\n", fieldName)
	}

	for _, expPath := range expPaths {
		expURL := api.CompleteTestURL(test.cfg, expPath)
		_, ok := gotMap[expURL]
		if !ok {
			test.t.Errorf("'%s' doesn't contain all wanted resources. missing '%s'.\n\n", fieldName, expPath)
		}
	}
}

func checkResIDsInSlice[T api.HasAPIResource](test test, fieldName string, endpoint api.EndpointName, expIDs []int32, gotRes []T) {
	test.t.Helper()
	sliceBasicChecks(test, fieldName, expIDs, gotRes)

	gotMap := api.GetResourceURLMap(gotRes)
	if len(gotMap) != len(gotRes) {
		test.t.Errorf("there appear to be duplicates in '%s'.\n\n", fieldName)
	}

	for _, expID := range expIDs {
		expURL := api.CreateResourceURL(test.cfg, endpoint, expID)
		_, ok := gotMap[expURL]
		if !ok {
			test.t.Errorf("'%s' doesn't contain all wanted resources. missing '%s'.\n\n", fieldName, api.CompleteTestPath(endpoint, expID))
		}
	}
}
