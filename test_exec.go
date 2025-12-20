package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func setupTest(t *testing.T, tc testInOut, testFunc string, testNum int, handlerFunc func(http.ResponseWriter, *http.Request)) (*httptest.ResponseRecorder, string, bool) {
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

func testExpectedUnique(t *testing.T, testName string, tc expectedUnique, gotID int32, gotName string) {
	if gotID != tc.id {
		t.Fatalf("%s: expected id %d, got %d", testName, tc.id, gotID)
	}

	if gotName != tc.name {
		t.Fatalf("%s: expected name %s, got %s", testName, tc.name, gotName)
	}
}

func testExpectedNameVer(t *testing.T, testName string, tc expectedNameVer, gotID int32, gotName string, gotVer *int32) {
	if gotID != tc.id {
		t.Fatalf("%s: expected id %d, got %d", testName, tc.id, gotID)
	}

	if gotName != tc.name {
		t.Fatalf("%s: expected name %s, got %s", testName, tc.name, gotName)
	}

	if h.DerefOrNil(gotVer) != h.DerefOrNil(tc.version) {
		t.Fatalf("%s: expected version %d, got %d", testName, h.DerefOrNil(tc.version), h.DerefOrNil(gotVer))
	}
}

func testAPIResourceList[T IsAPIResourceList](t *testing.T, testCfg *Config, testName string, resList T, tc expectedList) {
	res := resList.getListParams()
	if res.Count != tc.count {
		t.Fatalf("%s: expected count %d, got %d", testName, tc.count, res.Count)
	}

	testPaginationURLs(t, testCfg, testName, "previous url", res.Previous, tc.previous)
	testPaginationURLs(t, testCfg, testName, "next url", res.Next, tc.next)

	checks := []testCheck{
		{
			name:     "results",
			got:      resList.getResults(),
			expected: tc.results,
		},
	}

	testResponseChecks(t, testCfg, testName, checks, nil)
}

func testResponseChecks(t *testing.T, testCfg *Config, testName string, checks []testCheck, lenMap map[string]int) {
	for _, c := range checks {
		if len(c.expected) == 0 {
			continue
		}
		if !containsAllResources(testCfg, c.got, c.expected) {
			t.Fatalf("%s: %s doesn't contain all wanted resources of %v", testName, c.name, c.expected)
		}

		expLen, ok := lenMap[c.name]
		if !ok {
			continue
		}

		if !hasExpectedLength(c.got, expLen) {
			t.Fatalf("%s: %s expected length: %d, got: %d", testName, c.name, expLen, len(c.got))
		}
	}
}