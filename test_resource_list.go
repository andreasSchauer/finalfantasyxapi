package main


// checks the basic fields of an APIResourceList (count, pagination urls) and then checks for the stated resources
func compareAPIResourceLists[T APIResourceList](test test, endpoint string, expList expListIDs, gotList T) {
	test.t.Helper()
	compareListParams(test, expList.getListParams(), gotList.getListParams())
	checkResIDsInSlice(test, "results", endpoint, expList.results, gotList.getResults())
}


// checks if both slices are present and if the provided slice of resources contains all stated resources
func checkResPathsInSlice[T HasAPIResource](test test, fieldName string, expectedPaths []string, gotRes []T) {
	sliceBasicChecks(test, fieldName, expectedPaths, gotRes)

	gotMap := getResourceURLMap(gotRes)
	if len(gotMap) != len(gotRes) {
		test.t.Fatalf("%s: there appear to be duplicates in '%s'.", test.name, fieldName)
	}

	for _, expPath := range expectedPaths {
		expURL := completeTestURL(test.cfg, expPath)
		_, ok := gotMap[expURL]
		if !ok {
			test.t.Fatalf("%s: '%s' doesn't contain all wanted resources. missing '%s'.", test.name, fieldName, expURL)
		}
	}
}


func checkResIDsInSlice[T HasAPIResource](test test, fieldName, endpoint string, expectedIDs []int32, gotRes []T) {
	sliceBasicChecks(test, fieldName, expectedIDs, gotRes)

	gotMap := getResourceURLMap(gotRes)
	if len(gotMap) != len(gotRes) {
		test.t.Fatalf("%s: there appear to be duplicates in '%s'.", test.name, fieldName)
	}

	for _, expID := range expectedIDs {
		expURL := createResourceURL(test.cfg, endpoint, expID)
		_, ok := gotMap[expURL]
		if !ok {
			test.t.Fatalf("%s: '%s' doesn't contain all wanted resources. missing '%s'.", test.name, fieldName, expURL)
		}
	}
}
