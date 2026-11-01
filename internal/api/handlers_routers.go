package api

import (
	"fmt"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func (i handlerInput[T, R, A]) routerIDOnly(cfg *Config, w http.ResponseWriter, r *http.Request) {
	i.usage = []string{"/{id}"}
	segments := getPathSegments(r.URL.Path, i.Endpoint)

	switch len(segments) {
	case 0:
		i.handleList(cfg, w, r)
		return

	case 1:
		i.handleIDOnly(cfg, w, r, segments)
		return

	case 2:
		i.handleSubsections(cfg, w, r, segments)
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: %s", i.compileUsage()), nil)
		return
	}
}

func (i handlerInput[T, R, A]) routerNameID(cfg *Config, w http.ResponseWriter, r *http.Request) {
	i.usage = []string{"/{id}", "/{name}"}
	segments := getPathSegments(r.URL.Path, i.Endpoint)

	switch len(segments) {
	case 0:
		i.handleList(cfg, w, r)
		return

	case 1:
		i.handleNameID(cfg, w, r, segments)
		return

	case 2:
		i.handleSubsections(cfg, w, r, segments)
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: %s", i.compileUsage()), nil)
		return
	}
}

func (i handlerInput[T, R, A]) routerNameVersion(cfg *Config, w http.ResponseWriter, r *http.Request) {
	i.usage = []string{"/{id}", "/{name}", "/{name}/{version}"}
	segments := getPathSegments(r.URL.Path, i.Endpoint)

	switch len(segments) {
	case 0:
		i.handleList(cfg, w, r)
		return

	case 1:
		i.handleNameID(cfg, w, r, segments)
		return

	case 2:
		i.handleSubOrNameVer(cfg, w, r, segments)
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: %s", i.compileUsage()), nil)
		return
	}
}

func (i handlerInputEnums) router(cfg *Config, w http.ResponseWriter, r *http.Request) {
	segments := getPathSegments(r.URL.Path, i.Endpoint)

	switch len(segments) {
	case 0:
		i.handleList(cfg, w, r)
		return

	case 1:
		i.handleSingle(cfg, w, r, segments)
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: %s", h.FormatStringSlice(i.usage)), nil)
		return
	}
}

func (i handlerInputEndpoints) router(cfg *Config, w http.ResponseWriter, r *http.Request) {
	segments := getPathSegments(r.URL.Path, i.Endpoint)

	switch len(segments) {
	case 0:
		i.handleList(cfg, w, r)
		return

	case 1:
		i.handleSections(cfg, w, r, segments)
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: %s", h.FormatStringSlice(i.usage)), nil)
		return
	}
}

func (i handlerInputService[P, R]) routerGet(cfg *Config, w http.ResponseWriter, r *http.Request) {
	segments := getPathSegments(r.URL.Path, i.Endpoint)

	switch len(segments) {
	case 1:
		i.handleBody(cfg, w, r, segments)
		return

	case 2:
		i.handleBodySections(cfg, w, r, segments)
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: 'GET /api/%s/body', 'POST /api/%s'.", i.Endpoint, i.Endpoint), nil)
		return
	}
}

func (i handlerInputService[P, R]) routerPost(cfg *Config, w http.ResponseWriter, r *http.Request) {
	segments := getPathSegments(r.URL.Path, i.Endpoint)

	switch len(segments) {
	case 0:
		i.handlePost(cfg, w, r)
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: /api/%s. Use /api/%s/body to learn more about the required json format for the request body.", i.Endpoint, i.Endpoint), nil)
		return
	}
}
