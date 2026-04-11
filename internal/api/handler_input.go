package api

import (
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type handlerInput[T h.HasID, R any, A APIResource, L APIResourceList] struct {
	endpoint         string
	resourceType     string
	usage            []string
	objLookup        map[string]T
	objLookupID      map[int32]T
	queryLookup      map[string]QueryParam
	getMultipleQuery DbQueryStringMany
	retrieveQuery    DbQueryNoInput
	idToResFunc      func(*Config, handlerInput[T, R, A, L], int32) A
	resToListFunc    func(*Config, *http.Request, []A) (L, error)
	getSingleFunc    func(*http.Request, handlerInput[T, R, A, L], int32) (R, error)
	retrieveFunc     func(*http.Request, handlerInput[T, R, A, L]) (L, error)
	subsections      map[string]SubSectionFns
}

type SubSectionFns struct {
	dbQuery     DbQueryIntMany
	createSubFn func(*Config, *http.Request, int32) (SimpleResource, error)
}
