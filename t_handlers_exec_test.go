package main

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
	return fmt.Sprintf("%s: %d, requestURL: %s", name, caseNum, requestURL)
}


func testSingleResources[E testCase, G any](t *testing.T, tests []E, testFuncName string, handlerFunc func(http.ResponseWriter, *http.Request), compFunc func(test, E, G)) {
	for i, exp := range tests {
		test, got, err := setupTest[G](t, exp.GetTestGeneral(), testFuncName, i+1, handlerFunc)
		if errors.Is(err, errCorrect) {
			continue
		}

		compFunc(test, exp, got)
	}
}

// compareAPIResourceLists for normal API Resources
// compareSubResourceLists for Subsections
func testIdList[T any](t *testing.T, tests []expListIDs, endpoint, testFuncName string, handlerFunc func(http.ResponseWriter, *http.Request), compFunc func(test, string, expListIDs, T)) {
	for i, exp := range tests {
		expHandler := handlerFunc
		if handlerFunc == nil {
			expHandler = exp.handler
		}

		test, got, err := setupTest[T](t, exp.testGeneral, testFuncName, i+1, expHandler)
		if errors.Is(err, errCorrect) {
			continue
		}

		compFunc(test, endpoint, exp, got)
	}
}

// compareParameterLists for /parameters lists
// compareSectionLists for /sections lists
func testNameList[T any](t *testing.T, tests []expListNames, endpoint, testFuncName string, handlerFunc func(http.ResponseWriter, *http.Request), compFunc func(test, string, expListNames, T)) {
	for i, exp := range tests {
		expHandler := handlerFunc
		if handlerFunc == nil {
			expHandler = exp.handler
		}

		expEndpoint := endpoint
		if endpoint == "" {
			expEndpoint = exp.endpoint
		}

		test, got, err := setupTest[T](t, exp.testGeneral, testFuncName, i+1, expHandler)
		if errors.Is(err, errCorrect) {
			continue
		}

		compFunc(test, expEndpoint, exp, got)
	}
}


// makes the http request for the test and returns the result, as well as a test struct
func setupTest[T any](t *testing.T, tc testGeneral, testFunc string, testNum int, handlerFunc func(http.ResponseWriter, *http.Request)) (test, T, error) {
	t.Helper()
	testName := getTestName(testFunc, tc.requestURL, testNum)
	var zeroType T

	req := httptest.NewRequest(http.MethodGet, tc.requestURL, nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(handlerFunc)
	handler.ServeHTTP(rr, req)

	if rr.Code != tc.expectedStatus {
		t.Fatalf("%s: expected %d, got %d, body=%s", testName, tc.expectedStatus, rr.Code, rr.Body.String())
	}

	if tc.expectedErr != "" {
		rawErr := rr.Body.String()
		if !strings.Contains(rawErr, tc.expectedErr) {
			t.Fatalf("%s: expected error message to contain %s, got %q", testName, tc.expectedErr, rawErr)
		}
		return test{}, zeroType, errCorrect
	}

	var got T
	if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
		t.Fatalf("%s: failed to decode: %v", testName, err)
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