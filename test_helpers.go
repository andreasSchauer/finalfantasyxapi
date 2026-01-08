package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type test struct {
	t *testing.T
	cfg *Config
	name string
	expLengths map[string]int
	dontCheck map[string]bool
}

type resListTest struct {
	name string
	exp  []string
	got  []HasAPIResource
}

func newResListTestFromIDs[T HasAPIResource](fieldName, endpoint string, expIDs []int32, got []T) resListTest {
	exp := []string{}

	for _, id := range expIDs {
		path := completeTestPath(endpoint, id)
		exp = append(exp, path)
	}

	return newResListTest(fieldName, exp, got)
}

func newResListTest[T HasAPIResource](fieldName string, exp []string, got []T) resListTest {
	return resListTest{
		name: fieldName,
		exp:  exp,
		got:  toHasAPIResSlice(got),
	}
}

func getTestName(name, requestURL string, caseNum int) string {
	return fmt.Sprintf("%s: %d, requestURL: %s", name, caseNum, requestURL)
}

// makes the http request for the test and returns the responseRecorder needed to decode the json
func setupTest(t *testing.T, tc testGeneral, testFunc string, testNum int, handlerFunc func(http.ResponseWriter, *http.Request)) (*httptest.ResponseRecorder, string, bool) {
	t.Helper()
	testName := getTestName(testFunc, tc.requestURL, testNum)
	caughtErr := false

	req := httptest.NewRequest(http.MethodGet, tc.requestURL, nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(handlerFunc)
	handler.ServeHTTP(rr, req)

	if rr.Code != tc.expectedStatus {
		t.Fatalf("%s: expected %d, got %d, body=%s", testName, tc.expectedStatus, rr.Code, rr.Body.String())
	}

	if tc.expectedErr != "" {
		raw := rr.Body.String()
		if !strings.Contains(raw, tc.expectedErr) {
			t.Fatalf("%s: expected error message to contain %s, got %q", testName, tc.expectedErr, raw)
		}
		caughtErr = true
		return nil, "", caughtErr
	}

	return rr, testName, caughtErr
}

// checks, if all basic fields of a resource with a unique name are equal
func testExpectedUnique(test test, tc expUnique, gotID int32, gotName string) {
	test.t.Helper()
	compare(test, "id", tc.id, gotID)
	compare(test, "name", tc.name, gotName)
}

// checks, if all basic fields of a resource that uses a name/version pattern are equal
func testExpectedNameVer(test test, tc expNameVer, gotID int32, gotName string, gotVer *int32) {
	test.t.Helper()
	compare(test, "id", tc.id, gotID)
	compare(test, "name", tc.name, gotName)
	compare(test, "version", tc.version, gotVer)
}

// checks the basic fields of an APIResourceList (count, pagination urls) and then checks for the stated resources
func testAPIResourceList[T IsAPIResourceList](test test, endpoint string, expList expList, gotList T) {
	test.t.Helper()
	got := gotList.getListParams()
	compare(test, "count", expList.count, got.Count)

	compPageURL(test, "previous", expList.previous, got.Previous)
	compPageURL(test, "next", expList.next, got.Next)

	listTest := newResListTestFromIDs("results", endpoint, expList.results, gotList.getResults())
	testResourceList(test, listTest)
}

// checks if all provided slices of resources contains all stated resources and also checks for their length, if stated
func testResourceLists(test test, checks []resListTest) {
	test.t.Helper()

	for _, c := range checks {
		testResourceList(test, c)
	}
}

// checks if the provided slice of resources contains all stated resources and also checks for its length, if stated
func testResourceList(test test, expList resListTest) {
	test.t.Helper()

	if len(expList.exp) == 0 {
		return
	}
	checkResourcesInSlice(test, expList.name, expList.exp, expList.got)

	expLen, ok := test.expLengths[expList.name]
	if !ok {
		return
	}

	compare(test, expList.name+" length", expLen, len(expList.got))
}

// checks if the provided slice of resources contains all stated resources
func checkResourcesInSlice[T HasAPIResource](test test, fieldName string, expectedPaths []string, gotRes []T) {
	gotMap := getResourceMap(gotRes)
	if len(gotMap) != len(gotRes) {
		test.t.Fatalf("%s: there appear to be duplicates in %s", test.name, fieldName)
	}

	for _, expPath := range expectedPaths {
		expURL := test.cfg.completeTestURL(expPath)
		_, ok := gotMap[expURL]
		if !ok {
			test.t.Fatalf("%s: %s doesn't contain all wanted resources. missing %s", test.name, fieldName, expURL)
		}
	}
}

// will need to use name-version pattern in monsterAmount's GetName() method for getResourceAmountMap() to work
// checks if stated ResourceAmount entries are in slices (used for baseStats, itemAmounts, monsterAmounts) and if their amount values match
func checkResAmtsInSlice[T ResourceAmount](test test, fieldName string, expAmounts map[string]int32, gotAmounts []T) {
	expLen, ok := test.expLengths[fieldName]
	if !ok {
		return
	}
	compare(test, fieldName+" length", expLen, len(gotAmounts))

	gotMap := getResourceAmountMap(gotAmounts)

	for key, exp := range expAmounts {
		got, ok := gotMap[key]
		if !ok {
			test.t.Fatalf("%s: %s doesn't contain resource %s", test.name, fieldName, key)
		}
		compare(test, key, exp, got)
	}
}
