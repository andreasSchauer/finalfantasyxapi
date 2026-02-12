package api

func compareParameterLists(test test, _ string, exp expListNames, got QueryParameterList) {
	test.t.Helper()
	compareListParams(test, exp.getListParams(), got.getListParams())
	gotNames := []string{}

	for _, param := range got.Results {
		gotNames = append(gotNames, param.Name)
	}

	checkStringsInSlice(test, "results", exp.results, gotNames)
}

func compareSectionLists(test test, endpoint string, exp expListNames, got SectionList) {
	test.t.Helper()
	compareListParams(test, exp.getListParams(), got.getListParams())
	expURLs := []string{}

	for _, section := range exp.results {
		url := createSectionURL(test.cfg, endpoint, section)
		expURLs = append(expURLs, url)
	}

	checkStringsInSlice(test, "results", expURLs, got.Results)
}

// checks if the provided slice of strings contains all stated strings
func checkStringsInSlice(test test, fieldName string, expNames, gotNames []string) {
	sliceBasicChecks(test, fieldName, expNames, gotNames)

	gotMap := make(map[string]bool)
	for _, gotName := range gotNames {
		gotMap[gotName] = true
	}

	if len(gotMap) != len(gotNames) {
		test.t.Fatalf("%s: there appear to be duplicates in '%s'", test.name, fieldName)
	}

	for _, expName := range expNames {
		_, ok := gotMap[expName]
		if !ok {
			test.t.Fatalf("%s: %s doesn't contain all wanted names. missing '%s'.", test.name, fieldName, expName)
		}
	}
}
