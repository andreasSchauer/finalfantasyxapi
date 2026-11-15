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
				RequestURL:     "/api/endpoints/1",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "wrong format. '/endpoints' doesn't support single-resource requests",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/endpoints/1/1",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "wrong format. usage: '/api/endpoints'",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/endpoints",
				ExpectedStatus: http.StatusOK,
			},
			Count: len(apilookups.E.Endpoints.Slice),
			Results: []string{
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
