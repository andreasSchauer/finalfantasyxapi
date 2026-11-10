package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/servicetypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)


type routerServiceGet[P servicetypes.ServiceParams, R servicetypes.ServiceResponse] struct {
	routerParams
	i handlerinput.Service[P, R]
}

func (r routerServiceGet[P, R]) route() {
	switch len(r.segments) {
	case 1:
		r.body()
		return

	case 2:
		r.subsections()
		return

	default:
		r.formatError()
		return
	}
}

func (r routerServiceGet[P, R]) body() {
	err := r.i.VerifyQueries(r.r)
	if HandleHTTPError(r.w, err) {
		return
	}

	if r.segments[0] == string(ts.SnBody) {
		respondWithJSON(r.w, http.StatusOK, getBody(r.i.Endpoint))
		return
	}

	respondWithError(r.w, http.StatusBadRequest, fmt.Sprintf("'GET api/%s' only supports '/body'", r.i.Endpoint), nil)
}

func (r routerServiceGet[P, R]) subsections() {
	err := r.i.VerifyQueries(r.r)
	if HandleHTTPError(r.w, err) {
		return
	}

	if r.segments[0] != string(ts.SnBody) {
		respondWithError(r.w, http.StatusBadRequest, fmt.Sprintf("'GET api/%s' only supports '/body'", r.i.Endpoint), nil)
		return
	}

	switch r.segments[1] {
	case string(ts.SnDocs):
		respondWithJSON(r.w, http.StatusOK, r.i.DocsFn())
		return

	case string(ts.SnSamples):
		respondWithJSON(r.w, http.StatusOK, r.i.Samples)
		return

	default:
		respondWithError(r.w, http.StatusBadRequest, fmt.Sprintf("'GET api/%s/body' only supports the following sections: 'docs', 'samples'", r.i.Endpoint), nil)
		return
	}
}

func (r routerServiceGet[P, R]) formatError() {
	respondWithError(r.w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: 'GET /api/%s/body', 'POST /api/%s'", r.i.Endpoint, r.i.Endpoint), nil)
}

func newRouterServiceGet[P servicetypes.ServiceParams, R servicetypes.ServiceResponse](w http.ResponseWriter, r *http.Request, i handlerinput.Service[P, R]) routerServiceGet[P, R] {
	return routerServiceGet[P, R]{
		routerParams: newRouterParams(w, r, i.Endpoint),
		i:            i,
	}
}

type routerServicePost[P servicetypes.ServiceParams, R servicetypes.ServiceResponse] struct {
	routerParams
	i handlerinput.Service[P, R]
}

func (r routerServicePost[P, R]) route() {
	switch len(r.segments) {
	case 0:
		r.handle()
		return

	default:
		r.formatError()
		return
	}
}

func (r routerServicePost[P, R]) handle() {
	err := r.i.VerifyQueries(r.r)
	if HandleHTTPError(r.w, err) {
		return
	}

	response, err := r.i.ServicePost(r.r)
	if HandleHTTPError(r.w, err) {
		return
	}

	respondWithJSON(r.w, http.StatusOK, response)
}

func (r routerServicePost[P, R]) formatError() {
	respondWithError(r.w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: /api/%s. Use /api/%s/body to learn more about the required json format for the request body", r.i.Endpoint, r.i.Endpoint), nil)
}

func newRouterServicePost[P servicetypes.ServiceParams, R servicetypes.ServiceResponse](w http.ResponseWriter, r *http.Request, i handlerinput.Service[P, R]) routerServicePost[P, R] {
	return routerServicePost[P, R]{
		routerParams: newRouterParams(w, r, i.Endpoint),
		i:            i,
	}
}
