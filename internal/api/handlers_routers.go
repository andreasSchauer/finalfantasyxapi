package api

import (
	"fmt"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func routerIDOnly[T seedtypes.Lookupable, R any, A APIResource](w http.ResponseWriter, r *http.Request, i handlerInput[T, R, A]) {
	i.usage = []string{"/{id}"}
	segments := getPathSegments(r.URL.Path, i.Endpoint)

	switch len(segments) {
	case 0:
		i.handleList(w, r)
		return

	case 1:
		i.handleIDOnly(w, r, segments)
		return

	case 2:
		i.handleSubsections(w, r, segments)
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: %s", i.compileUsage()), nil)
		return
	}
}

func routerNameID[T seedtypes.Lookupable, R any, A APIResource](w http.ResponseWriter, r *http.Request, i handlerInput[T, R, A]) {
	i.usage = []string{"/{id}", "/{name}"}
	segments := getPathSegments(r.URL.Path, i.Endpoint)

	switch len(segments) {
	case 0:
		i.handleList(w, r)
		return

	case 1:
		i.handleNameID(w, r, segments)
		return

	case 2:
		i.handleSubsections(w, r, segments)
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: %s", i.compileUsage()), nil)
		return
	}
}

func routerNameVersion[T seedtypes.Lookupable, R any, A APIResource](w http.ResponseWriter, r *http.Request, i handlerInput[T, R, A]) {
	i.usage = []string{"/{id}", "/{name}", "/{name}/{version}"}
	segments := getPathSegments(r.URL.Path, i.Endpoint)

	switch len(segments) {
	case 0:
		i.handleList(w, r)
		return

	case 1:
		i.handleNameID(w, r, segments)
		return

	case 2:
		i.handleSubOrNameVer(w, r, segments)
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: %s", i.compileUsage()), nil)
		return
	}
}

func routerEnums(w http.ResponseWriter, r *http.Request, i handlerInputEnums) {
	segments := getPathSegments(r.URL.Path, i.Endpoint)

	switch len(segments) {
	case 0:
		i.handleList(w, r)
		return

	case 1:
		i.handleSingle(w, r, segments)
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: %s", h.FormatStringSlice(i.usage)), nil)
		return
	}
}

func routerEndpoints(w http.ResponseWriter, r *http.Request, i handlerInputEndpoints) {
	segments := getPathSegments(r.URL.Path, i.Endpoint)

	switch len(segments) {
	case 0:
		i.handleList(w, r)
		return

	case 1:
		i.handleSections(w, r, segments)
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: %s", h.FormatStringSlice(i.usage)), nil)
		return
	}
}

func routerServiceGet[P ServiceParams, R ServiceResponse](w http.ResponseWriter, r *http.Request, i handlerInputService[P, R]) {
	segments := getPathSegments(r.URL.Path, i.Endpoint)

	switch len(segments) {
	case 1:
		i.handleBody(w, r, segments)
		return

	case 2:
		i.handleBodySections(w, r, segments)
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: 'GET /api/%s/body', 'POST /api/%s'", i.Endpoint, i.Endpoint), nil)
		return
	}
}

func routerServicePost[P ServiceParams, R ServiceResponse](w http.ResponseWriter, r *http.Request, i handlerInputService[P, R]) {
	segments := getPathSegments(r.URL.Path, i.Endpoint)

	switch len(segments) {
	case 0:
		i.handlePost(w, r)
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: /api/%s. Use /api/%s/body to learn more about the required json format for the request body", i.Endpoint, i.Endpoint), nil)
		return
	}
}
