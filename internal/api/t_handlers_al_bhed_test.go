package api

import (
	"net/http"
	"testing"
)

func TestAlBhed(t *testing.T) {
	t.Parallel()
	tests := []expAlBhedResponse{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/al-bhed",
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid or corrupt payload in request body.",
				requestBody: 	[]byte(string("asjdasoi")),
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/al-bhed",
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "field 'text' can't be empty.",
				requestBody: AlBhedParams{
					Direction: "to-english",
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/al-bhed",
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "value 'bla' of field 'direction' is not a valid enum value. allowed values: 'to-al-bhed', 'to-english'.",
				requestBody: AlBhedParams{
					Text: "father",
					Direction: "bla",
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/al-bhed",
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "wrong format. all brackets must be closed.",
				requestBody: AlBhedParams{
					Text: "v[ath",
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/al-bhed",
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "wrong format. nested brackets are not allowed, and the first bracket must be an opening bracket.",
				requestBody: AlBhedParams{
					Text: "v]ath",
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/al-bhed",
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "wrong format. nested brackets are not allowed, and the first bracket must be an opening bracket.",
				requestBody: AlBhedParams{
					Text: "v[a[th",
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/al-bhed",
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "wrong format. all brackets must be closed.",
				requestBody: AlBhedParams{
					Text: "v[ath",
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/al-bhed",
				method: 		http.MethodPost,
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid field: 'monster'.",
				requestBody: DropChanceParams{
					Monster: 1,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/al-bhed",
				method: 		http.MethodPost,
				expectedStatus: http.StatusOK,
				requestBody: AlBhedParams{
					Text: "I am cool.",
				},
			},
			AlBhedResponse: AlBhedResponse{
				TranslatedText: "E ys luum.",
				OriginalText: "I am cool.",
				Direction: "to-al-bhed",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/al-bhed",
				method: 		http.MethodPost,
				expectedStatus: http.StatusOK,
				requestBody: AlBhedParams{
					Text: "father",
				},
			},
			AlBhedResponse: AlBhedResponse{
				TranslatedText: "vydran",
				OriginalText: "father",
				Direction: "to-al-bhed",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/al-bhed",
				method: 		http.MethodPost,
				expectedStatus: http.StatusOK,
				requestBody: AlBhedParams{
					Text: "f[y]the[n]",
				},
			},
			AlBhedResponse: AlBhedResponse{
				TranslatedText: "vydran",
				OriginalText: "f[y]the[n]",
				Direction: "to-al-bhed",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/al-bhed",
				method: 		http.MethodPost,
				expectedStatus: http.StatusOK,
				requestBody: AlBhedParams{
					Text: "f[Ydr]er",
				},
			},
			AlBhedResponse: AlBhedResponse{
				TranslatedText: "vYdran",
				OriginalText: "f[Ydr]er",
				Direction: "to-al-bhed",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/al-bhed",
				method: 		http.MethodPost,
				expectedStatus: http.StatusOK,
				requestBody: AlBhedParams{
					Text: "v[ath]an",
					Direction: "to-english",
				},
			},
			AlBhedResponse: AlBhedResponse{
				TranslatedText: "father",
				OriginalText: "v[ath]an",
				Direction: "to-english",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/al-bhed",
				method: 		http.MethodPost,
				expectedStatus: http.StatusOK,
				requestBody: AlBhedParams{
					Text: "!?,. 1234567890",
				},
			},
			AlBhedResponse: AlBhedResponse{
				TranslatedText: "!?,. 1234567890",
				OriginalText: "!?,. 1234567890",
				Direction: "to-al-bhed",
			},
		},
	}

	testServiceResponses(t, tests, "AlBhed", testCfg.HandleAlBhedGet, testCfg.HandleAlBhedPost, compareAlBhedResponses)
}