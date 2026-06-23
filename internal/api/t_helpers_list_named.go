package api

func compareParameterLists(test test, _ EndpointName, exp expListNames, got QueryParameterList) {
	test.t.Helper()
	compareListParams(test, exp.getListParams(), got.getListParams())
	gotNames := []string{}

	for _, param := range got.Results {
		gotNames = append(gotNames, string(param.Name))
	}

	checkStringsInSlice(test, "results", exp.results, gotNames)
}

func compareSectionLists(test test, endpoint EndpointName, exp expListNames, got SectionList) {
	test.t.Helper()
	compareListParams(test, exp.getListParams(), got.getListParams())

	checkStringsInSlice(test, "results", exp.results, got.Results)
}

// checks if the provided slice of strings contains all stated strings
func checkStringsInSlice(test test, fieldName string, expNames []NamedParam, gotNames []string) {
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
