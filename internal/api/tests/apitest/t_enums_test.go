package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/tests/apitesttypes"
)

func TestGetEnum(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.TestGeneral{
		{
			RequestURL:     "/api/enums/1/",
			ExpectedStatus: http.StatusNotFound,
			ExpectedErr:    "enum '1' not found",
		},
		{
			RequestURL:     "/api/enums/1/1",
			ExpectedStatus: http.StatusBadRequest,
			ExpectedErr:    "wrong format. usage: '/api/enums', '/api/enums/{enum_name}'",
		},
		{
			RequestURL:     "/api/enums/ability_type/1/",
			ExpectedStatus: http.StatusBadRequest,
			ExpectedErr:    "wrong format. usage: '/api/enums', '/api/enums/{enum_name}'",
		},
		{
			RequestURL:     "/api/enums/ability_type/1/a",
			ExpectedStatus: http.StatusBadRequest,
			ExpectedErr:    "wrong format. usage: '/api/enums', '/api/enums/{enum_name}'",
		},
		{
			RequestURL:     "/api/enums/ability_type?ass=1",
			ExpectedStatus: http.StatusBadRequest,
			ExpectedErr:    "parameter 'ass' does not exist for endpoint /enums. use /api/enums/parameters for available parameters",
		},
		{
			RequestURL:     "/api/enums/ability_type?limit=1",
			ExpectedStatus: http.StatusBadRequest,
			ExpectedErr:    "invalid usage of parameter 'limit'. parameter 'limit' can only be used with list-endpoints",
		},
		{
			RequestURL:     "/api/enums/ability_type",
			ExpectedStatus: http.StatusOK,
		},
	}

	testStatusses(t, tests, "GetEnum", testSrv.HandleEnums)
}

func TestRetrieveEnums(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.ListPaths{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/enums/ability_type?ass=1",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "parameter 'ass' does not exist for endpoint /enums. use /api/enums/parameters for available parameters",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/enums?limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count: 62,
			Results: []string{
				"/enums/ability_type",
				"/enums/blitzball_position_slot",
				"/enums/damage_formula",
				"/enums/loot_type",
				"/enums/node_state",
				"/enums/shop_category",
				"/enums/target_type",
				"/enums/weapon_type",
			},
		},
	}

	testPathList(t, tests, "RetrieveEnums", testSrv.HandleEnums, compareAPIResourceListPaths[apiresource.Named])
}
