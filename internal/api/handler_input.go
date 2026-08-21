package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type handlerInput[T seeding.Lookupable, R any, A APIResource, L APIResourceList] struct {
	endpoint         EndpointName
	resTypeSingle    ResTypeSingle
	resTypePlural    ResTypePlural
	usage            []string
	objLookup        map[string]T
	objLookupID      map[int32]T
	queryLookup      map[QueryParamName]QueryParam
	getMultipleQuery DbQueryStringMany
	retrieveQuery    DbQueryNoInput
	idToResFunc      func(*Config, handlerInput[T, R, A, L], int32) A
	resToListFunc    func(*Config, *http.Request, []A) (L, error)
	getSingleFunc    func(*http.Request, handlerInput[T, R, A, L], int32) (R, error)
	retrieveFunc     func(*http.Request, handlerInput[T, R, A, L]) ([]int32, error)
	avlFunc          func(*Config, *http.Request, []int32) ([]int32, error)
	subsections      map[SectionName]Subsection
}

type handlerInputEnums struct {
	endpoint    EndpointName
	usage       []string
	lookup      map[string]EnumResponse
	queryLookup map[QueryParamName]QueryParam
}

type handlerInputEndpoints struct {
	endpoint    EndpointName
	usage       []string
	slice       []EndpointName
	queryLookup map[QueryParamName]QueryParam
}

type handlerInputService[P ServiceParams, R ServiceResponse] struct {
	endpoint    EndpointName
	usage       []string
	queryLookup map[QueryParamName]QueryParam
	paramsDoc   ParamsDoc
	verifyFn    func(*Config, P, map[string]any) (P, error)
	executeFn   func(*Config, P, string) (R, error)
}
