package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)


type routerEnums struct {
	routerParams
	i handlerinput.Enums
}

func newRouterEnums(w http.ResponseWriter, r *http.Request, i handlerinput.Enums) routerEnums {
	return routerEnums{
		routerParams: newRouterParams(w, r, i.Endpoint),
		i:            i,
	}
}

func (r routerEnums) route() {
	switch len(r.segments) {
	case 0:
		r.list()
		return

	case 1:
		r.single()
		return

	default:
		r.formatError()
		return
	}
}

func (r routerEnums) list() {
	err := r.i.VerifyQueriesRetrieve(r.r)
	if HandleHTTPError(r.w, err) {
		return
	}

	resourceList, err := r.i.Retrieve(r.r)
	if HandleHTTPError(r.w, err) {
		return
	}

	respondWithJSON(r.w, http.StatusOK, resourceList)
}

func (r routerEnums) single() {
	if r.segments[0] == string(ts.SnParameters) {
		HandleParametersEnums(r.w, r.r, r.i)
		return
	}

	enum, key, err := ParseSegment(r.i, r.segments[0])
	if HandleHTTPError(r.w, err) {
		return
	}

	// happens after enum parsing, just for error priority reasons. if I prioritize wrong query param errors over wrong enum errors, I can simply swap the two functions
	err = r.i.VerifyQueriesGet(r.r, key)
	if HandleHTTPError(r.w, err) {
		return
	}

	respondWithJSON(r.w, http.StatusOK, enum)
}

func (r routerEnums) formatError() {
	respondWithError(r.w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: %s", h.FormatStringSlice(r.i.Usage)), nil)
}