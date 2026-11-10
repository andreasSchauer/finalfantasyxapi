package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

type routerEndpoints struct {
	routerParams
	i handlerinput.Endpoints
}

func newRouterEndpoints(w http.ResponseWriter, r *http.Request, i handlerinput.Endpoints) routerEndpoints {
	return routerEndpoints{
		routerParams: newRouterParams(w, r, i.Endpoint),
		i:            i,
	}
}

func (r routerEndpoints) route() {
	switch len(r.segments) {
	case 0:
		r.list()
		return

	case 1:
		r.subsections()
		return

	default:
		r.formatError()
		return
	}
}

func (r routerEndpoints) list() {
	err := r.i.VerifyQueries(r.r)
	if HandleHTTPError(r.w, err) {
		return
	}

	resourceList, err := r.i.Retrieve(r.r)
	if HandleHTTPError(r.w, err) {
		return
	}
	respondWithJSON(r.w, http.StatusOK, resourceList)
}

func (r routerEndpoints) subsections() {
	if r.segments[0] == string(ts.SnParameters) {
		HandleParametersEndpoints(r.w, r.r, r.i)
		return
	}

	respondWithError(r.w, http.StatusBadRequest, "wrong format. '/endpoints' doesn't support single-resource requests", nil)
}

func (r routerEndpoints) formatError() {
	respondWithError(r.w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: %s", h.FormatStringSlice(r.i.Usage)), nil)
}