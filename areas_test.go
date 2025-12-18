package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)


func TestGetArea(t *testing.T) {
	tests := []struct {
		testInOut
		expectedSingle
		expectedResourcesAreas
	}{
		{
			testInOut: testInOut{
				requestURL:     "/api/areas/145",
				expectedStatus: http.StatusOK,
			},
			expectedSingle: expectedSingle{
				id:                145,
				name:              "north",
				version:           h.GetInt32Ptr(1),
				expectedLen: map[string]int{
					"connected areas": 	2,
					"monsters":       	6,
					"formations":     	6,
				},
			},
			expectedResourcesAreas: expectedResourcesAreas{
				parentLocation:    "/locations/15",
				parentSublocation: "/sublocations/25",
				sidequest: h.GetStrPtr("/sidequests/6"),
				connectedAreas: []string{
					"/areas/144",
					"/areas/149",
				},
				monsters: []string{
					"/monsters/81",
					"/monsters/84",
					"/monsters/85",
				},
				formations: []string{
					"/monster-formations/204",
					"/monster-formations/208",
				},
			},
		},
	}

	for i, tc := range tests {
		testName := fmt.Sprintf("GetArea: %d, requestURL: %s", i+1, tc.requestURL)

		req := httptest.NewRequest(http.MethodGet, tc.requestURL, nil)
		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(testCfg.HandleAreas)
		handler.ServeHTTP(rr, req)

		if rr.Code != tc.expectedStatus {
			t.Fatalf("%s: expected %d, got %d, body=%s", testName, tc.expectedStatus, rr.Code, rr.Body.String())
		}

		if tc.expectedErr != "" {
			raw := rr.Body.String()
			if !strings.Contains(raw, tc.expectedErr) {
				t.Fatalf("%s: expected error message to contain %s, got %q", testName, tc.expectedErr, raw)
			}
			continue
		}

		var a Area
		if err := json.NewDecoder(rr.Body).Decode(&a); err != nil {
			t.Fatalf("%s: failed to decode: %v", testName, err)
		}

		if a.ID != tc.id {
			t.Fatalf("%s: expected id %d, got %d", testName, tc.id, a.ID)
		}

		if a.Name != tc.name {
			t.Fatalf("%s: expected name %s, got %s", testName, tc.name, a.Name)
		}

		if h.DerefOrNil(a.Version) != h.DerefOrNil(tc.version) {
			t.Fatalf("%s: expected version %d, got %d", testName, h.DerefOrNil(tc.version), h.DerefOrNil(a.Version))
		}

		if !resourcesMatch(testCfg, a.ParentLocation, tc.parentLocation) {
			t.Fatalf("%s: expected location %s, got %s", testName, tc.parentLocation, a.ParentLocation.URL)
		}

		if !resourcesMatch(testCfg, a.ParentSublocation, tc.parentSublocation) {
			t.Fatalf("%s: expected sublocation %s, got %s", testName, tc.parentSublocation, a.ParentSublocation)
		}

		if !resourcePtrsMatch(testCfg, a.Sidequest, tc.sidequest) {
			t.Fatalf("%s: expected sidequest %s, got %s", testName, h.DerefOrNil(tc.sidequest), derefResourcePtr(a.Sidequest).ToKeyFields())
		}

		checks := []testCheck{
			{name: "connected areas", got: toIface(a.ConnectedAreas), 			expected: tc.connectedAreas},
			{name: "characters",      got: toIface(a.Characters),      			expected: tc.characters},
			{name: "aeons",           got: toIface(a.Aeons),           			expected: tc.aeons},
			{name: "shops",           got: toIface(a.Shops),           			expected: tc.shops},
			{name: "treasures",       got: toIface(a.Treasures),       			expected: tc.treasures},
			{name: "monsters",        got: toIface(a.Monsters),        			expected: tc.monsters},
			{name: "formations",      got: toIface(a.Formations),      			expected: tc.formations},
			{name: "bg music",        got: toIface(a.Music.BackgroundMusic), 	expected: tc.bgMusic},
			{name: "cues music",      got: toIface(a.Music.Cues),          		expected: tc.cuesMusic},
			{name: "fmvs music",      got: toIface(a.Music.FMVs),          		expected: tc.fmvsMusic},
			{name: "boss music",      got: toIface(a.Music.BossFights),     	expected: tc.bossMusic},
			{name: "fmvs",            got: toIface(a.FMVs),                 	expected: tc.fmvs},
		}

		testResponseChecks(t, testCfg, testName, checks, tc.expectedLen)
	}
}