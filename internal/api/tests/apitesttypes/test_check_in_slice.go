package apitesttypes

import (
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
)

// CheckResPathsInSlice checks if both slices are present and if the provided slice of resources contains all stated resources
func (test Test) CheckResPathsInSlice[T apiresource.HasAPIResource](fn string, expPaths []string, gotRes []T) {
	test.T.Helper()
	test.checkSliceBasic(fn, expPaths, gotRes)

	gotMap := handlerinput.GetResourceURLMap(gotRes)
	if len(gotMap) != len(gotRes) {
		test.T.Errorf("there appear to be duplicates in '%s'.\n\n", fn)
	}

	for _, expPath := range expPaths {
		expURL := ah.CompleteTestURL(expPath)
		_, ok := gotMap[expURL]
		if !ok {
			test.T.Errorf("'%s' doesn't contain all wanted resources. missing '%s'.\n\n", fn, expPath)
		}
	}
}

func (test Test) CheckResIDsInSlice[T apiresource.HasAPIResource](fn string, endpoint ts.EndpointName, expIDs []int32, gotRes []T) {
	test.T.Helper()
	test.checkSliceBasic(fn, expIDs, gotRes)

	gotMap := handlerinput.GetResourceURLMap(gotRes)
	if len(gotMap) != len(gotRes) {
		test.T.Errorf("there appear to be duplicates in '%s'.\n\n", fn)
	}

	for _, expID := range expIDs {
		expURL := ah.CreateResourceURL(endpoint, expID)
		_, ok := gotMap[expURL]
		if !ok {
			test.T.Errorf("'%s' doesn't contain all wanted resources. missing '%s'.\n\n", fn, ah.CompleteTestPath(endpoint, expID))
		}
	}
}

// CheckStringsInSlice checks if the provided slice of strings contains all stated strings
func (test Test) CheckStringsInSlice(fn string, expNames []ts.NamedParam, gotNames []string) {
	test.T.Helper()
	test.checkSliceBasic(fn, expNames, gotNames)

	gotMap := make(map[string]bool)
	for _, gotName := range gotNames {
		gotMap[gotName] = true
	}

	if len(gotMap) != len(gotNames) {
		test.T.Errorf("there appear to be duplicates in '%s'\n\n", fn)
	}

	for _, expName := range expNames {
		_, ok := gotMap[string(expName)]
		if !ok {
			test.T.Errorf("%s doesn't contain all wanted names. missing '%s'.\n\n", fn, expName)
		}
	}
}

func (test Test) CheckSubResIDsInSlice[T apiresource.SimpleResource](fn string, endpoint ts.EndpointName, expIDs []int32, gotRes []T) {
	test.T.Helper()
	test.checkSliceBasic(fn, expIDs, gotRes)

	gotMap := handlerinput.GetSimpleResourceURLMap(gotRes)
	if len(gotMap) != len(gotRes) {
		test.T.Errorf("there appear to be duplicates in '%s'.\n\n", fn)
	}

	for _, expID := range expIDs {
		expURL := ah.CreateResourceURL(endpoint, expID)
		_, ok := gotMap[expURL]
		if !ok {
			test.T.Errorf("'%s' doesn't contain all wanted resources. missing '%s'\n\n", fn, ah.CompleteTestPath(endpoint, expID))
		}
	}
}
