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

func getTestName(name, requestURL string, caseNum int) string {
	return fmt.Sprintf("%s: %d, requestURL: %s", name, caseNum, requestURL)
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


func testIdList[T any](t *testing.T, tests []expListIDs, endpoint, testFuncName string, handlerFunc func(http.ResponseWriter, *http.Request), compFunc func(test, string, expListIDs, T)) {
	for i, tc := range tests {
		test, got, err := setupTest[T](t, tc.testGeneral, testFuncName, i+1, handlerFunc)
		if errors.Is(err, errCorrect) {
			continue
		}

		compFunc(test, endpoint, tc, got)
	}
}


func testNameList[T any](t *testing.T, tests []expListNames, endpoint, testFuncName string, handlerFunc func(http.ResponseWriter, *http.Request), compFunc func(test, string, expListNames, T)) {
	for i, tc := range tests {
		test, got, err := setupTest[T](t, tc.testGeneral, testFuncName, i+1, handlerFunc)
		if errors.Is(err, errCorrect) {
			continue
		}

		compFunc(test, endpoint, tc, got)
	}
}