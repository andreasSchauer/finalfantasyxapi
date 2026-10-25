package api

import (
	"net/http"
	"testing"
)

func TestServiceBody(t *testing.T) {
	t.Parallel()

	tests := []testGeneral{
		{
			requestURL:     "/api/al-bhed/body",
			expectedStatus: http.StatusOK,
			handler: 		testCfg.HandleAlBhedGet,
		},
		{
			requestURL:     "/api/turn-order/body",
			expectedStatus: http.StatusOK,
			handler: 		testCfg.HandleTurnOrderGet,
		},
		{
			requestURL:     "/api/drop-chance/body",
			expectedStatus: http.StatusOK,
			handler: 		testCfg.HandleDropChanceGet,
		},
		{
			requestURL:     "/api/al-bhed/",
			expectedStatus: http.StatusBadRequest,
			expectedErr: 	"wrong format. usage: 'GET /api/al-bhed/body', 'POST /api/al-bhed'.",
			handler: 		testCfg.HandleAlBhedGet,
		},
		{
			requestURL:     "/api/al-bhed/body/2/2",
			expectedStatus: http.StatusBadRequest,
			expectedErr: 	"wrong format. usage: 'GET /api/al-bhed/body', 'POST /api/al-bhed'.",
			handler: 		testCfg.HandleAlBhedGet,
		},
		{
			requestURL:     "/api/al-bhed/body/asd",
			expectedStatus: http.StatusBadRequest,
			expectedErr: 	"'GET api/al-bhed/body' only supports the following sections: 'docs', 'samples'.",
			handler: 		testCfg.HandleAlBhedGet,
		},
		{
			requestURL:     "/api/al-bhed/bodsa",
			expectedStatus: http.StatusBadRequest,
			expectedErr: 	"'GET api/al-bhed' only supports '/body'.",
			handler: 		testCfg.HandleAlBhedGet,
		},
		{
			requestURL:     "/api/al-bhed/body?limit=23",
			expectedStatus: http.StatusBadRequest,
			expectedErr: 	"computational endpoints don't have any query parameters.",
			handler: 		testCfg.HandleAlBhedGet,
		},
	}

	testServiceResponsesGet[Folder](t, tests, "ServiceBody")
}

func TestServiceBodyDocs(t *testing.T) {
	t.Parallel()

	tests := []testGeneral{
		{
			requestURL:     "/api/al-bhed/body/docs",
			expectedStatus: http.StatusOK,
			handler: 		testCfg.HandleAlBhedGet,
		},
		{
			requestURL:     "/api/turn-order/body/docs",
			expectedStatus: http.StatusOK,
			handler: 		testCfg.HandleTurnOrderGet,
		},
		{
			requestURL:     "/api/drop-chance/body/docs",
			expectedStatus: http.StatusOK,
			handler: 		testCfg.HandleDropChanceGet,
		},
	}

	testServiceResponsesGet[ParamsDoc](t, tests, "ServiceBodyDocs")
}

func TestAlBhedBodySamples(t *testing.T) {
	t.Parallel()

	tests := []testGeneral{
		{
			requestURL:     "/api/al-bhed/body/samples",
			expectedStatus: http.StatusOK,
			handler: 		testCfg.HandleAlBhedGet,
		},
	}

	testServiceResponsesGet[[]SampleParams[AlBhedParams]](t, tests, "AlBhedBodySamples")
}

func TestTurnOrderBodySamples(t *testing.T) {
	t.Parallel()

	tests := []testGeneral{
		{
			requestURL:     "/api/turn-order/body/samples",
			expectedStatus: http.StatusOK,
			handler: 		testCfg.HandleTurnOrderGet,
		},
	}

	testServiceResponsesGet[[]SampleParams[TurnOrderParams]](t, tests, "TurnOrderBodySamples")
}

func TestDropChanceBodySamples(t *testing.T) {
	t.Parallel()

	tests := []testGeneral{
		{
			requestURL:     "/api/drop-chance/body/samples",
			expectedStatus: http.StatusOK,
			handler: 		testCfg.HandleDropChanceGet,
		},
	}

	testServiceResponsesGet[[]SampleParams[DropChanceParams]](t, tests, "DropChanceBodySamples")
}