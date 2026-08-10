package api

import (
	"net/http"
	"testing"
)

func TestRetrieveEndpoints(t *testing.T) {
	t.Parallel()
	tests := []expListPaths{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/endpoints/1",
				expectedStatus: http.StatusBadRequest,
				expectedErr: "wrong format. '/endpoints' doesn't support single-resource requests.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/endpoints/1/1",
				expectedStatus: http.StatusBadRequest,
				expectedErr: "wrong format. usage: '/api/endpoints'",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/endpoints",
				expectedStatus: http.StatusOK,
			},
			count:   51,
			results: []string{
				"/areas",
				"/equipment-tables",
				"/overdrive-commands",
				"/enums",
				"/stats",
				"/topmenus",
				"/treasures",
				"/key-items",
				"/endpoints",
			},
		},
	}

	testPathList(t, tests, "RetrieveEndpoints", testCfg.HandleEndpoints, compareAPIResourceListPaths[NamedApiResourceList])
}