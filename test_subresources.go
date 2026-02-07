package main

type SubResourceListTest[T APIResource, S SubResource] struct {
	ListParams
	ParentResource T   `json:"parent_resource,omitempty"`
	Results        []S `json:"results"`
}

func (l SubResourceListTest[T, S]) getListParams() ListParams {
	return l.ListParams
}

func (l SubResourceListTest[T, S]) getResults() []S {
	return l.Results
}

type subResListTest struct {
	name string
	exp  []string
	got  []SubResource
}

func newSubResListTestFromIDs[T SubResource](fieldName, endpoint string, expIDs []int32, got []T) subResListTest {
	exp := []string{}

	for _, id := range expIDs {
		path := completeTestPath(endpoint, id)
		exp = append(exp, path)
	}

	return subResListTest{
		name: fieldName,
		exp:  exp,
		got:  toSubResourceSlice(got),
	}
}

func compareSubResourceLists[T APIResource, S SubResource](test test, endpoint string, expList expListIDs, gotList SubResourceListTest[T, S]) {
	test.t.Helper()
	got := gotList.getListParams()
	compare(test, "count", expList.count, got.Count)

	compPageURL(test, "previous", expList.previous, got.Previous)
	compPageURL(test, "next", expList.next, got.Next)

	gotParentURL := gotList.ParentResource.GetURL()
	compPageURL(test, "parent resource", expList.parentResource, &gotParentURL)

	listTest := newSubResListTestFromIDs("results", endpoint, expList.results, gotList.getResults())
	testSubResourceListResults(test, listTest)
}

func testSubResourceListResults(test test, expList subResListTest) {
	test.t.Helper()

	if len(expList.exp) == 0 {
		return
	}
	checkSubResourcesInSlice(test, expList.name, expList.exp, expList.got)

	expLen, ok := test.expLengths[expList.name]
	if !ok {
		return
	}

	compare(test, expList.name+" length", expLen, len(expList.got))
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
