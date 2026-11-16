package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/api"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/service/servicedocs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/service/servicetypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/tests/apitesttypes"
)

func TestServiceBody(t *testing.T) {
	t.Parallel()

	tests := []apitesttypes.TestGeneral{
		{
			RequestURL:     "/api/al-bhed/body",
			ExpectedStatus: http.StatusOK,
			Handler:        testSrv.HandleAlBhedGet,
		},
		{
			RequestURL:     "/api/turn-order/body",
			ExpectedStatus: http.StatusOK,
			Handler:        testSrv.HandleTurnOrderGet,
		},
		{
			RequestURL:     "/api/drop-chance/body",
			ExpectedStatus: http.StatusOK,
			Handler:        testSrv.HandleDropChanceGet,
		},
		{
			RequestURL:     "/api/al-bhed/",
			ExpectedStatus: http.StatusBadRequest,
			ExpectedErr:    "wrong format. usage: 'GET /api/al-bhed/body', 'POST /api/al-bhed'",
			Handler:        testSrv.HandleAlBhedGet,
		},
		{
			RequestURL:     "/api/al-bhed/body/2/2",
			ExpectedStatus: http.StatusBadRequest,
			ExpectedErr:    "wrong format. usage: 'GET /api/al-bhed/body', 'POST /api/al-bhed'",
			Handler:        testSrv.HandleAlBhedGet,
		},
		{
			RequestURL:     "/api/al-bhed/body/asd",
			ExpectedStatus: http.StatusBadRequest,
			ExpectedErr:    "'GET api/al-bhed/body' only supports the following sections: 'docs', 'samples'",
			Handler:        testSrv.HandleAlBhedGet,
		},
		{
			RequestURL:     "/api/al-bhed/bodsa",
			ExpectedStatus: http.StatusBadRequest,
			ExpectedErr:    "'GET api/al-bhed' only supports '/body'",
			Handler:        testSrv.HandleAlBhedGet,
		},
		{
			RequestURL:     "/api/al-bhed/body?limit=23",
			ExpectedStatus: http.StatusBadRequest,
			ExpectedErr:    "computational endpoints don't have any query parameters",
			Handler:        testSrv.HandleAlBhedGet,
		},
	}

	testServiceResponsesGet[api.Folder](t, tests, "ServiceBody")
}

func TestServiceBodyDocs(t *testing.T) {
	t.Parallel()

	tests := []apitesttypes.TestGeneral{
		{
			RequestURL:     "/api/al-bhed/body/docs",
			ExpectedStatus: http.StatusOK,
			Handler:        testSrv.HandleAlBhedGet,
		},
		{
			RequestURL:     "/api/turn-order/body/docs",
			ExpectedStatus: http.StatusOK,
			Handler:        testSrv.HandleTurnOrderGet,
		},
		{
			RequestURL:     "/api/drop-chance/body/docs",
			ExpectedStatus: http.StatusOK,
			Handler:        testSrv.HandleDropChanceGet,
		},
	}

	testServiceResponsesGet[servicedocs.ParamsDoc](t, tests, "ServiceBodyDocs")
}

func TestAlBhedBodySamples(t *testing.T) {
	t.Parallel()

	tests := []apitesttypes.TestGeneral{
		{
			RequestURL:     "/api/al-bhed/body/samples",
			ExpectedStatus: http.StatusOK,
			Handler:        testSrv.HandleAlBhedGet,
		},
	}

	testServiceResponsesGet[[]servicetypes.SampleParams[servicetypes.AlBhedParams]](t, tests, "AlBhedBodySamples")
}

func TestTurnOrderBodySamples(t *testing.T) {
	t.Parallel()

	tests := []apitesttypes.TestGeneral{
		{
			RequestURL:     "/api/turn-order/body/samples",
			ExpectedStatus: http.StatusOK,
			Handler:        testSrv.HandleTurnOrderGet,
		},
	}

	testServiceResponsesGet[[]servicetypes.SampleParams[servicetypes.TurnOrderParams]](t, tests, "TurnOrderBodySamples")
}

func TestDropChanceBodySamples(t *testing.T) {
	t.Parallel()

	tests := []apitesttypes.TestGeneral{
		{
			RequestURL:     "/api/drop-chance/body/samples",
			ExpectedStatus: http.StatusOK,
			Handler:        testSrv.HandleDropChanceGet,
		},
	}

	testServiceResponsesGet[[]servicetypes.SampleParams[servicetypes.DropChanceParams]](t, tests, "DropChanceBodySamples")
}
