package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/service/servicetypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/tests/apitesttypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
)

func TestAlBhed(t *testing.T) {
	t.Parallel()
	requestURL := "/api/al-bhed"

	tests := []apitesttypes.AlBhedResponse{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     requestURL,
				Method:         http.MethodPost,
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "invalid or corrupt payload in request body",
				RequestBody:    []byte(string("asjdasoi")),
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     requestURL,
				Method:         http.MethodPost,
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "field 'text' can't be empty",
				RequestBody: map[ts.FieldName]any{
					ts.PfnDirection: "to-english",
				},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     requestURL,
				Method:         http.MethodPost,
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "value 'bla' of field 'direction' is not a valid enum value. allowed values: 'to-al-bhed', 'to-english'",
				RequestBody: map[ts.FieldName]any{
					ts.PfnText:      "father",
					ts.PfnDirection: "bla",
				},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     requestURL,
				Method:         http.MethodPost,
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "wrong format. all brackets must be closed",
				RequestBody: map[ts.FieldName]any{
					ts.PfnText: "v[ath",
				},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     requestURL,
				Method:         http.MethodPost,
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "wrong format. nested brackets are not allowed, and the first bracket must be an opening bracket",
				RequestBody: map[ts.FieldName]any{
					ts.PfnText: "v]ath",
				},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     requestURL,
				Method:         http.MethodPost,
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "wrong format. nested brackets are not allowed, and the first bracket must be an opening bracket",
				RequestBody: map[ts.FieldName]any{
					ts.PfnText: "v[a[th",
				},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     requestURL,
				Method:         http.MethodPost,
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "wrong format. all brackets must be closed",
				RequestBody: map[ts.FieldName]any{
					ts.PfnText: "v[ath",
				},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     requestURL,
				Method:         http.MethodPost,
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "invalid field: 'monster_id'",
				RequestBody: map[ts.FieldName]any{
					ts.PfnMonsterID: 1,
				},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     requestURL,
				Method:         http.MethodPost,
				ExpectedStatus: http.StatusOK,
				RequestBody: map[ts.FieldName]any{
					ts.PfnText: "I am cool",
				},
			},
			AlBhedResponse: servicetypes.AlBhedResponse{
				TranslatedText: "E ys luum",
				OriginalText:   "I am cool",
				Direction:      "to-al-bhed",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     requestURL,
				Method:         http.MethodPost,
				ExpectedStatus: http.StatusOK,
				RequestBody: map[ts.FieldName]any{
					ts.PfnText: "father",
				},
			},
			AlBhedResponse: servicetypes.AlBhedResponse{
				TranslatedText: "vydran",
				OriginalText:   "father",
				Direction:      "to-al-bhed",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     requestURL,
				Method:         http.MethodPost,
				ExpectedStatus: http.StatusOK,
				RequestBody: map[ts.FieldName]any{
					ts.PfnText: "f[y]the[n]",
				},
			},
			AlBhedResponse: servicetypes.AlBhedResponse{
				TranslatedText: "vydran",
				OriginalText:   "f[y]the[n]",
				Direction:      "to-al-bhed",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     requestURL,
				Method:         http.MethodPost,
				ExpectedStatus: http.StatusOK,
				RequestBody: map[ts.FieldName]any{
					ts.PfnText: "f[Ydr]er",
				},
			},
			AlBhedResponse: servicetypes.AlBhedResponse{
				TranslatedText: "vYdran",
				OriginalText:   "f[Ydr]er",
				Direction:      "to-al-bhed",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     requestURL,
				Method:         http.MethodPost,
				ExpectedStatus: http.StatusOK,
				RequestBody: map[ts.FieldName]any{
					ts.PfnText:      "v[ath]an",
					ts.PfnDirection: "to-english",
				},
			},
			AlBhedResponse: servicetypes.AlBhedResponse{
				TranslatedText: "father",
				OriginalText:   "v[ath]an",
				Direction:      "to-english",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     requestURL,
				Method:         http.MethodPost,
				ExpectedStatus: http.StatusOK,
				RequestBody: map[ts.FieldName]any{
					ts.PfnText: "!?,. 1234567890",
				},
			},
			AlBhedResponse: servicetypes.AlBhedResponse{
				TranslatedText: "!?,. 1234567890",
				OriginalText:   "!?,. 1234567890",
				Direction:      "to-al-bhed",
			},
		},
	}

	testServiceResponsesPost(t, tests, "AlBhed", testSrv.HandleAlBhedPost)
}
