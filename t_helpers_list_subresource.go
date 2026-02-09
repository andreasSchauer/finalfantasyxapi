package main

type gotSubResourceList[T APIResource, S SubResource] struct {
	ListParams
	ParentResource T   `json:"parent_resource,omitempty"`
	Results        []S `json:"results"`
}

func (l gotSubResourceList[T, S]) getListParams() ListParams {
	return l.ListParams
}

func (l gotSubResourceList[T, S]) getResults() []S {
	return l.Results
}


func compareSubResourceLists[T APIResource, S SubResource](test test, endpoint string, expList expListIDs, gotList gotSubResourceList[T, S]) {
	test.t.Helper()
	compareListParams(test, expList.getListParams(), gotList.getListParams())

	gotParentURL := gotList.ParentResource.GetURL()
	compPageURL(test, "parent resource", expList.parentResource, &gotParentURL)

	checkSubResIDsInSlice(test, "results", endpoint, expList.results, gotList.getResults())
}


func checkSubResIDsInSlice[T SubResource](test test, fieldName, endpoint string, expIDs []int32, gotRes []T) {
	sliceBasicChecks(test, fieldName, expIDs, gotRes)
	
	gotMap := getSubResourceURLMap(gotRes)
	if len(gotMap) != len(gotRes) {
		test.t.Fatalf("%s: there appear to be duplicates in '%s'.", test.name, fieldName)
	}

	for _, expID := range expIDs {
		expURL := createResourceURL(test.cfg, endpoint, expID)
		_, ok := gotMap[expURL]
		if !ok {
			test.t.Fatalf("%s: '%s' doesn't contain all wanted resources. missing '%s'.", test.name, fieldName, expURL)
		}
	}
}
