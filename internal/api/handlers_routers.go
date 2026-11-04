package api

import (
	"fmt"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func (i handlerInput[T, R, A]) routerIDOnly(server *Server, w http.ResponseWriter, r *http.Request) {
	i.usage = []string{"/{id}"}
	segments := getPathSegments(r.URL.Path, i.Endpoint)

	switch len(segments) {
	case 0:
		i.handleList(server, w, r)
		return

	case 1:
		i.handleIDOnly(server, w, r, segments)
		return

	case 2:
		i.handleSubsections(server, w, r, segments)
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: %s", i.compileUsage()), nil)
		return
	}
}

func (i handlerInput[T, R, A]) routerNameID(server *Server, w http.ResponseWriter, r *http.Request) {
	i.usage = []string{"/{id}", "/{name}"}
	segments := getPathSegments(r.URL.Path, i.Endpoint)

	switch len(segments) {
	case 0:
		i.handleList(server, w, r)
		return

	case 1:
		i.handleNameID(server, w, r, segments)
		return

	case 2:
		i.handleSubsections(server, w, r, segments)
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: %s", i.compileUsage()), nil)
		return
	}
}

func (i handlerInput[T, R, A]) routerNameVersion(server *Server, w http.ResponseWriter, r *http.Request) {
	i.usage = []string{"/{id}", "/{name}", "/{name}/{version}"}
	segments := getPathSegments(r.URL.Path, i.Endpoint)

	switch len(segments) {
	case 0:
		i.handleList(server, w, r)
		return

	case 1:
		i.handleNameID(server, w, r, segments)
		return

	case 2:
		i.handleSubOrNameVer(server, w, r, segments)
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: %s", i.compileUsage()), nil)
		return
	}
}

func (i handlerInputEnums) router(server *Server, w http.ResponseWriter, r *http.Request) {
	segments := getPathSegments(r.URL.Path, i.Endpoint)

	switch len(segments) {
	case 0:
		i.handleList(server, w, r)
		return

	case 1:
		i.handleSingle(server, w, r, segments)
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: %s", h.FormatStringSlice(i.usage)), nil)
		return
	}
}

func (i handlerInputEndpoints) router(server *Server, w http.ResponseWriter, r *http.Request) {
	segments := getPathSegments(r.URL.Path, i.Endpoint)

	switch len(segments) {
	case 0:
		i.handleList(server, w, r)
		return

	case 1:
		i.handleSections(server, w, r, segments)
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: %s", h.FormatStringSlice(i.usage)), nil)
		return
	}
}

func (i handlerInputService[P, R]) routerGet(server *Server, w http.ResponseWriter, r *http.Request) {
	segments := getPathSegments(r.URL.Path, i.Endpoint)

	switch len(segments) {
	case 1:
		i.handleBody(server, w, r, segments)
		return

	case 2:
		i.handleBodySections(server, w, r, segments)
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: 'GET /api/%s/body', 'POST /api/%s'", i.Endpoint, i.Endpoint), nil)
		return
	}
}

func (i handlerInputService[P, R]) routerPost(server *Server, w http.ResponseWriter, r *http.Request) {
	segments := getPathSegments(r.URL.Path, i.Endpoint)

	switch len(segments) {
	case 0:
		i.handlePost(server, w, r)
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: /api/%s. Use /api/%s/body to learn more about the required json format for the request body", i.Endpoint, i.Endpoint), nil)
		return
	}
}
