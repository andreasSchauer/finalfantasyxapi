package apitest

import (
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

type gotSimpleResourceList[T apiresource.APIResource, S apiresource.SimpleResource] struct {
	apiresource.ListParams
	ParentResource T   `json:"parent_resource,omitempty"`
	Results        []S `json:"results"`
}

func (l gotSimpleResourceList[T, S]) GetListParams() apiresource.ListParams {
	return l.ListParams
}

func (l gotSimpleResourceList[T, S]) GetResults() []S {
	return l.Results
}

func compareSimpleResourceLists[T apiresource.APIResource, S apiresource.SimpleResource](test test, endpoint ts.EndpointName, expList expListIDs, gotList gotSimpleResourceList[T, S]) {
	test.t.Helper()
	test.compareListParams(expList.GetListParams(), gotList.GetListParams())

	gotParentURL := gotList.ParentResource.GetURL()
	test.compPageURL("parent resource", expList.parentResource, &gotParentURL)

	test.checkSubResIDsInSlice("results", endpoint, expList.results, gotList.GetResults())
}

func (test test) checkSubResIDsInSlice[T apiresource.SimpleResource](fieldName string, endpoint ts.EndpointName, expIDs []int32, gotRes []T) {
	test.t.Helper()
	test.checkSliceBasic(fieldName, expIDs, gotRes)

	gotMap := handlerinput.GetSimpleResourceURLMap(gotRes)
	if len(gotMap) != len(gotRes) {
		test.t.Errorf("there appear to be duplicates in '%s'.\n\n", fieldName)
	}

	for _, expID := range expIDs {
		expURL := ah.CreateResourceURL(endpoint, expID)
		_, ok := gotMap[expURL]
		if !ok {
			test.t.Errorf("'%s' doesn't contain all wanted resources. missing '%s'\n\n", fieldName, ah.CompleteTestPath(endpoint, expID))
		}
	}
}
