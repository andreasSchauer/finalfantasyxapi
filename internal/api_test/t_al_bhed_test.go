package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api"
	"github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func TestAlBhed(t *testing.T) {
	t.Parallel()
	requestURL := "/api/al-bhed"

	tests := []expAlBhedResponse{
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid or corrupt payload in request body",
				requestBody:    []byte(string("asjdasoi")),
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "field 'text' can't be empty",
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnDirection: "to-english",
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "value 'bla' of field 'direction' is not a valid enum value. allowed values: 'to-al-bhed', 'to-english'",
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnText:      "father",
					typestrings.PfnDirection: "bla",
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "wrong format. all brackets must be closed",
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnText: "v[ath",
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "wrong format. nested brackets are not allowed, and the first bracket must be an opening bracket",
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnText: "v]ath",
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "wrong format. nested brackets are not allowed, and the first bracket must be an opening bracket",
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnText: "v[a[th",
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "wrong format. all brackets must be closed",
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnText: "v[ath",
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid field: 'monster_id'",
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnMonsterID: 1,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnText: "I am cool",
				},
			},
			AlBhedResponse: api.AlBhedResponse{
				TranslatedText: "E ys luum",
				OriginalText:   "I am cool",
				Direction:      "to-al-bhed",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnText: "father",
				},
			},
			AlBhedResponse: api.AlBhedResponse{
				TranslatedText: "vydran",
				OriginalText:   "father",
				Direction:      "to-al-bhed",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnText: "f[y]the[n]",
				},
			},
			AlBhedResponse: api.AlBhedResponse{
				TranslatedText: "vydran",
				OriginalText:   "f[y]the[n]",
				Direction:      "to-al-bhed",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnText: "f[Ydr]er",
				},
			},
			AlBhedResponse: api.AlBhedResponse{
				TranslatedText: "vYdran",
				OriginalText:   "f[Ydr]er",
				Direction:      "to-al-bhed",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnText:      "v[ath]an",
					typestrings.PfnDirection: "to-english",
				},
			},
			AlBhedResponse: api.AlBhedResponse{
				TranslatedText: "father",
				OriginalText:   "v[ath]an",
				Direction:      "to-english",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				requestBody: map[typestrings.FieldName]any{
					typestrings.PfnText: "!?,. 1234567890",
				},
			},
			AlBhedResponse: api.AlBhedResponse{
				TranslatedText: "!?,. 1234567890",
				OriginalText:   "!?,. 1234567890",
				Direction:      "to-al-bhed",
			},
		},
	}

	testServiceResponsesPost(t, tests, "AlBhed", testCfg.HandleAlBhedPost, compareAlBhedResponses)
}
