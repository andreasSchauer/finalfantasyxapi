package api

import (
	"fmt"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func routerIdOnly[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, w http.ResponseWriter, r *http.Request, i handlerInput[T, R, A, L]) {
	i.usage = []string{"/{id}"}
	segments := getPathSegments(r.URL.Path, i.endpoint)

	switch len(segments) {
	case 0:
		handleEndpointList(cfg, w, r, i)
		return

	case 1:
		handleEndpointIDOnly(cfg, w, r, i, segments)
		return

	case 2:
		handleEndpointSubsections(cfg, w, r, i, segments)
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: %s", getUsageString(i)), nil)
		return
	}
}

func routerNameOrID[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, w http.ResponseWriter, r *http.Request, i handlerInput[T, R, A, L]) {
	i.usage = []string{"/{id}", "/{name}"}
	segments := getPathSegments(r.URL.Path, i.endpoint)

	switch len(segments) {
	case 0:
		handleEndpointList(cfg, w, r, i)
		return

	case 1:
		handleEndpointNameOrID(cfg, w, r, i, segments)
		return

	case 2:
		handleEndpointSubsections(cfg, w, r, i, segments)
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: %s", getUsageString(i)), nil)
		return
	}
}

func routerNameVersion[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, w http.ResponseWriter, r *http.Request, i handlerInput[T, R, A, L]) {
	i.usage = []string{"/{id}", "/{name}", "/{name}/{version}"}
	segments := getPathSegments(r.URL.Path, i.endpoint)

	switch len(segments) {
	case 0:
		handleEndpointList(cfg, w, r, i)
		return

	case 1:
		handleEndpointNameOrID(cfg, w, r, i, segments)
		return

	case 2:
		handleEndpointSubOrNameVer(cfg, w, r, i, segments)
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: %s", getUsageString(i)), nil)
		return
	}
}

func routerEnums(cfg *Config, w http.ResponseWriter, r *http.Request, i handlerInputEnums) {
	segments := getPathSegments(r.URL.Path, i.endpoint)

	switch len(segments) {
	case 0:
		handleEnumsEndpointList(cfg, w, r, i)
		return

	case 1:
		handleEnumsEndpointSingle(cfg, w, r, i, segments)
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: %s", h.FormatStringSlice(i.usage)), nil)
		return
	}
}

func routerEndpoints(cfg *Config, w http.ResponseWriter, r *http.Request, i handlerInputEndpoints) {
	segments := getPathSegments(r.URL.Path, i.endpoint)

	switch len(segments) {
	case 0:
		handleEndpointsEndpointList(cfg, w, r, i)
		return

	case 1:
		handleEndpointsEndpointSections(cfg, w, r, i, segments)
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: %s", h.FormatStringSlice(i.usage)), nil)
		return
	}
}

func routerServiceGet[R any](cfg *Config, w http.ResponseWriter, r *http.Request, i handlerInputService[R]) {
	segments := getPathSegments(r.URL.Path, i.endpoint)

	switch len(segments) {
	case 0:
		handleEndpointServiceGet(cfg, w, r, i)
		return

	case 1:
		handleServiceEndpointSections(cfg, w, r, i, segments)
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: /api/%s?state={json_string}", i.endpoint), nil)
		return
	}
}

func routerServicePost[R any](cfg *Config, w http.ResponseWriter, r *http.Request, i handlerInputService[R]) {
	segments := getPathSegments(r.URL.Path, i.endpoint)

	switch len(segments) {
	case 0:
		handleEndpointServicePost(cfg, w, r, i)
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: /api/%s. Use /api/%s/body to learn more about the required json format for the request body.", i.endpoint, i.endpoint), nil)
		return
	}
}
