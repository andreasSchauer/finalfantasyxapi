package api

import (
	"net/http"
	"testing"
)

func TestGetEnum(t *testing.T) {
	t.Parallel()
	tests := []testGeneral{
		{
			requestURL:     "/api/enums/1/",
			expectedStatus: http.StatusNotFound,
			expectedErr: 	"enum '1' not found.",
		},
		{
			requestURL:     "/api/enums/1/1",
			expectedStatus: http.StatusBadRequest,
			expectedErr: "wrong format. usage: '/api/enums', '/api/enums/{enum_name}'",
		},
		{
			requestURL:     "/api/enums/ability_type/1/",
			expectedStatus: http.StatusBadRequest,
			expectedErr: 	"wrong format. usage: '/api/enums', '/api/enums/{enum_name}'",
		},
		{
			requestURL:     "/api/enums/ability_type/1/a",
			expectedStatus: http.StatusBadRequest,
			expectedErr: 	"wrong format. usage: '/api/enums', '/api/enums/{enum_name}'",
		},
		{
			requestURL:     "/api/enums/ability_type?ass=1",
			expectedStatus: http.StatusBadRequest,
			expectedErr: 	"parameter 'ass' does not exist for endpoint /enums. use /api/enums/parameters for available parameters.",
		},
		{
			requestURL:     "/api/enums/ability_type?limit=1",
			expectedStatus: http.StatusBadRequest,
			expectedErr: 	"invalid usage of parameter 'limit'. parameter 'limit' can only be used with list-endpoints.",
		},
		{
			requestURL:     "/api/enums/ability_type",
			expectedStatus: http.StatusOK,
		},
	}

	testStatusses(t, tests, "GetEnum", testCfg.HandleEnums)
}

func TestRetrieveEnums(t *testing.T) {
	t.Parallel()
	tests := []expListPaths{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/enums/ability_type?ass=1",
				expectedStatus: http.StatusBadRequest,
				expectedErr: 	"parameter 'ass' does not exist for endpoint /enums. use /api/enums/parameters for available parameters.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/enums?limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   59,
			results: []string{
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

	testPathList(t, tests, "RetrieveEnums", testCfg.HandleEnums, compareAPIResourceListPaths[NamedApiResourceList])
}