package api

// checks the basic fields of an APIResourceList (count, pagination urls) and then checks for the stated resources
func compareAPIResourceLists[T APIResourceList](test test, endpoint EndpointName, expList expListIDs, gotList T) {
	test.t.Helper()
	compareListParams(test, expList.getListParams(), gotList.getListParams())
	checkResIDsInSlice(test, "results", endpoint, expList.results, gotList.getResults())
}

// checks if both slices are present and if the provided slice of resources contains all stated resources
func checkResPathsInSlice[T HasAPIResource](test test, fieldName string, expPaths []string, gotRes []T) {
	test.t.Helper()
	sliceBasicChecks(test, fieldName, expPaths, gotRes)

	gotMap := getResourceURLMap(gotRes)
	if len(gotMap) != len(gotRes) {
		test.t.Errorf("there appear to be duplicates in '%s'.\n\n", fieldName)
	}

	for _, expPath := range expPaths {
		expURL := completeTestURL(test.cfg, expPath)
		_, ok := gotMap[expURL]
		if !ok {
			test.t.Errorf("'%s' doesn't contain all wanted resources. missing '%s'.\n\n", fieldName, expPath)
		}
	}
}

func checkResIDsInSlice[T HasAPIResource](test test, fieldName string, endpoint EndpointName, expIDs []int32, gotRes []T) {
	test.t.Helper()
	sliceBasicChecks(test, fieldName, expIDs, gotRes)

	gotMap := getResourceURLMap(gotRes)
	if len(gotMap) != len(gotRes) {
		test.t.Errorf("there appear to be duplicates in '%s'.\n\n", fieldName)
	}

	for _, expID := range expIDs {
		expURL := createResourceURL(test.cfg, endpoint, expID)
		_, ok := gotMap[expURL]
		if !ok {
			test.t.Errorf("'%s' doesn't contain all wanted resources. missing '%s'.\n\n", fieldName, completeTestPath(endpoint, expID))
		}
	}
}
