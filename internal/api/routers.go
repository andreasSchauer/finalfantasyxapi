package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func routerIDOnly[T seedtypes.Lookupable, R any, A apiresource.APIResource](w http.ResponseWriter, r *http.Request, i handlerinput.HandlerInput[T, R, A]) {
	i.Usage = []string{"/{id}"}
	segments := getPathSegments(r.URL.Path, i.Endpoint)

	switch len(segments) {
	case 0:
		HandleList(w, r, i)
		return

	case 1:
		HandleIDOnly(w, r, i, segments[0])
		return

	case 2:
		HandleSubsections(w, r, i, segments[0], segments[1])
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: %s", i.CompileUsage()), nil)
		return
	}
}

func routerNameID[T seedtypes.Lookupable, R any, A apiresource.APIResource](w http.ResponseWriter, r *http.Request, i handlerinput.HandlerInput[T, R, A]) {
	i.Usage = []string{"/{id}", "/{name}"}
	segments := getPathSegments(r.URL.Path, i.Endpoint)

	switch len(segments) {
	case 0:
		HandleList(w, r, i)
		return

	case 1:
		HandleNameID(w, r, i, segments[0])
		return

	case 2:
		HandleSubsections(w, r, i, segments[0], segments[1])
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: %s", i.CompileUsage()), nil)
		return
	}
}

func routerNameVersion[T seedtypes.Lookupable, R any, A apiresource.APIResource](w http.ResponseWriter, r *http.Request, i handlerinput.HandlerInput[T, R, A]) {
	i.Usage = []string{"/{id}", "/{name}", "/{name}/{version}"}
	segments := getPathSegments(r.URL.Path, i.Endpoint)

	switch len(segments) {
	case 0:
		HandleList(w, r, i)
		return

	case 1:
		HandleNameID(w, r, i, segments[0])
		return

	case 2:
		HandleSubOrNameVer(w, r, i, segments[0], segments[1])
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: %s", i.CompileUsage()), nil)
		return
	}
}

func routerEnums(w http.ResponseWriter, r *http.Request, i handlerinput.HandlerInputEnums) {
	segments := getPathSegments(r.URL.Path, i.Endpoint)

	switch len(segments) {
	case 0:
		HandleListEnums(w, r, i)
		return

	case 1:
		HandleSingleEnums(w, r, i, segments[0])
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: %s", h.FormatStringSlice(i.Usage)), nil)
		return
	}
}

func routerEndpoints(w http.ResponseWriter, r *http.Request, i handlerinput.HandlerInputEndpoints) {
	segments := getPathSegments(r.URL.Path, i.Endpoint)

	switch len(segments) {
	case 0:
		HandleListEndpoints(w, r, i)
		return

	case 1:
		HandleSectionsEndpoints(w, r, i, segments[0])
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: %s", h.FormatStringSlice(i.Usage)), nil)
		return
	}
}

func routerServiceGet[P ServiceParams, R ServiceResponse](w http.ResponseWriter, r *http.Request, i HandlerInputService[P, R]) {
	segments := getPathSegments(r.URL.Path, i.Endpoint)

	switch len(segments) {
	case 1:
		HandleBody(w, r, i, segments[0])
		return

	case 2:
		HandleBodySections(w, r, i, segments)
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
		HandlePost(w, r,i)
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: /api/%s. Use /api/%s/body to learn more about the required json format for the request body", i.Endpoint, i.Endpoint), nil)
		return
	}
}
