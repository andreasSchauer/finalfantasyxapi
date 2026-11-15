package apitest

import (
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
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

func compareSimpleResourceLists[T apiresource.APIResource, S apiresource.SimpleResource](test Test, endpoint ts.EndpointName, expList ListIDs, gotList gotSimpleResourceList[T, S]) {
	test.T.Helper()
	test.compareListParams(expList.GetListParams(), gotList.GetListParams())

	gotParentURL := gotList.ParentResource.GetURL()
	test.compPageURL("parent resource", expList.ParentResource, &gotParentURL)

	test.checkSubResIDsInSlice("results", endpoint, expList.Results, gotList.GetResults())
}

func (test Test) checkSubResIDsInSlice[T apiresource.SimpleResource](fieldName string, endpoint ts.EndpointName, expIDs []int32, gotRes []T) {
	test.T.Helper()
	test.checkSliceBasic(fieldName, expIDs, gotRes)

	gotMap := handlerinput.GetSimpleResourceURLMap(gotRes)
	if len(gotMap) != len(gotRes) {
		test.T.Errorf("there appear to be duplicates in '%s'.\n\n", fieldName)
	}

	for _, expID := range expIDs {
		expURL := ah.CreateResourceURL(endpoint, expID)
		_, ok := gotMap[expURL]
		if !ok {
			test.T.Errorf("'%s' doesn't contain all wanted resources. missing '%s'\n\n", fieldName, ah.CompleteTestPath(endpoint, expID))
		}
	}
}
