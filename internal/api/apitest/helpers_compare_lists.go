package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
)

// checks the basic fields of an apiresource.APIResourceList (count, pagination urls) and then checks for the stated resources
func compareAPIResourceLists[A apiresource.APIResource](test Test, endpoint ts.EndpointName, expList ListIDs, gotList apiresource.APIResourceList[A]) {
	test.T.Helper()
	test.compareListParams(expList.GetListParams(), gotList.GetListParams())
	test.checkResIDsInSlice("results", endpoint, expList.Results, gotList.GetResults())
}

// checks the basic fields of an apiresource.APIResourceList (count, pagination urls) and then checks for the stated resource paths
func compareAPIResourceListPaths[A apiresource.APIResource](test Test, expList ListPaths, gotList apiresource.APIResourceList[A]) {
	test.T.Helper()
	test.compareListParams(expList.GetListParams(), gotList.GetListParams())
	test.checkResPathsInSlice("results", expList.Results, gotList.GetResults())
}

func compareParameterLists(test Test, _ ts.EndpointName, exp ListNames, got handlerinput.QueryParameterList) {
	test.T.Helper()
	test.compareListParams(exp.GetListParams(), got.GetListParams())
	gotNames := []string{}

	for _, param := range got.Results {
		gotNames = append(gotNames, string(param.Name))
	}

	test.checkStringsInSlice("results", exp.Results, gotNames)
}

func compareSectionLists(test Test, endpoint ts.EndpointName, exp ListNames, got handlerinput.SectionList) {
	test.T.Helper()

	test.compareListParams(exp.GetListParams(), got.GetListParams())
	test.checkStringsInSlice("results", exp.Results, got.Results)
}
