package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type handlerInput[T seeding.Lookupable, R any, A APIResource, L APIResourceList] struct {
	endpoint         EndpointName
	resTypeSing      ResTypeSingular
	resTypePlural	 ResTypePlural
	usage            []string
	objLookup        map[string]T
	objLookupID      map[int32]T
	queryLookup      map[QueryParamName]QueryParam
	getMultipleQuery DbQueryStringMany
	retrieveQuery    DbQueryNoInput
	idToResFunc      func(*Config, handlerInput[T, R, A, L], int32) A
	resToListFunc    func(*Config, *http.Request, []A) (L, error)
	getSingleFunc    func(*http.Request, handlerInput[T, R, A, L], int32) (R, error)
	retrieveFunc     func(*http.Request, handlerInput[T, R, A, L]) (L, error)
	avlFunc          func(*Config, *http.Request, []int32) ([]int32, error)
	subsections      map[string]Subsection
}
