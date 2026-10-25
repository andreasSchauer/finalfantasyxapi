package apiTest

import "github.com/andreasSchauer/finalfantasyxapi/internal/api"

func compareParameterLists(test test, _ api.EndpointName, exp expListNames, got api.QueryParameterList) {
	test.t.Helper()
	compareListParams(test, exp.GetListParams(), got.GetListParams())
	gotNames := []string{}

	for _, param := range got.Results {
		gotNames = append(gotNames, string(param.Name))
	}

	checkStringsInSlice(test, "results", exp.results, gotNames)
}

func compareSectionLists(test test, endpoint api.EndpointName, exp expListNames, got api.SectionList) {
	test.t.Helper()
	compareListParams(test, exp.GetListParams(), got.GetListParams())

	checkStringsInSlice(test, "results", exp.results, got.Results)
}

// checks if the provided slice of strings contains all stated strings
func checkStringsInSlice(test test, fieldName string, expNames []api.NamedParam, gotNames []string) {
	test.t.Helper()
	sliceBasicChecks(test, fieldName, expNames, gotNames)

	gotMap := make(map[string]bool)
	for _, gotName := range gotNames {
		gotMap[gotName] = true
	}

	if len(gotMap) != len(gotNames) {
		test.t.Errorf("there appear to be duplicates in '%s'\n\n", fieldName)
	}

	for _, expName := range expNames {
		_, ok := gotMap[string(expName)]
		if !ok {
			test.t.Errorf("%s doesn't contain all wanted names. missing '%s'.\n\n", fieldName, expName)
		}
	}
}
