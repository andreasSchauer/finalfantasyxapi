package api

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api/apihelpers"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
)

type routerParams struct {
	w        http.ResponseWriter
	r        *http.Request
	segments []string
}

func newRouterParams(w http.ResponseWriter, r *http.Request, endpoint ts.EndpointName) routerParams {
	return routerParams{
		w:        w,
		r:        r,
		segments: ah.GetPathSegments(r.URL.Path, endpoint),
	}
}
