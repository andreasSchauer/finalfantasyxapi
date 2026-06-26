package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type testCase interface {
	GetTestGeneral() testGeneral
}

func getTestName(name, requestURL string, caseNum int) string {
	return fmt.Sprintf("%s-%d: %s", name, caseNum, requestURL)
}

func testSingleResources[E testCase, G any](t *testing.T, tests []E, testFuncName string, handlerFunc func(http.ResponseWriter, *http.Request), compFunc func(test, E, G)) {
	t.Helper()

	for i, exp := range tests {
		tc := exp.GetTestGeneral()
		name := getTestName(testFuncName, tc.requestURL, i+1)

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			testObj, got, err := setupTest[G](t, tc, name, handlerFunc)
			if errors.Is(err, errCorrect) {
				return
			}

			compFunc(testObj, exp, got)
		})
	}
}

func testStatusses(t *testing.T, tests []testGeneral, testFuncName string, handlerFunc func(http.ResponseWriter, *http.Request)) {
	t.Helper()
	for i, exp := range tests {
		name := getTestName(testFuncName, exp.requestURL, i+1)

		expHandler := handlerFunc
		if handlerFunc == nil {
			expHandler = exp.handler
		}

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			_, _, err := setupTest[any](t, exp, name, expHandler)
			if errors.Is(err, errCorrect) {
				return
			}
		})
	}
}


// compareAPIResourceLists for normal API Resources
// compareSimpleResourceLists for Subsections
func testIdList[G any](t *testing.T, tests []expListIDs, endpoint EndpointName, testFuncName string, handlerFunc func(http.ResponseWriter, *http.Request), compFunc func(test, EndpointName, expListIDs, G)) {
	t.Helper()
	for i, exp := range tests {
		tc := exp.testGeneral
		name := getTestName(testFuncName, tc.requestURL, i+1)

		expHandler := handlerFunc
		if handlerFunc == nil {
			expHandler = exp.handler
		}

		expEndpoint := endpoint
		if endpoint == "" {
			expEndpoint = exp.endpoint
		}

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			test, got, err := setupTest[G](t, exp.testGeneral, name, expHandler)
			if errors.Is(err, errCorrect) {
				return
			}

			compFunc(test, expEndpoint, exp, got)
		})
	}
}

// compareParameterLists for /parameters lists
// compareSectionLists for /sections lists
func testNameList[G any](t *testing.T, tests []expListNames, endpoint EndpointName, testFuncName string, handlerFunc func(http.ResponseWriter, *http.Request), compFunc func(test, EndpointName, expListNames, G)) {
	t.Helper()
	for i, exp := range tests {
		tc := exp.testGeneral
		name := getTestName(testFuncName, tc.requestURL, i+1)

		expHandler := handlerFunc
		if handlerFunc == nil {
			expHandler = exp.handler
		}

		expEndpoint := endpoint
		if endpoint == "" {
			expEndpoint = exp.endpoint
		}

		t.Run(name, func(t *testing.T) {
			t.Parallel()
			
			test, got, err := setupTest[G](t, exp.testGeneral, name, expHandler)
			if errors.Is(err, errCorrect) {
				return
			}
			
			compFunc(test, expEndpoint, exp, got)
		})
	}
}



// makes the http request for the test and returns the result, as well as a test struct
func setupTest[T any](t *testing.T, tc testGeneral, testName string, handlerFunc func(http.ResponseWriter, *http.Request)) (test, T, error) {
	t.Helper()
	var zeroType T

	req := httptest.NewRequest(http.MethodGet, tc.requestURL, nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(handlerFunc)
	handler.ServeHTTP(rr, req)

	if rr.Code != tc.expectedStatus {
		if rr.Code >= 400 {
			t.Fatalf("expected status: %d, got %d, body=%s\n\n", tc.expectedStatus, rr.Code, rr.Body.String())
		}
			
		t.Fatalf("expected status: %d, got %d\n\n", tc.expectedStatus, rr.Code,)
	}

	if tc.expectedErr != "" {
		rawErr := rr.Body.String()
		if !strings.Contains(rawErr, tc.expectedErr) {
			t.Fatalf("error mismatch.\n- want: %s\n- got %s\n\n", tc.expectedErr, rawErr)
		}
		return test{}, zeroType, errCorrect
	}

	var got T
	if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
		t.Fatalf("failed to decode: %v\n\n", err)
	}

	test := test{
		t:          t,
		cfg:        testCfg,
		name:       testName,
		expLengths: tc.expLengths,
		dontCheck:  tc.dontCheck,
	}

	return test, got, nil
}
