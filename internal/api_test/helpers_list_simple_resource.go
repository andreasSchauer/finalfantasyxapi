package apiTest

import "github.com/andreasSchauer/finalfantasyxapi/internal/api"

type gotSimpleResourceList[T api.APIResource, S api.SimpleResource] struct {
	api.ListParams
	ParentResource T   `json:"parent_resource,omitempty"`
	Results        []S `json:"results"`
}

func (l gotSimpleResourceList[T, S]) GetListParams() api.ListParams {
	return l.ListParams
}

func (l gotSimpleResourceList[T, S]) GetResults() []S {
	return l.Results
}

func compareSimpleResourceLists[T api.APIResource, S api.SimpleResource](test test, endpoint api.EndpointName, expList expListIDs, gotList gotSimpleResourceList[T, S]) {
	test.t.Helper()
	compareListParams(test, expList.GetListParams(), gotList.GetListParams())

	gotParentURL := gotList.ParentResource.GetURL()
	compPageURL(test, "parent resource", expList.parentResource, &gotParentURL)

	checkSubResIDsInSlice(test, "results", endpoint, expList.results, gotList.GetResults())
}

func checkSubResIDsInSlice[T api.SimpleResource](test test, fieldName string, endpoint api.EndpointName, expIDs []int32, gotRes []T) {
	test.t.Helper()
	sliceBasicChecks(test, fieldName, expIDs, gotRes)

	gotMap := api.GetSimpleResourceURLMap(gotRes)
	if len(gotMap) != len(gotRes) {
		test.t.Errorf("there appear to be duplicates in '%s'.\n\n", fieldName)
	}

	for _, expID := range expIDs {
		expURL := api.CreateResourceURL(test.cfg, endpoint, expID)
		_, ok := gotMap[expURL]
		if !ok {
			test.t.Errorf("'%s' doesn't contain all wanted resources. missing '%s'\n\n", fieldName, api.CompleteTestPath(endpoint, expID))
		}
	}
}
