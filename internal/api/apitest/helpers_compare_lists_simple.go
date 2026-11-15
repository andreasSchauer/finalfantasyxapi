package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apitesttypes"
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

func compareSimpleResourceLists[T apiresource.APIResource, S apiresource.SimpleResource](test apitesttypes.Test, endpoint ts.EndpointName, expList apitesttypes.ListIDs, gotList gotSimpleResourceList[T, S]) {
	test.T.Helper()
	test.CompareListParams(expList.GetListParams(), gotList.GetListParams())

	gotParentURL := gotList.ParentResource.GetURL()
	test.CompPageURL("parent resource", expList.ParentResource, &gotParentURL)

	test.CheckSubResIDsInSlice("results", endpoint, expList.Results, gotList.GetResults())
}
