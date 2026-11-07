package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func routerIDOnly[T seedtypes.Lookupable, R any, A apiresource.APIResource](w http.ResponseWriter, r *http.Request, i HandlerInput[T, R, A]) {
	i.usage = []string{"/{id}"}
	segments := getPathSegments(r.URL.Path, i.Endpoint)

	switch len(segments) {
	case 0:
		i.HandleList(w, r)
		return

	case 1:
		i.HandleIDOnly(w, r, segments)
		return

	case 2:
		i.HandleSubsections(w, r, segments)
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: %s", i.CompileUsage()), nil)
		return
	}
}

func routerNameID[T seedtypes.Lookupable, R any, A apiresource.APIResource](w http.ResponseWriter, r *http.Request, i HandlerInput[T, R, A]) {
	i.usage = []string{"/{id}", "/{name}"}
	segments := getPathSegments(r.URL.Path, i.Endpoint)

	switch len(segments) {
	case 0:
		i.HandleList(w, r)
		return

	case 1:
		i.HandleNameID(w, r, segments)
		return

	case 2:
		i.HandleSubsections(w, r, segments)
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: %s", i.CompileUsage()), nil)
		return
	}
}

func routerNameVersion[T seedtypes.Lookupable, R any, A apiresource.APIResource](w http.ResponseWriter, r *http.Request, i HandlerInput[T, R, A]) {
	i.usage = []string{"/{id}", "/{name}", "/{name}/{version}"}
	segments := getPathSegments(r.URL.Path, i.Endpoint)

	switch len(segments) {
	case 0:
		i.HandleList(w, r)
		return

	case 1:
		i.HandleNameID(w, r, segments)
		return

	case 2:
		i.HandleSubOrNameVer(w, r, segments)
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: %s", i.CompileUsage()), nil)
		return
	}
}

func routerEnums(w http.ResponseWriter, r *http.Request, i HandlerInputEnums) {
	segments := getPathSegments(r.URL.Path, i.Endpoint)

	switch len(segments) {
	case 0:
		i.HandleList(w, r)
		return

	case 1:
		i.HandleSingle(w, r, segments)
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: %s", h.FormatStringSlice(i.usage)), nil)
		return
	}
}

func routerEndpoints(w http.ResponseWriter, r *http.Request, i HandlerInputEndpoints) {
	segments := getPathSegments(r.URL.Path, i.Endpoint)

	switch len(segments) {
	case 0:
		i.HandleList(w, r)
		return

	case 1:
		i.HandleSections(w, r, segments)
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: %s", h.FormatStringSlice(i.usage)), nil)
		return
	}
}

func routerServiceGet[P ServiceParams, R ServiceResponse](w http.ResponseWriter, r *http.Request, i HandlerInputService[P, R]) {
	segments := getPathSegments(r.URL.Path, i.Endpoint)

	switch len(segments) {
	case 1:
		i.HandleBody(w, r, segments)
		return

	case 2:
		i.HandleBodySections(w, r, segments)
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: 'GET /api/%s/body', 'POST /api/%s'", i.Endpoint, i.Endpoint), nil)
		return
	}
}

func routerServicePost[P ServiceParams, R ServiceResponse](w http.ResponseWriter, r *http.Request, i HandlerInputService[P, R]) {
	segments := getPathSegments(r.URL.Path, i.Endpoint)

	switch len(segments) {
	case 0:
		i.HandlePost(w, r)
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: /api/%s. Use /api/%s/body to learn more about the required json format for the request body", i.Endpoint, i.Endpoint), nil)
		return
	}
}
