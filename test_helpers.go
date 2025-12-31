package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type resListTest struct {
	name string
	exp  []string
	got  []HasAPIResource
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
func testExpectedUnique(t *testing.T, testName string, tc expUnique, gotID int32, gotName string) {
	t.Helper()
	compare(t, testName, "id", tc.id, gotID, nil)
	compare(t, testName, "name", tc.name, gotName, nil)
}

// checks, if all basic fields of a resource that uses a name/version pattern are equal
func testExpectedNameVer(t *testing.T, testName string, tc expNameVer, gotID int32, gotName string, gotVer *int32) {
	t.Helper()
	compare(t, testName, "id", tc.id, gotID, nil)
	compare(t, testName, "name", tc.name, gotName, nil)
	compare(t, testName, "version", tc.version, gotVer, nil)
}

// checks the basic fields of an APIResourceList (count, pagination urls) and then checks for the stated resources
func testAPIResourceList[T IsAPIResourceList](t *testing.T, testCfg *Config, testName string, expList expList, gotList T, dontCheck map[string]bool) {
	t.Helper()
	got := gotList.getListParams()
	compare(t, testName, "count", expList.count, got.Count, nil)

	compPageURL(t, testCfg, testName, "previous", expList.previous, got.Previous, dontCheck)
	compPageURL(t, testCfg, testName, "next", expList.next, got.Next, dontCheck)

	listTest := resListTest{
		name: "results",
		exp:  expList.results,
		got:  gotList.getResults(),
	}

	testResourceList(t, testCfg, testName, listTest, nil)
}

// checks if all provided slices of resources contains all stated resources and also checks for their length, if stated
func testResourceLists(t *testing.T, testCfg *Config, testName string, checks []resListTest, expLengths map[string]int) {
	t.Helper()

	for _, c := range checks {
		testResourceList(t, testCfg, testName, c, expLengths)
	}
}

// checks if the provided slice of resources contains all stated resources and also checks for its length, if stated
func testResourceList(t *testing.T, testCfg *Config, testName string, expList resListTest, expLengths map[string]int) {
	t.Helper()

	if len(expList.exp) == 0 {
		return
	}
	checkResourcesInSlice(t, testCfg, testName, expList.name, expList.exp, expList.got)

	expLen, ok := expLengths[expList.name]
	if !ok {
		return
	}

	compare(t, testName, expList.name+" length", expLen, len(expList.got), nil)
}

// checks if stated resources are in resource slice
func checkResourcesInSlice[T HasAPIResource](t *testing.T, cfg *Config, testName, fieldName string, expectedPaths []string, gotRes []T) {
	gotMap := getResourceMap(gotRes)
	if len(gotMap) != len(gotRes) {
		t.Fatalf("%s: there appear to be duplicates in %s", testName, fieldName)
	}

	for _, expPath := range expectedPaths {
		expURL := cfg.completeURL(expPath)
		_, ok := gotMap[expURL]
		if !ok {
			t.Fatalf("%s: %s doesn't contain all wanted resources. missing %s", testName, fieldName, expURL)
		}
	}
}

// will need to use name-version pattern in GetName() for monsterAmount
// checks if stated ResourceAmount entries are in slices (used for baseStats, itemAmounts, monsterAmounts) and if their amount values match
func checkResAmtsInSlice[T ResourceAmount](t *testing.T, testName, fieldName string, expAmounts map[string]int32, gotAmounts []T, expLengths map[string]int) {
	expLen, ok := expLengths[fieldName]
	if !ok {
		return
	}
	compare(t, testName, fieldName+" length", expLen, len(gotAmounts), nil)

	gotMap := getResourceAmountMap(gotAmounts)

	for key, exp := range expAmounts {
		got, ok := gotMap[key]
		if !ok {
			t.Fatalf("%s: %s doesn't contain resource %s", testName, fieldName, key)
		}
		compare(t, testName, key, exp, got, nil)
	}
}
