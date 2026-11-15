// Package apitest ajsdlaösdalk
package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apitesttypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
)

// checks the basic fields of an apiresource.APIResourceList (count, pagination urls) and then checks for the stated resources
func compareAPIResourceLists[A apiresource.APIResource](test apitesttypes.Test, endpoint ts.EndpointName, expList apitesttypes.ListIDs, gotList apiresource.APIResourceList[A]) {
	test.T.Helper()
	test.CompareListParams(expList.GetListParams(), gotList.GetListParams())
	test.CheckResIDsInSlice("results", endpoint, expList.Results, gotList.GetResults())
}

// checks the basic fields of an apiresource.APIResourceList (count, pagination urls) and then checks for the stated resource paths
func compareAPIResourceListPaths[A apiresource.APIResource](test apitesttypes.Test, expList apitesttypes.ListPaths, gotList apiresource.APIResourceList[A]) {
	test.T.Helper()
	test.CompareListParams(expList.GetListParams(), gotList.GetListParams())
	test.CheckResPathsInSlice("results", expList.Results, gotList.GetResults())
}

func compareParameterLists(test apitesttypes.Test, _ ts.EndpointName, exp apitesttypes.ListNames, got handlerinput.QueryParameterList) {
	test.T.Helper()
	test.CompareListParams(exp.GetListParams(), got.GetListParams())
	gotNames := []string{}

	for _, param := range got.Results {
		gotNames = append(gotNames, string(param.Name))
	}

	test.CheckStringsInSlice("results", exp.Results, gotNames)
}

func compareSectionLists(test apitesttypes.Test, endpoint ts.EndpointName, exp apitesttypes.ListNames, got handlerinput.SectionList) {
	test.T.Helper()

	test.CompareListParams(exp.GetListParams(), got.GetListParams())
	test.CheckStringsInSlice("results", exp.Results, got.Results)
}
