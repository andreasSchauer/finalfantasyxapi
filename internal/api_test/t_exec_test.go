package apiTest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api"
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

func testServiceResponsesGet[G any](t *testing.T, tests []testGeneral, testFuncName string) {
	t.Helper()
	var zero G

	for i, tc := range tests {
		name := getTestName(testFuncName, tc.requestURL, i+1)

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			test, got, err := setupTest[G](t, tc, name, tc.handler)
			if errors.Is(err, errCorrect) {
				return
			}

			if reflect.DeepEqual(zero, got) {
				test.t.Errorf("expected %s to have a response body, but didn't get one.\n\n", name)
			}
		})
	}
}

func testServiceResponsesPost[E testCase, G any](t *testing.T, tests []E, testFuncName string, handlerFunc func(http.ResponseWriter, *http.Request), compFunc func(test, E, G)) {
	t.Helper()

	for i, exp := range tests {
		tc := exp.GetTestGeneral()
		name := getTestName(testFuncName, tc.requestURL, i+1)

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			test, got, err := setupTest[G](t, tc, name, handlerFunc)
			if errors.Is(err, errCorrect) {
				return
			}

			compFunc(test, exp, got)
		})
	}
}

// compareAPIResourceLists for normal API Resources
// compareSimpleResourceLists for Subsections
func testIdList[G any](t *testing.T, tests []expListIDs, endpoint api.EndpointName, testFuncName string, handlerFunc func(http.ResponseWriter, *http.Request), compFunc func(test, api.EndpointName, expListIDs, G)) {
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

// compareAPIResourceListPaths for api resources
func testPathList[G any](t *testing.T, tests []expListPaths, testFuncName string, handlerFunc func(http.ResponseWriter, *http.Request), compFunc func(test, expListPaths, G)) {
	t.Helper()
	for i, exp := range tests {
		tc := exp.testGeneral
		name := getTestName(testFuncName, tc.requestURL, i+1)

		expHandler := handlerFunc
		if handlerFunc == nil {
			expHandler = exp.handler
		}

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			test, got, err := setupTest[G](t, exp.testGeneral, name, expHandler)
			if errors.Is(err, errCorrect) {
				return
			}

			compFunc(test, exp, got)
		})
	}
}

// compareParameterLists for /parameters lists
// compareSectionLists for /sections lists
func testNameList[G any](t *testing.T, tests []expListNames, endpoint api.EndpointName, testFuncName string, handlerFunc func(http.ResponseWriter, *http.Request), compFunc func(test, api.EndpointName, expListNames, G)) {
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
func setupTest[G any](t *testing.T, tc testGeneral, testName string, handlerFunc func(http.ResponseWriter, *http.Request)) (test, G, error) {
	t.Helper()
	var zero G

	if tc.method == "" {
		tc.method = http.MethodGet
	}

	var body io.Reader
	if tc.method == http.MethodPost && tc.requestBody != nil {
		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(tc.requestBody)
		if err != nil {
			t.Fatalf("failed to encode test request body: %v", err)
		}
		body = &buf
	}

	req := httptest.NewRequest(tc.method, tc.requestURL, body)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	res := httptest.NewRecorder()

	handler := http.HandlerFunc(handlerFunc)
	handler.ServeHTTP(res, req)

	if res.Code != tc.expectedStatus {
		if res.Code >= 400 {
			t.Fatalf("expected status: %d, got %d, body=%s\n\n", tc.expectedStatus, res.Code, res.Body.String())
		}

		t.Fatalf("expected status: %d, got %d\n\n", tc.expectedStatus, res.Code)
	}

	if tc.expectedErr != "" {
		rawErr := res.Body.String()
		if !strings.Contains(rawErr, tc.expectedErr) {
			t.Fatalf("error mismatch.\n- want: %s\n- got %s\n\n", tc.expectedErr, rawErr)
		}
		return test{}, zero, errCorrect
	}

	var got G
	err := json.NewDecoder(res.Body).Decode(&got)
	if err != nil {
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
