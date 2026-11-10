package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/servicetypes"
)

func routeIDOnly[T seedtypes.Lookupable, R any, A apiresource.APIResource](w http.ResponseWriter, r *http.Request, i handlerinput.Main[T, R, A]) {
	router := newRouterIDOnly(w, r, i)
	router.route()
}

func routeNameID[T seedtypes.Lookupable, R any, A apiresource.APIResource](w http.ResponseWriter, r *http.Request, i handlerinput.Main[T, R, A]) {
	router := newRouterNameID(w, r, i)
	router.route()
}

func routeNameVersion[T seedtypes.Lookupable, R any, A apiresource.APIResource](w http.ResponseWriter, r *http.Request, i handlerinput.Main[T, R, A]) {
	router := newRouterNameVersion(w, r, i)
	router.route()
}

func routeEnums(w http.ResponseWriter, r *http.Request, i handlerinput.Enums) {
	router := newRouterEnums(w, r, i)
	router.route()
}

func routeEndpoints(w http.ResponseWriter, r *http.Request, i handlerinput.Endpoints) {
	router := newRouterEndpoints(w, r, i)
	router.route()
}

func routeServiceGet[P servicetypes.ServiceParams, R servicetypes.ServiceResponse](w http.ResponseWriter, r *http.Request, i handlerinput.Service[P, R]) {
	router := newRouterServiceGet(w, r, i)
	router.route()
}

func routeServicePost[P servicetypes.ServiceParams, R servicetypes.ServiceResponse](w http.ResponseWriter, r *http.Request, i handlerinput.Service[P, R]) {
	router := newRouterServicePost(w, r, i)
	router.route()
}