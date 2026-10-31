package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type handlerInput[T seeding.Lookupable, R any, A APIResource, L APIResourceList] struct {
	Endpoint         EndpointName
	resTypeSingle    ResTypeSingle
	resTypePlural    ResTypePlural
	usage            []string
	objLookup        map[string]T
	objLookupID      map[int32]T
	queryLookup      map[QueryParamName]QueryParam
	getMultipleQuery DBQueryStringMany
	retrieveQuery    DBQueryNoInput
	idToResFunc      func(*Config, handlerInput[T, R, A, L], int32) A
	resToListFunc    func(*Config, *http.Request, []A) (L, error)
	getSingleFunc    func(*http.Request, handlerInput[T, R, A, L], int32) (R, error)
	retrieveFunc     func(*http.Request, handlerInput[T, R, A, L]) ([]int32, error)
	avlFunc          func(*Config, *http.Request, []int32) ([]int32, error)
	subsections      map[SectionName]Subsection
}

type handlerInputEnums struct {
	Endpoint    EndpointName
	usage       []string
	lookup      map[string]EnumResponse
	queryLookup map[QueryParamName]QueryParam
}

type handlerInputEndpoints struct {
	Endpoint    EndpointName
	usage       []string
	slice       []EndpointName
	queryLookup map[QueryParamName]QueryParam
}

type handlerInputService[P, R any] struct {
	Endpoint  EndpointName
	usage     []string
	verifyFn  func(*Config, P, ParamsRules, map[FieldName]any) (P, error)
	executeFn func(*Config, P) (R, error)
	docsFn    func(*Config) ParamsDoc
	samplesFn func() []SampleParams[P]
}
