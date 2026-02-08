package main

type resListTest struct {
	name string
	exp  []string
	got  []HasAPIResource
}

func rltIDs[T HasAPIResource](fieldName, endpoint string, expIDs []int32, got []T) resListTest {
	exp := []string{}

	for _, id := range expIDs {
		path := completeTestPath(endpoint, id)
		exp = append(exp, path)
	}

	return rlt(fieldName, exp, got)
}

func rlt[T HasAPIResource](fieldName string, exp []string, got []T) resListTest {
	return resListTest{
		name: fieldName,
		exp:  exp,
		got:  toHasAPIResSlice(got),
	}
}


// checks the basic fields of an APIResourceList (count, pagination urls) and then checks for the stated resources
func compareAPIResourceLists[T APIResourceList](test test, endpoint string, expList expListIDs, gotList T) {
	test.t.Helper()
	compareListParams(test, expList.getListParams(), gotList.getListParams())

	listTest := rltIDs("results", endpoint, expList.results, gotList.getResults())
	compareResListTest(test, listTest)
}

// checks if all provided slices of resources contains all stated resources and also checks their length, if stated
func compareResListTests(test test, checks []resListTest) {
	test.t.Helper()

	for _, c := range checks {
		compareResListTest(test, c)
	}
}

// checks if the provided slice of resources contains all stated resources and also checks its length, if stated
func compareResListTest(test test, expList resListTest) {
	test.t.Helper()

	compLength(test, expList.name, len(expList.got))
	checkResourcesInSlice(test, expList.name, expList.exp, expList.got)
}

// checks if the provided slice of resources contains all stated resources
func checkResourcesInSlice[T HasAPIResource](test test, fieldName string, expectedPaths []string, gotRes []T) {
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
