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

type subResListTest struct {
	name string
	exp  []string
	got  []SubResource
}

func srltIDs[T SubResource](fieldName, endpoint string, expIDs []int32, got []T) subResListTest {
	exp := []string{}

	for _, id := range expIDs {
		path := completeTestPath(endpoint, id)
		exp = append(exp, path)
	}

	return srlt(fieldName, exp, got)
}

func srlt[T SubResource](fieldName string, exp []string, got []T) subResListTest {
	return subResListTest{
		name: fieldName,
		exp:  exp,
		got:  toSubResourceSlice(got),
	}
}

func compareSubResourceLists[T APIResource, S SubResource](test test, endpoint string, expList expListIDs, gotList gotSubResourceList[T, S]) {
	test.t.Helper()
	compareListParams(test, expList.getListParams(), gotList.getListParams())

	gotParentURL := gotList.ParentResource.GetURL()
	compPageURL(test, "parent resource", expList.parentResource, &gotParentURL)

	testSubResourceListResults(test, srltIDs("results", endpoint, expList.results, gotList.getResults()))
}

func testSubResourceListResults(test test, expList subResListTest) {
	test.t.Helper()

	compLength(test, expList.name, len(expList.got))
	checkSubResourcesInSlice(test, expList.name, expList.exp, expList.got)
}

func checkSubResourcesInSlice[T SubResource](test test, fieldName string, expectedPaths []string, gotRes []T) {
	gotMap := getSubResourceURLMap(gotRes)
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
