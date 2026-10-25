package api

import (
	"net/http"
	"testing"
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
				expectedErr:    "invalid or corrupt payload in request body.",
				requestBody:    []byte(string("asjdasoi")),
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "field 'text' can't be empty.",
				requestBody: map[FieldName]any{
					pfnDirection: "to-english",
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "value 'bla' of field 'direction' is not a valid enum value. allowed values: 'to-al-bhed', 'to-english'.",
				requestBody: map[FieldName]any{
					pfnText:      "father",
					pfnDirection: "bla",
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "wrong format. all brackets must be closed.",
				requestBody: map[FieldName]any{
					pfnText: "v[ath",
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "wrong format. nested brackets are not allowed, and the first bracket must be an opening bracket.",
				requestBody: map[FieldName]any{
					pfnText: "v]ath",
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "wrong format. nested brackets are not allowed, and the first bracket must be an opening bracket.",
				requestBody: map[FieldName]any{
					pfnText: "v[a[th",
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "wrong format. all brackets must be closed.",
				requestBody: map[FieldName]any{
					pfnText: "v[ath",
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid field: 'monster_id'.",
				requestBody: map[FieldName]any{
					pfnMonsterID: 1,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				requestBody: map[FieldName]any{
					pfnText: "I am cool.",
				},
			},
			AlBhedResponse: AlBhedResponse{
				TranslatedText: "E ys luum.",
				OriginalText:   "I am cool.",
				Direction:      "to-al-bhed",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     requestURL,
				method:         http.MethodPost,
				expectedStatus: http.StatusOK,
				requestBody: map[FieldName]any{
					pfnText: "father",
				},
			},
			AlBhedResponse: AlBhedResponse{
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
				requestBody: map[FieldName]any{
					pfnText: "f[y]the[n]",
				},
			},
			AlBhedResponse: AlBhedResponse{
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
				requestBody: map[FieldName]any{
					pfnText: "f[Ydr]er",
				},
			},
			AlBhedResponse: AlBhedResponse{
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
				requestBody: map[FieldName]any{
					pfnText:      "v[ath]an",
					pfnDirection: "to-english",
				},
			},
			AlBhedResponse: AlBhedResponse{
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
				requestBody: map[FieldName]any{
					pfnText: "!?,. 1234567890",
				},
			},
			AlBhedResponse: AlBhedResponse{
				TranslatedText: "!?,. 1234567890",
				OriginalText:   "!?,. 1234567890",
				Direction:      "to-al-bhed",
			},
		},
	}

	testServiceResponsesPost(t, tests, "AlBhed", testCfg.HandleAlBhedPost, compareAlBhedResponses)
}
