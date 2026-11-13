package apitest

import (
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
)



// checks if both slices are present and if the provided slice of resources contains all stated resources
func (test test) checkResPathsInSlice[T apiresource.HasAPIResource](fieldName string, expPaths []string, gotRes []T) {
	test.t.Helper()
	test.checkSliceBasic(fieldName, expPaths, gotRes)

	gotMap := handlerinput.GetResourceURLMap(gotRes)
	if len(gotMap) != len(gotRes) {
		test.t.Errorf("there appear to be duplicates in '%s'.\n\n", fieldName)
	}

	for _, expPath := range expPaths {
		expURL := ah.CompleteTestURL(expPath)
		_, ok := gotMap[expURL]
		if !ok {
			test.t.Errorf("'%s' doesn't contain all wanted resources. missing '%s'.\n\n", fieldName, expPath)
		}
	}
}

func (test test) checkResIDsInSlice[T apiresource.HasAPIResource](fieldName string, endpoint ts.EndpointName, expIDs []int32, gotRes []T) {
	test.t.Helper()
	test.checkSliceBasic(fieldName, expIDs, gotRes)

	gotMap := handlerinput.GetResourceURLMap(gotRes)
	if len(gotMap) != len(gotRes) {
		test.t.Errorf("there appear to be duplicates in '%s'.\n\n", fieldName)
	}

	for _, expID := range expIDs {
		expURL := ah.CreateResourceURL(endpoint, expID)
		_, ok := gotMap[expURL]
		if !ok {
			test.t.Errorf("'%s' doesn't contain all wanted resources. missing '%s'.\n\n", fieldName, ah.CompleteTestPath(endpoint, expID))
		}
	}
}

// checks if the provided slice of strings contains all stated strings
func (test test) checkStringsInSlice(fieldName string, expNames []ts.NamedParam, gotNames []string) {
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