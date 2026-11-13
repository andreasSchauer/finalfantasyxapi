package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
)

func TestRetrieveEndpoints(t *testing.T) {
	t.Parallel()
	tests := []ListPaths{
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/endpoints/1",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "wrong format. '/endpoints' doesn't support single-resource requests",
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/endpoints/1/1",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "wrong format. usage: '/api/endpoints'",
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/endpoints",
				expectedStatus: http.StatusOK,
			},
			count: len(apilookups.E.Endpoints.Slice),
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

	testPathList(t, tests, "RetrieveEndpoints", testSrv.HandleEndpoints, compareAPIResourceListPaths[apiresource.Named])
}
