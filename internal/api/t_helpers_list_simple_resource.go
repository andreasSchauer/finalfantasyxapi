package api

type gotSimpleResourceList[T APIResource, S SimpleResource] struct {
	ListParams
	ParentResource T   `json:"parent_resource,omitempty"`
	Results        []S `json:"results"`
}

func (l gotSimpleResourceList[T, S]) getListParams() ListParams {
	return l.ListParams
}

func (l gotSimpleResourceList[T, S]) getResults() []S {
	return l.Results
}

func compareSimpleResourceLists[T APIResource, S SimpleResource](test test, endpoint string, expList expListIDs, gotList gotSimpleResourceList[T, S]) {
	test.t.Helper()
	compareListParams(test, expList.getListParams(), gotList.getListParams())

	gotParentURL := gotList.ParentResource.GetURL()
	compPageURL(test, "parent resource", expList.parentResource, &gotParentURL)

	checkSubResIDsInSlice(test, "results", endpoint, expList.results, gotList.getResults())
}

func checkSubResIDsInSlice[T SimpleResource](test test, fieldName, endpoint string, expIDs []int32, gotRes []T) {
	test.t.Helper()
	sliceBasicChecks(test, fieldName, expIDs, gotRes)

	gotMap := getSimpleResourceURLMap(gotRes)
	if len(gotMap) != len(gotRes) {
		test.t.Errorf("there appear to be duplicates in '%s'.\n\n", fieldName)
	}

	for _, expID := range expIDs {
		expURL := createResourceURL(test.cfg, endpoint, expID)
		_, ok := gotMap[expURL]
		if !ok {
			test.t.Errorf("'%s' doesn't contain all wanted resources. missing '%s'\n\n", fieldName, completeTestPath(endpoint, expID))
		}
	}
}
