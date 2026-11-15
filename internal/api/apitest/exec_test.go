package apitest

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apitesttypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/errs"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
)

type testCase[G any] interface {
	GetTestGeneral() apitesttypes.TestGeneral
	testType[G]
}

type testType[G any] interface {
	Compare(apitesttypes.Test, G)
}

func testSingleResources[E testCase[G], G any](t *testing.T, tests []E, testFuncName string, handlerFunc func(http.ResponseWriter, *http.Request)) {
	t.Helper()

	for i, exp := range tests {
		tc := exp.GetTestGeneral()
		name := tc.GetTestName(testFuncName, i+1)

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			testObj, got, err := setupTest[G](t, tc, name, handlerFunc)
			if errors.Is(err, errs.ErrCorrect) {
				return
			}

			exp.Compare(testObj, got)
		})
	}
}

func testStatusses(t *testing.T, tests []apitesttypes.TestGeneral, testFuncName string, handlerFunc func(http.ResponseWriter, *http.Request)) {
	t.Helper()
	for i, tc := range tests {
		name := tc.GetTestName(testFuncName, i+1)

		expHandler := handlerFunc
		if handlerFunc == nil {
			expHandler = tc.Handler
		}

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			_, _, err := setupTest[any](t, tc, name, expHandler)
			if errors.Is(err, errs.ErrCorrect) {
				return
			}
		})
	}
}

func testServiceResponsesGet[G any](t *testing.T, tests []apitesttypes.TestGeneral, testFuncName string) {
	t.Helper()
	var zero G

	for i, tc := range tests {
		name := tc.GetTestName(testFuncName, i+1)

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			test, got, err := setupTest[G](t, tc, name, tc.Handler)
			if errors.Is(err, errs.ErrCorrect) {
				return
			}

			if reflect.DeepEqual(zero, got) {
				test.T.Errorf("expected %s to have a response body, but didn't get one.\n\n", name)
			}
		})
	}
}

func testServiceResponsesPost[E testCase[G], G any](t *testing.T, tests []E, testFuncName string, handlerFunc func(http.ResponseWriter, *http.Request)) {
	t.Helper()

	for i, exp := range tests {
		tc := exp.GetTestGeneral()
		name := tc.GetTestName(testFuncName, i+1)

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			test, got, err := setupTest[G](t, tc, name, handlerFunc)
			if errors.Is(err, errs.ErrCorrect) {
				return
			}

			exp.Compare(test, got)
		})
	}
}

// compareAPIResourceLists for normal API Resources
// compareSimpleResourceLists for Subsections
func testIDList[G any](t *testing.T, tests []apitesttypes.ListIDs, endpoint ts.EndpointName, testFuncName string, handlerFunc func(http.ResponseWriter, *http.Request), compFunc func(apitesttypes.Test, ts.EndpointName, apitesttypes.ListIDs, G)) {
	t.Helper()
	for i, exp := range tests {
		tc := exp.GetTestGeneral()
		name := tc.GetTestName(testFuncName, i+1)

		expHandler := handlerFunc
		if handlerFunc == nil {
			expHandler = exp.Handler
		}

		expEndpoint := endpoint
		if endpoint == "" {
			expEndpoint = exp.Endpoint
		}

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			test, got, err := setupTest[G](t, exp.TestGeneral, name, expHandler)
			if errors.Is(err, errs.ErrCorrect) {
				return
			}

			compFunc(test, expEndpoint, exp, got)
		})
	}
}

// compareAPIResourceListPaths for api resources
func testPathList[G any](t *testing.T, tests []apitesttypes.ListPaths, testFuncName string, handlerFunc func(http.ResponseWriter, *http.Request), compFunc func(apitesttypes.Test, apitesttypes.ListPaths, G)) {
	t.Helper()
	for i, exp := range tests {
		tc := exp.TestGeneral
		name := tc.GetTestName(testFuncName, i+1)

		expHandler := handlerFunc
		if handlerFunc == nil {
			expHandler = exp.Handler
		}

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			test, got, err := setupTest[G](t, exp.TestGeneral, name, expHandler)
			if errors.Is(err, errs.ErrCorrect) {
				return
			}

			compFunc(test, exp, got)
		})
	}
}

// compareParameterLists for /parameters lists
// compareSectionLists for /sections lists
func testNameList[G any](t *testing.T, tests []apitesttypes.ListNames, testFuncName string, compFunc func(apitesttypes.Test, ts.EndpointName, apitesttypes.ListNames, G)) {
	t.Helper()
	for i, exp := range tests {
		tc := exp.TestGeneral
		name := tc.GetTestName(testFuncName, i+1)

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			test, got, err := setupTest[G](t, exp.TestGeneral, name, exp.Handler)
			if errors.Is(err, errs.ErrCorrect) {
				return
			}

			compFunc(test, exp.Endpoint, exp, got)
		})
	}
}

// makes the http request for the test and returns the result, as well as a test struct
func setupTest[G any](t *testing.T, tc apitesttypes.TestGeneral, testName string, handlerFunc func(http.ResponseWriter, *http.Request)) (apitesttypes.Test, G, error) {
	t.Helper()
	var zero G

	if tc.Method == "" {
		tc.Method = http.MethodGet
	}

	var body io.Reader
	if tc.Method == http.MethodPost && tc.RequestBody != nil {
		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(tc.RequestBody)
		if err != nil {
			t.Fatalf("failed to encode test request body: %v", err)
		}
		body = &buf
	}

	req := httptest.NewRequest(tc.Method, tc.RequestURL, body)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	res := httptest.NewRecorder()

	handler := http.HandlerFunc(handlerFunc)
	handler.ServeHTTP(res, req)

	if res.Code != tc.ExpectedStatus {
		if res.Code >= 400 {
			t.Fatalf("expected status: %d, got %d, body=%s\n\n", tc.ExpectedStatus, res.Code, res.Body.String())
		}

		t.Fatalf("expected status: %d, got %d\n\n", tc.ExpectedStatus, res.Code)
	}

	if tc.ExpectedErr != "" {
		rawErr := res.Body.String()
		if !strings.Contains(rawErr, tc.ExpectedErr) {
			t.Fatalf("error mismatch.\n- want: %s\n- got %s\n\n", tc.ExpectedErr, rawErr)
		}
		return apitesttypes.Test{}, zero, errs.ErrCorrect
	}

	var got G
	err := json.NewDecoder(res.Body).Decode(&got)
	if err != nil {
		t.Fatalf("failed to decode: %v\n\n", err)
	}

	test := apitesttypes.Test{
		T:          t,
		Srv:        &testSrv,
		Name:       testName,
		ExpLengths: tc.ExpLengths,
		DontCheck:  tc.DontCheck,
	}

	return test, got, nil
}
