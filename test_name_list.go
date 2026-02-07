package main

type nameListTest struct {
	name string
	exp  []string
	got  []string
}

func compareSectionLists(test test, endpoint string, exp expListNames, got SectionList) {
	expURLs := []string{}

	for _, section := range exp.results {
		url := createSectionURL(test.cfg, endpoint, section)
		expURLs = append(expURLs, url)
	}

	nameListTest := nameListTest{
		name: "results",
		exp:  expURLs,
		got:  got.Results,
	}

	compareNameLists(test, nameListTest)
}

func compareParameterLists(test test, _ string, exp expListNames, got QueryParameterList) {
	gotNames := []string{}

	for _, param := range got.Results {
		gotNames = append(gotNames, param.Name)
	}

	nameListTest := nameListTest{
		name: "results",
		exp:  exp.results,
		got:  gotNames,
	}

	compareNameLists(test, nameListTest)
}

// checks if the provided slice of names contains all stated names and also checks its length, if stated
func compareNameLists(test test, nameTest nameListTest) {
	test.t.Helper()

	if len(nameTest.exp) == 0 {
		return
	}

	checkNamesInSlice(test, nameTest)

	expLen, ok := test.expLengths[nameTest.name]
	if !ok {
		return
	}

	compare(test, nameTest.name+" length", expLen, len(nameTest.got))
}

// checks if the provided slice of names contains all stated names
func checkNamesInSlice(test test, nameTest nameListTest) {
	gotMap := make(map[string]bool)
	for _, gotName := range nameTest.got {
		gotMap[gotName] = true
	}

	if len(gotMap) != len(nameTest.got) {
		test.t.Fatalf("%s: there appear to be duplicates in '%s'", test.name, nameTest.name)
	}

	for _, expName := range nameTest.exp {
		_, ok := gotMap[expName]
		if !ok {
			test.t.Fatalf("%s: %s doesn't contain all wanted names. missing '%s'.", test.name, nameTest.name, expName)
		}
	}
}
