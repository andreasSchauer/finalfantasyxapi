package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api"
	"github.com/andreasSchauer/finalfantasyxapi/internal/servicetypes"
)

func TestServiceBody(t *testing.T) {
	t.Parallel()

	tests := []testGeneral{
		{
			requestURL:     "/api/al-bhed/body",
			expectedStatus: http.StatusOK,
			handler:        testSrv.HandleAlBhedGet,
		},
		{
			requestURL:     "/api/turn-order/body",
			expectedStatus: http.StatusOK,
			handler:        testSrv.HandleTurnOrderGet,
		},
		{
			requestURL:     "/api/drop-chance/body",
			expectedStatus: http.StatusOK,
			handler:        testSrv.HandleDropChanceGet,
		},
		{
			requestURL:     "/api/al-bhed/",
			expectedStatus: http.StatusBadRequest,
			expectedErr:    "wrong format. usage: 'GET /api/al-bhed/body', 'POST /api/al-bhed'",
			handler:        testSrv.HandleAlBhedGet,
		},
		{
			requestURL:     "/api/al-bhed/body/2/2",
			expectedStatus: http.StatusBadRequest,
			expectedErr:    "wrong format. usage: 'GET /api/al-bhed/body', 'POST /api/al-bhed'",
			handler:        testSrv.HandleAlBhedGet,
		},
		{
			requestURL:     "/api/al-bhed/body/asd",
			expectedStatus: http.StatusBadRequest,
			expectedErr:    "'GET api/al-bhed/body' only supports the following sections: 'docs', 'samples'",
			handler:        testSrv.HandleAlBhedGet,
		},
		{
			requestURL:     "/api/al-bhed/bodsa",
			expectedStatus: http.StatusBadRequest,
			expectedErr:    "'GET api/al-bhed' only supports '/body'",
			handler:        testSrv.HandleAlBhedGet,
		},
		{
			requestURL:     "/api/al-bhed/body?limit=23",
			expectedStatus: http.StatusBadRequest,
			expectedErr:    "computational endpoints don't have any query parameters",
			handler:        testSrv.HandleAlBhedGet,
		},
	}

	testServiceResponsesGet[api.Folder](t, tests, "ServiceBody")
}

func TestServiceBodyDocs(t *testing.T) {
	t.Parallel()

	tests := []testGeneral{
		{
			requestURL:     "/api/al-bhed/body/docs",
			expectedStatus: http.StatusOK,
			handler:        testSrv.HandleAlBhedGet,
		},
		{
			requestURL:     "/api/turn-order/body/docs",
			expectedStatus: http.StatusOK,
			handler:        testSrv.HandleTurnOrderGet,
		},
		{
			requestURL:     "/api/drop-chance/body/docs",
			expectedStatus: http.StatusOK,
			handler:        testSrv.HandleDropChanceGet,
		},
	}

	testServiceResponsesGet[servicetypes.ParamsDoc](t, tests, "ServiceBodyDocs")
}

func TestAlBhedBodySamples(t *testing.T) {
	t.Parallel()

	tests := []testGeneral{
		{
			requestURL:     "/api/al-bhed/body/samples",
			expectedStatus: http.StatusOK,
			handler:        testSrv.HandleAlBhedGet,
		},
	}

	testServiceResponsesGet[[]servicetypes.SampleParams[servicetypes.AlBhedParams]](t, tests, "AlBhedBodySamples")
}

func TestTurnOrderBodySamples(t *testing.T) {
	t.Parallel()

	tests := []testGeneral{
		{
			requestURL:     "/api/turn-order/body/samples",
			expectedStatus: http.StatusOK,
			handler:        testSrv.HandleTurnOrderGet,
		},
	}

	testServiceResponsesGet[[]servicetypes.SampleParams[servicetypes.TurnOrderParams]](t, tests, "TurnOrderBodySamples")
}

func TestDropChanceBodySamples(t *testing.T) {
	t.Parallel()

	tests := []testGeneral{
		{
			requestURL:     "/api/drop-chance/body/samples",
			expectedStatus: http.StatusOK,
			handler:        testSrv.HandleDropChanceGet,
		},
	}

	testServiceResponsesGet[[]servicetypes.SampleParams[servicetypes.DropChanceParams]](t, tests, "DropChanceBodySamples")
}
