package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
)

func compareParameterLists(test Test, _ ts.EndpointName, exp ListNames, got handlerinput.QueryParameterList) {
	test.t.Helper()
	test.compareListParams(exp.GetListParams(), got.GetListParams())
	gotNames := []string{}

	for _, param := range got.Results {
		gotNames = append(gotNames, string(param.Name))
	}

	test.checkStringsInSlice("results", exp.results, gotNames)
}

func compareSectionLists(test Test, endpoint ts.EndpointName, exp ListNames, got handlerinput.SectionList) {
	test.t.Helper()
	test.compareListParams(exp.GetListParams(), got.GetListParams())

	test.checkStringsInSlice("results", exp.results, got.Results)
}

// checks if the provided slice of strings contains all stated strings
func (test Test) checkStringsInSlice(fieldName string, expNames []ts.NamedParam, gotNames []string) {
	test.t.Helper()
	test.checkSliceBasic(fieldName, expNames, gotNames)

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
